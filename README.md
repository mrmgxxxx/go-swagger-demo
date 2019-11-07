GO Project Swagger Demo
========================

# KEYS

* golang
* protobuf
* gRPC interfaces
* HTTP RESTful API (gRPC-gateway)
* swagger-ui

# Getting Started

```shell
build:
cd examples && make

run echoserver:
cd echoserver/build/echoserver && ./echoserver

run gateway:
cd gateway/build/gateway && ./gateway

swagger-ui explore:
enter "http://your.gateway.endpoint/swagger/swagger-ui/" to your web-browser;
enter "http://your.gateway.endpoint/swagger/api.swagger.json" to swagger explore;
```

# Swagger UI explore automation

```shell
modify swagger-ui/index.html file,
change SwaggerUIBundle.url to "http://your.gateway.endpoint/swagger/api.swagger.json"
```
