package repository

import "database/sql"

type KampusRepository struct {
	db *sql.DB
}

type Kampus struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Jurusan1 string `json:"jurusan1"`
	Jurusan2 string `json:"jurusan2"`
}

func NewKampusRepository(db *sql.DB) *KampusRepository {
	return &KampusRepository{db: db}
}

func (k *KampusRepository) FetchKampusByID(id int64) ([]*Kampus, error) {
	var kampus []*Kampus
	query := `
		SELECT 	id, name, email,jurusan1, jurusan2 FROM kampus WHERE id = ?
	`
	rows, err := k.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
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
		SELECT 	id, name, email,jurusan1, jurusan2 FROM kampus WHERE name = ?
	`
	rows, err := k.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var kampusTemp Kampus
		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
		if err != nil {
			return nil, err
		}
		kampus = append(kampus, &kampusTemp)
	}
	return kampus, nil
}

func (u *KampusRepository) InsertKampus(name string, email string, jurusan1 string, jurusan2 string) error {

	_, err := u.db.Exec("INSERT INTO kampus (name string, email string,jurusan1 string, jurusan2 string) VALUES (?, ?, ?, ?)", name, email, jurusan1, jurusan2)
	if err != nil {
		return err
	}
	return nil

}

// func (k *KampusRepository) InsertKampus() ([]*Kampus, error) {
// 	var kampus []*Kampus
// 	query := `
// 		SELECT 	id, name, email,jurusan1, jurusan2  FROM kampus
// 	`
// 	rows, err := k.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var kampusTemp Kampus
// 		err := rows.Scan(&kampusTemp.Id, &kampusTemp.Name, &kampusTemp.Email, &kampusTemp.Jurusan1, &kampusTemp.Jurusan2)
// 		if err != nil {
// 			return nil, err
// 		}
// 		kampus = append(kampus, &kampusTemp)
// 	}
// 	return kampus, nil
// }
