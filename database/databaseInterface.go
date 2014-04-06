package database

import (
)

type database interface {
    Init()
    Close()
    Flushall()
    AddPageToQueue()
    AddPageToDone()
    AddPagesToQueue()
    RandomPageFromQueue()
    AddWordRelation()
    GetPopularWordRelations()
    GetMostPopularWords()
    GetAnalysedTextBlocksCounter()
    RaiseScrapedTextBlocksCounter()
}
