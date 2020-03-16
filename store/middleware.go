package store

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/Frosin/shoplist-api-client-go/ent/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
		dontNeedToken := func(request *http.Request) bool {
			return strings.Contains(request.RequestURI, "users") &&
				(request.Method == "GET" ||
					request.Method == "POST")
		}

		//if it create or get users api we dont need token
		if dontNeedToken(c.Request()) {
			log.Infof("request: %s, %s", c.Request().RequestURI, "pass token check")
			return next(c)
		}
		log.Infof("request: %s", c.Request().RequestURI)

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
		log.Infof("userID=%v, userTelegramID=%v, userToken=%v, userComunityID=%v", usr.ID, usr.TelegramID, usr.Token, usr.ComunityID)
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
