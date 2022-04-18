package user

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func (u *User) GetHistory() map[string]Order {
	return u.history
}

func NewOrder(items map[string]int) Order {
	b := make([]byte, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return Order{
		id:    string(b),
		items: items,
	}
}

func (o Order) GetItems() map[string]int {
	return o.items
}
