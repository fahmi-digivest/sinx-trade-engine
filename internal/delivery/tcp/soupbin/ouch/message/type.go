package message

// MessageType identifies each OUCH message.
type MessageType = int8

const (
	// Inbound
	MsgTypeEnterOrder      MessageType = 'O'
	MsgTypeReplaceOrder    MessageType = 'U'
	MsgTypeCancelOrder     MessageType = 'X'
	MsgTypeCancelOrderByID MessageType = 'Y'

	// Outbound
	MsgTypeOrderAccepted MessageType = 'A'
	MsgTypeOrderRejected MessageType = 'J'
	MsgTypeOrderReplaced MessageType = 'U' // same byte as ReplaceOrder (inbound vs outbound)
	MsgTypeOrderCanceled MessageType = 'C'
	MsgTypeOrderExecuted MessageType = 'E'
)

// Side represents the order side.
type Side byte

const (
	SideBuy       Side = 'B'
	SideSell      Side = 'S'
	SideShortSell Side = 'T'
)

func (s Side) String() string {
	switch s {
	case SideBuy:
		return "Buy"
	case SideSell:
		return "Sell"
	case SideShortSell:
		return "ShortSell"
	default:
		return "Unknown"
	}
}

// TimeInForce represents TIF values.
type TimeInForce int8

const (
	TIFUndefined TimeInForce = 0
	TIFDay       TimeInForce = 1
	TIFFak       TimeInForce = 3
	TIFFok       TimeInForce = 4
	TIFGTS       TimeInForce = 5
)

func (t TimeInForce) String() string {
	switch t {
	case TIFUndefined:
		return "Undefined"
	case TIFDay:
		return "Day"
	case TIFFak:
		return "FAK"
	case TIFFok:
		return "FOK"
	case TIFGTS:
		return "GTS"
	default:
		return "Unknown"
	}
}

// OpenClose specifies position handling.
type OpenClose int8

const (
	OpenCloseDefault        OpenClose = 0
	OpenCloseOpen           OpenClose = 1
	OpenCloseCloseNet       OpenClose = 2
	OpenCloseMandatoryClose OpenClose = 3
)

// OrderType defines type of order.
type OrderType int8

const (
	OrderTypeLimit         OrderType = 1
	OrderTypeMarket        OrderType = 2
	OrderTypeMarketToLimit OrderType = 3
)

func (o OrderType) String() string {
	switch o {
	case OrderTypeLimit:
		return "Limit"
	case OrderTypeMarket:
		return "Market"
	case OrderTypeMarketToLimit:
		return "MarketToLimit"
	default:
		return "Unknown"
	}
}

// OrderCapacity defines capacity values.
type OrderCapacity int8

const (
	OrderCapacityUndefined         OrderCapacity = 0
	OrderCapacityAgency            OrderCapacity = 1
	OrderCapacityProprietary       OrderCapacity = 2
	OrderCapacityIndividual        OrderCapacity = 3
	OrderCapacityPrincipal         OrderCapacity = 4
	OrderCapacityRiskLessPrincipal OrderCapacity = 5
)

// OrderState represents order book state.
type OrderState int8

const (
	OrderStateOnBook    OrderState = 1
	OrderStateNotOnBook OrderState = 2
)

// Attributes bitmask values.
type Attributes int16

const (
	AttributesUndefined          Attributes = 0
	AttributesMarketBid          Attributes = 1
	AttributesPriceStabilization Attributes = 2
	AttributesMargin             Attributes = 3
)

// SettlementMethod for ExchangeInfo.
type SettlementMethod byte

const (
	SettlementDVP  SettlementMethod = '1'
	SettlementDFOP SettlementMethod = '2'
)

// CancelReason codes returned in OrderCanceled.
type CancelReason int8

const (
	CancelReasonByUser                    CancelReason = 1
	CancelReasonBySystemAfterTrade        CancelReason = 3
	CancelReasonBySystemAfterNewOrder     CancelReason = 6
	CancelReasonBySystemAfterConverted    CancelReason = 8
	CancelReasonBySystem                  CancelReason = 9
	CancelReasonByProxyUser               CancelReason = 10
	CancelReasonBySystemAfterNewTriggered CancelReason = 12
	CancelReasonBySystemForHiddenOrder    CancelReason = 13
	CancelReasonBySystemOrderChanged      CancelReason = 19
	CancelReasonBySystemInstrumentSession CancelReason = 20
	CancelReasonSelfMatchPrevention       CancelReason = 43
	CancelReasonCircuitBreaker            CancelReason = 44
	CancelReasonCreditLimits              CancelReason = 45
	CancelReasonCorporateAction           CancelReason = 58
	CancelReasonSelfMatchPrevDefault      CancelReason = 59
)

// DealSource values in OrderExecuted.
type DealSource int8

const (
	DealSourceMatchedBySystem          DealSource = 1
	DealSourceComboMatchCombo          DealSource = 7
	DealSourceMatchedInAuction         DealSource = 20
	DealSourceTailorMadeCombination    DealSource = 36
	DealSourceComboMatchedOutrightLegs DealSource = 43
)

// MinLongPrice is the sentinel value meaning "no price / market price".
const MinLongPrice int64 = -9223372036854775808
