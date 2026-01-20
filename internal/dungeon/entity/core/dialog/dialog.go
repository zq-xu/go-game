package dialog

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type Dialog interface {
	DrawIcon(screen *ebiten.Image, x, y float64)
	DrawBox(screen *ebiten.Image, text string, x, y float64)
}

type dialog struct {
	icon imagekit.Image
	box  imagekit.Image
}

// NewDialog
func NewDialog() (Dialog, error) {
	var err error
	d := &dialog{}

	cfg := GetDialogConfig()
	d.icon, err = resources.NewDungeonImage(cfg.Icon)
	if err != nil {
		return nil, eris.Wrap(err, "load dialog icon image failed.")
	}
	d.box, err = resources.NewDungeonImage(cfg.Box)
	if err != nil {
		return nil, eris.Wrap(err, "load dialog box image failed.")
	}

	return d, nil
}

func (d *dialog) DrawIcon(screen *ebiten.Image, x, y float64) {
	graphics.DrawImage(screen, d.icon.Image(), x, y)
}

func (d *dialog) DrawBox(screen *ebiten.Image, text string, x, y float64) {
	// 背景宽高（你可以根据文本长度动态计算）
	w, h := 300.0, 100.0

	// 创建半透明背景
	bg := ebiten.NewImage(int(w), int(h))
	bg.Fill(color.RGBA{0, 0, 0, 180}) // 黑色，半透明

	// 绘制背景到屏幕
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(bg, op)

	graphics.DrawImage(screen, d.box.Image(), x, y)
}
