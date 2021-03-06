BIN = $(GOPATH)/bin
NODE_BIN      = $(shell cd goblin-client ; npm bin)
BINDATA = bindata.go
BINDATA_FLAGS = -pkg=main -prefix=server/data
IMPORT_PATH   = $(shell echo `pwd` | sed "s|^$(GOPATH)/src/||g")
GIT_HASH      = $(shell git rev-parse HEAD)
APP_NAME      = $(shell echo $(IMPORT_PATH) | sed 's:.*/::')
TARGET        = $(BIN)/$(APP_NAME)/goblin
LDFLAGS       = -w -X main.commitHash=$(GIT_HASH)
CLIENTBUILD   = goblin-client/static


build: clean buildclient $(TARGET)

buildclient:
	@cd goblin-client ; npm i
	@cd goblin-client ; $(NODE_BIN)/webpack --progress --colors

clean:
	@rm -rf goblin/data/static/build/*
	@rm -rf $(BINDATA)
	@rm -rf goblin/server.bin

test: clean $(TARGET)
	@cd platform; go test -cover ./...

$(TARGET): buildclient
	@export GO15VENDOREXPERIMENT=1
	@cp -r $(CLIENTBUILD) platform/data
	@cd platform; go build -ldflags '$(LDFLAGS)' -o server.bin

$(BINDATA):
	$(BIN)/go-bindata $(BINDATA_FLAGS) -o=$@ goblin/data/...
