package praktikum

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getalluser() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}
