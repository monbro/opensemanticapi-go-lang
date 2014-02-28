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
    "github.com/jmoiron/jsonq"
    "strings"
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

type jsonobject struct {
    Query ObjectType `json:"query"`
    QueryContinue TestObj `json:"query-continue"`
}

type TestObj struct {
    Categories Cat `json:"categories"`
}

type Cat struct {
    Value string `json:"clcontinue"`
}

// .

type ObjectType struct {
    Pages map[string]interface{} `json:"pages"`
}

type PageObject struct {
    Number Page `json:"233953"`
}

type Page struct {
    Pageid int `json:"pageid"`
    Title string  `json:"title"`
}

/*
 * TODO:
 * - add more inline comments
 */
func WikiGrab(word string) string {

    url := "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=" + word

    // DebugJson(url)

    b := GrabJson(url)

    var jsontype jsonobject
    json.Unmarshal(b, &jsontype)
    // log.Printf("%+v", jsontype)

    log.Printf("%+v", jsontype.Query)

    // for k, _ := range jsontype.Query.Pages {
    //     log.Printf("%+v", k)
    // }

    // searchTerm := decoded[0].(string)
    // resultArray := decoded.([]interface{})
    // resultArray := decoded.([]interface{})

    // log.Printf("%+v", resultArray)

    return word
}

func WikiGrabViaJsonq(word string) string {

    url := "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=" + word

    b := GrabJson(url)

    s := string(b[:])

    data := map[string]interface{}{}
    dec := json.NewDecoder(strings.NewReader(s))
    dec.Decode(&data)
    jq := jsonq.NewQuery(data)

    log.Printf("%+v", jq)

    return word
}

// func WikiGrab(word string) string {

//     url := "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=" + word

//     type UserType struct {
//         UserTypeId int
//         UserTypeName string
//     }

//     type User struct {
//         Session   string   `jpath:"userContext.cobrandConversationCredentials.sessionToken"`
//         CobrandId int      `jpath:"userContext.cobrandId"`
//         UserType  UserType `jpath:"userType"`
//         LoginName string   `jpath:"loginName"`
//     }

//     docScript := []byte(document)
//     docMap := map[string]interface{}{}
//     json.Unmarshal(docScript, &docMap)

//     user := User{}
//     DecodePath(docMap, &user)

//     fmt.Printf("%#v", user)

//     return word
// }
