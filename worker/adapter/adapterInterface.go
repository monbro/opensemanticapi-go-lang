package adapter

import (
)

type AdapterInterface interface {
    Run()
    RunNext(string)
    CreateSnippetWordsRelation(string)
    Configuration(string, int, bool, bool)
}
