build: bin vendor
	go build -o bin/xhgui-agent .

bin:
	mkdir -p bin

vendor:
	dep ensure

pull:
	docker pull bearstech/golang-dep

docker-build:
	docker run --rm \
		-v `pwd`:/go/src/github.com/factorysh/xhgui-agent \
		-w /go/src/github.com/factorysh/xhgui-agent \
		bearstech/golang-dep \
		make

test: vendor
	go test -v github.com/factorysh/xhgui-agent/fixedqueue
	go test -v github.com/factorysh/xhgui-agent/agent