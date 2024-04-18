# URL Shortener Service

This is a URL shortener service created using Go. It provides a simple API to shorten URLs and retrieve original URLs from the shortened versions.

## Features

- Shorten URLs to a concise form
- Retrieve original URLs from shortened paths
- In-memory caching of URLs with Redis
- tests for key components
- Collision rate 1 in 1000


### Prerequisites

- Go (Programming language)
- Redis (For caching shortened URLs)
## Usage
```bash
{
  "long_url": "https://www.example.com"
}
```