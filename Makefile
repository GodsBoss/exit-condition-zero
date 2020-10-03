all: \
	dist/gfx.png \
	dist/index.html \
	dist/main.wasm \
	dist/wasm_exec.js

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js
	@mkdir -p dist
	cp $< $@

dist/main.wasm:
	@mkdir -p dist
	GOOS=js GOARCH=wasm go build -o $@ ./game

dist/index.html: static/index.html
	@mkdir -p dist
	cp $< $@

dist/gfx.png: gfx/gfx.xcf gfx/gfx.sh
	@mkdir -p dist
	gfx/gfx.sh $< $@

serve: all
	@echo "Serving files at http://127.0.0.1:8080/ ..."
	@echo "Press Strg+C to end ..."
	docker run \
		--rm \
		-it \
		-v ${PWD}/dist:/usr/share/nginx/html:ro \
		-v ${PWD}/misc/docker-nginx/etc/nginx/mime.types:/etc/nginx/mime.types:ro \
		-p 8080:80 \
		nginx:1.19.2

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
	dist/main.wasm \
	serve
