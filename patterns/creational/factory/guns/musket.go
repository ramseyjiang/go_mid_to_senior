package guns

type musket struct {
	Gun
}

func newMusket() GunInfo {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
