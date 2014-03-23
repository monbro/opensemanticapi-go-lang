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

const (
    QUEUED_PAGES = "queued_page_title"
    DONE_PAGES = "done_page_title"
)

type Database struct {
    Client redis.Conn
}

func (db *Database) Init(Password string, DbNum int) {
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

func (db *Database) Close() {
    db.Client.Close()
}

func (db *Database) Flushall() {
    var err error
    _, err = db.Client.Do("FLUSHALL")
    if err != nil {
        log.Println("failed to create the client", err)
    }
}

func (db *Database) AddPageToQueue(pageName string) {

    // only add page if it is not done
    wasProcessed, e2 := redis.Bool(db.Client.Do("SISMEMBER", DONE_PAGES, pageName))
    if e2 != nil {
        log.Println("failed to create the client", e2)
        return
    }

    if !wasProcessed {
        _, e2 := db.Client.Do("SADD", QUEUED_PAGES, pageName)

        if e2 != nil {
            log.Println("failed to create the client", e2)
            return
        }
        log.Printf("Added page to queue: '%+v'", pageName)
    } else {
        log.Println("Page is already member in queued pages: ", pageName)
    }
}

func (db *Database) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
    }
}

func (db *Database) RandomPageFromQueue() string {
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

func (db *Database) AddWordRelation(word string, relation string) {
    log.Printf("Added new relation '%+v'", relation)
    log.Printf("to word '%+v'", word)

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

func (db *Database) GetPopularWordRelations(word string) []string {
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

// keeping for memories or what?? :-)
// func (db *Database) GetPopularWordRelationsOld(word string) []string {
//     // @TODO implement
//     // - return a array ordered by the most popular DESC
//     // - will be limited to a certain amount within the db query

//     log.Printf("Recevive for '%+v'", word)

//     wordRelationsBytes, e := db.Client.Smembers(word)
//     if e != nil {
//         log.Println("failed to get relations for this one", e)
//         // panic("No Url provided!")
//     }

//     log.Printf("Len '%+v'", len(wordRelationsBytes))

//     // allRelations := []string{"test", "test1"}

//     var allRelations []string
//     for i := range wordRelationsBytes {
//         allRelations = append(allRelations, string(wordRelationsBytes[i][:]))
//         log.Printf("Loop '%+v'", string(wordRelationsBytes[i][:]))
//     }

//     return allRelations
// }

