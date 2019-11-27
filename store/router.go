package store

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Router - basic route func type
type Router struct {
	DB DB
}

// GetGoods returns all products by shoppingId
func (router *Router) GetGoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("shoppingID")
		shoppingID, err := strconv.ParseUint(param, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Bad param", param)})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": true, "shoppingID": shoppingID, "goods": router.DB.GetGoodsByShoppingID(shoppingID)})
		}
	}
}

// LastShopping returns LastShopping information
func (router *Router) LastShopping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "shopping": router.DB.GetLastShopping()})
	}
}

// AddShopping inserts new shopping
func (router *Router) AddShopping() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shopping Shopping
		err := c.BindJSON(&shopping)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}
		err = router.DB.AddShopping(&shopping)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "shoppingID": shopping.ID})
	}
}

// AddItem inserts new product to shopping cart
func (router *Router) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productItem ShopList
		err := c.BindJSON(&productItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}
		err = router.DB.AddProductItem(&productItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "productItem": productItem})
	}
}

// GetComingShoppings returns coming shoppings
func (router *Router) GetComingShoppings() gin.HandlerFunc {
	return func(c *gin.Context) {
		date := c.Param("date")
		c.JSON(http.StatusOK, gin.H{"success": true, "shoppings": router.DB.GetComingShoppings(date)})
	}
}
