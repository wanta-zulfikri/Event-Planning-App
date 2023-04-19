package handler

import (
	"Event-Planning-App/app/features/users"
	"Event-Planning-App/helper"
	"Event-Planning-App/middlewares"
	"net/http"
	"strconv"

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
			c.Logger().Error(err.Error())
			code, res := helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil)
			return c.JSON(code, res)
		}

		res, err := uc.s.Login(input.Email, input.Password)
		if err != nil {
			if err.Error() == "Email not found" || err.Error() == "Invalid password" {
				code, res := helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found. Please check your email and password input.", nil)
				return c.JSON(code, res)
			}
			c.Logger().Error(err.Error())
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil)
			return c.JSON(code, res)
		}

		token, _ := middlewares.CreateToken(res.ID)
		result := &LoginResponse{Token: token}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful login, please use this token for further access.",
			Data:    result,
		})
	}
}

func (uc *UserController) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := int(middlewares.ExtractToken(c))
		data, err := uc.s.GetProfile(userID)
		if err != nil {
			c.Logger().Error("User not found", err.Error())
			code, res := helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found. Please check your email and password input.", nil)
			return c.JSON(code, res)
		}
		res := UserResponse{}
		copier.Copy(&res, &data)
		return c.JSON(helper.SuccessResponse(http.StatusOK, "profile successfully displayed", res))
	}
}

func (uc *UserController) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		var req UserUpdateInput
		if err := c.Bind(&req); err != nil {
			c.Logger().Errorf("Error binding request: %v", err)
			code, res := helper.ResponseFormat(http.StatusBadRequest, "Bad Request: "+err.Error(), nil)
			return c.JSON(code, res)
		}
		if err := uc.s.UpdateProfile(id, req.Username, req.Email, req.Password); err != nil {
			c.Logger().Errorf("Error updating profile: %v", err)
			code, res := helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error: "+err.Error(), nil)
			return c.JSON(code, res)
		}
		code, res := helper.ResponseFormat(http.StatusOK, "Profile updated successfully", nil)
		return c.JSON(code, res)
	}
}

func (uc *UserController) DeleteProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error("Failed to parse user ID", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
		}

		if err := uc.s.DeleteProfile(uint(userID)); err != nil {
			c.Logger().Error("Failed to delete profile", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success delete profile",
		})
	}
}
