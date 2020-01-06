package store

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/store/sqlc"

	"github.com/labstack/echo/v4"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04:05"
)

var (
	SuccessMessage             = "success"
	InternalServerErrorMessage = "Internal server error"
	NotFoundMessage            = "Entity not found"
	AccessDeniedMessage        = "Access denied"
	UnknownPathMessage         = "Unknown path"
	MethodNotAllowedMessage    = "Method not allowed"

	ErrValidation = errors.New("Validation error")
)

// Server - basic route func type
type Server struct {
	Version string
	Queries *sqlc.Queries
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
	response400 := func(err error, validation *[]interface{}) error {
		return s.error(ctx, http.StatusBadRequest, err, validation)
	}
	response404 := func() error {
		return s.error(ctx, http.StatusNotFound, nil, nil)
	}

	response500 := func(err error) error {
		return s.error(ctx, http.StatusInternalServerError, err, nil)
	}
	sID, err := strconv.Atoi(string(shoppingID))
	if err != nil {
		validationError := map[string]map[string]string{
			"validation": map[string]string{
				"shoppingID": "format",
			},
		}
		validation := []interface{}{validationError}
		return response400(err, &validation)
	}

	goods, err := s.Queries.GetGoodsByShoppingID(context.Background(), int32ToNullInt32(int32(sID)))
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
		var response api.LastShopping200
		var data []api.ShoppingWithId
		data = append(data, shopping)
		response.Version = &s.Version
		response.Message = SuccessMessage
		response.Data = &data
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
		return response400(ErrValidation, nil)
	}
	_, err := time.Parse(dateLayout, shParams.Date)
	if err != nil {
		validation := api.ShoppingValidation{
			Date: strPtr("format"),
		}
		return response400(ErrValidation, &validation)
	}
	_, err = time.Parse(timeLayout, shParams.Time)
	if err != nil {
		validation := api.ShoppingValidation{
			Time: strPtr("format"),
		}
		return response400(ErrValidation, &validation)
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

func (s *Server) error(ctx echo.Context, httpCode int, err error, validation *[]interface{}) error {
	if err != nil {
		log.Println(err.Error())
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
