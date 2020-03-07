package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/store/sqlc"
	"github.com/jmoiron/sqlx"

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

	ErrValidation    = errors.New("Validation error")
	ErrNilParameters = errors.New("One or more params are nil")
)

// Server - basic route func type
type Server struct {
	Version string
	Queries *sqlc.Queries
	DB      *sqlx.DB
}

// GetGoods returns all products by shoppingId
func (s *Server) GetGoods(ctx echo.Context, shoppingID api.ShoppingID) error {
	response200 := func(items *[]api.ShoppingItem) error {
		var response api.Goods200
		response.Version = &s.Version
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
	goods, err := s.Queries.GetGoodsByShoppingID(context.Background(), int32ToNullInt32(int32(shoppingID)))
	if err != nil {
		return response500(err)
	}
	if len(goods) == 0 {
		return response404()
	}
	items := sqlcToShoppingItems(goods)
	return response200(&items)
}

func sqlcToShoppingItems(goods []sqlc.ShopList) (shoppingItems []api.ShoppingItem) {
	for _, i := range goods {
		var item api.ShoppingItem
		id := int(i.ID)
		category := int(i.CategoryID.Int32)
		complete := true
		if i.Complete.Int32 == 0 {
			complete = false
		}
		listID := int(i.ListID.Int32)
		productName := i.ProductName.String
		quantity := int(i.Quantity.Int32)

		item.Id = &id
		item.CategoryID = category
		item.Complete = complete
		item.ListID = listID
		item.ProductName = productName
		item.Quantity = quantity
		shoppingItems = append(shoppingItems, item)
	}
	return
}

// LastShopping returns LastShopping information
func (s *Server) LastShopping(ctx echo.Context) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.LastShopping200
		var data []api.ShoppingWithId
		data = append(data, shopping)
		response.Version = &s.Version
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
	lastShopping, err := s.Queries.GetLastShopping(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response404()
		}
		return response500(err)
	}
	data, err := s.sqlcToShopping(lastShopping)
	if err != nil {
		return response500(err)
	}
	return response200(data)
}

func (s *Server) sqlcToShopping(shopping sqlc.Shopping) (api.ShoppingWithId, error) {
	var sqlcShopping api.ShoppingWithId
	id := int(shopping.ID)
	sqlcShopping.Id = &id
	sqlcShopping.Date = shopping.Date.String
	ownerID := int(shopping.OwnerID.Int32)
	sqlcShopping.OwnerID = ownerID
	sqlcShopping.Time = shopping.Time.String

	shop, err := s.Queries.GetShopByID(context.Background(), shopping.ShopID.Int32)
	if err != nil {
		return api.ShoppingWithId{}, err
	}
	sqlcShopping.Name = shop.Name.String

	return sqlcShopping, nil
}

// AddShopping inserts new shopping
func (s *Server) AddShopping(ctx echo.Context) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.Shopping200
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = &shopping
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error, validation *api.ShoppingValidation) error {
		var response api.Shopping400
		response.Version = &s.Version
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
	var shParams api.ShoppingParams
	if err := ctx.Bind(&shParams); err != nil {
		return response400(err, nil)
	}
	_, err := time.Parse(dateLayout, shParams.Date)
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
	shopID, err := s.getShopID(shParams.Name)
	if err != nil {
		return response500(err)
	}
	shopping := shoppingToSqlc(shParams, shopID)
	insRes, err := s.Queries.AddShopping(context.Background(), shopping)
	if err != nil {
		return response500(err)
	}
	responseBody := shoppingToShoppingWithID(shParams, insRes)
	return response200(responseBody)
}

func shoppingToShoppingWithID(shopping api.ShoppingParams, shID int64) api.ShoppingWithId {
	shoppingID := int(shID)
	return api.ShoppingWithId{
		ShoppingParams: shopping,
		Id:             &shoppingID,
	}
}

func (s *Server) getShopID(name string) (int32, error) {
	shopName := sql.NullString{
		String: name,
		Valid:  true,
	}
	findRes, err := s.Queries.GetShopByName(context.Background(), shopName)
	if err != nil {
		insRes, err := s.Queries.AddShop(context.Background(), shopName)
		if err != nil {
			return 0, err
		}
		return int32(insRes), nil
	}
	return findRes.ID, nil
}

