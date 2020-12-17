package g

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

var (
	DB         *gorm.DB
	VP         *viper.Viper
	Logger     *zap.Logger
	HttpClient *http.Client
	Redis      *redis.Client
)
