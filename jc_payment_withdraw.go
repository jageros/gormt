package gormt

import (
	"time"
)

// JcPaymentWithdraw 用户提币订单表
type JcPaymentWithdraw struct {
	ID             uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'主键ID'" json:"id"`                                         // 主键ID
	InvoiceCode    string    `gorm:"unique;column:invoice_code;type:varchar(64);not null;default:'';comment:'发票编码'" json:"invoice_code"`                                     // 发票编码
	AppID          string    `gorm:"column:app_id;type:varchar(64);not null;default:'';comment:'应用ID'" json:"app_id"`                                                        // 应用ID
	OrderID        string    `gorm:"column:order_id;type:varchar(64);not null;default:'';comment:'游戏端订单ID'" json:"order_id"`                                                 // 游戏端订单ID
	UserID         string    `gorm:"index:user_id;column:user_id;type:varchar(64);not null;default:'';comment:'用户游戏ID'" json:"user_id"`                                      // 用户游戏ID
	UserType       int8      `gorm:"column:user_type;type:tinyint;not null;default:0;comment:'用户类型，0-玩家，1-商家，2-其他'" json:"user_type"`                                        // 用户类型，0-玩家，1-商家，2-其他
	FromAddress    string    `gorm:"column:from_address;type:varchar(64);not null;default:'';comment:'发币账号地址'" json:"from_address"`                                          // 发币账号地址
	ToAddress      string    `gorm:"column:to_address;type:varchar(64);not null;default:'';comment:'用户提币地址'" json:"to_address"`                                              // 用户提币地址
	ChainNet       string    `gorm:"index:chain_net;column:chain_net;type:varchar(16);not null;default:'';comment:'网络类型, options:BTC|ERC20|BEP20|TRC20'" json:"chain_net"`   // 网络类型, options:BTC|ERC20|BEP20|TRC20
	Token          string    `gorm:"index:token;column:token;type:varchar(16);not null;default:'';comment:'代币， options:BTC|USDT|ETH|BNB|TRX|USDP|USDC'" json:"token"`        // 代币， options:BTC|USDT|ETH|BNB|TRX|USDP|USDC
	AmountInput    float64   `gorm:"column:amount_input;type:decimal(40,18);not null;default:0.000000000000000000;comment:'用户输入的提币金额'" json:"amount_input"`                  // 用户输入的提币金额
	AmountCredited float64   `gorm:"column:amount_credited;type:decimal(40,18);not null;default:0.000000000000000000;comment:'实际到账金额'" json:"amount_credited"`               // 实际到账金额
	GasToken       float64   `gorm:"column:gas_token;type:decimal(40,18);not null;default:0.000000000000000000;comment:'手续费换算对应token， 扣掉的， 基于gas_limit计算'" json:"gas_token"` // 手续费换算对应token， 扣掉的， 基于gas_limit计算
	GasRate        float64   `gorm:"column:gas_rate;type:decimal(40,18);not null;default:0.000000000000000000;comment:'燃料费(基础币)兑token的汇率'" json:"gas_rate"`                  // 燃料费(基础币)兑token的汇率
	UserFeeRate    float64   `gorm:"column:user_fee_rate;type:decimal(40,18);not null;default:0.000000000000000000;comment:'扣除玩家手续费的倍率，从配置文件中取'" json:"user_fee_rate"`       // 扣除玩家手续费的倍率，从配置文件中取
	GasLimit       float64   `gorm:"column:gas_limit;type:decimal(40,18);not null;default:0.000000000000000000;comment:'提币的燃料费限制'" json:"gas_limit"`                         // 提币的燃料费限制
	GasPrice       float64   `gorm:"column:gas_price;type:decimal(40,18);not null;default:0.000000000000000000;comment:'燃料费价格，用于计算需要花费的基础币'" json:"gas_price"`               // 燃料费价格，用于计算需要花费的基础币
	GasFee         float64   `gorm:"column:gas_fee;type:decimal(40,18);not null;default:0.000000000000000000;comment:'实际转账燃料费'" json:"gas_fee"`                              // 实际转账燃料费
	SendCount      int       `gorm:"column:send_count;type:int;not null;comment:'广播交易次数'" json:"send_count"`                                                                 // 广播交易次数
	SendAt         time.Time `gorm:"column:send_at;type:timestamp;default:null;comment:'发送交易时间'" json:"send_at"`                                                             // 发送交易时间
	SendRemark     string    `gorm:"column:send_remark;type:varchar(255);not null;default:'';comment:'交易出错备注'" json:"send_remark"`                                           // 交易出错备注
	CheckTimes     int       `gorm:"column:check_times;type:int;not null;default:0;comment:'查询结果次数'" json:"check_times"`                                                     // 查询结果次数
	CheckRemark    string    `gorm:"column:check_remark;type:varchar(255);not null;default:'';comment:'检查出错备注'" json:"check_remark"`                                         // 检查出错备注
	SuccessAt      time.Time `gorm:"column:success_at;type:timestamp;default:null;comment:'成功的时间'" json:"success_at"`                                                        // 成功的时间
	TxHash         string    `gorm:"index:tx_hash;column:tx_hash;type:varchar(80);not null;default:'';comment:'交易哈希'" json:"tx_hash"`                                        // 交易哈希
	Status         int8      `gorm:"column:status;type:tinyint;not null;default:0;comment:'订单状态:0待转账,1转账中,2发送成功待确认,3确认成功,4确认失败,5已退款,6已关闭'" json:"status"`                    // 订单状态:0待转账,1转账中,2发送成功待确认,3确认成功,4确认失败,5已退款,6已关闭
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"created_at"`                                   // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'" json:"updated_at"`                                   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *JcPaymentWithdraw) TableName() string {
	return "jc_payment_withdraw"
}

