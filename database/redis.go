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
    key := "queued_page_title"
    input := []byte(pageName)

    _, e2 := db.Client.Sadd(key, input)

    if e2 != nil {
        log.Println("failed to create the client", e2)
        return
    }
}

func (db *Database) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
        log.Printf("Added page to queue: '%+v'",pagesToQueue[i])
    }
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
