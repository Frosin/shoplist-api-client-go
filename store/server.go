package store

import (
	"fmt"
	"net/http"

	"github.com/Frosin/shoplist-api-client-go/api"

	"github.com/labstack/echo/v4"
)

// Server - basic route func type
type Server struct {
	DB DB
}

// GetGoods returns all products by shoppingId
func (s *Server) GetGoods(ctx echo.Context, shoppingID api.ShoppingID) error {
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shoppingID, "goods": s.DB.GetGoodsByShoppingID(uint64(shoppingID))})
}

// LastShopping returns LastShopping information
func (s *Server) LastShopping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shopping": s.DB.GetLastShopping()})
}

// AddShopping inserts new shopping
func (s *Server) AddShopping(ctx echo.Context) error {
	var shopping Shopping
	if err := ctx.Bind(&shopping); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
	}

	if err := s.DB.AddShopping(&shopping); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shopping.ID})
}

// AddItem inserts new product to shopping cart
func (s *Server) AddItem(ctx echo.Context) error {
	var productItem ShopList
	if err := ctx.Bind(&productItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
	}

	if err := s.DB.AddProductItem(&productItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "productItem": productItem})
}

// GetComingShoppings returns coming shoppings
func (s *Server) GetComingShoppings(ctx echo.Context, date api.Date) error {
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppings": s.DB.GetComingShoppings(string(date))})
}
