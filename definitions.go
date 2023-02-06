package tagmap

type Tag int
type TagName string

const UnknownTag = Tag(-1)

type List[V any] []V

func (l *List[V]) ToIList() []interface{} {
	if l == nil {
		return []interface{}{}
	}
	out := make([]interface{}, len(*l))
	for idx := range *l {
		out[idx] = (*l)[idx]
	}
	return out
}
