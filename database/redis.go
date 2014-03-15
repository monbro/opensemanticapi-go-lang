/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 */

package database

import (
    "github.com/alphazero/Go-Redis"
    "log"
)

const (
    QUEUED_PAGES = "queued_page_title"
    DONE_PAGES = "done_page_title"
)

type Database struct {
    Client redis.Client
}

func (db *Database) Init(Password string, DbNum int) {
    spec := redis.DefaultSpec().Db(DbNum).Password(Password)

    var e redis.Error
    db.Client, e = redis.NewSynchClientWithSpec(spec)

    if e != nil {
        log.Println("failed to create the client", e)
        return
    }
}

func (db *Database) AddPageToQueue(pageName string) {
    input := []byte(pageName)

    // only add page if it is not done
    isMember, e2 := db.Client.Sismember(DONE_PAGES, input)
    if e2 != nil {
        log.Println("failed to create the client", e2)
        return
    }

    if !isMember {
        _, e2 = db.Client.Sadd(QUEUED_PAGES, input)

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
    pageName, e2 := db.Client.Srandmember(QUEUED_PAGES)

    if e2 != nil {
        log.Println("failed to create the client", e2)
        panic("No Url provided!")
    }

    // remove page as well
    input := []byte(pageName)
    _, _ = db.Client.Srem(QUEUED_PAGES, input);

    // add page to be done
    _, e2 = db.Client.Sadd(DONE_PAGES, input)

    if e2 != nil {
        log.Println("failed to create the client", e2)
        panic("No Url provided!")
    }

    output := string(pageName[:])

    if output == "" {
        panic("No page in queue anymore!!!")
    }

    return output
}

func (db *Database) AddWordRelation(word string, relation string) {
    log.Printf("Added new relation '%+v'", relation)
    log.Printf("to word '%+v'", word)
    // @TODO implement
}

func (db *Database) FindWordRelations(word string) []string {
    // @TODO implement
    return []string{"test", "test1"}
}

