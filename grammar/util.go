package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/cache"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func bindSpan[T any](t *nom.Span[rune], p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return bindValue(t, fn.Spanning(p))
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

func bindValue[T any](t *T, p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
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
	return cache.CacheN(2, trace.TraceN(2, p))
}

func tBind[T ast.HasSpan, U any](t T, p nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	return top(to[T](bake(fn.Value(t, bindSpanT(t, p)))))
}

func tBindSeq[T ast.HasSpan, U any](t T, ps ...nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	return top(to[T](bake(fn.Value(t, bindSpanT(t, fn.Seq(ps...))))))
}

func tBindPhrase[T ast.HasSpan, U any](t T, ps ...nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	return top(to[T](bake(fn.Value(t, bindSpanT(t, phrase(ps...))))))
}

func tJoinSeq(ps ...nom.ParseFn[rune, rune]) nom.ParseFn[rune, string] {
	return top(runes.Join(fn.Seq(ps...)))
}

func tConcatSeq(ps ...nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return top(runes.Concat(fn.Seq(ps...)))
}

func tCons(p nom.ParseFn[rune, rune], ps nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return top(runes.Cons(p, ps))
}

func tAlt[T any](ps ...nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return top(fn.Alt(ps...))
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
