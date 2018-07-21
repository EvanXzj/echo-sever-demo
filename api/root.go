package api

import (
	"fmt"
	"time"

	"github.com/evanxzj/echo-server/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type rootRoutes struct {
	route
}

/*

	Map root routes

*/
func (r rootRoutes) mount() {

	r.group.GET("/", r.getRoot)
	r.group.POST("/login", r.login)

}

/*

	Serve the root route

*/
func (r rootRoutes) getRoot(c echo.Context) error {

	return c.String(200, "this is the root route")

}

/*

	Login with username / password and get an auth token

*/
func (r rootRoutes) login(c echo.Context) error {

	jwtSecret := utils.GetEnv("JWT_SECRET")

	un := c.FormValue("username")
	pw := c.FormValue("password")

	fmt.Printf(un)
	fmt.Printf(pw)

	user, err := r.db.Users.Login(un, pw)

	if err != nil {
		fmt.Println(err)
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.First + " " + user.Last
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return err
	}

	// attach token value to response
	resp := utils.NewResponse("ok", true)
	resp.Set("token", t)

	return c.JSON(200, resp.Payload)

}
