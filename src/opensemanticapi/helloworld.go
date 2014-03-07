package main
import (
    // "flag"
    "log"
    "redis"
    "encoding/json"
    "opensemanticapi/scraper"
    "opensemanticapi/requestStruct"
)

func main() {
    searchWikipedia()
    getWikipediaPage()
    testRedis()
}

func testRedis() {
    // Parse command-line flags; needed to let flags used by Go-Redis be parsed.
    // flag.Parse()

    spec := redis.DefaultSpec().Db(13).Password("")
    client, e := redis.NewSynchClientWithSpec(spec)
    if e != nil {
        log.Println("failed to create the client", e)
        return
    }

    key := "testkey"

    input := []byte("Testinput")

    client.Set(key, input)

    value, e := client.Get(key)
    if e != nil {
        log.Println("error on Get", e)
        return
    }

    log.Printf("This was stored and fetched in Redis: %s!\n", value)
}

func searchWikipedia() {
    rb := new(scraper.RequestBit)
    rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
    rb.Work()

    result := new(requestStruct.WikiSearch)

    if err := json.Unmarshal(rb.ResponseObjectRawJson[0], &result.SearchTerm); err != nil {
        log.Fatalln("expect string:", err)
    }

    if err := json.Unmarshal(rb.ResponseObjectRawJson[1], &result.Results); err != nil {
        log.Fatalln("expect []string:", err)
    }

    log.Printf("Searchterm: %+v", result)
    log.Printf("Second result item: %+v", result.Results[1])
}

func getWikipediaPage() {
    // test to grab a page content
    rb2 := new(scraper.RequestBit)
    rb2.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=Yanqing_County"
    rb2.ResponseObjectInterface = new(requestStruct.WikiPage)
    rb2.Work()

    w2 := *rb2.ResponseObjectInterface.(*requestStruct.WikiPage)
    log.Printf("Page Response: %+v", w2)
}
