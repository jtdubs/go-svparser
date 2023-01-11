package grammar

import (
	"context"
	"reflect"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

func bindSpan[T any](t *nom.Span[rune], p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return bindValue(t, fn.Spanning(p))
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
		return reflect.ValueOf(i).Interface().(O)
	})
}

func bake[T ast.Bakeable](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res T, err error) {
		end, res, err = p(ctx, start)
		if err == nil {
			err = res.Bake()
		}
		return
	}
}

func tBind[T, U any](t T, s *nom.Span[rune], p nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	if b, ok := reflect.ValueOf(t).Interface().(ast.Bakeable); ok {
		return trace.TraceN(1, to[T](bake(fn.Value(b, bindSpan(s, p)))))
	} else {
		return trace.TraceN(1, fn.Value(t, bindSpan(s, p)))
	}
}

func tBindSeq[T, U any](t T, s *nom.Span[rune], ps ...nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	if b, ok := reflect.ValueOf(t).Interface().(ast.Bakeable); ok {
		return trace.TraceN(1, to[T](bake(fn.Value(b, bindSpan(s, fn.Seq(ps...))))))
	} else {
		return trace.TraceN(1, fn.Value(t, bindSpan(s, fn.Seq(ps...))))
	}
}

func tBindPhrase[T, U any](t T, s *nom.Span[rune], ps ...nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	if b, ok := reflect.ValueOf(t).Interface().(ast.Bakeable); ok {
		return trace.TraceN(1, to[T](bake(fn.Value(b, bindSpan(s, phrase(ps...))))))
	} else {
		return trace.TraceN(1, fn.Value(t, bindSpan(s, phrase(ps...))))
	}
}

func tJoinSeq(ps ...nom.ParseFn[rune, rune]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Join(fn.Seq(ps...)))
}

func tConcatSeq(ps ...nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Concat(fn.Seq(ps...)))
}

func tCons(p nom.ParseFn[rune, rune], ps nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Cons(p, ps))
}

func parens[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return fn.Surrounded(
		word(runes.Rune('(')),
		word(runes.Rune(')')),
		word(p))
}

func word[T any](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return fn.Preceded(Whitespace0, p)
}

func phrase[T any](ps ...nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	var words []nom.ParseFn[rune, T]
	for _, p := range ps {
		words = append(words, word(p))
	}
	return fn.Discard(fn.Seq(words...))
}
