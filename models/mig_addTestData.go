package models

type addTestData {}

func (m addTestData) up(db *DB) error {
	db.Users.Create("admin","testFirst", "testLast", "testPassword")

	return nil
}

func (m addTestData) getName() string {
	return "$addTestData"
}
