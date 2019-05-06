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

docker-image:
	docker build -t xhgui-agent .

test: vendor
	go test -v github.com/factorysh/xhgui-agent/fixedqueue

clean:
	rm -rf bin data
	make -C contribs/server clean
	make -C contribs/client clean

demo: | docker-build docker-image
	make -C contribs/server
	make -C contribs/client
	docker-compose up -d

down:
	docker-compose down