func shoppingToSqlc(shopping api.ShoppingParams, shopID int32) (params sqlc.AddShoppingParams) {
	params.Date = stringToNullString(shopping.Date)
	params.ShopID = int32ToNullInt32(shopID)
	params.Time = stringToNullString(shopping.Time)
	params.OwnerID = int32ToNullInt32(int32(shopping.OwnerID))
	return
}

// AddItem inserts new product to shopping cart
func (s *Server) AddItem(ctx echo.Context) error {
	response200 := func(item api.ShoppingItemParamsWithId) error {
		var response api.Item200
		var data []api.ShoppingItemParamsWithId
		data = append(data, item)
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = &data
		return ctx.JSON(http.StatusOK, response)
	}

	response400 := func(err error, validation *[]interface{}) error {
		var response api.Item400
		response.Version = &s.Version
		response.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}

	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	var itemParams api.ShoppingItemParams
	if err := ctx.Bind(&itemParams); err != nil {
		return response400(ErrValidation, nil)
	}

	_, err := s.Queries.GetShoppingByID(context.Background(), int32(itemParams.ListID))
	if err != nil {
		return response404()
	}

	sqlcItem := itemToSqlc(itemParams)
	insID, err := s.Queries.AddProductItem(context.Background(), sqlcItem)

	if err != nil {
		return response500(err)
	}

	responseBody := itemToItemWithID(itemParams, insID)
	return response200(responseBody)
}

func itemToSqlc(itemParams api.ShoppingItemParams) (params sqlc.AddProductItemParams) {
	params.ProductName = stringToNullString(itemParams.ProductName)
	params.Quantity = int32ToNullInt32(int32(itemParams.Quantity))
	params.ListID = int32ToNullInt32(int32(itemParams.ListID))
	params.Complete = sql.NullBool{
		Bool:  false,
		Valid: true,
	}
	params.CategoryID = int32ToNullInt32(int32(itemParams.CategoryID))

	return
}

func itemToItemWithID(item api.ShoppingItemParams, id int64) api.ShoppingItemParamsWithId {
	intID := int(id)
	return api.ShoppingItemParamsWithId{
		ShoppingItemParams: item,
		Id:                 &intID,
	}

}

// GetComingShoppings returns coming shoppings
func (s *Server) GetComingShoppings(ctx echo.Context, date api.Date) error {
	response200 := func(shoppings []api.ShoppingWithId) error {
		var response api.ComingShoppings200
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = &shoppings
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		var response api.ComingShoppings400
		response.Version = &s.Version
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
	_, err := time.Parse(dateLayout, dateParam)
	if err != nil {
		return response400(err)
	}
	commingShoppings, err := s.Queries.GetComingShoppings(context.Background(), stringToNullString(dateParam))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response404()
		}
		return response500(err)
	}
	var result []api.ShoppingWithId
	for _, v := range commingShoppings {
		data, err := s.sqlcToShopping(v)
		if err != nil {
			return response500(err)
		}
		result = append(result, data)
	}
	return response200(result)
}

