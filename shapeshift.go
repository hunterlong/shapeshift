package shapeshift

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/shopspring/decimal"
)

var apiUrl string = "https://shapeshift.io"

type Pair struct {
	Name string `json:"pair,omitempty"`
}

type RateResponse struct {
	Pair string `json:"pair,omitempty"`
	Rate decimal.Decimal `json:"rate"`
	Error
}

type LimitResponse struct {
	Pair  string `json:"pair,omitempty"`
	Limit decimal.Decimal `json:"limit"`
	Error
}

type MarketInfoResponse struct {
	Pair     string  `json:"pair,omitempty"`
	Rate     decimal.Decimal `json:"rate,omitempty"`
	Limit    decimal.Decimal `json:"limit,omitempty"`
	Min      decimal.Decimal `json:"min,omitempty"`
	MinerFee decimal.Decimal `json:"minerFee,omitempty"`
	Error
}

type RecentTranxResponse []struct {
	CurIn     string  `json:"curIn"`
	CurOut    string  `json:"curOut"`
	Timestamp decimal.Decimal `json:"timestamp"`
	Amount    decimal.Decimal `json:"amount"`
	Error
}

type DepositStatusResponse struct {
	Status       string  `json:"status"`
	Address      string  `json:"address"`
	Withdraw     string  `json:"withdraw,omitempty"`
	IncomingCoin decimal.Decimal `json:"incomingCoin,omitempty"`
	IncomingType string  `json:"incomingType,omitempty"`
	OutgoingCoin string  `json:"outgoingCoin,omitempty"`
	OutgoingType string  `json:"outgoingType,omitempty"`
	Transaction  string  `json:"transaction,omitempty"`
	Error
}

type Receipt struct {
	Email         string `json:"email"`
	TransactionID string `json:"txid"`
}

type ValidateResponse struct {
	Valid bool `json:"isValid"`
	Error
}

type CancelResponse struct {
	Success string `json:"success,omitempty"`
	Error
}

type Address struct {
	Id string `json:"address"`
}

type New struct {
	Pair        string  `json:"pair,omitempty"`
	ToAddress   string  `json:"withdrawal"`
	FromAddress string  `json:"returnAddress,omitempty"`
	DestTag     string  `json:"destTag,omitempty"`
	RsAddress   string  `json:"rsAddress,omitempty"`
	ApiKey      string  `json:"apiKey,omitempty"`
	Amount      decimal.Decimal `json:"amount,omitempty"`
}

type NewTransactionResponse struct {
	SendTo     string `json:"deposit"`
	SendType   string `json:"depositType"`
	ReturnTo   string `json:"withdrawal"`
	ReturnType string `json:"withdrawalType"`
	Public     string `json:"public"`
	XrpDestTag string `json:"xrpDestTag"`
	ApiKey     string `json:"apiPubKey"`
	Error
}

type FixedTransactionResponse struct {
	Response NewFixedTransactionResponse `json:"success"`
	Error
}

type NewFixedTransactionResponse struct {
	OrderID          string  `json:"orderId"`
	Pair             string  `json:"pair,omitempty"`
	Withdrawal       string  `json:"withdrawal"`
	WithdrawalAmount string  `json:"withdrawalAmount"`
	Deposit          string  `json:"deposit"`
	DepositAmount    string  `json:"depositAmount"`
	Expiration       int64   `json:"expiration"`
	QuotedRate       string  `json:"quotedRate"`
	MaxLimit         decimal.Decimal `json:"maxLimit"`
	ReturnAddress    string  `json:"returnAddress"`
	APIPubKey        string  `json:"apiPubKey"`
	MinerFee         string  `json:"minerFee"`
	Error
}

type ListTransactionsAPIResponse struct {
	Transactions []Transaction
	Error
}

type ErrorMsg interface {
	ErrorMsg() string
	isOk() bool
}

func (e Error) ErrorMsg() string {
	return e.Message
}

func (e Error) isOk() bool {
	if e.Message == "" {
		return true
	}
	return false
}

type Error struct {
	Message string `json:"error,omitempty"`
}

