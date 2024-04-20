# workers go d1 sqlc

A Cloudflare Workers example project using Go and WebAssembly.

Libraries and tools used

- [syumai/workers](https://github.com/syumai/workers)
- [go-michi/michi](https://github.com/go-michi/michi)
- [sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)

## Requirements

- Node.js
- [wrangler](https://developers.cloudflare.com/workers/wrangler/)
  - just run `npm install -g wrangler`
- Go 1.22 or later
- [sqlc](https://github.com/sqlc-dev/sqlc)

## Development

### Getting Started

```bash
cp example.wrangler.toml wrangler.toml

# create d1 database
make create-db

# initialize d1 local database
make init-db-local

# edit wrangler.toml

# run dev server
make dev
```

### Commands

```
make create-db # create d1 database
make dev     # run dev server
make build   # build Go Wasm binary
make deploy # deploy worker
```
