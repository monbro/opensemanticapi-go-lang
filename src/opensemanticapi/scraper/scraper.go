/**
 * https://gocasts.io/gocasts/simple-http-get
 */

package scraper

import (
    "log"
    "io/ioutil"
    "net/http"
)

func FetchUrlContent(url string) string{
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    s := string(body[:])

    if err != nil {
        log.Fatal(err)
    }

    return s
}