type Transaction struct {
	InputTXID      string  `json:"inputTXID"`
	InputAddress   string  `json:"inputAddress"`
	InputCurrency  string  `json:"inputCurrency,omitempty"`
	InputAmount    decimal.Decimal `json:"inputAmount,omitempty"`
	OutputTXID     string  `json:"outputTXID,omitempty"`
	OutputAddress  string  `json:"outputAddress,omitempty"`
	OutputCurrency string  `json:"outputCurrency,omitempty"`
	OutputAmount   string  `json:"outputAmount,omitempty"`
	ShiftRate      string  `json:"shiftRate,omitempty"`
	Status         string  `json:"status,omitempty"`
}

type API struct {
	Key     string
	Address string
}

type TimeRemainingResponse struct {
	Status  string `json:"status"`
	Seconds int    `json:"seconds_remaining,string"`
	Error
}

func (p Pair) GetRate() (decimal.Decimal, error) {
	r, err := DoHttp("GET", "rate", p.Name)
	if err != nil {
		return decimal.Zero, err
	}
	var g RateResponse
	err = json.Unmarshal(r, &g)
	return g.Rate, err
}

func (p Pair) GetLimit() (decimal.Decimal, error) {
	r, err := DoHttp("GET", "limit", p.Name)
	if err != nil {
		return decimal.Zero, err
	}
	var g LimitResponse
	err = json.Unmarshal(r, &g)
	return g.Limit, err
}

func (p Pair) GetInfo() (*MarketInfoResponse, error) {
	r, err := DoHttp("GET", "marketinfo", p.Name)
	if err != nil {
		return nil, err
	}
	var g MarketInfoResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func RecentTransactions(count string) (*RecentTranxResponse, error) {
	r, err := DoHttp("GET", "recenttx", count)
	if err != nil {
		return nil, err
	}
	var g RecentTranxResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func DepositStatus(addr string) (*DepositStatusResponse, error) {
	r, err := DoHttp("GET", "txStat", addr)
	if err != nil {
		return nil, err
	}
	var g DepositStatusResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func TimeRemaining(addr string) (*TimeRemainingResponse, error) {
	r, err := DoHttp("GET", "timeremaining", addr)
	if err != nil {
		return nil, err
	}
	var g TimeRemainingResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

type ReceiptResponse struct {
	Email struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"email"`
	Error
}

func Coins() (*CoinsResponse, error) {
	r, err := DoHttp("GET", "getcoins", "")
	if err != nil {
		return nil, err
	}
	var g CoinsResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (r Receipt) Send() (*ReceiptResponse, error) {
	q, err := DoPostHttp("POST", "mail", r)
	if err != nil {
		return nil, err
	}
	var g ReceiptResponse
	json.Unmarshal(q, &g)
	return &g, err
}

func (n New) Shift() (*NewTransactionResponse, error) {
	r, err := DoPostHttp("POST", "shift", n)
	if err != nil {
		return nil, err
	}
	var g NewTransactionResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (n New) FixedShift() (*NewFixedTransactionResponse, error) {
	r, err := DoPostHttp("POST", "sendamount", n)
	if err != nil {
		return nil, err
	}
	var g FixedTransactionResponse
	err = json.Unmarshal(r, &g)
	return &g.Response, err
}

func (n Address) Cancel() (*CancelResponse, error) {
	r, err := DoPostHttp("POST", "cancelpending", n)
	if err != nil {
		return nil, err
	}
	var g CancelResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func Validate(addr string, coin string) (*ValidateResponse, error) {
	r, err := DoHttp("GET", "validateAddress/"+addr, coin)
	if err != nil {
		return nil, err
	}
	var g ValidateResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (i API) ListTransactions() ([]Transaction, error) {
	var r []byte
	//var err error
	var err error
	if i.Address != "" {
		r, err = DoHttp("GET", "txbyaddress/"+i.Address, i.Key)
	} else {
		r, err = DoHttp("GET", "txbyapikey", i.Key)
	}
	if err != nil {
		return nil, err
	}
	var g []Transaction
	err = json.Unmarshal(r, &g)
	return g, err
}

func DoPostHttp(method string, apimethod string, data interface{}) ([]byte, error) {
	new, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod, bytes.NewBuffer(new))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}

func DoHttp(method string, apimethod string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod+"/"+url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}
