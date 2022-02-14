package xconvert

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex2RGB(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			"white1",
			args{
				s: "ffffff",
			},
			255,
			255,
			255,
		},
		{
			"white2",
			args{
				s: "fff",
			},
			255,
			255,
			255,
		},
		{
			"invalid",
			args{
				s: "f",
			},
			0,
			0,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Hex2RGB(tt.args.s)
			assert.Equalf(t, tt.want, got, "Hex2RGB(%v)", tt.args.s)
			assert.Equalf(t, tt.want1, got1, "Hex2RGB(%v)", tt.args.s)
			assert.Equalf(t, tt.want2, got2, "Hex2RGB(%v)", tt.args.s)
		})
	}
}

func TestRGB2Hex(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"white",
			args{
				r: 255,
				g: 255,
				b: 255,
			},
			"ffffff",
		},
		{
			"black",
			args{
				r: 0,
				g: 0,
				b: 0,
			},
			"000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RGB2Hex(tt.args.r, tt.args.g, tt.args.b), "RGB2Hex(%v, %v, %v)", tt.args.r, tt.args.g, tt.args.b)
		})
	}
}

func TestHex2Color(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want color.Color
	}{
		{
			"white",
			args{
				s: "ffffff",
			},
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Hex2Color(tt.args.s), "Hex2Color(%v)", tt.args.s)
		})
	}
}
