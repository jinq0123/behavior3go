package core

type INodeCreator interface {
	CheckElem(name string) bool
	New(name string) (interface{}, error)
}
