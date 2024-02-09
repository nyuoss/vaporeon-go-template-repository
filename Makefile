.PHONY: tools
tools:
	cd ./tools; \
	cat tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %

test: fmt
	go test -cover ./...

fmt:
	go fmt

echo:
	echo "hello"
