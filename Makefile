all: help

.PHONY: deps
# Update dependencies
deps:
	go mod tidy -v

.PHONY: run
# Run the application
run:
	go run .

.PHONY: help
# Show this help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
