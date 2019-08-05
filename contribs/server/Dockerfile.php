FROM bearstech/php:7.2

ARG uid=1001
RUN useradd xhgui_app --uid ${uid} --shell /bin/bash --home /var/www

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    php-mongodb \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY src/xhgui /opt/xhgui
COPY config.php /opt/xhgui/config/

RUN mkdir -p /var/www \
        && ln -s /opt/xhgui/webroot /var/www/web

RUN mkdir -p /opt/xhgui/cache \
    && chown xhgui_app /opt/xhgui/cache \
    && chmod 700 /opt/xhgui/cache
USER xhgui_app
