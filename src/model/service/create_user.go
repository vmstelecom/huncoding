package service

import (
	"fmt"

	"github.com/vmstelecom/huncoding/src/configuration/logger"
	"github.com/vmstelecom/huncoding/src/configuration/rest_err"
	"github.com/vmstelecom/huncoding/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
