package factory

import "TodoAPI/database"

type Factory struct {
	Db interface{}
}

func NewFactory(dbType string) *Factory {
	f := &Factory{}

	f.setupDB(dbType)
	f.SetupRepository()

	return f
}

func (f *Factory) setupDB(dbType string) {
	//DB TYPE FROM ENV TERMINAL
	db, err := database.Connection(dbType)

	if err != nil {
		panic("Failed setup db, connection is undefined")
	}

	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}
}
