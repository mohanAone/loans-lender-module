package business

import (
	"context"
	"github.com/angel-one/make init/constants"
	"github.com/angel-one/make init/external"
	"github.com/angel-one/make init/models"
	"github.com/angel-one/make init/utils/configs"
)

// GetMoxy is used to get the moxy response
func GetMoxy(ctx context.Context) (models.MoxyResponse, error) {
	response := models.MoxyResponse{}

	moxyConfig, err := configs.Get(constants.MoxyConfig)
	if err != nil {
		return response, err
	}

	data, err := external.GetMoxy(ctx, moxyConfig.GetString(constants.URLConfigKey))
	if err != nil {
		return response, err
	}

	response.Data = data

	return response, nil
}
