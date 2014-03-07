/**
 * provides functions to process requests to given urls
 * some example:
 * - https://gist.github.com/border/775526
 * - https://stackoverflow.com/questions/16931499/in-go-language-how-do-i-unmarshal-json-to-array-of-object
 * - https://stackoverflow.com/questions/13593519/how-do-i-parse-an-inner-field-in-a-nested-json-object-in-golang
 * - https://stackoverflow.com/questions/17209111/unable-to-parse-a-complex-json-in-golang
 * - http://play.golang.org/p/AEC_TyXE3B
 * - http://play.golang.org/p/TFUgJsWNhq
 * - https://stackoverflow.com/questions/19482612/go-golang-array-type-inside-struct-missing-type-composite-literal
 * - https://gobyexample.com/json
 * - http://mattyjwilliams.blogspot.co.uk/2013/01/using-go-to-unmarshal-json-lists-with.html
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
