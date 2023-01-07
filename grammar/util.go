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

func BindSpan[T any](t *nom.Span[rune], p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return BindValue(t, fn.Spanning(p))
}

func BindValue[T any](t *T, p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res struct{}, err error) {
		var val T
		end, val, err = p(ctx, start)
		if err == nil {
			*t = val
		}
		return
	}
}

func To[O, I any](p nom.ParseFn[rune, I]) nom.ParseFn[rune, O] {
	return fn.Map(p, func(i I) O {
		return reflect.ValueOf(i).Interface().(O)
	})
}

func Bake[T ast.Bakeable](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return func(ctx context.Context, start nom.Cursor[rune]) (end nom.Cursor[rune], res T, err error) {
		end, res, err = p(ctx, start)
		if err == nil {
			err = res.Bake()
		}
		return
	}
}

func TBind[T, U any](t T, s *nom.Span[rune], p nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	if b, ok := reflect.ValueOf(t).Interface().(ast.Bakeable); ok {
		return trace.TraceN(1, To[T](Bake(fn.Value(b, BindSpan(s, p)))))
	} else {
		return trace.TraceN(1, fn.Value(t, BindSpan(s, p)))
	}
}

func TBindSeq[T, U any](t T, s *nom.Span[rune], ps ...nom.ParseFn[rune, U]) nom.ParseFn[rune, T] {
	if b, ok := reflect.ValueOf(t).Interface().(ast.Bakeable); ok {
		return trace.TraceN(1, To[T](Bake(fn.Value(b, BindSpan(s, fn.Seq(ps...))))))
	} else {
		return trace.TraceN(1, fn.Value(t, BindSpan(s, fn.Seq(ps...))))
	}
}

func TJoinSeq(ps ...nom.ParseFn[rune, rune]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Join(fn.Seq(ps...)))
}

func TConcatSeq(ps ...nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Concat(fn.Seq(ps...)))
}

func TCons(p nom.ParseFn[rune, rune], ps nom.ParseFn[rune, string]) nom.ParseFn[rune, string] {
	return trace.TraceN(1, runes.Cons(p, ps))
}
