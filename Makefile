GOBUILD = go build -v github.com/rfaulhaber/fcal

all: linux mac windows

linux: fcal.go
	env GOOS=linux arch=amd64 $(GOBUILD)

mac: fcal.go
	env GOOS=darwin arch=amd64 $(GOBUILD)

windows: fcal.go
	env GOOS=windows arch=amd64 $(GOBUILD)
