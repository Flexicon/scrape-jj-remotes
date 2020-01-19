package main

import (
	"github.com/flexicon/scrape-jj-remotes/pkg/fetch"
	"github.com/flexicon/scrape-jj-remotes/pkg/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	jobFetchingService *fetch.JobFetchingService
}

func main() {
	r := router.New()
	h := &Handler{jobFetchingService: fetch.NewJobFetchingService()}

	r.GET("/jobs", h.JobsHandler)

	router.Run(r)
}

func (h *Handler) JobsHandler(c echo.Context) error {
	jobs, err := h.jobFetchingService.Fetch()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, jobs)
}
