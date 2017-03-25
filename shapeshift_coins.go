package shapeshift

type Coin struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	Image           string `json:"image,omitempty"`
	SpecialReturn   bool   `json:"specialReturn,omitempty"`
	SpecialOutgoing bool   `json:"specialOutgoing,omitempty"`
	SpecialIncoming bool   `json:"specialIncoming,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	FieldKey        string `json:"fieldKey,omitempty"`
	QrName          string `json:"qrName,omitempty"`
}

type CoinsResponse struct {
	BTC struct {
		Coin
	} `json:"BTC"`
	BCY struct {
		Coin
	} `json:"BCY"`
	BLK struct {
		Coin
	} `json:"BLK"`
	BTCD struct {
		Coin
	} `json:"BTCD"`
	BTS struct {
		Coin
	} `json:"BTS"`
	CLAM struct {
		Coin
	} `json:"CLAM"`
	DASH struct {
		Coin
	} `json:"DASH"`
	DGB struct {
		Coin
	} `json:"DGB"`
	DGD struct {
		Coin
	} `json:"DGD"`
	DOGE struct {
		Coin
	} `json:"DOGE"`
	EMC struct {
		Coin
	} `json:"EMC"`
	ETH struct {
		Coin
	} `json:"ETH"`
	ETC struct {
		Coin
	} `json:"ETC"`
	FCT struct {
		Coin
	} `json:"FCT"`
	GNT struct {
		Coin
	} `json:"GNT"`
	LBC struct {
		Coin
	} `json:"LBC"`
	LSK struct {
		Coin
	} `json:"LSK"`
	LTC struct {
		Coin
	} `json:"LTC"`
	MAID struct {
		Coin
	} `json:"MAID"`
	MONA struct {
		Coin
	} `json:"MONA"`
	MSC struct {
		Coin
	} `json:"MSC"`
	NBT struct {
		Coin
	} `json:"NBT"`
	NMC struct {
		Coin
	} `json:"NMC"`
	NVC struct {
		Coin
	} `json:"NVC"`
	NXT struct {
		Coin
	} `json:"NXT"`
	POT struct {
		Coin
	} `json:"POT"`
	PPC struct {
		Coin
	} `json:"PPC"`
	REP struct {
		Coin
	} `json:"REP"`
	RDD struct {
		Coin
	} `json:"RDD"`
	SDC struct {
		Coin
	} `json:"SDC"`
	SC struct {
		Coin
	} `json:"SC"`
	SJCX struct {
		Coin
	} `json:"SJCX"`
	START struct {
		Coin
	} `json:"START"`
	STEEM struct {
		Coin
	} `json:"STEEM"`
	SNGLS struct {
		Coin
	} `json:"SNGLS"`
	USDT struct {
		Coin
	} `json:"USDT"`
	VOX struct {
		Coin
	} `json:"VOX"`
	VRC struct {
		Coin
	} `json:"VRC"`
	VTC struct {
		Coin
	} `json:"VTC"`
	XCP struct {
		Coin
	} `json:"XCP"`
	XMR struct {
		Coin
	} `json:"XMR"`
	XRP struct {
		Coin
	} `json:"XRP"`
	ZEC struct {
		Coin
	} `json:"ZEC"`
}
