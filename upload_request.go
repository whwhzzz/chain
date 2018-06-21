package main

type UploadRequest struct {
	BuyerCert  string `json:"buyer_cert"`
	SellerSig  string `json:"seller_sig"`
	SellerCert string `json:"seller_cert"`
	OfferID    string `json:"offer_id"`
	TimeExpire int64  `json:"time_expire"`
	Payload    string `json:"payload"`
	Content    string `json:"content"`
}

type DownloadRequest struct {
	BuyerCert  string `json:"buyer_cert"`
	OfferID    string `json:"offer_id"`
	UseDeposit bool   `json:"use_deposit"`
}
