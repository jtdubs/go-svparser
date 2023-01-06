package grammar

import (
	"reflect"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-svparser/ast"
)

// func Phrase[T any](ps ...nom.ParseFn[rune, T]) nom.ParseFn[rune, []T] {
// 	var parts []nom.ParseFn[rune, T]
// 	for _, p := range ps {
// 		parts = append(parts, nom.Preceded(nom.Many0(nom.Alt(runes.Space0(), Comment)), p))
// 	}
// 	return nom.Trace(nom.Seq(parts...))
// }

func Bind[T any](t *ast.Token, p nom.ParseFn[rune, T]) nom.ParseFn[rune, struct{}] {
	return func(start nom.Cursor[rune]) (end nom.Cursor[rune], res struct{}, err error) {
		res = struct{}{}
		end, _, err = p(start)
		if err == nil {
			t.Start = start
			t.End = end
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
