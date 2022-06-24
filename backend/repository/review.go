package repository

import "database/sql"

type ReviewRepository struct {
	db *sql.DB
}

type Review struct {
	ID        int64  `json:"id"`
	IdUSer    int64  `json:"id_user"`
	IdKampus  int64  `json:"id_kampus"`
	IdJurusan int64  `json:"id_jurusan"`
	Isian     string `json:"isian"`
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) FetchReviewByID(id int64) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT id, id_user, id_kampus, id_jurusan, isian FROM review WHERE id = ?
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.IdUSer, &reviewTemp.IdKampus, &reviewTemp.IdJurusan, &reviewTemp.Isian)
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
	select table review from review where id_kampus = ?
	`
	rows, err := r.db.Query(query, Kampus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.IdUSer, &reviewTemp.IdKampus, &reviewTemp.IdJurusan, &reviewTemp.Isian)
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
		SELECT * FROM review WHERE id_user = ?
	`
	rows, err := r.db.Query(query, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.IdUSer, &reviewTemp.IdKampus, &reviewTemp.IdJurusan, &reviewTemp.Isian)
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
		SELECT * FROM review WHERE isian = ? 
	`
	rows, err := r.db.Query(query, Isian)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		//review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.IdUSer, &reviewTemp.IdKampus, &reviewTemp.IdJurusan, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}
