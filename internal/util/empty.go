package util

import (
	"fmt"
	"reflect"
	"strings"
)

type checkChain struct {
	stack []string
}

func EmptyCheck(v any) error {
	c := checkChain{}
	return c.emptyCheck(reflect.ValueOf(v))
}

func (c *checkChain) emptyCheck(v reflect.Value) (err error) {
	if !v.IsValid() {
		return fmt.Errorf("invalid value %v", v)
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return fmt.Errorf("nil pointer %v", v)
		} else {
			c.popStack()
			return nil
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			c.pushStack(v.Type().Field(i).Name)
			if err := c.emptyCheck(v.Field(i)); err != nil {
				return err
			}
		}
		c.popStack()
		return nil
	default:
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return fmt.Errorf("%s 存在空字段或未定义 %s", c.joinStack(), v.String())
		} else {
			c.popStack()
			return nil
		}
	}
}

func (c *checkChain) joinStack() string {
	b := strings.Builder{}
	for _, s := range c.stack {
		b.WriteString(s)
		b.WriteByte('.')
	}
	return b.String()[:b.Len()-1]
}

func (c *checkChain) popStack() {
	if len(c.stack) <= 0 {
		return
	}
	c.stack = c.stack[:len(c.stack)-1]
}

func (c *checkChain) pushStack(s string) {
	c.stack = append(c.stack, s)
}
