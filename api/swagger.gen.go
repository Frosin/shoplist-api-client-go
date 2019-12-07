// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// Base defines model for Base.
type Base struct {
	Version *string `json:"version,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Embedded struct due to allOf(#/components/schemas/Base)
	Base
	// Embedded fields due to inline allOf schema
	Data interface{} `json:"data"`
}

// Error400 defines model for Error_400.
type Error400 struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error
	// Embedded fields due to inline allOf schema
	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}

// Error404 defines model for Error_404.
type Error404 struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error
	// Embedded fields due to inline allOf schema
	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}

// Error405 defines model for Error_405.
type Error405 struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error
	// Embedded fields due to inline allOf schema
	Errors  *[]interface{} `json:"errors,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// Error500 defines model for Error_500.
type Error500 struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error
	// Embedded fields due to inline allOf schema
	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
}

// Shop defines model for Shop.
type Shop struct {
	ID   *int    `json:"ID,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Shopping defines model for Shopping.
type Shopping struct {
	// Embedded fields due to inline allOf schema
	// Embedded struct due to allOf(#/components/schemas/Shopping_params)
	ShoppingParams
	// Embedded struct due to allOf(#/components/schemas/Shop)
	Shop
}

// ShoppingItem defines model for Shopping_item.
type ShoppingItem struct {
	// Embedded fields due to inline allOf schema
	// Embedded struct due to allOf(#/components/schemas/Shopping_item_params_with_id)
	ShoppingItemParamsWithId
}

// ShoppingItemParams defines model for Shopping_item_params.
type ShoppingItemParams struct {
	CategoryID  *int    `json:"categoryID,omitempty"`
	Complete    *bool   `json:"complete,omitempty"`
	ListID      *int    `json:"listID,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	Quantity    *int    `json:"quantity,omitempty"`
}

// ShoppingItemParamsWithId defines model for Shopping_item_params_with_id.
type ShoppingItemParamsWithId struct {
	// Embedded fields due to inline allOf schema
	ID *int `json:"ID,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/Shopping_item_params)
	ShoppingItemParams
}

// ShoppingParams defines model for Shopping_params.
type ShoppingParams struct {
	Complete *bool   `json:"complete,omitempty"`
	Date     *string `json:"date,omitempty"`
	OwnerID  *int    `json:"ownerID,omitempty"`
	ShopID   *int    `json:"shopID,omitempty"`
	Sum      *int    `json:"sum,omitempty"`
	Time     *string `json:"time,omitempty"`
}

// ShoppingWithId defines model for Shopping_with_id.
type ShoppingWithId struct {
	// Embedded fields due to inline allOf schema
	// Embedded fields due to inline allOf schema
	ID *int `json:"ID,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/Shopping_params)
	ShoppingParams
	// Embedded struct due to allOf(#/components/schemas/Shop)
	Shop
}

// Success defines model for Success.
type Success struct {
	// Embedded struct due to allOf(#/components/schemas/Base)
	Base
	// Embedded fields due to inline allOf schema
	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}

// Date defines model for date.
type Date string

// ShoppingID defines model for shoppingID.
type ShoppingID int

// Base404 defines model for Base_404.
type Base404 struct {
	// Embedded struct due to allOf(#/components/schemas/Error_404)
	Error404
}

// Base405 defines model for Base_405.
type Base405 struct {
	// Embedded struct due to allOf(#/components/schemas/Error_405)
	Error405
}

// Base500 defines model for Base_500.
type Base500 struct {
	// Embedded struct due to allOf(#/components/schemas/Error_500)
	Error500
}

// ComingShoppings200 defines model for ComingShoppings_200.
type ComingShoppings200 struct {
	// Embedded struct due to allOf(#/components/schemas/Success)
	Success
	// Embedded fields due to inline allOf schema
	Data *[]ShoppingWithId `json:"data,omitempty"`
}

// Goods200 defines model for Goods_200.
type Goods200 struct {
	// Embedded struct due to allOf(#/components/schemas/Success)
	Success
	// Embedded struct due to allOf(#/components/schemas/Shopping_item)
	ShoppingItem
}

// Item200 defines model for Item_200.
type Item200 struct {
	// Embedded struct due to allOf(#/components/schemas/Success)
	Success
	// Embedded struct due to allOf(#/components/schemas/Shopping_item_params_with_id)
	ShoppingItemParamsWithId
}

