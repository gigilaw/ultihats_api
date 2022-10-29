package validation

var UpsertDiscSkills struct {
	PrimaryRole         int    `binding:"omitempty,numeric"`
	Throwing            int    `binding:"omitempty,numeric"`
	Catching            int    `binding:"omitempty,numeric"`
	OffensiveStrategies int    `binding:"omitempty,numeric"`
	DefensiveStrategies int    `binding:"omitempty,numeric"`
	Public              string `binding:"omitempty,boolean"`
}
