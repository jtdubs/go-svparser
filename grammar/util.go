package grammar

import (
	"context"
	"unicode"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/cache"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func bindSpan[T any](t *nom.Span[rune], p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return bind(t, fn.Spanning(p))
}

func bindSpanT[T ast.HasSpan, U any](t T, p nom.ParseFn[rune, U]) nom.ParseFn[rune, struct{}] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res struct{}, err error) {
		var span nom.Span[rune]
		end, span, err = fn.Spanning(p)(ctx, start)
		if err == nil {
			t.SetSpan(span)
		}
		return
	}
}

func bind[T any](t *T, p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res struct{}, err error) {
		var val T
		end, val, err = p(ctx, start)
		if err == nil {
			*t = val
		}
		return
	}
}

func to[O, I any](p nom.ParseFn[rune, I]) nom.ParseFn[rune, O] {
	return fn.Map(p, func(i I) O {
		return any(i).(O)
	})
}

func bake[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res T, err error) {
		end, res, err = p(ctx, start)
		if err == nil {
			if b, ok := any(res).(ast.Bakeable); ok {
				err = b.Bake()
			}
		}
		return
	}
}

func top[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return cache.CacheN(1, trace.TraceN(1, p))
}

func token[T ast.HasSpan, U any](t T, p nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	return bake(fn.Value(t, bindSpanT(t, p)))
}

func join(ps ...nom.ParseFn[rune, rune]) nom.ParseFn[rune, string] {
	return runes.Join(fn.Seq(ps...))
}

func concat(ps ...nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return runes.Concat(fn.Seq(ps...))
}

func parens[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return trace.Trace(
		fn.Surrounded(
			word(runes.Rune('(')),
			word(runes.Rune(')')),
			word(p),
		),
	)
}

func brackets[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return trace.Trace(
		fn.Surrounded(
			word(runes.Rune('[')),
			word(runes.Rune(']')),
			word(p),
		),
	)
}

func word[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return trace.Trace(fn.Preceded(Whitespace0, p))
}

func phrase[T any](ps ...nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	var words []nom.ParseFn[rune, T]
	for _, p := range ps {
		words = append(words, word(p))
	}
	return fn.Discard(fn.Seq(words...))
}

var comma = word(fn.Discard(runes.Rune(',')))

var asciiPrintNonWS = fn.Satisfy(func(r rune) bool {
	return r < 128 && unicode.IsPrint(r) && !unicode.IsSpace(r)
})

var alpha_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ_")
var alphanumeric_ = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_")
var alphanumeric_S = runes.OneOf("abcdefghijklmnoprqstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ0123456789_$")

func Whitespace0(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return trace.Hidden(
		top(
			fn.Many0(Whitespace),
		),
	)(ctx, start)
}

func Whitespace1(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], []ast.Whitespace, error) {
	return trace.Hidden(
		top(
			fn.Many1(Whitespace),
		),
	)(ctx, start)
}

func Whitespace(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Whitespace, error) {
	return trace.Hidden(
		top(
			fn.Alt(
				to[ast.Whitespace](Comment),
				to[ast.Whitespace](Spaces),
			),
		),
	)(ctx, start)
}

func Spaces(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.Spaces, error) {
	res := &ast.Spaces{}
	return trace.Hidden(
		top(
			bake(
				fn.Value(res,
					bindSpan(&res.Span,
						fn.Many1(whitespace),
					),
				),
			),
		),
	)(ctx, start)
}
