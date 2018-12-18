package org

type Repository interface {
	List(...string) (string, error)
	Get(string) (string, error)
	Create(Org) error
	//edit
	Edit(Org) error
	Delete(string) error
}


