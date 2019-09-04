package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db DB
}

func (router *Router) getGoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("shoppingId")
		shoppingId, err := strconv.ParseUint(param, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Bad param", param)})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": true, "shoppingId": shoppingId, "goods": router.db.getGoodsByShoppingId(shoppingId)})
		}
	}
}

func (router *Router) lastShopping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "shopping": router.db.getLastShopping()})
	}
}

func (router *Router) addShopping() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shopping Shopping
		err := c.BindJSON(&shopping)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}
		err = router.db.addShopping(&shopping)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "shoppingId": shopping.Id})
	}
}

func (router *Router) addItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productItem ShopList
		err := c.BindJSON(&productItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
		}
		err = router.db.addProductItem(&productItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "productItem": productItem})
	}
}

func (router *Router) getComingShoppings() gin.HandlerFunc {
	return func(c *gin.Context) {
		date := c.Param("date")
		c.JSON(http.StatusOK, gin.H{"success": true, "shoppings": router.db.getComingShoppings(date)})
	}
}
