/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 */

package database

import (
    "github.com/garyburd/redigo/redis"
    // "errors"
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

    // spec := redis.DefaultSpec().Db(DbNum).Password(Password)

    // var e redis.Error
    // db.Client, e = redis.NewSynchClientWithSpec(spec)

    // if e != nil {
    //     log.Println("failed to create the client", e)
    //     return
    // }
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

// func (db *Database) AddPageToQueue(pageName string) {
//     input := []byte(pageName)

//     // only add page if it is not done
//     isMember, e2 := db.Client.Sismember(DONE_PAGES, input)
//     if e2 != nil {
//         log.Println("failed to create the client", e2)
//         return
//     }

//     if !isMember {
//         _, e2 = db.Client.Sadd(QUEUED_PAGES, input)

//         if e2 != nil {
//             log.Println("failed to create the client", e2)
//             return
//         }
//         log.Printf("Added page to queue: '%+v'", pageName)
//     } else {
//         log.Println("Page is already member in queued pages: ", pageName)
//     }
// }

// func (db *Database) AddPagesToQueue(pagesToQueue []string) {
//     for i := range pagesToQueue {
//         db.AddPageToQueue(pagesToQueue[i])
//     }
// }

// func (db *Database) RandomPageFromQueue() string {
//     pageName, e2 := db.Client.Srandmember(QUEUED_PAGES)

//     if e2 != nil {
//         log.Println("failed to create the client", e2)
//         panic("No Url provided!")
//     }

//     // remove page as well
//     input := []byte(pageName)
//     _, _ = db.Client.Srem(QUEUED_PAGES, input);

//     // add page to be done
//     _, e2 = db.Client.Sadd(DONE_PAGES, input)

//     if e2 != nil {
//         log.Println("failed to create the client", e2)
//         panic("No Url provided!")
//     }

//     output := string(pageName[:])

//     if output == "" {
//         panic("No page in queue anymore!!!")
//     }

//     return output
// }

// func (db *Database) AddWordRelation(word string, relation string) {
//     // @TODO implement
//     // - if relations exists already putt the counter higher
//     log.Printf("Added new relation '%+v'", relation)
//     log.Printf("to word '%+v'", word)

//     input := []byte(relation)
//     _, e := db.Client.Sadd(word, input)

//     if e != nil {
//         log.Println("failed to add relations for this one", e)
//         // panic("No Url provided!")
//     }

//     // increase counter for relation by one
//     db.Client.Incrby(word+":"+relation, 1);
// }

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

// func (db *Database) GetPopularWordRelations(word string) []string {
//     db.Client.Sort(owner, "by", owner+":*", 'LIMIT', 0, 120, 'DESC', "get", "#",


//     return allRelations
// }



