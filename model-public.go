package shopeego

type GetShopsByPartnerRequest struct {
	//
	RequestLiteBase
}

type GetShopsByPartnerResponse struct {
	// A list of shops authorized to the partner.
	AuthedShops []Shop `json:"authed_shops"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetCategoriesByCountryRequest struct {
	//
	RequestLiteBase
	// Two-digit country code(capital letter).
	Country string `json:"country"`
	// Is cross border or not. 1: cross border; 0: not cross border
	IsCB bool `json:"is_cb"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language"`
}

type GetCategoriesByCountryResponse struct {
	// List of categories info.
	Categories []Category `json:"categories"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetPaymentListRequest struct {
	//
	RequestLiteBase
	// ID/MY/SG/VN/PH/TH/TW
	Country string `json:"country"`
}

type GetPaymentListResponse struct {
	//
	PaymentMethodList []Method `json:"payment_method_list"`
}