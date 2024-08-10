package formater

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	var b Builder
	b.WriteStringItem("Decode//", "hello world")
	b.WriteStringItem("Encode", "testdata")
	fmt.Println(b.String())
}

func TestItem(t *testing.T) {
	var b Builder
	item := b.ItemWriter("Endode//")
	item.WriteLine("okokok")
	item.WriteLine("okokok")
	item.Done()
	item = b.ItemWriter("Decode//")
	item.WriteLine("okokok")
	item.WriteLine("okokok")
	item.Done()
	fmt.Println(b.String())
}
