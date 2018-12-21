package subneter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	fmt.Println("Hello")
}

func Test_getFirstLast(t *testing.T) {
	tests := []struct {
		name string
		num  uint8
		mask uint8
		f    uint8
		l    uint8
	}{
		{"0", 201, 1, 128, 255},
		{"1", 201, 2, 128, 255},
		{"2", 201, 3, 128, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, l := getFirstLast(tt.num, tt.mask)
			fmt.Println(f)
			fmt.Println(l)
			assert.Equal(t, tt.f, f)
			assert.Equal(t, tt.l, l)
		})
	}
}
