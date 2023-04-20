package handler

import (
	"Event-Planning-App/app/features/users"
	"Event-Planning-App/helper"
	"Event-Planning-App/middlewares"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	s users.Service
}

func New(h users.Service) users.Handler {
	return &UserController{s: h}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			code, res := helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil)
			return c.JSON(code, res)
		}
		err := uc.s.Register(users.Core{Username: input.Username, Email: input.Email, Password: input.Password})
		if err != nil {
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil)
			return c.JSON(code, res)
		}
		code, res := helper.ResponseFormat(http.StatusCreated, "Success Created an Account", nil)
		return c.JSON(code, res)
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			code, res := helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil)
			return c.JSON(code, res)
		}

		user, err := uc.s.Login(input.Email, input.Password)
		if err != nil {
			c.Logger().Error("Failed to login: ", err)
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil)
			return c.JSON(code, res)
		}

		token, err := middlewares.CreateJWT(user.Email)
		if err != nil {
			c.Logger().Error("Failed to create JWT token: ", err)
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil)
			return c.JSON(code, res)
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful login, please use this token for further access.",
			Data:    map[string]interface{}{"token": token},
		})
	}
}

func (uc *UserController) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil)
			return c.JSON(code, res)
		}

		email, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil)
			return c.JSON(code, res)
		}

		data, err := uc.s.GetProfile(email)
		if err != nil {
			c.Logger().Error("User not found", err.Error())
			code, res := helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found. Please check your email and password input.", nil)
			return c.JSON(code, res)
		}
		res := UserResponse{}
		copier.Copy(&res, &data)
		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    map[string]interface{}{"res": res},
		})
	}
}

func (uc *UserController) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil)
			return c.JSON(code, res)
		}

		email, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil)
			return c.JSON(code, res)
		}

		var input UpdateInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			code, res := helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil)
			return c.JSON(code, res)
		}

		err = uc.s.UpdateProfile(email, input.Username, input.Email, input.Password)
		if err != nil {
			c.Logger().Error("Failed to update profile: ", err)
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil)
			return c.JSON(code, res)
		}

		code, res := helper.ResponseFormat(http.StatusOK, "Profile updated successfully", nil)
		return c.JSON(code, res)
	}
}

func (uc *UserController) DeleteProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil)
			return c.JSON(code, res)
		}

		email, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			code, res := helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil)
			return c.JSON(code, res)
		}

		err = uc.s.DeleteProfile(email)
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Account deleted successfully",
		})
	}
}
