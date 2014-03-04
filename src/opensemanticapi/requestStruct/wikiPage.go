/**
 * provides functions to process requests to given urls
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
