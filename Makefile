all: README.md LICENSE_reespel.txt samples

README.md: README_standard.md reespel
	./reespel -d data/cmudict-0.7b <README_standard.md >$@

LICENSE_reespel.txt: LICENSE reespel
	./reespel -d data/cmudict-0.7b <LICENSE >$@

samples: reespel
	./reespel -d data/cmudict-0.7b <test/code.txt
	./reespel -d data/cmudict-0.7b <test/udhr.txt
	./reespel -d data/cmudict-0.7b <test/google-10000-english-usa.txt


run: reespel
	./reespel -d data/cmudict-0.7b <input.txt

reespel: go/src/reespel/main.go
	GOPATH=`pwd`/go go build reespel
