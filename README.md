opensemanticapi
===============

[![Build Status](https://travis-ci.org/monbro/opensemanticapi.png?branch=go-lang)](https://travis-ci.org/monbro/opensemanticapi)

**Works with GO LANG**

**This is a developer branch**

**Current target is to find a suitable database engine**

**Description**

Will allow you to create your own semantic wording database with redis. Otherwise there will be a open api to get related words by meaning. You could say, this implementation is a light version of the idea behind http://en.wikipedia.org/wiki/Latent_semantic_analysis in combination with http://en.wikipedia.org/wiki/Open-source_intelligence

![ScreenShot](https://raw.github.com/monbro/opensemanticapi/master/infographic.png)

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
* ```export GOPATH=/absolute/path/to/opensemanticapi```
* ```git submodule update --init```
* ```go get -d -v ./... && go build -v ./...``` to install all go dependencies

**How to run the main program**

* ```go run src/opensemanticapi/main.go``` to run the test script

**How to run all test's manually**

* ```go test opensemanticapi/tests```

**How to run the goconvey test's server**

* ```export GOROOT=/path/to/go/projects/opensemanticapi```
* switch into the folder where you want to watch the tests, e.g. 'cd src/opensemanticapi/'
* ```$GOPATH/bin/goconvey```
* access the webinterface on http://localhost:8080

**NOTE's**

* function names always starting with a uppercase letter
* GOROOT should be set to go lang itself
* GOPATH should be set to where your go projects are
* http://www.golangpatterns.info/object-oriented/classes
* http://blog.golang.org/laws-of-reflection
* http://golang.org/pkg/reflect/


This software is published under the MIT-License. See 'license' for more information.
