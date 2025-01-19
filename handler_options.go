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
func ReplaceAttr(funcs ...replaceAttrFunc) Option {
	return func(o *builder) {
		o.options.ReplaceAttr = mergeReplaceAttr(funcs...)
	}
}

func mergeReplaceAttr(funcs ...replaceAttrFunc) replaceAttrFunc {
	if len(funcs) == 0 {
		return func(groups []string, a slog.Attr) slog.Attr { return a }
	}
	if len(funcs) == 1 {
		return funcs[0]
	}
	return func(groups []string, a slog.Attr) slog.Attr {
		for _, fn := range funcs {
			a = fn(groups, a)
		}
		return a
	}
}
