package database

import (
)

type DatabaseInterface interface {
    Init(string, int)
    Close()
    Flushall()
    AddPageToQueue(string)
    AddPageToDone(string)
    AddPagesToQueue([]string)
    RandomPageFromQueue() string
    AddWordRelation(string, string)
    GetPopularWordRelations(string) []string
    GetMostPopularWords() []string
    GetAnalysedTextBlocksCounter() string
    RaiseScrapedTextBlocksCounter()
}
