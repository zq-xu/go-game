package controls

type ApprochObject interface {
	MoveNearby()
	MoveAway()
}

var NpcSet = make(map[string]ApprochObject, 0)

func RegisterNPC(name string, npc ApprochObject) { NpcSet[name] = npc }

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
