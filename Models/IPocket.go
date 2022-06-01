package models

type Pocket interface {
	CreatePocket
}

type CreatePocket interface {
	createPocket() Pocket
}

// func create(p IPocket) {
// 	p.createPocket()
// }
