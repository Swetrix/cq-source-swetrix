package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LiveVisitors struct {
	LiveVisitors []any
}

type LiveVisitorsRequest struct {
	ProjectId string
}

func (swetrixClient *SwetrixRestClient) GetLiveVisitors(ctx context.Context, request *LiveVisitorsRequest) (*LiveVisitors, error) {
	query := url.Values{}

	if request.ProjectId != "" {
		query.Add("pid", request.ProjectId)
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "api.swetrix.com",
		Path:     "/v1/log/liveVisitors",
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

	var item LiveVisitors

	err = json.NewDecoder(response.Body).Decode(&item.LiveVisitors)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
