.PHONY: clean
clean:
	@rm -rf igo/igopb/*

.PHONY: generate
generate:
	@$(MAKE) -s clean
	@docker run \
		--rm \
		-v ${PWD}:/defs namely/protoc-all \
		-d ./proto \
		-l go \
		-o igo/igopb  ; \

.PHONY: test
test:
	@go test ./...
