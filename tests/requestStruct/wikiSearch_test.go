package tests

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/monbro/opensemanticapi-go-lang/requestStruct"
    "encoding/json"
    "log"
)

func TestWikiSearch(t *testing.T) {
    Convey("should allow to parse a wikisearch Json response into this struct", t, func() {

        // test database configuration
        fakeWikipediaJsonResponse := `{"warnings":{"main":{"*":"Unrecognized parameter: 'limit'"}},"query-continue":{"search":{"sroffset":10}},"query":{"searchinfo":{"totalhits":182461},"search":[{"ns":0,"title":"Database","snippet":"A <span class='searchmatch'>database</span> is an organized collection of data .  The data are typically organized to model relevant aspects of reality in a way that  <b>...</b> ","size":63701,"wordcount":9022,"timestamp":"2014-03-15T07:01:42Z"},{"ns":0,"title":"Geographic Names Information System","snippet":"The Geographic Names Information System (GNIS) is a <span class='searchmatch'>database</span> that contains name and locative information about more than two million  <b>...</b> ","size":3977,"wordcount":541,"timestamp":"2014-02-20T13:15:03Z"},{"ns":0,"title":"Relational database","snippet":"A relational <span class='searchmatch'>database</span> is a <span class='searchmatch'>database</span>  that has a collection of tables  of data items, all of which is formally described and organized  <b>...</b> ","size":16495,"wordcount":2280,"timestamp":"2014-02-13T12:25:43Z"},{"ns":0,"title":"Relational database management system","snippet":"A relational <span class='searchmatch'>database</span> management system (RDBMS) is a <span class='searchmatch'>database</span> management system  (DBMS) that is based on the relational model  as  <b>...</b> ","size":6747,"wordcount":914,"timestamp":"2014-03-01T06:44:34Z"},{"ns":0,"title":"Bibliographic database","snippet":"A bibliographic <span class='searchmatch'>database</span> is a <span class='searchmatch'>database</span>  of bibliographic record s, an organized digital collection of references to published literature,  <b>...</b> ","size":3999,"wordcount":504,"timestamp":"2014-03-09T17:46:06Z"},{"ns":0,"title":"List of recurring The Simpsons characters","snippet":"<span class='searchmatch'>Database</span>: <span class='searchmatch'>Database</span>, or Data, (real name is Kyle according to Yellow Subterfuge ) is a nerd y student who attends Springfield Elementary  <b>...</b> ","size":215120,"wordcount":32439,"timestamp":"2014-03-12T19:02:38Z"},{"ns":0,"title":"Database index","snippet":"A <span class='searchmatch'>database</span> index is a data structure  that improves the speed of data retrieval operations on a <span class='searchmatch'>database</span> table  at the cost of additional  <b>...</b> ","size":15615,"wordcount":2455,"timestamp":"2014-03-04T22:06:54Z"},{"ns":0,"title":"Biological database","snippet":"Biological <span class='searchmatch'>databases</span> are libraries of life sciences information, collected from scientific experiments, published literature, high- <b>...</b> ","size":7549,"wordcount":900,"timestamp":"2014-03-08T16:35:49Z"},{"ns":0,"title":"Online database","snippet":"An online <span class='searchmatch'>database</span> is a <span class='searchmatch'>database</span>  accessible from a network, including from the Internet . It differs from a local <span class='searchmatch'>database</span>, held in an  <b>...</b> ","size":1270,"wordcount":168,"timestamp":"2013-10-06T10:39:48Z"},{"ns":0,"title":"List of academic databases and search engines","snippet":"This page contains a representative list of major <span class='searchmatch'>databases</span> and search engines useful in an academic setting for finding and accessing  <b>...</b> ","size":38000,"wordcount":3880,"timestamp":"2014-02-26T04:02:56Z"}]}}`
        fakeWikipediaJsonResponseByte := []byte(fakeWikipediaJsonResponse)

        wikiSearchStruct := new(requestStruct.WikiSearch)

        if err := json.Unmarshal(fakeWikipediaJsonResponseByte, &wikiSearchStruct); err != nil {
            log.Printf("error within json.Unmarshal for wikiSearchStruct:", err)
        }

        log.Printf("Added page to queue: '%+v'", wikiSearchStruct)

        So(len(wikiSearchStruct.Query.Search), ShouldEqual, 10)
        So(wikiSearchStruct.Query.Search[1].Title, ShouldEqual, "Geographic Names Information System")
        So(wikiSearchStruct.Query.Searchinfo.TotalHits, ShouldEqual, 182461)
    })
}
