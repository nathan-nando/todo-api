package abstraction

type Repository struct {
	Connection interface{}
	Db         interface{}
}
