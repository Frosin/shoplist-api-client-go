package store

import (
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/ent"

	"github.com/labstack/echo/v4"
)

func (s *Server) error(ctx echo.Context, httpCode int, err error, validation *[]interface{}) error {
	if err != nil {
		log.Errorf("code=%v, msg=%v", httpCode, err.Error())
	}

	switch httpCode {
	case 400: // BadRequest
		var response api.Error400
		response.Version = &s.version

		if validation != nil {
			response.Errors = *validation
		}
		response.Message = err.Error()
		return ctx.JSON(httpCode, response)
	case 401: // UnAuthorized
		var response api.Error401
		response.Version = &s.version
		response.Message = err.Error()
		return ctx.JSON(httpCode, response)
	case 404: // NotFound
		return ctx.JSON(httpCode, api.Error404{
			Error: api.Error{
				Base: api.Base{
					Version: &s.version,
				},
			},
			Message: NotFoundMessage,
		})
	case 405: // MethodNotAllowed
		return ctx.JSON(httpCode, api.Error405{
			Error: api.Error{
				Base: api.Base{
					Version: &s.version,
				},
			},
			Message: &MethodNotAllowedMessage,
		})
	}

	return ctx.JSON(httpCode, api.Error500{
		Error: api.Error{
			Base: api.Base{
				Version: &s.version,
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

func int32Ptr(i int32) *int32 {
	return &i
}

func (s *Server) FillFixtures() {
	ctx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	shop, err := s.ent.Shop.Create().SetName("Ашан").Save(ctx)
	if err != nil {
		log.Fatal("create shop error " + err.Error())
	}
	shg, err := s.ent.Shopping.Create().SetShop(shop).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.ent.Item.Create().SetShopping(shg).SetProductName("Хлеб").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.ent.Item.Create().SetShopping(shg).SetProductName("Батон").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.ent.Item.Create().SetShopping(shg).SetProductName("Молоко").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func entToShoppingItems(goods []*ent.Item) (shoppingItems []api.ShoppingItem) {
	for _, i := range goods {
		var item api.ShoppingItem
		item.Id = &i.ID
		item.CategoryID = i.CategoryID
		item.Complete = i.Complete
		item.ListID = i.Edges.Shopping.ID
		item.ProductName = i.ProductName
		item.Quantity = i.Quantity
		shoppingItems = append(shoppingItems, item)
	}
	return
}

func entToShopping(shopping *ent.Shopping) (apiSh api.ShoppingWithId) {
	apiSh.Id = &shopping.ID
	apiSh.Date = shopping.Date.Format(dateLayout)
	apiSh.Name = shopping.Edges.Shop.Name
	apiSh.OwnerID = shopping.Edges.User.ID
	apiSh.Time = shopping.Date.Format(timeLayout)
	return
}

func entToShoppings(shoppings []*ent.Shopping) (apiShs []api.ShoppingWithId) {
	for _, v := range shoppings {
		apiShs = append(apiShs, entToShopping(v))
	}
	return
}

func entToUser(u *ent.User) (apiUser api.UserWithID) {
	apiUser.Id = &u.ID
	chatID := int(u.ChatID)
	apiUser.ChatId = &chatID
	apiUser.ComunityId = &u.ComunityID
	TelegramID := int(u.TelegramID)
	apiUser.TelegramId = &TelegramID
	apiUser.TelegramUsername = &u.TelegramUsername
	apiUser.Token = &u.Token
	return
}

func entToUsers(users []*ent.User) (apiUsers []api.UserWithID) {
	for _, v := range users {
		apiUsers = append(apiUsers, entToUser(v))
	}
	return
}
