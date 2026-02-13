# URL Shortener

A simple and efficient URL shortening and redirection service built in **Go**.  
This project was developed as part of the **Gophercises** coding exercises and demonstrates clean architecture, CLI design, and database integration.

## Features

* **URL Shortening**  
  Encodes long URLs into compact 6-character Base62 short codes.

* **Redirection Service**  
  Redirects incoming requests to their original URLs.

* **CLI Interface**  
  Built with Cobra and for a clean and extensible command line experience.

* **MongoDB Integration**  
  Stores URL mappings persistently using MongoDB.

## Tech Stack

* **Go** (1.24.4+)
* **Chi** – HTTP router
* **Cobra** – CLI framework
* **MongoDB** – Database
* **Base62** – Short code encoding

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Stealthhy7512/gophercises.git
cd urlshort
```

### 2. Download dependencies

```bash
go mod download
```

## Configuration

Set the required environment variables before running the application.

```bash
MONGO_URI=<your_db_connection_string>
DATABASE_NAME=<database_name>
COLLECTION_NAME=<collection_name>
```

## Usage

To see all available commands:

```bash
go run main.go --help
```

### Available Commands

* `serve` – Start the HTTP server
* `shorten` – Generate a shortened URL

## Running the Server

Start the server:

```bash
go run main.go serve
```

Run on a custom port:

```bash
go run main.go serve --port 8000
```

By default, the server runs on:

`http://localhost:8080`

### CLI Flags

| Flag           | Description                               |
|----------------|-------------------------------------------|
| `-p`, `--port` | Port to run the server on (default: 8080) |

## Shortening a URL

To generate a short URL:

```bash
go run main.go shorten https://example.com
```

The command will output the generated short code.

## API Endpoints

### `GET /`

Returns a simple JSON health response:

```json
{
  "message": "Hello, world!"
}
```

### `GET /{shortCode}`

Redirects to the original URL if the short code exists.

If not found:

```json
{
  "error": "not found"
}
```

Status: `404 Not Found`

## Project Structure

```bash
cmd/        -> CLI commands (serve, shorten)
handler/    -> HTTP handlers
repository/ -> MongoDB interactions
service/    -> Business logic
model/      -> Data models
```

## TODO

* URL click count.
* Encoding row index instead of truncated internal ID for shorter and more consistent results.
* In memory caching for faster retrieval.
