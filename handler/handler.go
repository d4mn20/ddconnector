package handler

import (
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
