package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//NewRequest build a http.Request
func NewRequest(method, url string, headers map[string]string, body map[string]interface{}) (*http.Request, error) {
	requestBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}

	return request, nil
}