type IJcPaymentWithdraw interface {
	GetID() uint64
	GetInvoiceCode() string
	GetAppID() string
	GetOrderID() string
	GetUserID() string
	GetUserType() int8
	GetFromAddress() string
	GetToAddress() string
	GetChainNet() string
	GetToken() string
	GetAmountInput() float64
	GetAmountCredited() float64
	GetGasToken() float64
	GetGasRate() float64
	GetUserFeeRate() float64
	GetGasLimit() float64
	GetGasPrice() float64
	GetGasFee() float64
	GetSendCount() int
	GetSendAt() time.Time
	GetSendRemark() string
	GetCheckTimes() int
	GetCheckRemark() string
	GetSuccessAt() time.Time
	GetTxHash() string
	GetStatus() int8
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func (m *JcPaymentWithdraw) GetID() uint64 {
	return m.ID
}

func (m *JcPaymentWithdraw) GetInvoiceCode() string {
	return m.InvoiceCode
}

func (m *JcPaymentWithdraw) GetAppID() string {
	return m.AppID
}

func (m *JcPaymentWithdraw) GetOrderID() string {
	return m.OrderID
}

func (m *JcPaymentWithdraw) GetUserID() string {
	return m.UserID
}

func (m *JcPaymentWithdraw) GetUserType() int8 {
	return m.UserType
}

func (m *JcPaymentWithdraw) GetFromAddress() string {
	return m.FromAddress
}

func (m *JcPaymentWithdraw) GetToAddress() string {
	return m.ToAddress
}

func (m *JcPaymentWithdraw) GetChainNet() string {
	return m.ChainNet
}

func (m *JcPaymentWithdraw) GetToken() string {
	return m.Token
}

func (m *JcPaymentWithdraw) GetAmountInput() float64 {
	return m.AmountInput
}

func (m *JcPaymentWithdraw) GetAmountCredited() float64 {
	return m.AmountCredited
}

func (m *JcPaymentWithdraw) GetGasToken() float64 {
	return m.GasToken
}

func (m *JcPaymentWithdraw) GetGasRate() float64 {
	return m.GasRate
}

func (m *JcPaymentWithdraw) GetUserFeeRate() float64 {
	return m.UserFeeRate
}

func (m *JcPaymentWithdraw) GetGasLimit() float64 {
	return m.GasLimit
}

func (m *JcPaymentWithdraw) GetGasPrice() float64 {
	return m.GasPrice
}

func (m *JcPaymentWithdraw) GetGasFee() float64 {
	return m.GasFee
}

func (m *JcPaymentWithdraw) GetSendCount() int {
	return m.SendCount
}

func (m *JcPaymentWithdraw) GetSendAt() time.Time {
	return m.SendAt
}

func (m *JcPaymentWithdraw) GetSendRemark() string {
	return m.SendRemark
}

func (m *JcPaymentWithdraw) GetCheckTimes() int {
	return m.CheckTimes
}

func (m *JcPaymentWithdraw) GetCheckRemark() string {
	return m.CheckRemark
}

func (m *JcPaymentWithdraw) GetSuccessAt() time.Time {
	return m.SuccessAt
}

func (m *JcPaymentWithdraw) GetTxHash() string {
	return m.TxHash
}

func (m *JcPaymentWithdraw) GetStatus() int8 {
	return m.Status
}

func (m *JcPaymentWithdraw) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *JcPaymentWithdraw) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// ToMap  struct to map 结构体转成map
func (m *JcPaymentWithdraw) ToMap() map[string]any {
	return map[string]any{
		"id":              m.ID,
		"invoice_code":    m.InvoiceCode,
		"app_id":          m.AppID,
		"order_id":        m.OrderID,
		"user_id":         m.UserID,
		"user_type":       m.UserType,
		"from_address":    m.FromAddress,
		"to_address":      m.ToAddress,
		"chain_net":       m.ChainNet,
		"token":           m.Token,
		"amount_input":    m.AmountInput,
		"amount_credited": m.AmountCredited,
		"gas_token":       m.GasToken,
		"gas_rate":        m.GasRate,
		"user_fee_rate":   m.UserFeeRate,
		"gas_limit":       m.GasLimit,
		"gas_price":       m.GasPrice,
		"gas_fee":         m.GasFee,
		"send_count":      m.SendCount,
		"send_at":         m.SendAt,
		"send_remark":     m.SendRemark,
		"check_times":     m.CheckTimes,
		"check_remark":    m.CheckRemark,
		"success_at":      m.SuccessAt,
		"tx_hash":         m.TxHash,
		"status":          m.Status,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
	}
}

