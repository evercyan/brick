package xcolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	assert.NotEmpty(t, New("a").Sty(StyBold).Fg(FgRed).Bg(BgBlack).Text())
	assert.NotEmpty(t, New("a").Sty(1000).Fg(1000).Bg(1000).Text())
	New("a", "b", "c").Fg(FgGreen).Render()
}
