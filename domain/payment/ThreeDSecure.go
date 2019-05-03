// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ThreeDSecure represents class ThreeDSecure
type ThreeDSecure struct {
	AuthenticationFlow                   *string                               `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize                  *string                               `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator                   *string                               `json:"challengeIndicator,omitempty"`
	ExternalCardholderAuthenticationData *ExternalCardholderAuthenticationData `json:"externalCardholderAuthenticationData,omitempty"`
	PriorThreeDSecureData                *ThreeDSecureData                     `json:"priorThreeDSecureData,omitempty"`
	RedirectionData                      *RedirectionData                      `json:"redirectionData,omitempty"`
	SdkData                              *SdkDataInput                         `json:"sdkData,omitempty"`
	SkipAuthentication                   *bool                                 `json:"skipAuthentication,omitempty"`
}

// NewThreeDSecure constructs a new ThreeDSecure
func NewThreeDSecure() *ThreeDSecure {
	return &ThreeDSecure{}
}
