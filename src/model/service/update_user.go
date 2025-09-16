package service

import (
	"github.com/vmstelecom/huncoding/src/configuration/rest_err"
	"github.com/vmstelecom/huncoding/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
