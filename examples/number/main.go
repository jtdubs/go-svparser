package main

import (
	"fmt"

	"github.com/jtdubs/go-nom"
	"github.com/jtdubs/go-nom/printtracer"
	"github.com/jtdubs/go-nom/runes"
	"github.com/jtdubs/go-svparser/grammar"
)

func init() {
	nom.EnableTrace()
}

func main() {
	tracer := func() nom.Tracer[rune] {
		var opts printtracer.Options[rune]
		opts.IncludePackage("main")
		return opts.Tracer()
	}()

	c := runes.Cursor("64'h0123_4567_89ab_cdef").WithTracer(tracer)
	rest, result, err := grammar.Number(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Number: %v\n", result)
	if !rest.EOF() {
		fmt.Printf("Remaining: %q\n", string(rest.Rest()))
	}
}
