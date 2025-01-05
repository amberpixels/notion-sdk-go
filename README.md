<div align="center">
    <h1>Notion SDK for Go</h1>
    <p>
        <b>A simple and feature-rich Go client for the <a href="https://developers.notion.com">Notion API</a></b>
    </p>
    <br>
</div>

👋 This project is a *hard fork* of [jomei/notionapi](https://github.com/jomei/notionapi),
focusing on introducing new features, breaking changes, and enhanced API coverage.

[![GitHub tag (latest)](https://img.shields.io/github/v/tag/amberpixels/notion-sdk-go?label=go%20module)](https://github.com/amberpixels/notion-sdk-go/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/amberpixels/notion-sdk-go.svg)](https://pkg.go.dev/github.com/amberpixels/notion-sdk-go)
[![Test](https://github.com/amberpixels/notion-sdk-go/actions/workflows/go.yml/badge.svg)](https://github.com/amberpixels/notion-sdk-go/actions/workflows/go.yml)

## Supported APIs

It supports all APIs of the Notion API version `2022-06-28`.

## Installation

```bash
go get github.com/amberpixels/notion-sdk-go
```

## Usage

First, please follow the [Getting Started Guide](https://developers.notion.com/docs/getting-started) to obtain an integration token.

### Initialization

Import this library and initialize the API client using the obtained integration token.

```go
import "github.com/amberpixels/notion-sdk-go"

client := notion.NewClient("your_integration_token")
```

### Calling the API

You can use the methods of the initialized client to call the Notion API. Here is an example of how to retrieve a page:

```go
page, err := client.Page.Get(context.Background(), "your_page_id")
if err != nil {
    // Handle the error
}
```
