package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi-go-lang/database"
    // "github.com/garyburd/redigo/redis"
    // "github.com/golang/glog"
    "time"
    "log"
)

/**
 * will test the database adapter:
 * github.com/monbro/opensemanticapi-go-lang/database/redis_do.go
 */
// func TestRedisMulti(t *testing.T) {

//     // test database configuration
//     pwd := ""
//     dbPort := 13

//     // establish a connection with the database
//     Db := new(database.RedisMulti)
//     Db.Init(pwd, dbPort)

//     testRunner(t, Db)
// }


/**
 * will test the database adapter
 * github.com/monbro/opensemanticapi-go-lang/database/redis_multi.go
 */
// func TestRedisDo(t *testing.T) {

//     // test database configuration
//     pwd := ""
//     dbPort := 13

//     // establish a connection with the database
//     Db := new(database.RedisDo)
//     Db.Init(pwd, dbPort)
//     defer Db.Close();

//     testRunner(t, Db)
// }

/**
 * will test the database adapter
 * github.com/monbro/opensemanticapi-go-lang/database/mongodb.go
 */
func TestMongoDb(t *testing.T) {

    // test database configuration
    pwd := ""
    dbPort := 0

    // establish a connection with the database
    Db := new(database.MongoDb)
    Db.Init(pwd, dbPort)
    defer Db.Close();

    testRunner(t, Db)
}

/**
 * the actual test runner to test the required functionality accross all database adapters
 */
// func testRunner(t *testing.T, Db database.DatabaseInterface) {
func testRunner(t *testing.T, Db *database.MongoDb) {

    // flush database before tests
    Db.Flushall()
    log.Println("=====================================")
    log.Println("Flush Database")

    Convey("should allow to initialise and establish a connection ", t, func() {

        Convey("should save and receive a key from the database by accessing the redis client directly", func() {
            key := "testkey for test "+time.Now().String()
            keyValue := time.Now().String()

            Convey("should save the key", func() {
                _, e := Db.Set(key, keyValue)
                So(e, ShouldEqual, nil)

                Convey("should receive the key", func() {
                    // value, e := redis.String(Db.Get(key))
                    value, e := Db.String(Db.Get(key))
                    So(e, ShouldEqual, nil)
                    So(keyValue, ShouldEqual, value)
                })
            })
        })

        Convey("should handle queued pages proper", func() {

            pagesToQueue := []string{"pageOne", "Geographic Names Information System"}

            Convey("should add pages to the queue by a function call", func() {

                So(len(pagesToQueue), ShouldEqual, 2)

                // get a slice which will exclude the first element as we processing this one soon
                Db.AddPagesToQueue(pagesToQueue)
                members, errr := Db.Members(database.QUEUED_PAGES)
                // glog.Infof("Got random item from queued pages.",membersNew[0])
                queuedPages, _ := Db.Strings(members, errr)

                So(len(queuedPages), ShouldEqual, 2)

                Convey("should get a random page from the queued pages and remove it from the database", func() {
                    randomPage := Db.RandomPageFromQueue()
                    // glog.Infof("failed to create the client", randomPage)

                    So(randomPage, ShouldBeIn, pagesToQueue)

                    // refresh local variable
                    queuedPages, _ := Db.Strings(Db.Members(database.QUEUED_PAGES))

                    So(len(queuedPages), ShouldEqual, 1)

                    donePages, _ := Db.Strings(Db.Members(database.DONE_PAGES))
                    So(len(donePages), ShouldEqual, 1)
                })
            })
        })

        Convey("should add a new relation to a word with a counter and receive it", func() {

            word := "bacon"
            relations := []string{"cheese" ,"ham" ,"gorgonzola" ,"salami" ,"ham" ,"ham", "cheese"}

            Convey("should add all relation", func() {
                // loop trough all relations
                for _, relation := range relations {
                    Db.AddWordRelation(word, relation)
                }
            })

            // Convey("should have added a coutner for ham", func() {
            //     hamCounter, _ := Db.String(Db.GetCount("bacon","ham"))
            //     So(hamCounter, ShouldEqual, 3)
            // })

            // Convey("should have added a coutner for salami", func() {
            //     hamCounter, _ := Db.String(Db.Get("bacon:salami"))
            //     So(hamCounter, ShouldEqual, 1)
            // })

            // Convey("should receive all distinct existing relation", func() {
            //     amountRelations := Db.GetPopularWordRelations(word)
            //     So(len(amountRelations), ShouldEqual, 4) // because we added 5 distinct keys
            //     So(amountRelations[0], ShouldEqual, "ham") // because it was added 3 times
            //     So(amountRelations[1], ShouldEqual, "cheese") // because it was added 2 times
            // })

        })

    })
}
