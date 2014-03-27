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

type RedisMulti struct {
    Client redis.Conn
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

func (db *RedisMulti) Close() {
    db.Client.Close()
}

func (db *RedisMulti) Flushall() {
    db.Client.Do("FLUSHALL")
}

func (db *RedisMulti) Flush() {
    db.Client.Flush()
}

func (db *RedisMulti) AddPageToQueue(pageName string) {
    // only add page if it is not done
    wasProcessed, e := redis.Bool(db.Client.Do("SISMEMBER", DONE_PAGES, pageName))
    if e != nil {
        log.Println("failed to create the client", e)
    }

    if !wasProcessed {
        db.Client.Send("SADD", QUEUED_PAGES, pageName)
        log.Printf("Added page to queue: '%+v'", pageName)
    } else {
        log.Println("Page is already member in queued pages: ", pageName)
    }
}

func (db *RedisMulti) AddPageToDone(pageName string) {
    db.Client.Send("SADD", DONE_PAGES, pageName)
    log.Printf("Added page to done: '%+v'", pageName)
}

func (db *RedisMulti) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
    }
}

func (db *RedisMulti) RandomPageFromQueue() string {
    pageName, e := redis.String(db.Client.Do("SRANDMEMBER", QUEUED_PAGES))

    if e != nil {
        log.Println("failed to create the client", e)
    }

    // remove page as well
    db.Client.Send("SREM", QUEUED_PAGES, pageName)

    // add page to be done
    db.Client.Send("SADD", DONE_PAGES, pageName)

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

func (db *RedisMulti) createWordRelation(word string, relation string) {
    // add the actual maybe related word in the db
    db.Client.Send("SADD", word, relation)

    // increase counter for relation by one
    db.Client.Send("INCR", word+":"+relation)
}

func (db *RedisMulti) getPopularRelationsByDensity(word string) []string {
    allRelations, err := redis.Strings(db.Client.Do("SORT", word,
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
    value, e := redis.String(db.Client.Do("GET", key))
    if e != nil {
        return "0" // @TODO thats not a proper solution, refactor this one
    }

    return value
}

/**
 * functions for statistics
 */

func (db *RedisMulti) RaiseScrapedTextBlocksCounter() {
    // increase counter for relation by one
    db.Client.Send("INCR", TEXTBLOCKS_COUNTER)
}
