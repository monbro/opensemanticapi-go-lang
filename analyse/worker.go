/**
 * provides functions to create wording context relations
 */

package analyse

import (
    "log"
    "net/url"
    "encoding/json"
    "github.com/monbro/opensemanticapi/scraper"
    "github.com/monbro/opensemanticapi/requestStruct"
    "github.com/monbro/opensemanticapi/database"
)

type Worker struct {
    Debug bool
    Db *database.Database
    START_SEARCH_TERM string
    SNIPPET_LENGTH int
}

/**
 * will initially start the process
 */
func (w *Worker) Run() {
    // initial start
    w.RunNext(w.START_SEARCH_TERM)
}

/**
 * will run the process of storing words that are related in its context
 */
func (w *Worker) RunNext(searchTerm string) {

    log.Printf("Searchterm Now: '%+v'", searchTerm)

    // declare needed variables
    var pages []string

    w.Db = new(database.Database)
    w.Db.Init("", 13)

    // initial search request to get some pages back
    pages = SearchWikipedia(searchTerm)
    searchTerm = pages[0]

    // get a slice which will exclude the first element as we processing this one soon
    pagesToQueue := pages[1:]
    w.Db.AddPagesToQueue(pagesToQueue)

    rawContent := GetWikipediaPage(pages[0])

    snippetsRaw := GetSnippetsFromText(rawContent)
    snippets := GetSnippetsFromText(rawContent)
    snippets = CleanUpSnippets(snippets)

    for i := range snippets {
        if w.SNIPPET_LENGTH < len(snippets[i]) {
            log.Printf("LENG SNIPPET: '%+v'",len(snippets[i]))
            log.Printf("Snippet raw: %+v", len(snippetsRaw[i]))
            log.Printf("Snippet cleaned: %+v", len(snippets[i]))
            log.Printf("==========================================================================================")

            // w.CreateSnippetWordsReplations(snippets[i])
        }
    }

    // create aloop by calling it self for the next search term
    w.RunNext(w.Db.RandomPageFromQueue())
}

/**
 * will search wikipedia for a search term and return existing matching pages
 */
func SearchWikipedia(searchTerm string) []string {
    // lets create a new http request object
    rb := new(scraper.RequestBit)

    rb.Url = "https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch="+url.QueryEscape(searchTerm)+"&format=json&limit=3"
    log.Printf("Url crawling in SearchWikipedia: %+v", rb.Url)

    // inject the struct for the json response
    rb.ResponseObjectInterface = new(requestStruct.WikiSearch)
    rb.Work() // fire the request

    w2 := *rb.ResponseObjectInterface.(*requestStruct.WikiSearch)

    var results []string

    for _, value := range w2.Query.Search {
        results[] = value.Rev[0].RawContent
    }

    return wikiOpenSearch.Results
}

/**
 * will get the content of a wikipedia page
 */
func GetWikipediaPage(firstPage string) string {
    // lets create a new http request object
    rb := new(scraper.RequestBit)
    rb.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles="+url.QueryEscape(firstPage)
    log.Printf("Url crawling in GetWikipediaPage: %+v", rb.Url)

    // inject the struct for the json response
    rb.ResponseObjectInterface = new(requestStruct.WikiPage)
    rb.Work() // fire the request

    // for type assertion we need to explicite set the type of the returned interface object again
    // @TODO it would be nice to work without an interface here at all, but on the other hand to be flexible on the struct
    w2 := *rb.ResponseObjectInterface.(*requestStruct.WikiPage)

    // as the attribute 'Pages' is a map we neet to iterate trough it and return the first result, assuming this one is the page content
    for _, value := range w2.Query.Pages {
        return value.Rev[0].RawContent
    }

    // otherwise return an empty string
    return ""
}

/**
 * will analyse a snippet by spinning relations between each word within this snippet
 */
func (w *Worker) CreateSnippetWordsReplations(snippet string) {
    words := GetWordsFromSnippet(snippet)
    for _, word := range words {
        if len(word) > 3 &&
            word != " " {
            for _, relation := range words {
                if word != relation &&
                    len(relation) > 3 &&
                    relation != " " {
                    w.Db.AddWordRelation(word, relation)
                }
            }
        }

    }
}

/**
 * will search wikipedia for a search term and return existing matching pages
 * will be outdated soon!
 *
 * NOT IN USE CURRENTLY
 */
func OpenSearchWikipedia(searchTerm string) []string {
    // lets create a new http request object
    rb := new(scraper.RequestBit)

    rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search="+url.QueryEscape(searchTerm)+"&format=json&limit=3"
    log.Printf("Url crawling in SearchWikipedia: %+v", rb.Url)
    rb.Work() // fire the request

    // as wikipedia returns a sh*t formatted json we need to assign the result in two steps
    wikiOpenSearch := new(requestStruct.WikiOpenSearch)

    // first step is to assign the result term which is the first item in the returned array
    if err := json.Unmarshal(rb.ResponseObjectRawJson[0], &wikiOpenSearch.SearchTerm); err != nil {
        log.Fatalln("expect string:", err)
    }

    // second step is to assign the second item into the 'Results' array
    if err := json.Unmarshal(rb.ResponseObjectRawJson[1], &wikiOpenSearch.Results); err != nil {
        log.Fatalln("expect []string:", err)
    }

    // log.Printf("wikiOpenSearch is: %+v", wikiOpenSearch)

    return wikiOpenSearch.Results
}
