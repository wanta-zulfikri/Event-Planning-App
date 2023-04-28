package handler

import (
	"net/http"

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
		err := uc.s.Register(users.Core{Username: input.Username, Email: input.Email, Phone: input.Phone, Password: input.Password})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success Created an Account", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		user, err := uc.s.Login(input.Email, input.Password)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		token, err := middlewares.CreateJWT(user.ID, user.Email, user.Username)
		if err != nil {
			c.Logger().Error(err.Error())
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
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID
		data, err := uc.s.GetProfile(uint(id))
		if err != nil {
			c.Logger().Error(err.Error())
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
		var input UpdateInput
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		image, err := helper.UploadImage(c, file)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		updatedUser := users.Core{
			ID:       uint(id),
			Username: input.Username,
			Email:    input.Email,
			Phone:    input.Phone,
			Password: input.Password,
			Image:    image,
		}

		err = uc.s.UpdateProfile(uint(id), updatedUser)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success Updated an Account", nil))
	}
}

func (uc *UserController) DeleteProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID
		err = uc.s.DeleteProfile(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Deleted an Account", nil))
	}
}
