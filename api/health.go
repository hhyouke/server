package api

import (
	"net/http"

	"github.com/hhyouke/server/models"
)

// HealthCheck endpoint
func (a *API) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	payload := map[string]string{
		"name":        "50tin",
		"description": "50tin is a team intellectual network",
	}

	apiResult := models.NewAPIResult(models.ErrNil, payload, "health check succeeded")

	return models.SendJSON(w, apiResult)
}
