all: README.md LICENSE_reespel.txt samples

README.md: README_standard.md go/bin/reespel
	go/bin/reespel <README_standard.md >$@

LICENSE_reespel.txt: LICENSE go/bin/reespel
	go/bin/reespel <LICENSE >$@

samples: go/bin/reespel
	go/bin/reespel <test/code.txt
	go/bin/reespel <test/udhr.txt
	go/bin/reespel <test/google-10000-english-usa.txt

go/src/diksh/dikshaneree.go: go/bin/jenerait
	go/bin/jenerait -d data/cmudict-0.7b -o $@

go/bin/reespel: go/src/reespel/main.go go/src/diksh/dikshaneree.go
	GOPATH=`pwd`/go go install reespel

go/bin/jenerait: go/src/jenerait/main.go data/cmudict-0.7b
	GOPATH=`pwd`/go go install jenerait

clean:
	rm -rf go/bin go/pkg go/src/diksh/dikshaneree.go README.md LICENSE_reespel.txt
