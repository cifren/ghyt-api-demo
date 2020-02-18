package Repository

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	//"fmt"
	"github.com/stretchr/testify/require"
  "github.com/cifren/ghyt-api/youtrack/core"
)

func TestRun(t *testing.T) {

  client := TestClient{}
  jobRepo := JobRepository{Client: client}
  jobs := jobRepo.GetJobs()

	assert := require.New(t)
	assert.Equal(1, len(jobs))
	assert.Equal(2, len(jobs[0].Conditions))
	assert.Equal(2, len(jobs[0].Actions))

  // Conditions
	assert.Equal("equal", jobs[0].Conditions[0].Name)
	assert.Equal("event.pull_request.state", jobs[0].Conditions[0].Arguments["variableName"])
  // Actions
	assert.Equal("youtrack", jobs[0].Actions[0].To)
	assert.Equal("addTag", jobs[0].Actions[0].Name)
	assert.Equal("yt_id", jobs[0].Actions[0].Arguments["youtrackId"])
}

type TestClient struct {
}
func (this TestClient) Get(request core.Request) (http.Response, error) {
  w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	switch true {
    case request.Endpoint == "/jobs":
      io.WriteString(w, `[{"conditions":[{"id":1},{"id":2}],"actions":[{"id":1},{"id":2}],"id":1}]`)
    case request.Endpoint == "/conditions/1":
      io.WriteString(w, `{"Name":"equal","Arguments":{"variableName":"event.pull_request.state", "value": "open"},"jobId":1,"id":1}`)
    case request.Endpoint == "/conditions/2":
      io.WriteString(w, `{"Name":"regex","Arguments":{"variableName":"event.pull_request.title","value":"connect-[^-][0-9]*","persistName":"yt_id"},"jobId":1,"id":2}`)
    case request.Endpoint == "/actions/1":
      io.WriteString(w, `{"To":"youtrack","Name":"addTag","Arguments":{"youtrackId":"yt_id","tagName":"nok"},"jobId":1,"id":1}`)
    case request.Endpoint == "/actions/2":
      io.WriteString(w, `{"To":"youtrack","Name":"removeTag","Arguments":{"youtrackId":"yt_id","tagName":"nok"},"jobId":1,"id":2}`)
  }
	resp := w.Result()

  return *resp, nil
}
func(this TestClient) Post(request core.Request)(http.Response, error){
	return http.Response{}, nil
}
func(this TestClient) Request(request core.Request)(http.Response, error){
	return http.Response{}, nil
}
