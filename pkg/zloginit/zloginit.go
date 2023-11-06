package zloginit

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		if l == zerolog.PanicLevel {
			return zerolog.FatalLevel.String()
		}
		return l.String()
	}
	log.Logger = log.With().Caller().Logger()
}
