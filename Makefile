dependency:
	go get -t -v ./...

test:
	echo "" > coverage.txt
	for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out; \
	done