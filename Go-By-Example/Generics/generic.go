package main

/// MapKeys takes a map of any type and returns a slice of its keys
/// function has two type parameters - K and V; K has the comparable constraint, meaning that we can compare values of this type with the == and != operators
/// V has the any constraint, meaning that it’s not restricted in any way (any is an alias for interface{})
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
// List is a singly-linked list with values of any type
type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val T
}
//  define methods on generic types just like we do on regular types
func (lst *List[T]) Push(v, T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}