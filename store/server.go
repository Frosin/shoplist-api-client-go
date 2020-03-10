package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/ent"
	"github.com/Frosin/shoplist-api-client-go/ent/item"
	"github.com/Frosin/shoplist-api-client-go/ent/predicate"
	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/Frosin/shoplist-api-client-go/ent/user"
	entSql "github.com/facebookincubator/ent/dialect/sql"

	"github.com/labstack/echo/v4"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04:05"
)

var (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 10 * time.Second

	SuccessMessage             = "success"
	InternalServerErrorMessage = "Internal server error"
	NotFoundMessage            = "Entity not found"
	AccessDeniedMessage        = "Access denied"
	UnknownPathMessage         = "Unknown path"
	MethodNotAllowedMessage    = "Method not allowed"

	ErrValidation            = errors.New("Validation error")
	ErrNilParameters         = errors.New("One or more params are nil")
	ErrTypeAssertion         = errors.New("type assertion error")
	ErrUpdateUserConstFields = errors.New("can't update constant user fields")
)

// Server - basic route func type
type Server struct {
	version string
	//Queries *sqlc.Queries
	ent *ent.Client
	db  *sql.DB
}

func NewServer(v string, e *ent.Client, db *sql.DB) *Server {
	return &Server{
		version: v,
		ent:     e,
		db:      db,
	}
}

// GetGoods returns all products by shoppingId
func (s *Server) GetGoods(ctx echo.Context, shoppingID api.ShoppingID) error {
	response200 := func(items *[]api.ShoppingItem) error {
		var response api.Goods200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = items
		return ctx.JSON(http.StatusOK, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}

	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	goods, err := s.ent.Item.
		Query().
		WithShopping().
		Where(item.HasShoppingWith(shopping.IDEQ(int(shoppingID)))).
		All(contx)
	// goods, err := s.Queries.GetGoodsByShoppingID(context.Background(), int32ToNullInt32(int32(shoppingID)))
	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}
	// if len(goods) == 0 {
	// 	return response404()
	// }
	items := entToShoppingItems(goods)
	return response200(&items)
}

// // Добавление товара в покупку
// // (POST /addItem)
// func (s *Server) AddItem(ctx echo.Context) error {
// 	return nil
// }

// // Добавление покупки
// // (POST /addShopping)
// func (s *Server) AddShopping(ctx echo.Context) error {
// 	return nil
// }

// // Удаление товаров
// // (POST /deleteItems)
// func (s *Server) DeleteItems(ctx echo.Context) error {
// 	return nil
// }

// // Удаление покупок
// // (POST /deleteShoppings)
// func (s *Server) DeleteShoppings(ctx echo.Context) error {
// 	return nil
// }

// // Ближайшие 5 покупок
// // (GET /getComingShoppings/{date})
// func (s *Server) GetComingShoppings(ctx echo.Context, date api.Date) error {
// 	return nil
// }

// // Даные покупки
// // (GET /getShopping/{shoppingID})
// func (s *Server) GetShopping(ctx echo.Context, shoppingID api.ShoppingID) error {
// 	return nil
// }

// // Получение списка дней с покупками по месяцу и году
// // (GET /getShoppingDays/{year}/{month})
// func (s *Server) GetShoppingDays(ctx echo.Context, year api.Year, month api.Month) error {
// 	return nil
// }

// // Получение списка покупок по конекретному дню
// // (GET /getShoppingsByDay/{year}/{month}/{day})
// func (s *Server) GetShoppingsByDay(ctx echo.Context, year api.Year, month api.Month, day api.Day) error {
// 	return nil
// }

// // Последняя покупка
// // (GET /lastShopping)
// func (s *Server) LastShopping(ctx echo.Context) error {
// 	return nil
// }

// // Получение юзера по telegram user id
// // (GET /users)
// func (s *Server) GetUsers(ctx echo.Context, params api.GetUsersParams) error {
// 	return nil
// }

// // Добавление юзера
// // (PATCH /users)
// func (s *Server) UpdateUser(ctx echo.Context, params api.UpdateUserParams) error {
// 	return nil
// }

// // Добавление юзера
// // (POST /users)
// func (s *Server) CreateUser(ctx echo.Context) error {
// 	return nil
// }

