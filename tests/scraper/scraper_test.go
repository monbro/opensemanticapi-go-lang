package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi/scraper"
    "github.com/monbro/opensemanticapi/requestStruct"
    "net/url"
    "reflect"
    "log"
)

func TestScraper(t *testing.T) {
    Convey("Testing the scraper to search within wikipedia", t, func() {
        rb := new(scraper.RequestBit)
        rb.Url = "https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch="+url.QueryEscape("database")+"&format=json"

        rb.ResponseObjectInterface = new(requestStruct.WikiSearch)
        rb.Work()
        response := *rb.ResponseObjectInterface.(*requestStruct.WikiSearch)

        Convey("should have the correct url", func() {
            So(rb.Url, ShouldEqual, "https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=database&format=json")
        })

        Convey("should store the plain response string", func() {
            So(rb.PlainResponse, ShouldContainSubstring, `"title":"Database","snippet":"A <span class='searchmatch'>database</span> is an organized collection of data .`)
            So(rb.PlainResponse, ShouldContainSubstring, `{"query-continue":{"search":{"sroffset":10}},"query":{"searchinfo":`)
        })

        Convey("should store the given struct as a the response object", func() {
            log.Printf("Length word list: %+v", rb.ResponseObjectRawJson)
            So(reflect.TypeOf(rb.ResponseObjectRawJson).String(), ShouldEqual, "json.RawMessage")
            So(reflect.TypeOf(rb.ResponseArrayRawJson).String(), ShouldEqual, "[]json.RawMessage")
        })

        Convey("should Unmarschal the actual response into the response object and the first result should be ", func() {
            So(response.Query.Search[0].Title, ShouldEqual, "Database")
            So(len(response.Query.Search), ShouldBeGreaterThan, 0)
        })
    })

    Convey("Testing the scraper to get a page from wikipedia", t, func() {
        rb := new(scraper.RequestBit)
        rb.Url = "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=Yanqing_County"
        rb.ResponseObjectInterface = new(requestStruct.WikiPage)
        rb.Work()
        response := *rb.ResponseObjectInterface.(*requestStruct.WikiPage)

        Convey("should have the correct url", func() {
            So(rb.Url, ShouldEqual, "http://en.wikipedia.org/w/api.php?rvprop=content&format=json&prop=revisions|categories&rvprop=content&action=query&titles=Yanqing_County")
        })

        Convey("should store the plain response string", func() {
            So(rb.PlainResponse, ShouldNotEqual, nil)
        })

        Convey("should store the given struct as a the response object", func() {
            So(reflect.TypeOf(response).String(), ShouldEqual, "requestStruct.WikiPage")
        })

        Convey("should Unmarschal the actual response into the response object", func() {
            So(response.Query.Pages["2256752"].Title, ShouldEqual, "Yanqing County")
            So(response.Query.Pages["2256752"].PageId, ShouldEqual, 2256752)
        })

        Convey("should Unmarschal the actual response into the response object and contain the wanted data", func() {
            So(response.Query.Pages["2256752"].Rev[0].RawContent, ShouldContainSubstring, `name = {{raise|0.2em|Yanqing County}}`)
        })
    })
}
