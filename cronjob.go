package main

import (
    "github.com/monbro/opensemanticapi-go-lang/analyse"
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

    /**
     * comment the following line in if you want to loop for more results
     */
    // worker.InfiniteWorking = true

    worker.Run()
}