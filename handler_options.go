package slogopts

import "log/slog"

func Level(lv slog.Leveler) Option {
	return func(o *Builder) { o.options.Level = lv }
}

func AddSource(v bool) Option {
	return func(o *Builder) { o.options.AddSource = v }
}

type ReplaceAttrFunc = func(groups []string, a slog.Attr) slog.Attr

func ReplaceAttr(fn ReplaceAttrFunc) Option {
	return func(o *Builder) { o.options.ReplaceAttr = fn }
}

func MergeReplaceAttr(funcs ...ReplaceAttrFunc) ReplaceAttrFunc {
	return func(groups []string, a slog.Attr) slog.Attr {
		for _, fn := range funcs {
			a = fn(groups, a)
		}
		return a
	}
}
