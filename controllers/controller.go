package controllers

import (
	"github.com/apex/log"
	"gorm.io/gorm"
)

type Controller struct {
	Logger *log.Entry
	DB     *gorm.DB
}
