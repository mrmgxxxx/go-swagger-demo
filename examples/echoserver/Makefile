TARGET = echo-server

PWD    = $(shell pwd)
BUILD  = $(PWD)/build
BINDIR = $(BUILD)/$(TARGET)
BIN    = $(BINDIR)/$(TARGET)

GOENVPATH = $(shell go env GOPATH)

default:
	@echo -e "\e[34;1mBuilding $(TARGET)...\033[0m"
	go build -gcflags "-N -l" -o $(BIN) main.go
	@echo -e "\e[34;1mBuild $(TARGET) success!\n\033[0m"

clean:
	@rm -rf $(BUILD)
