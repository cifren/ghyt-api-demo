package Repository

import (
	"net/http"
  "github.com/cifren/ghyt-api/youtrack/core"
	"io"
)

type Client struct {
	Url string
	http http.Client
}
func (this Client) Get(request core.Request) (http.Response, error){
	request.Method = "GET"
	return this.Request(request)
}
func (this Client) Post(request core.Request) (http.Response, error){
	request.Method = "POST"
	return this.Request(request)
}
func (this Client) Request(request core.Request) (http.Response, error){
  var body io.Reader

	req, err := http.NewRequest(
		request.Method,
		request.Endpoint,
		body,
	)
	if err != nil {
	  panic(err)
	}

	res, err := this.http.Do(req)
	if err != nil {
	  panic(err)
	}

	return *res, nil
}
