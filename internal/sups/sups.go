package sups

const (
	KngID      = 0
	PlanetID   = 1
	ProgressID = 2
	EsscoID    = 3
	//VolgaID = 4

)

var SupNames = map[int64]string{0: "КНГ", 1: "ПЛАНЕТА", 2: "ПРОГРЕСС", 3: "ЕССКО", 4: "ВОЛГА", 5: ""}

var SupNamesLong = map[int64]string{0: "КНЯГИНЯ", 1: "ПЛАНЕТА-АВТО", 2: "ПРОГРЕСС-АВТО", 3: "ЕССКО-НН", 4: "ВОЛГА", 5: ""}

const (
	KngMask      = 1 << iota // 0001 (1)
	PlanetMask               // 0010 (2)
	ProgressMask             // 0100 (4)
	EsscoMask                // 1000 (8)
)
