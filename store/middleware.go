package store

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/Frosin/shoplist-api-client-go/ent/user"
	"github.com/labstack/echo/v4"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

func (s *Server) TokenHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		response401 := func(err error) error {
			return s.error(c, http.StatusUnauthorized, err, nil)
		}
		response500 := func(err error) error {
			return s.error(c, http.StatusInternalServerError, err, nil)
		}
		token := c.QueryParam("token")
		if token == "" {
			return response401(ErrTokenNotFound)
		}
		contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
		defer cancel()
		usr, err := s.ent.User.
			Query().
			Where(user.TokenEQ(token)).
			Only(contx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return response401(ErrTokenNotFound)
			}
			return response500(err)
		}
		c.Set("ownerID", usr.ID)

		comunityUsers, err := s.ent.User.
			Query().
			Where(user.ComunityIDEQ(usr.ComunityID)).
			All(contx)
		if err != nil {
			return response500(err)
		}
		comUserIDs := []int{}
		for _, v := range comunityUsers {
			comUserIDs = append(comUserIDs, v.ID)
		}
		c.Set("comunityUserIDs", comUserIDs)

		return next(c)
	}
}
