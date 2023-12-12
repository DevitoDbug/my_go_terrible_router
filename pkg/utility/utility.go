package utility

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var pathRegex = regexp.MustCompile(`/tasks/{(\d+)}`)

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

func GetIDFromRequest(r *http.Request) (int64, error) {
	// Find matches in the URL path
	matches := pathRegex.FindStringSubmatch(r.URL.Path)

	// Check if the URL path matches the expected format
	if len(matches) < 2 {
		return 0, errors.New("not found the id from the URL path")
	}

	// Convert the matched ID to int64
	taskID, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return taskID, nil
}