// Item400 defines model for Item_400.
type Item400 struct {
	// Embedded struct due to allOf(#/components/schemas/Error_400)
	Error400
	// Embedded fields due to inline allOf schema
	Errors *struct {
		Validation *struct {
			CategoryID  *int    `json:"categoryID,omitempty"`
			Complete    *bool   `json:"complete,omitempty"`
			ListID      *int    `json:"listID,omitempty"`
			ProductName *string `json:"productName,omitempty"`
			Quantity    *int    `json:"quantity,omitempty"`
		} `json:"validation,omitempty"`
	} `json:"errors,omitempty"`
}

// LastShopping200 defines model for LastShopping_200.
type LastShopping200 struct {
	// Embedded struct due to allOf(#/components/schemas/Success)
	Success
	// Embedded struct due to allOf(#/components/schemas/Shopping_with_id)
	ShoppingWithId
}

// Shopping200 defines model for Shopping_200.
type Shopping200 struct {
	// Embedded struct due to allOf(#/components/schemas/Success)
	Success
	// Embedded fields due to inline allOf schema
	// Embedded struct due to allOf(#/components/schemas/Shopping)
	Shopping
}

// Shopping400 defines model for Shopping_400.
type Shopping400 struct {
	// Embedded struct due to allOf(#/components/schemas/Error_400)
	Error400
	// Embedded fields due to inline allOf schema
	Errors *struct {
		Validation *struct {
			Complete *bool   `json:"complete,omitempty"`
			Date     *string `json:"date,omitempty"`
			OwnerID  *int    `json:"ownerID,omitempty"`
			ShopID   *int    `json:"shopID,omitempty"`
			Sum      *int    `json:"sum,omitempty"`
			Time     *string `json:"time,omitempty"`
		} `json:"validation,omitempty"`
	} `json:"errors,omitempty"`
}

// ItemRequest defines model for Item_request.
type ItemRequest struct {
	// Embedded struct due to allOf(#/components/schemas/Shopping_item_params)
	ShoppingItemParams
}

// ShoppingRequest defines model for Shopping_request.
type ShoppingRequest struct {
	// Embedded struct due to allOf(#/components/schemas/Shopping)
	Shopping
}

// AddItemRequestBody defines body for AddItem for application/json ContentType.
type AddItemJSONRequestBody ItemRequest

// AddShoppingRequestBody defines body for AddShopping for application/json ContentType.
type AddShoppingJSONRequestBody ShoppingRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Добавление товара в покупку// (POST /addItem)
	AddItem(ctx echo.Context) error
	// Добавление покупки// (POST /addShopping)
	AddShopping(ctx echo.Context) error
	// Ближайшие 5 покупок// (GET /getComingShoppings/{date})
	GetComingShoppings(ctx echo.Context, date Date) error
	// Список покупок// (GET /getGoods/{shoppingID})
	GetGoods(ctx echo.Context, shoppingID ShoppingID) error
	// Последняя покупка// (GET /lastShopping)
	LastShopping(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AddItem converts echo context to params.
func (w *ServerInterfaceWrapper) AddItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddItem(ctx)
	return err
}

// AddShopping converts echo context to params.
func (w *ServerInterfaceWrapper) AddShopping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddShopping(ctx)
	return err
}

// GetComingShoppings converts echo context to params.
func (w *ServerInterfaceWrapper) GetComingShoppings(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "date" -------------
	var date Date

	err = runtime.BindStyledParameter("simple", false, "date", ctx.Param("date"), &date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComingShoppings(ctx, date)
	return err
}

// GetGoods converts echo context to params.
func (w *ServerInterfaceWrapper) GetGoods(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shoppingID" -------------
	var shoppingID ShoppingID

	err = runtime.BindStyledParameter("simple", false, "shoppingID", ctx.Param("shoppingID"), &shoppingID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shoppingID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetGoods(ctx, shoppingID)
	return err
}

// LastShopping converts echo context to params.
func (w *ServerInterfaceWrapper) LastShopping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LastShopping(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/addItem", wrapper.AddItem)
	router.POST("/addShopping", wrapper.AddShopping)
	router.GET("/getComingShoppings/:date", wrapper.GetComingShoppings)
	router.GET("/getGoods/:shoppingID", wrapper.GetGoods)
	router.GET("/lastShopping", wrapper.LastShopping)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
