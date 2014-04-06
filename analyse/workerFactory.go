/**
 * will take care about the worker, no matter which one you choose
 */

package analyse

import (
    "github.com/monbro/opensemanticapi-go-lang/analyse/adapter"
)

func WorkerFactory(
                    adapterName string,
                    StartSearchTerm string,
                    SnippetLength int,
                    isFastMode bool,
                    isInfiniteCronjobRun bool,
                    ) adapter.AdapterInterface {

    var worker adapter.AdapterInterface

    switch adapterName {
        case "concurrencyA":
            worker = new(adapter.ConcurrencyA)
        case "serialA":
            worker = new(adapter.SerialA)
        default:
            panic("No adapter with this name found!")
    }

    worker.Configuration(
        StartSearchTerm,
        SnippetLength,
        isFastMode,
        isInfiniteCronjobRun,
        )

    return worker
}
