package bcrypts

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	s := "123456"
	c := New(&Config{
		Cost: 10,
	})
	e, _ := c.Digest(s)
	fmt.Println(e)
}
