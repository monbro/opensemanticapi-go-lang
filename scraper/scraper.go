/**
 * provides functions to process requests to given urls
 */

package scraper

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/golang/glog"
)

/**
 * will prove a struct to hold a custom struct called 'ResponseObjectInterface' which will match to the json response
 * and in the case of a not well formated json response 'ResponseObjectRawJson' will allow the access of the raw json response string
 */
type RequestBit struct {
    Url string
    PlainResponse string
    ResponseObjectInterface interface{}
    ResponseObjectRawJson json.RawMessage
    ResponseArrayRawJson []json.RawMessage
}

/**
 * will actually do the http get request
 */
func (rb *RequestBit) Work() {
    if rb.Url == "" {
        panic("No Url provided!")
    }

    b := rb.GetByteCodeFromUrl()
    rb.PlainResponse = string(b[:])

    // fill the interface if is set
    if err := json.Unmarshal(b, &rb.ResponseObjectInterface); err != nil {
        glog.Warningf("error within json.Unmarshal for ResponseObjectInterface:", err)
    }

    // fill the raw json object
    if err := json.Unmarshal(b, &rb.ResponseObjectRawJson); err != nil {
        glog.Warningf("error within json.Unmarshal for ResponseObjectRawJson:", err)
    }

    // fill the raw json
    if err := json.Unmarshal(b, &rb.ResponseArrayRawJson); err != nil {
        glog.Warningf("error within json.Unmarshal for ResponseArrayRawJson:", err)
    }
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
