package main

import (
    "github.com/monbro/opensemanticapi/analyse"
)

/**
 * configurations constantes
 */
const (
    SNIPPET_LENGTH = 120
    START_SEARCH_TERM = "database"
)

func main() {
    worker := new(analyse.Worker)
    worker.START_SEARCH_TERM = START_SEARCH_TERM
    worker.SNIPPET_LENGTH = SNIPPET_LENGTH
    worker.Run()
}


// func mainOld() {
//     // declare needed variables
//     var pages []string

//     db := new(database.Database)
//     db.Init("", 13)

//     // initial search request to get some pages back
//     pages = searchWikipedia("database")
//     searchTerm := pages[0]
//     log.Printf("Searchterm: %+v", searchTerm)

//     // get a slice which will exclude the first element as we processing this one soon
//     pagesToQueue := pages[1:]

//     // add all other pages to the queue
//     for i := range pagesToQueue {
//         db.AddPageToQueue(pagesToQueue[i])
//         log.Printf("Added page to queue: '%+v'",pagesToQueue[i])
//     }

//     // process first page now
//     rawContent := getWikipediaPage(pages[1])
//     // log.Printf("Page Response Processed: %+v", requestStruct.GetWikiRawTextRegexpr(rawContent))
//     // log.Printf("Page Response Processed: %+v", rawContent)

//     re := regexp.MustCompile("\n|\r")
//     snippets := re.Split(rawContent, -1)

//     log.Printf("Length snippets: %+v", len(snippets))
//     // log.Printf("Length snippets: %+v", snippets[10])

//     for i := range snippets {
//         if SNIPPET_LENGTH < len(snippets[i]) {
//             log.Printf("LENG SNIPPET: '%+v'",len(snippets[i]))
//             log.Printf("Snippets raw content: %+v", requestStruct.GetWikiRawTextRegexpr(snippets[i]))
//             log.Printf("==========================================================================================")
//         }
//     }

//     // @TODO: strip also content between a block

// }
