package grammar

import (
	"reflect"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-svparser/ast"
)

func BindSpan[T any](t *nom.Span[rune], p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return BindValue(t, nom.Spanning(p))
}

func BindValue[T any](t *T, p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return func(start nom.Cursor[rune]) (end nom.Cursor[rune], res struct{}, err error) {
		var val T
		end, val, err = p(start)
		if err == nil {
			*t = val
		}
		return
	}
}

func To[O, I any](p nom.ParseFn[rune, I]) nom.ParseFn[rune, O] {
	return nom.Map(p, func(i I) O {
		return reflect.ValueOf(i).Interface().(O)
	})
}

func Bake[T ast.Bakeable](p nom.ParseFn[rune, T]) nom.ParseFn[rune, T] {
	return func(start nom.Cursor[rune]) (end nom.Cursor[rune], res T, err error) {
		end, res, err = p(start)
		if err == nil {
			err = res.Bake()
		}
		return
	}
}
