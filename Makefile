BINPATH = bin/app

.PHONY: dev
dev: $(MAKE) watch-app

.PHONY: watch-app
watch-app:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) build-app" \
	--build.bin "$(BINPATH)" \
	--build.include_ext "go" \
	--build.exclude_dir "bin,web"

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/main.go