all:
	go get -u -d
	go build -ldflags="-s -w"
