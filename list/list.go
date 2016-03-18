type Element struct {
  next, prev *Element
  list *List
  Value interface{}
}

func (e *Element) Next() *Element {
  if p := e.next; e.list !=nil && p != &e.list.root {
    return p
  }
  return nil
}

func (e *Element) Prev *Element {
  if p := e.prev; e.list != nil && p != &e.list.root {
    return p
  }
  return nil
}


