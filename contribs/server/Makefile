build: build-php build-nginx

src:
	mkdir -p src

src/xhgui/README.md: src
	git submodule update --init

build-php-tool:
	docker build \
		-t xhgui-php-tool \
		-f Dockerfile.php-tool \
		.

src/xhgui/vendor: src/xhgui/README.md build-php-tool
	docker run -ti \
		--rm \
		-u `id -u` \
		-v $(CURDIR)/src/xhgui:/toto \
		-v ~/.composer:/.composer \
		-w /toto \
		xhgui-php-tool \
		composer install

build-php: src/xhgui/vendor
	docker build \
		-t xhgui-php \
		-f Dockerfile.php \
		.

build-nginx: src/xhgui/vendor
	docker build \
		-t xhgui-nginx \
		-f Dockerfile.nginx \
		.

clean:
	rm -rf vendor src/xhgui

