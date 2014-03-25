**Testing**

**Add new tests**

* check the goconvey documenation: https://github.com/smartystreets/goconvey/wiki/Documentation
* and check the available assertions: https://github.com/smartystreets/goconvey/wiki/Assertions

**How to run all test's manually**

* ```go test github.com/monbro/opensemanticapi-go-lang/tests/scraper```
* ```go test github.com/monbro/opensemanticapi-go-lang/tests/database```
* ```go test github.com/monbro/opensemanticapi-go-lang/tests/analyse``` // @TODO write tests
* ```go test github.com/monbro/opensemanticapi-go-lang/tests/api``` // @TODO write tests
* ```go test github.com/monbro/opensemanticapi-go-lang/tests/requestStruct```

**How to run the goconvey test's server**

* ensure that you installed goconvey proper ```https://github.com/smartystreets/goconvey#installation```
* switch into the folder where you want to watch the tests, e.g. 'cd src/github.com/monbro/opensemanticapi-go-lang/'
* run the test server```$GOPATH/bin/goconvey```
* access the webinterface on http://localhost:8080
