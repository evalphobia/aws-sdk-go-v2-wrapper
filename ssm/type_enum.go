package ssm

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/ssm"
)

type ParameterTier string

const (
	ParameterTierStandard           ParameterTier = ParameterTier(SDK.ParameterTierStandard)
	ParameterTierAdvanced           ParameterTier = ParameterTier(SDK.ParameterTierAdvanced)
	ParameterTierIntelligentTiering ParameterTier = ParameterTier(SDK.ParameterTierIntelligentTiering)
)

type ParameterType string

const (
	ParameterTypeString       ParameterType = ParameterType(SDK.ParameterTypeString)
	ParameterTypeStringList   ParameterType = ParameterType(SDK.ParameterTypeStringList)
	ParameterTypeSecureString ParameterType = ParameterType(SDK.ParameterTypeSecureString)
)
