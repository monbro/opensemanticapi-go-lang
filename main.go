package main

import (
    "github.com/monbro/opensemanticapi/analyse"
)

/**
 * configurations constantes
 */
const (
    SNIPPET_LENGTH = 120
    START_SEARCH_TERM = "database"
)

func main() {
    worker := new(analyse.Worker)
    worker.START_SEARCH_TERM = START_SEARCH_TERM
    worker.SNIPPET_LENGTH = SNIPPET_LENGTH
    worker.Run()
}
