**Installation**

* install go language (```brew install go``` on a mac or ```http://golang.org/doc/install``` or via gvm ```bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)```)
* install your redis server (http://redis.io/topics/quickstart) on a disk with several free GB
* download redis and unzip it
* clone this repo ```git clone https://github.com/monbro/opensemanticapi.git```
* ```export GOPATH=/your/go/workspace/folder/```
* ```cd $GOPATH```
* ```go get -d -v ./... && go build -v ./...``` to install all go dependencies - go will lookup on all import path's and grab the needed repos itself - pretty awesome
