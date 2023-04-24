package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"

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
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}
		err := uc.s.Register(users.Core{Username: input.Username, Email: input.Email, Password: input.Password})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success created an account", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		user, err := uc.s.Login(input.Email, input.Password)
		if err != nil {
			c.Logger().Error("Failed to login: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		token, err := middlewares.CreateJWT(user.ID, user.Email, user.Username)
		if err != nil {
			c.Logger().Error("Failed to create JWT token: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
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
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		idString, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(fmt.Sprint(idString), 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		data, err := uc.s.GetProfile(uint(id))
		if err != nil {
			c.Logger().Error("User not found", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found. Please check your email and password input.", nil))
		}
		res := UserResponse{}
		copier.Copy(&res, &data)
		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    res,
		})
	}
}

func (uc *UserController) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		idString, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(fmt.Sprint(idString), 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		var input UpdateInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.Logger().Error("Failed to get image: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		image, err := helper.UploadImage(c, file)
		if err != nil {
			c.Logger().Error("Failed to upload image: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		err = uc.s.UpdateProfile(uint(id), input.Username, input.Email, input.Password, image)
		if err != nil {
			c.Logger().Error("Failed to update profile: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success updated an account", nil))
	}
}

func (uc *UserController) DeleteProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		idString, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(fmt.Sprint(idString), 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		err = uc.s.DeleteProfile(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an account", nil))
	}
}
