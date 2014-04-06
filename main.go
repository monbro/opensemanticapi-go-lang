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
    SnippetLength = 120
    StartSearchTerm = "london"
)

func main() {

    // preparation to use flags
    isApiServer := flag.Bool("api", false, "Do you want to start the api server?")
    isFastMode := flag.Bool("fast", false, "Do you want to run in super fast mode (heavy cpu usage etc.)?")
    isInfiniteCronjobRun := flag.Bool("infinite", false, "Do you want to run the cronjob infinite?")
    flag.Parse()

    if *isApiServer {
        fmt.Println("Starting API server ...")
        api.StartServer()
    } else {

        adapterName := "concurrencyA"

        worker := analyse.WorkerFactory(
                            adapterName,
                            StartSearchTerm,
                            SnippetLength,
                            *isFastMode,
                            *isInfiniteCronjobRun,
                        )

        fmt.Println("Starting cronjob ...")
        worker.Run()
    }
}
