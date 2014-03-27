package api

/**
 * @TODO do the json encoding the proper way
 * extend the api to more possibilities
 */

import (
    // "github.com/monbro/opensemanticapi-go-lang/api" // @TODO put the logic into this file
    "log"
    "github.com/codegangsta/martini"
     "github.com/monbro/opensemanticapi-go-lang/database"
     "strings"
)

func StartServer() {
    log.Println("Starting API server now.")

    // configurate the database connection
    Db := new(database.Database)
    Db.Init("", 10)

    // set up a nice glas of martini
    m := martini.Classic()

    // http://localhost:3000
    m.Get("/", func() string {
        return `{"result":"Hello world!"}`
    })

    // eg: http://localhost:3000/relation/database
    m.Get("/relation/:word", func(params martini.Params) string {
        relations := Db.GetPopularWordRelations(params["word"])
        return `{"result": ["`+strings.Join(relations, `", "`)+`"]}`
    })

    // eg: http://localhost:3000/popular-relations/
    m.Get("/popular-relations/", func(params martini.Params) string {
        relations := Db.GetMostPopularWords()
        return `{"result": ["`+strings.Join(relations, `", "`)+`"]}`
    })

    // actual launch the server
    m.Run()
}
