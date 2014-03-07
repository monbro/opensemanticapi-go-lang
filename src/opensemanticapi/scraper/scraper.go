/**
 * provides functions to process requests to given urls
 */

package scraper

import (
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
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
    if rb.Url == "" {
        panic("No Url provided!")
    }

    b := rb.GetByteCodeFromUrl()
    rb.PlainResponse = string(b[:])

    // fill the interface if is set
    if err := json.Unmarshal(b, &rb.ResponseObjectInterface); err != nil {
        log.Printf("error within json.Unmarshal for ResponseObjectInterface:", err)
    }

    // fill the raw json
    if err := json.Unmarshal(b, &rb.ResponseObjectRawJson); err != nil {
        log.Printf("error within json.Unmarshal for ResponseObjectRawJson:", err)
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

/**
 * hopefully used later as a mock or something
 */
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
