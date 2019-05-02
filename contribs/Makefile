build: src/xhgui-collector
	docker build \
		-t php-with-xhgui-agent \
		.

src/xhgui-collector: src
	cd src && git clone https://github.com/perftools/xhgui-collector

src:
	mkdir -p src

clean:
	rm -rf src