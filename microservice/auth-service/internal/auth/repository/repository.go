package repository

import (
	"database/sql"

	"github.com/username/auth-service/internal/auth/models"
)

type AuthRepository interface {
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(username string) error
	GetUserByUsername(username string) (*models.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) InsertUser(user *models.User) error {
	// implementasi query untuk menyimpan data pengguna baru ke dalam database
}

func (a *authRepository) UpdateUser(user *models.User) error {
	// implementasi query untuk memperbarui data pengguna di dalam database
}

func (a *authRepository) DeleteUser(username string) error {
	// implementasi query untuk menghapus data pengguna dari database
}

func (a *authRepository) GetUserByUsername(username string) (*models.User, error) {
	// implementasi query untuk mengambil data pengguna berdasarkan username dari database
}
