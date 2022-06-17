package repository

import "database/sql"

type KampusRepository struct {
	db *sql.DB
}

func NewKampusRepository(db *sql.DB) *KampusRepository {
	return &KampusRepository{db: db}
}

func (k *KampusRepository) FetchKampusByID(id int64) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT * FROM kampus WHERE id = ?
	`
	rows, err := k.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.ID, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Info, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FetchKampusByName(name string) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT * FROM kampus WHERE name = ?
	`
	rows, err := k.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.ID, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Info, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FetchKampusByEmail(email string) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT * FROM kampus WHERE email = ?
	`
	rows, err := k.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.ID, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Info, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FetchKampusByJurusan(jurusan string) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT * FROM kampus WHERE jurusan1 = ? OR jurusan2 = ?
	`
	rows, err := k.db.Query(query, jurusan, jurusan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.ID, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Info, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (k *KampusRepository) FetchAllKampus() ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT * FROM kampus
	`
	rows, err := k.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.ID, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Info, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}
