package pdf

type Object interface {
	ToString() string
	GetObjectNumber() int
}