// GetShoppingDays returns days with shopping by month and year
func (s *Server) GetShoppingDays(ctx echo.Context, year api.Year, month api.Month) error {
	response200 := func(days []int) error {
		var response api.ShoppingDays200
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = &days
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(validation *api.ShoppingDaysValidation) error {
		var response api.ShoppingDays400
		response.Version = &s.Version
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

	strMonth := strconv.Itoa(int(month))
	if month < 10 {
		strMonth = "0" + strMonth
	}
	queryParam := fmt.Sprintf("%v-%s", year, strMonth)
	queryParam = queryParam + "%"

	days, err := s.Queries.GetShoppingDays(ctx.Request().Context(), sql.NullString{
		String: queryParam,
		Valid:  true,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response404()
		}
		return response500(err)
	}
	var result []int
	for _, v := range days {
		fDate, err := time.Parse(dateLayout, v.String)
		if err != nil {
			return response500(err)
		}
		result = append(result, fDate.Day())
	}
	return response200(result)
}

func (s *Server) sqlcToShoppingWithId(sh []sqlc.Shopping) (*[]api.ShoppingWithId, error) {
	result := []api.ShoppingWithId{}
	shopIDs := []interface{}{}
	paramLabels := []string{}
	for i, v := range sh {
		shopIDs = append(shopIDs, v.ShopID.Int32)
		paramLabels = append(paramLabels, "$"+strconv.Itoa(i))
	}
	params := strings.Join(paramLabels, ",")
	query := strings.Replace(sqlc.GetShopNamesQuery, "$", params, 1)
	rows, err := s.DB.Query(query, shopIDs...)

	if err != nil {
		return nil, err
	}
	shopNames := []string{}
	for rows.Next() {
		name := ""
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		shopNames = append(shopNames, name)
	}
	for i, v := range sh {
		var shopping api.ShoppingWithId
		shopping.Id = intPtr(int(v.ID))
		shopping.Date = v.Date.String
		shopping.Name = shopNames[i]
		shopping.OwnerID = int(v.OwnerID.Int32)
		shopping.Time = v.Time.String
		result = append(result, shopping)
	}
	return &result, nil
}

//GetShoppingsByDay returns shoppings by day
func (s *Server) GetShoppingsByDay(ctx echo.Context, year api.Year, month api.Month, day api.Day) error {
	response200 := func(data *[]api.ShoppingWithId) error {
		var response api.Shoppings200
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = data
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(validation *api.ShoppingsByDayValidation) error {
		var response api.Shoppings400
		response.Version = &s.Version
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

	result, err := s.Queries.GetShoppingsByDay(ctx.Request().Context(), sql.NullString{
		String: queryParam,
		Valid:  true,
	})

	shoppings, err := s.sqlcToShoppingWithId(result)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response404()
		}
		return response500(err)
	}

	return response200(shoppings)
}

func (s *Server) GetShopping(ctx echo.Context, shoppingID api.ShoppingID) error {
	response200 := func(shopping api.ShoppingWithId) error {
		var response api.Shopping200
		response.Version = &s.Version
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

	shopping, err := s.Queries.GetShoppingByID(context.Background(), int32(shoppingID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response404()
		}
		return response500(err)
	}
	data, err := s.sqlcToShopping(shopping)
	if err != nil {
		return response500(err)
	}
	return response200(data)
}

// Удаление товаров
// (DELETE /deleteItems)
func (s *Server) DeleteItems(ctx echo.Context) error {
	response200 := func() error {
		response := api.Base200{}
		response.Version = &s.Version
		response.Message = SuccessMessage
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	var deleteNumbers api.DeleteItemsRequest
	if err := ctx.Bind(&deleteNumbers); err != nil {
		return response400(err)
	}

	query, args, err := sqlx.In("DELETE FROM shop_list WHERE id IN (?);", deleteNumbers.Ids)
	if err != nil {
		return response500(err)
	}
	_, err = s.DB.Exec(query, args...)
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
		response.Version = &s.Version
		response.Message = SuccessMessage
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {
		return s.error(ctx, http.StatusBadRequest, err, nil)
	}
	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}

	var deleteNumbers api.DeleteItemsRequest
	if err := ctx.Bind(&deleteNumbers); err != nil {
		return response400(err)
	}

	query, args, err := sqlx.In("DELETE FROM shopping WHERE id IN (?);", deleteNumbers.Ids)
	if err != nil {
		return response500(err)
	}
	_, err = s.DB.Exec(query, args...)
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
		response.Version = &s.Version
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

	var users []sqlc.User
	var err error

	if params.ComunityId == nil && params.TelegramUserId == nil {
		return response400(ErrNilParameters)
	}

	switch params.ComunityId {
	case nil:
		user, err := s.Queries.GetUserByTelegramID(contx, int32(*params.TelegramUserId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return response404()
			}
			return response500(err)
		}
		users = append(users, user)
	default:
		users, err = s.Queries.GetUsersByComunityID(contx, string(*params.ComunityId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return response404()
			}
			return response500(err)
		}
	}
	apiUsers := sqlcToAPIUsers(users)
	return response200(apiUsers)
}

// Добавление юзера
// (PATCH /users)
func (s *Server) UpdateUser(ctx echo.Context, params api.UpdateUserParams) error {
	return nil
}

// Добавление юзера
// (POST /users)
func (s *Server) CreateUser(ctx echo.Context) error {
	return nil
}
