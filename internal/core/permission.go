package core

type Permission struct {
	tableName   struct{} `sql:"permissions"`
	ID          int
	Name        string
	DisplayName string
	Description string
}
