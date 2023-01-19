package core

type setupRoute struct {
	find   bool
	post   bool
	put    bool
	delete bool
}

func NewSetup(find, post, put, delete bool) setupRoute {
	setup := setupRoute{find, post, put, delete}

	return setup
}
