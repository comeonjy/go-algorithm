package structure

type Tree struct {
	Data interface{}
	Leafs []*Tree
}
