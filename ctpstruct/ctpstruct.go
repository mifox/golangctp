package ctpstruct

type CThostFtdcDisseminationField struct {
	//信息分发
	SequenceSeries TThostFtdcSequenceSeriesType
	// 序列系列号
	SequenceNo TThostFtdcSequenceNoType
	// 序列号
}
type CThostFtdcReqUserLoginField struct {
	//用户登录请求
	TradingDay TThostFtdcDateType
	// 交易日
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	Password TThostFtdcPasswordType
	// 密码
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// 协议信息
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	OneTimePassword TThostFtdcPasswordType
	// 动态密码
	ClientIPAddress TThostFtdcIPAddressType
	// 终端IP地址
	LoginRemark TThostFtdcLoginRemarkType
	// 登录备注
}
type CThostFtdcRspUserLoginField struct {
	//用户登录应答
	TradingDay TThostFtdcDateType
	// 交易日
	LoginTime TThostFtdcTimeType
	// 登录成功时间
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	SystemName TThostFtdcSystemNameType
	// 交易系统名称
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	MaxOrderRef TThostFtdcOrderRefType
	// 最大报单引用
	SHFETime TThostFtdcTimeType
	// 上期所时间
	DCETime TThostFtdcTimeType
	// 大商所时间
	CZCETime TThostFtdcTimeType
	// 郑商所时间
	FFEXTime TThostFtdcTimeType
	// 中金所时间
	INETime TThostFtdcTimeType
	// 能源中心时间
}
type CThostFtdcUserLogoutField struct {
	//用户登出请求
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcForceUserLogoutField struct {
	//强制交易员退出
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcReqAuthenticateField struct {
	//客户端认证请求
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	AuthCode TThostFtdcAuthCodeType
	// 认证码
}
type CThostFtdcRspAuthenticateField struct {
	//客户端认证响应
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
}
type CThostFtdcAuthenticationInfoField struct {
	//客户端认证信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	AuthInfo TThostFtdcAuthInfoType
	// 认证信息
	IsResult TThostFtdcBoolType
	// 是否为认证结果
}
type CThostFtdcRspUserLogin2Field struct {
	//用户登录应答2
	TradingDay TThostFtdcDateType
	// 交易日
	LoginTime TThostFtdcTimeType
	// 登录成功时间
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	SystemName TThostFtdcSystemNameType
	// 交易系统名称
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	MaxOrderRef TThostFtdcOrderRefType
	// 最大报单引用
	SHFETime TThostFtdcTimeType
	// 上期所时间
	DCETime TThostFtdcTimeType
	// 大商所时间
	CZCETime TThostFtdcTimeType
	// 郑商所时间
	FFEXTime TThostFtdcTimeType
	// 中金所时间
	INETime TThostFtdcTimeType
	// 能源中心时间
	RandomString TThostFtdcRandomStringType
	// 随机串
}
type CThostFtdcTransferHeaderField struct {
	//银期转帐报文头
	Version TThostFtdcVersionType
	// 版本号，常量，1.0
	TradeCode TThostFtdcTradeCodeType
	// 交易代码，必填
	TradeDate TThostFtdcTradeDateType
	// 交易日期，必填，格式：yyyymmdd
	TradeTime TThostFtdcTradeTimeType
	// 交易时间，必填，格式：hhmmss
	TradeSerial TThostFtdcTradeSerialType
	// 发起方流水号，N/A
	FutureID TThostFtdcFutureIDType
	// 期货公司代码，必填
	BankID TThostFtdcBankIDType
	// 银行代码，根据查询银行得到，必填
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码，根据查询银行得到，必填
	OperNo TThostFtdcOperNoType
	// 操作员，N/A
	DeviceID TThostFtdcDeviceIDType
	// 交易设备类型，N/A
	RecordNum TThostFtdcRecordNumType
	// 记录数，N/A
	SessionID TThostFtdcSessionIDType
	// 会话编号，N/A
	RequestID TThostFtdcRequestIDType
	// 请求编号，N/A
}
type CThostFtdcTransferBankToFutureReqField struct {
	//银行资金转期货请求，TradeCode=202001
	FutureAccount TThostFtdcAccountIDType
	// 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 密码
	TradeAmt TThostFtdcMoneyType
	// 转账金额
	CustFee TThostFtdcMoneyType
	// 客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
}
type CThostFtdcTransferBankToFutureRspField struct {
	//银行资金转期货请求响应
	RetCode TThostFtdcRetCodeType
	// 响应代码
	RetInfo TThostFtdcRetInfoType
	// 响应信息
	FutureAccount TThostFtdcAccountIDType
	// 资金账户
	TradeAmt TThostFtdcMoneyType
	// 转帐金额
	CustFee TThostFtdcMoneyType
	// 应收客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种
}
type CThostFtdcTransferFutureToBankReqField struct {
	//期货资金转银行请求，TradeCode=202002
	FutureAccount TThostFtdcAccountIDType
	// 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 密码
	TradeAmt TThostFtdcMoneyType
	// 转账金额
	CustFee TThostFtdcMoneyType
	// 客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
}
type CThostFtdcTransferFutureToBankRspField struct {
	//期货资金转银行请求响应
	RetCode TThostFtdcRetCodeType
	// 响应代码
	RetInfo TThostFtdcRetInfoType
	// 响应信息
	FutureAccount TThostFtdcAccountIDType
	// 资金账户
	TradeAmt TThostFtdcMoneyType
	// 转帐金额
	CustFee TThostFtdcMoneyType
	// 应收客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种
}
type CThostFtdcTransferQryBankReqField struct {
	//查询银行资金请求，TradeCode=204002
	FutureAccount TThostFtdcAccountIDType
	// 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 密码
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
}
type CThostFtdcTransferQryBankRspField struct {
	//查询银行资金请求响应
	RetCode TThostFtdcRetCodeType
	// 响应代码
	RetInfo TThostFtdcRetInfoType
	// 响应信息
	FutureAccount TThostFtdcAccountIDType
	// 资金账户
	TradeAmt TThostFtdcMoneyType
	// 银行余额
	UseAmt TThostFtdcMoneyType
	// 银行可用余额
	FetchAmt TThostFtdcMoneyType
	// 银行可取余额
	CurrencyCode TThostFtdcCurrencyCodeType
	// 币种
}
type CThostFtdcTransferQryDetailReqField struct {
	//查询银行交易明细请求，TradeCode=204999
	FutureAccount TThostFtdcAccountIDType
	// 期货资金账户
}
type CThostFtdcTransferQryDetailRspField struct {
	//查询银行交易明细请求响应
	TradeDate TThostFtdcDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	TradeCode TThostFtdcTradeCodeType
	// 交易代码
	FutureSerial TThostFtdcTradeSerialNoType
	// 期货流水号
	FutureID TThostFtdcFutureIDType
	// 期货公司代码
	FutureAccount TThostFtdcFutureAccountType
	// 资金帐号
	BankSerial TThostFtdcTradeSerialNoType
	// 银行流水号
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码
	BankAccount TThostFtdcBankAccountType
	// 银行账号
	CertCode TThostFtdcCertCodeType
	// 证件号码
	CurrencyCode TThostFtdcCurrencyCodeType
	// 货币代码
	TxAmount TThostFtdcMoneyType
	// 发生金额
	Flag TThostFtdcTransferValidFlagType
	// 有效标志
}
type CThostFtdcRspInfoField struct {
	//响应信息
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcExchangeField struct {
	//交易所
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeName TThostFtdcExchangeNameType
	// 交易所名称
	ExchangeProperty TThostFtdcExchangePropertyType
	// 交易所属性
}
type CThostFtdcProductField struct {
	//产品
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ProductName TThostFtdcProductNameType
	// 产品名称
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ProductClass TThostFtdcProductClassType
	// 产品类型
	VolumeMultiple TThostFtdcVolumeMultipleType
	// 合约数量乘数
	PriceTick TThostFtdcPriceType
	// 最小变动价位
	MaxMarketOrderVolume TThostFtdcVolumeType
	// 市价单最大下单量
	MinMarketOrderVolume TThostFtdcVolumeType
	// 市价单最小下单量
	MaxLimitOrderVolume TThostFtdcVolumeType
	// 限价单最大下单量
	MinLimitOrderVolume TThostFtdcVolumeType
	// 限价单最小下单量
	PositionType TThostFtdcPositionTypeType
	// 持仓类型
	PositionDateType TThostFtdcPositionDateTypeType
	// 持仓日期类型
	CloseDealType TThostFtdcCloseDealTypeType
	// 平仓处理类型
	TradeCurrencyID TThostFtdcCurrencyIDType
	// 交易币种类型
	MortgageFundUseRange TThostFtdcMortgageFundUseRangeType
	// 质押资金可用范围
	ExchangeProductID TThostFtdcInstrumentIDType
	// 交易所产品代码
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType
	// 合约基础商品乘数
}
type CThostFtdcInstrumentField struct {
	//合约
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InstrumentName TThostFtdcInstrumentNameType
	// 合约名称
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ProductClass TThostFtdcProductClassType
	// 产品类型
	DeliveryYear TThostFtdcYearType
	// 交割年份
	DeliveryMonth TThostFtdcMonthType
	// 交割月
	MaxMarketOrderVolume TThostFtdcVolumeType
	// 市价单最大下单量
	MinMarketOrderVolume TThostFtdcVolumeType
	// 市价单最小下单量
	MaxLimitOrderVolume TThostFtdcVolumeType
	// 限价单最大下单量
	MinLimitOrderVolume TThostFtdcVolumeType
	// 限价单最小下单量
	VolumeMultiple TThostFtdcVolumeMultipleType
	// 合约数量乘数
	PriceTick TThostFtdcPriceType
	// 最小变动价位
	CreateDate TThostFtdcDateType
	// 创建日
	OpenDate TThostFtdcDateType
	// 上市日
	ExpireDate TThostFtdcDateType
	// 到期日
	StartDelivDate TThostFtdcDateType
	// 开始交割日
	EndDelivDate TThostFtdcDateType
	// 结束交割日
	InstLifePhase TThostFtdcInstLifePhaseType
	// 合约生命周期状态
	IsTrading TThostFtdcBoolType
	// 当前是否交易
	PositionType TThostFtdcPositionTypeType
	// 持仓类型
	PositionDateType TThostFtdcPositionDateTypeType
	// 持仓日期类型
	LongMarginRatio TThostFtdcRatioType
	// 多头保证金率
	ShortMarginRatio TThostFtdcRatioType
	// 空头保证金率
	MaxMarginSideAlgorithm TThostFtdcMaxMarginSideAlgorithmType
	// 是否使用大额单边保证金算法
	UnderlyingInstrID TThostFtdcInstrumentIDType
	// 基础商品代码
	StrikePrice TThostFtdcPriceType
	// 执行价
	OptionsType TThostFtdcOptionsTypeType
	// 期权类型
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType
	// 合约基础商品乘数
	CombinationType TThostFtdcCombinationTypeType
	// 组合类型
}
type CThostFtdcBrokerField struct {
	//经纪公司
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	BrokerAbbr TThostFtdcBrokerAbbrType
	// 经纪公司简称
	BrokerName TThostFtdcBrokerNameType
	// 经纪公司名称
	IsActive TThostFtdcBoolType
	// 是否活跃
}
type CThostFtdcTraderField struct {
	//交易所交易员
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	Password TThostFtdcPasswordType
	// 密码
	InstallCount TThostFtdcInstallCountType
	// 安装数量
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcInvestorField struct {
	//投资者
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组代码
	InvestorName TThostFtdcPartyNameType
	// 投资者名称
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	IsActive TThostFtdcBoolType
	// 是否活跃
	Telephone TThostFtdcTelephoneType
	// 联系电话
	Address TThostFtdcAddressType
	// 通讯地址
	OpenDate TThostFtdcDateType
	// 开户日期
	Mobile TThostFtdcMobileType
	// 手机
	CommModelID TThostFtdcInvestorIDType
	// 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
}
type CThostFtdcTradingCodeField struct {
	//交易编码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	IsActive TThostFtdcBoolType
	// 是否活跃
	ClientIDType TThostFtdcClientIDTypeType
	// 交易编码类型
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	BizType TThostFtdcBizTypeType
	// 业务类型
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcPartBrokerField struct {
	//会员编码和经纪公司编码对照表
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	IsActive TThostFtdcBoolType
	// 是否活跃
}
type CThostFtdcSuperUserField struct {
	//管理用户
	UserID TThostFtdcUserIDType
	// 用户代码
	UserName TThostFtdcUserNameType
	// 用户名称
	Password TThostFtdcPasswordType
	// 密码
	IsActive TThostFtdcBoolType
	// 是否活跃
}
type CThostFtdcSuperUserFunctionField struct {
	//管理用户功能权限
	UserID TThostFtdcUserIDType
	// 用户代码
	FunctionCode TThostFtdcFunctionCodeType
	// 功能代码
}
type CThostFtdcInvestorGroupField struct {
	//投资者组
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组代码
	InvestorGroupName TThostFtdcInvestorGroupNameType
	// 投资者分组名称
}
type CThostFtdcTradingAccountField struct {
	//资金账户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	PreMortgage TThostFtdcMoneyType
	// 上次质押金额
	PreCredit TThostFtdcMoneyType
	// 上次信用额度
	PreDeposit TThostFtdcMoneyType
	// 上次存款额
	PreBalance TThostFtdcMoneyType
	// 上次结算准备金
	PreMargin TThostFtdcMoneyType
	// 上次占用的保证金
	InterestBase TThostFtdcMoneyType
	// 利息基数
	Interest TThostFtdcMoneyType
	// 利息收入
	Deposit TThostFtdcMoneyType
	// 入金金额
	Withdraw TThostFtdcMoneyType
	// 出金金额
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenCash TThostFtdcMoneyType
	// 冻结的资金
	FrozenCommission TThostFtdcMoneyType
	// 冻结的手续费
	CurrMargin TThostFtdcMoneyType
	// 当前保证金总额
	CashIn TThostFtdcMoneyType
	// 资金差额
	Commission TThostFtdcMoneyType
	// 手续费
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 持仓盈亏
	Balance TThostFtdcMoneyType
	// 期货结算准备金
	Available TThostFtdcMoneyType
	// 可用资金
	WithdrawQuota TThostFtdcMoneyType
	// 可取资金
	Reserve TThostFtdcMoneyType
	// 基本准备金
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	Credit TThostFtdcMoneyType
	// 信用额度
	Mortgage TThostFtdcMoneyType
	// 质押金额
	ExchangeMargin TThostFtdcMoneyType
	// 交易所保证金
	DeliveryMargin TThostFtdcMoneyType
	// 投资者交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType
	// 交易所交割保证金
	ReserveBalance TThostFtdcMoneyType
	// 保底期货结算准备金
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	PreFundMortgageIn TThostFtdcMoneyType
	// 上次货币质入金额
	PreFundMortgageOut TThostFtdcMoneyType
	// 上次货币质出金额
	FundMortgageIn TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageOut TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageAvailable TThostFtdcMoneyType
	// 货币质押余额
	MortgageableFund TThostFtdcMoneyType
	// 可质押货币金额
	SpecProductMargin TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductFrozenMargin TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductCommission TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductFrozenCommission TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductPositionProfit TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin TThostFtdcMoneyType
	// 特殊产品交易所保证金
	BizType TThostFtdcBizTypeType
	// 业务类型
	FrozenSwap TThostFtdcMoneyType
	// 延时换汇冻结金额
	RemainSwap TThostFtdcMoneyType
	// 剩余换汇额度
}
type CThostFtdcInvestorPositionField struct {
	//投资者持仓
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	PosiDirection TThostFtdcPosiDirectionType
	// 持仓多空方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	PositionDate TThostFtdcPositionDateType
	// 持仓日期
	YdPosition TThostFtdcVolumeType
	// 上日持仓
	Position TThostFtdcVolumeType
	// 今日持仓
	LongFrozen TThostFtdcVolumeType
	// 多头冻结
	ShortFrozen TThostFtdcVolumeType
	// 空头冻结
	LongFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	OpenVolume TThostFtdcVolumeType
	// 开仓量
	CloseVolume TThostFtdcVolumeType
	// 平仓量
	OpenAmount TThostFtdcMoneyType
	// 开仓金额
	CloseAmount TThostFtdcMoneyType
	// 平仓金额
	PositionCost TThostFtdcMoneyType
	// 持仓成本
	PreMargin TThostFtdcMoneyType
	// 上次占用的保证金
	UseMargin TThostFtdcMoneyType
	// 占用的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenCash TThostFtdcMoneyType
	// 冻结的资金
	FrozenCommission TThostFtdcMoneyType
	// 冻结的手续费
	CashIn TThostFtdcMoneyType
	// 资金差额
	Commission TThostFtdcMoneyType
	// 手续费
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 持仓盈亏
	PreSettlementPrice TThostFtdcPriceType
	// 上次结算价
	SettlementPrice TThostFtdcPriceType
	// 本次结算价
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OpenCost TThostFtdcMoneyType
	// 开仓成本
	ExchangeMargin TThostFtdcMoneyType
	// 交易所保证金
	CombPosition TThostFtdcVolumeType
	// 组合成交形成的持仓
	CombLongFrozen TThostFtdcVolumeType
	// 组合多头冻结
	CombShortFrozen TThostFtdcVolumeType
	// 组合空头冻结
	CloseProfitByDate TThostFtdcMoneyType
	// 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	TodayPosition TThostFtdcVolumeType
	// 今日持仓
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率
	MarginRateByVolume TThostFtdcRatioType
	// 保证金率(按手数)
	StrikeFrozen TThostFtdcVolumeType
	// 执行冻结
	StrikeFrozenAmount TThostFtdcMoneyType
	// 执行冻结金额
	AbandonFrozen TThostFtdcVolumeType
	// 放弃执行冻结
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	YdStrikeFrozen TThostFtdcVolumeType
	// 执行冻结的昨仓
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcInstrumentMarginRateField struct {
	//合约保证金率
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金费
	IsRelative TThostFtdcBoolType
	// 是否相对交易所收取
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcInstrumentCommissionRateField struct {
	//合约手续费率
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType
	// 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType
	// 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 平今手续费
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	BizType TThostFtdcBizTypeType
	// 业务类型
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcDepthMarketDataField struct {
	//深度行情
	TradingDay TThostFtdcDateType
	// 交易日
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	LastPrice TThostFtdcPriceType
	// 最新价
	PreSettlementPrice TThostFtdcPriceType
	// 上次结算价
	PreClosePrice TThostFtdcPriceType
	// 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType
	// 昨持仓量
	OpenPrice TThostFtdcPriceType
	// 今开盘
	HighestPrice TThostFtdcPriceType
	// 最高价
	LowestPrice TThostFtdcPriceType
	// 最低价
	Volume TThostFtdcVolumeType
	// 数量
	Turnover TThostFtdcMoneyType
	// 成交金额
	OpenInterest TThostFtdcLargeVolumeType
	// 持仓量
	ClosePrice TThostFtdcPriceType
	// 今收盘
	SettlementPrice TThostFtdcPriceType
	// 本次结算价
	UpperLimitPrice TThostFtdcPriceType
	// 涨停板价
	LowerLimitPrice TThostFtdcPriceType
	// 跌停板价
	PreDelta TThostFtdcRatioType
	// 昨虚实度
	CurrDelta TThostFtdcRatioType
	// 今虚实度
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	UpdateMillisec TThostFtdcMillisecType
	// 最后修改毫秒
	BidPrice1 TThostFtdcPriceType
	// 申买价一
	BidVolume1 TThostFtdcVolumeType
	// 申买量一
	AskPrice1 TThostFtdcPriceType
	// 申卖价一
	AskVolume1 TThostFtdcVolumeType
	// 申卖量一
	BidPrice2 TThostFtdcPriceType
	// 申买价二
	BidVolume2 TThostFtdcVolumeType
	// 申买量二
	AskPrice2 TThostFtdcPriceType
	// 申卖价二
	AskVolume2 TThostFtdcVolumeType
	// 申卖量二
	BidPrice3 TThostFtdcPriceType
	// 申买价三
	BidVolume3 TThostFtdcVolumeType
	// 申买量三
	AskPrice3 TThostFtdcPriceType
	// 申卖价三
	AskVolume3 TThostFtdcVolumeType
	// 申卖量三
	BidPrice4 TThostFtdcPriceType
	// 申买价四
	BidVolume4 TThostFtdcVolumeType
	// 申买量四
	AskPrice4 TThostFtdcPriceType
	// 申卖价四
	AskVolume4 TThostFtdcVolumeType
	// 申卖量四
	BidPrice5 TThostFtdcPriceType
	// 申买价五
	BidVolume5 TThostFtdcVolumeType
	// 申买量五
	AskPrice5 TThostFtdcPriceType
	// 申卖价五
	AskVolume5 TThostFtdcVolumeType
	// 申卖量五
	AveragePrice TThostFtdcPriceType
	// 当日均价
	ActionDay TThostFtdcDateType
	// 业务日期
}
type CThostFtdcInstrumentTradingRightField struct {
	//投资者合约交易权限
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	TradingRight TThostFtdcTradingRightType
	// 交易权限
}
type CThostFtdcBrokerUserField struct {
	//经纪公司用户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserName TThostFtdcUserNameType
	// 用户名称
	UserType TThostFtdcUserTypeType
	// 用户类型
	IsActive TThostFtdcBoolType
	// 是否活跃
	IsUsingOTP TThostFtdcBoolType
	// 是否使用令牌
	IsAuthForce TThostFtdcBoolType
	// 是否强制终端认证
}
type CThostFtdcBrokerUserPasswordField struct {
	//经纪公司用户口令
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	Password TThostFtdcPasswordType
	// 密码
	LastUpdateTime TThostFtdcDateTimeType
	// 上次修改时间
	LastLoginTime TThostFtdcDateTimeType
	// 上次登陆时间
	ExpireDate TThostFtdcDateType
	// 密码过期时间
	WeakExpireDate TThostFtdcDateType
	// 弱密码过期时间
}
type CThostFtdcBrokerUserFunctionField struct {
	//经纪公司用户功能权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	BrokerFunctionCode TThostFtdcBrokerFunctionCodeType
	// 经纪公司功能代码
}
type CThostFtdcTraderOfferField struct {
	//交易所交易员报盘机
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	Password TThostFtdcPasswordType
	// 密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	TraderConnectStatus TThostFtdcTraderConnectStatusType
	// 交易所交易员连接状态
	ConnectRequestDate TThostFtdcDateType
	// 发出连接请求的日期
	ConnectRequestTime TThostFtdcTimeType
	// 发出连接请求的时间
	LastReportDate TThostFtdcDateType
	// 上次报告日期
	LastReportTime TThostFtdcTimeType
	// 上次报告时间
	ConnectDate TThostFtdcDateType
	// 完成连接日期
	ConnectTime TThostFtdcTimeType
	// 完成连接时间
	StartDate TThostFtdcDateType
	// 启动日期
	StartTime TThostFtdcTimeType
	// 启动时间
	TradingDay TThostFtdcDateType
	// 交易日
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	MaxTradeID TThostFtdcTradeIDType
	// 本席位最大成交编号
	MaxOrderMessageReference TThostFtdcReturnCodeType
	// 本席位最大报单备拷
}
type CThostFtdcSettlementInfoField struct {
	//投资者结算结果
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	Content TThostFtdcContentType
	// 消息正文
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcInstrumentMarginRateAdjustField struct {
	//合约保证金率调整
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金费
	IsRelative TThostFtdcBoolType
	// 是否相对交易所收取
}
type CThostFtdcExchangeMarginRateField struct {
	//交易所保证金率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金费
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcExchangeMarginRateAdjustField struct {
	//交易所保证金率调整
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 跟随交易所投资者多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 跟随交易所投资者多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 跟随交易所投资者空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 跟随交易所投资者空头保证金费
	ExchLongMarginRatioByMoney TThostFtdcRatioType
	// 交易所多头保证金率
	ExchLongMarginRatioByVolume TThostFtdcMoneyType
	// 交易所多头保证金费
	ExchShortMarginRatioByMoney TThostFtdcRatioType
	// 交易所空头保证金率
	ExchShortMarginRatioByVolume TThostFtdcMoneyType
	// 交易所空头保证金费
	NoLongMarginRatioByMoney TThostFtdcRatioType
	// 不跟随交易所投资者多头保证金率
	NoLongMarginRatioByVolume TThostFtdcMoneyType
	// 不跟随交易所投资者多头保证金费
	NoShortMarginRatioByMoney TThostFtdcRatioType
	// 不跟随交易所投资者空头保证金率
	NoShortMarginRatioByVolume TThostFtdcMoneyType
	// 不跟随交易所投资者空头保证金费
}
type CThostFtdcExchangeRateField struct {
	//汇率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源币种
	FromCurrencyUnit TThostFtdcCurrencyUnitType
	// 源币种单位数量
	ToCurrencyID TThostFtdcCurrencyIDType
	// 目标币种
	ExchangeRate TThostFtdcExchangeRateType
	// 汇率
}
type CThostFtdcSettlementRefField struct {
	//结算引用
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
}
type CThostFtdcCurrentTimeField struct {
	//当前时间
	CurrDate TThostFtdcDateType
	// 当前日期
	CurrTime TThostFtdcTimeType
	// 当前时间
	CurrMillisec TThostFtdcMillisecType
	// 当前时间（毫秒）
	ActionDay TThostFtdcDateType
	// 业务日期
}
type CThostFtdcCommPhaseField struct {
	//通讯阶段
	TradingDay TThostFtdcDateType
	// 交易日
	CommPhaseNo TThostFtdcCommPhaseNoType
	// 通讯时段编号
	SystemID TThostFtdcSystemIDType
	// 系统编号
}
type CThostFtdcLoginInfoField struct {
	//登录信息
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	LoginDate TThostFtdcDateType
	// 登录日期
	LoginTime TThostFtdcTimeType
	// 登录时间
	IPAddress TThostFtdcIPAddressType
	// IP地址
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// 协议信息
	SystemName TThostFtdcSystemNameType
	// 系统名称
	PasswordDeprecated TThostFtdcPasswordType
	// 密码,已弃用
	MaxOrderRef TThostFtdcOrderRefType
	// 最大报单引用
	SHFETime TThostFtdcTimeType
	// 上期所时间
	DCETime TThostFtdcTimeType
	// 大商所时间
	CZCETime TThostFtdcTimeType
	// 郑商所时间
	FFEXTime TThostFtdcTimeType
	// 中金所时间
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	OneTimePassword TThostFtdcPasswordType
	// 动态密码
	INETime TThostFtdcTimeType
	// 能源中心时间
	IsQryControl TThostFtdcBoolType
	// 查询时是否需要流控
	LoginRemark TThostFtdcLoginRemarkType
	// 登录备注
	Password TThostFtdcPasswordType
	// 密码
}
type CThostFtdcLogoutAllField struct {
	//登录信息
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	SystemName TThostFtdcSystemNameType
	// 系统名称
}
type CThostFtdcFrontStatusField struct {
	//前置状态
	FrontID TThostFtdcFrontIDType
	// 前置编号
	LastReportDate TThostFtdcDateType
	// 上次报告日期
	LastReportTime TThostFtdcTimeType
	// 上次报告时间
	IsActive TThostFtdcBoolType
	// 是否活跃
}
type CThostFtdcUserPasswordUpdateField struct {
	//用户口令变更
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	OldPassword TThostFtdcPasswordType
	// 原来的口令
	NewPassword TThostFtdcPasswordType
	// 新的口令
}
type CThostFtdcInputOrderField struct {
	//输入报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	UserForceClose TThostFtdcBoolType
	// 用户强评标志
	IsSwapOrder TThostFtdcBoolType
	// 互换单标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcOrderField struct {
	//报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	OrderSource TThostFtdcOrderSourceType
	// 报单来源
	OrderStatus TThostFtdcOrderStatusType
	// 报单状态
	OrderType TThostFtdcOrderTypeType
	// 报单类型
	VolumeTraded TThostFtdcVolumeType
	// 今成交数量
	VolumeTotal TThostFtdcVolumeType
	// 剩余数量
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 委托时间
	ActiveTime TThostFtdcTimeType
	// 激活时间
	SuspendTime TThostFtdcTimeType
	// 挂起时间
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ActiveTraderID TThostFtdcTraderIDType
	// 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	UserForceClose TThostFtdcBoolType
	// 用户强评标志
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 经纪公司报单编号
	RelativeOrderSysID TThostFtdcOrderSysIDType
	// 相关报单
	ZCETotalTradedVolume TThostFtdcVolumeType
	// 郑商所成交数量
	IsSwapOrder TThostFtdcBoolType
	// 互换单标志
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExchangeOrderField struct {
	//交易所报单
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	OrderSource TThostFtdcOrderSourceType
	// 报单来源
	OrderStatus TThostFtdcOrderStatusType
	// 报单状态
	OrderType TThostFtdcOrderTypeType
	// 报单类型
	VolumeTraded TThostFtdcVolumeType
	// 今成交数量
	VolumeTotal TThostFtdcVolumeType
	// 剩余数量
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 委托时间
	ActiveTime TThostFtdcTimeType
	// 激活时间
	SuspendTime TThostFtdcTimeType
	// 挂起时间
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ActiveTraderID TThostFtdcTraderIDType
	// 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExchangeOrderInsertErrorField struct {
	//交易所报单插入失败
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcInputOrderActionField struct {
	//输入报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeChange TThostFtdcVolumeType
	// 数量变化
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcOrderActionField struct {
	//报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeChange TThostFtdcVolumeType
	// 数量变化
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExchangeOrderActionField struct {
	//交易所报单操作
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeChange TThostFtdcVolumeType
	// 数量变化
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExchangeOrderActionErrorField struct {
	//交易所报单操作失败
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcExchangeTradeField struct {
	//交易所成交
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TradeID TThostFtdcTradeIDType
	// 成交编号
	Direction TThostFtdcDirectionType
	// 买卖方向
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	TradingRole TThostFtdcTradingRoleType
	// 交易角色
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	Price TThostFtdcPriceType
	// 价格
	Volume TThostFtdcVolumeType
	// 数量
	TradeDate TThostFtdcDateType
	// 成交时期
	TradeTime TThostFtdcTimeType
	// 成交时间
	TradeType TThostFtdcTradeTypeType
	// 成交类型
	PriceSource TThostFtdcPriceSourceType
	// 成交价来源
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	TradeSource TThostFtdcTradeSourceType
	// 成交来源
}
type CThostFtdcTradeField struct {
	//成交
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TradeID TThostFtdcTradeIDType
	// 成交编号
	Direction TThostFtdcDirectionType
	// 买卖方向
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	TradingRole TThostFtdcTradingRoleType
	// 交易角色
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	Price TThostFtdcPriceType
	// 价格
	Volume TThostFtdcVolumeType
	// 数量
	TradeDate TThostFtdcDateType
	// 成交时期
	TradeTime TThostFtdcTimeType
	// 成交时间
	TradeType TThostFtdcTradeTypeType
	// 成交类型
	PriceSource TThostFtdcPriceSourceType
	// 成交价来源
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 经纪公司报单编号
	TradeSource TThostFtdcTradeSourceType
	// 成交来源
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcUserSessionField struct {
	//用户会话
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	LoginDate TThostFtdcDateType
	// 登录日期
	LoginTime TThostFtdcTimeType
	// 登录时间
	IPAddress TThostFtdcIPAddressType
	// IP地址
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// 协议信息
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	LoginRemark TThostFtdcLoginRemarkType
	// 登录备注
}
type CThostFtdcQueryMaxOrderVolumeField struct {
	//查询最大报单数量
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	MaxVolume TThostFtdcVolumeType
	// 最大允许报单数量
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcSettlementInfoConfirmField struct {
	//投资者结算结果确认信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ConfirmDate TThostFtdcDateType
	// 确认日期
	ConfirmTime TThostFtdcTimeType
	// 确认时间
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcSyncDepositField struct {
	//出入金同步
	DepositSeqNo TThostFtdcDepositSeqNoType
	// 出入金流水号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Deposit TThostFtdcMoneyType
	// 入金金额
	IsForce TThostFtdcBoolType
	// 是否强制进行
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcSyncFundMortgageField struct {
	//货币质押同步
	MortgageSeqNo TThostFtdcDepositSeqNoType
	// 货币质押流水号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源币种
	MortgageAmount TThostFtdcMoneyType
	// 质押金额
	ToCurrencyID TThostFtdcCurrencyIDType
	// 目标币种
}
type CThostFtdcBrokerSyncField struct {
	//经纪公司同步
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcSyncingInvestorField struct {
	//正在同步中的投资者
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组代码
	InvestorName TThostFtdcPartyNameType
	// 投资者名称
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	IsActive TThostFtdcBoolType
	// 是否活跃
	Telephone TThostFtdcTelephoneType
	// 联系电话
	Address TThostFtdcAddressType
	// 通讯地址
	OpenDate TThostFtdcDateType
	// 开户日期
	Mobile TThostFtdcMobileType
	// 手机
	CommModelID TThostFtdcInvestorIDType
	// 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
}
type CThostFtdcSyncingTradingCodeField struct {
	//正在同步中的交易代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	IsActive TThostFtdcBoolType
	// 是否活跃
	ClientIDType TThostFtdcClientIDTypeType
	// 交易编码类型
}
type CThostFtdcSyncingInvestorGroupField struct {
	//正在同步中的投资者分组
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组代码
	InvestorGroupName TThostFtdcInvestorGroupNameType
	// 投资者分组名称
}
type CThostFtdcSyncingTradingAccountField struct {
	//正在同步中的交易账号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	PreMortgage TThostFtdcMoneyType
	// 上次质押金额
	PreCredit TThostFtdcMoneyType
	// 上次信用额度
	PreDeposit TThostFtdcMoneyType
	// 上次存款额
	PreBalance TThostFtdcMoneyType
	// 上次结算准备金
	PreMargin TThostFtdcMoneyType
	// 上次占用的保证金
	InterestBase TThostFtdcMoneyType
	// 利息基数
	Interest TThostFtdcMoneyType
	// 利息收入
	Deposit TThostFtdcMoneyType
	// 入金金额
	Withdraw TThostFtdcMoneyType
	// 出金金额
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenCash TThostFtdcMoneyType
	// 冻结的资金
	FrozenCommission TThostFtdcMoneyType
	// 冻结的手续费
	CurrMargin TThostFtdcMoneyType
	// 当前保证金总额
	CashIn TThostFtdcMoneyType
	// 资金差额
	Commission TThostFtdcMoneyType
	// 手续费
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 持仓盈亏
	Balance TThostFtdcMoneyType
	// 期货结算准备金
	Available TThostFtdcMoneyType
	// 可用资金
	WithdrawQuota TThostFtdcMoneyType
	// 可取资金
	Reserve TThostFtdcMoneyType
	// 基本准备金
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	Credit TThostFtdcMoneyType
	// 信用额度
	Mortgage TThostFtdcMoneyType
	// 质押金额
	ExchangeMargin TThostFtdcMoneyType
	// 交易所保证金
	DeliveryMargin TThostFtdcMoneyType
	// 投资者交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType
	// 交易所交割保证金
	ReserveBalance TThostFtdcMoneyType
	// 保底期货结算准备金
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	PreFundMortgageIn TThostFtdcMoneyType
	// 上次货币质入金额
	PreFundMortgageOut TThostFtdcMoneyType
	// 上次货币质出金额
	FundMortgageIn TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageOut TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageAvailable TThostFtdcMoneyType
	// 货币质押余额
	MortgageableFund TThostFtdcMoneyType
	// 可质押货币金额
	SpecProductMargin TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductFrozenMargin TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductCommission TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductFrozenCommission TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductPositionProfit TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin TThostFtdcMoneyType
	// 特殊产品交易所保证金
	FrozenSwap TThostFtdcMoneyType
	// 延时换汇冻结金额
	RemainSwap TThostFtdcMoneyType
	// 剩余换汇额度
}
type CThostFtdcSyncingInvestorPositionField struct {
	//正在同步中的投资者持仓
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	PosiDirection TThostFtdcPosiDirectionType
	// 持仓多空方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	PositionDate TThostFtdcPositionDateType
	// 持仓日期
	YdPosition TThostFtdcVolumeType
	// 上日持仓
	Position TThostFtdcVolumeType
	// 今日持仓
	LongFrozen TThostFtdcVolumeType
	// 多头冻结
	ShortFrozen TThostFtdcVolumeType
	// 空头冻结
	LongFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	OpenVolume TThostFtdcVolumeType
	// 开仓量
	CloseVolume TThostFtdcVolumeType
	// 平仓量
	OpenAmount TThostFtdcMoneyType
	// 开仓金额
	CloseAmount TThostFtdcMoneyType
	// 平仓金额
	PositionCost TThostFtdcMoneyType
	// 持仓成本
	PreMargin TThostFtdcMoneyType
	// 上次占用的保证金
	UseMargin TThostFtdcMoneyType
	// 占用的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenCash TThostFtdcMoneyType
	// 冻结的资金
	FrozenCommission TThostFtdcMoneyType
	// 冻结的手续费
	CashIn TThostFtdcMoneyType
	// 资金差额
	Commission TThostFtdcMoneyType
	// 手续费
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 持仓盈亏
	PreSettlementPrice TThostFtdcPriceType
	// 上次结算价
	SettlementPrice TThostFtdcPriceType
	// 本次结算价
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OpenCost TThostFtdcMoneyType
	// 开仓成本
	ExchangeMargin TThostFtdcMoneyType
	// 交易所保证金
	CombPosition TThostFtdcVolumeType
	// 组合成交形成的持仓
	CombLongFrozen TThostFtdcVolumeType
	// 组合多头冻结
	CombShortFrozen TThostFtdcVolumeType
	// 组合空头冻结
	CloseProfitByDate TThostFtdcMoneyType
	// 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	TodayPosition TThostFtdcVolumeType
	// 今日持仓
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率
	MarginRateByVolume TThostFtdcRatioType
	// 保证金率(按手数)
	StrikeFrozen TThostFtdcVolumeType
	// 执行冻结
	StrikeFrozenAmount TThostFtdcMoneyType
	// 执行冻结金额
	AbandonFrozen TThostFtdcVolumeType
	// 放弃执行冻结
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	YdStrikeFrozen TThostFtdcVolumeType
	// 执行冻结的昨仓
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcSyncingInstrumentMarginRateField struct {
	//正在同步中的合约保证金率
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金费
	IsRelative TThostFtdcBoolType
	// 是否相对交易所收取
}
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	//正在同步中的合约手续费率
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType
	// 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType
	// 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 平今手续费
}
type CThostFtdcSyncingInstrumentTradingRightField struct {
	//正在同步中的合约交易权限
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	TradingRight TThostFtdcTradingRightType
	// 交易权限
}
type CThostFtdcQryOrderField struct {
	//查询报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryTradeField struct {
	//查询成交
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TradeID TThostFtdcTradeIDType
	// 成交编号
	TradeTimeStart TThostFtdcTimeType
	// 开始时间
	TradeTimeEnd TThostFtdcTimeType
	// 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInvestorPositionField struct {
	//查询投资者持仓
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryTradingAccountField struct {
	//查询资金账户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	BizType TThostFtdcBizTypeType
	// 业务类型
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
}
type CThostFtdcQryInvestorField struct {
	//查询投资者
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcQryTradingCodeField struct {
	//查询交易编码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ClientIDType TThostFtdcClientIDTypeType
	// 交易编码类型
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInvestorGroupField struct {
	//查询投资者组
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcQryInstrumentMarginRateField struct {
	//查询合约保证金率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInstrumentCommissionRateField struct {
	//查询手续费率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInstrumentTradingRightField struct {
	//查询合约交易权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcQryBrokerField struct {
	//查询经纪公司
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcQryTraderField struct {
	//查询交易员
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQrySuperUserFunctionField struct {
	//查询管理用户功能权限
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcQryUserSessionField struct {
	//查询用户会话
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcQryPartBrokerField struct {
	//查询经纪公司会员代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
}
type CThostFtdcQryFrontStatusField struct {
	//查询前置状态
	FrontID TThostFtdcFrontIDType
	// 前置编号
}
type CThostFtdcQryExchangeOrderField struct {
	//查询交易所报单
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQryOrderActionField struct {
	//查询报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryExchangeOrderActionField struct {
	//查询交易所报单操作
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQrySuperUserField struct {
	//查询管理用户
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcQryExchangeField struct {
	//查询交易所
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryProductField struct {
	//查询产品
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ProductClass TThostFtdcProductClassType
	// 产品类型
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryInstrumentField struct {
	//查询合约
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
}
type CThostFtdcQryDepthMarketDataField struct {
	//查询行情
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryBrokerUserField struct {
	//查询经纪公司用户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcQryBrokerUserFunctionField struct {
	//查询经纪公司用户权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcQryTraderOfferField struct {
	//查询交易员报盘机
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQrySyncDepositField struct {
	//查询出入金流水
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	DepositSeqNo TThostFtdcDepositSeqNoType
	// 出入金流水号
}
type CThostFtdcQrySettlementInfoField struct {
	//查询投资者结算结果
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	TradingDay TThostFtdcDateType
	// 交易日
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcQryExchangeMarginRateField struct {
	//查询交易所保证金率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	//查询交易所调整保证金率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
}
type CThostFtdcQryExchangeRateField struct {
	//查询汇率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源币种
	ToCurrencyID TThostFtdcCurrencyIDType
	// 目标币种
}
type CThostFtdcQrySyncFundMortgageField struct {
	//查询货币质押流水
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	MortgageSeqNo TThostFtdcDepositSeqNoType
	// 货币质押流水号
}
type CThostFtdcQryHisOrderField struct {
	//查询报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
}
type CThostFtdcOptionInstrMiniMarginField struct {
	//当前期权合约最小保证金
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	MinMargin TThostFtdcMoneyType
	// 单位（手）期权合约最小保证金
	ValueMethod TThostFtdcValueMethodType
	// 取值方式
	IsRelative TThostFtdcBoolType
	// 是否跟随交易所收取
}
type CThostFtdcOptionInstrMarginAdjustField struct {
	//当前期权合约保证金调整系数
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	SShortMarginRatioByMoney TThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType
	// 投机空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType
	// 保值空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType
	// 套利空头保证金调整系数
	IsRelative TThostFtdcBoolType
	// 是否跟随交易所收取
	MShortMarginRatioByMoney TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType
	// 做市商空头保证金调整系数
}
type CThostFtdcOptionInstrCommRateField struct {
	//当前期权合约手续费的详细内容
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType
	// 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType
	// 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 平今手续费
	StrikeRatioByMoney TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByVolume TThostFtdcRatioType
	// 执行手续费
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcOptionInstrTradeCostField struct {
	//期权交易成本
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	FixedMargin TThostFtdcMoneyType
	// 期权合约保证金不变部分
	MiniMargin TThostFtdcMoneyType
	// 期权合约最小保证金
	Royalty TThostFtdcMoneyType
	// 期权合约权利金
	ExchFixedMargin TThostFtdcMoneyType
	// 交易所期权合约保证金不变部分
	ExchMiniMargin TThostFtdcMoneyType
	// 交易所期权合约最小保证金
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryOptionInstrTradeCostField struct {
	//期权交易成本查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	InputPrice TThostFtdcPriceType
	// 期权合约报价
	UnderlyingPrice TThostFtdcPriceType
	// 标的价格,填0则用昨结算价
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryOptionInstrCommRateField struct {
	//期权手续费率查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcIndexPriceField struct {
	//股指现货指数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ClosePrice TThostFtdcPriceType
	// 指数现货收盘价
}
type CThostFtdcInputExecOrderField struct {
	//输入的执行宣告
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionType TThostFtdcActionTypeType
	// 执行类型
	PosiDirection TThostFtdcPosiDirectionType
	// 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 期权行权后生成的头寸是否自动平仓
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcInputExecOrderActionField struct {
	//输入执行宣告操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExecOrderField struct {
	//执行宣告
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionType TThostFtdcActionTypeType
	// 执行类型
	PosiDirection TThostFtdcPosiDirectionType
	// 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 期权行权后生成的头寸是否自动平仓
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 本地执行宣告编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 执行宣告提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ExecResult TThostFtdcExecResultType
	// 执行结果
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerExecOrderSeq TThostFtdcSequenceNoType
	// 经纪公司报单编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExecOrderActionField struct {
	//执行宣告操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 本地执行宣告编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	ActionType TThostFtdcActionTypeType
	// 执行类型
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExecOrderField struct {
	//执行宣告查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告编号
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
}
type CThostFtdcExchangeExecOrderField struct {
	//交易所执行宣告信息
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionType TThostFtdcActionTypeType
	// 执行类型
	PosiDirection TThostFtdcPosiDirectionType
	// 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 期权行权后生成的头寸是否自动平仓
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 本地执行宣告编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 执行宣告提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ExecResult TThostFtdcExecResultType
	// 执行结果
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExchangeExecOrderField struct {
	//交易所执行宣告查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQryExecOrderActionField struct {
	//执行宣告操作查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcExchangeExecOrderActionField struct {
	//交易所执行宣告操作
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 本地执行宣告编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	ActionType TThostFtdcActionTypeType
	// 执行类型
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExchangeExecOrderActionField struct {
	//交易所执行宣告操作查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcErrExecOrderField struct {
	//错误执行宣告
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionType TThostFtdcActionTypeType
	// 执行类型
	PosiDirection TThostFtdcPosiDirectionType
	// 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 期权行权后生成的头寸是否自动平仓
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcQryErrExecOrderField struct {
	//查询错误执行宣告
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcErrExecOrderActionField struct {
	//错误执行宣告操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType
	// 执行宣告引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcQryErrExecOrderActionField struct {
	//查询错误执行宣告操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcOptionInstrTradingRightField struct {
	//投资者期权合约交易权限
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	TradingRight TThostFtdcTradingRightType
	// 交易权限
}
type CThostFtdcQryOptionInstrTradingRightField struct {
	//查询期权合约交易权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	Direction TThostFtdcDirectionType
	// 买卖方向
}
type CThostFtdcInputForQuoteField struct {
	//输入的询价
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ForQuoteRef TThostFtdcOrderRefType
	// 询价引用
	UserID TThostFtdcUserIDType
	// 用户代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcForQuoteField struct {
	//询价
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ForQuoteRef TThostFtdcOrderRefType
	// 询价引用
	UserID TThostFtdcUserIDType
	// 用户代码
	ForQuoteLocalID TThostFtdcOrderLocalIDType
	// 本地询价编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	ForQuoteStatus TThostFtdcForQuoteStatusType
	// 询价状态
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerForQutoSeq TThostFtdcSequenceNoType
	// 经纪公司询价编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryForQuoteField struct {
	//询价查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcExchangeForQuoteField struct {
	//交易所询价信息
	ForQuoteLocalID TThostFtdcOrderLocalIDType
	// 本地询价编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	ForQuoteStatus TThostFtdcForQuoteStatusType
	// 询价状态
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExchangeForQuoteField struct {
	//交易所询价查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcInputQuoteField struct {
	//输入的报价
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	QuoteRef TThostFtdcOrderRefType
	// 报价引用
	UserID TThostFtdcUserIDType
	// 用户代码
	AskPrice TThostFtdcPriceType
	// 卖价格
	BidPrice TThostFtdcPriceType
	// 买价格
	AskVolume TThostFtdcVolumeType
	// 卖数量
	BidVolume TThostFtdcVolumeType
	// 买数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	AskOrderRef TThostFtdcOrderRefType
	// 衍生卖报单引用
	BidOrderRef TThostFtdcOrderRefType
	// 衍生买报单引用
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 应价编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcInputQuoteActionField struct {
	//输入报价操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	QuoteActionRef TThostFtdcOrderActionRefType
	// 报价操作引用
	QuoteRef TThostFtdcOrderRefType
	// 报价引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQuoteField struct {
	//报价
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	QuoteRef TThostFtdcOrderRefType
	// 报价引用
	UserID TThostFtdcUserIDType
	// 用户代码
	AskPrice TThostFtdcPriceType
	// 卖价格
	BidPrice TThostFtdcPriceType
	// 买价格
	AskVolume TThostFtdcVolumeType
	// 卖数量
	BidVolume TThostFtdcVolumeType
	// 买数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 本地报价编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	NotifySequence TThostFtdcSequenceNoType
	// 报价提示序号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报价提交状态
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	QuoteStatus TThostFtdcOrderStatusType
	// 报价状态
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	AskOrderSysID TThostFtdcOrderSysIDType
	// 卖方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType
	// 买方报单编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerQuoteSeq TThostFtdcSequenceNoType
	// 经纪公司报价编号
	AskOrderRef TThostFtdcOrderRefType
	// 衍生卖报单引用
	BidOrderRef TThostFtdcOrderRefType
	// 衍生买报单引用
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 应价编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQuoteActionField struct {
	//报价操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	QuoteActionRef TThostFtdcOrderActionRefType
	// 报价操作引用
	QuoteRef TThostFtdcOrderRefType
	// 报价引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 本地报价编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryQuoteField struct {
	//报价查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价编号
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcExchangeQuoteField struct {
	//交易所报价信息
	AskPrice TThostFtdcPriceType
	// 卖价格
	BidPrice TThostFtdcPriceType
	// 买价格
	AskVolume TThostFtdcVolumeType
	// 卖数量
	BidVolume TThostFtdcVolumeType
	// 买数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 本地报价编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	NotifySequence TThostFtdcSequenceNoType
	// 报价提示序号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报价提交状态
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	QuoteStatus TThostFtdcOrderStatusType
	// 报价状态
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	AskOrderSysID TThostFtdcOrderSysIDType
	// 卖方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType
	// 买方报单编号
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 应价编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExchangeQuoteField struct {
	//交易所报价查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQryQuoteActionField struct {
	//报价操作查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcExchangeQuoteActionField struct {
	//交易所报价操作
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType
	// 报价操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 本地报价编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryExchangeQuoteActionField struct {
	//交易所报价操作查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcOptionInstrDeltaField struct {
	//期权合约delta值
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Delta TThostFtdcRatioType
	// Delta值
}
type CThostFtdcForQuoteRspField struct {
	//发给做市商的询价请求
	TradingDay TThostFtdcDateType
	// 交易日
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 询价编号
	ForQuoteTime TThostFtdcTimeType
	// 询价时间
	ActionDay TThostFtdcDateType
	// 业务日期
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcStrikeOffsetField struct {
	//当前期权合约执行偏移值的详细内容
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Offset TThostFtdcMoneyType
	// 执行偏移值
	OffsetType TThostFtdcStrikeOffsetTypeType
	// 执行偏移类型
}
type CThostFtdcQryStrikeOffsetField struct {
	//期权执行偏移值查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcInputBatchOrderActionField struct {
	//输入批量报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	UserID TThostFtdcUserIDType
	// 用户代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcBatchOrderActionField struct {
	//批量报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcExchangeBatchOrderActionField struct {
	//交易所批量报单操作
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryBatchOrderActionField struct {
	//查询批量报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcCombInstrumentGuardField struct {
	//组合合约安全系数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	GuarantRatio TThostFtdcRatioType
	//
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryCombInstrumentGuardField struct {
	//组合合约安全系数查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcInputCombActionField struct {
	//输入的申请组合
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	CombActionRef TThostFtdcOrderRefType
	// 组合引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	Volume TThostFtdcVolumeType
	// 数量
	CombDirection TThostFtdcCombDirectionType
	// 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcCombActionField struct {
	//申请组合
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	CombActionRef TThostFtdcOrderRefType
	// 组合引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	Volume TThostFtdcVolumeType
	// 数量
	CombDirection TThostFtdcCombDirectionType
	// 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionLocalID TThostFtdcOrderLocalIDType
	// 本地申请组合编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ActionStatus TThostFtdcOrderActionStatusType
	// 组合状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	ComTradeID TThostFtdcTradeIDType
	// 组合编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryCombActionField struct {
	//申请组合查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcExchangeCombActionField struct {
	//交易所申请组合信息
	Direction TThostFtdcDirectionType
	// 买卖方向
	Volume TThostFtdcVolumeType
	// 数量
	CombDirection TThostFtdcCombDirectionType
	// 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ActionLocalID TThostFtdcOrderLocalIDType
	// 本地申请组合编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	ActionStatus TThostFtdcOrderActionStatusType
	// 组合状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	ComTradeID TThostFtdcTradeIDType
	// 组合编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
}
type CThostFtdcQryExchangeCombActionField struct {
	//交易所申请组合查询
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcProductExchRateField struct {
	//产品报价汇率
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	QuoteCurrencyID TThostFtdcCurrencyIDType
	// 报价币种类型
	ExchangeRate TThostFtdcExchangeRateType
	// 汇率
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryProductExchRateField struct {
	//产品报价汇率查询
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcQryForQuoteParamField struct {
	//查询询价价差参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcForQuoteParamField struct {
	//询价价差参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	LastPrice TThostFtdcPriceType
	// 最新价
	PriceInterval TThostFtdcPriceType
	// 价差
}
type CThostFtdcMMOptionInstrCommRateField struct {
	//当前做市商期权合约手续费的详细内容
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType
	// 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType
	// 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 平今手续费
	StrikeRatioByMoney TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByVolume TThostFtdcRatioType
	// 执行手续费
}
type CThostFtdcQryMMOptionInstrCommRateField struct {
	//做市商期权手续费率查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcMMInstrumentCommissionRateField struct {
	//做市商合约手续费率
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType
	// 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType
	// 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 平今手续费
}
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	//查询做市商合约手续费率
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcInstrumentOrderCommRateField struct {
	//当前报单手续费的详细内容
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	OrderCommByVolume TThostFtdcRatioType
	// 报单手续费
	OrderActionCommByVolume TThostFtdcRatioType
	// 撤单手续费
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInstrumentOrderCommRateField struct {
	//报单手续费率查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcTradeParamField struct {
	//交易参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	TradeParamID TThostFtdcTradeParamIDType
	// 参数代码
	TradeParamValue TThostFtdcSettlementParamValueType
	// 参数代码值
	Memo TThostFtdcMemoType
	// 备注
}
type CThostFtdcInstrumentMarginRateULField struct {
	//合约保证金率调整
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金费
}
type CThostFtdcFutureLimitPosiParamField struct {
	//期货持仓限制参数
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	SpecOpenVolume TThostFtdcVolumeType
	// 当日投机开仓数量限制
	ArbiOpenVolume TThostFtdcVolumeType
	// 当日套利开仓数量限制
	OpenVolume TThostFtdcVolumeType
	// 当日投机+套利开仓数量限制
}
type CThostFtdcLoginForbiddenIPField struct {
	//禁止登录IP
	IPAddress TThostFtdcIPAddressType
	// IP地址
}
type CThostFtdcIPListField struct {
	//IP列表
	IPAddress TThostFtdcIPAddressType
	// IP地址
	IsWhite TThostFtdcBoolType
	// 是否白名单
}
type CThostFtdcInputOptionSelfCloseField struct {
	//输入的期权自对冲
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 期权自对冲引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 期权行权的头寸是否自对冲
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcInputOptionSelfCloseActionField struct {
	//输入期权自对冲操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType
	// 期权自对冲操作引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 期权自对冲引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcOptionSelfCloseField struct {
	//期权自对冲
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 期权自对冲引用
	UserID TThostFtdcUserIDType
	// 用户代码
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 期权行权的头寸是否自对冲
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 本地期权自对冲编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 期权自对冲提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ExecResult TThostFtdcExecResultType
	// 自对冲结果
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerOptionSelfCloseSeq TThostFtdcSequenceNoType
	// 经纪公司报单编号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcOptionSelfCloseActionField struct {
	//期权自对冲操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType
	// 期权自对冲操作引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 期权自对冲引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 本地期权自对冲编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryOptionSelfCloseField struct {
	//期权自对冲查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲编号
	InsertTimeStart TThostFtdcTimeType
	// 开始时间
	InsertTimeEnd TThostFtdcTimeType
	// 结束时间
}
type CThostFtdcExchangeOptionSelfCloseField struct {
	//交易所期权自对冲信息
	Volume TThostFtdcVolumeType
	// 数量
	RequestID TThostFtdcRequestIDType
	// 请求编号
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 期权行权的头寸是否自对冲
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 本地期权自对冲编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 期权自对冲提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲编号
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 插入时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ExecResult TThostFtdcExecResultType
	// 自对冲结果
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryOptionSelfCloseActionField struct {
	//期权自对冲操作查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcExchangeOptionSelfCloseActionField struct {
	//交易所期权自对冲操作
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 本地期权自对冲编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcSyncDelaySwapField struct {
	//延时换汇同步
	DelaySwapSeqNo TThostFtdcDepositSeqNoType
	// 换汇流水号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源币种
	FromAmount TThostFtdcMoneyType
	// 源金额
	FromFrozenSwap TThostFtdcMoneyType
	// 源换汇冻结金额(可用冻结)
	FromRemainSwap TThostFtdcMoneyType
	// 源剩余换汇额度(可提冻结)
	ToCurrencyID TThostFtdcCurrencyIDType
	// 目标币种
	ToAmount TThostFtdcMoneyType
	// 目标金额
}
type CThostFtdcQrySyncDelaySwapField struct {
	//查询延时换汇同步
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	DelaySwapSeqNo TThostFtdcDepositSeqNoType
	// 延时换汇流水号
}
type CThostFtdcInvestUnitField struct {
	//投资单元
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	InvestorUnitName TThostFtdcPartyNameType
	// 投资者单元名称
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组代码
	CommModelID TThostFtdcInvestorIDType
	// 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcQryInvestUnitField struct {
	//查询投资单元
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcSecAgentCheckModeField struct {
	//二级代理商资金校验模式
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种
	BrokerSecAgentID TThostFtdcAccountIDType
	// 境外中介机构资金帐号
	CheckSelfAccount TThostFtdcBoolType
	// 是否需要校验自己的资金账户
}
type CThostFtdcMarketDataField struct {
	//市场行情
	TradingDay TThostFtdcDateType
	// 交易日
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	LastPrice TThostFtdcPriceType
	// 最新价
	PreSettlementPrice TThostFtdcPriceType
	// 上次结算价
	PreClosePrice TThostFtdcPriceType
	// 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType
	// 昨持仓量
	OpenPrice TThostFtdcPriceType
	// 今开盘
	HighestPrice TThostFtdcPriceType
	// 最高价
	LowestPrice TThostFtdcPriceType
	// 最低价
	Volume TThostFtdcVolumeType
	// 数量
	Turnover TThostFtdcMoneyType
	// 成交金额
	OpenInterest TThostFtdcLargeVolumeType
	// 持仓量
	ClosePrice TThostFtdcPriceType
	// 今收盘
	SettlementPrice TThostFtdcPriceType
	// 本次结算价
	UpperLimitPrice TThostFtdcPriceType
	// 涨停板价
	LowerLimitPrice TThostFtdcPriceType
	// 跌停板价
	PreDelta TThostFtdcRatioType
	// 昨虚实度
	CurrDelta TThostFtdcRatioType
	// 今虚实度
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	UpdateMillisec TThostFtdcMillisecType
	// 最后修改毫秒
	ActionDay TThostFtdcDateType
	// 业务日期
}
type CThostFtdcMarketDataBaseField struct {
	//行情基础属性
	TradingDay TThostFtdcDateType
	// 交易日
	PreSettlementPrice TThostFtdcPriceType
	// 上次结算价
	PreClosePrice TThostFtdcPriceType
	// 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType
	// 昨持仓量
	PreDelta TThostFtdcRatioType
	// 昨虚实度
}
type CThostFtdcMarketDataStaticField struct {
	//行情静态属性
	OpenPrice TThostFtdcPriceType
	// 今开盘
	HighestPrice TThostFtdcPriceType
	// 最高价
	LowestPrice TThostFtdcPriceType
	// 最低价
	ClosePrice TThostFtdcPriceType
	// 今收盘
	UpperLimitPrice TThostFtdcPriceType
	// 涨停板价
	LowerLimitPrice TThostFtdcPriceType
	// 跌停板价
	SettlementPrice TThostFtdcPriceType
	// 本次结算价
	CurrDelta TThostFtdcRatioType
	// 今虚实度
}
type CThostFtdcMarketDataLastMatchField struct {
	//行情最新成交属性
	LastPrice TThostFtdcPriceType
	// 最新价
	Volume TThostFtdcVolumeType
	// 数量
	Turnover TThostFtdcMoneyType
	// 成交金额
	OpenInterest TThostFtdcLargeVolumeType
	// 持仓量
}
type CThostFtdcMarketDataBestPriceField struct {
	//行情最优价属性
	BidPrice1 TThostFtdcPriceType
	// 申买价一
	BidVolume1 TThostFtdcVolumeType
	// 申买量一
	AskPrice1 TThostFtdcPriceType
	// 申卖价一
	AskVolume1 TThostFtdcVolumeType
	// 申卖量一
}
type CThostFtdcMarketDataBid23Field struct {
	//行情申买二、三属性
	BidPrice2 TThostFtdcPriceType
	// 申买价二
	BidVolume2 TThostFtdcVolumeType
	// 申买量二
	BidPrice3 TThostFtdcPriceType
	// 申买价三
	BidVolume3 TThostFtdcVolumeType
	// 申买量三
}
type CThostFtdcMarketDataAsk23Field struct {
	//行情申卖二、三属性
	AskPrice2 TThostFtdcPriceType
	// 申卖价二
	AskVolume2 TThostFtdcVolumeType
	// 申卖量二
	AskPrice3 TThostFtdcPriceType
	// 申卖价三
	AskVolume3 TThostFtdcVolumeType
	// 申卖量三
}
type CThostFtdcMarketDataBid45Field struct {
	//行情申买四、五属性
	BidPrice4 TThostFtdcPriceType
	// 申买价四
	BidVolume4 TThostFtdcVolumeType
	// 申买量四
	BidPrice5 TThostFtdcPriceType
	// 申买价五
	BidVolume5 TThostFtdcVolumeType
	// 申买量五
}
type CThostFtdcMarketDataAsk45Field struct {
	//行情申卖四、五属性
	AskPrice4 TThostFtdcPriceType
	// 申卖价四
	AskVolume4 TThostFtdcVolumeType
	// 申卖量四
	AskPrice5 TThostFtdcPriceType
	// 申卖价五
	AskVolume5 TThostFtdcVolumeType
	// 申卖量五
}
type CThostFtdcMarketDataUpdateTimeField struct {
	//行情更新时间属性
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	UpdateMillisec TThostFtdcMillisecType
	// 最后修改毫秒
	ActionDay TThostFtdcDateType
	// 业务日期
}
type CThostFtdcMarketDataExchangeField struct {
	//行情交易所代码属性
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcSpecificInstrumentField struct {
	//指定的合约
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcInstrumentStatusField struct {
	//合约状态
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	SettlementGroupID TThostFtdcSettlementGroupIDType
	// 结算组代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InstrumentStatus TThostFtdcInstrumentStatusType
	// 合约交易状态
	TradingSegmentSN TThostFtdcTradingSegmentSNType
	// 交易阶段编号
	EnterTime TThostFtdcTimeType
	// 进入本状态时间
	EnterReason TThostFtdcInstStatusEnterReasonType
	// 进入本状态原因
}
type CThostFtdcQryInstrumentStatusField struct {
	//查询合约状态
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
}
type CThostFtdcInvestorAccountField struct {
	//投资者账户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcPositionProfitAlgorithmField struct {
	//浮动盈亏算法
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Algorithm TThostFtdcAlgorithmType
	// 盈亏算法
	Memo TThostFtdcMemoType
	// 备注
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcDiscountField struct {
	//会员资金折扣
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Discount TThostFtdcRatioType
	// 资金折扣比例
}
type CThostFtdcQryTransferBankField struct {
	//查询转帐银行
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码
}
type CThostFtdcTransferBankField struct {
	//转帐银行
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码
	BankName TThostFtdcBankNameType
	// 银行名称
	IsActive TThostFtdcBoolType
	// 是否活跃
}
type CThostFtdcQryInvestorPositionDetailField struct {
	//查询投资者持仓明细
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcInvestorPositionDetailField struct {
	//投资者持仓明细
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	Direction TThostFtdcDirectionType
	// 买卖
	OpenDate TThostFtdcDateType
	// 开仓日期
	TradeID TThostFtdcTradeIDType
	// 成交编号
	Volume TThostFtdcVolumeType
	// 数量
	OpenPrice TThostFtdcPriceType
	// 开仓价
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	TradeType TThostFtdcTradeTypeType
	// 成交类型
	CombInstrumentID TThostFtdcInstrumentIDType
	// 组合合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	CloseProfitByDate TThostFtdcMoneyType
	// 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	PositionProfitByDate TThostFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByTrade TThostFtdcMoneyType
	// 逐笔对冲持仓盈亏
	Margin TThostFtdcMoneyType
	// 投资者保证金
	ExchMargin TThostFtdcMoneyType
	// 交易所保证金
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率
	MarginRateByVolume TThostFtdcRatioType
	// 保证金率(按手数)
	LastSettlementPrice TThostFtdcPriceType
	// 昨结算价
	SettlementPrice TThostFtdcPriceType
	// 结算价
	CloseVolume TThostFtdcVolumeType
	// 平仓量
	CloseAmount TThostFtdcMoneyType
	// 平仓金额
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcTradingAccountPasswordField struct {
	//资金账户口令域
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 密码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcMDTraderOfferField struct {
	//交易所行情报盘机
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	Password TThostFtdcPasswordType
	// 密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	TraderConnectStatus TThostFtdcTraderConnectStatusType
	// 交易所交易员连接状态
	ConnectRequestDate TThostFtdcDateType
	// 发出连接请求的日期
	ConnectRequestTime TThostFtdcTimeType
	// 发出连接请求的时间
	LastReportDate TThostFtdcDateType
	// 上次报告日期
	LastReportTime TThostFtdcTimeType
	// 上次报告时间
	ConnectDate TThostFtdcDateType
	// 完成连接日期
	ConnectTime TThostFtdcTimeType
	// 完成连接时间
	StartDate TThostFtdcDateType
	// 启动日期
	StartTime TThostFtdcTimeType
	// 启动时间
	TradingDay TThostFtdcDateType
	// 交易日
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	MaxTradeID TThostFtdcTradeIDType
	// 本席位最大成交编号
	MaxOrderMessageReference TThostFtdcReturnCodeType
	// 本席位最大报单备拷
}
type CThostFtdcQryMDTraderOfferField struct {
	//查询行情报盘机
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
}
type CThostFtdcQryNoticeField struct {
	//查询客户通知
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcNoticeField struct {
	//客户通知
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	Content TThostFtdcContentType
	// 消息正文
	SequenceLabel TThostFtdcSequenceLabelType
	// 经纪公司通知内容序列号
}
type CThostFtdcUserRightField struct {
	//用户权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserRightType TThostFtdcUserRightTypeType
	// 客户权限类型
	IsForbidden TThostFtdcBoolType
	// 是否禁止
}
type CThostFtdcQrySettlementInfoConfirmField struct {
	//查询结算信息确认域
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcLoadSettlementInfoField struct {
	//装载结算信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcBrokerWithdrawAlgorithmField struct {
	//经纪公司可提资金算法表
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	WithdrawAlgorithm TThostFtdcAlgorithmType
	// 可提资金算法
	UsingRatio TThostFtdcRatioType
	// 资金使用率
	IncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 可提是否包含平仓盈利
	AllWithoutTrade TThostFtdcAllWithoutTradeType
	// 本日无仓且无成交客户是否受可提比例限制
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 可用是否包含平仓盈利
	IsBrokerUserEvent TThostFtdcBoolType
	// 是否启用用户事件
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	FundMortgageRatio TThostFtdcRatioType
	// 货币质押比率
	BalanceAlgorithm TThostFtdcBalanceAlgorithmType
	// 权益算法
}
type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	//资金账户口令变更域
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OldPassword TThostFtdcPasswordType
	// 原来的口令
	NewPassword TThostFtdcPasswordType
	// 新的口令
}
type CThostFtdcTradingAccountPasswordUpdateField struct {
	//资金账户口令变更域
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	OldPassword TThostFtdcPasswordType
	// 原来的口令
	NewPassword TThostFtdcPasswordType
	// 新的口令
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcQryCombinationLegField struct {
	//查询组合合约分腿
	CombInstrumentID TThostFtdcInstrumentIDType
	// 组合合约代码
	LegID TThostFtdcLegIDType
	// 单腿编号
	LegInstrumentID TThostFtdcInstrumentIDType
	// 单腿合约代码
}
type CThostFtdcQrySyncStatusField struct {
	//查询组合合约分腿
	TradingDay TThostFtdcDateType
	// 交易日
}
type CThostFtdcCombinationLegField struct {
	//组合交易合约的单腿
	CombInstrumentID TThostFtdcInstrumentIDType
	// 组合合约代码
	LegID TThostFtdcLegIDType
	// 单腿编号
	LegInstrumentID TThostFtdcInstrumentIDType
	// 单腿合约代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	LegMultiple TThostFtdcLegMultipleType
	// 单腿乘数
	ImplyLevel TThostFtdcImplyLevelType
	// 派生层数
}
type CThostFtdcSyncStatusField struct {
	//数据同步状态
	TradingDay TThostFtdcDateType
	// 交易日
	DataSyncStatus TThostFtdcDataSyncStatusType
	// 数据同步状态
}
type CThostFtdcQryLinkManField struct {
	//查询联系人
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcLinkManField struct {
	//联系人
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	PersonType TThostFtdcPersonTypeType
	// 联系人类型
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	PersonName TThostFtdcPartyNameType
	// 名称
	Telephone TThostFtdcTelephoneType
	// 联系电话
	Address TThostFtdcAddressType
	// 通讯地址
	ZipCode TThostFtdcZipCodeType
	// 邮政编码
	Priority TThostFtdcPriorityType
	// 优先级
	UOAZipCode TThostFtdcUOAZipCodeType
	// 开户邮政编码
	PersonFullName TThostFtdcInvestorFullNameType
	// 全称
}
type CThostFtdcQryBrokerUserEventField struct {
	//查询经纪公司用户事件
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserEventType TThostFtdcUserEventTypeType
	// 用户事件类型
}
type CThostFtdcBrokerUserEventField struct {
	//查询经纪公司用户事件
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	UserEventType TThostFtdcUserEventTypeType
	// 用户事件类型
	EventSequenceNo TThostFtdcSequenceNoType
	// 用户事件序号
	EventDate TThostFtdcDateType
	// 事件发生日期
	EventTime TThostFtdcTimeType
	// 事件发生时间
	UserEventInfo TThostFtdcUserEventInfoType
	// 用户事件信息
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcQryContractBankField struct {
	//查询签约银行请求
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码
}
type CThostFtdcContractBankField struct {
	//查询签约银行响应
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行分中心代码
	BankName TThostFtdcBankNameType
	// 银行名称
}
type CThostFtdcInvestorPositionCombineDetailField struct {
	//投资者组合持仓明细
	TradingDay TThostFtdcDateType
	// 交易日
	OpenDate TThostFtdcDateType
	// 开仓日期
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ComTradeID TThostFtdcTradeIDType
	// 组合编号
	TradeID TThostFtdcTradeIDType
	// 撮合编号
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	Direction TThostFtdcDirectionType
	// 买卖
	TotalAmt TThostFtdcVolumeType
	// 持仓量
	Margin TThostFtdcMoneyType
	// 投资者保证金
	ExchMargin TThostFtdcMoneyType
	// 交易所保证金
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率
	MarginRateByVolume TThostFtdcRatioType
	// 保证金率(按手数)
	LegID TThostFtdcLegIDType
	// 单腿编号
	LegMultiple TThostFtdcLegMultipleType
	// 单腿乘数
	CombInstrumentID TThostFtdcInstrumentIDType
	// 组合持仓合约编码
	TradeGroupID TThostFtdcTradeGroupIDType
	// 成交组号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcParkedOrderField struct {
	//预埋单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	UserForceClose TThostFtdcBoolType
	// 用户强评标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParkedOrderID TThostFtdcParkedOrderIDType
	// 预埋报单编号
	UserType TThostFtdcUserTypeType
	// 用户类型
	Status TThostFtdcParkedOrderStatusType
	// 预埋单状态
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	IsSwapOrder TThostFtdcBoolType
	// 互换单标志
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcParkedOrderActionField struct {
	//输入预埋单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeChange TThostFtdcVolumeType
	// 数量变化
	UserID TThostFtdcUserIDType
	// 用户代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType
	// 预埋撤单单编号
	UserType TThostFtdcUserTypeType
	// 用户类型
	Status TThostFtdcParkedOrderStatusType
	// 预埋撤单状态
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryParkedOrderField struct {
	//查询预埋单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryParkedOrderActionField struct {
	//查询预埋撤单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcRemoveParkedOrderField struct {
	//删除预埋单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ParkedOrderID TThostFtdcParkedOrderIDType
	// 预埋报单编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcRemoveParkedOrderActionField struct {
	//删除预埋撤单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType
	// 预埋撤单编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcInvestorWithdrawAlgorithmField struct {
	//经纪公司可提资金算法表
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	UsingRatio TThostFtdcRatioType
	// 可提资金比例
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	FundMortgageRatio TThostFtdcRatioType
	// 货币质押比率
}
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	//查询组合持仓明细
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	CombInstrumentID TThostFtdcInstrumentIDType
	// 组合持仓合约编码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcMarketDataAveragePriceField struct {
	//成交均价
	AveragePrice TThostFtdcPriceType
	// 当日均价
}
type CThostFtdcVerifyInvestorPasswordField struct {
	//校验投资者密码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Password TThostFtdcPasswordType
	// 密码
}
type CThostFtdcUserIPField struct {
	//用户IP
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	IPMask TThostFtdcIPAddressType
	// IP地址掩码
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcTradingNoticeInfoField struct {
	//用户事件通知信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	SendTime TThostFtdcTimeType
	// 发送时间
	FieldContent TThostFtdcContentType
	// 消息正文
	SequenceSeries TThostFtdcSequenceSeriesType
	// 序列系列号
	SequenceNo TThostFtdcSequenceNoType
	// 序列号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcTradingNoticeField struct {
	//用户事件通知
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者范围
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	SequenceSeries TThostFtdcSequenceSeriesType
	// 序列系列号
	UserID TThostFtdcUserIDType
	// 用户代码
	SendTime TThostFtdcTimeType
	// 发送时间
	SequenceNo TThostFtdcSequenceNoType
	// 序列号
	FieldContent TThostFtdcContentType
	// 消息正文
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryTradingNoticeField struct {
	//查询交易事件通知
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryErrOrderField struct {
	//查询错误报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcErrOrderField struct {
	//错误报单
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	UserForceClose TThostFtdcBoolType
	// 用户强评标志
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	IsSwapOrder TThostFtdcBoolType
	// 互换单标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	ClientID TThostFtdcClientIDType
	// 交易编码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcErrorConditionalOrderField struct {
	//查询错误报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	UserID TThostFtdcUserIDType
	// 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 报单价格条件
	Direction TThostFtdcDirectionType
	// 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 组合投机套保标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeTotalOriginal TThostFtdcVolumeType
	// 数量
	TimeCondition TThostFtdcTimeConditionType
	// 有效期类型
	GTDDate TThostFtdcDateType
	// GTD日期
	VolumeCondition TThostFtdcVolumeConditionType
	// 成交量类型
	MinVolume TThostFtdcVolumeType
	// 最小成交量
	ContingentCondition TThostFtdcContingentConditionType
	// 触发条件
	StopPrice TThostFtdcPriceType
	// 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 强平原因
	IsAutoSuspend TThostFtdcBoolType
	// 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	RequestID TThostFtdcRequestIDType
	// 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约在交易所的代码
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提交状态
	NotifySequence TThostFtdcSequenceNoType
	// 报单提示序号
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	OrderSource TThostFtdcOrderSourceType
	// 报单来源
	OrderStatus TThostFtdcOrderStatusType
	// 报单状态
	OrderType TThostFtdcOrderTypeType
	// 报单类型
	VolumeTraded TThostFtdcVolumeType
	// 今成交数量
	VolumeTotal TThostFtdcVolumeType
	// 剩余数量
	InsertDate TThostFtdcDateType
	// 报单日期
	InsertTime TThostFtdcTimeType
	// 委托时间
	ActiveTime TThostFtdcTimeType
	// 激活时间
	SuspendTime TThostFtdcTimeType
	// 挂起时间
	UpdateTime TThostFtdcTimeType
	// 最后修改时间
	CancelTime TThostFtdcTimeType
	// 撤销时间
	ActiveTraderID TThostFtdcTraderIDType
	// 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType
	// 结算会员编号
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	UserProductInfo TThostFtdcProductInfoType
	// 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	UserForceClose TThostFtdcBoolType
	// 用户强评标志
	ActiveUserID TThostFtdcUserIDType
	// 操作用户代码
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 经纪公司报单编号
	RelativeOrderSysID TThostFtdcOrderSysIDType
	// 相关报单
	ZCETotalTradedVolume TThostFtdcVolumeType
	// 郑商所成交数量
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	IsSwapOrder TThostFtdcBoolType
	// 互换单标志
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	AccountID TThostFtdcAccountIDType
	// 资金账号
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
}
type CThostFtdcQryErrOrderActionField struct {
	//查询错误报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcErrOrderActionField struct {
	//错误报单操作
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单操作引用
	OrderRef TThostFtdcOrderRefType
	// 报单引用
	RequestID TThostFtdcRequestIDType
	// 请求编号
	FrontID TThostFtdcFrontIDType
	// 前置编号
	SessionID TThostFtdcSessionIDType
	// 会话编号
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	OrderSysID TThostFtdcOrderSysIDType
	// 报单编号
	ActionFlag TThostFtdcActionFlagType
	// 操作标志
	LimitPrice TThostFtdcPriceType
	// 价格
	VolumeChange TThostFtdcVolumeType
	// 数量变化
	ActionDate TThostFtdcDateType
	// 操作日期
	ActionTime TThostFtdcTimeType
	// 操作时间
	TraderID TThostFtdcTraderIDType
	// 交易所交易员代码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ClientID TThostFtdcClientIDType
	// 客户代码
	BusinessUnit TThostFtdcBusinessUnitType
	// 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 报单操作状态
	UserID TThostFtdcUserIDType
	// 用户代码
	StatusMsg TThostFtdcErrorMsgType
	// 状态信息
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	BranchID TThostFtdcBranchIDType
	// 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
	MacAddress TThostFtdcMacAddressType
	// Mac地址
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcQryExchangeSequenceField struct {
	//查询交易所状态
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcExchangeSequenceField struct {
	//交易所状态
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	SequenceNo TThostFtdcSequenceNoType
	// 序号
	MarketStatus TThostFtdcInstrumentStatusType
	// 合约交易状态
}
type CThostFtdcQueryMaxOrderVolumeWithPriceField struct {
	//根据价格查询最大报单数量
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	OffsetFlag TThostFtdcOffsetFlagType
	// 开平标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	MaxVolume TThostFtdcVolumeType
	// 最大允许报单数量
	Price TThostFtdcPriceType
	// 报单价格
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryBrokerTradingParamsField struct {
	//查询经纪公司交易参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
}
type CThostFtdcBrokerTradingParamsField struct {
	//经纪公司交易参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	MarginPriceType TThostFtdcMarginPriceTypeType
	// 保证金价格类型
	Algorithm TThostFtdcAlgorithmType
	// 盈亏算法
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 可用是否包含平仓盈利
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	OptionRoyaltyPriceType TThostFtdcOptionRoyaltyPriceTypeType
	// 期权权利金价格类型
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
}
type CThostFtdcQryBrokerTradingAlgosField struct {
	//查询经纪公司交易算法
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
}
type CThostFtdcBrokerTradingAlgosField struct {
	//经纪公司交易算法
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	HandlePositionAlgoID TThostFtdcHandlePositionAlgoIDType
	// 持仓处理算法编号
	FindMarginRateAlgoID TThostFtdcFindMarginRateAlgoIDType
	// 寻找保证金率算法编号
	HandleTradingAccountAlgoID TThostFtdcHandleTradingAccountAlgoIDType
	// 资金处理算法编号
}
type CThostFtdcQueryBrokerDepositField struct {
	//查询经纪公司资金
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcBrokerDepositField struct {
	//经纪公司资金
	TradingDay TThostFtdcTradeDateType
	// 交易日期
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType
	// 会员代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	PreBalance TThostFtdcMoneyType
	// 上次结算准备金
	CurrMargin TThostFtdcMoneyType
	// 当前保证金总额
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	Balance TThostFtdcMoneyType
	// 期货结算准备金
	Deposit TThostFtdcMoneyType
	// 入金金额
	Withdraw TThostFtdcMoneyType
	// 出金金额
	Available TThostFtdcMoneyType
	// 可提资金
	Reserve TThostFtdcMoneyType
	// 基本准备金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
}
type CThostFtdcQryCFMMCBrokerKeyField struct {
	//查询保证金监管系统经纪公司密钥
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
}
type CThostFtdcCFMMCBrokerKeyField struct {
	//保证金监管系统经纪公司密钥
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType
	// 经纪公司统一编码
	CreateDate TThostFtdcDateType
	// 密钥生成日期
	CreateTime TThostFtdcTimeType
	// 密钥生成时间
	KeyID TThostFtdcSequenceNoType
	// 密钥编号
	CurrentKey TThostFtdcCFMMCKeyType
	// 动态密钥
	KeyKind TThostFtdcCFMMCKeyKindType
	// 动态密钥类型
}
type CThostFtdcCFMMCTradingAccountKeyField struct {
	//保证金监管系统经纪公司资金账户密钥
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType
	// 经纪公司统一编码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	KeyID TThostFtdcSequenceNoType
	// 密钥编号
	CurrentKey TThostFtdcCFMMCKeyType
	// 动态密钥
}
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	//请求查询保证金监管系统经纪公司资金账户密钥
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
type CThostFtdcBrokerUserOTPParamField struct {
	//用户动态令牌参数
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	OTPVendorsID TThostFtdcOTPVendorsIDType
	// 动态令牌提供商
	SerialNumber TThostFtdcSerialNumberType
	// 动态令牌序列号
	AuthKey TThostFtdcAuthKeyType
	// 令牌密钥
	LastDrift TThostFtdcLastDriftType
	// 漂移值
	LastSuccess TThostFtdcLastSuccessType
	// 成功值
	OTPType TThostFtdcOTPTypeType
	// 动态令牌类型
}
type CThostFtdcManualSyncBrokerUserOTPField struct {
	//手工同步用户动态令牌
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	OTPType TThostFtdcOTPTypeType
	// 动态令牌类型
	FirstOTP TThostFtdcPasswordType
	// 第一个动态密码
	SecondOTP TThostFtdcPasswordType
	// 第二个动态密码
}
type CThostFtdcCommRateModelField struct {
	//投资者手续费率模板
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	CommModelID TThostFtdcInvestorIDType
	// 手续费率模板代码
	CommModelName TThostFtdcCommModelNameType
	// 模板名称
}
type CThostFtdcQryCommRateModelField struct {
	//请求查询投资者手续费率模板
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	CommModelID TThostFtdcInvestorIDType
	// 手续费率模板代码
}
type CThostFtdcMarginModelField struct {
	//投资者保证金率模板
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	MarginModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelName TThostFtdcCommModelNameType
	// 模板名称
}
type CThostFtdcQryMarginModelField struct {
	//请求查询投资者保证金率模板
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	MarginModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
}
type CThostFtdcEWarrantOffsetField struct {
	//仓单折抵信息
	TradingDay TThostFtdcTradeDateType
	// 交易日期
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	Direction TThostFtdcDirectionType
	// 买卖方向
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	Volume TThostFtdcVolumeType
	// 数量
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryEWarrantOffsetField struct {
	//查询仓单折抵信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQryInvestorProductGroupMarginField struct {
	//查询投资者品种/跨品种保证金
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	ProductGroupID TThostFtdcInstrumentIDType
	// 品种/跨品种标示
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcInvestorProductGroupMarginField struct {
	//投资者品种/跨品种保证金
	ProductGroupID TThostFtdcInstrumentIDType
	// 品种/跨品种标示
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	TradingDay TThostFtdcDateType
	// 交易日
	SettlementID TThostFtdcSettlementIDType
	// 结算编号
	FrozenMargin TThostFtdcMoneyType
	// 冻结的保证金
	LongFrozenMargin TThostFtdcMoneyType
	// 多头冻结的保证金
	ShortFrozenMargin TThostFtdcMoneyType
	// 空头冻结的保证金
	UseMargin TThostFtdcMoneyType
	// 占用的保证金
	LongUseMargin TThostFtdcMoneyType
	// 多头保证金
	ShortUseMargin TThostFtdcMoneyType
	// 空头保证金
	ExchMargin TThostFtdcMoneyType
	// 交易所保证金
	LongExchMargin TThostFtdcMoneyType
	// 交易所多头保证金
	ShortExchMargin TThostFtdcMoneyType
	// 交易所空头保证金
	CloseProfit TThostFtdcMoneyType
	// 平仓盈亏
	FrozenCommission TThostFtdcMoneyType
	// 冻结的手续费
	Commission TThostFtdcMoneyType
	// 手续费
	FrozenCash TThostFtdcMoneyType
	// 冻结的资金
	CashIn TThostFtdcMoneyType
	// 资金差额
	PositionProfit TThostFtdcMoneyType
	// 持仓盈亏
	OffsetAmount TThostFtdcMoneyType
	// 折抵总金额
	LongOffsetAmount TThostFtdcMoneyType
	// 多头折抵总金额
	ShortOffsetAmount TThostFtdcMoneyType
	// 空头折抵总金额
	ExchOffsetAmount TThostFtdcMoneyType
	// 交易所折抵总金额
	LongExchOffsetAmount TThostFtdcMoneyType
	// 交易所多头折抵总金额
	ShortExchOffsetAmount TThostFtdcMoneyType
	// 交易所空头折抵总金额
	HedgeFlag TThostFtdcHedgeFlagType
	// 投机套保标志
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	//查询监控中心用户令牌
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资单元代码
}
type CThostFtdcCFMMCTradingAccountTokenField struct {
	//监控中心用户令牌
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType
	// 经纪公司统一编码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	KeyID TThostFtdcSequenceNoType
	// 密钥编号
	Token TThostFtdcCFMMCTokenType
	// 动态令牌
}
type CThostFtdcQryProductGroupField struct {
	//查询产品组
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
}
type CThostFtdcProductGroupField struct {
	//投资者品种/跨品种保证金产品组
	ProductID TThostFtdcInstrumentIDType
	// 产品代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	ProductGroupID TThostFtdcInstrumentIDType
	// 产品组代码
}
type CThostFtdcBulletinField struct {
	//交易所公告
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	TradingDay TThostFtdcDateType
	// 交易日
	BulletinID TThostFtdcBulletinIDType
	// 公告编号
	SequenceNo TThostFtdcSequenceNoType
	// 序列号
	NewsType TThostFtdcNewsTypeType
	// 公告类型
	NewsUrgency TThostFtdcNewsUrgencyType
	// 紧急程度
	SendTime TThostFtdcTimeType
	// 发送时间
	Abstract TThostFtdcAbstractType
	// 消息摘要
	ComeFrom TThostFtdcComeFromType
	// 消息来源
	Content TThostFtdcContentType
	// 消息正文
	URLLink TThostFtdcURLLinkType
	// WEB地址
	MarketID TThostFtdcMarketIDType
	// 市场代码
}
type CThostFtdcQryBulletinField struct {
	//查询交易所公告
	ExchangeID TThostFtdcExchangeIDType
	// 交易所代码
	BulletinID TThostFtdcBulletinIDType
	// 公告编号
	SequenceNo TThostFtdcSequenceNoType
	// 序列号
	NewsType TThostFtdcNewsTypeType
	// 公告类型
	NewsUrgency TThostFtdcNewsUrgencyType
	// 紧急程度
}
type CThostFtdcReqOpenAccountField struct {
	//转帐开户请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 汇钞标志
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	TID TThostFtdcTIDType
	// 交易ID
	UserID TThostFtdcUserIDType
	// 用户标识
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcReqCancelAccountField struct {
	//转帐销户请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 汇钞标志
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	TID TThostFtdcTIDType
	// 交易ID
	UserID TThostFtdcUserIDType
	// 用户标识
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcReqChangeAccountField struct {
	//变更银行账户请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	NewBankAccount TThostFtdcBankAccountType
	// 新银行帐号
	NewBankPassWord TThostFtdcPasswordType
	// 新银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	TID TThostFtdcTIDType
	// 交易ID
	Digest TThostFtdcDigestType
	// 摘要
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcReqTransferField struct {
	//转账请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType
	// 费用支付标志
	CustFee TThostFtdcCustFeeType
	// 应收客户费用
	BrokerFee TThostFtdcFutureFeeType
	// 应收期货公司费用
	Message TThostFtdcAddInfoType
	// 发送方给接收方的消息
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	TransferStatus TThostFtdcTransferStatusType
	// 转账交易状态
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcRspTransferField struct {
	//银行发起银行资金转期货响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType
	// 费用支付标志
	CustFee TThostFtdcCustFeeType
	// 应收客户费用
	BrokerFee TThostFtdcFutureFeeType
	// 应收期货公司费用
	Message TThostFtdcAddInfoType
	// 发送方给接收方的消息
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	TransferStatus TThostFtdcTransferStatusType
	// 转账交易状态
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcReqRepealField struct {
	//冲正请求
	RepealTimeInterval TThostFtdcRepealTimeIntervalType
	// 冲正时间间隔
	RepealedTimes TThostFtdcRepealedTimesType
	// 已经冲正次数
	BankRepealFlag TThostFtdcBankRepealFlagType
	// 银行冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType
	// 期商冲正标志
	PlateRepealSerial TThostFtdcPlateSerialType
	// 被冲正平台流水号
	BankRepealSerial TThostFtdcBankSerialType
	// 被冲正银行流水号
	FutureRepealSerial TThostFtdcFutureSerialType
	// 被冲正期货流水号
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType
	// 费用支付标志
	CustFee TThostFtdcCustFeeType
	// 应收客户费用
	BrokerFee TThostFtdcFutureFeeType
	// 应收期货公司费用
	Message TThostFtdcAddInfoType
	// 发送方给接收方的消息
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	TransferStatus TThostFtdcTransferStatusType
	// 转账交易状态
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcRspRepealField struct {
	//冲正响应
	RepealTimeInterval TThostFtdcRepealTimeIntervalType
	// 冲正时间间隔
	RepealedTimes TThostFtdcRepealedTimesType
	// 已经冲正次数
	BankRepealFlag TThostFtdcBankRepealFlagType
	// 银行冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType
	// 期商冲正标志
	PlateRepealSerial TThostFtdcPlateSerialType
	// 被冲正平台流水号
	BankRepealSerial TThostFtdcBankSerialType
	// 被冲正银行流水号
	FutureRepealSerial TThostFtdcFutureSerialType
	// 被冲正期货流水号
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType
	// 费用支付标志
	CustFee TThostFtdcCustFeeType
	// 应收客户费用
	BrokerFee TThostFtdcFutureFeeType
	// 应收期货公司费用
	Message TThostFtdcAddInfoType
	// 发送方给接收方的消息
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	TransferStatus TThostFtdcTransferStatusType
	// 转账交易状态
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcReqQueryAccountField struct {
	//查询账户信息请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcRspQueryAccountField struct {
	//查询账户信息响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	BankUseAmount TThostFtdcTradeAmountType
	// 银行可用金额
	BankFetchAmount TThostFtdcTradeAmountType
	// 银行可取金额
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcFutureSignIOField struct {
	//期商签到签退
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
}
type CThostFtdcRspFutureSignInField struct {
	//期商签到响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	PinKey TThostFtdcPasswordKeyType
	// PIN密钥
	MacKey TThostFtdcPasswordKeyType
	// MAC密钥
}
type CThostFtdcReqFutureSignOutField struct {
	//期商签退请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
}
type CThostFtdcRspFutureSignOutField struct {
	//期商签退响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcReqQueryTradeResultBySerialField struct {
	//查询指定流水号的交易结果请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	Reference TThostFtdcSerialType
	// 流水号
	RefrenceIssureType TThostFtdcInstitutionTypeType
	// 本流水号发布者的机构类型
	RefrenceIssure TThostFtdcOrganCodeType
	// 本流水号发布者机构编码
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	Digest TThostFtdcDigestType
	// 摘要
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcRspQueryTradeResultBySerialField struct {
	//查询指定流水号的交易结果响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	Reference TThostFtdcSerialType
	// 流水号
	RefrenceIssureType TThostFtdcInstitutionTypeType
	// 本流水号发布者的机构类型
	RefrenceIssure TThostFtdcOrganCodeType
	// 本流水号发布者机构编码
	OriginReturnCode TThostFtdcReturnCodeType
	// 原始返回代码
	OriginDescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
	// 原始返回码描述
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 转帐金额
	Digest TThostFtdcDigestType
	// 摘要
}
type CThostFtdcReqDayEndFileReadyField struct {
	//日终文件就绪请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	FileBusinessCode TThostFtdcFileBusinessCodeType
	// 文件业务功能
	Digest TThostFtdcDigestType
	// 摘要
}
type CThostFtdcReturnResultField struct {
	//返回结果
	ReturnCode TThostFtdcReturnCodeType
	// 返回代码
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
	// 返回码描述
}
type CThostFtdcVerifyFuturePasswordField struct {
	//验证期货资金密码
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	TID TThostFtdcTIDType
	// 交易ID
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcVerifyCustInfoField struct {
	//验证客户信息
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	//验证期货资金密码和客户信息
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcDepositResultInformField struct {
	//验证期货资金密码和客户信息
	DepositSeqNo TThostFtdcDepositSeqNoType
	// 出入金流水号，该流水号为银期报盘返回的流水号
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	Deposit TThostFtdcMoneyType
	// 入金金额
	RequestID TThostFtdcRequestIDType
	// 请求编号
	ReturnCode TThostFtdcReturnCodeType
	// 返回代码
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
	// 返回码描述
}
type CThostFtdcReqSyncKeyField struct {
	//交易核心向银期报盘发出密钥同步请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Message TThostFtdcAddInfoType
	// 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
}
type CThostFtdcRspSyncKeyField struct {
	//交易核心向银期报盘发出密钥同步响应
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Message TThostFtdcAddInfoType
	// 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcNotifyQueryAccountField struct {
	//查询账户信息通知
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	BankUseAmount TThostFtdcTradeAmountType
	// 银行可用金额
	BankFetchAmount TThostFtdcTradeAmountType
	// 银行可取金额
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcTransferSerialField struct {
	//银期转账交易流水表
	PlateSerial TThostFtdcPlateSerialType
	// 平台流水号
	TradeDate TThostFtdcTradeDateType
	// 交易发起方日期
	TradingDay TThostFtdcDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	TradeCode TThostFtdcTradeCodeType
	// 交易代码
	SessionID TThostFtdcSessionIDType
	// 会话编号
	BankID TThostFtdcBankIDType
	// 银行编码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构编码
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	BrokerID TThostFtdcBrokerIDType
	// 期货公司编码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	FutureAccType TThostFtdcFutureAccTypeType
	// 期货公司帐号类型
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
	FutureSerial TThostFtdcFutureSerialType
	// 期货公司流水号
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	TradeAmount TThostFtdcTradeAmountType
	// 交易金额
	CustFee TThostFtdcCustFeeType
	// 应收客户费用
	BrokerFee TThostFtdcFutureFeeType
	// 应收期货公司费用
	AvailabilityFlag TThostFtdcAvailabilityFlagType
	// 有效标志
	OperatorCode TThostFtdcOperatorCodeType
	// 操作员
	BankNewAccount TThostFtdcBankAccountType
	// 新银行帐号
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcQryTransferSerialField struct {
	//请求查询转帐流水
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	BankID TThostFtdcBankIDType
	// 银行编码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcNotifyFutureSignInField struct {
	//期商签到通知
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	PinKey TThostFtdcPasswordKeyType
	// PIN密钥
	MacKey TThostFtdcPasswordKeyType
	// MAC密钥
}
type CThostFtdcNotifyFutureSignOutField struct {
	//期商签退通知
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Digest TThostFtdcDigestType
	// 摘要
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcNotifySyncKeyField struct {
	//交易核心向银期报盘发出密钥同步处理结果的通知
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	InstallID TThostFtdcInstallIDType
	// 安装编号
	UserID TThostFtdcUserIDType
	// 用户标识
	Message TThostFtdcAddInfoType
	// 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	OperNo TThostFtdcOperNoType
	// 交易柜员
	RequestID TThostFtdcRequestIDType
	// 请求编号
	TID TThostFtdcTIDType
	// 交易ID
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcQryAccountregisterField struct {
	//请求查询银期签约关系
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	BankID TThostFtdcBankIDType
	// 银行编码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构编码
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcAccountregisterField struct {
	//客户开销户信息表
	TradeDay TThostFtdcTradeDateType
	// 交易日期
	BankID TThostFtdcBankIDType
	// 银行编码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构编码
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BrokerID TThostFtdcBrokerIDType
	// 期货公司编码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期货公司分支机构编码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	OpenOrDestroy TThostFtdcOpenOrDestroyType
	// 开销户类别
	RegDate TThostFtdcTradeDateType
	// 签约日期
	OutDate TThostFtdcTradeDateType
	// 解约日期
	TID TThostFtdcTIDType
	// 交易ID
	CustType TThostFtdcCustTypeType
	// 客户类型
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcOpenAccountField struct {
	//银期开户信息
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 汇钞标志
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	TID TThostFtdcTIDType
	// 交易ID
	UserID TThostFtdcUserIDType
	// 用户标识
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcCancelAccountField struct {
	//银期销户信息
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 汇钞标志
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	DeviceID TThostFtdcDeviceIDType
	// 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType
	// 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	OperNo TThostFtdcOperNoType
	// 交易柜员
	TID TThostFtdcTIDType
	// 交易ID
	UserID TThostFtdcUserIDType
	// 用户标识
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcChangeAccountField struct {
	//银期变更银行账号信息
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	NewBankAccount TThostFtdcBankAccountType
	// 新银行帐号
	NewBankPassWord TThostFtdcPasswordType
	// 新银行密码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	BankPwdFlag TThostFtdcPwdFlagType
	// 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	TID TThostFtdcTIDType
	// 交易ID
	Digest TThostFtdcDigestType
	// 摘要
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType
	// 长客户姓名
}
type CThostFtdcSecAgentACIDMapField struct {
	//二级代理操作员银期权限
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	AccountID TThostFtdcAccountIDType
	// 资金账户
	CurrencyID TThostFtdcCurrencyIDType
	// 币种
	BrokerSecAgentID TThostFtdcAccountIDType
	// 境外中介机构资金帐号
}
type CThostFtdcQrySecAgentACIDMapField struct {
	//二级代理操作员银期权限查询
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	AccountID TThostFtdcAccountIDType
	// 资金账户
	CurrencyID TThostFtdcCurrencyIDType
	// 币种
}
type CThostFtdcUserRightsAssignField struct {
	//灾备中心交易权限
	BrokerID TThostFtdcBrokerIDType
	// 应用单元代码
	UserID TThostFtdcUserIDType
	// 用户代码
	DRIdentityID TThostFtdcDRIdentityIDType
	// 交易中心代码
}
type CThostFtdcBrokerUserRightAssignField struct {
	//经济公司是否有在本标示的交易权限
	BrokerID TThostFtdcBrokerIDType
	// 应用单元代码
	DRIdentityID TThostFtdcDRIdentityIDType
	// 交易中心代码
	Tradeable TThostFtdcBoolType
	// 能否交易
}
type CThostFtdcDRTransferField struct {
	//灾备交易转换报文
	OrigDRIdentityID TThostFtdcDRIdentityIDType
	// 原交易中心代码
	DestDRIdentityID TThostFtdcDRIdentityIDType
	// 目标交易中心代码
	OrigBrokerID TThostFtdcBrokerIDType
	// 原应用单元代码
	DestBrokerID TThostFtdcBrokerIDType
	// 目标易用单元代码
}
type CThostFtdcFensUserInfoField struct {
	//Fens用户信息
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	LoginMode TThostFtdcLoginModeType
	// 登录模式
}
type CThostFtdcCurrTransferIdentityField struct {
	//当前银期所属交易中心
	IdentityID TThostFtdcDRIdentityIDType
	// 交易中心代码
}
type CThostFtdcLoginForbiddenUserField struct {
	//禁止登录用户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
	IPAddress TThostFtdcIPAddressType
	// IP地址
}
type CThostFtdcQryLoginForbiddenUserField struct {
	//查询禁止登录用户
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcMulticastGroupInfoField struct {
	//UDP组播组信息
	GroupIP TThostFtdcIPAddressType
	// 组播组IP地址
	GroupPort TThostFtdcIPPortType
	// 组播组IP端口
	SourceIP TThostFtdcIPAddressType
	// 源地址
}
type CThostFtdcTradingAccountReserveField struct {
	//资金账户基本准备金
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Reserve TThostFtdcMoneyType
	// 基本准备金
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcQryLoginForbiddenIPField struct {
	//查询禁止登录IP
	IPAddress TThostFtdcIPAddressType
	// IP地址
}
type CThostFtdcQryIPListField struct {
	//查询IP列表
	IPAddress TThostFtdcIPAddressType
	// IP地址
}
type CThostFtdcQryUserRightsAssignField struct {
	//查询用户下单权限分配表
	BrokerID TThostFtdcBrokerIDType
	// 应用单元代码
	UserID TThostFtdcUserIDType
	// 用户代码
}
type CThostFtdcReserveOpenAccountConfirmField struct {
	//银期预约开户确认请求
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcLongIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	TID TThostFtdcTIDType
	// 交易ID
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	Password TThostFtdcPasswordType
	// 期货密码
	BankReserveOpenSeq TThostFtdcBankSerialType
	// 预约开户银行流水号
	BookDate TThostFtdcTradeDateType
	// 预约开户日期
	BookPsw TThostFtdcPasswordType
	// 预约开户验证密码
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcReserveOpenAccountField struct {
	//银期预约开户
	TradeCode TThostFtdcTradeCodeType
	// 业务功能码
	BankID TThostFtdcBankIDType
	// 银行代码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType
	// 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期商分支机构代码
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradeTime TThostFtdcTradeTimeType
	// 交易时间
	BankSerial TThostFtdcBankSerialType
	// 银行流水号
	TradingDay TThostFtdcTradeDateType
	// 交易系统日期
	PlateSerial TThostFtdcSerialType
	// 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType
	// 最后分片标志
	SessionID TThostFtdcSessionIDType
	// 会话号
	CustomerName TThostFtdcLongIndividualNameType
	// 客户姓名
	IdCardType TThostFtdcIdCardTypeType
	// 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 证件号码
	Gender TThostFtdcGenderType
	// 性别
	CountryCode TThostFtdcCountryCodeType
	// 国家代码
	CustType TThostFtdcCustTypeType
	// 客户类型
	Address TThostFtdcAddressType
	// 地址
	ZipCode TThostFtdcZipCodeType
	// 邮编
	Telephone TThostFtdcTelephoneType
	// 电话号码
	MobilePhone TThostFtdcMobilePhoneType
	// 手机
	Fax TThostFtdcFaxType
	// 传真
	EMail TThostFtdcEMailType
	// 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 资金账户状态
	BankAccount TThostFtdcBankAccountType
	// 银行帐号
	BankPassWord TThostFtdcPasswordType
	// 银行密码
	InstallID TThostFtdcInstallIDType
	// 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
	Digest TThostFtdcDigestType
	// 摘要
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货公司银行编码
	TID TThostFtdcTIDType
	// 交易ID
	ReserveOpenAccStas TThostFtdcReserveOpenAccStasType
	// 预约开户状态
	ErrorID TThostFtdcErrorIDType
	// 错误代码
	ErrorMsg TThostFtdcErrorMsgType
	// 错误信息
}
type CThostFtdcAccountPropertyField struct {
	//银行账户属性
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	AccountID TThostFtdcAccountIDType
	// 投资者帐号
	BankID TThostFtdcBankIDType
	// 银行统一标识类型
	BankAccount TThostFtdcBankAccountType
	// 银行账户
	OpenName TThostFtdcInvestorFullNameType
	// 银行账户的开户人名称
	OpenBank TThostFtdcOpenBankType
	// 银行账户的开户行
	IsActive TThostFtdcBoolType
	// 是否活跃
	AccountSourceType TThostFtdcAccountSourceTypeType
	// 账户来源
	OpenDate TThostFtdcDateType
	// 开户日期
	CancelDate TThostFtdcDateType
	// 注销日期
	OperatorID TThostFtdcOperatorIDType
	// 录入员代码
	OperateDate TThostFtdcDateType
	// 录入日期
	OperateTime TThostFtdcTimeType
	// 录入时间
	CurrencyID TThostFtdcCurrencyIDType
	// 币种代码
}
type CThostFtdcQryCurrDRIdentityField struct {
	//查询当前交易中心
	DRIdentityID TThostFtdcDRIdentityIDType
	// 交易中心代码
}
type CThostFtdcCurrDRIdentityField struct {
	//当前交易中心
	DRIdentityID TThostFtdcDRIdentityIDType
	// 交易中心代码
}
type CThostFtdcQrySecAgentCheckModeField struct {
	//查询二级代理商资金校验模式
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者代码
}
