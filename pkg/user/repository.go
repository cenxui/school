package user

type Repository interface {
	List(...string) (string, error)
	Get(string) (string, error)
	Create(User) error
	//edit
	Edit(User) error
	Delete(string) error
}

