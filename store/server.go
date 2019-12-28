package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Frosin/shoplist-api-client-go/api"
	"github.com/Frosin/shoplist-api-client-go/store/sqlc"

	"github.com/labstack/echo/v4"
)

var (
	messageSuccess = "success"
)

// Server - basic route func type
type Server struct {
	Version string
	DB      DB
	Queries *sqlc.Queries
}

// GetGoods returns all products by shoppingId
func (s *Server) GetGoods(ctx echo.Context, shoppingID api.ShoppingID) error {

	response200 := func(items *[]api.ShoppingItem) error {
		var response api.Goods200
		response.Message = messageSuccess
		response.Data = items
		return ctx.JSON(http.StatusOK, response)
	}
	response400 := func(err error) error {

	}
	response404 := func(err error) error {

	}
	response405 := func(err error) error {

	}
	response500 := func(err error) error {

	}

	goods, err := s.Queries.GetGoodsByShoppingID(context.Background(), *int32ToNullInt32(int32(shoppingID)))
	if err != nil {
		// 500 case
	}
	items := goodsToShoppingItems(goods)

	//return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shoppingID, "goods": s.DB.GetGoodsByShoppingID(uint64(shoppingID))})
}

func goodsToShoppingItems(goods []sqlc.ShopList) (shoppingItems []api.ShoppingItem) {
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

		item.ID = &id
		item.CategoryID = &category
		item.Complete = &complete
		item.ListID = &listID
		item.ProductName = &productName
		item.Quantity = &quantity
		shoppingItems = append(shoppingItems, item)
	}
	return
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

func int32ToNullInt32(i int32) *sql.NullInt32 {
	return &sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// // GetGoods returns all products by shoppingId
// func (s *Server) GetGoods(ctx echo.Context, shoppingID api.ShoppingID) error {
// 	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shoppingID, "goods": s.DB.GetGoodsByShoppingID(uint64(shoppingID))})
// }

// // LastShopping returns LastShopping information
// func (s *Server) LastShopping(ctx echo.Context) error {
// 	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shopping": s.DB.GetLastShopping()})
// }

// // AddShopping inserts new shopping
// func (s *Server) AddShopping(ctx echo.Context) error {
// 	var shopping Shopping
// 	if err := ctx.Bind(&shopping); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
// 	}

// 	if err := s.DB.AddShopping(&shopping); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
// 	}
// 	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppingID": shopping.ID})
// }

// // AddItem inserts new product to shopping cart
// func (s *Server) AddItem(ctx echo.Context) error {
// 	var productItem ShopList
// 	if err := ctx.Bind(&productItem); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "Parse error", err)})
// 	}

// 	if err := s.DB.AddProductItem(&productItem); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": fmt.Sprintf("%s:%s", "DB insert error", err)})
// 	}
// 	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "productItem": productItem})
// }

// // GetComingShoppings returns coming shoppings
// func (s *Server) GetComingShoppings(ctx echo.Context, date api.Date) error {
// 	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "shoppings": s.DB.GetComingShoppings(string(date))})
// }
