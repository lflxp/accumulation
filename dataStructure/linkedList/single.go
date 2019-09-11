package linkedList

type LiskedNode struct {
	index int
	value interface{}
	next  *LiskedNode
}

func (this *LiskedNode) GetNext() *LiskedNode {
	return this.next
}

func (this *LiskedNode) GetValue() interface{} {
	return this.value
}

func (this *LiskedNode) GetIndex() int {
	return this.index
}
