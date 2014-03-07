/**
 * provides functions to process requests to given urls
 * some example:
 * - https://gist.github.com/border/775526
 * - https://stackoverflow.com/questions/16931499/in-go-language-how-do-i-unmarshal-json-to-array-of-object
 */

package requestStruct

import (
)

type WikiPage struct {
    Query SubType `json:"query"`
}

type ActualPage struct {
    PageId int  `json:"pageid"`
    Title string `json:"title"`
}

type SubType struct {
    Pages map[string]ActualPage `json:"pages"`
}
