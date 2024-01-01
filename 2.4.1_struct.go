package main

type Robot struct {
	On bool
	Ammo int
	Power int
}

func (r *Robot) Shoot() bool {
	if r.On == false {
		return false
	} else if r.Ammo > 0 {
		r.Ammo--
		return true
	} else {
		return false
	}
}

func (r *Robot) RideBike() bool {
	if r.On == false {
		return false
	} else if r.Power > 0 {
		r.Power--
		return true
	} else {
		return false
	}
}