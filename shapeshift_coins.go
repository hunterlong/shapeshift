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
	FIRST struct {
		Coin
	} `json:"1ST"`
	ANT struct {
		Coin
	} `json:"ANT"`
	BAT struct {
		Coin
	} `json:"BAT"`
	BCH struct {
		Coin
	} `json:"BCH"`
	BCY struct {
		Coin
	} `json:"BCY"`
	BLK struct {
		Coin
	} `json:"BLK"`
	BNT struct {
		Coin
	} `json:"BNT"`
	BTC struct {
		Coin
	} `json:"BTC"`
	BTCD struct {
		Coin
	} `json:"BTCD"`
	BTG struct {
		Coin
	} `json:"BTG"`
	BTS struct {
		Coin
	} `json:"BTS"`
	CLAM struct {
		Coin
	} `json:"CLAM"`
	CVC struct {
		Coin
	} `json:"CVC"`
	DASH struct {
		Coin
	} `json:"DASH"`
	DCR struct {
		Coin
	} `json:"DCR"`
	DGB struct {
		Coin
	} `json:"DGB"`
	DNT struct {
		Coin
	} `json:"DNT"`
	DOGE struct {
		Coin
	} `json:"DOGE"`
	EDG struct {
		Coin
	} `json:"EDG"`
	EMC struct {
		Coin
	} `json:"EMC"`
	EOS struct {
		Coin
	} `json:"EOS"`
	ETC struct {
		Coin
	} `json:"ETC"`
	ETH struct {
		Coin
	} `json:"ETH"`
	FCT struct {
		Coin
	} `json:"FCT"`
	FUN struct {
		Coin
	} `json:"FUN"`
	GAME struct {
		Coin
	} `json:"GAME"`
	GNO struct {
		Coin
	} `json:"GNO"`
	GNT struct {
		Coin
	} `json:"GNT"`
	GUP struct {
		Coin
	} `json:"GUP"`
	KMD struct {
		Coin
	} `json:"KMD"`
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
	MLN struct {
		Coin
	} `json:"MLN"`
	MONA struct {
		Coin
	} `json:"MONA"`
	MSC struct {
		Coin
	} `json:"MSC"`
	MTL struct {
		Coin
	} `json:"MTL"`
	NBT struct {
		Coin
	} `json:"NBT"`
	NEO struct {
		Coin
	} `json:"NEO"`
	NMC struct {
		Coin
	} `json:"NMC"`
	NMR struct {
		Coin
	} `json:"NMR"`
	NVC struct {
		Coin
	} `json:"NVC"`
	NXT struct {
		Coin
	} `json:"NXT"`
	OMG struct {
		Coin
	} `json:"OMG"`
	POT struct {
		Coin
	} `json:"POT"`
	PPC struct {
		Coin
	} `json:"PPC"`
	QTUM struct {
		Coin
	} `json:"QTUM"`
	RCN struct {
		Coin
	} `json:"RCN"`
	RDD struct {
		Coin
	} `json:"RDD"`
	REP struct {
		Coin
	} `json:"REP"`
	RLC struct {
		Coin
	} `json:"RLC"`
	SALT struct {
		Coin
	} `json:"SALT"`
	SC struct {
		Coin
	} `json:"SC"`
	SNGLS struct {
		Coin
	} `json:"SNGLS"`
	SNT struct {
		Coin
	} `json:"SNT"`
	START struct {
		Coin
	} `json:"START"`
	STEEM struct {
		Coin
	} `json:"STEEM"`
	STORJ struct {
		Coin
	} `json:"STORJ"`
	SWT struct {
		Coin
	} `json:"SWT"`
	TRST struct {
		Coin
	} `json:"TRST"`
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
	WAVES struct {
		Coin
	} `json:"WAVES"`
	WINGS struct {
		Coin
	} `json:"WINGS"`
	XCP struct {
		Coin
	} `json:"XCP"`
	XEM struct {
		Coin
	} `json:"XEM"`
	XMR struct {
		Coin
	} `json:"XMR"`
	XRP struct {
		Coin
	} `json:"XRP"`
	ZEC struct {
		Coin
	} `json:"ZEC"`
	ZRX struct {
		Coin
	} `json:"ZRX"`
}
