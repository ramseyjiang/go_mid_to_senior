package guns

type Ak47 struct {
	Gun
}

func newAk47() GunInfo {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
