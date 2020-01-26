build: generate
	go build -o refdocs . 

generate: go-bindata
	$(GOBINDATA) template/

go-bindata:
ifeq (, $(shell which go-bindata))
	go get -u github.com/go-bindata/go-bindata/...
GOBINDATA=$(GOBIN)/go-bindata
else
GOBINDATA=$(shell which go-bindata)
endif
