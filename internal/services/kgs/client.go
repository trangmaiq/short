package kgs

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type (
	useKeyResponse struct {
		Key string `json:"key,omitempty"`
	}
	errorResponse struct {
		ErrorCode string `json:"error_code,omitempty"`
		Message   string `json:"message,omitempty"`
	}
	client struct {
		c *http.Client
	}
)

func NewClient() *client {
	return &client{c: http.DefaultClient}
}

func (c *client) UseKey(ctx context.Context) (string, error) {
	// TODO: kgs host should be configurable
	req, err := http.NewRequest(http.MethodPut, "http://localhost:9009/keys/use", nil)
	if err != nil {
		return "", err
	}

	res, err := c.c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var j errorResponse
		err = json.NewDecoder(res.Body).Decode(&j)
		if err != nil {
			return "", err
		}

		return "", errors.New(j.Message)
	}

	var keyResponse useKeyResponse
	err = json.NewDecoder(res.Body).Decode(&keyResponse)
	if err != nil {
		return "", err
	}

	return keyResponse.Key, nil
}
