/**
 * https://gocasts.io/gocasts/simple-http-get
 */

/*
 * json response:
 *
 * ["database",["Database","Database transaction","Database index"]]
 */

package scraper

import (
    // "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

func WikiSearch(url string) []interface{} {
    r, _ := http.Get(url)
    var b []byte

    b, _ = ioutil.ReadAll(r.Body)
    r.Body.Close()

    var decoded []interface{}

    json.Unmarshal(b, &decoded)

    // searchTerm := decoded[0].(string)
    resultArray := decoded[1].([]interface{})
    // firstResult := resultArray[0].(string)

    // log.Printf("%+v", searchTerm)
    // log.Printf("%+v", firstResult)

    return resultArray
}

func WikiGrab(word string) {

}
