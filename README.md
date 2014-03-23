opensemanticapi
===============

[![Build Status](https://travis-ci.org/monbro/opensemanticapi.png?branch=go-lang)](https://travis-ci.org/monbro/opensemanticapi)

**Works with GO LANG**

**This branch is unstable and experimental, but includes a working prototype**

**Description**

Will allow you to create your own semantic wording database with redis. Otherwise there will be a open api to get related words by meaning. You could say, this implementation is a light version of the idea behind http://en.wikipedia.org/wiki/Latent_semantic_analysis in combination with http://en.wikipedia.org/wiki/Open-source_intelligence

![ScreenShot](https://raw.githubusercontent.com/monbro/opensemanticapi/go-lang/osapi_explanation.jpg)

**Concerns**

I thought about using a wikipedia database dump file which is available here http://dumps.wikimedia.org.
But I came to the conclusion that I do want to work in a direction to work with api data as a dump is not always available. So please see this wikipedia api scraping method as a temporary solution as it is inefficient for our case.

**Examples**

The following examples where given after the system was collecting for about one hour only.

**1. Example (http://localhost:8080/relations/ship):**

* Input: "ship"

* Output: ["midshipmen", "aboard", "ships", "rating", "master", "served", "seaman", "sea", "officers", "santa", "sailing", "cadets", "able", "sail", "navigation", "lieutenant", "hms", "expected", "yahoo", "storm", "rated", "promotion", "maría", "lewis", "false", "era", "boys", "wind", "voyage", "volunteer", "servants", "required", "passing", "palos"]

**2. Example (http://localhost:8080/relations/human):**

* Input: "human"

* Output: ["humans", "evolution", "primates", "ago", "ado", "studies", "physiology", "bonobo"]

**3. Example (http://localhost:8080/relations/dog):**

* Input: "dog"

* Output: ["infant", "wildlife", "offspring", "mother", "future", "southwest", "koalas", "conflict", "animals", "aitken", "wolf", "urban", "rehabilitation", "pet", "perspective", "nursing", "mexico", "evolutionary", "weaning", "ticket", "texas", "speech", "special", "retrospective", "primate", "holtcamp", "fund", "enough", "domestic", "cost", "arizona", "210–217", "variety", "trivers", "trauma", "terms", "sprawl", "southwestern", "sense", "river", "received", "questions", "point", "perhaps", "parent", "otter", "makes", "little", "less", "himself", "gray", "gorilla", "frequently"]

**Installation Guide**

* install go language (```brew install go``` on a mac or ```http://golang.org/doc/install``` or via gvm ```bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)```)
* install your redis server (http://redis.io/topics/quickstart) on a disk with several free GB
* start your redis server with ```redis-server``` on a new console tab
* clone this repo ```git clone https://github.com/monbro/opensemanticapi.git```
* ```export GOPATH=/your/go/workspace/folder/```
* ```cd $GOPATH``
* ```go get -d -v ./... && go build -v ./...``` to install all go dependencies - go will lookup on all import path's and grab the needed repos itself - pretty awesome

**How to run the cronjob / scraper**

* ```go run src/github.com/monbro/opensemanticapi/cronjob.go``` to run the test script
if you are in the terminal of the redis-cli you can check the result with ```select 10``` to be in the correct datbase
and with ```SORT database by database:* Limit 0 120 DESC GET #``` you should get a example

**How to run the REST API**

* ```go run src/github.com/monbro/opensemanticapi/api_server.go``` to launch the REST API server script
* shoule be now accessible via http://localhost:3000/relation/database

**Testing**

* needs to be written

**Add new tests**

* check the goconvey documenation: https://github.com/smartystreets/goconvey/wiki/Documentation
* and check the available assertions: https://github.com/smartystreets/goconvey/wiki/Assertions

**How to run all test's manually**

* ```go test github.com/monbro/opensemanticapi/tests/scraper```
* ```go test github.com/monbro/opensemanticapi/tests/database```
* ```go test github.com/monbro/opensemanticapi/tests/analyse``` // @TODO write tests
* ```go test github.com/monbro/opensemanticapi/tests/api``` // @TODO write tests
* ```go test github.com/monbro/opensemanticapi/tests/requestStruct```

**How to run the goconvey test's server**

* ensure that you installed goconvey proper ```https://github.com/smartystreets/goconvey#installation```
* switch into the folder where you want to watch the tests, e.g. 'cd src/github.com/monbro/opensemanticapi/'
* run the test server```$GOPATH/bin/goconvey```
* access the webinterface on http://localhost:8080

**NOTE's**

* function names always starting with a uppercase letter
* GOROOT should be set to go lang itself
* GOPATH should be set to where your go projects are
* http://www.golangpatterns.info/object-oriented/classes
* http://blog.golang.org/laws-of-reflection
* http://golang.org/pkg/reflect/
* https://code.google.com/p/go-wiki/wiki/GithubCodeLayout
* http://golang.org/doc/code.html
* http://godoc.org/?q=markdown // good to search anything go related

This software is published under the MIT-License. See 'license' for more information.
