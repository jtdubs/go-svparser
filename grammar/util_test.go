package grammar

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-nom/trace/printtracer"
)

type testCase[T any] struct {
	in        string
	want      T
	wantRest  string
	wantError bool
}

func validateTrace[T, U any](t *testing.T, name string, fn nom.ParseFn[rune, T], tc testCase[U]) {
	t.Helper()
	trace.TraceSupported()
	ctx := trace.WithTracing(
		trace.WithTracer(
			context.Background(),
			func() trace.Tracer[rune] {
				op := &printtracer.Options[rune]{}
				op.IncludePackage("grammar")
				op.Exclude(`grammar\.[a-z]`)
				return op.Tracer()
			}(),
		),
	)
	validateHelper(t, ctx, name, fn, tc)
}

func validate[T, U any](t *testing.T, name string, fn nom.ParseFn[rune, T], tc testCase[U]) {
	t.Helper()
	validateHelper(t, context.Background(), name, fn, tc)
}

func validateHelper[T, U any](t *testing.T, ctx context.Context, name string, fn nom.ParseFn[rune, T], tc testCase[U]) {
	t.Helper()

	c := runes.Cursor(tc.in)
	gotRest, got, err := fn(ctx, c)
	gotError := (err != nil)

	if gotError != tc.wantError {
		if tc.wantError {
			t.Errorf("%v(%q) = %v, want error", name, tc.in, got)
		} else {
			t.Errorf("%v(%q) unexpected error: %v", name, tc.in, err)
		}
		return
	}

	if string(gotRest.Rest()) != tc.wantRest {
		t.Errorf("%v(%q) rest = %q, want %q", name, tc.in, string(gotRest.Rest()), tc.wantRest)
		return
	}

	if diff := cmp.Diff(got, tc.want, cmpopts.IgnoreTypes(nom.Span[rune]{})); diff != "" {
		t.Errorf("%v(%q) = %v, want %v", name, tc.in, got, tc.want)
		return
	}
}
