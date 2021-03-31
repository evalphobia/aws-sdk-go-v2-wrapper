package pinpoint

import (
	"context"
)

// XUpdateEndpointEmail updates email endpoint.
// (attrs[0]=Attributes, attrs[1]=UserAttributes)
func (svc *Pinpoint) XUpdateEndpointEmail(ctx context.Context, appID, endpointID, userID, email string, attrs ...map[string][]string) error {
	req := UpdateEndpointRequest{
		ApplicationID: appID,
		EndpointID:    endpointID,
		EndpointRequest: EndpointRequest{
			Address:     email,
			ChannelType: ChannelTypeEmail,
			User: EndpointUser{
				UserID: userID,
			},
		},
	}

	attrSize := len(attrs)
	switch {
	case attrSize >= 2:
		req.EndpointRequest.Attributes = attrs[0]
		req.EndpointRequest.User.UserAttributes = attrs[1]
	case attrSize == 1:
		req.EndpointRequest.Attributes = attrs[0]
	}

	_, err := svc.UpdateEndpoint(ctx, req)
	return err
}
