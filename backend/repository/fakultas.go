package repository

import "database/sql"

type fakultasRepository struct {
	db *sql.DB
}

func NewFakultasRepository(db *sql.DB) *fakultasRepository {
	return &fakultasRepository{db: db}

}

func (f *fakultasRepository) FetchFakultasByID(id int64) (fakultas, error) {
	var fakultas Fakultas
	err := f.db.QueryRow("SELECT * FROM fakultas WHERE id = ?", id).Scan(&fakultas.ID, fakultas.Name, &fakultas.Akreditasi)
	if err != nil {
		return fakultas, err
	}
	return fakultas, nil
}

func (f *fakultasRepository) FetchFakultasByName(fakultasName string) (fakultas, error) {
	var fakultas Fakultas
	err := f.db.QueryRow("SELECT * FROM fakultas WHERE name = ?", fakultasName).Scan(&fakultas.ID, fakultas.Name, &fakultas.Akreditasi)
	if err != nil {
		return fakultas, err
	}
	return fakultas, nil
}

func (f *fakultasRepository) FetchFakultass() ([]fakultas, error) {
	rows, err := f.db.Query("SELECT * FROM fakultas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fakultass []Fakultas
	for rows.Next() {
		var fakultas Fakultas
		err := rows.Scan(&fakultas.ID, fakultas.Name, &fakultas.Akreditasi)
		if err != nil {
			return nil, err
		}
		fakultass = append(fakultass, fakultas)
	}
	return fakultass, nil

}
