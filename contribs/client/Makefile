build: vendor/xhgui-collector
	docker build \
		-t php-with-xhgui-agent \
		.

vendor/xhgui-collector:
	docker run \
		--rm \
		-u `id -u` \
		-v $(CURDIR):/toto \
		-v ~/.composer:/.composer \
		-w /toto \
		bearstech/php-composer:7.2 \
		composer install

clean:
	rm -rf vendor
