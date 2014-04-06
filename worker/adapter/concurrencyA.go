/**
 * provides functions to create wording context relations
 */

package adapter

import (
    "net/url"
    // "encoding/json"
    "sync"
    "strconv"
    "github.com/golang/glog"
    "github.com/monbro/opensemanticapi-go-lang/worker/util"
    "github.com/monbro/opensemanticapi-go-lang/scraper"
    "github.com/monbro/opensemanticapi-go-lang/requestStruct"
    "github.com/monbro/opensemanticapi-go-lang/database"
)

type ConcurrencyA struct {
    Db *database.RedisMulti
    StartSearchTerm string
    SnippetLength int
    IsInfiniteWorking bool
    IsFastMode bool
    Wg sync.WaitGroup
}

/**
 * configuration
 */
func (w *ConcurrencyA) Configuration(
            StartSearchTerm string,
            SnippetLength int,
            IsFastMode bool,
            IsInfiniteCronjobRun bool) {
    w.StartSearchTerm = StartSearchTerm
    w.SnippetLength = SnippetLength
    w.IsFastMode = IsFastMode
    w.IsInfiniteWorking = IsInfiniteCronjobRun
}

/**
 * will initially start the process
 */
func (w *ConcurrencyA) Start() {

    glog.Info("Using adapter concurrencyA ...")

    if w.IsFastMode {
        util.MaximumUlimit()
    }

    // init database
    w.Db = new(database.RedisMulti)
    w.Db.Init("", 10)

    // initial start
    w.Runner(w.StartSearchTerm)
}

/**
 * will run the process of storing words that are related in its context
 */
func (w *ConcurrencyA) Runner(searchTerm string) {

    glog.Infof("Searchterm Now: '%+v'", searchTerm)

    // declare needed variables
    var pages []string

    // initial search request to get some pages back
    pages = SearchWikipedia(searchTerm)
    searchTerm = pages[0]

    // add this page to the done collection
    w.Db.AddPageToDone(searchTerm)

    // get a slice which will exclude the first element as we processing this one soon
    pagesToQueue := pages[1:]

    // pagesToQueue = []string{"pageOne", "Geographic Names Information System"}
    w.Db.AddPagesToQueue(pagesToQueue)

    rawContent := GetWikipediaPage(pages[0])

    snippetsRaw := util.GetSnippetsFromText(rawContent)
    snippets := util.GetSnippetsFromText(rawContent)
    snippets = util.CleanUpSnippets(snippets)

    w.Db.Multi()

    for i := range snippets {
        if w.SnippetLength < len(snippets[i]) {
            // log.Println("Snippet "+strconv.Itoa(i)+"/"+strconv.Itoa(len(snippets))+" with a length of "+strconv.Itoa(len(snippetsRaw[i])))
            glog.Info("Snippet "+strconv.Itoa(i)+"/"+strconv.Itoa(len(snippets))+" with a length of "+strconv.Itoa(len(snippetsRaw[i])))

            // analyse the text block
            go w.CreateSnippetWordsRelation(snippets[i])

            // raise counter for text blocks
            w.Db.RaiseScrapedTextBlocksCounter()
        }
    }

    if !w.IsFastMode {
        // wait for all snippets to be finished
        glog.Info("Waiting now to have all the go routines completed")
        w.Wg.Wait()
    }

    // flush the queued commands from the pipeline
    w.Db.Flush()

    if w.IsInfiniteWorking {

        // create aloop by calling it self for the next search term
        w.Runner(w.Db.RandomPageFromQueue())
    }
}

/**
 * will search wikipedia for a search term and return existing matching pages
 */
func SearchWikipedia(searchTerm string) []string {
    // lets create a new http request object
    rb := new(scraper.RequestBit)

    rb.Url = "https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch="+url.QueryEscape(searchTerm)+"&format=json"
    glog.Infof("Url crawling: %+v", rb.Url)

    // inject the struct for the json response
    rb.ResponseObjectInterface = new(requestStruct.WikiSearch)
    rb.Work() // fire the request

    w2 := *rb.ResponseObjectInterface.(*requestStruct.WikiSearch)

    var results []string

    for _, value := range w2.Query.Search {
        // results[len(results)] = value.Rev[0].RawContent
        // http://blog.golang.org/slices
        results = append(results, value.Title)
    }

    return results
}

/**
 * will get the content of a wikipedia page
 */
func GetWikipediaPage(pageTitle string) string {
    // lets create a new http request object
    rb := new(scraper.RequestBit)
    rb.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles="+url.QueryEscape(pageTitle)
    glog.Infof("Url crawling: %+v", rb.Url)

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
 *
 * @TODO should be changed to fit https://gobyexample.com/ConcurrencyA-pools probably?
 */
func (w *ConcurrencyA) CreateSnippetWordsRelation(snippet string) {

    if !w.IsFastMode {
        w.Wg.Add(1)
    }

    words := util.GetWordsFromSnippet(snippet)
    for _, word := range words {
        // check if word has more than 2 letters and this includes checking for an empty string
        if len(word) > 2 {
            for _, relation := range words {
                // check if we not adding a relation to the word itself
                // check if the relation is more than 2 letters long and not an empty string
                if word != relation &&
                    len(relation) > 2 {
                        w.Db.AddWordRelation(word, relation)
                }
            }
        }
    }

    if !w.IsFastMode {
        w.Wg.Done()
    }
}

/**
 * thats how the method would like with as a autarkic function
 */
// func snippetConcurrencyA(snippets []string, Db *database.RedisMulti, SnippetLength int) {
//     for i := range snippets {
//         if SnippetLength < len(snippets[i]) {
//             log.Println("Snippet "+strconv.Itoa(i)+"/"+strconv.Itoa(len(snippets))+" with a length of "+strconv.Itoa(len(snippets[i])))

//             // analyse the text block
//             words := GetWordsFromSnippet(snippets[i])
//             for _, word := range words {
//                 // check if word has more than 2 letters and this includes checking for an empty string
//                 if len(word) > 2 {
//                     for _, relation := range words {
//                         // check if we not adding a relation to the word itself
//                         // check if the relation is more than 2 letters long and not an empty string
//                         if word != relation &&
//                             len(relation) > 2 {
//                                 Db.AddWordRelation(word, relation)
//                         }
//                     }
//                 }
//             }

//             // raise counter for text blocks
//             Db.RaiseScrapedTextBlocksCounter()
//         }
//     }
// }

/**
 * will search wikipedia for a search term and return existing matching pages
 * will be outdated soon!
 *
 * NOT IN USE CURRENTLY
 */
// func OpenSearchWikipedia(searchTerm string) []string {
//     // lets create a new http request object
//     rb := new(scraper.RequestBit)

//     rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search="+url.QueryEscape(searchTerm)+"&format=json&limit=3"
//     glog.Infof("Url crawling in SearchWikipedia: %+v", rb.Url)
//     rb.Work() // fire the request

//     // as wikipedia returns a sh*t formatted json we need to assign the result in two steps
//     wikiOpenSearch := new(requestStruct.WikiOpenSearch)

//     // first step is to assign the result term which is the first item in the returned array
//     if err := json.Unmarshal(rb.ResponseArrayRawJson[0], &wikiOpenSearch.SearchTerm); err != nil {
//         glog.Fatalf("expect string: %+v", err)
//     }

//     // second step is to assign the second item into the 'Results' array
//     if err := json.Unmarshal(rb.ResponseArrayRawJson[1], &wikiOpenSearch.Results); err != nil {
//         glog.Fatalf("expect []string: %+v", err)
//     }

//     return wikiOpenSearch.Results
// }
