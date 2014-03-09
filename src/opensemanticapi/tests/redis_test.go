package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/alphazero/Go-Redis"
    "log"
    "time"
)

func TestRedis(t *testing.T) {
    Convey("Test if we cann use the redis database here", t, func() {

        spec := redis.DefaultSpec().Db(13).Password("")
        client, e := redis.NewSynchClientWithSpec(spec)
        if e != nil {
            log.Println("failed to create the client", e)
            return
        }

        key := "testkey for test "+time.Now().String()
        input := []byte(key)

        Convey("should save the key to redis", func() {
            e := client.Set(key, input)
            So(e, ShouldEqual, nil)
        })

        Convey("should get the key to redis", func() {
            value, e := client.Get(key)
            So(e, ShouldEqual, nil)
            stringValue := string(value[:])
            So(stringValue, ShouldEqual, key)
        })

    })
}
