package main

import (
	"os"
	"runtime"
	"fmt"
    "path/filepath"
	"github.com/kataras/iris"
	herolib "github.com/kataras/iris/hero"
	"github.com/joho/godotenv"
	. "github.com/cifren/ghyt-api/ghyt/core"
	. "github.com/cifren/ghyt-api/ghyt/core/handler"
	"github.com/cifren/ghyt-api-demo/src/config"
)

func main() {
	loadEnv()
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

func loadEnv(){
	var env string
	if os.Getenv("APP_ENV") != "" {
		env = "." + os.Getenv("APP_ENV")
	}

	filename := ".env" + env
	err := godotenv.Load(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("'%s' file not loaded but could be if created", filename))
	}
}

func register() herolib.Hero {
	def := herolib.New()

	all := make(map[string]interface{})
	all["parameters"] = config.GetParameters()
	all["jobConfig"] = config.GetJobConf()
	container := Container{All: all}
	container.InitContainer()
	def.Register(container)

	return *def
}
