package entity

import "time"

type (
	User struct {
		ID           uint         `json:"id" gorm:"primaryKey"`
		NamaLengkap  string       `json:"nama_lengkap" gorm:"type:varchar(100);not null"`
		Email        string       `json:"email" binding:"email" gorm:"type:varchar(100);unique;not null"`
		Usia         int          `json:"usia" gorm:"not null"`
		JenisKelamin string       `json:"jenis_kelamin" gorm:"type:enum('Laki-laki', 'Perempuan');not null"`
		Alamat       string       `json:"alamat" gorm:"type:varchar(200);not null"`
		Password     string       `json:"password" gorm:"type:varchar(100);not null"`
		CreatedAt    time.Time    `json:"created_at"`
		UpdatedAt    time.Time    `json:"updated_at"`
		HasilTests   []*HasilTest `json:"hasil_tests"`
	}

	UserRegisterInputParam struct {
		NamaLengkap   string `json:"nama_lengkap" binding:"required"`
		Email         string `json:"email" binding:"email"`
		Usia          int    `json:"usia" binding:"required"`
		JenisKelamin  string `json:"jenis_kelamin" binding:"required"`
		Alamat        string `json:"alamat" binding:"required"`
		Password      string `json:"password" binding:"required"`
		VerifPassword string `json:"verif_password" binding:"required"`
	}

	UserLoginInputParam struct {
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required"`
	}

	UserUpdateInputParam struct {
		NamaLengkap  string `json:"nama_lengkap" binding:"required"`
		Email        string `json:"email" binding:"email"`
		Usia         int    `json:"usia" binding:"required"`
		JenisKelamin string `json:"jenis_kelamin" binding:"required"`
		Alamat       string `json:"alamat" binding:"required"`
	}

	UserLoginResponse struct {
		User  UserResponse `json:"user"`
		Token string       `json:"token"`
	}

	UserResponse struct {
		ID           uint                     `json:"id"`
		NamaLengkap  string                   `json:"nama_lengkap"`
		Email        string                   `json:"email"`
		Usia         int                      `json:"usia"`
		JenisKelamin string                   `json:"jenis_kelamin"`
		Alamat       string                   `json:"alamat"`
		Password     string                   `json:"password"`
		CreatedAt    time.Time                `json:"created_at"`
		UpdatedAt    time.Time                `json:"updated_at"`
		HasilTests   []*UserHasilTestResponse `json:"hasil_tests"`
	}

	UserHasilTestResponse struct {
		ID        uint   `json:"id"`
		Poin      int    `json:"poin"`
		CreatedAt string `json:"created_at"`
	}
)
