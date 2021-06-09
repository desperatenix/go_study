package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	frontItem *ListItem
	backItem  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.frontItem
}

func (l *list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	nList := ListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
	if l.frontItem != nil {
		l.insertBefore(l.frontItem, &nList)
	} else {
		l.backItem = &nList
	}
	l.frontItem = &nList
	l.len++

	return &nList
}

func (l *list) PushBack(v interface{}) *ListItem {
	nList := ListItem{
		Value: v,
		Prev:  l.backItem,
		Next:  nil,
	}
	if l.backItem != nil {
		l.backItem.Next = &nList
	} else {
		l.frontItem = &nList
	}
	l.backItem = &nList
	l.len++

	return &nList
}

func (l *list) Remove(i *ListItem) {
	l.pop(i)
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	l.insertBefore(l.frontItem, l.pop(i))
	l.frontItem = i
}

func (l *list) pop(i *ListItem) *ListItem {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.frontItem = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.backItem = i.Prev
	}
	i.Next = nil
	i.Prev = nil

	return i
}

func (l *list) insertBefore(root *ListItem, nList *ListItem) {
	nList.Next = root
	if root != nil {
		root.Prev = nList
	}
}

func NewList() List {
	return new(list)
}
