# Why FROM scratch? Check https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07#.6bm47xu3m
FROM scratch
MAINTAINER Patrick Tsai <patrick@simplypatrick.com>

ADD ph2slack ph2slack

ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/ph2slack"]