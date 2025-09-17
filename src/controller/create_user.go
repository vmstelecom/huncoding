package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmstelecom/huncoding/src/configuration/logger"
	"github.com/vmstelecom/huncoding/src/configuration/validation"
	"github.com/vmstelecom/huncoding/src/controller/model/request"
	"github.com/vmstelecom/huncoding/src/model"
	"github.com/vmstelecom/huncoding/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trting to validate user info", err,
			zap.String("journey", "createUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}
