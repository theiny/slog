# SLOG - Simple Log for Go

A simple Golang package for printing colourised log statements, with an option to configure the log level to control output.

# Install

`go get -u github.com/theiny/slog`

# Example

```
package main

import (
  log "github.com/theiny/slog"
)

func main() {

	log.Debug("Some extra information that you want to hide sometimes because it's flooding your console or log storage.")

	log.Info("Some useful information")

	log.Error("An error message of some sort")
  
}
```

![log output](log.png?raw=true "Log Output screenshot")

# Set Log Level

In code:

`log.SetLevel("DEBUG")`

or via environment variable:

`LOG_LEVEL: "DEBUG"`

# Log Levels

You can suppress the log layers by setting the appropriate log level e.g. setting it to `ERROR` will only print error statements, setting it to `INFO` will print error and info statements. `DEBUG` prints everything. 

