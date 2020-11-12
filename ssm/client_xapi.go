package ssm

import (
	"context"
	"strings"
)

// XGetParameterValue fetches a parameter value from SSM.
func (svc *SSM) XGetParameterValue(ctx context.Context, name string) (value string, notFound bool, err error) {
	return svc.xGetParameterValue(ctx, name, false)
}

// XGetParameterValueWithDecryption fetches a decrypted value.
func (svc *SSM) XGetParameterValueWithDecryption(ctx context.Context, name string) (value string, notFound bool, err error) {
	return svc.xGetParameterValue(ctx, name, true)
}

func (svc *SSM) xGetParameterValue(ctx context.Context, name string, decryption bool) (value string, notFound bool, err error) {
	result, err := svc.GetParameter(ctx, GetParameterRequest{
		Name:           name,
		WithDecryption: decryption,
	})
	if err != nil {
		return "", false, err
	}

	return result.Value, result.NotFound, nil
}

// XGetParameterValueList fetches parameter value list from SSM.
func (svc *SSM) XGetParameterValueList(ctx context.Context, list []string) (values map[string]string, err error) {
	return svc.xGetParameterValueList(ctx, list, false)
}

// XGetParameterValueListWithDecryption fetches decrypted value list.
func (svc *SSM) XGetParameterValueListWithDecryption(ctx context.Context, list []string) (values map[string]string, err error) {
	return svc.xGetParameterValueList(ctx, list, true)
}

func (svc *SSM) xGetParameterValueList(ctx context.Context, list []string, decryption bool) (values map[string]string, err error) {
	result, err := svc.GetParameters(ctx, GetParametersRequest{
		Names:          list,
		WithDecryption: decryption,
	})
	m := make(map[string]string)
	if err != nil {
		return m, err
	}

	for _, v := range result.Parameters {
		m[v.Name] = v.Value
	}
	return m, nil
}

// XPutParameter puts the value to SSM parameter store.
func (svc *SSM) XPutParameter(ctx context.Context, name, value string) (alreadyExists bool, err error) {
	result, err := svc.PutParameter(ctx, PutParameterRequest{
		Name:  name,
		Value: value,
		Type:  ParameterTypeString,
	})
	if err != nil {
		return false, err
	}
	return result.AlreadyExists, nil
}

// XPutParameterList puts the value set (StringList type) to SSM parameter store.
func (svc *SSM) XPutParameterList(ctx context.Context, name string, values []string) (alreadyExists bool, err error) {
	result, err := svc.PutParameter(ctx, PutParameterRequest{
		Name:  name,
		Value: strings.Join(values, ","),
		Type:  ParameterTypeStringList,
	})
	if err != nil {
		return false, err
	}
	return result.AlreadyExists, nil
}

// XPutParameterSecrure puts the value to SSM parameter store with encryption by default AWS KMS key.
func (svc *SSM) XPutParameterSecrure(ctx context.Context, name, value string) (alreadyExists bool, err error) {
	result, err := svc.PutParameter(ctx, PutParameterRequest{
		Name:  name,
		Value: value,
		Type:  ParameterTypeSecureString,
	})
	if err != nil {
		return false, err
	}
	return result.AlreadyExists, nil
}

// XPutParameterSecureCMK puts the value to SSM parameter store with encryption by your AWS KMS key.
func (svc *SSM) XPutParameterSecureCMK(ctx context.Context, name, value, keyID string) (alreadyExists bool, err error) {
	result, err := svc.PutParameter(ctx, PutParameterRequest{
		Name:  name,
		Value: value,
		KeyID: keyID,
		Type:  ParameterTypeSecureString,
	})
	if err != nil {
		return false, err
	}
	return result.AlreadyExists, nil
}
