# URL Shortener

This project is a URL shortening/redirection service built in Go that maps specific URL paths to target URLs. It is designed as part of the Gophercises coding exercises.

## Features

* **URL Shortening:** Encodes a given URL to 6 character codes.
* **CLI Interface:** Built with Cobra to provide a clean command-line interface with flags for configuration paths.
* **Dynamic Redirects:** Redirects incoming requests based on the provided mappings or returns a JSON error if the path is not found.

## Prerequisites

* **Go**: version 1.24.4 or higher.
* **Dependencies**: The project utilizes `chi` for routing, `base62` for encoding, `cobra` for the CLI and `mongodb` for database.

## Installation

1. Clone the repository to your local machine.
2. Install the required dependencies:

    ```bash
    go mod download
    ```

## Usage

Make sure to add database variables to your environment

```bash
MONGO_URI = <your_db_connection_string>
DATABASE_NAME = <database_name>
COLLECTION_NAME = <table_name>
```

You can start by typing:

```bash
go run main.go --help
```

**Available commands**

* `serve`
* `shorten`

### Running the Server

To run the server for redirections:

```bash
go run main.go serve
```

To run the server with a specific port:

```bash
go run main.go serve --port 8000
```

The server starts on `localhost:8080` by default.

**CLI Flags**

* `-p`, `--port`: Specify the port to run the server on.

### URL Shortening

To shorten a URL

```bash
go run main.go shorten [url]
```

## API Endpoints

* `GET /`: Returns a "Hello, world!" JSON message.

* `GET /{path}`: Redirects to the mapped URL if the path exists. If the path is missing, it returns a 404 Not Found JSON response.
