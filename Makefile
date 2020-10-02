all: dist/wasm_exec.js

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist
	cp $< $@

dist:
	mkdir -p dist

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
