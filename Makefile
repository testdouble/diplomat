.PHONY: format lint watch clean e2e test generate

main: main.go */*.go
	go build -o main

format:
	@go fmt ./...

bin/golint:
	GOBIN=`pwd`/bin go get golang.org/x/lint/golint

bin/mockery:
	GOBIN=`pwd`/bin go get github.com/vektra/mockery/.../

bin/templify:
	GOBIN=`pwd`/bin go get github.com/wlbr/templify

generate: bin/mockery bin/templify
	bin/mockery -all
	go generate ./...

lint: bin/golint
	@bin/golint -set_exit_status ./...

watch:
	rg --files | entr -rc sh -c "make format && make main&& make test && make e2e && make lint"

clean:
	rm -f *.go.* */*.go.*
	rm -f main

e2e:
	@if ! curl -sS localhost:7357 &> /dev/null; then\
		echo "Starting local httpbin...";\
		docker run -d -p 7357:80 --rm --name httpbin kennethreitz/httpbin;\
		wget --spider localhost:7357 &> /dev/null;\
	fi
	# When run within `entr -r`, STDIN is closed, and `bats` really doesn't like
	# that. To assuage it, we create a "fake STDIN" with `echo`.
	# More: https://bitbucket.org/eradman/entr/commits/ec5e793ae710
	#
	# "We're professionals..."
	echo | bats --pretty test/*.bats

test:
	@go test ./...
