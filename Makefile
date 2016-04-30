PKG = github.com/Azure/azure-sdk-for-go

godeps:
	GOPATH=`godep path` godep save ./...
	GOPATH=`godep path` go get ${PKG}

build: main.go
	godep go build -o azure-vhd-cli main.go

dockerfile: azure-vhd-cli
	docker build -t azure-vhd-cli .

clean:
	rm -rf Godeps azure-vhd-cli


all: godeps build dockerfile
