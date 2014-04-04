/**
 * provides functions to create wording context relations
 *
 * THIS CLASS IS FOR PERFORMANCE TEST PURPOSE ONLY
 */

package analyse

import (
    "github.com/golang/glog"
    "strconv"
    "github.com/monbro/opensemanticapi-go-lang/database"
)

type WorkerA struct {
    Debug bool
    Db *database.RedisMulti
    START_SEARCH_TERM string
    SNIPPET_LENGTH int
    InfiniteWorking bool
}

/**
 * will run the process of storing words that are related in its context
 */
func (w *WorkerA) RunNext(searchTerm string) {

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

    snippetsRaw := GetSnippetsFromText(rawContent)
    snippets := GetSnippetsFromText(rawContent)
    snippets = CleanUpSnippets(snippets)

    w.Db.Multi()

    for i := range snippets {
        if w.SNIPPET_LENGTH < len(snippets[i]) {
            glog.Info("Snippet "+strconv.Itoa(i)+"/"+strconv.Itoa(len(snippets))+" with a length of "+strconv.Itoa(len(snippetsRaw[i])))

            // analyse the text block
            w.CreateSnippetWordsRelation(snippets[i])

            // raise counter for text blocks
            w.Db.RaiseScrapedTextBlocksCounter()
        }
    }

    // flush the queued commands from the pipeline
    w.Db.Flush()

    if w.InfiniteWorking {

        // create aloop by calling it self for the next search term
        w.RunNext(w.Db.RandomPageFromQueue())
    }
}

/**
 * will analyse a snippet by spinning relations between each word within this snippet
 *
 * @TODO should be changed to fit https://gobyexample.com/worker-pools probably?
 */
func (w *WorkerA) CreateSnippetWordsRelation(snippet string) {

    words := GetWordsFromSnippet(snippet)
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

}
