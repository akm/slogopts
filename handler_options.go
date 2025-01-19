package slogopts

import "log/slog"

// Set the level of the handler.
func Level(lv slog.Leveler) Option {
	return func(o *builder) { o.options.Level = lv }
}

// Set AddSource to true to add source information to the log entry.
func AddSource(v bool) Option {
	return func(o *builder) { o.options.AddSource = v }
}

type replaceAttrFunc = func(groups []string, a slog.Attr) slog.Attr

// Set ReplaceAttr function to replace the attribute.
func ReplaceAttr(fn replaceAttrFunc) Option {
	return func(o *builder) { o.options.ReplaceAttr = fn }
}

func MergeReplaceAttr(funcs ...replaceAttrFunc) replaceAttrFunc {
	return func(groups []string, a slog.Attr) slog.Attr {
		for _, fn := range funcs {
			a = fn(groups, a)
		}
		return a
	}
}
