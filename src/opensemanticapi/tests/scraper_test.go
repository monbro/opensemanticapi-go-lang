package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "opensemanticapi/scraper"
)

func TestScraper(t *testing.T) {
    Convey("Testing the scraper", t, func() {
        // catch a suggested list of results for a random keyword
        url := "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
        val := scraper.WikiSearch(url)

        // this test is very basic and of course the result of this api request will change someday
        Convey(`The result should be a string`, func() {
            So(val[1], ShouldEqual, "Database transaction")
        })

        Convey("val should not be nil", func() {
            So(val, ShouldNotBeNil)
        })

        Convey("val should not be nil", func() {
            res := scraper.WikiGrab(val[1].(string))

            So(res, ShouldNotBeNil)
            So(res, ShouldEqual, "Database transaction")
        })
    })
}
