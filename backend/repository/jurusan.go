package repository

import "database/sql"

type jurusanRepository struct {
	db *sql.DB
}

func NewJurusanRepository(db *sql.DB) *jurusanRepository {
	return &jurusanRepository{db: db}
}

func (j *jurusanRepository) FetchJurusanByID(id int64) (jurusan, error) {
	var jurusan Jurusan
	err := j.db.QueryRow("SELECT * FROM jurusan WHERE id = ?", id).Scan(&jurusan.ID, jurusan.Name, &jurusan.Akreditasi)
	if err != nil {
		return jurusan, err
	}
	return jurusan, nil
}

func (j *jurusanRepository) FetchJurusanByName(jurusanName string) (jurusan, error) {
	var jurusan Jurusan
	err := j.db.QueryRow("SELECT * FROM jurusan WHERE name = ?", jurusanName).Scan(&jurusan.ID, jurusan.Name, &jurusan.Akreditasi)
	if err != nil {
		return jurusan, err
	}
	return jurusan, nil
}

func (j *jurusanRepository) FetchJurusans() ([]jurusan, error) {
	rows, err := j.db.Query("SELECT * FROM jurusan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jurusans []Jurusan
	for rows.Next() {
		var jurusan Jurusan
		err := rows.Scan(&jurusan.ID, jurusan.Name, &jurusan.Akreditasi)
		if err != nil {
			return nil, err
		}
		jurusans = append(jurusans, jurusan)
	}
	return jurusans, nil