package workflows

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DomainStatusInput struct {
	Domain string `json:"domain"`
}

type DomainStatusOutput struct {
	Status string `json:"status"`
}

func GetDomainStatus(ctx context.Context, input DomainStatusInput) (DomainStatusOutput, error) {
	var res DomainStatusOutput
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:8091/domains/%s/status", input.Domain), nil)
	if err != nil {
		return res, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if err = errorHook(resp); err != nil {
		return res, err
	}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}
