package Repository

import (
  "github.com/cifren/ghyt-api/ghyt/core/config"
)

type JobRepository struct {}
func (JobRepository) GetJobs() []config.Job {
  return
}
