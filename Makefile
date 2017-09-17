GOBUILD = go build -v 
SOURCE = github.com/rfaulhaber/fcal

all: linux mac windows freebsd

linux: fcal.go
	env GOOS=linux arch=amd64 $(GOBUILD) -o fcal-linux $(SOURCE)

mac: fcal.go
	env GOOS=darwin arch=amd64 $(GOBUILD) -o fcal-mac $(SOURCE)

windows: fcal.go
	env GOOS=windows arch=amd64 $(GOBUILD) -o fcal-windows.exe $(SOURCE)

freebsd: fcal.go
	env GOOS=freebsd arch=amd64 $(GOBUILD) -o fcal-freebsd $(SOURCE)
