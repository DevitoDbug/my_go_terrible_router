package utility

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
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

func GetIDFromRequest(r *http.Request) (int64, error) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	if id == "" {
		return 0, errors.New("id parameter not found in the URL")
	}

	number, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error parsing id parameter: %v\n", err)
		return 0, err
	}

	return number, nil
}