// LastShopping returns LastShopping information
func (s *Server) LastShopping(ctx echo.Context) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.LastShopping200
		var data []api.ShoppingWithId
		data = append(data, shopping)
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &data
		return ctx.JSON(http.StatusOK, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	lastShopping, err := s.ent.Shopping.Query().
		WithShop().
		Order(ent.Desc("rowid")).
		Where(shopping.HasUserWith(
			user.IDIn(userIDs...),
		)).
		Limit(1).
		Only(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}
	data := entToShopping(lastShopping)
	if err != nil {
		return response500(err)
	}
	return response200(data)
}

// AddShopping inserts new shopping
func (s *Server) AddShopping(ctx echo.Context) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.Shopping200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &shopping
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error, validation *api.ShoppingValidation) error {
		var response api.Shopping400
		response.Version = &s.version
		response.Message = err.Error()

		if validation != nil {
			response.Errors = &api.ShoppingProperty{
				Validation: validation,
			}
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()
	userID, ok := ctx.Get("ownerID").(int)
	if !ok {
		return response500(ErrTypeAssertion)
	}

	var shParams api.ShoppingParams
	if err := ctx.Bind(&shParams); err != nil {
		return response400(err, nil)
	}
	date, err := time.Parse(dateLayout, shParams.Date)
	if err != nil {
		validation := api.ShoppingValidation{
			Date: strPtr("format"),
		}
		return response400(err, &validation)
	}
	_, err = time.Parse(timeLayout, shParams.Time)
	if err != nil {
		validation := api.ShoppingValidation{
			Time: strPtr("format"),
		}
		return response400(err, &validation)
	}

	shop, err := s.ent.Shop.
		Query().
		Where(shop.NameEQ(shParams.Name)).
		First(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			shop, err = s.ent.Shop.
				Create().
				SetName(shParams.Name).Save(contx)
			if err != nil {
				return response500(err)
			}
		}
		return response500(err)
	}

	var newShopping *ent.Shopping
	err = s.ent.WithTx(contx, func(tx *ent.Tx) error {
		newShopping, err = tx.Shopping.
			Create().
			SetShop(shop).
			SetDate(date).
			SetUserID(userID).
			Save(contx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return response500(err)
	}

	return response200(api.ShoppingWithId{
		ShoppingParams: shParams,
		Id:             &newShopping.ID,
	})
}

// AddItem inserts new product to shopping cart
func (s *Server) AddItem(ctx echo.Context) error {
	response200 := func(item api.ShoppingItemParamsWithId) error {
		var response api.Item200
		var data []api.ShoppingItemParamsWithId
		data = append(data, item)
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &data
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error, validation *[]interface{}) error {
		var response api.Item400
		response.Version = &s.version
		response.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var itemParams api.ShoppingItemParams
	if err := ctx.Bind(&itemParams); err != nil {
		return response400(ErrValidation, nil)
	}

	shopping, err := s.ent.Shopping.
		Query().
		Where(shopping.IDEQ(itemParams.ListID)).
		Only(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}

	var newItem *ent.Item
	err = s.ent.WithTx(contx, func(tx *ent.Tx) error {
		newItem, err = tx.Item.
			Create().
			SetProductName(itemParams.ProductName).
			SetShopping(shopping).
			Save(contx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return response500(err)
	}

	return response200(api.ShoppingItemParamsWithId{
		Id:                 &newItem.ID,
		ShoppingItemParams: itemParams,
	})
}

// GetComingShoppings returns coming shoppings
func (s *Server) GetComingShoppings(ctx echo.Context, date api.Date) error {
	response200 := func(shoppings []api.ShoppingWithId) error {
		var response api.ComingShoppings200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &shoppings
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		var response api.ComingShoppings400
		response.Version = &s.version
		response.Message = err.Error()
		response.Errors = &api.ComingShoppingsProperty{
			Validation: &api.ComingShoppingsValidation{
				Date: strPtr("format"),
			},
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	dateParam := string(date)
	dTime, err := time.Parse(dateLayout, dateParam)
	if err != nil {
		return response400(err)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	commingShoppings, err := s.ent.Shopping.
		Query().
		Where(
			shopping.DateGTE(dTime),
			shopping.HasUserWith(user.IDIn(userIDs...)),
		).
		Order(ent.Desc("rowid")).
		Limit(5).
		All(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}

	return response200(entToShoppings(commingShoppings))
}

// GetShoppingDays returns days with shopping by month and year
func (s *Server) GetShoppingDays(ctx echo.Context, year api.Year, month api.Month) error {
	response200 := func(days []int) error {
		var response api.ShoppingDays200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &days
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(validation *api.ShoppingDaysValidation) error {
		var response api.ShoppingDays400
		response.Version = &s.version
		response.Message = ErrValidation.Error()
		response.Errors = &api.ShoppingDaysErrors{
			Validation: validation,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	curYear := time.Now().Year()
	var validation api.ShoppingDaysValidation
	var valCount = 0
	if math.Abs(float64(curYear-int(year))) > 1 {
		validation.Year = strPtr("format")
		valCount++
	}
	if month < 1 || month > 12 {
		validation.Month = strPtr("format")
		valCount++
	}
	if valCount != 0 {
		return response400(&validation)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	strMonth := strconv.Itoa(int(month))
	if month < 10 {
		strMonth = "0" + strMonth
	}
	queryParam := fmt.Sprintf("%v-%s", year, strMonth)

	monthShoppings, err := s.ent.Shopping.
		Query().
		Where(
			shopping.HasUserWith(
				user.IDIn(userIDs...),
			),
			predicate.Shopping(func(s *entSql.Selector) {
				s.Where(entSql.Like(s.C(shopping.FieldDate), queryParam))
			})).
		All(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}

	var result []int
	for _, v := range monthShoppings {
		result = append(result, v.Date.Day())
	}

	return response200(result)
}

//GetShoppingsByDay returns shoppings by day
func (s *Server) GetShoppingsByDay(ctx echo.Context, year api.Year, month api.Month, day api.Day) error {
	response200 := func(data []api.ShoppingWithId) error {
		var response api.Shoppings200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &data
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(validation *api.ShoppingsByDayValidation) error {
		var response api.Shoppings400
		response.Version = &s.version
		response.Message = ErrValidation.Error()
		response.Errors = &api.ShoppingsByDayErrors{
			Validation: validation,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	curYear := time.Now().Year()
	var validation api.ShoppingsByDayValidation
	var valCount = 0
	if math.Abs(float64(curYear-int(year))) > 1 {
		validation.Year = strPtr("format")
		valCount++
	}
	if month < 1 || month > 12 {
		validation.Month = strPtr("format")
		valCount++
	}
	if day < 1 || day > 31 {
		validation.Day = strPtr("format")
		valCount++
	}
	if valCount != 0 {
		return response400(&validation)
	}

	strMonth := strconv.Itoa(int(month))
	if month < 10 {
		strMonth = "0" + strMonth
	}

	strDay := strconv.Itoa(int(day))
	if day < 10 {
		strDay = "0" + strDay
	}

	queryParam := fmt.Sprintf("%v-%s-%s", year, strMonth, strDay)

	shoppings, err := s.ent.Shopping.
		Query().
		Where(
			shopping.HasUserWith(
				user.IDIn(userIDs...),
			),
			predicate.Shopping(func(s *entSql.Selector) {
				s.Where(entSql.Like(s.C(shopping.FieldDate), queryParam))
			})).
		All(contx)

	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}

	return response200(entToShoppings(shoppings))
}

func (s *Server) GetShopping(ctx echo.Context, shoppingID api.ShoppingID) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.Shopping200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = &shopping
		return ctx.JSON(http.StatusOK, response)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	shopping, err := s.ent.Shopping.
		Query().
		Where(
			shopping.IDEQ(int(shoppingID)),
			shopping.HasUserWith(
				user.IDIn(userIDs...),
			)).
		Only(contx)
	if err != nil {
		if ent.IsNotFound(err) {
			return response404()
		}
		return response500(err)
	}

	return response200(entToShopping(shopping))
}

// Удаление товаров
// (DELETE /deleteItems)
func (s *Server) DeleteItems(ctx echo.Context) error {
	response200 := func() error {
		response := api.Base200{}
		response.Version = &s.version
		response.Message = SuccessMessage
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var deleteNumbers api.DeleteItemsRequest
	if err := ctx.Bind(&deleteNumbers); err != nil {
		return response400(err)
	}

	_, err := s.ent.Item.Delete().Where(item.IDIn(deleteNumbers.Ids...)).Exec(contx)
	if err != nil {
		return response500(err)
	}

	return response200()
}

// Удаление покупок
// (DELETE /deleteShoppings)
func (s *Server) DeleteShoppings(ctx echo.Context) error {
	response200 := func() error {
		response := api.Base200{}
		response.Version = &s.version
		response.Message = SuccessMessage
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	userIDs, ok := ctx.Get("comunityUserIDs").([]int)
	if !ok {
		return response500(ErrTypeAssertion)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var deleteNumbers api.DeleteItemsRequest
	if err := ctx.Bind(&deleteNumbers); err != nil {
		return response400(err)
	}

	_, err := s.ent.Shopping.
		Delete().
		Where(
			shopping.HasUserWith(
				user.IDIn(userIDs...),
			),
			shopping.IDIn(deleteNumbers.Ids...),
		).
		Exec(contx)
	if err != nil {
		return response500(err)
	}
	return response200()
}

// Получение юзера по telegram user id
// (GET /users)
func (s *Server) GetUsers(ctx echo.Context, params api.GetUsersParams) error {
	response200 := func(users *[]api.UserWithID) error {
		var response api.Users200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = users
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var users []*ent.User
	var err error

	if params.ComunityId == nil && params.TelegramUserId == nil {
		return response400(ErrNilParameters)
	}

	switch params.ComunityId {
	case nil:
		user, err := s.ent.User.
			Query().
			Where(user.TelegramIDEQ(int64(*params.TelegramUserId))).
			Only(contx)
		if err != nil {
			if ent.IsNotFound(err) {
				return response404()
			}
			return response500(err)
		}
		users = append(users, user)
	default:
		users, err = s.ent.User.
			Query().
			Where(user.ComunityIDEQ(string(*params.ComunityId))).
			All(contx)
		if err != nil {
			if ent.IsNotFound(err) {
				return response404()
			}
			return response500(err)
		}
	}
	apiUsers := entToUsers(users)
	return response200(&apiUsers)
}

// Добавление юзера
// (PATCH /users)
func (s *Server) UpdateUser(ctx echo.Context, params api.UpdateUserParams) error {
	response200 := func() error {
		var response api.Base200
		response.Version = &s.version
		response.Message = SuccessMessage
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var user api.UpdateUserJSONRequestBody
	var err error
	if err = ctx.Bind(&user); err != nil {
		return response400(err)
	}
	updNum := 0

	if user.ChatId != nil || user.TelegramId != nil || user.Token != nil {
		return response400(ErrUpdateUserConstFields)
	}

	if user.ComunityId != nil {
		_, err = s.ent.User.
			UpdateOneID(int(params.UserId)).
			SetComunityID(*user.ComunityId).
			Save(contx)
		if err != nil {
			if ent.IsNotFound(err) {
				return response404()
			}
			return response500(err)
		}
		updNum++
	}

	if user.TelegramUsername != nil {
		_, err = s.ent.User.
			UpdateOneID(int(params.UserId)).
			SetTelegramUsername(*user.TelegramUsername).
			Save(contx)
		if err != nil {
			if ent.IsNotFound(err) {
				return response404()
			}
			return response500(err)
		}
		updNum++
	}

	if updNum == 0 {
		return response400(ErrNilParameters)
	}

	return response200()
}

// Добавление юзера
// (POST /users)
func (s *Server) CreateUser(ctx echo.Context) error {
	response200 := func(users *[]api.UserWithID) error {
		var response api.Users200
		response.Version = &s.version
		response.Message = SuccessMessage
		response.Data = users
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	contx, cancel := context.WithTimeout(context.Background(), ReadTimeout)
	defer cancel()

	var u api.UpdateUserJSONRequestBody
	if err := ctx.Bind(&u); err != nil {
		return response400(err)
	}

	if u.ComunityId == nil ||
		u.TelegramId == nil ||
		u.ChatId == nil ||
		u.TelegramUsername == nil ||
		u.Token == nil {
		return response400(ErrNilParameters)
	}

	user, err := s.ent.User.
		Create().
		SetChatID(int64(*u.ChatId)).
		SetComunityID(*u.ComunityId).
		SetTelegramID(int64(*u.TelegramId)).
		SetTelegramUsername(*u.TelegramUsername).
		SetToken(*u.Token).Save(contx)

	if err != nil {
		return response500(err)
	}

	return response200(&[]api.UserWithID{
		entToUser(user),
	})
}
