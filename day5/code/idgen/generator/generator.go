package generator

type idGenerator interface {
	Generate() int32
	Current() int32
}
