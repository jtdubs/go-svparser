package ast

type Bakeable interface {
	Bake() error
}
