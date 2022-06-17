package repository

import "database/sql"

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) FetchReviewByID(id int64) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT * FROM review WHERE id = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 10 OFFSET 0
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.KampusID, &reviewTemp.UserID, &reviewTemp.Jurusan, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) FetchReviewByKampusID(Jurusan string) ([]*Review, error) {
	var review []*Review
	query := `
	select * from review where nama_kampus = 'Universitas Indonesia' and jurusan = 'Teknik'
	`
	rows, err := r.db.Query(query, Jurusan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.KampusID, &reviewTemp.UserID, &reviewTemp.Jurusan, &reviewTemp.Isian)
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
		SELECT * FROM review WHERE user_id = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 10 OFFSET 0
	`
	rows, err := r.db.Query(query, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.KampusID, &reviewTemp.UserID, &reviewTemp.Jurusan, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}

func (r *ReviewRepository) FetchReviewByJurusan(Jurusan string) ([]*Review, error) {
	var review []*Review
	query := `
		SELECT * FROM review WHERE jurusan = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 10 OFFSET 0
	`
	rows, err := r.db.Query(query, Jurusan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.KampusID, &reviewTemp.UserID, &reviewTemp.Jurusan, &reviewTemp.Isian)
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
		SELECT * FROM review WHERE isian = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 10 OFFSET 0
	`
	rows, err := r.db.Query(query, Isian)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var reviewTemp Review
		err := rows.Scan(&reviewTemp.ID, &reviewTemp.KampusID, &reviewTemp.UserID, &reviewTemp.Jurusan, &reviewTemp.Isian)
		if err != nil {
			return nil, err
		}
		review = append(review, &reviewTemp)
	}
	return review, nil
}
