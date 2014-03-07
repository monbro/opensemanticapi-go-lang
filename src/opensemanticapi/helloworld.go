package main
import (
    // "flag"
    // "fmt"
    "log"
    // "redis"
    "encoding/json"
    "opensemanticapi/scraper"
    "opensemanticapi/requestStruct"
)

// func mainOld() {
//     // Redis test
//     // Parse command-line flags; needed to let flags used by Go-Redis be parsed.
//     flag.Parse()

//     spec := redis.DefaultSpec().Db(13).Password("")
//     client, e := redis.NewSynchClientWithSpec(spec)
//     if e != nil {
//         log.Println("failed to create the client", e)
//         return
//     }

//     key := "testkey"

//     input := []byte("Testinput")

//     client.Set(key, input)

//     value, e := client.Get(key)
//     if e != nil {
//         log.Println("error on Get", e)
//         return
//     }

//     fmt.Printf("Hey, ciao %s!\n", fmt.Sprintf("%s", value))



//     // a) this is using the scraper section

//     // catch a suggested list of results for a random keyword
//     url := "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
//     val := scraper.WikiSearch(url)

//     fmt.Printf("%+v \n", val[1])
//     // res := scraper.WikiGrab(val[1].(string))
//     scraper.WikiGrab("Yanqing_County")



//     // b) another test

//     rb := new(scraper.RequestBit)
//     rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
//     rb.ResponseObject = new(requestStruct.WikiSearch)
//     rb.Work()

//     w := *rb.ResponseObject.(*requestStruct.WikiSearch)
//     // or do a type switch http://golang.org/doc/effective_go.html#type_switch ??
//     // http://grokbase.com/t/gg/golang-nuts/14173g1p35/go-nuts-empty-interface-and-type-switch

//     log.Printf("CUSTOM: %+v", w[0])
// }

func main() {
    rb := new(scraper.RequestBit)
    rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
    rb.Work()

    // w := *rb.ResponseObjectRawJson.(*json.RawMessage)
    log.Printf("CUSTOM: %+v", rb.ResponseObjectRawJson)
    // log.Printf("CUSTOM: %+v", w[1][2])

    result := new(requestStruct.WikiSearchTwo)

    if err := json.Unmarshal(rb.ResponseObjectRawJson[0], &result.SearchTerm); err != nil {
        log.Fatalln("expect string:", err)
    }

    if err := json.Unmarshal(rb.ResponseObjectRawJson[1], &result.Results); err != nil {
        log.Fatalln("expect []string:", err)
    }

    log.Printf("CUSTOM: %+v", result)
    log.Printf("CUSTOM: %+v", result.Results[1])
    // log.Printf("Plain Response: %s\n", b)


    // test to grab a page content
    rb2 := new(scraper.RequestBit)
    rb2.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=Yanqing_County"
    rb2.ResponseObjectInterface = new(requestStruct.WikiPage)
    rb2.Work()

    w2 := *rb2.ResponseObjectInterface.(*requestStruct.WikiPage)
    log.Printf("CUSTOM: %+v", w2)
}
