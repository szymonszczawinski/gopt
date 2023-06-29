package messenger

var (
	HELLO Topic = Topic{"hello"}
)

type Topic struct {
	name string
}
