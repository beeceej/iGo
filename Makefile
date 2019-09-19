.PHONY: clean
clean:
	@rm -rf igo/igopb/*.go

.PHONY: generate
generate: clean
	@docker run \
		--rm \
		-v ${PWD}:/defs namely/protoc-all \
		-d ./proto \
		-l go \
		-o igo/igopb

.PHONY: image
image: clean
	@docker build -t beeceej/igo:$(TAG) .

.PHONY: test
test: generate
	@go test ./...
