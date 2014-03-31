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
    "log"
)

type RedisMulti struct {
    Client redis.Conn
    Pool *redis.Pool
}

func (db *RedisMulti) Init(Password string, DbNum int) {
    var err error
    db.Client, err = redis.Dial("tcp", ":6379")
    if err != nil {
        log.Println("failed to create the client", err)
        return
    }

    db.Client.Do("SELECT", DbNum)
}

// https://stackoverflow.com/questions/19971968/go-golang-redis-too-many-open-files-error
func (db *RedisMulti) InitPool(Password string, DbNum int) {
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

func (db *RedisMulti) Close() {
    db.Client.Close()
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

func (db *RedisMulti) AddPageToQueue(pageName string) {
    r := db.Pool.Get()
    defer r.Close()

    // only add page if it is not done
    wasProcessed, e := redis.Bool(r.Do("SISMEMBER", DONE_PAGES, pageName))
    if e != nil {
        log.Println("failed to create the client", e)
    }

    if !wasProcessed {
        r.Send("SADD", QUEUED_PAGES, pageName)
        log.Printf("Added page to queue: '%+v'", pageName)
    } else {
        log.Println("Page is already member in queued pages: ", pageName)
    }
}

func (db *RedisMulti) AddPageToDone(pageName string) {
    r := db.Pool.Get()
    defer r.Close()

    r.Send("SADD", DONE_PAGES, pageName)
    log.Printf("Added page to done: '%+v'", pageName)
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
        log.Println("failed to create the client", e)
    }

    // remove page as well
    r.Send("SREM", QUEUED_PAGES, pageName)

    // add page to be done
    r.Send("SADD", DONE_PAGES, pageName)

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
    return db.getPopularRelationsByDensity(word);
}

func (db *RedisMulti) GetMostPopularWords() []string {
    return db.getPopularRelationsByDensity(MOST_POPULAR_WORDS);
}

func (db *RedisMulti) GetAnalysedTextBlocksCounter() string {
    return db.getValueFromKey(TEXTBLOCKS_COUNTER);
}

/**
 * private helper functions
 */

func (db *RedisMulti) stripOverlappingListContent(contextList []string, stripList []string) {
    // check the words that are
    // maybe try http://redis.io/commands/sdiff
}

func (db *RedisMulti) createWordRelation(word string, relation string) {
    r := db.Pool.Get()
    defer r.Close()

    // add the actual maybe related word in the db
    r.Send("SADD", word, relation)

    // increase counter for relation by one
    r.Send("INCR", word+":"+relation)
}

func (db *RedisMulti) getPopularRelationsByDensity(word string) []string {
    r := db.Pool.Get()
    defer r.Close()

    allRelations, err := redis.Strings(r.Do("SORT", word,
        "BY", word+":*",
        "Limit", 0, 120,
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

/**
 * functions for statistics
 */

func (db *RedisMulti) RaiseScrapedTextBlocksCounter() {
    r := db.Pool.Get()
    defer r.Close()

    // increase counter for relation by one
    r.Send("INCR", TEXTBLOCKS_COUNTER)
}
