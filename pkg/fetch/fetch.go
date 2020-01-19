package fetch

import (
	"encoding/json"
	"github.com/flexicon/scrape-jj-remotes/pkg/model"
	"io/ioutil"
	"net/http"
	"time"
)

const apiUrl = "https://justjoin.it/api/offers"

type JobFetchingService struct {
	httpClient *http.Client
}

func NewJobFetchingService() *JobFetchingService {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	return &JobFetchingService{
		httpClient: httpClient,
	}
}

func (s *JobFetchingService) Fetch() ([]model.Job, error) {
	res, err := s.httpClient.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jobs []model.Job
	err = json.Unmarshal(buf, &jobs)
	if err != nil {
		return nil, err
	}

	return onlyRemoteJobs(jobs), nil
}

func onlyRemoteJobs(jobs []model.Job) []model.Job {
	var i int
	for _, job := range jobs {
		if job.Remote {
			jobs[i] = job
			i++
		}
	}

	return jobs[:i]
}
