package repository

import (
// 	"io/ioutil"
// 	"encoding/json"
// 	"strconv"
// 	"fmt"
// 	"errors"
  "github.com/cifren/ghyt-api/ghyt/core/logger"
  "github.com/cifren/ghyt-api/youtrack/core"
  "github.com/cifren/ghyt-api/ghyt/core/job"
)

type JobLogRepository struct {
	Client core.ClientInterface
	Logger logger.Logger
}
func (this JobLogRepository) Persist(log job.JobLog) (error) {
  request := core.Request{
    Endpoint: "logs",
    Body: log,
  }

  this.Client.Post(request)

  return nil
}
