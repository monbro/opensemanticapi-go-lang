package main

import (
    // "github.com/monbro/opensemanticapi/api"
    "log"
    "github.com/codegangsta/martini"
)

/**
 * configurations constantes
 */
const (
    SOMETHING = 120
)

func main() {
    log.Println("Starting API now")

    m := martini.Classic()
    m.Get("/", func() string {
    return "Hello world!"
    })
    m.Run()
}
