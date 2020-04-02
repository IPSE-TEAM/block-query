FROM golang:1.12
MAINTAINER TransX team

COPY token-query /application/token-query

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /application

CMD ["./token-query"]
