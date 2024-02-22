package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go-backend/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GlobDB     *gorm.DB
	GlobConfig config.ServerConfig
	GlobLogger *zap.Logger
	Redis      *redis.Client
	Trans      ut.Translator
)
