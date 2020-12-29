FROM golang:latest as build
ENV GO111MODULE=on

COPY . /go/src/github.com/eivy/daikin-aircon-exporter
WORKDIR /go/src/github.com/eivy/daikin-aircon-exporter
RUN go build

FROM scratch
COPY --from=build /daikin-aircon-exporter/daikin-aircon-exporter .
ENTRYPOINT ["daikin-aircon-exporter"]
