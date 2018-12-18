package zone

type Repository interface {
	List(...string) (string, error)
	Get(string) (string, error)
	Create(Zone) error
	//edit
	Edit(Zone) error
	Delete(string) error
}

