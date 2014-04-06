/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 *
 * https://github.com/luanjunyi/gossipd/blob/master/mqtt/redis.go would be a good abstraction class to use
 * also interesting https://github.com/vube/redigolock/blob/master/redigolock.go
 */

package database

import (
    "github.com/garyburd/redigo/redis"
    "github.com/golang/glog"
)

type RedisMulti struct {
    Pool *redis.Pool
}

// https://stackoverflow.com/questions/19971968/go-golang-redis-too-many-open-files-error
func (db *RedisMulti) Init(Password string, DbNum int) {
    db.Pool = &redis.Pool {
        MaxIdle: 6000,
        MaxActive: 60000, // max number of connections
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", ":6379")
            if err != nil {
                    panic(err.Error())
            }
            c.Do("SELECT", DbNum)
            return c, err
        },
    }
}

func (db *RedisMulti) AddPageToQueue(pageName string) {
    r := db.Pool.Get()
    defer r.Close()

    // only add page if it is not done
    wasProcessed, e := redis.Bool(r.Do("SISMEMBER", DONE_PAGES, pageName))
    if e != nil {
        glog.Infof("failed to create the client", e)
    }

    if !wasProcessed {
        r.Send("SADD", QUEUED_PAGES, pageName)
        glog.Infof("Added page to queue: '%+v'", pageName)
    } else {
        glog.Infof("Page is already member in queued pages: ", pageName)
    }
}

func (db *RedisMulti) AddPageToDone(pageName string) {
    r := db.Pool.Get()
    defer r.Close()

    r.Send("SADD", DONE_PAGES, pageName)
    glog.Infof("Added page to done: '%+v'", pageName)
}

func (db *RedisMulti) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
    }
}

func (db *RedisMulti) RandomPageFromQueue() string {
    r := db.Pool.Get()
    defer r.Close()

    pageName, e := redis.String(r.Do("SRANDMEMBER", QUEUED_PAGES))

    if e != nil {
        glog.Errorf("failed to create the client", e)
    }

    // remove page as well
    _, e1 := r.Do("SREM", QUEUED_PAGES, pageName)
    if e1 != nil {
        glog.Errorf("failed to create the client", e)
    }

    // add page to be done
    _, e2 :=r.Do("SADD", DONE_PAGES, pageName)
    if e2 != nil {
        glog.Errorf("failed to create the client", e)
    }

    if pageName == "" {
        panic("No page in queue anymore!!!")
    }

    return pageName
}

func (db *RedisMulti) AddWordRelation(word string, relation string) {
    // first we do want to add the relation to the current word
    db.createWordRelation(word, relation);

    // second we want to overall count the density of every word
    db.createWordRelation(MOST_POPULAR_WORDS, relation);
}

/**
 * api RedisMulti functions
 */

func (db *RedisMulti) GetPopularWordRelations(word string) []string {
    mongoDBWords := db.getPopularRelationsByDensity(word, 120)
    popularWords := db.getPopularRelationsByDensity(MOST_POPULAR_WORDS, 500);

    v := make([]string, 0, len(mongoDBWords))

    for  _, value := range mongoDBWords {
        if !stringInSlice(value, popularWords) {
            v = append(v, value)
        }
    }
    return v
}

func (db *RedisMulti) GetMostPopularWords() []string {
    return db.getPopularRelationsByDensity(MOST_POPULAR_WORDS, 120)
}

func (db *RedisMulti) GetAnalysedTextBlocksCounter() string {
    return db.getValueFromKey(TEXTBLOCKS_COUNTER)
}

/**
 * public helper methods
 */

func (db *RedisMulti) Set(key string, value string) (interface{}, error) {
    r := db.Pool.Get()
    defer r.Close()

    return r.Do("SET", key, value)
}

func (db *RedisMulti) Get(key string) (interface{}, error) {
    r := db.Pool.Get()
    defer r.Close()

    return r.Do("GET", key)
}

func (db *RedisMulti) Members(key string) (interface{}, error) {
    r := db.Pool.Get()
    defer r.Close()

    return r.Do("SMEMBERS", key)
}

func (db *RedisMulti) Close() {
    db.Pool.Close()
}

func (db *RedisMulti) Flushall() {
    r := db.Pool.Get()
    defer r.Close()

    r.Do("FLUSHALL")
}

func (db *RedisMulti) Flush() {
    r := db.Pool.Get()
    defer r.Close()

    r.Flush()
}

func (db *RedisMulti) Multi() {
    r := db.Pool.Get()
    defer r.Close()

    r.Send("MULTI")
}

/**
 * private helper methods
 */

func (db *RedisMulti) createWordRelation(word string, relation string) {
    r := db.Pool.Get()
    defer r.Close()

    // add the actual maybe related word in the db
    r.Send("SADD", word, relation)

    // increase counter for relation by one
    r.Send("INCR", word+":"+relation)
}

func (db *RedisMulti) getPopularRelationsByDensity(word string, limit int) []string {
    r := db.Pool.Get()
    defer r.Close()

    allRelations, err := redis.Strings(r.Do("SORT", word,
        "BY", word+":*",
        "Limit", 0, limit,
        "DESC",
        "GET", "#"))
    if err != nil {
        panic(err)
    }

    return allRelations
}

func (db *RedisMulti) getValueFromKey(key string) string {
    r := db.Pool.Get()
    defer r.Close()

    value, e := redis.String(r.Do("GET", key))
    if e != nil {
        return "0" // @TODO thats not a proper solution, refactor this one
    }

    return value
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

/**
 * methods for statistics
 */

func (db *RedisMulti) RaiseScrapedTextBlocksCounter() {
    r := db.Pool.Get()
    defer r.Close()

    // increase counter for relation by one
    r.Send("INCR", TEXTBLOCKS_COUNTER)
}
