FROM bearstech/debian

COPY bin/xhgui-agent /usr/local/bin/xhgui-agent

USER nobody
CMD ["xhgui-agent"]