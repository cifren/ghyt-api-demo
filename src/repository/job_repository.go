package repository

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"errors"
  "github.com/cifren/ghyt-api/ghyt/core/config"
  "github.com/cifren/ghyt-api/ghyt/core/logger"
  "github.com/cifren/ghyt-api/youtrack/core"
)

type JobRepository struct {
	Client core.ClientInterface
	Logger logger.Logger
}
func (this JobRepository) GetJobs() ([]config.Job, error) {
  jobs, err := this.getJobsRequest()
  fmt.Printf("jobs %+v", jobs)

  if err != nil {
    return jobs, err
  }

  return jobs, nil
}
func (this JobRepository) getJsonJobs() []byte {
  req := core.Request{
    Endpoint: "jobs",
  }

  resp, err := this.Client.Get(req)

  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  return body
}
func (this JobRepository) getJobsRequest() ([]config.Job, error) {
  jsonRequest := this.getJsonJobs()

  var jobs []config.Job

  err := json.Unmarshal(jsonRequest, &jobs)
  if err != nil {
    err := errors.New(fmt.Sprintf(
      "Une erreur s'est produite '%s', sur la route '%s' donnant une reponse '%s'",
      err,
      "/jobs",
      jsonRequest,
    ))
    this.Logger.Debug(fmt.Sprintf("%s", err))
    return jobs, err
  }

  return jobs, nil
}

