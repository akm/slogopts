/*
slogopts is a package that provides new functions to build slog.Handler or *slog.Logger with options.

# How to create a new logger with options

The most simple way

	slogopts.New(os.Stdout)

With JSONHandler

	slogopts.New(os.Stdout, slogopts.JSON())

With TextHandler

	slogopts.New(os.Stdout, slogopts.Text())

With log level

	slogopts.New(os.Stdout, slogopts.Level(slog.LevelDebug))

With modifying key change

	replTime := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{Key: "t", Value: a.Value}
		}
		return a
	}
	slogopts.New(os.Stdout, slogopts.ReplaceAttr(replTime))

With AddSource

	slogopts.New(os.Stdout, slogopts.AddSource(true))

With multiple options

	slogopts.New(os.Stdout, slogopts.JSON(), slogopts.Level(slog.LevelDebug), slogopts.AddSource(true))

# How to create a new handler with options

You can create a handler by using slogopts.NewHandler instead of slogopts.New. The usage is the same as slogopts.New.

The most simple way

	slogopts.NewHandler(os.Stdout)

With JSONHandler

	slogopts.NewHandler(os.Stdout, slogopts.JSON())
*/
package slogopts
