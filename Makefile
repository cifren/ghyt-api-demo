dk=docker-compose
dk_run=$(dk) run --rm
## GOLANG
dkr_go=$(dk_run) go
r_go=$(dkr_go) go

## DOCKER
console:
	$(dkr_go) bash

build:
	$(dk) build go

stop:
	$(dk) down --remove-orphan

start:
	$(dk) up -d go

restart: stop start

logs:
	$(dk) logs -f

## DEV
dev@mod.clean:
	$(r_go) mod tidy

init:
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
ngrok.up:
	$(ngrok) http go:8080

c:
	$(dkr_go) env

