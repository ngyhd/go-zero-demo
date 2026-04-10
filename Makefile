.PHONY: bff
# generate bff -- make bff
api:
	goctl api go -api ./bff/bff.api -dir ./bff/

.PHONY: user
# generate rpc -- make user
rpc:
	goctl rpc protoc ./user/user.proto --go_out=./user --go-grpc_out=./user --zrpc_out=./user

.PHONY: post
# generate rpc -- make post
rpc:
	goctl rpc protoc ./post/post.proto --go_out=./post --go-grpc_out=./post --zrpc_out=./post

.PHONY: doc
# generate doc -- make doc
doc:
	goctl api plugin -plugin goctl-swagger="swagger -filename bff.json" -api ./bff/bff.api -dir ./doc


.PHONY: model
# generate model
# make model F=xxx.sql
model:$(F)
$(F):
	goctl model mysql ddl -src="./docs/sql/$(F)" -dir="./v1/rpc/dao" --style goZero

# ========== Version Management ==========
VERSION_FILE := version
CURRENT_VERSION := $(shell cat $(VERSION_FILE) 2>/dev/null || echo "1.0.0")
MAJOR := $(shell echo $(CURRENT_VERSION) | cut -d. -f1)
MINOR := $(shell echo $(CURRENT_VERSION) | cut -d. -f2)
PATCH := $(shell echo $(CURRENT_VERSION) | cut -d. -f3)

# Helper to get new version
NEW_PATCH := $(MAJOR).$(MINOR).$(shell echo "$$((PATCH+1))")
NEW_MINOR := $(MAJOR).$(shell echo "$$((MINOR+1))").0
NEW_MAJOR := $(shell echo "$$((MAJOR+1))").0.0

.PHONY: version
version:
	@echo "Current version: $(CURRENT_VERSION)"

# Git tag name (with 'v' prefix)
TAG_NAME := v$(shell cat $(VERSION_FILE))

.PHONY: patch
patch: ## Increment patch version and create git tag
	@echo "Incrementing patch version..."
	@echo "$(NEW_PATCH)" > $(VERSION_FILE)
	@echo "New version: $$(cat $(VERSION_FILE))"
	@git add $(VERSION_FILE)
	@git commit -m "release: v$$(cat $(VERSION_FILE))"
	@git tag -a v$$(cat $(VERSION_FILE)) -m "release: v$$(cat $(VERSION_FILE))"
	@echo "Git tag created: v$$(cat $(VERSION_FILE))"
	@git push && git push --tags
	@echo "Pushed to remote"

.PHONY: minor
minor: ## Increment minor version and create git tag
	@echo "Incrementing minor version..."
	@echo "$(NEW_MINOR)" > $(VERSION_FILE)
	@echo "New version: $$(cat $(VERSION_FILE))"
	@git add $(VERSION_FILE)
	@git commit -m "release: v$$(cat $(VERSION_FILE))"
	@git tag -a v$$(cat $(VERSION_FILE)) -m "release: v$$(cat $(VERSION_FILE))"
	@echo "Git tag created: v$$(cat $(VERSION_FILE))"
	@git push && git push --tags
	@echo "Pushed to remote"

.PHONY: major
major: ## Increment major version and create git tag
	@echo "Incrementing major version..."
	@echo "$(NEW_MAJOR)" > $(VERSION_FILE)
	@echo "New version: $$(cat $(VERSION_FILE))"
	@git add $(VERSION_FILE)
	@git commit -m "release: v$$(cat $(VERSION_FILE))"
	@git tag -a v$$(cat $(VERSION_FILE)) -m "release: v$$(cat $(VERSION_FILE))"
	@echo "Git tag created: v$$(cat $(VERSION_FILE))"
	@git push && git push --tags
	@echo "Pushed to remote"

.PHONY: tag
tag: ## Create git tag for current version (without incrementing)
	@git tag -a $(TAG_NAME) -m "release: $(TAG_NAME)"
	@echo "Git tag created: $(TAG_NAME)"
	@git push && git push --tags
	@echo "Pushed to remote"

# ========== Build ==========
.PHONY: build-bff
build-bff:
	go build -ldflags "-X main.Version=$$(cat $(VERSION_FILE))" -o bin/bff ./bff

.PHONY: build-user
build-user:
	go build -ldflags "-X main.Version=$$(cat $(VERSION_FILE))" -o bin/user ./user

.PHONY: build-post
build-post:
	go build -ldflags "-X main.Version=$$(cat $(VERSION_FILE))" -o bin/post ./post

.PHONY: build
build: build-bff build-user build-post
	@echo "Build complete. Version: $$(cat $(VERSION_FILE))"

.PHONY: run-bff
run-bff:
	go run ./bff -f etc/bff.yaml
