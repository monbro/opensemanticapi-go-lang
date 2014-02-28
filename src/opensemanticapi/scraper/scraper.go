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
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

func DebugJson(url string) {
    r, _ := http.Get(url)
    var b []byte

    b, _ = ioutil.ReadAll(r.Body)
    r.Body.Close()

    s := string(b[:])

    log.Printf("DEBUG START ---------------------------")
    log.Printf("URL: %+v", url)
    log.Printf("RESPONSE: %+v", s)
    log.Printf("DEBUG END ---------------------------")
}

func GrabJson(url string) []byte{
    r, _ := http.Get(url)
    var b []byte

    b, _ = ioutil.ReadAll(r.Body)
    r.Body.Close()

    return b;
}

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

func getUrl() string {
    return "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles="
}

func WikiGrab(word string) string {

    // generate the url
    url := getUrl() + word

    // grab the wikipedia api response in bytecode
    b := GrabJson(url)

    // DebugJson(url)

    // declare structs for our api respoonse
    type ActualPage struct {
        PageId int  `json:"pageid"`
        Title string `json:"title"`
    }

    type SubType struct {
        Pages map[string]ActualPage `json:"pages"`
    }

    type JsonObject struct {
        Query SubType `json:"query"`
        // QueryContinue TestObj `json:"query-continue"`
    }

    // ini a variable called 't' of our custom type JsonObject
    var t JsonObject

    // actually inject the bytecode into our variable t
    json.Unmarshal(b, &t)

    // assign the value of the property Pages of the struct SubTypes into 'me', which is a hash map
    me := t.Query.Pages

    // iterate trough all items in the hash map and pringt the key, and the properties PageId and Title
    for k, v := range me {
        log.Printf("%+v", k)
        log.Printf("%+v", v.PageId)
        log.Printf("%+v", v.Title)
    }

    return word
}
