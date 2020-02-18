package Repository

import (
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
  "github.com/cifren/ghyt-api/ghyt/core/config"
  "github.com/cifren/ghyt-api/youtrack/core"
)

type JobRepository struct {
	Client core.ClientInterface
}
func (this JobRepository) GetJobs() []config.Job {

  jobMaps := this.getJobsRequest()
  fmt.Printf("jobMaps %#v\n", jobMaps)
  jobs := []config.Job{}

  for _, jobMap := range jobMaps {

    conditions := []config.Condition{}
    for _, conditionReq := range jobMap.Conditions {
      condition := this.getCondition(conditionReq.Id)
      conditions = append(conditions, condition)
    }

    actions := []config.Action{}
    for _, actionReq := range jobMap.Actions {
      action := this.getAction(actionReq.Id)
      actions = append(actions, action)
    }

    job := config.Job{Conditions: conditions, Actions: actions}
    jobs = append(jobs, job)
  }

  return jobs
}
func (this JobRepository) getJsonJobs() []byte {
  req := core.Request{
    Endpoint: "/jobs",
  }

  resp, err := this.Client.Get(req)

  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  return body
}
func (this JobRepository) getJsonConditions(id int) []byte {
  req := core.Request{
    Endpoint: "/conditions/" + strconv.Itoa(id),
  }

  resp, err := this.Client.Get(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  return body
}
func (this JobRepository) getJsonActions(id int) []byte {
  req := core.Request{
    Endpoint: "/actions/" + strconv.Itoa(id),
  }

  resp, err := this.Client.Get(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  return body
}
func (this JobRepository) getJobsRequest() ([]JobRequest) {
  jsonRequest := this.getJsonJobs()

  var jobsRequest []JobRequest

  err := json.Unmarshal(jsonRequest, &jobsRequest)
  if err != nil {
    panic(err)
  }

  return jobsRequest
}
func (this JobRepository) getCondition(id int) config.Condition {
  jsonRequest := this.getJsonConditions(id)
  fmt.Printf("jsonRequest %#v\n", string(jsonRequest))
  condition := config.Condition{}

  err := json.Unmarshal(jsonRequest, &condition)
  if err != nil {
    panic(err)
  }

  return condition
}
func (this JobRepository) getAction(id int) config.Action {
  jsonRequest := this.getJsonActions(id)

  action := config.Action{}

  err := json.Unmarshal(jsonRequest, &action)
  if err != nil {
    panic(err)
  }

  return action
}

type JobRequest struct {
  Actions []IdStruct
  Conditions []IdStruct
  Id int
}
type IdStruct struct {
  Id int
}
