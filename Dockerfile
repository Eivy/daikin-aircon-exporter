FROM golang:alpine as build
ENV GO111MODULE=on

COPY . /workdir
WORKDIR /workdir
RUN apk add make && make build/daikin-aircon-exporter

FROM alpine:latest
COPY --from=build /workdir/build/ /usr/local/bin
CMD ["daikin-aircon-exporter"]
LABEL org.opencontainers.image.source https://github.com/Eivy/daikin-aircon-exporter
