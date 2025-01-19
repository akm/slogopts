# slogopts

![CI](https://github.com/akm/slogopts/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/github/akm/slogopts/graph/badge.svg?token=0Qz9JKRlIr)](https://codecov.io/github/akm/slogopts)
[![Go Report Card](https://goreportcard.com/badge/github.com/akm/slogopts)](https://goreportcard.com/report/github.com/akm/slogopts)
[![Documentation](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/akm/slogopts)
![license](https://img.shields.io/github/license/akm/slogopts)

slogopts is a package that provides new functions to build slog.Handler or \*slog.Logger with options.

## How to create a new logger with options

### The most simple way

```golang
slogopts.New(os.Stdout)
```

### With JSONHandler

```golang
slogopts.New(os.Stdout, slogopts.JSON())
```

### With TextHandler

```golang
slogopts.New(os.Stdout, slogopts.Text())
```

### With log level

```golang
slogopts.New(os.Stdout, slogopts.Level(slog.LevelDebug))
```

### With modifying key change

```golang
replTime := func(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{Key: "t", Value: a.Value}
	}
	return a
}
slogopts.New(os.Stdout, slogopts.ReplaceAttr(replTime))
```

### With AddSource

```golang
slogopts.New(os.Stdout, slogopts.AddSource(true))
```

### With multiple options

```golang
slogopts.New(os.Stdout, slogopts.JSON(), slogopts.Level(slog.LevelDebug), slogopts.AddSource(true))
```

## How to create a new handler with options

You can create a handler by using slogopts.NewHandler instead of slogopts.New. The usage is the same as slogopts.New.

### The most simple way

```golang
slogopts.NewHandler(os.Stdout)
```

### With JSONHandler

```golang
slogopts.NewHandler(os.Stdout, slogopts.JSON())
```
