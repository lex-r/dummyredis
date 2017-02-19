PACKAGE=github.com/lex-r/dummyredis

test:
	docker run --rm \
		-v $(CURDIR):/go/src/$(PACKAGE) \
		-w /go/src/$(PACKAGE) \
		golang:1.7.5-alpine3.5 go test
