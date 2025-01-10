package notion_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

// RoundTripFunc allows defining custom RoundTrip behavior.
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip executes the custom RoundTrip function.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newMockedClient returns *http.Client which responds with content from a given file.
func newMockedClient(t *testing.T, requestMockFile string, statusCode int) http.RoundTripper {
	return RoundTripFunc(func(*http.Request) *http.Response {
		file, err := os.Open(requestMockFile)
		require.NoError(t, err, "failed to open mock file")

		return &http.Response{
			StatusCode: statusCode,
			Body:       file,
			Header:     make(http.Header),
		}
	})
}

func TestRateLimit(t *testing.T) {
	t.Run("should return error when rate limit is exceeded", func(t *testing.T) {
		transport := RoundTripFunc(func(*http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"0"}},
			}
		})
		client := notion.New(
			"some_token",
			notion.WithTransport(transport),
			notion.WithRetry(2),
		)

		_, err := client.Blocks.Get(context.Background(), "some_block_id")

		assert.Error(t, err, "expected error due to rate limit")
		assert.EqualError(t, err, "Retry request with 429 response failed after 2 retries")
	})

	t.Run("should make maxRetries attempts", func(t *testing.T) {
		attempts := 0
		maxRetries := 2
		transport := RoundTripFunc(func(*http.Request) *http.Response {
			attempts++
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"0"}},
			}
		})
		client := notion.New("some_token",
			notion.WithTransport(transport),
			notion.WithRetry(maxRetries),
		)

		_, err := client.Blocks.Get(context.Background(), "some_block_id")

		assert.Error(t, err, "expected error due to rate limit")
		assert.Equal(t, maxRetries, attempts, "number of retries should match maxRetries")
	})
}

func TestBasicAuthHeader(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		assert.Equal(t, "Basic bXkgaWQgaGVyZTpzZWNyZXQgc2hoaA==", auth, "expected correct basic auth header")
		writer.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	srvURL, err := url.Parse(srv.URL)
	require.NoError(t, err, "failed to parse test server URL")

	transport := RoundTripFunc(func(req *http.Request) *http.Response {
		req.URL = srvURL
		resp, err := http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err, "failed to make HTTP request")
		return resp
	})

	opts := []notion.ClientOpt{
		notion.WithTransport(transport),
		notion.WithOAuthAppCredentials("my id here", "secret shhh"),
	}
	client := notion.New("some_token", opts...)

	_, err = client.Auth.CreateToken(context.Background(), &notion.TokenCreateRequest{})
	assert.NoError(t, err, "unexpected error during token creation")
}
