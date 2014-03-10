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
    Password string
    DbNum int
    Client redis.Client
}

func (db *Database) AddPageToQueue(pageName string) {
    spec := redis.DefaultSpec().Db(db.DbNum).Password(db.Password)

    var e redis.Error
    db.Client, e = redis.NewSynchClientWithSpec(spec)

    if e != nil {
        log.Println("failed to create the client", e)
        return
    }

    key := "queued_page_title"
    input := []byte(pageName)

    _, e2 := db.Client.Sadd(key, input)

    if e2 != nil {
        log.Println("failed to create the client", e2)
        return
    }
}
