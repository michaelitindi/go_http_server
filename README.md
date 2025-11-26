# Go HTTP Server

A simple HTTP web server built with Go.

## Prerequisites

- Go 1.16 or higher installed on your system

## Installation

1. **Install Go** (if not already installed)
   - Visit [golang.org](https://golang.org/dl) and download the latest version
   - Follow the installation instructions for your operating system

2. **Clone or navigate to this repository**
   ```bash
   cd go_http_server
   ```

## Running the Server

Start the server with:

```bash
go run server.go main.go
```

You should see:
```
=== Go HTTP Server ===

Starting the HTTP server...

Server starting on http://localhost:8080
```

## Accessing the Server

Open your browser and visit:

- **Home**: http://localhost:8080/
- **About**: http://localhost:8080/about
- **Users**: http://localhost:8080/users

Or use `curl` from the command line:

```bash
curl http://localhost:8080/
curl http://localhost:8080/about
curl http://localhost:8080/users
```

## Stopping the Server

Press **Ctrl+C** in the terminal where the server is running.

## Project Files

- `main.go` - Entry point that starts the server
- `server.go` - HTTP server with route handlers
- `go.mod` - Go module file with dependencies
