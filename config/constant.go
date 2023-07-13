package config

import (
	"go.uber.org/zap"
	"tes-face/entity"
)

var (
	GLOBAL   *entity.OssInfo
	FILEPATH *entity.FilePath
	TES_LOG  *zap.Logger
)
