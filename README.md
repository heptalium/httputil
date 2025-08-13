# Package httputil

The package github.com/heptalium/httputil contains helper functions for
processing HTTP requests.

## ParseRequest

ParseRequest parses HTTP requests into a struct. It supports POST form
values, JSON encoded data and XML encoded as request. The request type
is determined by the Content-Type header.

## WriteHttpStatus

WriteHttpStatus replies to the HTTP request with the status code and a
message consisting of the status code and the corresponding mes
