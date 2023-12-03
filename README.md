# HTTP Response Application

## Overview
This application is a simple HTTP server written in Go. It responds to all incoming HTTP requests with a static response code and body. These values are configurable via environment variables.

## Configuration

### Environment Variables

- `HTTP_RESPONSE_BODY`: The body of the HTTP response. If not set, the default response body is "ok".
- `HTTP_RESPONSE_CODE`: The HTTP status code to be returned in responses. If this variable is not set or is set to a non-integer value, the default response code is 200 (OK).

### Default Values

- **Default Response Body**: "ok"
- **Default Response Code**: 200

## Running the Application

Set the environment variables `HTTP_RESPONSE_BODY` and `HTTP_RESPONSE_CODE` as needed. For example, in a Unix-like shell, you can set them as follows:
   
   ```shell
   export HTTP_RESPONSE_BODY="Your Response Here"
   export HTTP_RESPONSE_CODE=404
   go run main.go
   ```

The server will start and listen on port 8080.

## Accessing the Server

Once the server is running, it can be accessed using any HTTP client (like a web browser or curl) by navigating to http://localhost:8080. The server will respond with the configured response body and status code.