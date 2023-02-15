package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Performance struct {
	Params         map[string]any `json:"params"`
	Chart          map[string]any `json:"chart"`
	AppliedFilters []any          `json:"appliedFilters"`
}

type GetPerformanceRequest struct {
	ProjectId  string
	TimeBucket string
	Period     string
	From       string
	To         string
	Timezone   string
}

func (swetrixClient *SwetrixRestClient) GetPerformance(ctx context.Context, request *GetPerformanceRequest) (*Performance, error) {
	query := url.Values{}

	if request.ProjectId != "" {
		query.Add("pid", request.ProjectId)
	}

	if request.TimeBucket != "" {
		query.Add("timeBucket", request.TimeBucket)
	}

	if request.Period != "" {
		query.Add("period", request.Period)
	}

	if request.From != "" {
		query.Add("from", request.From)
	}

	if request.To != "" {
		query.Add("to", request.To)
	}

	if request.Timezone != "" {
		query.Add("timezone", request.Timezone)
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "api.swetrix.com",
		Path:     "/v1/log/performance",
		RawQuery: query.Encode(),
	}

	httpRequest, err := http.NewRequestWithContext(ctx, "GET", u.String() /*body=*/, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Add("X-API-Key", swetrixClient.ApiKey)

	response, err := swetrixClient.HttpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("status %v", response.StatusCode)
		} else {
			return nil, fmt.Errorf("status %v: %s", response.StatusCode, body)
		}
	}

	var item Performance

	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
