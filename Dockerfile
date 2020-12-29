FROM golang:latest as build
ENV GO111MODULE=on

COPY . /workdir
WORKDIR /workdir
RUN make

FROM scratch
COPY --from=build /workdir/build/daikin-aircon-exporter .
ENTRYPOINT ["daikin-aircon-exporter"]
