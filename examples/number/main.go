package main

import (
	"context"
	"fmt"

	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-nom/trace"
	"github.com/jtdubs/go-nom/trace/printtracer"
	"github.com/jtdubs/go-svparser/grammar"
)

func init() {
	trace.TraceSupported()
}

func main() {
	tracer := func() trace.Tracer[rune] {
		var opts printtracer.Options[rune]
		opts.IncludePackage("main")
		return opts.Tracer()
	}()
	ctx := trace.WithTracing(trace.WithTracer(context.Background(), tracer))

	c := runes.Cursor("64'h0123_4567_89ab_cdef")
	rest, result, err := grammar.Number(ctx, c)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Number: %v\n", result)
	if !rest.EOF() {
		fmt.Printf("Remaining: %q\n", string(rest.Rest()))
	}
}
