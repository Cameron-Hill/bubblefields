package bubbleform

type Field interface {
	Value() string
	SetError(err error) Form
}
