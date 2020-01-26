package fetch

import (
	"encoding/json"
	"github.com/flexicon/scrape-jj-remotes/pkg/model"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiUrl         = "https://justjoin.it/api/offers"
	apiResponseKey = "JJ_API_RESPONSE"
)

type JobFetchingService struct {
	httpClient *http.Client
	cache      *cache.Cache
}

func NewJobFetchingService() *JobFetchingService {
	return &JobFetchingService{
		httpClient: &http.Client{Timeout: time.Second * 5},
		cache:      cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (s *JobFetchingService) Fetch() ([]model.Job, error) {
	cachedResponse, found := s.cache.Get(apiResponseKey)
	if found {
		return s.interpretResponse(cachedResponse.([]byte))
	}

	res, err := s.httpClient.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return s.interpretResponse(buf)
}

func (s *JobFetchingService) interpretResponse(response []byte) ([]model.Job, error) {
	var jobs []model.Job
	err := json.Unmarshal(response, &jobs)
	if err != nil {
		return nil, err
	}

	s.cache.Set(apiResponseKey, response, cache.DefaultExpiration)

	return onlyRemoteNonB2BJobs(jobs), nil
}

func onlyRemoteNonB2BJobs(jobs []model.Job) []model.Job {
	var i int
	for _, job := range jobs {
		if job.Remote && job.EmploymentType != "b2b" {
			jobs[i] = job
			i++
		}
	}

	return jobs[:i]
}
