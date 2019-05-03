FROM bearstech/php:7.2

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    php-mongodb \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY src/xhgui /opt/xhgui
COPY config.php /opt/xhgui/config/

RUN mkdir -p /var/www \
        && ln -s /opt/xhgui/webroot /var/www/web

USER www-data
