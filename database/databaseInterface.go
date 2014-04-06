package database

import (
)

/**
 * common interface to describe all database adapter
 * @TODO tidy up the methods
 */
type DatabaseInterface interface {
    Init(string, int)
    Close()
    Set(string, string) (interface{}, error)
    Get(string) (interface{}, error)
    Members(string) (interface{}, error)
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
