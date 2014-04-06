**Usage**

**How to run the cronjob / scraper**

* ```go run src/github.com/monbro/opensemanticapi-go-lang/main.go``` to run the test script
if you are in the terminal of the redis-cli you can check the result with ```select 10``` to be in the correct datbase
and with ```SORT database by database:* Limit 0 120 DESC GET #``` you should get a example

Also you can watch the inserts to your databse by using ```monitor``` in your redis client.

If you want to run the cronjob ongoing, just use the flag "infinite" ```go run src/github.com/monbro/opensemanticapi-go-lang/main.go -infinite=true```

Availabel flags:
* `-api=false` Do you want to start the api server?
* `-infinite=false` Do you want to run the cronjob infinite?
* `-fast=true` Do you want to run in super fast mode (heavy cpu usage etc.)?
* `-v=1 -alsologtostderr=true` to show the log messages or see https://godoc.org/github.com/golang/glog for more details

**How to run the REST API**

* ```go run src/github.com/monbro/opensemanticapi-go-lang/main.go -api=true``` to launch the REST API server script
* shoule be now accessible via ```http://localhost:3000/relation/database``` after the cronjob was running once
