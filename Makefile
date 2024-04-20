.PHONY: gen
gen:
	sqlc generate

.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.24.0 -mode=go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm .

.PHONY: deploy
deploy:
	wrangler deploy

.PHONY: create-db
create-db:
	wrangler d1 create d1-todo-server

.PHONY: init-db
init-db:
	wrangler d1 execute d1-todo-server --file=./schema.sql

.PHONY: init-db-local
init-db-local:
	wrangler d1 execute d1-todo-server --file=./schema.sql --local
