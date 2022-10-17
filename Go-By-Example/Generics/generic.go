package main

// MapKeys takes a map of any type and returns a slice of its keys
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