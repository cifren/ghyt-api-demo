module github.com/cifren/ghyt-api-demo

go 1.12

require (
	github.com/cifren/ghyt-api v0.2.0
	github.com/kataras/iris v11.1.1+incompatible
	github.com/mattn/go-isatty v0.0.11 // indirect
)

replace github.com/cifren/ghyt-api => ./localvendor/ghyt-api
