# URL Shortener

This project is a URL redirection service built in Go that maps specific URL paths to target URLs. It is designed as part of the Gophercises coding exercises.

## Features

* **Flexible Configuration**: Supports loading URL mappings from both YAML and JSON files.
* **CLI Interface**: Built with Cobra to provide a clean command-line interface with flags for configuration paths.
* **Dynamic Redirects**: Redirects incoming requests based on the provided mappings or returns a JSON error if the path is not found.
* **Map Merging**: Capability to merge multiple configurations from different file sources at runtime.

## Prerequisites

* **Go**: version 1.24.4 or higher.
* **Dependencies**: The project utilizes `chi` for routing, `yaml.v3` for parsing, and `cobra` for the CLI.

## Installation

1. Clone the repository to your local machine.
2. Install the required dependencies:

    ```bash
    go mod download
    ```

## Usage

You can start the redirection server by passing the path to your configuration files using CLI flags.

### Running the Server

To run the server with a YAML configuration:

```bash
go run main.go --yaml-path=cfg.yaml
```

To run the server with a JSON configuration:

```bash
go run main.go --json-path=cfg.json
```

The server starts on `localhost:8080` by default.

### CLI Flags

* `-y`, `--yaml-path`: Specify the path to a YAML file containing path-to-URL mappings.

* `-j`, `--json-path`: Specify the path to a JSON file containing path-to-URL mappings.

## API Endpoints

* `GET /`: Returns a "Hello, world!" JSON message.

* `GET /{path}`: Redirects to the mapped URL if the path exists. If the path is missing, it returns a 404 Not Found JSON response.
