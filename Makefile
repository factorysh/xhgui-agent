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

demo: | docker-build docker-image
	cd contribs/server && make
	cd contribs/client && make
	docker-compose up -d

down:
	docker-compose down