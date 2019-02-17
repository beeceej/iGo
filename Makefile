.PHONY: clean
clean:
	@rm -rf igopb/*

.PHONY: generate
generate:
	@echo "Unfortunately since the generated code is built in docker\nwe have to chown the generated files back to the user [${USER}].\nThis requires sudo :-("
	@$(MAKE) -s clean
	@docker run -v ${PWD}:/defs namely/protoc-all \
		-d ./proto \
		-l go \
		-o igopb  ; \
		sudo chown -R "${USER}":"${USER}" igopb

.PHONY: test
test:
	@go test ./...
