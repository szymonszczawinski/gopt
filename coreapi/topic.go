package coreapi

var (
	HELLO Topic = Topic{"hello"}
)

type Topic struct {
	name string
}
