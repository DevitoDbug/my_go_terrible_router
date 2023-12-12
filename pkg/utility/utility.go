package utility

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
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
	parsedValues, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Printf("Could not parse URL\n")
		return 0, err
	}
	value := parsedValues.Query()
	number, err := strconv.ParseInt(value.Get("id"), 10, 64)
	if err != nil {
		log.Printf("Could not find key-id in the URL\n")
		return 0, err
	}
	return number, nil
}
