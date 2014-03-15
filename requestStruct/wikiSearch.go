/**
 * provides functions to process requests to given urls
 */

package requestStruct

import (
)

type WikiSearch struct {
    Query SubTypeFirst `json:"query"`
}

type SubTypeFirst struct {
    Searchinfo SubTypeSecond `json:"searchinfo"`
    Search []SearchResults `json:"search"`
}

type SubTypeSecond struct {
    TotalHits int `json:"totalhits"`
}

type SearchResults struct {
    Ns int `json:"ns"`
    Title string `json:"title"`
    Snippet string `json:"snippet"`
    Size int `json:"size"`
    Wordcount int `json:"wordcount"`
    Timestamp string `json:"timestamp"`
}
