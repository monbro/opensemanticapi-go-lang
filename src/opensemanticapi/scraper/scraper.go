/**
 * provides functions to process requests to given urls
 */

package scraper

import (
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    // "reflect"
    // "opensemanticapi/requestStruct"
)

type RequestBit struct {
    Url string
    PlainResponse string
    ResponseObjectInterface interface{}
    ResponseObjectRawJson []json.RawMessage
}

/**
 * will actually do the http get request
 */
func (rb *RequestBit) Work() {
    if(rb.Url == "") {
        panic("No Url provided!")
    }

    // if(rb.ResponseObject == nil) {
    //     panic("No Response Struct provided!")
    // }

    b := rb.GetByteCodeFromUrl()
    rb.PlainResponse = string(b[:])

    // now parse the response into the given struct object

    // fill the interface
    json.Unmarshal(b, &rb.ResponseObjectInterface)

    // fill the raw json
    json.Unmarshal(b, &rb.ResponseObjectRawJson)


    // log.Printf("CUSTOM: %+v", rb.ResponseObject)

    // typeName := reflect.TypeOf(rb.ResponseObject).String()
    // log.Printf("Type: %+v", typeName) // *requestStruct.WikiSearch

    // typeObj := reflect.TypeOf(rb.ResponseObject)
    // log.Printf("Type: %+v", typeObj) // *requestStruct.WikiSearch


    // w := reflect.New("*requestStruct.WikiSearch")
    // slice := reflect.MakeSlice(reflect.SliceOf(*rb.ResponseObject), 0, 0).Interface()
    // log.Printf("CUSTOM: %+v", reflect.TypeOf(slice))

    // w := *rb.ResponseObject.(*requestStruct.WikiSearch)
    // log.Printf("CUSTOM: %+v", w[0])
}

/**
 * will get the byte code from a url
 */
func (rb *RequestBit) GetByteCodeFromUrl() []byte{
    r, _ := http.Get(rb.Url)
    var b []byte

    b, _ = ioutil.ReadAll(r.Body)
    r.Body.Close()

    return b
}

func (rb *RequestBit) GetByteCodeFromUrlFake() []byte{
    b := []byte(`["database",["Database","Database transaction","Database index"]]`) // fake json response

    return b
}

/**
 *  will dump the url and the plain api response string
 */
func (rb *RequestBit) DebugMe() {
    log.Printf("DEBUG START ---------------------------")
    log.Printf("URL: %+v", rb.Url)
    log.Printf("RESPONSE: %+v", rb.PlainResponse)
    log.Printf("DEBUG END ---------------------------")
}

/**
 * the following should be obsolete soon
 */

/**
 * will get the byte code from a http get request
 */
func GrabJson(url string) []byte{
    r, _ := http.Get(url)
    var b []byte

    b, _ = ioutil.ReadAll(r.Body)
    r.Body.Close()

    return b;
}

/**
 * will search within wikipedia for pages matching the most for a given string,
 * then returns all results as an array
 */
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
