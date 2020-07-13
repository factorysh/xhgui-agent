build: bin
	go build -o bin/xhgui-agent .

bin:
	mkdir -p bin

pull:
	docker pull bearstech/golang-dev

docker-build: bin
	mkdir -p .cache
	docker run --rm \
		-u `id -u` \
		-v `pwd`/.cache:/.cache \
		-v `pwd`:/go/src/github.com/factorysh/xhgui-agent \
		-w /go/src/github.com/factorysh/xhgui-agent \
		bearstech/golang-dev \
		make

docker-image:
	docker build -t xhgui-agent .

test:
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
