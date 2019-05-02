FROM bearstech/php:7.2

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    php-tideways \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY src/xhgui-collector/ /opt/xhgui-collector/
COPY www.conf /opt/www.conf.tpl