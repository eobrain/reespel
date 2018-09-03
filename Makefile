run: main
	./main -d data/cmudict-0.7b

main: go/src/main/reespel.go
	GOPATH=`pwd`/go go build main
