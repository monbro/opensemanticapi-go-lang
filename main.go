package main

import (
    "log"
    "encoding/json"
    "github.com/monbro/opensemanticapi/scraper"
    "github.com/monbro/opensemanticapi/requestStruct"
)

func main() {
    searchWikipedia()
    getWikipediaPage()
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
