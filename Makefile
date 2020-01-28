dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go

## DOCKER
dk@console:
	$(dkr_go) bash

dk@build:
	$(dk) build go

dk@stop:
	$(dk) down --remove-orphan

dk@start:
	$(dk) up -d go

dk@restart: dk@stop dk@start

dk@logs:
	$(dk) logs -f

## DEV
dev@mod.clean:
	$(r_go) mod tidy

dev@init:
	$(r_go) mod init github.com/cifren/ghyt-api-demo

dev@go:
	$(r_go) $(c)

dev@go-run:
	$(r_go) run main.go

dev@go-watch:
	$(dkr_go) gin --port 9000 --appPort 9001 run main.go

dev@gin-install:
	$(r_go) get -v github.com/codegangsta/gin

## SERVICES
dev@ngrok.up:
	$(ngrok) http go:8080
