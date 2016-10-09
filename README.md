# vincent-demo

A simple example site using [Vincent](http://github.com/tomdionysus/vincent).

## Getting Started

If you don't have a golang development environment already, install [golang](http://golang.org).

This example is a simple HTTP server, configured using the golang `flag` package. 

### Creating the Server

To create the server, call `New` with an optional logger instance.

In [main.go](main.go):

```go
  // Create vincent server
  logger.Debug("Creating Server")
  svr, err := vincent.New(logger)
  if err!=nil { logger.Error("Cannot Create Server: %s", err); return }
```

### Templates 

Your site files should be in a directory ([templates](templates) in this example). Template files use the [handlebars](http://handlebarsjs.com/) format and should have the `.hbs` extension - all other files are served as-is. Templates are loaded using the `LoadTemplates` function on `Server`.

In [main.go](main.go):

```go
  // Load templates from the templates/ directory
  logger.Debug("Loading Templates")
  err = svr.LoadTemplates("","templates")
  if err!=nil { logger.Error("Cannot load Templates: %s", err); return }
```

### Controllers

Controllers are registered on routes and executed when a request traverses the route, before any templates are passed. All keys set in `Output` will be available to all templates under that route.

**Example**: Here, we bind a controller to the root path `/` - meaning it will be executed on all requests. We set two context variables in `Output` - `version` and `port`. 

In [main.go](main.go):

```go
  // This is an example controller
  svr.AddController("/", func(context *vincent.Context) (bool, error) {
    context.Output["version"]="1.0.1"
    context.Output["port"] = *cfg.HTTPPort
    return true, nil
  })
```

In [index.html.hbs](templates/index.html.hbs):

```handlebars
      <div class="header clearfix">
        <nav>
          <ul class="nav nav-pills pull-right">
            <li role="presentation" class="active"><a href="#">Home</a></li>
            <li role="presentation"><a href="http://github.com/tomdionysus/vincent">Github</a></li>
            <li role="presentation"><a href="https://godoc.org/github.com/tomdionysus/vincent">Godoc</a></li>
          </ul>
        </nav>
        <h3 class="text-muted">Vincent v{{ version }}</h3>
      </div>
```

### Starting the server

Starting the server causes Vincent to listen on the configured port and start serving requests. The Vincent server runs in its own goroutine, a call to `Start` will return immediately.

In [main.go](main.go):

```go
  // Start the server on the configured port
  logger.Debug("Starting HTTP Server")
  svr.Start(fmt.Sprintf(":%d", *cfg.HTTPPort))
  logger.Info("HTTP Listening on port %d", *cfg.HTTPPort)
```

## License

vincent-demo is licensed under the Open Source MIT license. Please see the [License File](LICENSE.txt) for more details.

## Code Of Conduct

The vincent project supports and enforces [The Contributor Covenant](http://contributor-covenant.org/). Please read [the code of conduct](CODE_OF_CONDUCT.md) before contributing.
