package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
}

func (m *Middleware) RecoveringMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if rerr := recover(); rerr != nil {
					err, ok := rerr.(error)
					if !ok {
						err = fmt.Errorf("%v", err)
					}
					ctx.Logger().Printf("[PANIC RECOVER] %v\n", err.Error())
				}
			}()
			return next(ctx)
		}
	}
}
