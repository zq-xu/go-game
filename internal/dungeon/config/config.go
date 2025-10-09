package config

import "github.com/zq-xu/go-game/assets/dungeon"

type Config struct{}

type DialogConfig struct {
	Icon string
	Box  string
}

func GetDialogConfig() *DialogConfig {
	return &DialogConfig{
		Icon: dungeon.GetDungeonImagePath("dialog/dialog_icon.png"),
		Box:  dungeon.GetDungeonImagePath("dialog/dialog_box.png"),
	}
}
