package models

type sessionForeignKey {}

func (s sessionForeignKey) up(db *DB) error {
	db.Client.Model(&Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return nil
}

func (s sessionForeignKey) getName() string {
	return "$sessionForeignKey"
}
