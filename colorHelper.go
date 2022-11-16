package pixels

import (
	"image/color"
)

func NewCol(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

var (
	GREY              = NewCol(192, 192, 192, 255)
	DARK_GREY         = NewCol(128, 128, 128, 255)
	VERY_DARK_GREY    = NewCol(64, 64, 64, 255)
	RED               = NewCol(255, 0, 0, 255)
	DARK_RED          = NewCol(128, 0, 0, 255)
	VERY_DARK_RED     = NewCol(64, 0, 0, 255)
	YELLOW            = NewCol(255, 255, 0, 255)
	DARK_YELLOW       = NewCol(128, 128, 0, 255)
	VERY_DARK_YELLOW  = NewCol(64, 64, 0, 255)
	GREEN             = NewCol(0, 255, 0, 255)
	DARK_GREEN        = NewCol(0, 128, 0, 255)
	VERY_DARK_GREEN   = NewCol(0, 64, 0, 255)
	CYAN              = NewCol(0, 255, 255, 255)
	DARK_CYAN         = NewCol(0, 128, 128, 255)
	VERY_DARK_CYAN    = NewCol(0, 64, 64, 255)
	BLUE              = NewCol(0, 0, 255, 255)
	DARK_BLUE         = NewCol(0, 0, 128, 255)
	VERY_DARK_BLUE    = NewCol(0, 0, 64, 255)
	MAGENTA           = NewCol(255, 0, 255, 255)
	DARK_MAGENTA      = NewCol(128, 0, 128, 255)
	VERY_DARK_MAGENTA = NewCol(64, 0, 64, 255)
	WHITE             = NewCol(255, 255, 255, 255)
	BLACK             = NewCol(0, 0, 0, 255)
	BLANK             = NewCol(0, 0, 0, 0)
)

func Alpha(c color.RGBA, a uint8) *color.RGBA {
	c.A = a
	return &c
}
