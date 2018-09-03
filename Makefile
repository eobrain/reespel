run: reespel
	./reespel -d data/cmudict-0.7b

reespel: go/src/reespel/main.go
	GOPATH=`pwd`/go go build reespel
