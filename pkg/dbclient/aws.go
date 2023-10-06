package dbclient

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func awsAuthBuilder(ctx context.Context, cliCfg Config) (string, error) {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	return awsAuthBuilderWithConfig(ctx, cliCfg, cfg)
}

func awsAuthBuilderWithConfig(ctx context.Context, cliCfg Config, awsCfg aws.Config) (string, error) {

	parsed, err := url.Parse(cliCfg.DSN)
	info := parsed.User

	authToken, err := auth.BuildAuthToken(
		ctx,
		fmt.Sprintf("%s:%s", parsed.Host, parsed.Port()),
		awsCfg.Region,
		info.Username(),
		awsCfg.Credentials,
	)

	parsed.User = url.UserPassword(info.Username(), authToken)

	return parsed.String(), err

}

func GetRdsClientfromDefualtConfig() (*rds.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Couldn't load default configuration : ", err)
		return nil, err
	}
	return rds.NewFromConfig(sdkConfig), nil
}
