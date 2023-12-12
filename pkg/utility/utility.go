package utility

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadBody(r *http.Request, task interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, task)
	if err != nil {
		return err
	}
	return nil
}
