package notion

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	apiURL     = "https://api.notion.com"
	apiVersion = "v1"
)

const (
	currentNotionVersion = "2022-06-28"
	maxRetries           = 3
)

type errJSONDecodeFunc func(data []byte) error

// clientAPI is an internal-use Notion API Client.
type clientAPI struct {
	apiVersion    string
	notionVersion string

	token Token

	transport     http.RoundTripper
	parsedBaseURL *url.URL

	maxRetries int

	errDecoder errJSONDecodeFunc

	// used in Authorization header only for requests that require Basic authentication.
	oauthID     string
	oauthSecret string
}

// newClientAPI creates a new API Client. It's used internally, as a wrapper on HTTP mechanics.
func newClientAPI(token Token) *clientAPI {
	u, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	return &clientAPI{
		token:         token,
		transport:     http.DefaultTransport,
		parsedBaseURL: u,
		apiVersion:    apiVersion,
		notionVersion: currentNotionVersion,
		maxRetries:    maxRetries,

		errDecoder: func(data []byte) error {
			var apiErr APIError
			err := json.Unmarshal(data, &apiErr)
			if err != nil {
				return err
			}
			return &apiErr
		},
	}
}

func (c *clientAPI) request(ctx context.Context, method string, path string, params map[string]string, payload any) (*http.Response, error) {
	return c.requestRaw(ctx, method, path, params, payload, false, c.errDecoder)
}

func (c *clientAPI) requestRaw(ctx context.Context, method string, path string, params map[string]string, payload any, basicAuth bool, customErrDecoder errJSONDecodeFunc) (*http.Response, error) {
	u, err := c.parsedBaseURL.Parse(fmt.Sprintf("%s/%s", c.apiVersion, path))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if payload != nil && !reflect.ValueOf(payload).IsNil() {
		body, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}

	if len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if basicAuth {
		cred := base64.StdEncoding.EncodeToString([]byte(c.oauthID + ":" + c.oauthSecret))
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", cred))
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.String()))
	}
	req.Header.Add("Notion-Version", c.notionVersion)
	req.Header.Add("Content-Type", "application/json")

	failedAttempts := 0
	var res *http.Response
	for {
		var err error
		res, err = c.transport.RoundTrip(req.WithContext(ctx))
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusTooManyRequests {
			break
		}

		failedAttempts++
		if failedAttempts == c.maxRetries {
			return nil, &RateLimitedError{Message: fmt.Sprintf("Retry request with 429 response failed after %d retries", failedAttempts)}
		}
		// https://developers.notion.com/reference/request-limits#rate-limits
		retryAfterHeader := res.Header["Retry-After"]
		if len(retryAfterHeader) == 0 {
			return nil, &RateLimitedError{Message: "Retry-After header missing from Notion API response headers for 429 response"}
		}
		retryAfter := retryAfterHeader[0]

		waitSeconds, err := strconv.Atoi(retryAfter)
		if err != nil {
			break // should not happen
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Duration(waitSeconds) * time.Second):
		}
	}

	if res.StatusCode != http.StatusOK {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, customErrDecoder(data)
	}

	return res, nil
}
