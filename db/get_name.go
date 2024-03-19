package db

type DB interface {
	GetNameByIndex(index int) string
}

func GetName(db DB, index int) string {
	return db.GetNameByIndex(index)
}

// mockgen -destination get_name_mock.go -package mock_db -source get_name.go
