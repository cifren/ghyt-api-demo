package repository

import (
	"io"
	"fmt"
	"bytes"
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
	if request.Body != nil {
		body = request.GetJsonBody()
		fmt.Println("herere")
	}

	fmt.Printf("jsonBody %#v\n", request.GetJsonBody().String())

  url, err := url.Parse(this.Url + "/" + request.Endpoint)

	req, err := http.NewRequest(
		request.Method,
		url.String(),
		body,
	)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}
	if body != nil {
		fmt.Println(fmt.Sprintf(
			"Request '%s' : %s, headers => %+v, body => %s",
			request.Method,
			url.String(),
			req.Header,
			body.(*bytes.Buffer).String(),
		))
	} else {
		fmt.Println(fmt.Sprintf(
			"Request '%s' : %s, headers => %+v",
			request.Method,
			url.String(),
			req.Header,
		))
	}

	if err != nil {
	  return http.Response{}, err
	}

	res, err := this.http.Do(req)
	if err != nil {
	  return http.Response{}, err
	}

	return *res, nil
}
