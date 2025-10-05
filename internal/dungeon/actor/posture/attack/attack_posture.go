package attack

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/config"
)

type AttackPosture interface {
	Attack(direction config.Direction)

	IsAttacking() bool

	AttackPostureItem
}

type AttackPostureItem interface {
	Draw(screen *ebiten.Image, x, y float64)
	isInProgress() bool

	Direction() config.Direction
	MoveToStart()
}

var (
	attackImagesPath = map[config.Direction]string{
		config.UpDirection:    dungeon.GetDungeonImagePath("/actor/attack/attack1_up.png"),
		config.DownDirection:  dungeon.GetDungeonImagePath("/actor/attack/attack1_down.png"),
		config.LeftDirection:  dungeon.GetDungeonImagePath("/actor/attack/attack1_left.png"),
		config.RightDirection: dungeon.GetDungeonImagePath("/actor/attack/attack1_right.png"),
	}
)

type attackPosture struct {
	AttackPostureItem // the current one
	attackItemSet     map[config.Direction]AttackPostureItem
}

func NewAttackPosture() (AttackPosture, error) {
	var err error
	ap := &attackPosture{
		attackItemSet: make(map[config.Direction]AttackPostureItem),
	}

	for k, v := range attackImagesPath {
		ap.attackItemSet[k], err = newImageTableItem(k, v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load attack posture imagetable item %d", k)
		}
	}

	ap.AttackPostureItem = ap.attackItemSet[config.DefaultDirection]
	return ap, nil
}

func NewGifAttackPosture() (AttackPosture, error) {
	var err error
	ap := &attackPosture{
		attackItemSet: make(map[config.Direction]AttackPostureItem),
	}

	for k, v := range attackImagesPath {
		ap.attackItemSet[k], err = newGifItem(k, v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load attack posture gif item %d", k)
		}
	}

	ap.AttackPostureItem = ap.attackItemSet[config.DefaultDirection]
	return ap, nil
}

func (ap *attackPosture) Attack(d config.Direction) {
	if d != ap.AttackPostureItem.Direction() {
		ap.AttackPostureItem = ap.attackItemSet[d]
	}

	ap.AttackPostureItem.MoveToStart()
}

func (ap *attackPosture) IsAttacking() bool { return ap.isInProgress() }

// func (ap *attackPosture) DebugDraw(screen *ebiten.Image, x, y float64) {
// 	ap.attackItemSet[config.DownDirection].rollImages.Range(func(index int, img imagekit.Image) {
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(x+float64(50*index), y)
// 		screen.DrawImage(img.Image(), op)
// 	})
// }
