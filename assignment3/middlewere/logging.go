package middlewere

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func ZerologMiddleware(nextHanler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		startTime := time.Now()

		err := nextHanler(context)

		log.Info().Str("method", context.Request().Method).Str("path", context.Request().URL.Path).Dur("latency", time.Since(startTime)).Msg("Request")

		return err

	}
}
