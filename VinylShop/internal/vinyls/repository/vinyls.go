package repository

import (
	"VinylShop/internal/vinyls/domain"
	"VinylShop/pkg/web/errors"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type VinylRepositoryImpl struct {
	DB *sql.DB
}

func NewVinylRepository(db *sql.DB) *VinylRepositoryImpl {
	return &VinylRepositoryImpl{
		DB: db,
	}
}
func (d *VinylRepositoryImpl) Insert(vinylInsert *domain.Vinyl) (int, error) {
	stmt := "INSERT INTO vinyls(title, artist, releasedate, price, rating) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var insertedId int
	row := d.DB.QueryRow(stmt, vinylInsert.Title, vinylInsert.Artist, vinylInsert.ReleaseDate, vinylInsert.Price, vinylInsert.Rating)
	if err := row.Scan(&insertedId); err != nil {
		return 0, err
	}
	return insertedId, nil
}
func (d *VinylRepositoryImpl) GetAll(filters *domain.Filters) ([]domain.Vinyl, error) {
	var vinyls []domain.Vinyl
	stmt := fmt.Sprintf("SELECT * FROM vinyls WHERE %s LIKE %s ORDER BY %s %s LIMIT $1 OFFSET $2", filters.FilterType, filters.Filter, filters.Sort, filters.Order)
	rows, err := d.DB.Query(stmt, filters.Limit, filters.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var vinyl domain.Vinyl
		err := rows.Scan(&vinyl.ID, &vinyl.Title, &vinyl.Artist, &vinyl.ReleaseDate, &vinyl.Price, &vinyl.Rating, &vinyl.CreatedAt, &vinyl.UpdatedAt)
		if err != nil {
			return nil, err
		}
		vinyls = append(vinyls, vinyl)
	}
	return vinyls, nil
}
func (d *VinylRepositoryImpl) Get(id int) (domain.Vinyl, error) {
	var vinyl domain.Vinyl
	stmt := `SELECT * FROM vinyls WHERE id=$1`
	fmt.Println("Executing query:", stmt, "with id:", id)
	err := d.DB.QueryRow(stmt, id).Scan(&vinyl.ID, &vinyl.Title, &vinyl.Artist, &vinyl.ReleaseDate, &vinyl.Price, &vinyl.Rating, &vinyl.CreatedAt, &vinyl.UpdatedAt)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return domain.Vinyl{}, err
	}
	return vinyl, nil
}
func (d *VinylRepositoryImpl) Length() int {
	var length int
	stmt := `SELECT COUNT(*) FROM vinyls`
	err := d.DB.QueryRow(stmt).Scan(&length)
	if err != nil {
		return 0
	}
	return length
}

//	func (d *VinylRepositoryImpl) Add(vinyl domain.Vinyl) (int64, error) {
//		result, err := d.DB.Exec("INSERT INTO vinyls(title, artist, price, releasedate) VALUES (?,?,?,?)", vinyl.Title, vinyl.Artist, vinyl.Price, vinyl.ReleaseDate)
//		if err != nil {
//			return 0, fmt.Errorf("Problem with vinyl at: %v", err)
//		}
//		id, err := result.LastInsertId()
//		if err != nil {
//			return 0, fmt.Errorf("Problem with vinyl at: %v", err)
//		}
//		return id, nil
//	}
/*func (d *VinylRepositoryImpl) Update(id int, vinylUpdate *domain.Vinyl) (*domain.Vinyl, error) {
	stmt := `UPDATE vinyls SET title=$1, artist=$2, releasedate=$3, price=$4, rating=$5 WHERE id=$6 RETURNING *;`
	row := d.DB.QueryRow(stmt, vinylUpdate.Title, vinylUpdate.Artist, vinylUpdate.ReleaseDate, vinylUpdate.Price, vinylUpdate.Rating, id)
	err := row.Scan(&vinylUpdate.ID, &vinylUpdate.Title, &vinylUpdate.Artist, &vinylUpdate.ReleaseDate, &vinylUpdate.Price, &vinylUpdate.Rating, &vinylUpdate.CreatedAt, &vinylUpdate.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return vinylUpdate, nil
}*/
func (d *VinylRepositoryImpl) Update(id int, vinylUpdate *domain.Vinyl) (*domain.Vinyl, error) {
	stmt := `UPDATE vinyls SET title=$1, artist=$2, releasedate=$3, price=$4, rating=$5, updatedat=now() WHERE id=$6 RETURNING id, title, artist, releasedate, price, rating, createdat, updatedat;`
	row := d.DB.QueryRow(stmt, vinylUpdate.Title, vinylUpdate.Artist, vinylUpdate.ReleaseDate, vinylUpdate.Price, vinylUpdate.Rating, id)
	err := row.Scan(&vinylUpdate.ID, &vinylUpdate.Title, &vinylUpdate.Artist, &vinylUpdate.ReleaseDate, &vinylUpdate.Price, &vinylUpdate.Rating, &vinylUpdate.CreatedAt, &vinylUpdate.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return vinylUpdate, nil
}
func (d *VinylRepositoryImpl) Delete(id int) error {
	stmt := "DELETE FROM vinyls WHERE id=$1"
	_, err := d.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
func (d *VinylRepositoryImpl) Exists(id int) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM vinyls where id=$1)`
	if err := d.DB.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}
func (d *VinylRepositoryImpl) Authenticate(email string, password []byte) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users WHERE email=$1"
	if err := d.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword); err != nil {
		return -1, errors.ErrEmailNotFound
	}
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return -1, errors.ErrIncorrectPassword
	}
	return id, nil
}
func (d *VinylRepositoryImpl) ExistsByEmail(email string) bool {
	var exists bool
	stmt := "SELECT EXISTS(SELECT TRUE FROM users WHERE email=$1)"
	if err := d.DB.QueryRow(stmt, email).Scan(&exists); err != nil {
		return false
	}
	return exists
}
