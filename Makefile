
all: clean build

build: esc-kv-amd64.so esc-kv-arm.so

amd64: esc-kv-amd64.so

arm: esc-kv-arm.so

esc-kv-amd64.so:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -tags netgo -ldflags "-w -s" -o files/esc-kv-amd64.so kv.go

esc-kv-arm.so:
	CC=arm-linux-gnueabi-cc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -buildmode=plugin -a -tags netgo -ldflags "-w -s" -o files/esc-kv-arm.so kv.go

deps:
	go get github.com/patrickmn/go-cache
	go get github.com/yuin/gopher-lua
	go get github.com/esc/esc

clean:
	rm files/esc-kv*.so -f

run:
	docker-compose up
