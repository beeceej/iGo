.PHONY: clean
clean:
	@rm -rf igopb/*

.PHONY: generate
generate:
	@$(MAKE) -s clean
	@docker run \
		--rm \
		-v ${PWD}:/defs namely/protoc-all \
		-d ./proto \
		-l go \
		-o igopb  ; \

.PHONY: test
test:
	@go test ./...
