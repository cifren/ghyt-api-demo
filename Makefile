dk=docker-compose
dk_run=$(dk) run --rm
dk_exec=$(dk) exec
## GOLANG
dkr_go=$(dk_run) go
dke_go=$(dk_exec) go
r_go=$(dkr_go) go
e_go=$(dke_go) go
## NGROK
dkr_ngrok=$(dk_run) ngrok
r_ngrok=$(dkr_ngrok) ngrok
## NGROK
dkr_npm=$(dk_run) npm
dke_npm=$(dk_exec) npm
r_npm=$(dkr_npm) npm
e_npm=$(dke_npm) npm

### DOCKER
dev@dk.go.console:
	$(dkr_go) bash

dev@dk.stop:
	$(dk) down --remove-orphan

dev@dk.go.build:
	$(dk) build go

dev@dk.go.stop:
	docker-compose rm -fsv go

dev@dk.go.start:
	$(dk) up -d go

dev@dk.go.restart: dev@dk.go.stop dev@dk.go.start

dev@dk.go.logs:
	$(dk) logs -f go

dev@dk.go.install: dev@go.env.file-copy dev@dk.go.stop dev@go.install-hot-reload dev@dk.go.restart

## DEV
dev@go:
	$(r_go) $(c)

# Display source code
dev@go.mod.vendor:
	$(r_go) mod vendor

dev@go.mod.tidy:
	$(r_go) mod tidy

dev@go.mod.download:
	$(e_go) mod download

dev@go.install-hot-reload:
	$(r_go) get github.com/githubnemo/CompileDaemon

# Generate main ".exe" in bin
dev@go.install-app:
	$(r_go) install main.go
	export APP_ENV=local;./main

dev@go.run-app:
	$(r_go) run main.go

# Use local code : create a folder vendor and change go.mod to use local code
# !! Do not push your go.mod anymore !!
dev@go.mod.local-dev: dev@go.git-clone-ghyt-api dev@go.mod.replace-ghyt-api

dev@go.mod.replace-ghyt-api:
	$(r_go) mod edit -replace github.com/cifren/ghyt-api=./localvendor/ghyt-api

dev@go.git-clone-ghyt-api:
	git clone git@github.com:cifren/ghyt-api.git ./localvendor/ghyt-api

dev@go.env.file-copy:
	cp .env.local.dist .env.local

test@go.install: dev@go.env.file-copy

test@go.run:
	$(r_go) test -race ./src/...

test@go.run-cover:
	$(r_go) test -cover -race ./src/...

# Ex : $ make test@run-specific file="github.com/cifren/ghyt-api/youtrack/repository"
test@go.run-specific:
	$(r_go) test -race $(p)

### NGROK
dev@ngrok.start:
	$(r_ngrok) http go:9001

### NPM
dev@dk.npm.install: dev@npm.install dev@dk.npm.restart

dev@dk.npm.start:
	$(dk) up -d npm

dev@dk.npm.stop:
	docker-compose rm -fsv npm

dev@dk.npm.restart: dev@dk.npm.stop dev@dk.npm.start

dev@dke.npm.console:
	$(dke_npm) bash

dev@dkr.npm.console:
	$(dkr_npm) bash

dev@dk.npm.console: dev@dke.npm.console

dev@dk.npm.logs:
	$(dk) logs -f npm

dev@dk.npm.build:
	$(dk) build npm

dev@npm.install:
	$(r_npm) install

dev@npm.install-pkg:
	$(e_npm) install --save $(p)

dev@npm.install-pkg-dev:
	$(e_npm) install --save-dev $(p)

dev@npm.run-prod:
	$(r_npm) run build
