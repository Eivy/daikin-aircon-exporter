# Daikin Aircon Exporter

Expose Daikin Aircon metrics to Prometheus.

## Usage
### Docker
```sh
docker run -d -p 9823:9823 -e DAIKIN_AIRCON_TARGET=192.168.0.2 ghcr.io/eivy/daikin-aircon-exporter:v0.0.1
```
