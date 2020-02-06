dk=docker-compose
dk_run=$(dk) run --rm
dk_exec=$(dk) exec
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go
## NGROK
dkr_ngrok=$(dk_run) ngrok
r_ngrok=$(dkr_ngrok) ngrok
## NGROK
dkr_npm=$(dk_run) npm
dke_npm=$(dk_exec) npm
r_npm=$(dkr_npm) npm
e_npm=$(dke_npm) npm
 
## DOCKER
dev@console:
	$(dkr_go) bash

dev@build:
	$(dk) build go

dev@stop:
	$(dk) down --remove-orphan

dev@start:
	$(dk) up -d go

dev@restart: dev@stop dev@start

dev@logs:
	$(dk) logs -f

dev@install: dev@stop dev@mod.download dev@go.install-hot-reload dev@restart

## DEV
dev@mod.clean:
	$(r_go) mod tidy

dev@mod.download:
	$(r_go) mod download

dev@go.install-hot-reload:
	$(r_go) get github.com/githubnemo/CompileDaemon

dev@init:
	$(r_go) mod init github.com/cifren/ghyt-api-demo

dev@go:
	$(r_go) $(c)

dev@go-exec:
	$(r_go) install main.go
	export APP_ENV=local;./bin/main

dev@go-run:
	$(r_go) run main.go

# Use local code : create a folder vendor and change go.mod to use local code
# !! Do not push your go.mod anymore !!
dev@mod.local-dev: dev@git.clone-ghyt-api dev@mod.replace-ghyt-api

dev@mod.replace-ghyt-api:
	$(r_go) mod edit -replace github.com/cifren/ghyt-api=./vendor/ghyt-api

dev@git.clone-ghyt-api:
	git clone git@github.com:cifren/ghyt-api.git ./vendor/ghyt-api

dev@env.file-copy:
	cp .env.local.dist .env.local

test@install: dev@env.file-copy

test@run:
	$(r_go) test -race ./main.go

test@run-cover:
	$(r_go) test -cover -race ./main.go

## SERVICES
dev@ngrok.up:
	$(r_ngrok) http go:9001

## NPM
dev@dknpm.install: dev@npm.install dev@dknpm.restart

dev@dknpm.start:
	$(dk) up -d npm

dev@dknpm.stop:
	docker-compose rm -fsv npm

dev@dknpm.restart: dev@dknpm.stop dev@dknpm.start

dev@dknpm.console:
	$(dkr_npm) bash

dev@dknpm.logs:
	$(dk) logs -f npm

dev@dknpm.build:
	$(dk) build npm

dev@npm.install:
	$(r_npm) install

dev@npm.install-pkg:
	$(e_npm) install --save $(p)

dev@npm.install-pkg-dev:
	$(e_npm) install --save-dev $(p)

dev@npm.run-prod:
	$(r_npm) run build
