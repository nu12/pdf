package pdf

type Body struct {
	Objects []Object
}

func (b *Body) AddObject(o Object) {
	b.Objects = append(b.Objects, o)
}
