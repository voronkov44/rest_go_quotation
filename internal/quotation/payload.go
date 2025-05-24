package quotation

type CreateQuotationResponse struct {
	ID     int    `json:"id" example:"1"`
	Author string `json:"author" example:"Sun Tzu"`
	Quote  string `json:"quote" example:"Appear weak when you are strong, and strong when you are weak."`
}

type CreateQuotationRequest struct {
	Author string `json:"author" example:"Sun Tzu"`
	Quote  string `json:"quote" example:"Appear weak when you are strong, and strong when you are weak."`
}
