
# Admin

[![Build Status](https://travis-ci.org/natebrennand/admin.svg)](https://travis-ci.org/natebrennand/admin)

Admin adds an HTTP listener to port 8001 (configurable by the environment variable `ADMIN_PORT`).

## Using it

Admin adds listeners and starts a HTTP listener using the `init` function.
Simply import and ignore the package.

```go
import _ "github.com/natebrennand/admin" // add healthcheck endpoint
```


## Endpoints

### /ping

Returns `"pong"`.

### /stats

Returns several stats of the current memory consumption of the application.



