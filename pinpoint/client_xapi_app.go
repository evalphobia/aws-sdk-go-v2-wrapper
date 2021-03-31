package pinpoint

import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/errors"
)

// XExistsApp checks if the application already exists or not.
func (svc *Pinpoint) XExistsApp(ctx context.Context, appID string) (bool, error) {
	_, err := svc.GetApp(ctx, GetAppRequest{
		ApplicationID: appID,
	})
	if err == nil {
		return true, nil
	}

	if e, ok := err.(errors.ErrorData); ok {
		if e.GetAWSErrCode() == "NotFoundException" {
			return false, nil
		}
	}
	return false, err
}

// XGetAppIDByName searches app by name and returns its ID.
func (svc *Pinpoint) XGetAppIDByName(ctx context.Context, appName string) (appID string, err error) {
	nextToken := ""
	for {
		resp, err := svc.GetApps(ctx, GetAppsRequest{
			PageSize: "1000",
			Token:    nextToken,
		})
		if err != nil {
			return "", err
		}
		for _, v := range resp.Item {
			if v.Name == appName {
				return v.ID, nil
			}
		}
		if resp.NextToken == "" {
			return "", nil
		}
		nextToken = resp.NextToken
	}
}
