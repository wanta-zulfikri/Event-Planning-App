
package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
)

type ReviewController struct {
	n reviews.Service
}

func New(o reviews.Service) reviews.Handler {
	return &ReviewController{n: o}
}

func (rc *ReviewController) WriteReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing. ", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		var input RequestWriteReview
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		newReview := reviews.Core{
			EventID: input.EventID,
			Review:  input.Review, 
		
		}

		err = rc.n.WriteReview(newReview)
		if err != nil {
			c.Logger().Error("Failed to write a review:", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "write a review successfully", nil))
	}
}

func (rc *ReviewController) UpdateReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		var input RequestUpdateReview
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input:", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		// file, err := c.FormFile("image")
		// var image string
		// if err != nil && err != http.ErrMissingFile {
		// 	c.Logger().Error("Failed to get image from form file: ", err)
		// 	return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		// } else if file != nil {
		// 	image, err = helper.UploadImage(c, file)
		// 	if err != nil {
		// 		c.Logger().Error("Failed to upload image: ", err)
		// 		return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		// 	}
		// }

		updatedReview := reviews.Core{
			ID:      uint(id),
			Review:  	input.Review, 
		}
		err = rc.n.UpdateReview(updatedReview.ID, updatedReview)
		if err != nil {
			c.Logger().Error("Failed to update review: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))

		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Review updated successfully", nil))
	}
}

func (rc *ReviewController) DeleteReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))

		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))

		}

		err = rc.n.DeleteReview(uint(id))
		if err != nil {
			c.Logger().Error("Error delleting review", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success deletedc an review", nil))
	}
}
