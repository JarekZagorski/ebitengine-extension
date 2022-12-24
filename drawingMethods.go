package pixels

import (
	"image/color"

	vec "github.com/JarekZagorski/go-vectors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func FillRect(image *ebiten.Image, pos, size vec.Vf2d, clr color.Color) {
	ebitenutil.DrawRect(image, pos.X, pos.Y, size.X, size.Y, clr)
}

func DrawRect(image *ebiten.Image, pos, size vec.Vf2d, clr color.Color) {
	// north
	ebitenutil.DrawLine(image, pos.X, pos.Y, pos.X+size.X-1, pos.Y, clr)
	// east
	ebitenutil.DrawLine(image, pos.X+size.X, pos.Y, pos.X+size.X, pos.Y+size.Y, clr)
	// south
	ebitenutil.DrawLine(image, pos.X, pos.Y+size.Y, pos.X+size.X, pos.Y+size.Y, clr)
	// west
	ebitenutil.DrawLine(image, pos.X, pos.Y, pos.X, pos.Y+size.Y+1, clr)
}

type Renderable struct {
	Image     *ebiten.Image
	generator func() *ebiten.Image
}

func NewRenderable(gen func() *ebiten.Image) *Renderable {
	return &Renderable{gen(), gen}
}

func (r *Renderable) Redraw() {
	r.Image = r.generator()
}
