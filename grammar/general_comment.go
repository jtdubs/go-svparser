package grammar

import (
	"context"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/fn"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-svparser/ast"
)

//
// A.9.2 Comments
//

/*
 * comment ::=
 *	one_line_comment
 * | block_comment"
 */
func Comment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], ast.Comment, error) {
	return trace.Trace(fn.Alt(
		to[ast.Comment](BlockComment),
		to[ast.Comment](OneLineComment),
	))(ctx, start)
}

// block_comment ::= /* comment_text */
func BlockComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.BlockComment, error) {
	res := &ast.BlockComment{}
	return tBindSeq(res, &res.Span,
		fn.Discard(runes.Tag("/*")),
		bindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Tag("*/")),
					),
				),
			),
		),
		fn.Discard(runes.Tag("*/")),
	)(ctx, start)
}

/*
 * one_line_comment ::= // comment_text \n
 */
func OneLineComment(ctx context.Context, start nom.Cursor[rune]) (nom.Cursor[rune], *ast.OneLineComment, error) {
	res := &ast.OneLineComment{}
	return tBindSeq(res, &res.Span,
		fn.Discard(runes.Tag("//")),
		bindSpan(&res.TextT,
			runes.Join(
				fn.First(
					fn.ManyTill(
						fn.Any[rune],
						fn.Peek(runes.Newline),
					),
				),
			),
		),
		fn.Discard(runes.Newline),
	)(ctx, start)
}
