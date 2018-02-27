default:

clean:
	find . -type f | grep .ego.go | xargs rm

generate: http

http:
	make -C http

run: generate
	goo install ./cmd/tiny-ego
	tiny-ego

.PHONY: http
