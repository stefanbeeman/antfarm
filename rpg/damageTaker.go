package rpg

type DamageTaker interface {
	TakeDamage(Damage)
	BL() int
	Shock() int
	Pain() int
}
