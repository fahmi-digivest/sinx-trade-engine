package message

// EnterOrder is the inbound message to submit a new order (type 'O').
// Total size: 113 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'O'
//	     1    8  int64    orderToken
//	     9    4  int32    orderBookId
//	    13    1  Alpha    side
//	    14    8  int64    quantity
//	    22    8  Price    price
//	    30    1  int8     timeInForce
//	    31    1  int8     openClose
//	    32   16  Alpha    clientAccount
//	    48   15  Alpha    customerInfo
//	    63   32  Alpha    exchangeInfo
//	    95    8  int64    displayQuantity
//	   103    1  int8     orderType
//	   104    2  int16    timeInForceData
//	   106    1  int8     orderCapacity
//	   107    4  int32    selfMatchPreventionKey
//	   111    2  int16    attributes
type EnterOrder struct {
	OuchMessageType        int8 // always 'O'
	OrderToken             int64
	OrderBookID            int32
	Side                   Side
	Quantity               int64
	Price                  int64 // MinLongPrice = market price
	TimeInForce            TimeInForce
	OpenClose              OpenClose
	ClientAccount          string // max 15 chars, 16th byte is null
	CustomerInfo           string // max 15 chars
	ExchangeInfo           ExchangeInfo
	DisplayQuantity        int64
	OrderType              OrderType
	TimeInForceData        int16
	OrderCapacity          OrderCapacity
	SelfMatchPreventionKey int32
	Attributes             Attributes
}

// ReplaceOrder is the inbound message to modify an existing order (type 'U').
// Total size: 112 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'U'
//	     1    8  int64    existingOrderToken
//	     9    8  int64    replacementOrderToken
//	    17    8  int64    quantity
//	    25    8  Price    price
//	    33    1  int8     openClose
//	    34   16  Alpha    clientAccount
//	    50   15  Alpha    customerInfo
//	    65   32  Alpha    exchangeInfo
//	    97    8  int64    displayQuantity
//	   105    1  int8     timeInForce
//	   106    2  int16    timeInForceData
//	   108    4  int32    selfMatchPreventionKey
type ReplaceOrder struct {
	OuchMessageType        int8 // always 'U'
	ExistingOrderToken     int64
	ReplacementOrderToken  int64
	Quantity               int64
	Price                  int64 // MinLongPrice = no change
	OpenClose              OpenClose
	ClientAccount          string // first 15 chars used; empty means no change
	CustomerInfo           string // max 15 chars
	ExchangeInfo           ExchangeInfo
	DisplayQuantity        int64 // 0 = no change
	TimeInForce            TimeInForce
	TimeInForceData        int16
	SelfMatchPreventionKey int32
}

// CancelOrder is the inbound message to cancel by order token (type 'X').
// Total size: 9 bytes.
//
//	Offset  Len  Type   Field
//	     0    1  int8   ouchMessageType = 'X'
//	     1    8  int64  orderToken
type CancelOrder struct {
	OuchMessageType int8 // always 'X'
	OrderToken      int64
}

// CancelOrderByID is the inbound message to cancel using the system-assigned Order ID (type 'Y').
// Total size: 14 bytes.
//
//	Offset  Len  Type     Field
//	     0    1  int8     ouchMessageType = 'Y'
//	     1    4  int32    orderBookId
//	     5    1  Alpha    side
//	     6    8  int64    orderId
type CancelOrderByID struct {
	OuchMessageType int8 // always 'Y'
	OrderBookID     int32
	Side            Side
	OrderID         int64
}
