package user

type User struct {
	name       string
	phone      string
	favourites map[string]bool
	history    map[string]Order
}

type Order struct {
	id    string
	items map[string]int
}

func NewUser(name, phone string) *User {
	return &User{
		name:       name,
		phone:      phone,
		favourites: make(map[string]bool),
		history:    make(map[string]Order),
	}
}

func (u *User) AddFavourite(name string) {
	u.favourites[name] = true
}

func (u *User) RemoveFavourite(name string) {
	u.favourites[name] = false
}

func (u *User) GetAllFavourites() []string {
	allFavourites := make([]string, 0, len(u.favourites))

	for name, favourite := range u.favourites {
		if !favourite {
			continue
		}
		allFavourites = append(allFavourites, name)
	}

	return allFavourites
}

func (u *User) Favours(name string) bool {
	return u.favourites[name]
}

func (u *User) AddHistory(order Order) {
	u.history[order.id] = order
}
