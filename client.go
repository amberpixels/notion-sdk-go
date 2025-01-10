package notion

import "net/http"

// Client is a Notion SDK client.
type Client struct {
	api *clientAPI

	Auth      *AuthenticationService
	Blocks    *BlocksService
	Pages     *PagesService
	Databases *DatabasesService
	Users     *UsersService
	Comments  *CommentsService
	Search    *SearchService
}

// Token is a type for Notion API tokens.
type Token string

// String returns the string representation of the Token.
func (t Token) String() string { return string(t) }

// New creates a new Client with the given token and options.
func New(token Token, opts ...ClientOpt) *Client {
	api := newClientAPI(token)

	c := &Client{
		api: api,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Auth = newAuthenticationService(api)
	c.Blocks = newBlocksService(api)
	c.Pages = newPagesService(api)
	c.Databases = newDatabasesService(api)
	c.Users = newUsersService(api)
	c.Comments = newCommentsService(api)
	c.Search = newSearchService(api)

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
