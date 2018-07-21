package api

import (
	"github.com/evanxzj/echo-server/utils"

	"github.com/labstack/echo"
)

type userRoutes struct {
	route
}

/*

	Map users routes

*/
func (ur userRoutes) mount() {

	ur.group.GET("/:id", ur.getUser)
	ur.group.GET("/", ur.getAll)
	ur.group.POST("/", ur.createUser)

}

/*

	Fetch a user by id

*/
func (ur userRoutes) getUser(c echo.Context) error {

	id, err := utils.ConvertID(c.Param("id"))

	if err != nil {
		return err
	}

	user := ur.db.Users.GetUserByID(id)

	return c.JSON(200, user)

}

/*

	Fetch all users

*/
func (ur userRoutes) getAll(c echo.Context) error {

	return c.JSON(200, ur.db.Users.GetAll())

}

/*

	Create a user

*/
func (ur userRoutes) createUser(c echo.Context) error {

	username := c.FormValue("username")
	first := c.FormValue("first")
	last := c.FormValue("last")
	pw := c.FormValue("password")

	ur.db.Users.Create(username, first, last, pw)
	resp := utils.NewResponse("ok", true)

	return c.JSON(200, resp.Payload)

}
