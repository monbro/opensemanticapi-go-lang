package main

import (
    "log"
    "net/url"
    "encoding/json"
    "github.com/monbro/opensemanticapi/scraper"
    "github.com/monbro/opensemanticapi/requestStruct"
    "github.com/monbro/opensemanticapi/database"
)

func main() {
    var pages []string

    pages = searchWikipedia("database")
    log.Printf("Searchterm: %+v", pages[2])

    // works but not used yet
    // rawContent := getWikipediaPage(pages[1])
    // log.Printf("Page Response: %+v", rawContent)

    db := new(database.Database)
    db.Password = ""
    db.DbNum = 13
    db.AddPageToQueue(pages[2])
}

func searchWikipedia(searchTerm string) []string {
    rb := new(scraper.RequestBit)
    rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search="+url.QueryEscape(searchTerm)+"&format=json&limit=3"
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

    return result.Results
}

func getWikipediaPage(firstPage string) string {
    // test to grab a page content
    rb2 := new(scraper.RequestBit)
    rb2.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles="+url.QueryEscape(firstPage)

    rb2.ResponseObjectInterface = new(requestStruct.WikiPage)
    rb2.Work()

    w2 := *rb2.ResponseObjectInterface.(*requestStruct.WikiPage)

    for _, value := range w2.Query.Pages {
        return value.Rev[0].RawContent
    }

    return ""
}
