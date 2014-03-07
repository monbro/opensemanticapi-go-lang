package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "opensemanticapi/scraper"
    "opensemanticapi/requestStruct"
    "reflect"
    "log"
    "encoding/json"
)

func TestScraper(t *testing.T) {
    Convey("Testing the scraper to search within wikipedia", t, func() {
        rb := new(scraper.RequestBit)
        rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
        rb.Work()
        response := new(requestStruct.WikiSearch)

        if err := json.Unmarshal(rb.ResponseObjectRawJson[0], &response.SearchTerm); err != nil {
            log.Fatalln("expect string:", err)
        }

        if err := json.Unmarshal(rb.ResponseObjectRawJson[1], &response.Results); err != nil {
            log.Fatalln("expect []string:", err)
        }

        Convey("should have the correct url", func() {
            So(rb.Url, ShouldEqual, "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3")
        })

        Convey("should store the plain response string", func() {
            So(rb.PlainResponse, ShouldEqual, "[\"database\",[\"Database\",\"Database transaction\",\"Database index\"]]")
        })

        Convey("should store the given struct as a the response object", func() {
            So(reflect.TypeOf(rb.ResponseObjectRawJson).String(), ShouldEqual, "[]json.RawMessage")
        })

        Convey("should Unmarschal the actual response into the response object", func() {
            So(response.SearchTerm, ShouldEqual, "database")
        })

        Convey("should Unmarschal the actual response into the response object", func() {
            So(response.Results[2], ShouldEqual, "Database index")
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
    })
}
