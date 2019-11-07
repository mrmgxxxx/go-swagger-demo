TARGET = gateway

PWD    = $(shell pwd)
BUILD  = $(PWD)/build
BINDIR = $(BUILD)/$(TARGET)
BIN    = $(BINDIR)/$(TARGET)

GOENVPATH = $(shell go env GOPATH)
PROTOCOL  = $(GOENVPATH)/src/go-swagger-demo/examples/protocol

SWAGGEDIR = $(BINDIR)/swagger
SWAGGERUI = $(GOENVPATH)/src/go-swagger-demo/third_party/swagger-ui
APIDOCS   = $(PROTOCOL)/echoserver/echoserver.swagger.json

default:
	@echo -e "\e[34;1mBuilding $(TARGET)...\033[0m"
	go build -gcflags "-N -l" -o $(BIN) main.go
	@mkdir -p $(SWAGGEDIR) && cp -r $(SWAGGERUI) $(SWAGGEDIR) && cp -r $(APIDOCS) $(SWAGGEDIR)/api.swagger.json
	@echo -e "\e[34;1mBuild $(TARGET) success!\n\033[0m"

clean:
	@rm -rf $(BUILD)
