package store

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Router - basic route func type
type Router struct {
	DB DB
}

// GetGoods returns all products by shoppingId
func (router *Router) GetGoods() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("shoppingID")
		shoppingID, err := strconv.ParseUint(param, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Bad param", param)})
		}

		return c.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shoppingID, "goods": router.DB.GetGoodsByShoppingID(shoppingID)})
	}
}

// LastShopping returns LastShopping information
func (router *Router) LastShopping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"success": true, "shopping": router.DB.GetLastShopping()})
	}
}

// AddShopping inserts new shopping
func (router *Router) AddShopping() echo.HandlerFunc {
	return func(c echo.Context) error {
		var shopping Shopping
		if err := c.Bind(&shopping); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}

		if err := router.DB.AddShopping(&shopping); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		return c.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shopping.ID})
	}
}

// AddItem inserts new product to shopping cart
func (router *Router) AddItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var productItem ShopList
		if err := c.Bind(&productItem); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}

		if err := router.DB.AddProductItem(&productItem); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		return c.JSON(http.StatusOK, echo.Map{"success": true, "productItem": productItem})
	}
}

// GetComingShoppings returns coming shoppings
func (router *Router) GetComingShoppings() echo.HandlerFunc {
	return func(c echo.Context) error {
		date := c.Param("date")
		return c.JSON(http.StatusOK, echo.Map{"success": true, "shoppings": router.DB.GetComingShoppings(date)})
	}
}
