package formater

import (
	"unicode/utf8"
	"unsafe"
)

const split = '-'

var splits = []byte{'-', '-', '-', '-', '-', '-'}

// computeWidth returns the display width of a string,
// assuming Chinese characters are double-width.
func computeWidth(s []byte) int {
	str := unsafe.String(&s[0], len(s))
	rs := []rune(str)

	width := 0
	for _, r := range rs {
		if utf8.RuneLen(r) > 1 {
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

type Builder struct {
	p []byte
}

func (b *Builder) ItemWriter(title string) Item {
	titleSli := unsafe.Slice(unsafe.StringData(title), len(title))
	width := computeWidth(titleSli) + len(splits)*2
	b.p = append(b.p, splits...)
	b.p = append(b.p, titleSli...)
	b.p = append(b.p, splits...)
	b.p = append(b.p, '\n')
	return Item{
		width: width,
		p:     &b.p,
	}
}

func (b *Builder) WriteStringItem(title string, content string) {
	contextSli := unsafe.Slice(unsafe.StringData(content), len(content))
	b.WriteBytesItem(title, contextSli)

}

func (b *Builder) WriteBytesItem(title string, content []byte) {
	titleSli := unsafe.Slice(unsafe.StringData(title), len(title))
	contextSli := content
	b.p = append(b.p, splits...)
	b.p = append(b.p, titleSli...)
	b.p = append(b.p, splits...)
	b.p = append(b.p, '\n')
	b.p = append(b.p, contextSli...)
	b.p = append(b.p, '\n')

	// aline
	width := computeWidth(titleSli) + len(splits)*2
	for i := 0; i < width; i++ {
		b.p = append(b.p, split)
	}

	b.p = append(b.p, '\n')
	b.p = append(b.p, '\n')
}

func (b *Builder) String() string {
	if len(b.p) <= 3 {
		return ""
	}
	b.p = b.p[:len(b.p)-2]
	return unsafe.String(&b.p[0], len(b.p))
}
