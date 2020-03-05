package store

import (
	"database/sql"

	"github.com/labstack/gommon/log"

	"github.com/Frosin/shoplist-api-client-go/api"

	"github.com/labstack/echo/v4"
)

func (s *Server) error(ctx echo.Context, httpCode int, err error, validation *[]interface{}) error {
	if err != nil {
		log.Info(err.Error())
	}

	switch httpCode {
	case 400: // BadRequest
		var response api.Error400
		response.Version = &s.Version

		if validation != nil {
			response.Errors = *validation
		}
		response.Message = err.Error()
		return ctx.JSON(httpCode, response)
	case 404: // NotFound
		return ctx.JSON(httpCode, api.Error404{
			Error: api.Error{
				Base: api.Base{
					Version: &s.Version,
				},
			},
			Message: NotFoundMessage,
		})
	case 405: // MethodNotAllowed
		return ctx.JSON(httpCode, api.Error405{
			Error: api.Error{
				Base: api.Base{
					Version: &s.Version,
				},
			},
			Message: &MethodNotAllowedMessage,
		})
	}

	return ctx.JSON(httpCode, api.Error500{
		Error: api.Error{
			Base: api.Base{
				Version: &s.Version,
			},
		},
		Errors:  err.Error(),
		Message: InternalServerErrorMessage,
	})
}

func int32ToNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

func stringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
