package repository

import (
	"database/sql"
)

type ReviewRepository struct {
	db *sql.DB
}

type Review struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	KampusName  string `json:"kampus_name"`
	JurusanName string `json:"jurusan_name"`
	Isian       string `json:"isian"`
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) FetchReviewByID(id int64) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT id, username, kampus_name, jurusan_name, isian FROM review WHERE id = ?
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.Username, &reviewTemp.KampusName, &reviewTemp.JurusanName, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) FetchReviewByKampusID(Kampus string) ([]*Review, error) {
	var review []*Review
	query := `
	SELECT id, username, kampus_name, jurusan_name, isian FROM review WHERE id = ?
	`
	rows, err := r.db.Query(query, Kampus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.Username, &reviewTemp.KampusName, &reviewTemp.JurusanName, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) FetchReviewByUserID(UserID int64) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT id, username, kampus_name, jurusan_name, isian FROM review WHERE id = ?
	`
	rows, err := r.db.Query(query, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.Username, &reviewTemp.KampusName, &reviewTemp.JurusanName, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) FetchReviewByIsian(Isian string) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT id, username, kampus_name, jurusan_name, isian FROM review WHERE id = ?
	`
	rows, err := r.db.Query(query, Isian)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		//review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.Username, &reviewTemp.KampusName, &reviewTemp.JurusanName, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) InsertReview(username string, kampus_name string, jurusan_name string) error {

	_, err := r.db.Exec("INSERT INTO review (username string, kampus_name string, jurusan_name string) VALUES (?, ?, ?)", username, kampus_name, jurusan_name)
	if err != nil {
		return err
	}
	return nil

}

func (r *ReviewRepository) InsertcreatReview(isian string) error {

	_, err := r.db.Exec("INSERT INTO kampus (isian string) VALUES (?)", isian)
	if err != nil {
		return err
	}
	return nil

}
