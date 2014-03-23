package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi/database"
    "github.com/garyburd/redigo/redis"
    "time"
    "log"
)

func TestRedis(t *testing.T) {

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.Database)
    Db.Init(pwd, dbPort)
    defer Db.Close();

    // flush database before tests
    Db.Flushall()
    log.Println("=====================================")
    log.Println("Flush Database")

    Convey("should allow to initialise and establish a connection ", t, func() {

        Convey("should save and receive a key from the database by accessing the redis client directly", func() {
            key := "testkey for test "+time.Now().String()
            keyValue := time.Now().String()

            Convey("should save the key", func() {
                _, e := Db.Client.Do("SET", key, keyValue)
                So(e, ShouldEqual, nil)

                Convey("should receive the key", func() {
                    value, e := redis.String(Db.Client.Do("GET", key))
                    So(e, ShouldEqual, nil)
                    So(keyValue, ShouldEqual, value)
                })
            })
        })

        Convey("should handle queued pages proper", func() {

            pagesToQueue := []string{"pageOne","pageTwo"}

            Convey("should add pages to the queue by a function call", func() {

                So(len(pagesToQueue), ShouldEqual, 2)

                // get a slice which will exclude the first element as we processing this one soon
                Db.AddPagesToQueue(pagesToQueue)
                // amountQueuedPages, _ := Db.Client.Smembers(database.QUEUED_PAGES)
                amountQueuedPages, _ := redis.Strings(Db.Client.Do("SMEMBERS", database.QUEUED_PAGES))

                So(len(amountQueuedPages), ShouldEqual, 2)

                Convey("should get a random page from the queued pages and remove it from the database", func() {
                    randomPage := Db.RandomPageFromQueue()

                    So(randomPage, ShouldBeIn, pagesToQueue)

                    // refresh local variable
                    amountQueuedPages, _ := redis.Strings(Db.Client.Do("SMEMBERS", database.QUEUED_PAGES))

                    So(len(amountQueuedPages), ShouldEqual, 1)

                    // amountDonePages, _ := Db.Client.Smembers(database.DONE_PAGES)
                    amountDonePages, _ := redis.Strings(Db.Client.Do("SMEMBERS", database.DONE_PAGES))
                    So(len(amountDonePages), ShouldEqual, 1)
                })
            })
        })

        // Convey("should add a new relation to a word with a counter and receive it", func() {

        //     word := "bacon"
        //     relations := []string{"cheese" ,"ham" ,"gorgonzola" ,"salami" ,"ham" ,"ham"}

        //     Convey("should add all relation", func() {
        //         // loop trough all relations
        //         for _, relation := range relations {
        //             Db.AddWordRelation(word, relation)
        //         }
        //     })

        //     Convey("should have added a coutner for ham", func() {
        //         hamCounterByte, _ := Db.Client.Get("bacon:ham")
        //         hamCounter := string(hamCounterByte[:]) // @TODO use proper type -> not string but int conversion
        //         So(hamCounter, ShouldEqual, "3")
        //     })

        //     Convey("should have added a coutner for salami", func() {
        //         hamCounterByte, _ := Db.Client.Get("bacon:salami")
        //         hamCounter := string(hamCounterByte[:]) // @TODO use proper type -> not string but int conversion
        //         So(hamCounter, ShouldEqual, "1")
        //     })

        //     Convey("should receive all distinct existing relation", func() {
        //         amountRelations := Db.GetPopularWordRelations(word)
        //         So(len(amountRelations), ShouldEqual, 4)
        //     })

        // })

    })

}
