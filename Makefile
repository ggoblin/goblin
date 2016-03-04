BIN = $(GOPATH)/bin
BINDATA = bindata.go
BINDATA_FLAGS = -pkg=main -prefix=server/data
IMPORT_PATH   = $(shell echo `pwd` | sed "s|^$(GOPATH)/src/||g")
GIT_HASH      = $(shell git rev-parse HEAD)
APP_NAME      = $(shell echo $(IMPORT_PATH) | sed 's:.*/::')
TARGET        = $(BIN)/$(APP_NAME)
LDFLAGS       = -w -X main.commitHash=$(GIT_HASH)


build: clean $(TARGET)

clean:
	@rm -rf goblin/data/static/build/*
	@rm -rf $(BINDATA)

$(TARGET): $(BINDATA)
	@go build -ldflags '$(LDFLAGS)' -o goblin-server.bin

$(BINDATA):
	$(BIN)/go-bindata $(BINDATA_FLAGS) -o=$@ goblin/data/...
