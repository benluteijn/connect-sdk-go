// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ThreeDSecureResults represents class ThreeDSecureResults
type ThreeDSecureResults struct {
	Cavv                         *string           `json:"cavv,omitempty"`
	DirectoryServerTransactionID *string           `json:"directoryServerTransactionId,omitempty"`
	Eci                          *string           `json:"eci,omitempty"`
	SdkData                      *SdkDataOutput    `json:"sdkData,omitempty"`
	ThreeDSecureData             *ThreeDSecureData `json:"threeDSecureData,omitempty"`
	ThreeDSecureVersion          *string           `json:"threeDSecureVersion,omitempty"`
	ThreeDServerTransactionID    *string           `json:"threeDServerTransactionId,omitempty"`
	Xid                          *string           `json:"xid,omitempty"`
}

// NewThreeDSecureResults constructs a new ThreeDSecureResults
func NewThreeDSecureResults() *ThreeDSecureResults {
	return &ThreeDSecureResults{}
}
