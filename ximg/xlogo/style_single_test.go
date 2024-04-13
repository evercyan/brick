package xlogo

import (
	"fmt"
	"testing"

	"github.com/evercyan/brick/xfile"
	"github.com/stretchr/testify/assert"
)

func TestStyleSingle(t *testing.T) {
	logo := New(WithRadious(1))

	fpath1 := xfile.Temp() + ".png"
	err1 := logo.Save("A", fpath1)
	assert.Nil(t, err1)
	fmt.Println("fpath1", fpath1)

	fpath2 := xfile.Temp() + ".png"
	err2 := logo.Save("革", fpath2)
	assert.Nil(t, err2)
	fmt.Println("fpath2", fpath2)
}
