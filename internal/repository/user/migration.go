package user

import "github.com/Haoyunforever/Micro-Scanner/internal/model"

func migration() {
	_db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.User{})
}
