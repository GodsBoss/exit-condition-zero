all: \
	dist/main.wasm \
	dist/wasm_exec.js

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist
	cp $< $@

dist/main.wasm:
	GOOS=js GOARCH=wasm go build -o $@ ./game

dist:
	mkdir -p dist

serve: all
	@echo "Serving files at http://127.0.0.1:8080/ ..."
	@echo "Press Strg+C to end ..."
	docker run \
		--rm \
		-it \
		-v ${PWD}:/usr/share/nginx/html:ro \
		-p 8080:80 \
		nginx:1.19.2

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
	dist/main.wasm \
	serve
