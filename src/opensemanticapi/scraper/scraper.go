/**
 * https://gocasts.io/gocasts/simple-http-get
 */

package scraper

import (
    "log"
    "io/ioutil"
    "net/http"
    "os"
)

func FetchUrlContent(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    _, err = os.Stdout.Write(body)

    if err != nil {
        log.Fatal(err)
    }
}
