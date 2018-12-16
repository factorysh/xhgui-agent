build: bin vendor
	go build -o bin/xhgui-agent .

bin:
	mkdir -p bin

vendor:
	dep ensure

test: vendor
	go test -v github.com/factorysh/xhgui-agent/fixedqueue
	go test -v github.com/factorysh/xhgui-agent/agent