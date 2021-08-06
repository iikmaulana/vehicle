package config

import (
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/vehicle/base/libs/token"
	"os"
)

func (c *Config) InitToken() serror.SError {
	c.Token = &token.Initialize{
		Alg:                    "HS256",
		SecretKey:              os.Getenv("JWT_SECRET_KEY"),
		RefreshTokenExpiration: helper.StringToInt(os.Getenv("JWT_REFRESH_DURATION"), 17520),
		AccesstokenExpiration:  helper.StringToInt(os.Getenv("JWT_ACCESS_DURATION"), 8760),
	}

	return nil
}
