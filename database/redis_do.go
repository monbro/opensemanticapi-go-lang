/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 */

package database

import (
    "github.com/garyburd/redigo/redis"
    "log"
)

type RedisDo struct {
    Client redis.Conn
}

func (db *RedisDo) Init(Password string, DbNum int) {
    var err error
    db.Client, err = redis.Dial("tcp", ":6379")
    if err != nil {
        log.Println("failed to create the client", err)
        return
    }

    var err2 error
    _, err2 = db.Client.Do("SELECT", DbNum)
    if err2 != nil {
        log.Println("failed to create the client", err2)
    }
}

func (db *RedisDo) AddPageToQueue(pageName string) {
    // only add page if it is not done
    wasProcessed, e := redis.Bool(db.Client.Do("SISMEMBER", DONE_PAGES, pageName))
    if e != nil {
        log.Println("failed to create the client", e)
    }

    if !wasProcessed {
        _, e := db.Client.Do("SADD", QUEUED_PAGES, pageName)

        if e != nil {
            log.Println("failed to create the client", e)
            return
        }
        log.Printf("Added page to queue: '%+v'", pageName)
    } else {
        log.Println("Page is already member in queued pages: ", pageName)
    }
}

func (db *RedisDo) AddPageToDone(pageName string) {
    _, e := db.Client.Do("SADD", DONE_PAGES, pageName)

    if e != nil {
        log.Println("failed to create the client", e)
        return
    }
    log.Printf("Added page to done: '%+v'", pageName)
}

func (db *RedisDo) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
    }
}

func (db *RedisDo) RandomPageFromQueue() string {
    pageName, e := redis.String(db.Client.Do("SRANDMEMBER", QUEUED_PAGES))

    if e != nil {
        log.Println("failed to create the client", e)
    }

    // remove page as well
    _, e = db.Client.Do("SREM", QUEUED_PAGES, pageName)

    // add page to be done
    _, e = db.Client.Do("SADD", DONE_PAGES, pageName)

    if e != nil {
        log.Println("failed to create the client", e)
    }

    if pageName == "" {
        panic("No page in queue anymore!!!")
    }

    return pageName
}

func (db *RedisDo) AddWordRelation(word string, relation string) {
    // first we do want to add the relation to the current word
    db.createWordRelation(word, relation);

    // second we want to overall count the density of every word
    db.createWordRelation(MOST_POPULAR_WORDS, relation);
}

/**
 * api RedisDo methods
 */

func (db *RedisDo) GetPopularWordRelations(word string) []string {
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

func (db *RedisDo) GetMostPopularWords() []string {
    return db.getPopularRelationsByDensity(MOST_POPULAR_WORDS, 120);
}

func (db *RedisDo) GetAnalysedTextBlocksCounter() string {
    return db.getValueFromKey(TEXTBLOCKS_COUNTER);
}

/**
 * public helper methods
 */

func (db *RedisDo) Set(key string, value string) (interface{}, error) {
    return db.Client.Do("SET", key, value)
}

func (db *RedisDo) Get(key string) (interface{}, error) {
    return db.Client.Do("GET", key)
}

func (db *RedisDo) Members(key string) (interface{}, error) {
    return db.Client.Do("SMEMBERS", key)
}

func (db *RedisDo) Close() {
    db.Client.Close()
}

func (db *RedisDo) Flushall() {
    var err error
    _, err = db.Client.Do("FLUSHALL")
    if err != nil {
        log.Println("failed to create the client", err)
    }
}

/**
 * private helper methods
 */

func (db *RedisDo) createWordRelation(word string, relation string) {
    // add the actual maybe related word in the db
    _, e := db.Client.Do("SADD", word, relation)

    if e != nil {
        log.Println("failed to add relations for this one", e)
    }

    // increase counter for relation by one
    _, e = db.Client.Do("INCR", word+":"+relation)

    if e != nil {
        log.Println("failed to create the client", e)
    }
}

func (db *RedisDo) getPopularRelationsByDensity(word string, limit int) []string {
    allRelations, err := redis.Strings(db.Client.Do("SORT", word,
        "BY", word+":*",
        "Limit", 0, limit,
        "DESC",
        "GET", "#"))
    if err != nil {
        panic(err)
    }

    return allRelations
}

func (db *RedisDo) getValueFromKey(key string) string {
    value, e := redis.String(db.Client.Do("GET", key))
    if e != nil {
        return "0" // @TODO thats not a proper solution, refactor this one
    }

    return value
}

/**
 * methods for statistics
 */

func (db *RedisDo) RaiseScrapedTextBlocksCounter() {
    // increase counter for relation by one
    _, e := db.Client.Do("INCR", TEXTBLOCKS_COUNTER)

    if e != nil {
        log.Println("failed to create the client", e)
    }
}
