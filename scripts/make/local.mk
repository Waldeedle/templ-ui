build: ## build the service
	templ generate \
	&& cd cmd/ui \
	&& GOOS=linux GOARCH=amd64 go build -ldflags "-linkmode external -extldflags -static" -gcflags="all=-N -l" -o bin/ui

run: ## run the service
	make build
	make watch-css
	air

build-css: ## build minified css
	tailwindcss -i $(ROOT_PATH)/web/static/input.css -o $(ROOT_PATH)/web/static/main.css --minify

watch-css: ## watch for css changes and rebuild css
	tailwindcss -i $(ROOT_PATH)/web/static/input.css -o $(ROOT_PATH)/web/static/main.css --watchrun

.PHONY: build-css watch-css