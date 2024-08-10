package formater

import "unsafe"

type Item struct {
	width int
	p     *[]byte
}

func (i *Item) WriteLine(content string) {
	contextSli := unsafe.Slice(unsafe.StringData(content), len(content))
	i.WriteLineBytes(contextSli)

}

func (i *Item) WriteLineBytes(content []byte) {
	contextSli := content
	*i.p = append(*i.p, contextSli...)
	*i.p = append(*i.p, '\n')
	for j := 0; j < i.width; j++ {
		*i.p = append(*i.p, split)
	}
	*i.p = append(*i.p, '\n')

}

func (i *Item) Done() {
	*i.p = append(*i.p, '\n')
}
