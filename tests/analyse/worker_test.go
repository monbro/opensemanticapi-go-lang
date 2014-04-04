package tests

import(
    "testing"
    // . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi-go-lang/database"
    "github.com/monbro/opensemanticapi-go-lang/analyse"
    // "github.com/garyburd/redigo/redis"
    "log"
    "strconv"
)

var searchTerm string = "garden"

func BenchmarkDefaultWorker(b *testing.B) {
    log.Println("Testing now Worker which is using go routines to process a text block and to store it into redis!")

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.RedisMulti)
    Db.InitPool(pwd, dbPort)
    Db.Flushall()

    worker := new(analyse.Worker)
    worker.Db = Db
    worker.SNIPPET_LENGTH = 120

    worker.IsFastMode = false
    worker.IsInfiniteWorking = false

    worker.RunNext(searchTerm)

    r := Db.Pool.Get()
    defer r.Close()

    amount, _ := r.Do("SCARD", searchTerm)
    log.Println("Items have been added to db: "+strconv.FormatInt(amount.(int64), 10))
}

func BenchmarkWorkerA(b *testing.B) {
    log.Println("Testing now WorkerA which is not using any go routines!")

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.RedisMulti)
    Db.InitPool(pwd, dbPort)
    Db.Flushall()

    worker := new(analyse.WorkerA)
    worker.Db = Db
    worker.SNIPPET_LENGTH = 120

    worker.InfiniteWorking = false

    // we do not have a heavy setup, but just in case we reset the timer
    b.ResetTimer()

    // run the actual task to test against
    worker.RunNext(searchTerm)

    r := Db.Pool.Get()
    defer r.Close()

    amount, _ := r.Do("SCARD", searchTerm)
    log.Println("Items have been added to db: "+strconv.FormatInt(amount.(int64), 10))
}
