FROM alpine
LABEL maintainer="rsmit@daltcore.com"

COPY release-tool-linux-amd64 /root/rt
COPY entrypoint.sh /root/entrypoint.sh

WORKDIR /root

ENTRYPOINT ["/root/entrypoint.sh"]
