package npcs

import (
	"github.com/rotisserie/eris"
	"github.com/zq-xu/gotools/logx"

	"github.com/zq-xu/go-game/internal/dungeon/characters/npcs/base"
	"github.com/zq-xu/go-game/internal/dungeon/core/controls"
)

func LoadNPCs() error {
	for k, v := range base.NPCConfigSet {
		logx.Logger.Infof("Loading NPC %s", k)
		mg, err := base.NewNPC(v)
		if err != nil {
			return eris.Wrapf(err, "failed to new %s", v.Name)
		}
		controls.RegisterNPC(mg)
		logx.Logger.Infof("Loaded NPC %s", k)
	}

	return nil
}
