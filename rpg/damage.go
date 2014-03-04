package rpg

type Damage interface {
	GetType() string
	GetLevel() int
}

type BasicDamage struct {
	Type  string
	Level int
}

func (this BasicDamage) GetType() string { return this.Type }
func (this BasicDamage) GetLevel() int   { return this.Level }

type DamageLevel interface {
	GetBL() int
	GetShock() int
	GetPain() int
	GetEffect() string
}

type BasicDamageLevel struct {
	BL     int
	Shock  int
	Pain   int
	Effect string
}

func (this BasicDamageLevel) GetBL() int        { return this.BL }
func (this BasicDamageLevel) GetShock() int     { return this.Shock }
func (this BasicDamageLevel) GetPain() int      { return this.Pain }
func (this BasicDamageLevel) GetEffect() string { return this.Effect }

type DamageType interface {
	GetName() string
	GetInjuryName() string
	GetGrazeName() string
	GetDamageLevel(int) DamageLevel
}

type BasicDamageType struct {
	Name         string
	InjuryName   string
	GrazeName    string
	DamageLevels [6]DamageLevel
}

func (this BasicDamageType) GetName() string       { return this.Name }
func (this BasicDamageType) GetInjuryName() string { return this.GrazeName }
func (this BasicDamageType) GetGrazeName() string  { return this.InjuryName }
