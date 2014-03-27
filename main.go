package main

import (
    "github.com/monbro/opensemanticapi-go-lang/analyse"
    "github.com/monbro/opensemanticapi-go-lang/api"
    "flag"
    "fmt"
)

/**
 * configurations constantes
 */
const (
    SNIPPET_LENGTH = 120
    START_SEARCH_TERM = "database"
)

func main() {

    // preparation to use flags
    isApiServer := flag.Bool("api", false, "Do you want to start the api server?")
    isInfiniteCronjobRun := flag.Bool("infinite", false, "Do you want to run the cronjob infinite?")
    flag.Parse()

    if *isApiServer {
        fmt.Println("Starting API server ...")
        api.StartServer()
    } else {
        fmt.Println("Starting cronjob ...")

        worker := new(analyse.Worker)
        worker.START_SEARCH_TERM = START_SEARCH_TERM
        worker.SNIPPET_LENGTH = SNIPPET_LENGTH

        // check if the cronjob should run continouing
        if *isInfiniteCronjobRun {
            worker.InfiniteWorking = true
        }

        worker.Run()
    }
}
