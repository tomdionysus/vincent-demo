# vincent-demo

A simple example site using [Vincent](http://github.com/tomdionysus/vincent).

## Getting Started

If you don't have a golang development environment already, install [golang](http://golang.org).

See [main.go](main.go) on how to bootstrap a vincent project - the example is a simple HTTP server, configured using the golang `flag` package. 

### Templates 
Your site files should be in a directory ([templates](templates) in this example). Template files use the [handlebars](http://handlebarsjs.com/) format and should have the `.hbs` extension - all other files are served as-is.

### Controllers

Controllers are registered on routes and executed when a request traverses the route, before any templates are passed. All keys set in `Output` will be available to all templates under that route.

**Example**: Here, we bind a controller to the root path `/` - meaning it will be executed on all requests. We set two context variables in `Output` - `version` and `port`. 

```golang
  // This is an example controller
  svr.AddController("/", func(context *vincent.Context) (bool, error) {
    context.Output["version"]="1.0.1"
    context.Output["port"] = *cfg.HTTPPort
    return true, nil
  })
```

In the [index.html.hbs](templates/index.html.hbs) file, we use one of these variables using the standard handlebars format:

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

## License

vincent-demo is licensed under the Open Source MIT license. Please see the [License File](LICENSE.txt) for more details.

## Code Of Conduct

The vincent project supports and enforces [The Contributor Covenant](http://contributor-covenant.org/). Please read [the code of conduct](CODE_OF_CONDUCT.md) before contributing.
