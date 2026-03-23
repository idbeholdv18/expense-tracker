package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

func RegisterRateLimiterKeyBuilder(r *http.Request) (string, error) {
	var body struct {
		Email string `json:"email"`
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return "", err
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if body.Email == "" {
		return fmt.Sprintf("email:unknown:ip:%s", host), nil
	}

	return fmt.Sprintf("email:%s:ip:%s", body.Email, host), nil
}

func LoginRateLimiterKeyBuilder(r *http.Request) (string, error) {
	var body struct {
		Login string `json:"login"`
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return "", err
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if body.Login == "" {
		return "login:unknown:" + host, nil
	}

	return fmt.Sprintf("login:%s:ip:%s", body.Login, host), nil
}
