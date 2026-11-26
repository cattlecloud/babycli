# babycli

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/babycli.svg)](https://pkg.go.dev/cattlecloud.net/go/babycli)
[![License](https://img.shields.io/github/license/cattlecloud/babycli?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/babycli/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/babycli/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/babycli/actions/workflows/ci.yaml)

`babycli` is a modern CLI arguments parser for Go.

### Requirements

The minimum Go version is `go1.27`.

### Getting Started

The `babycli` package can be added to a project with `go get`.

```shell
go get cattlecloud.net/go/babycli@latest
```

```go
import "cattlecloud.net/go/babycli"
```

### Examples

##### Simple example

A single root-level command with a few flags. Required flags use `Require`;
flags with fallback values use `Default`.

```go
func main() {
    os.Exit(babycli.New(&babycli.Configuration{
        Arguments: babycli.Arguments(),
        Version:   "1.0.0",
        Top: &babycli.Component{
            Name: "greet",
            Help: "say hello to someone",
            Flags: babycli.Flags{
                {
                    Type:    babycli.StringFlag,
                    Long:    "name",
                    Short:   "n",
                    Require: true,
                    Help:    "name of the person to greet",
                },
                {
                    Type:    babycli.IntFlag,
                    Long:    "times",
                    Short:   "t",
                    Default: &babycli.Default{Value: 1, Show: true},
                    Help:    "number of times to repeat the greeting",
                },
                {
                    Type:  babycli.BooleanFlag,
                    Long:  "shout",
                    Short: "s",
                    Help:  "print the greeting in uppercase",
                },
            },
            Function: func(c *babycli.Component) babycli.Code {
                name := c.GetString("name")
                greeting := fmt.Sprintf("hello, %s!", name)
                if c.GetBool("shout") {
                    greeting = strings.ToUpper(greeting)
                }
                for i := 0; i < c.GetInt("times"); i++ {
                    fmt.Println(greeting)
                }
                return babycli.Success
            },
        },
    }).Run())
}
```

##### Complex example

Commands can be nested arbitrarily. Here a root command has a `serve`
subcommand and a `db` subcommand that itself has `migrate` and `seed`.
Global flags like `--config` are available to every command.

```go
func main() {
    os.Exit(babycli.New(&babycli.Configuration{
        Arguments: babycli.Arguments(),
        Version:   "1.0.0",
        Globals: babycli.Flags{
            {
                Type:  babycli.StringFlag,
                Long:  "config",
                Short: "c",
                Help:  "path to configuration file",
            },
        },
        Top: &babycli.Component{
            Name:        "mytool",
            Help:        "an example command with subcommands",
            Description: "demonstrates a multi-level command hierarchy.",
            Components: babycli.Components{
                serveCmd(),
                dbCmd(),
            },
        },
    }).Run())
}

func serveCmd() *babycli.Component {
    return &babycli.Component{
        Name: "serve",
        Help: "run the web server",
        Flags: babycli.Flags{
            {
                Type:    babycli.StringFlag,
                Long:    "bind",
                Short:   "b",
                Require: true,
                Help:    "address to bind the server to",
            },
            {
                Type:    babycli.IntFlag,
                Long:    "port",
                Short:   "p",
                Default: &babycli.Default{Value: 8080},
                Help:    "port to listen on",
            },
            {
                Type:    babycli.StringFlag,
                Long:    "middleware",
                Short:   "m",
                Repeats: true,
                Help:    "middleware to enable (repeatable)",
            },
        },
        Function: func(c *babycli.Component) babycli.Code {
            fmt.Printf("serving on %s:%d\n", c.GetString("bind"), c.GetInt("port"))
            fmt.Printf("middleware: %s\n", strings.Join(c.GetStrings("middleware"), ", "))
            return babycli.Success
        },
    }
}

func dbCmd() *babycli.Component {
    return &babycli.Component{
        Name:        "db",
        Help:        "manage the database",
        Description: "run database maintenance commands.",
        Components: babycli.Components{
            {
                Name: "migrate",
                Help: "apply pending migrations",
                Function: func(c *babycli.Component) babycli.Code {
                    fmt.Println("migrating database...")
                    return babycli.Success
                },
            },
            {
                Name: "seed",
                Help: "seed the database with sample data",
                Flags: babycli.Flags{
                    {
                        Type:  babycli.BooleanFlag,
                        Long:  "force",
                        Short: "f",
                        Help:  "overwrite existing data",
                    },
                },
                Function: func(c *babycli.Component) babycli.Code {
                    fmt.Printf("seeding database (forced=%t)...\n", c.GetBool("force"))
                    return babycli.Success
                },
            },
        },
    }
}
```

### Influence

This library was made after many projects of using [urfave/cli](https://github.com/urfave/cli),
one of the better command line parser libraries for Go. `babycli` models itself
after this library but with a cleaner, more robust implementation under the hood.
Credit of inspiration belongs to the authors.

### License

The `cattlecloud.net/go/babycli` module is open source under the [BSD](LICENSE) license.
