run: main
	./main

main: go/src/main/reespel.go
	GOROOT=`pwd`/go go build main
