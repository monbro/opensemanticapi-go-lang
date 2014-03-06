package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "opensemanticapi/scraper"
    "opensemanticapi/requestStruct"
    "reflect"
    // "log"
)

// func TestScraper(t *testing.T) {
//     Convey("Testing the scraper", t, func() {
//         // catch a suggested list of results for a random keyword
//         url := "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
//         val := scraper.WikiSearch(url)

//         // this test is very basic and of course the result of this api request will change someday
//         Convey(`The result should be a string`, func() {
//             So(val[1], ShouldEqual, "Database transaction")
//         })

//         Convey("val should not be nil", func() {
//             So(val, ShouldNotBeNil)
//         })

//         Convey("val should not be nil", func() {
//             res := scraper.WikiGrab(val[1].(string))

//             So(res, ShouldNotBeNil)
//             So(res, ShouldEqual, "Database transaction")
//         })
//     })
// }

func TestScraper(t *testing.T) {
    Convey("Testing the scraper struct RequestBit with all its functions", t, func() {

        rb := new(scraper.RequestBit)
        rb.Url = "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
        rb.ResponseObject = new(requestStruct.WikiSearch)
        // w := *rb.ResponseObject.(*requestStruct.WikiSearch)


        // 1) actually responseobject is accessiable if it is an array slice
        // rb.ResponseObject = make(requestStruct.WikiSearch, 2, 100)

        rb.Work()

        Convey("should have the correct url", func() {
            So(rb.Url, ShouldEqual, "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3")
        })

        Convey("should store the plain response string", func() {
            So(rb.PlainResponse, ShouldEqual, "[\"database\",[\"Database\",\"Database transaction\",\"Database index\"]]")
        })

        Convey("should store the given struct as a the response object", func() {
            So(reflect.TypeOf(rb.ResponseObject).String(), ShouldEqual, "*requestStruct.WikiSearch")
            // So(reflect.TypeOf(rb.ResponseObject[1].([]interface{})[0]).String(), ShouldEqual, "*requestStruct.WikiSearch")
        })

        Convey("should Unmarschal the actual response into the response object", func() {
            So(rb.ResponseObject[0], ShouldEqual, "*requestStruct.WikiSearch")
        })

        Convey("should Unmarschal the actual response into the response object", func() {
            So(rb.ResponseObject[1], ShouldEqual, "[Database Database transaction Database index]")
        })
    })
}