// ToMapWithoutModel  struct to map 结构体转成map, 不带gorm.Model
func (m *JcPaymentWithdraw) ToMapWithoutModel() map[string]any {
	return map[string]any{
		"invoice_code":    m.InvoiceCode,
		"app_id":          m.AppID,
		"order_id":        m.OrderID,
		"user_id":         m.UserID,
		"user_type":       m.UserType,
		"from_address":    m.FromAddress,
		"to_address":      m.ToAddress,
		"chain_net":       m.ChainNet,
		"token":           m.Token,
		"amount_input":    m.AmountInput,
		"amount_credited": m.AmountCredited,
		"gas_token":       m.GasToken,
		"gas_rate":        m.GasRate,
		"user_fee_rate":   m.UserFeeRate,
		"gas_limit":       m.GasLimit,
		"gas_price":       m.GasPrice,
		"gas_fee":         m.GasFee,
		"send_count":      m.SendCount,
		"send_at":         m.SendAt,
		"send_remark":     m.SendRemark,
		"check_times":     m.CheckTimes,
		"check_remark":    m.CheckRemark,
		"success_at":      m.SuccessAt,
		"tx_hash":         m.TxHash,
		"status":          m.Status,
	}
}

// JcPaymentWithdrawColumns get sql column name.获取数据库列名
var JcPaymentWithdrawColumns = struct {
	ID             string
	InvoiceCode    string
	AppID          string
	OrderID        string
	UserID         string
	UserType       string
	FromAddress    string
	ToAddress      string
	ChainNet       string
	Token          string
	AmountInput    string
	AmountCredited string
	GasToken       string
	GasRate        string
	UserFeeRate    string
	GasLimit       string
	GasPrice       string
	GasFee         string
	SendCount      string
	SendAt         string
	SendRemark     string
	CheckTimes     string
	CheckRemark    string
	SuccessAt      string
	TxHash         string
	Status         string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	InvoiceCode:    "invoice_code",
	AppID:          "app_id",
	OrderID:        "order_id",
	UserID:         "user_id",
	UserType:       "user_type",
	FromAddress:    "from_address",
	ToAddress:      "to_address",
	ChainNet:       "chain_net",
	Token:          "token",
	AmountInput:    "amount_input",
	AmountCredited: "amount_credited",
	GasToken:       "gas_token",
	GasRate:        "gas_rate",
	UserFeeRate:    "user_fee_rate",
	GasLimit:       "gas_limit",
	GasPrice:       "gas_price",
	GasFee:         "gas_fee",
	SendCount:      "send_count",
	SendAt:         "send_at",
	SendRemark:     "send_remark",
	CheckTimes:     "check_times",
	CheckRemark:    "check_remark",
	SuccessAt:      "success_at",
	TxHash:         "tx_hash",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}
