package repository

import (
	"database/sql"
	"fmt"

	"github.com/username/admin-service/internal/admin/models"
)

type AdminRepository interface {
	InsertAdmin(admin *models.Admin) error
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(adminID string) error
	GetAdminByID(adminID string) (*models.Admin, error)
	GetAllAdmins() ([]*models.Admin, error)
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) InsertAdmin(admin *models.Admin) error {
	sql := fmt.Sprintf("INSERT INTO admins(id, name, email) VALUES ('%s', '%s', '%s')", admin.ID, admin.Name, admin.Email)
	_, err := r.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) UpdateAdmin(admin *models.Admin) error {
	sql := fmt.Sprintf("UPDATE admins SET name='%s', email='%s' WHERE id='%s'", admin.Name, admin.Email, admin.ID)
	_, err := r.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) DeleteAdmin(adminID string) error {
	sql := fmt.Sprintf("DELETE FROM admins WHERE id='%s'", adminID)
	_, err := r.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetAdminByID(adminID string) (*models.Admin, error) {
	sql := fmt.Sprintf("SELECT id, name, email FROM admins WHERE id='%s'", adminID)
	row := r.db.QueryRow(sql)
	admin := &models.Admin{}
	err := row.Scan(&admin.ID, &admin.Name, &admin.Email)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) GetAllAdmins() ([]*models.Admin, error) {
	sql := "SELECT id, name, email FROM admins"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	admins := []*models.Admin{}
	for rows.Next() {
		admin := &models.Admin{}
		err := rows.Scan(&admin.ID, &admin.Name, &admin.Email)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return admins, nil
}
