package message

// OrderAccepted is the outbound acknowledgement of a valid EnterOrder (type 'A').
// Total size: 130 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'A'
//	     1    8  Timestamp timestamp
//	     9    8  int64    orderToken
//	    17    4  int32    orderBookId
//	    21    1  Alpha    side
//	    22    8  int64    orderId
//	    30    8  int64    quantity
//	    38    8  Price    price
//	    46    1  int8     timeInForce
//	    47    1  int8     openClose
//	    48   16  Alpha    clientAccount
//	    64    1  int8     orderState
//	    65   15  Alpha    customerInfo
//	    80   32  Alpha    exchangeInfo
//	   112    8  int64    displayQuantity
//	   120    1  int8     orderType
//	   121    2  int16    timeInForceData
//	   123    1  int8     orderCapacity
//	   124    4  int32    selfMatchPreventionKey
//	   128    2  int16    attributes
type OrderAccepted struct {
	OuchMessageType        int8  // always 'A'
	Timestamp              int64 // nanoseconds since Unix epoch
	OrderToken             int64
	OrderBookID            int32
	Side                   Side
	OrderID                int64
	Quantity               int64
	Price                  int64
	TimeInForce            TimeInForce
	OpenClose              OpenClose
	ClientAccount          string // max 15 chars
	OrderState             OrderState
	CustomerInfo           string // max 15 chars
	ExchangeInfo           ExchangeInfo
	DisplayQuantity        int64
	OrderType              OrderType
	TimeInForceData        int16
	OrderCapacity          OrderCapacity
	SelfMatchPreventionKey int32
	Attributes             Attributes
}

// OrderRejected is returned when Enter, Replace, or Cancel is rejected (type 'J').
// Total size: 29 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'J'
//	     1    8  Timestamp timestamp
//	     9    8  int64    orderToken
//	    17    8  int64    orderId
//	    25    4  int32    rejectCode
type OrderRejected struct {
	OuchMessageType int8 // always 'J'
	Timestamp       int64
	OrderToken      int64
	OrderID         int64
	RejectCode      int32
}

// OrderReplaced acknowledges a successful ReplaceOrder (type 'U' outbound).
// Total size: 138 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'U'
//	     1    8  Timestamp timestamp
//	     9    8  int64    replacementOrderToken
//	    17    8  int64    previousOrderToken
//	    25    4  int32    orderBookId
//	    29    1  Alpha    side
//	    30    8  int64    orderId
//	    38    8  int64    quantity
//	    46    8  Price    price
//	    54    1  int8     timeInForce
//	    55    1  int8     openClose
//	    56   16  Alpha    clientAccount
//	    72    1  int8     orderState
//	    73   15  Alpha    customerInfo
//	    88   32  Alpha    exchangeInfo
//	   120    8  int64    displayQuantity
//	   128    1  int8     orderType
//	   129    2  int16    timeInForceData
//	   131    1  int8     orderCapacity
//	   132    4  int32    selfMatchPreventionKey
//	   136    2  int16    attributes
type OrderReplaced struct {
	OuchMessageType        int8 // always 'U'
	Timestamp              int64
	ReplacementOrderToken  int64
	PreviousOrderToken     int64
	OrderBookID            int32
	Side                   Side
	OrderID                int64
	Quantity               int64
	Price                  int64
	TimeInForce            TimeInForce
	OpenClose              OpenClose
	ClientAccount          string // max 15 chars
	OrderState             OrderState
	CustomerInfo           string // max 15 chars
	ExchangeInfo           ExchangeInfo
	DisplayQuantity        int64
	OrderType              OrderType
	TimeInForceData        int16
	OrderCapacity          OrderCapacity
	SelfMatchPreventionKey int32
	Attributes             Attributes
}

// OrderCanceled informs that an order has been canceled (type 'C').
// Total size: 31 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'C'
//	     1    8  Timestamp timestamp
//	     9    8  int64    orderToken
//	    17    4  int32    orderBookId
//	    21    1  Alpha    side
//	    22    8  int64    orderId
//	    30    1  int8     cancelReason
type OrderCanceled struct {
	OuchMessageType int8 // always 'C'
	Timestamp       int64
	OrderToken      int64
	OrderBookID     int32
	Side            Side
	OrderID         int64
	CancelReason    CancelReason
}

// OrderExecuted is returned on a partial or full fill (type 'E').
// Total size: 50 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'E'
//	     1    8  Timestamp timestamp
//	     9    8  int64    orderToken
//	    17    4  int32    orderBookId
//	    21    8  int64    tradeQuantity
//	    29    8  Price    tradePrice
//	    37    8  int64    matchId
//	    45    4  int32    comboGroupId
//	    49    1  int8     dealSource
type OrderExecuted struct {
	OuchMessageType int8 // always 'E'
	Timestamp       int64
	OrderToken      int64
	OrderBookID     int32
	TradeQuantity   int64
	TradePrice      int64
	MatchID         int64
	ComboGroupID    int32
	DealSource      DealSource
}
