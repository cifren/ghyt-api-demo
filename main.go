package main

import (
	"os"
	"runtime"
    "path/filepath"
	"github.com/kataras/iris"
	herolib "github.com/kataras/iris/hero"
	. "github.com/cifren/ghyt-api/ghyt/core"
	. "github.com/cifren/ghyt-api/ghyt/core/handler"

	demoConfig "github.com/cifren/ghyt-api-demo/config"
	demoRepository "github.com/cifren/ghyt-api-demo/src/repository"
)

func main() {
	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

	def := register()

	webhookHandler := def.Handler(GhWebhookHandler)
	app.Post("/webhook-gh", webhookHandler)

	app.Run(iris.Addr(":" + os.Getenv("PORT")), iris.WithoutServerError(iris.ErrServerClosed))
}

func getPath() string {
	_, b, _, _ := runtime.Caller(0)
	return  filepath.Dir(b)
}

func register() herolib.Hero {
	def := herolib.New()

	all := make(map[string]interface{})
	all["parameters"] = demoConfig.GetParameters()
	all["jobRepository"] = getJobRepository()
	container := Container{All: all}
	container.InitContainer()
	def.Register(container)

	return *def
}

func getJobRepository() demoRepository.JobRepository {
  return demoRepository.JobRepository{}
}
