package shopeego

// CommonInfoRequest request
type CommonInfoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}
