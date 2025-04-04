# Caching Proxy Server

## Overview
This project is a caching proxy server built using Go and Redis. It forwards requests to an origin server, caches responses for faster retrieval.


## Features
- Forward Proxy: Forwards client requests to an origin server.
- Response Caching: Uses Redis to store responses for a specified time.
- Cache HIT/MISS: Displays whether the response was served from cache or the origin server.

## Tech Stack
- Go (Golang)
- Redis (for caching)

## Requirements
- Go (>=1.21)
- Docker (to run Redis in a Docker container)

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/<user_name>/proxy_server.git
    cd proxy_server
    ```

2. Install the dependencies:
    ```bash
    go mod tidy
    ```

3. Run the redis docker compose file:
   ```bash
   docker compose -f redis_docker_compose.yml up -d
    ```

4. Start the proxy server:
   ```bash
   go run server.go --port 3000 --origin http://dummyjson.com
   ```
   You can change the port no. and the orgin server url. 


## Usage

### Send a test request:
   ```bash
   curl -v http://localhost:3000/some_path
   ```
- If the response is cached, you'll see ```Cache HIT```.
- If the response is fetched from the origin, you'll see ```Cache MISS```.

Read more about this project [here](https://medium.com/@avikesar2013/caching-proxy-server-in-go-38a6630125bc)!



