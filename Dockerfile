FROM bearstech/debian

COPY bin/xhgui-agent /usr/local/bin/xhgui-agent

EXPOSE 8002
USER nobody
CMD ["xhgui-agent"]