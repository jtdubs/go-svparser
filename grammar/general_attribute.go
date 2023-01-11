package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.9.1 Attributes
//

/*
 * attribute_instance ::= (* attr_spec { , attr_spec } *)
 */
func AttributeInstance(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.AttributeInstance, error) {
	res := &ast.AttributeInstance{}
	return tBindPhrase(res, &res.Span,
		fn.Discard(runes.Tag("(*")),
		fn.Bind(&res.Specs, fn.SeparatedList1(
			word(runes.Rune(',')),
			AttrSpec,
		)),
		fn.Discard(runes.Tag("*)")),
	)(ctx, start)
}

/*
 * attr_spec ::= attr_name [ = constant_expression ]
 */
func AttrSpec(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.AttrSpec, error) {
	res := &ast.AttrSpec{}
	return tBindPhrase(res, &res.Span,
		fn.Bind(&res.Name, AttrName),
		fn.Opt(phrase(
			fn.Discard(runes.Rune('=')),
			fn.Bind(&res.Expr, ConstantExpression),
		)),
	)(ctx, start)
}

/*
 * attr_name ::= identifier
 */
func AttrName(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.AttrName, error) {
	res := &ast.AttrName{}
	return tBind(res, &res.Span, fn.Bind(&res.ID, Identifier))(ctx, start)
}
