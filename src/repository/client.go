package Repository

import (
	"io"
	"fmt"
	"net/http"
	"net/url"
  "github.com/cifren/ghyt-api/youtrack/core"
	"github.com/cifren/ghyt-api/ghyt/core/logger"
)

type Client struct {
	Url string
	Logger logger.Logger
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

  url, err := url.Parse(this.Url + "/" + request.Endpoint)

	req, err := http.NewRequest(
		request.Method,
		url.String(),
		body,
	)

	this.Logger.Debug(fmt.Sprintf(
    "Request '%s' : %s, headers => %+v",
    request.Method,
    url.String(),
    req.Header,
  ))

	if err != nil {
	  panic(err)
	}

	res, err := this.http.Do(req)
	if err != nil {
	  panic(err)
	}

	return *res, nil
}
