package adapter

import (
)

type AdapterInterface interface {
    Start()
    Runner(string)
    CreateSnippetWordsRelation(string)
    Configuration(string, int, bool, bool)
}
