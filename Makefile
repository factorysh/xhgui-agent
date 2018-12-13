build: bin vendor
	go build -o bin/xhgui-agent .

bin:
	mkdir -p bin

vendor:
	dep ensure