/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 */

package database

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "strconv"
    "log"
)

type MongoDb struct {
    Client *mgo.Session
    Collection *mgo.Collection
}

func (db *MongoDb) Init(Password string, DbNum int) {
    mongoUrl := "mongodb://localhost:27017/osapi"

    var err error
    db.Client, err = mgo.Dial(mongoUrl)
    if err != nil {
        log.Println("failed to create the client", err)
    }

    db.Collection = db.Client.DB("osapi").C("testcollection")
    iter := db.Collection.Find(bson.M{"_id": "123"})

    // if err2 != nil {
    //     log.Println("failed to create the client", err2)
    // }
    log.Println("the result: ", iter)
}

// func (db *MongoDb) AddPageToQueue(pageName string) {

// }

// func (db *MongoDb) AddPageToDone(pageName string) {

// }

// func (db *MongoDb) AddPagesToQueue(pagesToQueue []string) {

// }

// func (db *MongoDb) RandomPageFromQueue() string {

// }

// func (db *MongoDb) AddWordRelation(word string, relation string) {

// }

// /**
//  * api MongoDb methods
//  */

// func (db *MongoDb) GetPopularWordRelations(word string) []string {

// }

// func (db *MongoDb) GetMostPopularWords() []string {
// }

// func (db *MongoDb) GetAnalysedTextBlocksCounter() string {
// }

/**
 * public helper methods
 */

func (db *MongoDb) String(key interface{}, e error) (string, error) {
    return key.(string), e
}

func (db *MongoDb) Int(key string, e error) (int, error) {
    i, err := strconv.Atoi(key)
    if err != nil {
        // handle error
        panic("Could not convert string to int")
    }
    return i, e
}

func (db *MongoDb) Set(key string, value string) (interface{}, error) {
    return db.Collection.Upsert(bson.M{"key": key}, bson.M{"key": key, "value": value})
}

func (db *MongoDb) Get(key string) (interface{}, error) {
    var result bson.M
    err := db.Collection.Find(bson.M{"key": key}).One(&result)
    return result["value"], err
}

// func (db *MongoDb) Members(key string) (interface{}, error) {
// }

func (db *MongoDb) Close() {

}

func (db *MongoDb) Flushall() {

}

/**
 * private helper methods
 */

// func (db *MongoDb) createWordRelation(word string, relation string) {

// }

// func (db *MongoDb) getPopularRelationsByDensity(word string, limit int) []string {

// }

// func (db *MongoDb) getValueFromKey(key string) string {

// }

// /**
//  * methods for statistics
//  */

// func (db *MongoDb) RaiseScrapedTextBlocksCounter() {
//     // increase counter for relation by one

// }
