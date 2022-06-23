package repository

import "database/sql"

type JurusanRepository struct {
	db *sql.DB
}

type Jurusan struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IdKampus int64  `json:"id_kampus"`
}

func NewJurusanRepository(db *sql.DB) *JurusanRepository {
	return &JurusanRepository{db: db}
}

func (j *JurusanRepository) FetchJurusanByID(id int64) ([]*Jurusan, error) {
	var jurusan []*Jurusan
	query := `
		SELECT 	id, name, id_kampus FROM jurusan WHERE id = ?
	`
	rows, err := j.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var jurusanTemp Jurusan
		err := rows.Scan(&jurusanTemp.ID, &jurusanTemp.Name, &jurusanTemp.IdKampus)
		if err != nil {
			return nil, err
		}
		jurusan = append(jurusan, &jurusanTemp)
	}
	return jurusan, nil
}

func (j *JurusanRepository) FetchJurusanByName(name string) ([]*Jurusan, error) {
	var jurusan []*Jurusan
	query := `
		SELECT 	id, name, id_kampus FROM jurusan WHERE name = ?
	`
	rows, err := j.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var jurusanTemp Jurusan
		err := rows.Scan(&jurusanTemp.ID, &jurusanTemp.Name, &jurusanTemp.IdKampus)
		if err != nil {
			return nil, err
		}
		jurusan = append(jurusan, &jurusanTemp)
	}
	return jurusan, nil
}
