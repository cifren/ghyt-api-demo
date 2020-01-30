dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go
## NGROK
dkr_ngrok=$(dk_run) ngrok
r_ngrok=$(dkr_ngrok) ngrok


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
	export APP_ENV=local;./gopath/bin/main

dev@go-run:
	$(r_go) run main.go

# Use local code : create a folder vendor and change go.mod to use local code
# !! Do not push your go.mod anymore !!
dev@mod.local-dev: dev@git.clone-ghyt-api dev@mod.replace-ghyt-api

dev@mod.replace-ghyt-api:
	$(r_go) mod edit -replace github.com/cifren/ghyt-api=./vendor/ghyt-api

dev@git.clone-ghyt-api:
	git clone git@github.com:cifren/ghyt-api.git ./vendor/ghyt-api

## SERVICES
dev@ngrok.up:
	$(r_ngrok) http go:9001
