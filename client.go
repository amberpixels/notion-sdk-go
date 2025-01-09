package notion

import "net/http"

type Client struct {
	api *clientAPI

	Auth      *AuthenticationService
	Blocks    *BlocksService
	Pages     *PagesService
	Databases *DatabasesService
	Comments  *CommentsService
	Search    *SearchService
}

// New creates a new Client. d
func New(token Token, opts ...ClientOpt) *Client {
	api := newClientAPI(token)

	c := &Client{
		api: api,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Auth = NewAuthenticationService(api)
	c.Blocks = NewBlocksService(api)
	c.Pages = NewPagesService(api)
	c.Databases = NewDatabasesService(api)
	c.Comments = NewCommentsService(api)
	c.Search = NewSearchService(api)

	return c
}

// ClientOpt to configure API client
type ClientOpt func(*Client)

// WithTransport overrides the default http.Client
func WithTransport(transport http.RoundTripper) ClientOpt {
	return func(c *Client) { c.api.transport = transport }
}

// WithVersion overrides the Notion API version
func WithVersion(version string) ClientOpt {
	return func(c *Client) { c.api.notionVersion = version }
}

// WithRetry overrides the default number of max retry attempts on 429 errors
func WithRetry(retries int) ClientOpt {
	return func(c *Client) { c.api.maxRetries = retries }
}

// WithOAuthAppCredentials sets the OAuth app ID and secret to use when fetching a token from Notion.
func WithOAuthAppCredentials(id, secret string) ClientOpt {
	return func(c *Client) {
		c.api.oauthID = id
		c.api.oauthSecret = secret
	}
}
