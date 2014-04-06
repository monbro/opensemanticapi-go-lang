package tests

import(
    "testing"
    // . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi-go-lang/database"
    // "github.com/monbro/opensemanticapi-go-lang/analyse"
    analyseAdapter "github.com/monbro/opensemanticapi-go-lang/analyse/adapter"
    "log"
    "strconv"
)

var searchTerm string = "garden"

func BenchmarkDefaultWorkerFastMode(b *testing.B) {
    log.Println("Testing now Worker which is using go routines to process a text block and to store it into redis! with fastmode on")

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.RedisMulti)
    Db.InitPool(pwd, dbPort)
    Db.Flushall()

    worker := new(analyseAdapter.ConcurrencyA)
    worker.Db = Db
    worker.SnippetLength = 120

    worker.IsFastMode = true
    worker.IsInfiniteWorking = false

    // reset the brenchmark timer
    b.ResetTimer()

    worker.RunNext(searchTerm)

    r := Db.Pool.Get()
    defer r.Close()

    amount, _ := r.Do("SCARD", searchTerm)
    log.Println("Items have been added to db: "+strconv.FormatInt(amount.(int64), 10))
}

func BenchmarkDefaultWorker(b *testing.B) {
    log.Println("Testing now Worker which is using go routines to process a text block and to store it into redis!")

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.RedisMulti)
    Db.InitPool(pwd, dbPort)
    Db.Flushall()

    worker := new(analyseAdapter.ConcurrencyA)
    worker.Db = Db
    worker.SnippetLength = 120

    worker.IsFastMode = false
    worker.IsInfiniteWorking = false

    // reset the brenchmark timer
    b.ResetTimer()

    worker.RunNext(searchTerm)

    r := Db.Pool.Get()
    defer r.Close()

    amount, _ := r.Do("SCARD", searchTerm)
    log.Println("Items have been added to db: "+strconv.FormatInt(amount.(int64), 10))
}

func BenchmarkWorkerSerial(b *testing.B) {
    log.Println("Testing now WorkerSerial which is not using any go routines!")

    // test database configuration
    pwd := ""
    dbPort := 13

    // establish a connection with the database
    Db := new(database.RedisMulti)
    Db.InitPool(pwd, dbPort)
    Db.Flushall()

    worker := new(analyseAdapter.SerialA)
    worker.Db = Db
    worker.SnippetLength = 120

    worker.InfiniteWorking = false

    // reset the brenchmark timer
    b.ResetTimer()

    // run the actual task to test against
    worker.RunNext(searchTerm)

    r := Db.Pool.Get()
    defer r.Close()

    amount, _ := r.Do("SCARD", searchTerm)
    log.Println("Items have been added to db: "+strconv.FormatInt(amount.(int64), 10))
}
