package controls

import "github.com/zq-xu/go-game/internal/dungeon/types"

var NpcSet = make(map[string]types.NPC, 0)

func RegisterNPC(npc types.NPC) { NpcSet[npc.Name()] = npc }

func GetNPC(name string) types.NPC { return NpcSet[name] }

func ApproachingNPC(name string) {
	npc, ok := NpcSet[name]
	if ok {
		npc.MoveNearby()
	}
}

func MoveAwaryNPC(name string) {
	npc, ok := NpcSet[name]
	if ok {
		npc.MoveAway()
	}
}
