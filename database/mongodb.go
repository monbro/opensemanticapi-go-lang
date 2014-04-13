/**
 * provides functions to process requests to given urls
 * - http://redis.io/commands/KEYS
 * - http://redis.io/commands
 */

package database

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "github.com/golang/glog"
    "strconv"
    "math/rand"
    "time"
    // "errors"
)


type Error string

func (err Error) Error() string { return string(err) }

type MongoDb struct {
    Client *mgo.Session
    CollectionRelations *mgo.Collection
}

func (db *MongoDb) Init(Password string, DbNum int) {
    mongoUrl := "mongodb://localhost:27017/osapi"

    var err error
    db.Client, err = mgo.Dial(mongoUrl)
    if err != nil {
        glog.Errorf("failed to create the client", err)
    }

    db.CollectionRelations = db.Client.DB("osapi").C("relations")
}

func (db *MongoDb) AddPageToQueue(pageName string) {
    var result bson.M
    change := mgo.Change{
        Update:    bson.D{{"$push", bson.D{{"items", pageName}}}},
        Upsert:    true,
        ReturnNew: true,
    }
    var err error
    if _, err = db.CollectionRelations.Find(bson.M{"key": QUEUED_PAGES}).Apply(change, &result); err != nil {
        glog.Errorf("failed to create the client", err)
    }

    glog.Infof("Added page to queue: '%+v'", pageName)
}

func (db *MongoDb) AddPageToDone(pageName string) {
    var result bson.M
    change := mgo.Change{
        Update:    bson.D{{"$push", bson.D{{"items", pageName}}}},
        Upsert:    true,
        ReturnNew: true,
    }
    var err error
    if _, err = db.CollectionRelations.Find(bson.M{"key": DONE_PAGES}).Apply(change, &result); err != nil {
        glog.Errorf("failed to create the client", err)
    }

    glog.Infof("Added page to done: '%+v'", pageName)
}

func (db *MongoDb) removePageFromQueued(pageName string) {
    var result bson.M
    change := mgo.Change{
        Update:    bson.D{{"$pull", bson.D{{"items", pageName}}}},
        Upsert:    true,
        ReturnNew: true,
    }
    var err error
    if _, err = db.CollectionRelations.Find(bson.M{"key": QUEUED_PAGES}).Apply(change, &result); err != nil {
        glog.Errorf("failed to create the client", err)
    }

    glog.Infof("removed page to queue: '%+v'", pageName)
}

func (db *MongoDb) AddPagesToQueue(pagesToQueue []string) {
    for i := range pagesToQueue {
        db.AddPageToQueue(pagesToQueue[i])
    }
}

func (db *MongoDb) RandomPageFromQueue() string {
    // { $pull: { <arrayField>: <query2> } }

    // result := struct{ N int }{}
    var result bson.M
    err := db.CollectionRelations.Find(bson.M{"key": QUEUED_PAGES}).One(&result)
    // e := db.CollectionRelations.Find(bson.M{"key": QUEUED_PAGES}, bson.M{"$pop": bson.M{ "t": 1 }}).One(&result)
    if err != nil {
        glog.Errorf("failed to create the client", err)
    }

    output, err := db.Strings(result["items"], err)

    randomNumber := random(0,(len(output)))

    returnPage := output[randomNumber]
    db.removePageFromQueued(returnPage)
    db.AddPageToDone(returnPage)
    return returnPage
}

func (db *MongoDb) AddWordRelation(word string, relation string) {
    // first we do want to add the relation to the current word
    db.createWordRelation(word, relation);

    // second we want to overall count the density of every word
    db.createWordRelation(MOST_POPULAR_WORDS, relation);
}

// /**
//  * api MongoDb methods
//  */

func (db *MongoDb) GetPopularWordRelations(word string) []string {
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

func (db *MongoDb) Strings(reply interface{}, err error) ([]string, error) {
    input := reply.([]interface{})

    newArray := make([]string, len(input))
    for i, v := range input {
        newArray[i] = string(v.(string))
    }

    return newArray, err
}

func (db *MongoDb) Set(key string, value string) (interface{}, error) {
    return db.CollectionRelations.Upsert(bson.M{"key": key}, bson.M{"key": key, "value": value})
}

func (db *MongoDb) Get(key string) (interface{}, error) {
    var result bson.M
    err := db.CollectionRelations.Find(bson.M{"key": key}).One(&result)
    glog.Infof("Test: '%+v'", result["t"])
    return result["value"], err
}

// func (db *MongoDb) GetCount(key string, relation string) (interface{}, error) {
//     var result bson.M
//     err := db.CollectionRelations.Find(bson.M{"key": key}).One(&result)
//     glog.Infof("Test: '%+v'", result["value"])
//     return result["value"], err
// }

func (db *MongoDb) Members(key string) (interface{}, error) {
    var result bson.M
    err := db.CollectionRelations.Find(bson.M{"key": key}).One(&result)
    return result["items"], err
}

func (db *MongoDb) Close() {

}

func (db *MongoDb) Flushall() {
    e := db.CollectionRelations.DropCollection()
    if e != nil {
        panic("Could not delete collection")
    }
}

/**
 * private helper methods
 */

func (db *MongoDb) createWordRelation(word string, relation string) {
    var result bson.M
    change := mgo.Change{
        Update:    bson.D{{"$inc", bson.D{{"items."+relation, 1}}}},
        Upsert:    true,
        ReturnNew: true,
    }
    var err error
    if _, err = db.CollectionRelations.Find(bson.M{"key": word}).Apply(change, &result); err != nil {
        glog.Errorf("failed to create the client", err)
    }

    glog.Infof("Added word relation: '%+v'", word)
}

func (db *MongoDb) getPopularRelationsByDensity(word string, limit int) []string {

    var result bson.M
    err := db.CollectionRelations.Find(bson.M{"key": word}).One(&result)
    // e := db.CollectionRelations.Find(bson.M{"key": QUEUED_PAGES}, bson.M{"$pop": bson.M{ "t": 1 }}).One(&result)
    if err != nil {
        glog.Errorf("failed to create the client", err)
    }

    return result["items"].([]string)
}

// func (db *MongoDb) getValueFromKey(key string) string {

// }

// /**
//  * methods for statistics
//  */

// func (db *MongoDb) RaiseScrapedTextBlocksCounter() {
//     // increase counter for relation by one

// }


func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
