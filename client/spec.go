package client

import "fmt"

type Spec struct {
	ProjectIds []string `json:"project_ids"`
	ApiKey     string   `json:"api_key"`

	TimeBucket string `json:"time_bucket"`
	Period     string `json:"period"`

	From     string `json:"from"`
	To       string `json:"to"`
	Timezone string `json:"timezone"`

	Concurrency int `json:"concurrency"`
}

func (spec *Spec) Validate() error {
	if len(spec.ProjectIds) == 0 {
		return fmt.Errorf("at least on project_id is required")
	}

	if spec.ApiKey == "" {
		return fmt.Errorf("api_key is required")
	}

	if spec.TimeBucket == "" {
		return fmt.Errorf("time_bucket is required")
	}

	if spec.Period == "" && (spec.From == "" || spec.To == "") {
		return fmt.Errorf("either 'period' or 'from/to' is required")
	}

	return nil
}
