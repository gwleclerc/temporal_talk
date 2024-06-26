package workflows

import (
	"errors"
	"net/http"
)

func errorHook(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return errors.New("invalid payload")
	}
	return errors.New("internal server error")
}
