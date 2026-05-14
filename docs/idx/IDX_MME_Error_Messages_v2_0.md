# IDX MME Error Messages

**Document Version:** 2.0  
**Date:** 25 February 2026  
**Source:** Indonesia Stock Exchange (IDX) — JATS MME Error Message Reference

---

## Introduction

Depending on where in the system, at what layer, etc. a validation takes place, JATS MME will send different kinds of error messages. Some are error message numbers listed in the Error Message Reference, which translates these numbers into readable text. FIX and OUCH users can receive these error message numbers.

### Field Descriptions

| Field | Description |
|-------|-------------|
| **Alarm Code** | The alarm code name |
| **Number** | The alarm code number |
| **Subsystem** | The subsystem where the error occurred |
| **Severity Code** | E = Error, W = Warning |
| **Message Text** | The message shown on the screen |
| **Explanation** | The reason for the error occurrence |
| **User Action** | The recommended action to take |

---

## Table of Contents

- [AUTH Error Messages](#auth-error-messages) (13 errors)
- [AUTHC Error Messages](#authc-error-messages) (13 errors)
- [AF Error Messages](#af-error-messages) (5 errors)
- [DE Error Messages](#de-error-messages) (7 errors)
- [ME Error Messages](#me-error-messages) (176 errors)
- [RDM Error Messages](#rdm-error-messages) (26 errors)
- [PH Error Messages](#ph-error-messages) (12 errors)
- [PTR Error Messages](#ptr-error-messages) (652 errors)
- [QM Error Messages](#qm-error-messages) (21 errors)
- [SM Error Messages](#sm-error-messages) (26 errors)
- [TH Error Messages](#th-error-messages) (7 errors)

---

## AUTH Error Messages

### `AUTH_ILL_TRT_FOR_USER` (-80003)

| Field | Value |
|-------|-------|
| **Number** | `-80003` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Transaction is disallowed for this user |

### `AUTH_OB_CLOSED` (-80005)

| Field | Value |
|-------|-------|
| **Number** | `-80005` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal transaction at this time |
| **Explanation** | The order book is closed |

### `AUTH_LIU_NOTFOU` (-80007)

| Field | Value |
|-------|-------|
| **Number** | `-80007` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Actor is not allowed to act in this market segment |

### `AUTH_CLI_NOTFOU` (-80009)

| Field | Value |
|-------|-------|
| **Number** | `-80009` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | The participant is not allowed to act in this market segment |

### `AUTH_USER_NOT_ACTIVE` (-80013)

| Field | Value |
|-------|-------|
| **Number** | `-80013` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Actor not in active state |

### `AUTH_PART_NOT_ACTIVE` (-80015)

| Field | Value |
|-------|-------|
| **Number** | `-80015` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Participant not in active state |

### `AUTH_NOT_TRADED` (-80017)

| Field | Value |
|-------|-------|
| **Number** | `-80017` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to trade in order book |

### `AUTH_OBL_NOTFOU` (-80021)

| Field | Value |
|-------|-------|
| **Number** | `-80021` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Member obligation not found |
| **Explanation** | Specified member does not exist in the member obligation table |

### `AUTH_TRC_INS_NOTFOUND` (-80023)

| Field | Value |
|-------|-------|
| **Number** | `-80023` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Trading report class not valid for the instrument |

### `AUTH_TRC_SST_NOTFOUND` (-80025)

| Field | Value |
|-------|-------|
| **Number** | `-80025` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Trading report class not valid for session state |

### `AUTH_TRC_UST_NOTFOUND` (-80027)

| Field | Value |
|-------|-------|
| **Number** | `-80027` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Trading report class not valid for the actor |

### `AUTH_LIU_NOTFOU_BID` (-80031)

| Field | Value |
|-------|-------|
| **Number** | `-80031` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Actor is not allowed to place bid orders in this market segment |
| **Explanation** | The actor is not allowed to place bid orders in this market segment |

### `AUTH_LIU_NOTFOU_ASK` (-80033)

| Field | Value |
|-------|-------|
| **Number** | `-80033` |
| **Subsystem** | AUTH |
| **Severity** | 🟡 WARNING |
| **Message** | Actor is not allowed to place ask orders in this market segment |
| **Explanation** | The actor is not allowed to place ask orders in this market segment |
| **User Action** | None 3 AUTHC Error Messages |

---

## AUTHC Error Messages

### `AUTHC_OTHER` (-50002)

| Field | Value |
|-------|-------|
| **Number** | `-50002` |
| **Subsystem** | AUTHC |
| **Severity** | 🔴 ERROR |
| **Message** | Login rejected for technical reasons |
| **Explanation** | Something undefined has happened. |

### `AUTHC_PASSWORD_EXPIRED` (-50007)

| Field | Value |
|-------|-------|
| **Number** | `-50007` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The password has expired |
| **Explanation** | The entered password is no longer valid and must be reset/updated. |
| **User Action** | Reset the password. |

### `AUTHC_PASSWORD_POLICY` (-50009)

| Field | Value |
|-------|-------|
| **Number** | `-50009` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The password does not conform to the password policy |
| **Explanation** | The password does not conform to the password policy. |
| **User Action** | Update the password according to the password policy. |

### `AUTHC_ACCOUNT_DISABLED` (-50011)

| Field | Value |
|-------|-------|
| **Number** | `-50011` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The account is disabled |
| **Explanation** | The account has been disabled and can not be used to login. |
| **User Action** | Contact the exchange. |

### `AUTHC_ACCOUNT_LOCKED` (-50013)

| Field | Value |
|-------|-------|
| **Number** | `-50013` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The account is locked |
| **Explanation** | The account has been locked and must be unlocked. |
| **User Action** | Contact the exchange. |

### `AUTHC_SYSTEM_LOGOUT` (-50015)

| Field | Value |
|-------|-------|
| **Number** | `-50015` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The user has been logged out by the system |
| **Explanation** | The user has been logged out by the system. |

### `AUTHC_AUTHENTICATOR_LOGOUT` (-50017)

| Field | Value |
|-------|-------|
| **Number** | `-50017` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The user has been logged out by the IDP |
| **Explanation** | The user has been logged out by the IDP. |

### `AUTHC_APPLICATION_LOGOUT` (-50019)

| Field | Value |
|-------|-------|
| **Number** | `-50019` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The user has been logged out by the application |
| **Explanation** | The user has been logged out by the application. |

### `AUTHC_SESSION_DROPPED` (-50023)

| Field | Value |
|-------|-------|
| **Number** | `-50023` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | The session was dropped |
| **Explanation** | The session was dropped. |

### `AUTHC_DISABLED` (-50025)

| Field | Value |
|-------|-------|
| **Number** | `-50025` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | Logon is disabled |
| **Explanation** | Logon has been disabled and users can not login. |
| **User Action** | Contact the exchange. |

### `AUTHC_BLOCKED` (-50027)

| Field | Value |
|-------|-------|
| **Number** | `-50027` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | Logon is blocked |
| **Explanation** | Logon is blocked. |

### `AUTHC_SYSTEM_BLOCKED` (-50029)

| Field | Value |
|-------|-------|
| **Number** | `-50029` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | Logon is blocked by the system |
| **Explanation** | Logon is blocked by the system. |

### `AUTHC_APPLICATION_BLOCKED` (-50033)

| Field | Value |
|-------|-------|
| **Number** | `-50033` |
| **Subsystem** | AUTHC |
| **Severity** | 🟡 WARNING |
| **Message** | Logon is blocked by the application |
| **Explanation** | Logon is blocked by the application. |
| **User Action** | None. 4 AF Error Messages |

---

## AF Error Messages

### `AF_REQUEST_NOT_FOUND` (-1200002)

| Field | Value |
|-------|-------|
| **Number** | `-1200002` |
| **Subsystem** | AF |
| **Severity** | 🔴 ERROR |
| **Message** | Countersign Request not found. |

### `AF_OWN_COUNTERSIGN` (-1200004)

| Field | Value |
|-------|-------|
| **Number** | `-1200004` |
| **Subsystem** | AF |
| **Severity** | 🔴 ERROR |
| **Message** | Actor is accepting or rejecting its own countersign request. |
| **Explanation** | An Actor is not allowed to accept or reject its own countersign request. |

### `AF_OTHER_CANCEL` (-1200006)

| Field | Value |
|-------|-------|
| **Number** | `-1200006` |
| **Subsystem** | AF |
| **Severity** | 🔴 ERROR |
| **Message** | Actor is cancelling somebody else's countersign request. |
| **Explanation** | An Actor is only allowed to cancel its own countersign request. |

### `AF_INV_SET_EXCHANGE_RATE` (-1200008)

| Field | Value |
|-------|-------|
| **Number** | `-1200008` |
| **Subsystem** | AF |
| **Severity** | 🔴 ERROR |
| **Message** | Set Exchange Rate message is not valid |
| **Explanation** | The set exchange rate does not contain valid information |

### `AF_INV_ACTOR` (-1200009)

| Field | Value |
|-------|-------|
| **Number** | `-1200009` |
| **Subsystem** | AF |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Actor. |
| **Explanation** | Actor is invalid. |
| **User Action** | None. 5 DE Error Messages |

---

## DE Error Messages

### `OUCH_DUPLICATE_TOKEN` (-800002)

| Field | Value |
|-------|-------|
| **Number** | `-800002` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | The token is not unique |
| **Explanation** | The token has already been used for this session. |

### `OUCH_UNKNOWN_TOKEN` (-800004)

| Field | Value |
|-------|-------|
| **Number** | `-800004` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | The token is not known |
| **Explanation** | Attempt to use a token that has not been previously sent. |

### `OUCH_INVALID_ORDERBOOK` (-800006)

| Field | Value |
|-------|-------|
| **Number** | `-800006` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | Invalid orderbook |
| **Explanation** | The orderbook specified is not available. |

### `OUCH_INVALID_SIDE` (-800008)

| Field | Value |
|-------|-------|
| **Number** | `-800008` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | Invalid side |
| **Explanation** | The side specified is invalid. |

### `OUCH_INVALID_TIF` (-800010)

| Field | Value |
|-------|-------|
| **Number** | `-800010` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | Invalid Time In Force |
| **Explanation** | The time in force specified is invalid. |

### `OUCH_CAN_NOT_CANCEL` (-800014)

| Field | Value |
|-------|-------|
| **Number** | `-800014` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | The order can not be cancelled |
| **Explanation** | The order was rejected or has not been accepted by the exchange yet. |

### `OUCH_THROTTLING` (-800020)

| Field | Value |
|-------|-------|
| **Number** | `-800020` |
| **Subsystem** | DE |
| **Severity** | 🔴 ERROR |
| **Message** | Throttling limit exceeded |
| **Explanation** | Throttling limit exceeded. |
| **User Action** | Enter transactions at a lower pace. 6 ME Error Messages |

---

## ME Error Messages

### `ME_MARKET_MAKER_PROTECTION` (-405044)

| Field | Value |
|-------|-------|
| **Number** | `-405044` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Market Maker Protection triggered on underlying. |
| **User Action** | Set override flag to send quote forcefully or wait for quotation frozen time to pass. |

### `ME_INVALID_LONG_CONSUMPTION_ACCOUNT` (-405058)

| Field | Value |
|-------|-------|
| **Number** | `-405058` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Available long quantity is less than the current consumption for the account. |
| **Explanation** | The received SetPositionLimit message is invalid. |
| **User Action** | Set higher limit. |

### `ME_INVALID_LONG_CONSUMPTION_INVESTOR` (-405060)

| Field | Value |
|-------|-------|
| **Number** | `-405060` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Available long quantity is less than the current consumption for the investor. |
| **Explanation** | The received SetPositionLimit message is invalid. |
| **User Action** | Set higher limit. |

### `ME_INVALID_SHORT_CONSUMPTION_ACCOUNT` (-405062)

| Field | Value |
|-------|-------|
| **Number** | `-405062` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Available loan quantity is less than the current consumption for the account. |
| **Explanation** | The received SetPositionLimit message is invalid. |
| **User Action** | Set higher limit. |

### `ME_INVALID_SHORT_CONSUMPTION_INVESTOR` (-405064)

| Field | Value |
|-------|-------|
| **Number** | `-405064` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Available loan quantity is less than the current consumption for the investor. |
| **Explanation** | The received SetPositionLimit message is invalid. |
| **User Action** | Set higher limit. |

### `ME_INVALID_DELTA_FACTOR` (-405065)

| Field | Value |
|-------|-------|
| **Number** | `-405065` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid delta factor. |
| **User Action** | Correct the delta factor. |

### `ME_INVALID_DELTA_ORDERBOOK` (-405067)

| Field | Value |
|-------|-------|
| **Number** | `-405067` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Delta protection is applicable for option/future/forward only. |
| **User Action** | Check orderBook. |

### `ME_INVALID_VEGA_FACTOR` (-405069)

| Field | Value |
|-------|-------|
| **Number** | `-405069` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid vega factor. |
| **User Action** | Correct the vega factor. |

### `ME_INVALID_VEGA_ORDERBOOK` (-405071)

| Field | Value |
|-------|-------|
| **Number** | `-405071` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Vega protection is applicable for options only. |
| **User Action** | Check orderBook. |

### `ME_MATCH_BID_ASK` (-420023)

| Field | Value |
|-------|-------|
| **Number** | `-420023` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order must specify bid or ask. |
| **Explanation** | The given order does not specify whether it is a bid or ask. The order is ignored by the system. |
| **User Action** | Correct your program. |

### `ME_MATCH_VALIDITY` (-420025)

| Field | Value |
|-------|-------|
| **Number** | `-420025` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Given time validity is not allowed. |
| **Explanation** | The time validity given in the order is not valid at the exchange. |
| **User Action** | Enter a valid time for this order. |

### `ME_MATCH_MARBOUNCE` (-420027)

| Field | Value |
|-------|-------|
| **Number** | `-420027` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market-price orders must be of type Fill or Kill or Immediate or Cancel in this trading state. |
| **Explanation** | The time validity must be zero for market-price orders. The system does not store market-price orders in the order book and thus these orders cannot have a time limit. |
| **User Action** | Correct your program. |

### `ME_MATCH_PREMIUM` (-420029)

| Field | Value |
|-------|-------|
| **Number** | `-420029` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Given premium is not allowed. |
| **Explanation** | The premium must be aligned at the price ticks for the given instrument. See trading rules for the instrument. |
| **User Action** | Correct the price. |

### `ME_MATCH_QUANTITY` (-420045)

| Field | Value |
|-------|-------|
| **Number** | `-420045` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal quantity in order. |
| **Explanation** | The quantity must be in the interval 1 to 32767. For block orders the quantity must also be a positive multiple of the block size. Orders with illegal quantities are ignored by the system. |
| **User Action** | Enter a valid quantity. |

### `ME_MATCH_ILL_ORDER_TYPE` (-420076)

| Field | Value |
|-------|-------|
| **Number** | `-420076` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Unknown order type. |
| **Explanation** | Order transaction must specify limit order (1) or market order (2) in order_type_c. |
| **User Action** | Correct your program. |

### `ME_MATCH_MKT_ORDER_PRICE` (-420078)

| Field | Value |
|-------|-------|
| **Number** | `-420078` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | A market order must not specify a price. |
| **Explanation** | A market order must not specify a price to avoid confusion. |
| **User Action** | Correct your program. |

### `ME_MATCH_BLOCK_MAX_LEGS` (-420093)

| Field | Value |
|-------|-------|
| **Number** | `-420093` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Too many legs for a block order. |
| **Explanation** | Illegal number of legs in a block order. |
| **User Action** | Correct your program. |

### `ME_MATCH_BLOCK_SERIES` (-420095)

| Field | Value |
|-------|-------|
| **Number** | `-420095` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Series appears twice on the same side in a block order. |
| **Explanation** | One block must not have the same duplicated series on the same side. |
| **User Action** | Correct your program. |

### `ME_MATCH_ILL_BUY_SELL_OP` (-420115)

| Field | Value |
|-------|-------|
| **Number** | `-420115` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal side on order |
| **Explanation** | The actor is not allowed to place an order on the given side of the order book |
| **User Action** | Change the side of the order |

### `ME_MATCH_MEM_CRS_NOT_ALLOWED` (-420129)

| Field | Value |
|-------|-------|
| **Number** | `-420129` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Crossing your own orders is not allowed in this instrument type. |
| **Explanation** | The incoming order resulted in an order book with crossed prices. |
| **User Action** | Modify your order and retry. |

### `ME_MATCH_PRICE_LIMIT` (-420131)

| Field | Value |
|-------|-------|
| **Number** | `-420131` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The price is outside the allowed price limits for this instrument. |
| **Explanation** | The price is outside the allowed price limits. |
| **User Action** | Modify your order according to the defined price limits and retry. |

### `ME_MATCH_INV_STP_COND` (-420133)

| Field | Value |
|-------|-------|
| **Number** | `-420133` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The given stop condition is invalid. |
| **Explanation** | A actor has entered a stop condition that is not allowed. |
| **User Action** | Modify your order so that it contains a valid stop condition. |

### `ME_MATCH_INV_HIDDEN` (-420135)

| Field | Value |
|-------|-------|
| **Number** | `-420135` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume is not allowed for this order type. |
| **Explanation** | The instrument type is configured to deny hidden volume. |
| **User Action** | Modify your order and retry. |

### `ME_MATCH_INV_SHOWN` (-420137)

| Field | Value |
|-------|-------|
| **Number** | `-420137` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The allowed ratio between shown quantity and total quantity has been exceeded. |
| **Explanation** | Shown volume must be greater than zero and less than total volume. |
| **User Action** | Modify your order and retry. |

### `ME_MATCH_NOT_AUTH` (-420139)

| Field | Value |
|-------|-------|
| **Number** | `-420139` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The actor is not authorized to do on-behalf transactions. |
| **User Action** | Modify your order and retry. |

### `ME_MATCH_SHOWN_TOO_SMALL` (-420141)

| Field | Value |
|-------|-------|
| **Number** | `-420141` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Shown volume too small. |
| **Explanation** | Shown volume must be greater than the value specified in the CDB. |
| **User Action** | Modify your order and retry. |

### `ME_MATCH_NO_WILD_CARD` (-420153)

| Field | Value |
|-------|-------|
| **Number** | `-420153` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Wildcards ( * and % ) and spaces are not allowed in the account. |
| **Explanation** | The supplied account contains spaces or wildcard characters ( * and % ). |
| **User Action** | Correct your account. |

### `ME_MATCH_INV_OP_CLS_REQ` (-420155)

| Field | Value |
|-------|-------|
| **Number** | `-420155` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid value in open_close_request. |
| **Explanation** | Invalid value in open_close_request. A value of 0 - 3 is valid + 4 and 5 for combo series. |
| **User Action** | Correct your order. |

### `ME_MATCH_INV_ACCOUNT` (-420159)

| Field | Value |
|-------|-------|
| **Number** | `-420159` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Missing or invalid account. |
| **Explanation** | The supplied account is either missing or invalid. |
| **User Action** | Correct your order. |

### `ME_MATCH_ORD_NOT_FOU` (-420177)

| Field | Value |
|-------|-------|
| **Number** | `-420177` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The given order was not found in the order book. |
| **User Action** | Send an order that exists in the order book. |

### `ME_MATCH_BLMT_ORDER_PRICE` (-420184)

| Field | Value |
|-------|-------|
| **Number** | `-420184` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | A best limit order must not specify a price. |
| **User Action** | Correct your program. |

### `ME_MATCH_INS_FILL_OR_KILL` (-420193)

| Field | Value |
|-------|-------|
| **Number** | `-420193` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill or kill orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_FILL_OR_KILL` (-420195)

| Field | Value |
|-------|-------|
| **Number** | `-420195` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill or kill orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_FILL_OR_KILL` (-420197)

| Field | Value |
|-------|-------|
| **Number** | `-420197` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill or kill orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_IMMEDIATE_OR_CANCEL` (-420199)

| Field | Value |
|-------|-------|
| **Number** | `-420199` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Immediate or Cancel orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_IMMEDIATE_OR_CANCEL` (-420201)

| Field | Value |
|-------|-------|
| **Number** | `-420201` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Immediate or Cancel orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_IMMEDIATE_OR_CANCEL` (-420203)

| Field | Value |
|-------|-------|
| **Number** | `-420203` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Immediate or Cancel orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_FILL_AND_STORE` (-420205)

| Field | Value |
|-------|-------|
| **Number** | `-420205` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill and store orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_FILL_AND_STORE` (-420207)

| Field | Value |
|-------|-------|
| **Number** | `-420207` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill and store orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_FILL_AND_STORE` (-420209)

| Field | Value |
|-------|-------|
| **Number** | `-420209` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Fill and store orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_LIMIT_ORD` (-420211)

| Field | Value |
|-------|-------|
| **Number** | `-420211` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Limit orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_LIMIT_ORD` (-420213)

| Field | Value |
|-------|-------|
| **Number** | `-420213` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Limit orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_LIMIT_ORD` (-420215)

| Field | Value |
|-------|-------|
| **Number** | `-420215` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Limit orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_MARKET_ORD` (-420217)

| Field | Value |
|-------|-------|
| **Number** | `-420217` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_MARKET_ORD` (-420219)

| Field | Value |
|-------|-------|
| **Number** | `-420219` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_MARKET_ORD` (-420221)

| Field | Value |
|-------|-------|
| **Number** | `-420221` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_ALL_OR_NONE` (-420223)

| Field | Value |
|-------|-------|
| **Number** | `-420223` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | All or none orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_ALL_OR_NONE` (-420225)

| Field | Value |
|-------|-------|
| **Number** | `-420225` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | All or none orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_ALL_OR_NONE` (-420227)

| Field | Value |
|-------|-------|
| **Number** | `-420227` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | All or none orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_ILL_ORDER_TYPE_INT` (-420237)

| Field | Value |
|-------|-------|
| **Number** | `-420237` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal order type for this instrument type. |
| **Explanation** | The order type is disabled for the instrument type. |
| **User Action** | Correct your program. |

### `ME_MATCH_INS_IMBALANCE` (-420245)

| Field | Value |
|-------|-------|
| **Number** | `-420245` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Imbalance order is not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_IMBALANCE` (-420247)

| Field | Value |
|-------|-------|
| **Number** | `-420247` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Imbalance order is not allowed in current trading state |
| **Explanation** | Imbalance order is not allowed in current state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_IMBALANCE` (-420249)

| Field | Value |
|-------|-------|
| **Number** | `-420249` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Imbalance order is not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_MTL_ROUND_LOT` (-420251)

| Field | Value |
|-------|-------|
| **Number** | `-420251` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market to limit is not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_MTL_ROUND_LOT` (-420253)

| Field | Value |
|-------|-------|
| **Number** | `-420253` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market to limit is not allowed in current trading state |
| **Explanation** | Market to limit is not allowed in current state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_MTL_ROUND_LOT` (-420255)

| Field | Value |
|-------|-------|
| **Number** | `-420255` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market to limit is not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_HIDDEN_AGGRESSIVE` (-420287)

| Field | Value |
|-------|-------|
| **Number** | `-420287` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume order is not allowed for this instrument |
| **Explanation** | Hidden volume for aggressive order is not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_HIDDEN_AGGRESSIVE` (-420289)

| Field | Value |
|-------|-------|
| **Number** | `-420289` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume order is not allowed in current trading state |
| **Explanation** | Hidden volume for aggressive order is not allowed in current state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_HIDDEN_AGGRESSIVE` (-420291)

| Field | Value |
|-------|-------|
| **Number** | `-420291` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume order is not allowed for this actor |
| **Explanation** | Hidden volume for aggressive order is not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_MIN_BLK_SIZE` (-420311)

| Field | Value |
|-------|-------|
| **Number** | `-420311` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Quantity is less than the minimum for the block size |
| **User Action** | Respecify the order. |

### `ME_MATCH_MAX_BLK_SIZE` (-420313)

| Field | Value |
|-------|-------|
| **Number** | `-420313` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Quantity exceeds the maximum for the block size |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_DECR_SHOWN_NOHIDD` (-420315)

| Field | Value |
|-------|-------|
| **Number** | `-420315` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Volume may not be decreased |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_DECR_SHOWN` (-420317)

| Field | Value |
|-------|-------|
| **Number** | `-420317` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Shown volume may not be decreased when hidden volume exists |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_DECR_HIDD` (-420319)

| Field | Value |
|-------|-------|
| **Number** | `-420319` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume may not be decreased |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_INCR_SHOWN` (-420321)

| Field | Value |
|-------|-------|
| **Number** | `-420321` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Shown volume may not be increased |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_INCR_HIDD` (-420323)

| Field | Value |
|-------|-------|
| **Number** | `-420323` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume may not be increased |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_CLIENT` (-420325)

| Field | Value |
|-------|-------|
| **Number** | `-420325` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Account (client) field may not be changed |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_PRICE_IMPR` (-420327)

| Field | Value |
|-------|-------|
| **Number** | `-420327` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The price may not be improved |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_PRICE_DISIMPR` (-420329)

| Field | Value |
|-------|-------|
| **Number** | `-420329` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The price may not be disimproved |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_TIME_EXT` (-420333)

| Field | Value |
|-------|-------|
| **Number** | `-420333` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The time validity may not be extended |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_TIME_DECR` (-420335)

| Field | Value |
|-------|-------|
| **Number** | `-420335` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The time validity may not be decreased |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_EXCH_ORDER_TYPE` (-420339)

| Field | Value |
|-------|-------|
| **Number** | `-420339` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The exchange specific order type may not be changed |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_DECR_SHOWN_NOHIDD` (-420343)

| Field | Value |
|-------|-------|
| **Number** | `-420343` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Volume may not be decreased in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_DECR_SHOWN` (-420345)

| Field | Value |
|-------|-------|
| **Number** | `-420345` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Shown volume may not be decreased when hidden volume exists in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_DECR_HIDD` (-420347)

| Field | Value |
|-------|-------|
| **Number** | `-420347` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume may not be decreased in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_INCR_SHOWN` (-420349)

| Field | Value |
|-------|-------|
| **Number** | `-420349` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Shown volume may not be increased in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_INCR_HIDD` (-420351)

| Field | Value |
|-------|-------|
| **Number** | `-420351` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Hidden volume may not be increased in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_CLIENT` (-420353)

| Field | Value |
|-------|-------|
| **Number** | `-420353` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Account (client) field may not be changed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_PRICE_IMPR` (-420355)

| Field | Value |
|-------|-------|
| **Number** | `-420355` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The price may not be improved in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_PRICE_DISIMPR` (-420357)

| Field | Value |
|-------|-------|
| **Number** | `-420357` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The price may not be disimproved in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_TIME_EXT` (-420361)

| Field | Value |
|-------|-------|
| **Number** | `-420361` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The time validity may not be extended in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_TIME_DECR` (-420363)

| Field | Value |
|-------|-------|
| **Number** | `-420363` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The time validity may not be decreased in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_PST_ORDER` (-420371)

| Field | Value |
|-------|-------|
| **Number** | `-420371` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price stabilization not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_PST_ORDER` (-420373)

| Field | Value |
|-------|-------|
| **Number** | `-420373` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price stabilization not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_PST_ORDER` (-420375)

| Field | Value |
|-------|-------|
| **Number** | `-420375` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price stabilization not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_SHORT_ORDER` (-420377)

| Field | Value |
|-------|-------|
| **Number** | `-420377` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Short Sell not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_SHORT_ORDER` (-420379)

| Field | Value |
|-------|-------|
| **Number** | `-420379` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Short Sell not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_SHORT_ORDER` (-420381)

| Field | Value |
|-------|-------|
| **Number** | `-420381` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Short Sell not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_MB_ORDER` (-420383)

| Field | Value |
|-------|-------|
| **Number** | `-420383` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market Bid not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_MB_ORDER` (-420385)

| Field | Value |
|-------|-------|
| **Number** | `-420385` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market Bid not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_MB_ORDER` (-420387)

| Field | Value |
|-------|-------|
| **Number** | `-420387` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market Bid not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_BL_ORDER` (-420389)

| Field | Value |
|-------|-------|
| **Number** | `-420389` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Best Limit not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_BL_ORDER` (-420391)

| Field | Value |
|-------|-------|
| **Number** | `-420391` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Best Limit not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_BL_ORDER` (-420393)

| Field | Value |
|-------|-------|
| **Number** | `-420393` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Best Limit not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_MB_SELL` (-420397)

| Field | Value |
|-------|-------|
| **Number** | `-420397` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Market Bid sell order |
| **User Action** | Respecify the order. |

### `ME_MATCH_MB_ORDER_TYPE` (-420399)

| Field | Value |
|-------|-------|
| **Number** | `-420399` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal order type for market bid order |
| **User Action** | Respecify the order. |

### `ME_MATCH_SH_BUY` (-420401)

| Field | Value |
|-------|-------|
| **Number** | `-420401` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Short Sell buy order |
| **User Action** | Respecify the order. |

### `ME_MATCH_FOK_IOC_NOT_ALLOWD` (-420411)

| Field | Value |
|-------|-------|
| **Number** | `-420411` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Entered order type is not allowed to have IoC or FoK. |
| **User Action** | Respecify the order. |

### `ME_MATCH_BEST_LIMIT_REQ` (-420415)

| Field | Value |
|-------|-------|
| **Number** | `-420415` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | No current market to establish default price. |
| **User Action** | Use best limit order type when there is a market. |

### `ME_MATCH_INV_ALWAYS_INACTIVE` (-420425)

| Field | Value |
|-------|-------|
| **Number** | `-420425` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Not possible to perform actions on the type of order. |
| **Explanation** | It is not possible to perform actions on this order type. |
| **User Action** | Respecify the order. |

### `ME_MATCH_ILL_EXCH_ORDER_TYPE` (-420426)

| Field | Value |
|-------|-------|
| **Number** | `-420426` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Invalid exchange order type. |
| **Explanation** | The exchange order type is exchange specific. The sent value is not valid for this exchange. |
| **User Action** | Correct your program. |

### `ME_MATCH_STOP_PREMIUM` (-420429)

| Field | Value |
|-------|-------|
| **Number** | `-420429` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Given stop premium is invalid. |
| **Explanation** | The stop premium must be aligned at the price ticks defined for the given instrument. See trading rules for the instrument. |
| **User Action** | Correct the stop price. |

### `ME_MATCH_SST_SSO` (-420443)

| Field | Value |
|-------|-------|
| **Number** | `-420443` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | State Type Order is not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_SSO` (-420445)

| Field | Value |
|-------|-------|
| **Number** | `-420445` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | State Type Order is not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_SSO` (-420447)

| Field | Value |
|-------|-------|
| **Number** | `-420447` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | State Type Order is not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_SSO_SPEC_NO_TYPE` (-420451)

| Field | Value |
|-------|-------|
| **Number** | `-420451` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | State Type Order must specify state type |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_EXT_T_NOT_VALID` (-420453)

| Field | Value |
|-------|-------|
| **Number** | `-420453` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | State Type valid only for State Type Orders |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_STOP_ORD` (-420459)

| Field | Value |
|-------|-------|
| **Number** | `-420459` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | STOP orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_STOP_ORD` (-420461)

| Field | Value |
|-------|-------|
| **Number** | `-420461` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Stop orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_STOP_ORD` (-420463)

| Field | Value |
|-------|-------|
| **Number** | `-420463` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Stop orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_GTS` (-420539)

| Field | Value |
|-------|-------|
| **Number** | `-420539` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Good till session orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_GTS` (-420541)

| Field | Value |
|-------|-------|
| **Number** | `-420541` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Good till session orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_GTS` (-420543)

| Field | Value |
|-------|-------|
| **Number** | `-420543` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Good till session orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_IMS_RESTRICTIONS` (-420551)

| Field | Value |
|-------|-------|
| **Number** | `-420551` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order cannot be updated/canceled during IMS session state |
| **Explanation** | When session state is Issuing IMS or Buy Back IMS it is not allowed for participants other than the issuer to update or cancel orders. |
| **User Action** | No action |

### `ME_MATCH_MAX_QUOTE_ITEMS` (-420585)

| Field | Value |
|-------|-------|
| **Number** | `-420585` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Number of items more than maximum allowed |
| **Explanation** | The number of items in the quote message is more than the maximum allowed. |
| **User Action** | Reduce the number of items. |

### `ME_MATCH_INS_VWAP` (-420603)

| Field | Value |
|-------|-------|
| **Number** | `-420603` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Vwap orders are not allowed for this instrument |
| **Explanation** | Vwap orders are not allowed for this instrument. |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_VWAP` (-420605)

| Field | Value |
|-------|-------|
| **Number** | `-420605` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Vwap orders are not allowed in this session state |
| **Explanation** | Vwap orders are not allowed in this session state. |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_VWAP` (-420607)

| Field | Value |
|-------|-------|
| **Number** | `-420607` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Vwap orders are not allowed for this actor |
| **Explanation** | Vwap orders are not allowed for this actor. |
| **User Action** | Respecify the order. |

### `ME_MATCH_MIN_QUOTE_ITEMS` (-420613)

| Field | Value |
|-------|-------|
| **Number** | `-420613` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Number of items less than minimum allowed |
| **Explanation** | The number of items in the quote message is less than the minimum allowed. |
| **User Action** | Increase the number of items. |

### `ME_MATCH_MUST_BE_ACTIVE` (-420615)

| Field | Value |
|-------|-------|
| **Number** | `-420615` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order must be specified as active. |
| **Explanation** | The order can only be entered as an active order. |
| **User Action** | Respecify the order. |

### `ME_MATCH_ATR_PST_ORDER` (-420621)

| Field | Value |
|-------|-------|
| **Number** | `-420621` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price stabilization not allowed with these order attributes |
| **Explanation** | Price stabilization not allowed with these order attributes. Typically this is due to the specified TimeValidity or OrderType. |
| **User Action** | Respecify the order. |

### `ME_MATCH_IMBALANCE_IOC` (-420623)

| Field | Value |
|-------|-------|
| **Number** | `-420623` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Imbalance orders must be of type Immediate or Cancel. |
| **User Action** | Respecify the order. |

### `ME_MATCH_FOK_NOT_ALLOWD` (-420625)

| Field | Value |
|-------|-------|
| **Number** | `-420625` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Entered order type is not allowed to be FoK. |
| **User Action** | Respecify the order. |

### `ME_MATCH_MM_PROTECTION_NOT_TRIGGERED` (-420626)

| Field | Value |
|-------|-------|
| **Number** | `-420626` |
| **Subsystem** | ME |
| **Severity** | 🔴 ERROR |
| **Message** | Market Maker Protection has not been triggered. |
| **Explanation** | Market Maker Protection has not been triggered and hence resetting MM protection is not allowed. |
| **User Action** | No action |

### `ME_MATCH_UPDATE_TO_FOK_IOC` (-420627)

| Field | Value |
|-------|-------|
| **Number** | `-420627` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | It is not allowed to update this order to IoC or FoK. |
| **User Action** | Respecify the update order or cancel the existing order and enter a new order. |

### `ME_MATCH_FIXED_PRICE_SESSION` (-420629)

| Field | Value |
|-------|-------|
| **Number** | `-420629` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Given price is not allowed in session with fixed price. |
| **Explanation** | Orders must be entered at the specified fixed price. If no price is specified then all orders are rejected |
| **User Action** | Correct the price. |

### `ME_MATCH_INV_QUOTE_ITEMS` (-420631)

| Field | Value |
|-------|-------|
| **Number** | `-420631` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The number of items is invalid |
| **Explanation** | The number of items in the quote message is invalid. |
| **User Action** | Correct the number of items. |

### `ME_MATCH_ORDER_LOCKED` (-420633)

| Field | Value |
|-------|-------|
| **Number** | `-420633` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order cannot be updated due to being locked. |
| **Explanation** | Orders that are locked due to being sent to the away market cannot be updated |
| **User Action** | Unlock the order before updating |

### `ME_MATCH_ORDER_NOT_LOCKED` (-420635)

| Field | Value |
|-------|-------|
| **Number** | `-420635` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order cannot be unlocked due to not being locked. |
| **Explanation** | Orders that are not locked cannot be unlocked |
| **User Action** | No action |

### `ME_MATCH_PEG_REF_PRICE_MISSING` (-420637)

| Field | Value |
|-------|-------|
| **Number** | `-420637` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | No reference price available for pegged order. |
| **Explanation** | The order book currently has no price available for the pegged order to peg against. |
| **User Action** | A reference price must be set first |

### `ME_MATCH_PEG_TRIGGER_ORDER` (-420639)

| Field | Value |
|-------|-------|
| **Number** | `-420639` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Pegged order not allowed to be a trigger order. |
| **Explanation** | Pegged orders are not allowed to also be trigger orders. |
| **User Action** | Respecify the order |

### `ME_MATCH_INS_TRACKING` (-420641)

| Field | Value |
|-------|-------|
| **Number** | `-420641` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Tracking orders are not allowed for this instrument. |
| **User Action** | Respecify the order |

### `ME_MATCH_SST_TRACKING` (-420643)

| Field | Value |
|-------|-------|
| **Number** | `-420643` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Tracking orders are not allowed in this session state. |
| **User Action** | Respecify the order |

### `ME_MATCH_USR_TRACKING` (-420645)

| Field | Value |
|-------|-------|
| **Number** | `-420645` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Tracking orders are not allowed for this actor. |
| **User Action** | Respecify the order |

### `ME_MATCH_INV_TRACKED_ORDER_BOOK` (-420647)

| Field | Value |
|-------|-------|
| **Number** | `-420647` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Order book either doesn't exist or cannot be tracked by this tracking order. |
| **Explanation** | Tracking orders can only track order books that have the same currency, ranking and decimals configuration as the order's order book. |
| **User Action** | Respecify the order |

### `ME_MATCH_SH_ZERO_TICK` (-420649)

| Field | Value |
|-------|-------|
| **Number** | `-420649` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price must be greater than or equal to Last Match Price |
| **Explanation** | The short sell order's price must be greater than or equal to Last Match Price. |
| **User Action** | Respecify the order. |

### `ME_MATCH_SH_PLUS_TICK` (-420651)

| Field | Value |
|-------|-------|
| **Number** | `-420651` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price must be greater than Last Match Price |
| **Explanation** | The short sell order's price must be greater than Last Match Price. |
| **User Action** | Respecify the order. |

### `ME_MATCH_SH_ILL_ORDER_TYPE` (-420653)

| Field | Value |
|-------|-------|
| **Number** | `-420653` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Illegal order type for short sell order. |
| **Explanation** | The short sell order must not be a market order. |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_MARGIN` (-420655)

| Field | Value |
|-------|-------|
| **Number** | `-420655` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Margin orders are not allowed for this instrument. |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_MARGIN` (-420657)

| Field | Value |
|-------|-------|
| **Number** | `-420657` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Margin orders are not allowed in this session state. |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_MARGIN` (-420659)

| Field | Value |
|-------|-------|
| **Number** | `-420659` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Margin orders are not allowed for this actor. |
| **User Action** | Respecify the order. |

### `ME_MATCH_MARGIN_SELL` (-420661)

| Field | Value |
|-------|-------|
| **Number** | `-420661` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Margin order cannot be a sell order. |
| **User Action** | Respecify the order. |

### `ME_MATCH_MARGIN_INV_EXCH_ORDER_TYPE` (-420663)

| Field | Value |
|-------|-------|
| **Number** | `-420663` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Margin exchange order type cannot be combined with order exchange order types. |
| **User Action** | Respecify the order. |

### `ME_MATCH_COMBO_MASS_QUOTE` (-420665)

| Field | Value |
|-------|-------|
| **Number** | `-420665` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Combination order books are not allowed in mass quotes |
| **User Action** | Do not send in mass quotes with combination order books. |

### `ME_MATCH_ACCOUNT_SUSPENDED` (-420667)

| Field | Value |
|-------|-------|
| **Number** | `-420667` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The account is suspended. |
| **Explanation** | The supplied account is suspended. |
| **User Action** | Correct your account. |

### `ME_MATCH_INVESTOR_SUSPENDED` (-420669)

| Field | Value |
|-------|-------|
| **Number** | `-420669` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The investor is suspended. |
| **Explanation** | The investor owning the supplied account is suspended. |
| **User Action** | Correct your account. |

### `ME_MATCH_ORDER_MIN_VAL` (-420671)

| Field | Value |
|-------|-------|
| **Number** | `-420671` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The order's value is lower than the minimum allowed value. |
| **User Action** | Specify higher order value |

### `ME_MATCH_ORDER_MAX_VAL` (-420673)

| Field | Value |
|-------|-------|
| **Number** | `-420673` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The order's value is higher than the maximum allowed value. |
| **User Action** | Specify lower order value |

### `ME_MATCH_ORDER_VAL_REF_PRICE_MISSING` (-420675)

| Field | Value |
|-------|-------|
| **Number** | `-420675` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | No reference price available to use for validating the unpriced order's value. |
| **Explanation** | A reference price is required to validate the value of an unpriced order. |
| **User Action** | A reference price must be set first |

### `ME_MATCH_INS_MINIMUM_FILL` (-420677)

| Field | Value |
|-------|-------|
| **Number** | `-420677` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum fill orders are not allowed for this instrument. |
| **User Action** | Respecify the order |

### `ME_MATCH_SST_MINIMUM_FILL` (-420679)

| Field | Value |
|-------|-------|
| **Number** | `-420679` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum fill orders are not allowed in this session state. |
| **User Action** | Respecify the order |

### `ME_MATCH_USR_MINIMUM_FILL` (-420681)

| Field | Value |
|-------|-------|
| **Number** | `-420681` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum fill orders are not allowed for this actor. |
| **User Action** | Respecify the order |

### `ME_MATCH_ILL_QUANTITY_RESTRICTION_CHANGE` (-420683)

| Field | Value |
|-------|-------|
| **Number** | `-420683` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The quantity restriction of an order is not allowed to change. |
| **Explanation** | An AoN order cannot be changed into a Min Fill order, or vice versa. Nor can an exisitng order without quantity restrictions be changed to an order with quantity restrictions. |
| **User Action** | Enter a valid quantity. |

### `ME_MATCH_RESERVE_ORDER_QUANTITY_RESTRICTION` (-420685)

| Field | Value |
|-------|-------|
| **Number** | `-420685` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | An AoN or Minimum fill order is not allowed to be a reserve order. |
| **User Action** | Enter a valid quantity. |

### `ME_MATCH_ILL_RESERVE_CHANGE` (-420687)

| Field | Value |
|-------|-------|
| **Number** | `-420687` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The reserve condition of an order is not allowed to change. |
| **User Action** | Enter a valid quantity. |

### `ME_MATCH_QUANTITY_RESTRICTION_NOT_ALLOWED` (-420689)

| Field | Value |
|-------|-------|
| **Number** | `-420689` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Entered order type is not allowed to have quantity restriction. |
| **User Action** | Respecify the order. |

### `ME_MATCH_IMB_TRIGGER_ORDER` (-420691)

| Field | Value |
|-------|-------|
| **Number** | `-420691` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Imbalance order is not allowed to be a trigger order. |
| **Explanation** | Imbalance orders are not allowed to also be trigger orders. |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_WHOLE_OR_NONE` (-420693)

| Field | Value |
|-------|-------|
| **Number** | `-420693` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Whole or none orders are not allowed for this instrument |
| **User Action** | Respecify the order. |

### `ME_MATCH_INS_MINIMUM_EXECUTION` (-420695)

| Field | Value |
|-------|-------|
| **Number** | `-420695` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum execution orders are not allowed for this instrument. |
| **User Action** | Respecify the order |

### `ME_MATCH_SST_WHOLE_OR_NONE` (-420697)

| Field | Value |
|-------|-------|
| **Number** | `-420697` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Whole or none orders are not allowed in this session state |
| **User Action** | Respecify the order. |

### `ME_MATCH_SST_MINIMUM_EXECUTION` (-420699)

| Field | Value |
|-------|-------|
| **Number** | `-420699` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum execution orders are not allowed in this session state. |
| **User Action** | Respecify the order |

### `ME_MATCH_USR_WHOLE_OR_NONE` (-420701)

| Field | Value |
|-------|-------|
| **Number** | `-420701` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Whole or none orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_USR_MINIMUM_EXECUTION` (-420703)

| Field | Value |
|-------|-------|
| **Number** | `-420703` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Minimum execution orders are not allowed for this actor. |
| **User Action** | Respecify the order |

### `ME_MATCH_INVALID_ASSET` (-420705)

| Field | Value |
|-------|-------|
| **Number** | `-420705` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid asset. |
| **Explanation** | The specified asset is invalid. |
| **User Action** | Specify a valid asset |

### `ME_MATCH_INVALID_INVESTOR` (-420707)

| Field | Value |
|-------|-------|
| **Number** | `-420707` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid investor. |
| **Explanation** | The specified investor is invalid. |
| **User Action** | Specify a valid investor |

### `ME_MATCH_POSITION_LIMIT_BREACH` (-420709)

| Field | Value |
|-------|-------|
| **Number** | `-420709` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Requested position exceeds the configured limit. |
| **Explanation** | The order quantity exceeds the remaining quantity allowed for the Account/Investor position. |
| **User Action** | Increase the allowed position or restate the order with a valid quantity. |

### `ME_MATCH_USR_POST_ONLY` (-420715)

| Field | Value |
|-------|-------|
| **Number** | `-420715` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Post Only orders are not allowed for this actor |
| **User Action** | Respecify the order. |

### `ME_MATCH_POST_ONLY` (-420717)

| Field | Value |
|-------|-------|
| **Number** | `-420717` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Not possible to place the order as post only. |
| **User Action** | Respecify the order. |

### `ME_MATCH_POST_ONLY_ADJUST` (-420719)

| Field | Value |
|-------|-------|
| **Number** | `-420719` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Not possible to adjust the post only order. |
| **User Action** | Respecify the order. |

### `ME_OB_NOT_SERIES` (-425007)

| Field | Value |
|-------|-------|
| **Number** | `-425007` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The given series has not been defined (in this instance). |
| **Explanation** | The series was not found in the order book. Order-book data may be corrupt. |
| **User Action** | Please send a description of the circumstances together with log files to OMT. |

### `ME_OB_NOT_CUSTOMER` (-425017)

| Field | Value |
|-------|-------|
| **Number** | `-425017` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | The given customer was not found or is not in a valid trade state. |
| **Explanation** | The customer is not allowed to trade or is missing |
| **User Action** | Add the customer in CDB if it is missing. 7 ME Error Messages |

### `ME_MATCH_IDX_INV_EXCHANGE_INFO` (-421001)

| Field | Value |
|-------|-------|
| **Number** | `-421001` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Wrong contents of exchange info |
| **Explanation** | The contents of the exchange info field is wrong |
| **User Action** | You need to change the exchange info field. The contents of the fields for ordersource or settlement method is not in the correct format. |

### `ME_MATCH_IDX_INV_PRICE_PST` (-421003)

| Field | Value |
|-------|-------|
| **Number** | `-421003` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price stabilization is not allowed at this price. |
| **Explanation** | Price entered for the Price stabilization order is invalid. |
| **User Action** | Enter valid price value. |

### `ME_MATCH_IDX_SH_ONE_LESS_TICK_ZERO` (-421005)

| Field | Value |
|-------|-------|
| **Number** | `-421005` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price must be greater than or equal to one tick below Last Match Price |
| **Explanation** | The short sell order's price must be greater than or equal to one price tick below the Last Match Price in the order book. |
| **User Action** | Respecify the order. |

### `ME_MATCH_IDX_SH_ONE_LESS_TICK_PLUS` (-421007)

| Field | Value |
|-------|-------|
| **Number** | `-421007` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Price must be greater than one tick below Last Match Price ME_MATCH_IDX_SH_ONE_LESS_TICK_PLUS |
| **Explanation** | The short sell order's price must be greater than one price tick below the Last Match Price in the order book. |
| **User Action** | Respecify the order. ME_MATCH_IDX_MAX_ORDER_QUANTITY_PERCENTAGE_OF_TRADABLE_ |

### `QUANTITY` (-421009)

| Field | Value |
|-------|-------|
| **Number** | `-421009` |
| **Subsystem** | ME |
| **Severity** | 🟡 WARNING |
| **Message** | Quantity exceeds the max order quantity percentage of tradable quantity. |
| **User Action** | Respecify the order/quote. |

---

## RDM Error Messages

### `MME_RDM_TMC_ILLEGAL_RATIO` (-49900009)

| Field | Value |
|-------|-------|
| **Number** | `-49900009` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Leg ratio not defined or the GCD (all ratios) is not one. |
| **Explanation** | Ratio in all legs or GCD for all legs must be one. |
| **User Action** | Send correct ratio in legs in TMC creation request. |

### `MME_RDM_TMC_DUPLICATE_LEGS` (-49900011)

| Field | Value |
|-------|-------|
| **Number** | `-49900011` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Same leg exists more than once in Tailor-made combination. |
| **Explanation** | One leg can occur only once in the Tailor-made combination. |
| **User Action** | Send unique legs in Tailor-made combination. |

### `MME_RDM_TMC_INV_LEG_ORDERBOOK` (-49900013)

| Field | Value |
|-------|-------|
| **Number** | `-49900013` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid orderbook in leg. |
| **Explanation** | Invalid orderbook sent in the Tailor-made combination. |
| **User Action** | Send valid orderbooks in legs. |

### `MME_RDM_TMC_STRATEGY_NOT_FOUND` (-49900015)

| Field | Value |
|-------|-------|
| **Number** | `-49900015` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Valid strategy not found. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_USER_NOT_ALLOW` (-49900017)

| Field | Value |
|-------|-------|
| **Number** | `-49900017` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | User not allowed to create TMC. |
| **Explanation** | User not allowed to send TMC Message or have Buy/Sell or have rights in legs market segment. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_MAX_RATIO` (-49900019)

| Field | Value |
|-------|-------|
| **Number** | `-49900019` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Leg ratio is greater than configured max ratio. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_MAX_RATIO_DIFFERENCE` (-49900021)

| Field | Value |
|-------|-------|
| **Number** | `-49900021` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Min and Max ratio difference is greater than configured max ratio difference. |
| **Explanation** | Leg's Min and Max ratio difference is greater than configured max ratio difference. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_PARTITION` (-49900023)

| Field | Value |
|-------|-------|
| **Number** | `-49900023` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid partition. |
| **Explanation** | All Legs must be traded within the same partition as the combination order book. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_LEG_MARKET` (-49900025)

| Field | Value |
|-------|-------|
| **Number** | `-49900025` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid market. |
| **Explanation** | All Legs must have same market. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_LEG_UNDERLYING` (-49900027)

| Field | Value |
|-------|-------|
| **Number** | `-49900027` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid underlying asset. |
| **Explanation** | All Legs must have same underlying asset. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_CONT_VAL_FACT` (-49900029)

| Field | Value |
|-------|-------|
| **Number** | `-49900029` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid contract value factor. |
| **Explanation** | All Legs must have same contract value factor. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_PRICE_UNIT` (-49900031)

| Field | Value |
|-------|-------|
| **Number** | `-49900031` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid price unit. |
| **Explanation** | All Legs must be traded in the same price unit as the combination order book. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_ASSET_CURRENCY` (-49900033)

| Field | Value |
|-------|-------|
| **Number** | `-49900033` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid asset currency. |
| **Explanation** | All Legs must be traded using same currency and currency unit as the combination order book. |
| **User Action** | Contact exchange. |

### `MME_RDM_TMC_INV_PRICE_TICK_DEC` (-49900035)

| Field | Value |
|-------|-------|
| **Number** | `-49900035` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid price tick decimal. |
| **Explanation** | All Legs must be traded in the same number of price decimals as the combination order book. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_FAILED` (-49900037)

| Field | Value |
|-------|-------|
| **Number** | `-49900037` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Failed to generate Repo order book. |
| **Explanation** | failed to generate Repo order book due to unexpected reason. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_CLEARING_HOLIDAY` (-49900039)

| Field | Value |
|-------|-------|
| **Number** | `-49900039` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Return date is a clear holiday. |
| **Explanation** | Return date cannot be a clearing holiday. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_CONTRACT_LENGTH` (-49900041)

| Field | Value |
|-------|-------|
| **Number** | `-49900041` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid contract length. |
| **Explanation** | The specified contract length must be on or within the allowed range in the repo rules. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_TYPE` (-49900043)

| Field | Value |
|-------|-------|
| **Number** | `-49900043` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid repo type. |
| **Explanation** | The specified asset must have repo rules configured for the repo type specified by the user. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_RECALL_ALLOWED` (-49900045)

| Field | Value |
|-------|-------|
| **Number** | `-49900045` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid recalled allowed. |
| **Explanation** | The “Recall Allowed” flag can only be set to true of the repo type is “Securities Lending”. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_INV_REFERENCE_PRICE` (-49900047)

| Field | Value |
|-------|-------|
| **Number** | `-49900047` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid reference price. |
| **Explanation** | Specified order book does not have a valid reference price. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_ASSET_NOT_ALLOWED` (-49900049)

| Field | Value |
|-------|-------|
| **Number** | `-49900049` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Either asset is not connected to assetGroup or repo type is not allowed for market segment. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_INV_ASSET` (-49900051)

| Field | Value |
|-------|-------|
| **Number** | `-49900051` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Given asset is either not exits, not an equity asset or not active. |
| **Explanation** | Repo order book is only allowed for valid equity order books. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_INV_EQUITY_ORDERBOOK` (-49900053)

| Field | Value |
|-------|-------|
| **Number** | `-49900053` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid equity order book. |
| **Explanation** | Repo order book is only allowed for valid equity order books. |
| **User Action** | Contact exchange. |

### `MME_RDM_REPO_NO_ASSET_REPO` (-49900055)

| Field | Value |
|-------|-------|
| **Number** | `-49900055` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | Asset repo is not configured for equity asset. |
| **User Action** | Contact exchange. |

### `MME_RDM_MAX_REPO_ORDER_BOOKS_REACHED` (-49900057)

| Field | Value |
|-------|-------|
| **Number** | `-49900057` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | The maximum number of allowed active repo order books has been reached. |
| **User Action** | Contact exchange. |

### `MME_RDM_NOT_INITIALIZED` (-49900059)

| Field | Value |
|-------|-------|
| **Number** | `-49900059` |
| **Subsystem** | RDM |
| **Severity** | 🟡 WARNING |
| **Message** | RDM OBG has not yet been initialized. |
| **Explanation** | RDM OBG has not yet been initialized when the transaction was received. |
| **User Action** | None. 8 PH Error Messages |

---

## PH Error Messages

### `PH_NOT_INITIALIZED` (-500001)

| Field | Value |
|-------|-------|
| **Number** | `-500001` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Price Handler has not yet been initialized. |
| **Explanation** | Price Handler has not yet been initialized when the transaction was received. |

### `PH_INV_ORDERBOOK` (-500003)

| Field | Value |
|-------|-------|
| **Number** | `-500003` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Order Book. |
| **Explanation** | The specified order book can not be found. |

### `PH_INV_ORD_BOOK_NOT_INDEX` (-500005)

| Field | Value |
|-------|-------|
| **Number** | `-500005` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Order Book, not an index. |
| **Explanation** | The specified order book is not an index. |

### `PH_INV_REF_PRICE_SOURCE` (-500007)

| Field | Value |
|-------|-------|
| **Number** | `-500007` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Reference Price Source, not supported. |
| **Explanation** | The specified Reference Price Source is not supported. |

### `PH_INV_ORD_BOOK_TRADABLE` (-500009)

| Field | Value |
|-------|-------|
| **Number** | `-500009` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Order Book, tradable. |
| **Explanation** | The specified order book is tradable. |

### `PH_NO_REFERENCE_PRICE` (-500011)

| Field | Value |
|-------|-------|
| **Number** | `-500011` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Reference price not exists to calculate closing price. |

### `PH_CLOSING_PRICE_CALCULATION` (-500013)

| Field | Value |
|-------|-------|
| **Number** | `-500013` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Closing price calculation not configured for market segment. |
| **User Action** | Contact exchange. |

### `PH_INV_ACTOR` (-500015)

| Field | Value |
|-------|-------|
| **Number** | `-500015` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Actor. |
| **Explanation** | Actor is invalid. |

### `PH_INV_ORD_BOOK_NOT_ACTIVE` (-500017)

| Field | Value |
|-------|-------|
| **Number** | `-500017` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Order Book not active. |
| **Explanation** | The specified order book is not active. |

### `PH_ORD_BOOK_NOT_MODIFIED` (-500019)

| Field | Value |
|-------|-------|
| **Number** | `-500019` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Order Book not modified. |
| **Explanation** | The specified change did not modify the order book |

### `PH_INV_OVERRIDE_CIRCUIT_BREAKER` (-500021)

| Field | Value |
|-------|-------|
| **Number** | `-500021` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid OverrideIndexCircuitBreaker message. |
| **Explanation** | The received OverrideIndexCircuitBreaker messages is invalid. |

### `PH_REF_PRICE_OUTSIDE_RANGE` (-500023)

| Field | Value |
|-------|-------|
| **Number** | `-500023` |
| **Subsystem** | PH |
| **Severity** | 🟡 WARNING |
| **Message** | Reference price is outside of the valid price tick range. |
| **Explanation** | The received reference price is invalid. |
| **User Action** | None. 9 PTR Error Messages |

---

## PTR Error Messages

### `RM_BLOCKED` (-1300002)

| Field | Value |
|-------|-------|
| **Number** | `-1300002` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | The operation was denied due to an administrative block. |
| **Explanation** | An entity or entities were put into administrative block, and the operation was denied. |

### `RM_BREACHED` (-1300004)

| Field | Value |
|-------|-------|
| **Number** | `-1300004` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | The operation was denied due to a breach. |
| **Explanation** | A risk check(s) were breacehed, and the operation was denied. |

### `RM_ERROR` (-1300006)

| Field | Value |
|-------|-------|
| **Number** | `-1300006` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | The operation was denied due to an internal error. |
| **Explanation** | An internal error occurred, and the operation was denied. |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_AGGREGATED` (-1300008)

| Field | Value |
|-------|-------|
| **Number** | `-1300008` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300010)

| Field | Value |
|-------|-------|
| **Number** | `-1300010` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VALUE_AGGREGATED` (-1300012)

| Field | Value |
|-------|-------|
| **Number** | `-1300012` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VOLUME_AGGREGATED` (-1300014)

| Field | Value |
|-------|-------|
| **Number** | `-1300014` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_QUANTITY_AGGREGATED` (-1300016)

| Field | Value |
|-------|-------|
| **Number** | `-1300016` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_TRANSIENT` (-1300018)

| Field | Value |
|-------|-------|
| **Number** | `-1300018` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300020)

| Field | Value |
|-------|-------|
| **Number** | `-1300020` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VALUE_TRANSIENT` (-1300022)

| Field | Value |
|-------|-------|
| **Number** | `-1300022` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VOLUME_TRANSIENT` (-1300024)

| Field | Value |
|-------|-------|
| **Number** | `-1300024` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_QUANTITY_TRANSIENT` (-1300026)

| Field | Value |
|-------|-------|
| **Number** | `-1300026` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_OVER_TIME` (-1300028)

| Field | Value |
|-------|-------|
| **Number** | `-1300028` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300030)

| Field | Value |
|-------|-------|
| **Number** | `-1300030` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VALUE_OVER_TIME` (-1300032)

| Field | Value |
|-------|-------|
| **Number** | `-1300032` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_VOLUME_OVER_TIME` (-1300034)

| Field | Value |
|-------|-------|
| **Number** | `-1300034` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_QUANTITY_OVER_TIME` (-1300036)

| Field | Value |
|-------|-------|
| **Number** | `-1300036` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_AGGREGATED` (-1300038)

| Field | Value |
|-------|-------|
| **Number** | `-1300038` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300040)

| Field | Value |
|-------|-------|
| **Number** | `-1300040` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VALUE_AGGREGATED` (-1300042)

| Field | Value |
|-------|-------|
| **Number** | `-1300042` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VOLUME_AGGREGATED` (-1300044)

| Field | Value |
|-------|-------|
| **Number** | `-1300044` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_QUANTITY_AGGREGATED` (-1300046)

| Field | Value |
|-------|-------|
| **Number** | `-1300046` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_TRANSIENT` (-1300048)

| Field | Value |
|-------|-------|
| **Number** | `-1300048` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300050)

| Field | Value |
|-------|-------|
| **Number** | `-1300050` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VALUE_TRANSIENT` (-1300052)

| Field | Value |
|-------|-------|
| **Number** | `-1300052` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VOLUME_TRANSIENT` (-1300054)

| Field | Value |
|-------|-------|
| **Number** | `-1300054` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_QUANTITY_TRANSIENT` (-1300056)

| Field | Value |
|-------|-------|
| **Number** | `-1300056` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_OVER_TIME` (-1300058)

| Field | Value |
|-------|-------|
| **Number** | `-1300058` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300060)

| Field | Value |
|-------|-------|
| **Number** | `-1300060` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VALUE_OVER_TIME` (-1300062)

| Field | Value |
|-------|-------|
| **Number** | `-1300062` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_VOLUME_OVER_TIME` (-1300064)

| Field | Value |
|-------|-------|
| **Number** | `-1300064` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_QUANTITY_OVER_TIME` (-1300066)

| Field | Value |
|-------|-------|
| **Number** | `-1300066` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_AGGREGATED` (-1300068)

| Field | Value |
|-------|-------|
| **Number** | `-1300068` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_COUNT_AGGREGATED` (-1300070)

| Field | Value |
|-------|-------|
| **Number** | `-1300070` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VALUE_AGGREGATED` (-1300072)

| Field | Value |
|-------|-------|
| **Number** | `-1300072` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VOLUME_AGGREGATED` (-1300074)

| Field | Value |
|-------|-------|
| **Number** | `-1300074` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_QUANTITY_AGGREGATED` (-1300076)

| Field | Value |
|-------|-------|
| **Number** | `-1300076` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_TRANSIENT` (-1300078)

| Field | Value |
|-------|-------|
| **Number** | `-1300078` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_COUNT_TRANSIENT` (-1300080)

| Field | Value |
|-------|-------|
| **Number** | `-1300080` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VALUE_TRANSIENT` (-1300082)

| Field | Value |
|-------|-------|
| **Number** | `-1300082` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VOLUME_TRANSIENT` (-1300084)

| Field | Value |
|-------|-------|
| **Number** | `-1300084` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_QUANTITY_TRANSIENT` (-1300086)

| Field | Value |
|-------|-------|
| **Number** | `-1300086` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_OVER_TIME` (-1300088)

| Field | Value |
|-------|-------|
| **Number** | `-1300088` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_EXPOSURE_COUNT_OVER_TIME` (-1300090)

| Field | Value |
|-------|-------|
| **Number** | `-1300090` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VALUE_OVER_TIME` (-1300092)

| Field | Value |
|-------|-------|
| **Number** | `-1300092` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_VOLUME_OVER_TIME` (-1300094)

| Field | Value |
|-------|-------|
| **Number** | `-1300094` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_QUANTITY_OVER_TIME` (-1300096)

| Field | Value |
|-------|-------|
| **Number** | `-1300096` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_AGGREGATED` (-1300098)

| Field | Value |
|-------|-------|
| **Number** | `-1300098` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_COUNT_AGGREGATED` (-1300100)

| Field | Value |
|-------|-------|
| **Number** | `-1300100` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VALUE_AGGREGATED` (-1300102)

| Field | Value |
|-------|-------|
| **Number** | `-1300102` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VOLUME_AGGREGATED` (-1300104)

| Field | Value |
|-------|-------|
| **Number** | `-1300104` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_QUANTITY_AGGREGATED` (-1300106)

| Field | Value |
|-------|-------|
| **Number** | `-1300106` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_TRANSIENT` (-1300108)

| Field | Value |
|-------|-------|
| **Number** | `-1300108` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_COUNT_TRANSIENT` (-1300110)

| Field | Value |
|-------|-------|
| **Number** | `-1300110` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VALUE_TRANSIENT` (-1300112)

| Field | Value |
|-------|-------|
| **Number** | `-1300112` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VOLUME_TRANSIENT` (-1300114)

| Field | Value |
|-------|-------|
| **Number** | `-1300114` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_QUANTITY_TRANSIENT` (-1300116)

| Field | Value |
|-------|-------|
| **Number** | `-1300116` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_OVER_TIME` (-1300118)

| Field | Value |
|-------|-------|
| **Number** | `-1300118` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_EXPOSURE_COUNT_OVER_TIME` (-1300120)

| Field | Value |
|-------|-------|
| **Number** | `-1300120` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VALUE_OVER_TIME` (-1300122)

| Field | Value |
|-------|-------|
| **Number** | `-1300122` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_VOLUME_OVER_TIME` (-1300124)

| Field | Value |
|-------|-------|
| **Number** | `-1300124` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_QUANTITY_OVER_TIME` (-1300126)

| Field | Value |
|-------|-------|
| **Number** | `-1300126` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_AGGREGATED` (-1300128)

| Field | Value |
|-------|-------|
| **Number** | `-1300128` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_COUNT_AGGREGATED` (-1300130)

| Field | Value |
|-------|-------|
| **Number** | `-1300130` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VALUE_AGGREGATED` (-1300132)

| Field | Value |
|-------|-------|
| **Number** | `-1300132` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VOLUME_AGGREGATED` (-1300134)

| Field | Value |
|-------|-------|
| **Number** | `-1300134` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_QUANTITY_AGGREGATED` (-1300136)

| Field | Value |
|-------|-------|
| **Number** | `-1300136` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_TRANSIENT` (-1300138)

| Field | Value |
|-------|-------|
| **Number** | `-1300138` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_COUNT_TRANSIENT` (-1300140)

| Field | Value |
|-------|-------|
| **Number** | `-1300140` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VALUE_TRANSIENT` (-1300142)

| Field | Value |
|-------|-------|
| **Number** | `-1300142` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VOLUME_TRANSIENT` (-1300144)

| Field | Value |
|-------|-------|
| **Number** | `-1300144` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_QUANTITY_TRANSIENT` (-1300146)

| Field | Value |
|-------|-------|
| **Number** | `-1300146` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_OVER_TIME` (-1300148)

| Field | Value |
|-------|-------|
| **Number** | `-1300148` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_EXPOSURE_COUNT_OVER_TIME` (-1300150)

| Field | Value |
|-------|-------|
| **Number** | `-1300150` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VALUE_OVER_TIME` (-1300152)

| Field | Value |
|-------|-------|
| **Number** | `-1300152` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_VOLUME_OVER_TIME` (-1300154)

| Field | Value |
|-------|-------|
| **Number** | `-1300154` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_QUANTITY_OVER_TIME` (-1300156)

| Field | Value |
|-------|-------|
| **Number** | `-1300156` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_AGGREGATED` (-1300158)

| Field | Value |
|-------|-------|
| **Number** | `-1300158` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_COUNT_AGGREGATED` (-1300160)

| Field | Value |
|-------|-------|
| **Number** | `-1300160` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VALUE_AGGREGATED` (-1300162)

| Field | Value |
|-------|-------|
| **Number** | `-1300162` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VOLUME_AGGREGATED` (-1300164)

| Field | Value |
|-------|-------|
| **Number** | `-1300164` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_QUANTITY_AGGREGATED` (-1300166)

| Field | Value |
|-------|-------|
| **Number** | `-1300166` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_TRANSIENT` (-1300168)

| Field | Value |
|-------|-------|
| **Number** | `-1300168` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_COUNT_TRANSIENT` (-1300170)

| Field | Value |
|-------|-------|
| **Number** | `-1300170` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VALUE_TRANSIENT` (-1300172)

| Field | Value |
|-------|-------|
| **Number** | `-1300172` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VOLUME_TRANSIENT` (-1300174)

| Field | Value |
|-------|-------|
| **Number** | `-1300174` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_QUANTITY_TRANSIENT` (-1300176)

| Field | Value |
|-------|-------|
| **Number** | `-1300176` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_OVER_TIME` (-1300178)

| Field | Value |
|-------|-------|
| **Number** | `-1300178` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_EXPOSURE_COUNT_OVER_TIME` (-1300180)

| Field | Value |
|-------|-------|
| **Number** | `-1300180` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VALUE_OVER_TIME` (-1300182)

| Field | Value |
|-------|-------|
| **Number** | `-1300182` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_VOLUME_OVER_TIME` (-1300184)

| Field | Value |
|-------|-------|
| **Number** | `-1300184` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_QUANTITY_OVER_TIME` (-1300186)

| Field | Value |
|-------|-------|
| **Number** | `-1300186` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_AGGREGATED` (-1300188)

| Field | Value |
|-------|-------|
| **Number** | `-1300188` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300190)

| Field | Value |
|-------|-------|
| **Number** | `-1300190` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VALUE_AGGREGATED` (-1300192)

| Field | Value |
|-------|-------|
| **Number** | `-1300192` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VOLUME_AGGREGATED` (-1300194)

| Field | Value |
|-------|-------|
| **Number** | `-1300194` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_QUANTITY_AGGREGATED` (-1300196)

| Field | Value |
|-------|-------|
| **Number** | `-1300196` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_TRANSIENT` (-1300198)

| Field | Value |
|-------|-------|
| **Number** | `-1300198` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300200)

| Field | Value |
|-------|-------|
| **Number** | `-1300200` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VALUE_TRANSIENT` (-1300202)

| Field | Value |
|-------|-------|
| **Number** | `-1300202` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VOLUME_TRANSIENT` (-1300204)

| Field | Value |
|-------|-------|
| **Number** | `-1300204` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_QUANTITY_TRANSIENT` (-1300206)

| Field | Value |
|-------|-------|
| **Number** | `-1300206` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_OVER_TIME` (-1300208)

| Field | Value |
|-------|-------|
| **Number** | `-1300208` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300210)

| Field | Value |
|-------|-------|
| **Number** | `-1300210` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VALUE_OVER_TIME` (-1300212)

| Field | Value |
|-------|-------|
| **Number** | `-1300212` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_VOLUME_OVER_TIME` (-1300214)

| Field | Value |
|-------|-------|
| **Number** | `-1300214` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_QUANTITY_OVER_TIME` (-1300216)

| Field | Value |
|-------|-------|
| **Number** | `-1300216` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_AGGREGATED` (-1300218)

| Field | Value |
|-------|-------|
| **Number** | `-1300218` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_AGGREGATED` (-1300220)

| Field | Value |
|-------|-------|
| **Number** | `-1300220` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VALUE_AGGREGATED` (-1300222)

| Field | Value |
|-------|-------|
| **Number** | `-1300222` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VOLUME_AGGREGATED` (-1300224)

| Field | Value |
|-------|-------|
| **Number** | `-1300224` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_QUANTITY_AGGREGATED` (-1300226)

| Field | Value |
|-------|-------|
| **Number** | `-1300226` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_TRANSIENT` (-1300228)

| Field | Value |
|-------|-------|
| **Number** | `-1300228` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_TRANSIENT` (-1300230)

| Field | Value |
|-------|-------|
| **Number** | `-1300230` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VALUE_TRANSIENT` (-1300232)

| Field | Value |
|-------|-------|
| **Number** | `-1300232` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VOLUME_TRANSIENT` (-1300234)

| Field | Value |
|-------|-------|
| **Number** | `-1300234` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_QUANTITY_TRANSIENT` (-1300236)

| Field | Value |
|-------|-------|
| **Number** | `-1300236` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_OVER_TIME` (-1300238)

| Field | Value |
|-------|-------|
| **Number** | `-1300238` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_OVER_TIME` (-1300240)

| Field | Value |
|-------|-------|
| **Number** | `-1300240` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VALUE_OVER_TIME` (-1300242)

| Field | Value |
|-------|-------|
| **Number** | `-1300242` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_VOLUME_OVER_TIME` (-1300244)

| Field | Value |
|-------|-------|
| **Number** | `-1300244` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_QUANTITY_OVER_TIME` (-1300246)

| Field | Value |
|-------|-------|
| **Number** | `-1300246` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_AGGREGATED` (-1300248)

| Field | Value |
|-------|-------|
| **Number** | `-1300248` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_AGGREGATED` (-1300250)

| Field | Value |
|-------|-------|
| **Number** | `-1300250` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VALUE_AGGREGATED` (-1300252)

| Field | Value |
|-------|-------|
| **Number** | `-1300252` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VOLUME_AGGREGATED` (-1300254)

| Field | Value |
|-------|-------|
| **Number** | `-1300254` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_QUANTITY_AGGREGATED` (-1300256)

| Field | Value |
|-------|-------|
| **Number** | `-1300256` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_TRANSIENT` (-1300258)

| Field | Value |
|-------|-------|
| **Number** | `-1300258` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_TRANSIENT` (-1300260)

| Field | Value |
|-------|-------|
| **Number** | `-1300260` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VALUE_TRANSIENT` (-1300262)

| Field | Value |
|-------|-------|
| **Number** | `-1300262` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VOLUME_TRANSIENT` (-1300264)

| Field | Value |
|-------|-------|
| **Number** | `-1300264` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_QUANTITY_TRANSIENT` (-1300266)

| Field | Value |
|-------|-------|
| **Number** | `-1300266` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_OVER_TIME` (-1300268)

| Field | Value |
|-------|-------|
| **Number** | `-1300268` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_OVER_TIME` (-1300270)

| Field | Value |
|-------|-------|
| **Number** | `-1300270` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VALUE_OVER_TIME` (-1300272)

| Field | Value |
|-------|-------|
| **Number** | `-1300272` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_VOLUME_OVER_TIME` (-1300274)

| Field | Value |
|-------|-------|
| **Number** | `-1300274` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_QUANTITY_OVER_TIME` (-1300276)

| Field | Value |
|-------|-------|
| **Number** | `-1300276` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_AGGREGATED` (-1300278)

| Field | Value |
|-------|-------|
| **Number** | `-1300278` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_COUNT_AGGREGATED` (-1300280)

| Field | Value |
|-------|-------|
| **Number** | `-1300280` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VALUE_AGGREGATED` (-1300282)

| Field | Value |
|-------|-------|
| **Number** | `-1300282` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VOLUME_AGGREGATED` (-1300284)

| Field | Value |
|-------|-------|
| **Number** | `-1300284` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_QUANTITY_AGGREGATED` (-1300286)

| Field | Value |
|-------|-------|
| **Number** | `-1300286` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_TRANSIENT` (-1300288)

| Field | Value |
|-------|-------|
| **Number** | `-1300288` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_COUNT_TRANSIENT` (-1300290)

| Field | Value |
|-------|-------|
| **Number** | `-1300290` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VALUE_TRANSIENT` (-1300292)

| Field | Value |
|-------|-------|
| **Number** | `-1300292` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VOLUME_TRANSIENT` (-1300294)

| Field | Value |
|-------|-------|
| **Number** | `-1300294` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_QUANTITY_TRANSIENT` (-1300296)

| Field | Value |
|-------|-------|
| **Number** | `-1300296` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_OVER_TIME` (-1300298)

| Field | Value |
|-------|-------|
| **Number** | `-1300298` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_EXPOSURE_COUNT_OVER_TIME` (-1300300)

| Field | Value |
|-------|-------|
| **Number** | `-1300300` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VALUE_OVER_TIME` (-1300302)

| Field | Value |
|-------|-------|
| **Number** | `-1300302` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_VOLUME_OVER_TIME` (-1300304)

| Field | Value |
|-------|-------|
| **Number** | `-1300304` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_QUANTITY_OVER_TIME` (-1300306)

| Field | Value |
|-------|-------|
| **Number** | `-1300306` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_AGGREGATED` (-1300308)

| Field | Value |
|-------|-------|
| **Number** | `-1300308` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300310)

| Field | Value |
|-------|-------|
| **Number** | `-1300310` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VALUE_AGGREGATED` (-1300312)

| Field | Value |
|-------|-------|
| **Number** | `-1300312` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VOLUME_AGGREGATED` (-1300314)

| Field | Value |
|-------|-------|
| **Number** | `-1300314` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_QUANTITY_AGGREGATED` (-1300316)

| Field | Value |
|-------|-------|
| **Number** | `-1300316` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_TRANSIENT` (-1300318)

| Field | Value |
|-------|-------|
| **Number** | `-1300318` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300320)

| Field | Value |
|-------|-------|
| **Number** | `-1300320` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VALUE_TRANSIENT` (-1300322)

| Field | Value |
|-------|-------|
| **Number** | `-1300322` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VOLUME_TRANSIENT` (-1300324)

| Field | Value |
|-------|-------|
| **Number** | `-1300324` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_QUANTITY_TRANSIENT` (-1300326)

| Field | Value |
|-------|-------|
| **Number** | `-1300326` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_OVER_TIME` (-1300328)

| Field | Value |
|-------|-------|
| **Number** | `-1300328` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300330)

| Field | Value |
|-------|-------|
| **Number** | `-1300330` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VALUE_OVER_TIME` (-1300332)

| Field | Value |
|-------|-------|
| **Number** | `-1300332` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_VOLUME_OVER_TIME` (-1300334)

| Field | Value |
|-------|-------|
| **Number** | `-1300334` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_QUANTITY_OVER_TIME` (-1300336)

| Field | Value |
|-------|-------|
| **Number** | `-1300336` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_AGGREGATED` (-1300338)

| Field | Value |
|-------|-------|
| **Number** | `-1300338` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_COUNT_AGGREGATED` (-1300340)

| Field | Value |
|-------|-------|
| **Number** | `-1300340` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VALUE_AGGREGATED` (-1300342)

| Field | Value |
|-------|-------|
| **Number** | `-1300342` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VOLUME_AGGREGATED` (-1300344)

| Field | Value |
|-------|-------|
| **Number** | `-1300344` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_QUANTITY_AGGREGATED` (-1300346)

| Field | Value |
|-------|-------|
| **Number** | `-1300346` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_TRANSIENT` (-1300348)

| Field | Value |
|-------|-------|
| **Number** | `-1300348` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_COUNT_TRANSIENT` (-1300350)

| Field | Value |
|-------|-------|
| **Number** | `-1300350` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VALUE_TRANSIENT` (-1300352)

| Field | Value |
|-------|-------|
| **Number** | `-1300352` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VOLUME_TRANSIENT` (-1300354)

| Field | Value |
|-------|-------|
| **Number** | `-1300354` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_QUANTITY_TRANSIENT` (-1300356)

| Field | Value |
|-------|-------|
| **Number** | `-1300356` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_OVER_TIME` (-1300358)

| Field | Value |
|-------|-------|
| **Number** | `-1300358` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_EXPOSURE_COUNT_OVER_TIME` (-1300360)

| Field | Value |
|-------|-------|
| **Number** | `-1300360` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VALUE_OVER_TIME` (-1300362)

| Field | Value |
|-------|-------|
| **Number** | `-1300362` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_VOLUME_OVER_TIME` (-1300364)

| Field | Value |
|-------|-------|
| **Number** | `-1300364` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_QUANTITY_OVER_TIME` (-1300366)

| Field | Value |
|-------|-------|
| **Number** | `-1300366` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_AGGREGATED` (-1300368)

| Field | Value |
|-------|-------|
| **Number** | `-1300368` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_COUNT_AGGREGATED` (-1300370)

| Field | Value |
|-------|-------|
| **Number** | `-1300370` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VALUE_AGGREGATED` (-1300372)

| Field | Value |
|-------|-------|
| **Number** | `-1300372` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VOLUME_AGGREGATED` (-1300374)

| Field | Value |
|-------|-------|
| **Number** | `-1300374` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_QUANTITY_AGGREGATED` (-1300376)

| Field | Value |
|-------|-------|
| **Number** | `-1300376` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_TRANSIENT` (-1300378)

| Field | Value |
|-------|-------|
| **Number** | `-1300378` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_COUNT_TRANSIENT` (-1300380)

| Field | Value |
|-------|-------|
| **Number** | `-1300380` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VALUE_TRANSIENT` (-1300382)

| Field | Value |
|-------|-------|
| **Number** | `-1300382` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VOLUME_TRANSIENT` (-1300384)

| Field | Value |
|-------|-------|
| **Number** | `-1300384` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_QUANTITY_TRANSIENT` (-1300386)

| Field | Value |
|-------|-------|
| **Number** | `-1300386` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_OVER_TIME` (-1300388)

| Field | Value |
|-------|-------|
| **Number** | `-1300388` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_EXPOSURE_COUNT_OVER_TIME` (-1300390)

| Field | Value |
|-------|-------|
| **Number** | `-1300390` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VALUE_OVER_TIME` (-1300392)

| Field | Value |
|-------|-------|
| **Number** | `-1300392` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_VOLUME_OVER_TIME` (-1300394)

| Field | Value |
|-------|-------|
| **Number** | `-1300394` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_QUANTITY_OVER_TIME` (-1300396)

| Field | Value |
|-------|-------|
| **Number** | `-1300396` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_AGGREGATED` (-1300398)

| Field | Value |
|-------|-------|
| **Number** | `-1300398` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300400)

| Field | Value |
|-------|-------|
| **Number** | `-1300400` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VALUE_AGGREGATED` (-1300402)

| Field | Value |
|-------|-------|
| **Number** | `-1300402` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VOLUME_AGGREGATED` (-1300404)

| Field | Value |
|-------|-------|
| **Number** | `-1300404` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_QUANTITY_AGGREGATED` (-1300406)

| Field | Value |
|-------|-------|
| **Number** | `-1300406` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_TRANSIENT` (-1300408)

| Field | Value |
|-------|-------|
| **Number** | `-1300408` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300410)

| Field | Value |
|-------|-------|
| **Number** | `-1300410` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VALUE_TRANSIENT` (-1300412)

| Field | Value |
|-------|-------|
| **Number** | `-1300412` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VOLUME_TRANSIENT` (-1300414)

| Field | Value |
|-------|-------|
| **Number** | `-1300414` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_QUANTITY_TRANSIENT` (-1300416)

| Field | Value |
|-------|-------|
| **Number** | `-1300416` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_OVER_TIME` (-1300418)

| Field | Value |
|-------|-------|
| **Number** | `-1300418` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300420)

| Field | Value |
|-------|-------|
| **Number** | `-1300420` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VALUE_OVER_TIME` (-1300422)

| Field | Value |
|-------|-------|
| **Number** | `-1300422` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_VOLUME_OVER_TIME` (-1300424)

| Field | Value |
|-------|-------|
| **Number** | `-1300424` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_QUANTITY_OVER_TIME` (-1300426)

| Field | Value |
|-------|-------|
| **Number** | `-1300426` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_AGGREGATED` (-1300428)

| Field | Value |
|-------|-------|
| **Number** | `-1300428` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300430)

| Field | Value |
|-------|-------|
| **Number** | `-1300430` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VALUE_AGGREGATED` (-1300432)

| Field | Value |
|-------|-------|
| **Number** | `-1300432` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VOLUME_AGGREGATED` (-1300434)

| Field | Value |
|-------|-------|
| **Number** | `-1300434` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_QUANTITY_AGGREGATED` (-1300436)

| Field | Value |
|-------|-------|
| **Number** | `-1300436` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_TRANSIENT` (-1300438)

| Field | Value |
|-------|-------|
| **Number** | `-1300438` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300440)

| Field | Value |
|-------|-------|
| **Number** | `-1300440` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VALUE_TRANSIENT` (-1300442)

| Field | Value |
|-------|-------|
| **Number** | `-1300442` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VOLUME_TRANSIENT` (-1300444)

| Field | Value |
|-------|-------|
| **Number** | `-1300444` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_QUANTITY_TRANSIENT` (-1300446)

| Field | Value |
|-------|-------|
| **Number** | `-1300446` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_OVER_TIME` (-1300448)

| Field | Value |
|-------|-------|
| **Number** | `-1300448` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300450)

| Field | Value |
|-------|-------|
| **Number** | `-1300450` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VALUE_OVER_TIME` (-1300452)

| Field | Value |
|-------|-------|
| **Number** | `-1300452` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_VOLUME_OVER_TIME` (-1300454)

| Field | Value |
|-------|-------|
| **Number** | `-1300454` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_QUANTITY_OVER_TIME` (-1300456)

| Field | Value |
|-------|-------|
| **Number** | `-1300456` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_AGGREGATED` (-1300458)

| Field | Value |
|-------|-------|
| **Number** | `-1300458` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300460)

| Field | Value |
|-------|-------|
| **Number** | `-1300460` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VALUE_AGGREGATED` (-1300462)

| Field | Value |
|-------|-------|
| **Number** | `-1300462` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VOLUME_AGGREGATED` (-1300464)

| Field | Value |
|-------|-------|
| **Number** | `-1300464` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_QUANTITY_AGGREGATED` (-1300466)

| Field | Value |
|-------|-------|
| **Number** | `-1300466` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_TRANSIENT` (-1300468)

| Field | Value |
|-------|-------|
| **Number** | `-1300468` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300470)

| Field | Value |
|-------|-------|
| **Number** | `-1300470` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VALUE_TRANSIENT` (-1300472)

| Field | Value |
|-------|-------|
| **Number** | `-1300472` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VOLUME_TRANSIENT` (-1300474)

| Field | Value |
|-------|-------|
| **Number** | `-1300474` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_QUANTITY_TRANSIENT` (-1300476)

| Field | Value |
|-------|-------|
| **Number** | `-1300476` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_OVER_TIME` (-1300478)

| Field | Value |
|-------|-------|
| **Number** | `-1300478` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300480)

| Field | Value |
|-------|-------|
| **Number** | `-1300480` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VALUE_OVER_TIME` (-1300482)

| Field | Value |
|-------|-------|
| **Number** | `-1300482` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_VOLUME_OVER_TIME` (-1300484)

| Field | Value |
|-------|-------|
| **Number** | `-1300484` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_QUANTITY_OVER_TIME` (-1300486)

| Field | Value |
|-------|-------|
| **Number** | `-1300486` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_AGGREGATED` (-1300488)

| Field | Value |
|-------|-------|
| **Number** | `-1300488` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300490)

| Field | Value |
|-------|-------|
| **Number** | `-1300490` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VALUE_AGGREGATED` (-1300492)

| Field | Value |
|-------|-------|
| **Number** | `-1300492` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VOLUME_AGGREGATED` (-1300494)

| Field | Value |
|-------|-------|
| **Number** | `-1300494` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_QUANTITY_AGGREGATED` (-1300496)

| Field | Value |
|-------|-------|
| **Number** | `-1300496` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_TRANSIENT` (-1300498)

| Field | Value |
|-------|-------|
| **Number** | `-1300498` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300500)

| Field | Value |
|-------|-------|
| **Number** | `-1300500` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VALUE_TRANSIENT` (-1300502)

| Field | Value |
|-------|-------|
| **Number** | `-1300502` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VOLUME_TRANSIENT` (-1300504)

| Field | Value |
|-------|-------|
| **Number** | `-1300504` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_QUANTITY_TRANSIENT` (-1300506)

| Field | Value |
|-------|-------|
| **Number** | `-1300506` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_OVER_TIME` (-1300508)

| Field | Value |
|-------|-------|
| **Number** | `-1300508` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300510)

| Field | Value |
|-------|-------|
| **Number** | `-1300510` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VALUE_OVER_TIME` (-1300512)

| Field | Value |
|-------|-------|
| **Number** | `-1300512` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_VOLUME_OVER_TIME` (-1300514)

| Field | Value |
|-------|-------|
| **Number** | `-1300514` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_QUANTITY_OVER_TIME` (-1300516)

| Field | Value |
|-------|-------|
| **Number** | `-1300516` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_AGGREGATED` (-1300518)

| Field | Value |
|-------|-------|
| **Number** | `-1300518` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_AGGREGATED` (-1300520)

| Field | Value |
|-------|-------|
| **Number** | `-1300520` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VALUE_AGGREGATED` (-1300522)

| Field | Value |
|-------|-------|
| **Number** | `-1300522` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VOLUME_AGGREGATED` (-1300524)

| Field | Value |
|-------|-------|
| **Number** | `-1300524` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_QUANTITY_AGGREGATED` (-1300526)

| Field | Value |
|-------|-------|
| **Number** | `-1300526` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_TRANSIENT` (-1300528)

| Field | Value |
|-------|-------|
| **Number** | `-1300528` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_TRANSIENT` (-1300530)

| Field | Value |
|-------|-------|
| **Number** | `-1300530` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VALUE_TRANSIENT` (-1300532)

| Field | Value |
|-------|-------|
| **Number** | `-1300532` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VOLUME_TRANSIENT` (-1300534)

| Field | Value |
|-------|-------|
| **Number** | `-1300534` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_QUANTITY_TRANSIENT` (-1300536)

| Field | Value |
|-------|-------|
| **Number** | `-1300536` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_OVER_TIME` (-1300538)

| Field | Value |
|-------|-------|
| **Number** | `-1300538` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_OVER_TIME` (-1300540)

| Field | Value |
|-------|-------|
| **Number** | `-1300540` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VALUE_OVER_TIME` (-1300542)

| Field | Value |
|-------|-------|
| **Number** | `-1300542` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_VOLUME_OVER_TIME` (-1300544)

| Field | Value |
|-------|-------|
| **Number** | `-1300544` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_QUANTITY_OVER_TIME` (-1300546)

| Field | Value |
|-------|-------|
| **Number** | `-1300546` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_AGGREGATED` (-1300548)

| Field | Value |
|-------|-------|
| **Number** | `-1300548` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300550)

| Field | Value |
|-------|-------|
| **Number** | `-1300550` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VALUE_AGGREGATED` (-1300552)

| Field | Value |
|-------|-------|
| **Number** | `-1300552` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VOLUME_AGGREGATED` (-1300554)

| Field | Value |
|-------|-------|
| **Number** | `-1300554` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_QUANTITY_AGGREGATED` (-1300556)

| Field | Value |
|-------|-------|
| **Number** | `-1300556` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_TRANSIENT` (-1300558)

| Field | Value |
|-------|-------|
| **Number** | `-1300558` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300560)

| Field | Value |
|-------|-------|
| **Number** | `-1300560` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VALUE_TRANSIENT` (-1300562)

| Field | Value |
|-------|-------|
| **Number** | `-1300562` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VOLUME_TRANSIENT` (-1300564)

| Field | Value |
|-------|-------|
| **Number** | `-1300564` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_QUANTITY_TRANSIENT` (-1300566)

| Field | Value |
|-------|-------|
| **Number** | `-1300566` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_OVER_TIME` (-1300568)

| Field | Value |
|-------|-------|
| **Number** | `-1300568` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300570)

| Field | Value |
|-------|-------|
| **Number** | `-1300570` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VALUE_OVER_TIME` (-1300572)

| Field | Value |
|-------|-------|
| **Number** | `-1300572` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_VOLUME_OVER_TIME` (-1300574)

| Field | Value |
|-------|-------|
| **Number** | `-1300574` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_QUANTITY_OVER_TIME` (-1300576)

| Field | Value |
|-------|-------|
| **Number** | `-1300576` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_AGGREGATED` (-1300578)

| Field | Value |
|-------|-------|
| **Number** | `-1300578` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300580)

| Field | Value |
|-------|-------|
| **Number** | `-1300580` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VALUE_AGGREGATED` (-1300582)

| Field | Value |
|-------|-------|
| **Number** | `-1300582` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VOLUME_AGGREGATED` (-1300584)

| Field | Value |
|-------|-------|
| **Number** | `-1300584` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_QUANTITY_AGGREGATED` (-1300586)

| Field | Value |
|-------|-------|
| **Number** | `-1300586` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_TRANSIENT` (-1300588)

| Field | Value |
|-------|-------|
| **Number** | `-1300588` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300590)

| Field | Value |
|-------|-------|
| **Number** | `-1300590` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VALUE_TRANSIENT` (-1300592)

| Field | Value |
|-------|-------|
| **Number** | `-1300592` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VOLUME_TRANSIENT` (-1300594)

| Field | Value |
|-------|-------|
| **Number** | `-1300594` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_QUANTITY_TRANSIENT` (-1300596)

| Field | Value |
|-------|-------|
| **Number** | `-1300596` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_OVER_TIME` (-1300598)

| Field | Value |
|-------|-------|
| **Number** | `-1300598` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300600)

| Field | Value |
|-------|-------|
| **Number** | `-1300600` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VALUE_OVER_TIME` (-1300602)

| Field | Value |
|-------|-------|
| **Number** | `-1300602` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_VOLUME_OVER_TIME` (-1300604)

| Field | Value |
|-------|-------|
| **Number** | `-1300604` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_QUANTITY_OVER_TIME` (-1300606)

| Field | Value |
|-------|-------|
| **Number** | `-1300606` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_AGGREGATED` (-1300608)

| Field | Value |
|-------|-------|
| **Number** | `-1300608` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_COUNT_AGGREGATED` (-1300610)

| Field | Value |
|-------|-------|
| **Number** | `-1300610` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VALUE_AGGREGATED` (-1300612)

| Field | Value |
|-------|-------|
| **Number** | `-1300612` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VOLUME_AGGREGATED` (-1300614)

| Field | Value |
|-------|-------|
| **Number** | `-1300614` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_QUANTITY_AGGREGATED` (-1300616)

| Field | Value |
|-------|-------|
| **Number** | `-1300616` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_TRANSIENT` (-1300618)

| Field | Value |
|-------|-------|
| **Number** | `-1300618` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_COUNT_TRANSIENT` (-1300620)

| Field | Value |
|-------|-------|
| **Number** | `-1300620` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VALUE_TRANSIENT` (-1300622)

| Field | Value |
|-------|-------|
| **Number** | `-1300622` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VOLUME_TRANSIENT` (-1300624)

| Field | Value |
|-------|-------|
| **Number** | `-1300624` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_QUANTITY_TRANSIENT` (-1300626)

| Field | Value |
|-------|-------|
| **Number** | `-1300626` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_OVER_TIME` (-1300628)

| Field | Value |
|-------|-------|
| **Number** | `-1300628` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_EXPOSURE_COUNT_OVER_TIME` (-1300630)

| Field | Value |
|-------|-------|
| **Number** | `-1300630` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VALUE_OVER_TIME` (-1300632)

| Field | Value |
|-------|-------|
| **Number** | `-1300632` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_VOLUME_OVER_TIME` (-1300634)

| Field | Value |
|-------|-------|
| **Number** | `-1300634` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_QUANTITY_OVER_TIME` (-1300636)

| Field | Value |
|-------|-------|
| **Number** | `-1300636` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_AGGREGATED` (-1300638)

| Field | Value |
|-------|-------|
| **Number** | `-1300638` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_COUNT_AGGREGATED` (-1300640)

| Field | Value |
|-------|-------|
| **Number** | `-1300640` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VALUE_AGGREGATED` (-1300642)

| Field | Value |
|-------|-------|
| **Number** | `-1300642` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VOLUME_AGGREGATED` (-1300644)

| Field | Value |
|-------|-------|
| **Number** | `-1300644` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_QUANTITY_AGGREGATED` (-1300646)

| Field | Value |
|-------|-------|
| **Number** | `-1300646` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_TRANSIENT` (-1300648)

| Field | Value |
|-------|-------|
| **Number** | `-1300648` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_COUNT_TRANSIENT` (-1300650)

| Field | Value |
|-------|-------|
| **Number** | `-1300650` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VALUE_TRANSIENT` (-1300652)

| Field | Value |
|-------|-------|
| **Number** | `-1300652` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VOLUME_TRANSIENT` (-1300654)

| Field | Value |
|-------|-------|
| **Number** | `-1300654` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_QUANTITY_TRANSIENT` (-1300656)

| Field | Value |
|-------|-------|
| **Number** | `-1300656` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_OVER_TIME` (-1300658)

| Field | Value |
|-------|-------|
| **Number** | `-1300658` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_EXPOSURE_COUNT_OVER_TIME` (-1300660)

| Field | Value |
|-------|-------|
| **Number** | `-1300660` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VALUE_OVER_TIME` (-1300662)

| Field | Value |
|-------|-------|
| **Number** | `-1300662` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_VOLUME_OVER_TIME` (-1300664)

| Field | Value |
|-------|-------|
| **Number** | `-1300664` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_QUANTITY_OVER_TIME` (-1300666)

| Field | Value |
|-------|-------|
| **Number** | `-1300666` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_AGGREGATED` (-1300668)

| Field | Value |
|-------|-------|
| **Number** | `-1300668` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_COUNT_AGGREGATED` (-1300670)

| Field | Value |
|-------|-------|
| **Number** | `-1300670` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VALUE_AGGREGATED` (-1300672)

| Field | Value |
|-------|-------|
| **Number** | `-1300672` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VOLUME_AGGREGATED` (-1300674)

| Field | Value |
|-------|-------|
| **Number** | `-1300674` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_QUANTITY_AGGREGATED` (-1300676)

| Field | Value |
|-------|-------|
| **Number** | `-1300676` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_TRANSIENT` (-1300678)

| Field | Value |
|-------|-------|
| **Number** | `-1300678` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_COUNT_TRANSIENT` (-1300680)

| Field | Value |
|-------|-------|
| **Number** | `-1300680` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VALUE_TRANSIENT` (-1300682)

| Field | Value |
|-------|-------|
| **Number** | `-1300682` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VOLUME_TRANSIENT` (-1300684)

| Field | Value |
|-------|-------|
| **Number** | `-1300684` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_QUANTITY_TRANSIENT` (-1300686)

| Field | Value |
|-------|-------|
| **Number** | `-1300686` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_OVER_TIME` (-1300688)

| Field | Value |
|-------|-------|
| **Number** | `-1300688` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_EXPOSURE_COUNT_OVER_TIME` (-1300690)

| Field | Value |
|-------|-------|
| **Number** | `-1300690` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VALUE_OVER_TIME` (-1300692)

| Field | Value |
|-------|-------|
| **Number** | `-1300692` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_VOLUME_OVER_TIME` (-1300694)

| Field | Value |
|-------|-------|
| **Number** | `-1300694` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_QUANTITY_OVER_TIME` (-1300696)

| Field | Value |
|-------|-------|
| **Number** | `-1300696` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_AGGREGATED` (-1300698)

| Field | Value |
|-------|-------|
| **Number** | `-1300698` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_COUNT_AGGREGATED` (-1300700)

| Field | Value |
|-------|-------|
| **Number** | `-1300700` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VALUE_AGGREGATED` (-1300702)

| Field | Value |
|-------|-------|
| **Number** | `-1300702` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VOLUME_AGGREGATED` (-1300704)

| Field | Value |
|-------|-------|
| **Number** | `-1300704` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_QUANTITY_AGGREGATED` (-1300706)

| Field | Value |
|-------|-------|
| **Number** | `-1300706` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_TRANSIENT` (-1300708)

| Field | Value |
|-------|-------|
| **Number** | `-1300708` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_COUNT_TRANSIENT` (-1300710)

| Field | Value |
|-------|-------|
| **Number** | `-1300710` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VALUE_TRANSIENT` (-1300712)

| Field | Value |
|-------|-------|
| **Number** | `-1300712` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VOLUME_TRANSIENT` (-1300714)

| Field | Value |
|-------|-------|
| **Number** | `-1300714` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_QUANTITY_TRANSIENT` (-1300716)

| Field | Value |
|-------|-------|
| **Number** | `-1300716` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_OVER_TIME` (-1300718)

| Field | Value |
|-------|-------|
| **Number** | `-1300718` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_EXPOSURE_COUNT_OVER_TIME` (-1300720)

| Field | Value |
|-------|-------|
| **Number** | `-1300720` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VALUE_OVER_TIME` (-1300722)

| Field | Value |
|-------|-------|
| **Number** | `-1300722` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_VOLUME_OVER_TIME` (-1300724)

| Field | Value |
|-------|-------|
| **Number** | `-1300724` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_QUANTITY_OVER_TIME` (-1300726)

| Field | Value |
|-------|-------|
| **Number** | `-1300726` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_AGGREGATED` (-1300728)

| Field | Value |
|-------|-------|
| **Number** | `-1300728` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_COUNT_AGGREGATED` (-1300730)

| Field | Value |
|-------|-------|
| **Number** | `-1300730` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VALUE_AGGREGATED` (-1300732)

| Field | Value |
|-------|-------|
| **Number** | `-1300732` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VOLUME_AGGREGATED` (-1300734)

| Field | Value |
|-------|-------|
| **Number** | `-1300734` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_QUANTITY_AGGREGATED` (-1300736)

| Field | Value |
|-------|-------|
| **Number** | `-1300736` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_TRANSIENT` (-1300738)

| Field | Value |
|-------|-------|
| **Number** | `-1300738` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_COUNT_TRANSIENT` (-1300740)

| Field | Value |
|-------|-------|
| **Number** | `-1300740` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VALUE_TRANSIENT` (-1300742)

| Field | Value |
|-------|-------|
| **Number** | `-1300742` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VOLUME_TRANSIENT` (-1300744)

| Field | Value |
|-------|-------|
| **Number** | `-1300744` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_QUANTITY_TRANSIENT` (-1300746)

| Field | Value |
|-------|-------|
| **Number** | `-1300746` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_OVER_TIME` (-1300748)

| Field | Value |
|-------|-------|
| **Number** | `-1300748` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_EXPOSURE_COUNT_OVER_TIME` (-1300750)

| Field | Value |
|-------|-------|
| **Number** | `-1300750` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VALUE_OVER_TIME` (-1300752)

| Field | Value |
|-------|-------|
| **Number** | `-1300752` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_VOLUME_OVER_TIME` (-1300754)

| Field | Value |
|-------|-------|
| **Number** | `-1300754` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_QUANTITY_OVER_TIME` (-1300756)

| Field | Value |
|-------|-------|
| **Number** | `-1300756` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_AGGREGATED` (-1300758)

| Field | Value |
|-------|-------|
| **Number** | `-1300758` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_AGGREGATED` (-1300760)

| Field | Value |
|-------|-------|
| **Number** | `-1300760` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VALUE_AGGREGATED` (-1300762)

| Field | Value |
|-------|-------|
| **Number** | `-1300762` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VOLUME_AGGREGATED` (-1300764)

| Field | Value |
|-------|-------|
| **Number** | `-1300764` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_QUANTITY_AGGREGATED` (-1300766)

| Field | Value |
|-------|-------|
| **Number** | `-1300766` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_TRANSIENT` (-1300768)

| Field | Value |
|-------|-------|
| **Number** | `-1300768` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_TRANSIENT` (-1300770)

| Field | Value |
|-------|-------|
| **Number** | `-1300770` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VALUE_TRANSIENT` (-1300772)

| Field | Value |
|-------|-------|
| **Number** | `-1300772` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VOLUME_TRANSIENT` (-1300774)

| Field | Value |
|-------|-------|
| **Number** | `-1300774` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_QUANTITY_TRANSIENT` (-1300776)

| Field | Value |
|-------|-------|
| **Number** | `-1300776` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_OVER_TIME` (-1300778)

| Field | Value |
|-------|-------|
| **Number** | `-1300778` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_EXPOSURE_COUNT_OVER_TIME` (-1300780)

| Field | Value |
|-------|-------|
| **Number** | `-1300780` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VALUE_OVER_TIME` (-1300782)

| Field | Value |
|-------|-------|
| **Number** | `-1300782` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_VOLUME_OVER_TIME` (-1300784)

| Field | Value |
|-------|-------|
| **Number** | `-1300784` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_QUANTITY_OVER_TIME` (-1300786)

| Field | Value |
|-------|-------|
| **Number** | `-1300786` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_AGGREGATED` (-1300788)

| Field | Value |
|-------|-------|
| **Number** | `-1300788` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_AGGREGATED` (-1300790)

| Field | Value |
|-------|-------|
| **Number** | `-1300790` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VALUE_AGGREGATED` (-1300792)

| Field | Value |
|-------|-------|
| **Number** | `-1300792` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VOLUME_AGGREGATED` (-1300794)

| Field | Value |
|-------|-------|
| **Number** | `-1300794` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_QUANTITY_AGGREGATED` (-1300796)

| Field | Value |
|-------|-------|
| **Number** | `-1300796` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_TRANSIENT` (-1300798)

| Field | Value |
|-------|-------|
| **Number** | `-1300798` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_TRANSIENT` (-1300800)

| Field | Value |
|-------|-------|
| **Number** | `-1300800` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VALUE_TRANSIENT` (-1300802)

| Field | Value |
|-------|-------|
| **Number** | `-1300802` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VOLUME_TRANSIENT` (-1300804)

| Field | Value |
|-------|-------|
| **Number** | `-1300804` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_QUANTITY_TRANSIENT` (-1300806)

| Field | Value |
|-------|-------|
| **Number** | `-1300806` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_OVER_TIME` (-1300808)

| Field | Value |
|-------|-------|
| **Number** | `-1300808` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_EXPOSURE_COUNT_OVER_TIME` (-1300810)

| Field | Value |
|-------|-------|
| **Number** | `-1300810` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VALUE_OVER_TIME` (-1300812)

| Field | Value |
|-------|-------|
| **Number** | `-1300812` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_VOLUME_OVER_TIME` (-1300814)

| Field | Value |
|-------|-------|
| **Number** | `-1300814` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_QUANTITY_OVER_TIME` (-1300816)

| Field | Value |
|-------|-------|
| **Number** | `-1300816` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_AGGREGATED` (-1300818)

| Field | Value |
|-------|-------|
| **Number** | `-1300818` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_COUNT_AGGREGATED` (-1300820)

| Field | Value |
|-------|-------|
| **Number** | `-1300820` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VALUE_AGGREGATED` (-1300822)

| Field | Value |
|-------|-------|
| **Number** | `-1300822` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VOLUME_AGGREGATED` (-1300824)

| Field | Value |
|-------|-------|
| **Number** | `-1300824` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_QUANTITY_AGGREGATED` (-1300826)

| Field | Value |
|-------|-------|
| **Number** | `-1300826` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_TRANSIENT` (-1300828)

| Field | Value |
|-------|-------|
| **Number** | `-1300828` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_COUNT_TRANSIENT` (-1300830)

| Field | Value |
|-------|-------|
| **Number** | `-1300830` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VALUE_TRANSIENT` (-1300832)

| Field | Value |
|-------|-------|
| **Number** | `-1300832` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VOLUME_TRANSIENT` (-1300834)

| Field | Value |
|-------|-------|
| **Number** | `-1300834` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_QUANTITY_TRANSIENT` (-1300836)

| Field | Value |
|-------|-------|
| **Number** | `-1300836` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_OVER_TIME` (-1300838)

| Field | Value |
|-------|-------|
| **Number** | `-1300838` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_EXPOSURE_COUNT_OVER_TIME` (-1300840)

| Field | Value |
|-------|-------|
| **Number** | `-1300840` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VALUE_OVER_TIME` (-1300842)

| Field | Value |
|-------|-------|
| **Number** | `-1300842` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_VOLUME_OVER_TIME` (-1300844)

| Field | Value |
|-------|-------|
| **Number** | `-1300844` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_QUANTITY_OVER_TIME` (-1300846)

| Field | Value |
|-------|-------|
| **Number** | `-1300846` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_AGGREGATED` (-1300848)

| Field | Value |
|-------|-------|
| **Number** | `-1300848` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300850)

| Field | Value |
|-------|-------|
| **Number** | `-1300850` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VALUE_AGGREGATED` (-1300852)

| Field | Value |
|-------|-------|
| **Number** | `-1300852` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VOLUME_AGGREGATED` (-1300854)

| Field | Value |
|-------|-------|
| **Number** | `-1300854` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_QUANTITY_AGGREGATED` (-1300856)

| Field | Value |
|-------|-------|
| **Number** | `-1300856` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_TRANSIENT` (-1300858)

| Field | Value |
|-------|-------|
| **Number** | `-1300858` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300860)

| Field | Value |
|-------|-------|
| **Number** | `-1300860` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VALUE_TRANSIENT` (-1300862)

| Field | Value |
|-------|-------|
| **Number** | `-1300862` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VOLUME_TRANSIENT` (-1300864)

| Field | Value |
|-------|-------|
| **Number** | `-1300864` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_QUANTITY_TRANSIENT` (-1300866)

| Field | Value |
|-------|-------|
| **Number** | `-1300866` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_OVER_TIME` (-1300868)

| Field | Value |
|-------|-------|
| **Number** | `-1300868` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300870)

| Field | Value |
|-------|-------|
| **Number** | `-1300870` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VALUE_OVER_TIME` (-1300872)

| Field | Value |
|-------|-------|
| **Number** | `-1300872` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_VOLUME_OVER_TIME` (-1300874)

| Field | Value |
|-------|-------|
| **Number** | `-1300874` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_QUANTITY_OVER_TIME` (-1300876)

| Field | Value |
|-------|-------|
| **Number** | `-1300876` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_AGGREGATED` (-1300878)

| Field | Value |
|-------|-------|
| **Number** | `-1300878` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_COUNT_AGGREGATED` (-1300880)

| Field | Value |
|-------|-------|
| **Number** | `-1300880` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VALUE_AGGREGATED` (-1300882)

| Field | Value |
|-------|-------|
| **Number** | `-1300882` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VOLUME_AGGREGATED` (-1300884)

| Field | Value |
|-------|-------|
| **Number** | `-1300884` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_QUANTITY_AGGREGATED` (-1300886)

| Field | Value |
|-------|-------|
| **Number** | `-1300886` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_TRANSIENT` (-1300888)

| Field | Value |
|-------|-------|
| **Number** | `-1300888` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_COUNT_TRANSIENT` (-1300890)

| Field | Value |
|-------|-------|
| **Number** | `-1300890` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VALUE_TRANSIENT` (-1300892)

| Field | Value |
|-------|-------|
| **Number** | `-1300892` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VOLUME_TRANSIENT` (-1300894)

| Field | Value |
|-------|-------|
| **Number** | `-1300894` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_QUANTITY_TRANSIENT` (-1300896)

| Field | Value |
|-------|-------|
| **Number** | `-1300896` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_OVER_TIME` (-1300898)

| Field | Value |
|-------|-------|
| **Number** | `-1300898` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_EXPOSURE_COUNT_OVER_TIME` (-1300900)

| Field | Value |
|-------|-------|
| **Number** | `-1300900` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VALUE_OVER_TIME` (-1300902)

| Field | Value |
|-------|-------|
| **Number** | `-1300902` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_VOLUME_OVER_TIME` (-1300904)

| Field | Value |
|-------|-------|
| **Number** | `-1300904` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_QUANTITY_OVER_TIME` (-1300906)

| Field | Value |
|-------|-------|
| **Number** | `-1300906` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_AGGREGATED` (-1300908)

| Field | Value |
|-------|-------|
| **Number** | `-1300908` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_COUNT_AGGREGATED` (-1300910)

| Field | Value |
|-------|-------|
| **Number** | `-1300910` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VALUE_AGGREGATED` (-1300912)

| Field | Value |
|-------|-------|
| **Number** | `-1300912` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VOLUME_AGGREGATED` (-1300914)

| Field | Value |
|-------|-------|
| **Number** | `-1300914` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_QUANTITY_AGGREGATED` (-1300916)

| Field | Value |
|-------|-------|
| **Number** | `-1300916` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_TRANSIENT` (-1300918)

| Field | Value |
|-------|-------|
| **Number** | `-1300918` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_COUNT_TRANSIENT` (-1300920)

| Field | Value |
|-------|-------|
| **Number** | `-1300920` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VALUE_TRANSIENT` (-1300922)

| Field | Value |
|-------|-------|
| **Number** | `-1300922` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VOLUME_TRANSIENT` (-1300924)

| Field | Value |
|-------|-------|
| **Number** | `-1300924` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_QUANTITY_TRANSIENT` (-1300926)

| Field | Value |
|-------|-------|
| **Number** | `-1300926` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_OVER_TIME` (-1300928)

| Field | Value |
|-------|-------|
| **Number** | `-1300928` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_EXPOSURE_COUNT_OVER_TIME` (-1300930)

| Field | Value |
|-------|-------|
| **Number** | `-1300930` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VALUE_OVER_TIME` (-1300932)

| Field | Value |
|-------|-------|
| **Number** | `-1300932` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_VOLUME_OVER_TIME` (-1300934)

| Field | Value |
|-------|-------|
| **Number** | `-1300934` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_QUANTITY_OVER_TIME` (-1300936)

| Field | Value |
|-------|-------|
| **Number** | `-1300936` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_AGGREGATED` (-1300938)

| Field | Value |
|-------|-------|
| **Number** | `-1300938` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300940)

| Field | Value |
|-------|-------|
| **Number** | `-1300940` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VALUE_AGGREGATED` (-1300942)

| Field | Value |
|-------|-------|
| **Number** | `-1300942` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VOLUME_AGGREGATED` (-1300944)

| Field | Value |
|-------|-------|
| **Number** | `-1300944` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_QUANTITY_AGGREGATED` (-1300946)

| Field | Value |
|-------|-------|
| **Number** | `-1300946` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_TRANSIENT` (-1300948)

| Field | Value |
|-------|-------|
| **Number** | `-1300948` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300950)

| Field | Value |
|-------|-------|
| **Number** | `-1300950` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VALUE_TRANSIENT` (-1300952)

| Field | Value |
|-------|-------|
| **Number** | `-1300952` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VOLUME_TRANSIENT` (-1300954)

| Field | Value |
|-------|-------|
| **Number** | `-1300954` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_QUANTITY_TRANSIENT` (-1300956)

| Field | Value |
|-------|-------|
| **Number** | `-1300956` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_OVER_TIME` (-1300958)

| Field | Value |
|-------|-------|
| **Number** | `-1300958` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300960)

| Field | Value |
|-------|-------|
| **Number** | `-1300960` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VALUE_OVER_TIME` (-1300962)

| Field | Value |
|-------|-------|
| **Number** | `-1300962` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_VOLUME_OVER_TIME` (-1300964)

| Field | Value |
|-------|-------|
| **Number** | `-1300964` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_QUANTITY_OVER_TIME` (-1300966)

| Field | Value |
|-------|-------|
| **Number** | `-1300966` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_AGGREGATED` (-1300968)

| Field | Value |
|-------|-------|
| **Number** | `-1300968` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_COUNT_AGGREGATED` (-1300970)

| Field | Value |
|-------|-------|
| **Number** | `-1300970` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VALUE_AGGREGATED` (-1300972)

| Field | Value |
|-------|-------|
| **Number** | `-1300972` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VOLUME_AGGREGATED` (-1300974)

| Field | Value |
|-------|-------|
| **Number** | `-1300974` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_QUANTITY_AGGREGATED` (-1300976)

| Field | Value |
|-------|-------|
| **Number** | `-1300976` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_TRANSIENT` (-1300978)

| Field | Value |
|-------|-------|
| **Number** | `-1300978` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_COUNT_TRANSIENT` (-1300980)

| Field | Value |
|-------|-------|
| **Number** | `-1300980` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VALUE_TRANSIENT` (-1300982)

| Field | Value |
|-------|-------|
| **Number** | `-1300982` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VOLUME_TRANSIENT` (-1300984)

| Field | Value |
|-------|-------|
| **Number** | `-1300984` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_QUANTITY_TRANSIENT` (-1300986)

| Field | Value |
|-------|-------|
| **Number** | `-1300986` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_OVER_TIME` (-1300988)

| Field | Value |
|-------|-------|
| **Number** | `-1300988` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_EXPOSURE_COUNT_OVER_TIME` (-1300990)

| Field | Value |
|-------|-------|
| **Number** | `-1300990` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VALUE_OVER_TIME` (-1300992)

| Field | Value |
|-------|-------|
| **Number** | `-1300992` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_VOLUME_OVER_TIME` (-1300994)

| Field | Value |
|-------|-------|
| **Number** | `-1300994` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_QUANTITY_OVER_TIME` (-1300996)

| Field | Value |
|-------|-------|
| **Number** | `-1300996` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_AGGREGATED` (-1300998)

| Field | Value |
|-------|-------|
| **Number** | `-1300998` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_COUNT_AGGREGATED` (-1301000)

| Field | Value |
|-------|-------|
| **Number** | `-1301000` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VALUE_AGGREGATED` (-1301002)

| Field | Value |
|-------|-------|
| **Number** | `-1301002` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VOLUME_AGGREGATED` (-1301004)

| Field | Value |
|-------|-------|
| **Number** | `-1301004` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_QUANTITY_AGGREGATED` (-1301006)

| Field | Value |
|-------|-------|
| **Number** | `-1301006` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_TRANSIENT` (-1301008)

| Field | Value |
|-------|-------|
| **Number** | `-1301008` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_COUNT_TRANSIENT` (-1301010)

| Field | Value |
|-------|-------|
| **Number** | `-1301010` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VALUE_TRANSIENT` (-1301012)

| Field | Value |
|-------|-------|
| **Number** | `-1301012` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VOLUME_TRANSIENT` (-1301014)

| Field | Value |
|-------|-------|
| **Number** | `-1301014` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_QUANTITY_TRANSIENT` (-1301016)

| Field | Value |
|-------|-------|
| **Number** | `-1301016` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_OVER_TIME` (-1301018)

| Field | Value |
|-------|-------|
| **Number** | `-1301018` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_EXPOSURE_COUNT_OVER_TIME` (-1301020)

| Field | Value |
|-------|-------|
| **Number** | `-1301020` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VALUE_OVER_TIME` (-1301022)

| Field | Value |
|-------|-------|
| **Number** | `-1301022` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_VOLUME_OVER_TIME` (-1301024)

| Field | Value |
|-------|-------|
| **Number** | `-1301024` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_QUANTITY_OVER_TIME` (-1301026)

| Field | Value |
|-------|-------|
| **Number** | `-1301026` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_AGGREGATED` (-1301028)

| Field | Value |
|-------|-------|
| **Number** | `-1301028` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_COUNT_AGGREGATED` (-1301030)

| Field | Value |
|-------|-------|
| **Number** | `-1301030` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VALUE_AGGREGATED` (-1301032)

| Field | Value |
|-------|-------|
| **Number** | `-1301032` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VOLUME_AGGREGATED` (-1301034)

| Field | Value |
|-------|-------|
| **Number** | `-1301034` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_QUANTITY_AGGREGATED` (-1301036)

| Field | Value |
|-------|-------|
| **Number** | `-1301036` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_TRANSIENT` (-1301038)

| Field | Value |
|-------|-------|
| **Number** | `-1301038` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_COUNT_TRANSIENT` (-1301040)

| Field | Value |
|-------|-------|
| **Number** | `-1301040` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VALUE_TRANSIENT` (-1301042)

| Field | Value |
|-------|-------|
| **Number** | `-1301042` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VOLUME_TRANSIENT` (-1301044)

| Field | Value |
|-------|-------|
| **Number** | `-1301044` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_QUANTITY_TRANSIENT` (-1301046)

| Field | Value |
|-------|-------|
| **Number** | `-1301046` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_OVER_TIME` (-1301048)

| Field | Value |
|-------|-------|
| **Number** | `-1301048` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_EXPOSURE_COUNT_OVER_TIME` (-1301050)

| Field | Value |
|-------|-------|
| **Number** | `-1301050` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VALUE_OVER_TIME` (-1301052)

| Field | Value |
|-------|-------|
| **Number** | `-1301052` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_VOLUME_OVER_TIME` (-1301054)

| Field | Value |
|-------|-------|
| **Number** | `-1301054` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_QUANTITY_OVER_TIME` (-1301056)

| Field | Value |
|-------|-------|
| **Number** | `-1301056` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_AGGREGATED` (-1301058)

| Field | Value |
|-------|-------|
| **Number** | `-1301058` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_AGGREGATED` (-1301060)

| Field | Value |
|-------|-------|
| **Number** | `-1301060` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VALUE_AGGREGATED` (-1301062)

| Field | Value |
|-------|-------|
| **Number** | `-1301062` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VOLUME_AGGREGATED` (-1301064)

| Field | Value |
|-------|-------|
| **Number** | `-1301064` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_QUANTITY_AGGREGATED` (-1301066)

| Field | Value |
|-------|-------|
| **Number** | `-1301066` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_TRANSIENT` (-1301068)

| Field | Value |
|-------|-------|
| **Number** | `-1301068` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation EXPOSURE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_TRANSIENT` (-1301070)

| Field | Value |
|-------|-------|
| **Number** | `-1301070` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation EXPOSURE_COUNT limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VALUE_TRANSIENT` (-1301072)

| Field | Value |
|-------|-------|
| **Number** | `-1301072` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation VALUE limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VOLUME_TRANSIENT` (-1301074)

| Field | Value |
|-------|-------|
| **Number** | `-1301074` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation VOLUME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_QUANTITY_TRANSIENT` (-1301076)

| Field | Value |
|-------|-------|
| **Number** | `-1301076` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation QUANTITY limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_OVER_TIME` (-1301078)

| Field | Value |
|-------|-------|
| **Number** | `-1301078` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation EXPOSURE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_EXPOSURE_COUNT_OVER_TIME` (-1301080)

| Field | Value |
|-------|-------|
| **Number** | `-1301080` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation EXPOSURE_COUNT OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VALUE_OVER_TIME` (-1301082)

| Field | Value |
|-------|-------|
| **Number** | `-1301082` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation VALUE OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_VOLUME_OVER_TIME` (-1301084)

| Field | Value |
|-------|-------|
| **Number** | `-1301084` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation VOLUME OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_QUANTITY_OVER_TIME` (-1301086)

| Field | Value |
|-------|-------|
| **Number** | `-1301086` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation QUANTITY OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_TRANSACTIONS_AGGREGATED` (-1301088)

| Field | Value |
|-------|-------|
| **Number** | `-1301088` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_TRANSACTIONS_TRANSIENT` (-1301090)

| Field | Value |
|-------|-------|
| **Number** | `-1301090` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_SELL_TRANSACTIONS_OVER_TIME` (-1301092)

| Field | Value |
|-------|-------|
| **Number** | `-1301092` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_TRANSACTIONS_AGGREGATED` (-1301094)

| Field | Value |
|-------|-------|
| **Number** | `-1301094` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_TRANSACTIONS_TRANSIENT` (-1301096)

| Field | Value |
|-------|-------|
| **Number** | `-1301096` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_SELL_TRANSACTIONS_OVER_TIME` (-1301098)

| Field | Value |
|-------|-------|
| **Number** | `-1301098` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_TRANSACTIONS_AGGREGATED` (-1301100)

| Field | Value |
|-------|-------|
| **Number** | `-1301100` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_TRANSACTIONS_TRANSIENT` (-1301102)

| Field | Value |
|-------|-------|
| **Number** | `-1301102` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_DEVIATION_TRANSACTIONS_OVER_TIME` (-1301104)

| Field | Value |
|-------|-------|
| **Number** | `-1301104` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price deviation TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_TRANSACTIONS_AGGREGATED` (-1301106)

| Field | Value |
|-------|-------|
| **Number** | `-1301106` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_TRANSACTIONS_TRANSIENT` (-1301108)

| Field | Value |
|-------|-------|
| **Number** | `-1301108` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_BOUGHT_TRANSACTIONS_OVER_TIME` (-1301110)

| Field | Value |
|-------|-------|
| **Number** | `-1301110` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded bought TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_TRANSACTIONS_AGGREGATED` (-1301112)

| Field | Value |
|-------|-------|
| **Number** | `-1301112` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_TRANSACTIONS_TRANSIENT` (-1301114)

| Field | Value |
|-------|-------|
| **Number** | `-1301114` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_BOUGHT_TRANSACTIONS_OVER_TIME` (-1301116)

| Field | Value |
|-------|-------|
| **Number** | `-1301116` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net bought TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_TRANSACTIONS_AGGREGATED` (-1301118)

| Field | Value |
|-------|-------|
| **Number** | `-1301118` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_TRANSACTIONS_TRANSIENT` (-1301120)

| Field | Value |
|-------|-------|
| **Number** | `-1301120` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_SOLD_TRANSACTIONS_OVER_TIME` (-1301122)

| Field | Value |
|-------|-------|
| **Number** | `-1301122` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded sold TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_TRANSACTIONS_AGGREGATED` (-1301124)

| Field | Value |
|-------|-------|
| **Number** | `-1301124` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_TRANSACTIONS_TRANSIENT` (-1301126)

| Field | Value |
|-------|-------|
| **Number** | `-1301126` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_SELL_TRANSACTIONS_OVER_TIME` (-1301128)

| Field | Value |
|-------|-------|
| **Number** | `-1301128` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_TRANSACTIONS_AGGREGATED` (-1301130)

| Field | Value |
|-------|-------|
| **Number** | `-1301130` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_TRANSACTIONS_TRANSIENT` (-1301132)

| Field | Value |
|-------|-------|
| **Number** | `-1301132` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_TRANSACTIONS_OVER_TIME` (-1301134)

| Field | Value |
|-------|-------|
| **Number** | `-1301134` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_AGGREGATED` (-1301136)

| Field | Value |
|-------|-------|
| **Number** | `-1301136` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open traded net TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_TRANSIENT` (-1301138)

| Field | Value |
|-------|-------|
| **Number** | `-1301138` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open traded net TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_OVER_TIME` (-1301140)

| Field | Value |
|-------|-------|
| **Number** | `-1301140` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open traded net TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRANSACTIONS_AGGREGATED` (-1301142)

| Field | Value |
|-------|-------|
| **Number** | `-1301142` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total open TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRANSACTIONS_TRANSIENT` (-1301144)

| Field | Value |
|-------|-------|
| **Number** | `-1301144` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total open TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_OPEN_TRANSACTIONS_OVER_TIME` (-1301146)

| Field | Value |
|-------|-------|
| **Number** | `-1301146` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total open TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_TRANSACTIONS_AGGREGATED` (-1301148)

| Field | Value |
|-------|-------|
| **Number** | `-1301148` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_TRANSACTIONS_TRANSIENT` (-1301150)

| Field | Value |
|-------|-------|
| **Number** | `-1301150` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_NET_BUY_TRANSACTIONS_OVER_TIME` (-1301152)

| Field | Value |
|-------|-------|
| **Number** | `-1301152` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total net buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_TRANSACTIONS_AGGREGATED` (-1301154)

| Field | Value |
|-------|-------|
| **Number** | `-1301154` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED traded net sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_TRANSACTIONS_TRANSIENT` (-1301156)

| Field | Value |
|-------|-------|
| **Number** | `-1301156` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT traded net sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TRADED_NET_SOLD_TRANSACTIONS_OVER_TIME` (-1301158)

| Field | Value |
|-------|-------|
| **Number** | `-1301158` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max traded net sold TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_TRANSACTIONS_AGGREGATED` (-1301160)

| Field | Value |
|-------|-------|
| **Number** | `-1301160` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_TRANSACTIONS_TRANSIENT` (-1301162)

| Field | Value |
|-------|-------|
| **Number** | `-1301162` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_TRADED_TRANSACTIONS_OVER_TIME` (-1301164)

| Field | Value |
|-------|-------|
| **Number** | `-1301164` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total traded TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_TRANSACTIONS_AGGREGATED` (-1301166)

| Field | Value |
|-------|-------|
| **Number** | `-1301166` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED total buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_TRANSACTIONS_TRANSIENT` (-1301168)

| Field | Value |
|-------|-------|
| **Number** | `-1301168` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT total buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_TOTAL_BUY_TRANSACTIONS_OVER_TIME` (-1301170)

| Field | Value |
|-------|-------|
| **Number** | `-1301170` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max total buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_TRANSACTIONS_AGGREGATED` (-1301172)

| Field | Value |
|-------|-------|
| **Number** | `-1301172` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_TRANSACTIONS_TRANSIENT` (-1301174)

| Field | Value |
|-------|-------|
| **Number** | `-1301174` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_BUY_TRANSACTIONS_OVER_TIME` (-1301176)

| Field | Value |
|-------|-------|
| **Number** | `-1301176` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_TRANSACTIONS_AGGREGATED` (-1301178)

| Field | Value |
|-------|-------|
| **Number** | `-1301178` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_TRANSACTIONS_TRANSIENT` (-1301180)

| Field | Value |
|-------|-------|
| **Number** | `-1301180` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_SELL_TRANSACTIONS_OVER_TIME` (-1301182)

| Field | Value |
|-------|-------|
| **Number** | `-1301182` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_TRANSACTIONS_AGGREGATED` (-1301184)

| Field | Value |
|-------|-------|
| **Number** | `-1301184` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED open net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_TRANSACTIONS_TRANSIENT` (-1301186)

| Field | Value |
|-------|-------|
| **Number** | `-1301186` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT open net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_OPEN_NET_BUY_TRANSACTIONS_OVER_TIME` (-1301188)

| Field | Value |
|-------|-------|
| **Number** | `-1301188` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max open net buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_TRANSACTIONS_AGGREGATED` (-1301190)

| Field | Value |
|-------|-------|
| **Number** | `-1301190` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max AGGREGATED price tick deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_TRANSACTIONS_TRANSIENT` (-1301192)

| Field | Value |
|-------|-------|
| **Number** | `-1301192` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max TRANSIENT price tick deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_MAX_PRICE_TICK_DEVIATION_TRANSACTIONS_OVER_TIME` (-1301194)

| Field | Value |
|-------|-------|
| **Number** | `-1301194` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the max price tick deviation TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_TRANSACTIONS_AGGREGATED` (-1301196)

| Field | Value |
|-------|-------|
| **Number** | `-1301196` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_TRANSACTIONS_TRANSIENT` (-1301198)

| Field | Value |
|-------|-------|
| **Number** | `-1301198` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_SELL_TRANSACTIONS_OVER_TIME` (-1301200)

| Field | Value |
|-------|-------|
| **Number** | `-1301200` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_TRANSACTIONS_AGGREGATED` (-1301202)

| Field | Value |
|-------|-------|
| **Number** | `-1301202` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_TRANSACTIONS_TRANSIENT` (-1301204)

| Field | Value |
|-------|-------|
| **Number** | `-1301204` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_SELL_TRANSACTIONS_OVER_TIME` (-1301206)

| Field | Value |
|-------|-------|
| **Number** | `-1301206` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_TRANSACTIONS_AGGREGATED` (-1301208)

| Field | Value |
|-------|-------|
| **Number** | `-1301208` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_TRANSACTIONS_TRANSIENT` (-1301210)

| Field | Value |
|-------|-------|
| **Number** | `-1301210` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_DEVIATION_TRANSACTIONS_OVER_TIME` (-1301212)

| Field | Value |
|-------|-------|
| **Number** | `-1301212` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price deviation TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_TRANSACTIONS_AGGREGATED` (-1301214)

| Field | Value |
|-------|-------|
| **Number** | `-1301214` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_TRANSACTIONS_TRANSIENT` (-1301216)

| Field | Value |
|-------|-------|
| **Number** | `-1301216` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_BOUGHT_TRANSACTIONS_OVER_TIME` (-1301218)

| Field | Value |
|-------|-------|
| **Number** | `-1301218` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded bought TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_TRANSACTIONS_AGGREGATED` (-1301220)

| Field | Value |
|-------|-------|
| **Number** | `-1301220` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_TRANSACTIONS_TRANSIENT` (-1301222)

| Field | Value |
|-------|-------|
| **Number** | `-1301222` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net bought TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_BOUGHT_TRANSACTIONS_OVER_TIME` (-1301224)

| Field | Value |
|-------|-------|
| **Number** | `-1301224` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net bought TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_TRANSACTIONS_AGGREGATED` (-1301226)

| Field | Value |
|-------|-------|
| **Number** | `-1301226` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_TRANSACTIONS_TRANSIENT` (-1301228)

| Field | Value |
|-------|-------|
| **Number** | `-1301228` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_SOLD_TRANSACTIONS_OVER_TIME` (-1301230)

| Field | Value |
|-------|-------|
| **Number** | `-1301230` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded sold TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_TRANSACTIONS_AGGREGATED` (-1301232)

| Field | Value |
|-------|-------|
| **Number** | `-1301232` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_TRANSACTIONS_TRANSIENT` (-1301234)

| Field | Value |
|-------|-------|
| **Number** | `-1301234` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_SELL_TRANSACTIONS_OVER_TIME` (-1301236)

| Field | Value |
|-------|-------|
| **Number** | `-1301236` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_TRANSACTIONS_AGGREGATED` (-1301238)

| Field | Value |
|-------|-------|
| **Number** | `-1301238` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_TRANSACTIONS_TRANSIENT` (-1301240)

| Field | Value |
|-------|-------|
| **Number** | `-1301240` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_TRANSACTIONS_OVER_TIME` (-1301242)

| Field | Value |
|-------|-------|
| **Number** | `-1301242` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_AGGREGATED` (-1301244)

| Field | Value |
|-------|-------|
| **Number** | `-1301244` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open traded net TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_TRANSIENT` (-1301246)

| Field | Value |
|-------|-------|
| **Number** | `-1301246` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open traded net TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRADED_NET_TRANSACTIONS_OVER_TIME` (-1301248)

| Field | Value |
|-------|-------|
| **Number** | `-1301248` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open traded net TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRANSACTIONS_AGGREGATED` (-1301250)

| Field | Value |
|-------|-------|
| **Number** | `-1301250` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total open TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRANSACTIONS_TRANSIENT` (-1301252)

| Field | Value |
|-------|-------|
| **Number** | `-1301252` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total open TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_OPEN_TRANSACTIONS_OVER_TIME` (-1301254)

| Field | Value |
|-------|-------|
| **Number** | `-1301254` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total open TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_TRANSACTIONS_AGGREGATED` (-1301256)

| Field | Value |
|-------|-------|
| **Number** | `-1301256` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_TRANSACTIONS_TRANSIENT` (-1301258)

| Field | Value |
|-------|-------|
| **Number** | `-1301258` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_NET_BUY_TRANSACTIONS_OVER_TIME` (-1301260)

| Field | Value |
|-------|-------|
| **Number** | `-1301260` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total net buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_TRANSACTIONS_AGGREGATED` (-1301262)

| Field | Value |
|-------|-------|
| **Number** | `-1301262` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED traded net sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_TRANSACTIONS_TRANSIENT` (-1301264)

| Field | Value |
|-------|-------|
| **Number** | `-1301264` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT traded net sold TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TRADED_NET_SOLD_TRANSACTIONS_OVER_TIME` (-1301266)

| Field | Value |
|-------|-------|
| **Number** | `-1301266` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled traded net sold TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_TRANSACTIONS_AGGREGATED` (-1301268)

| Field | Value |
|-------|-------|
| **Number** | `-1301268` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_TRANSACTIONS_TRANSIENT` (-1301270)

| Field | Value |
|-------|-------|
| **Number** | `-1301270` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total traded TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_TRADED_TRANSACTIONS_OVER_TIME` (-1301272)

| Field | Value |
|-------|-------|
| **Number** | `-1301272` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total traded TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_TRANSACTIONS_AGGREGATED` (-1301274)

| Field | Value |
|-------|-------|
| **Number** | `-1301274` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED total buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_TRANSACTIONS_TRANSIENT` (-1301276)

| Field | Value |
|-------|-------|
| **Number** | `-1301276` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT total buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_TOTAL_BUY_TRANSACTIONS_OVER_TIME` (-1301278)

| Field | Value |
|-------|-------|
| **Number** | `-1301278` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled total buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_TRANSACTIONS_AGGREGATED` (-1301280)

| Field | Value |
|-------|-------|
| **Number** | `-1301280` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_TRANSACTIONS_TRANSIENT` (-1301282)

| Field | Value |
|-------|-------|
| **Number** | `-1301282` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_BUY_TRANSACTIONS_OVER_TIME` (-1301284)

| Field | Value |
|-------|-------|
| **Number** | `-1301284` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_TRANSACTIONS_AGGREGATED` (-1301286)

| Field | Value |
|-------|-------|
| **Number** | `-1301286` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_TRANSACTIONS_TRANSIENT` (-1301288)

| Field | Value |
|-------|-------|
| **Number** | `-1301288` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net sell TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_SELL_TRANSACTIONS_OVER_TIME` (-1301290)

| Field | Value |
|-------|-------|
| **Number** | `-1301290` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net sell TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_TRANSACTIONS_AGGREGATED` (-1301292)

| Field | Value |
|-------|-------|
| **Number** | `-1301292` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED open net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_TRANSACTIONS_TRANSIENT` (-1301294)

| Field | Value |
|-------|-------|
| **Number** | `-1301294` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT open net buy TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_OPEN_NET_BUY_TRANSACTIONS_OVER_TIME` (-1301296)

| Field | Value |
|-------|-------|
| **Number** | `-1301296` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled open net buy TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_TRANSACTIONS_AGGREGATED` (-1301298)

| Field | Value |
|-------|-------|
| **Number** | `-1301298` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled AGGREGATED price tick deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_TRANSACTIONS_TRANSIENT` (-1301300)

| Field | Value |
|-------|-------|
| **Number** | `-1301300` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled TRANSIENT price tick deviation TRANSACTIONS limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_DISABLED_PRICE_TICK_DEVIATION_TRANSACTIONS_OVER_TIME` (-1301302)

| Field | Value |
|-------|-------|
| **Number** | `-1301302` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | User has breached the disabled price tick deviation TRANSACTIONS OVER_TIME limit |
| **Explanation** | This error code is generated when a user breaches a risk limit |

### `RM_UNSUPPORTED` (-1301304)

| Field | Value |
|-------|-------|
| **Number** | `-1301304` |
| **Subsystem** | PTR |
| **Severity** | 🔴 ERROR |
| **Message** | The operation was denied due to an unsupported request. |
| **Explanation** | The request was not supported by the system, and the operation was denied. |
| **User Action** | None 10 QM Error Messages |

---

## QM Error Messages

### `ME_QM_QR_NOT_ENABLED` (-490000)

| Field | Value |
|-------|-------|
| **Number** | `-490000` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quote Requests not enabled. |

### `ME_QM_QR_NOT_FOUND` (-490002)

| Field | Value |
|-------|-------|
| **Number** | `-490002` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quote Request not found. |

### `ME_QM_MULT_QR` (-490004)

| Field | Value |
|-------|-------|
| **Number** | `-490004` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Multiple Quote Requests not allowed. |

### `ME_QM_QRR_NOT_FOUND` (-490006)

| Field | Value |
|-------|-------|
| **Number** | `-490006` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quote Request Response not found. |

### `ME_QM_DUPL_RESPONSE` (-490008)

| Field | Value |
|-------|-------|
| **Number** | `-490008` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Duplicate Quote Request Response. |

### `ME_QM_MIN_QUANTITY` (-490010)

| Field | Value |
|-------|-------|
| **Number** | `-490010` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quantity is too small. |
| **Explanation** | The quantity in the Quote Request is lower than the MIN quantity. |

### `ME_QM_MATCH_QUANTITY` (-490012)

| Field | Value |
|-------|-------|
| **Number** | `-490012` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quantity is not same as in Quote Request. |
| **Explanation** | The Response Quantity must be the same as the Quote Request Quantity |

### `ME_QM_ACCEPT_QUANTITY` (-490014)

| Field | Value |
|-------|-------|
| **Number** | `-490014` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quantity is lower than in Quote Request. |
| **Explanation** | The Response Quantity must be the same as the Quote Request Quantity |

### `ME_QM_NOT_RECIPIENT` (-490016)

| Field | Value |
|-------|-------|
| **Number** | `-490016` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Quote Request Receiver not found. |
| **Explanation** | The process failed allocating necessary memory |

### `ME_QM_NOT_PARTICIPANT` (-490017)

| Field | Value |
|-------|-------|
| **Number** | `-490017` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | The participant has not been defined. |
| **Explanation** | The participant was not found in the quote monitor. The quote monitor may be corrupt. |
| **User Action** | Please send a description of the circumstances together with log files to the exchange operator. |

### `ME_QM_NOT_INSTRUMENT` (-490019)

| Field | Value |
|-------|-------|
| **Number** | `-490019` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | The given instrument has not been defined (in this instance). |
| **Explanation** | The instrument was not found in the quote monitor instance. The quote monitor may be corrupt. |
| **User Action** | Please send a description of the circumstances together with log files to the exchange operator. |

### `ME_QM_INVALID_SIDE` (-490021)

| Field | Value |
|-------|-------|
| **Number** | `-490021` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Order must specify bid or ask. |
| **Explanation** | The given order does not specify whether it is a bid or ask. The order is ignored by the system. |
| **User Action** | Correct your program. |

### `ME_QM_INVALID_DEAL_SIZE` (-490023)

| Field | Value |
|-------|-------|
| **Number** | `-490023` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid number of deals. |
| **Explanation** | The number of deals must not be zero. |

### `ME_QM_MAX_RECIPIENTS` (-490025)

| Field | Value |
|-------|-------|
| **Number** | `-490025` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Too many recipients specified. |
| **Explanation** | The number of recipients must be less than the maximum allowed number of items. |

### `ME_QM_IQO_NOT_FOUND` (-490027)

| Field | Value |
|-------|-------|
| **Number** | `-490027` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Indicative Quote Offer not found. |

### `ME_QM_IQ_NOT_FOUND` (-490029)

| Field | Value |
|-------|-------|
| **Number** | `-490029` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Indicative Quote not found. |

### `ME_QM_IQ_NOT_ENABLED` (-490030)

| Field | Value |
|-------|-------|
| **Number** | `-490030` |
| **Subsystem** | QM |
| **Severity** | 🔴 ERROR |
| **Message** | Indicative Quote not enabled. |

### `ME_QM_IQ_NO_TR` (-490031)

| Field | Value |
|-------|-------|
| **Number** | `-490031` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | The configuration does not allow this. |
| **Explanation** | The instrument has a configuration that does not allow an indicative quote to be sent in this instrument. |
| **User Action** | Please notify the exchange operator with a description of the problem. 11 QM Error Messages |

### `ME_QM_IDX_INV_EXCHANGE_INFO` (-491001)

| Field | Value |
|-------|-------|
| **Number** | `-491001` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Wrong contents of exchange info |
| **Explanation** | The contents of the exchange info field is wrong |
| **User Action** | You need to change the exchange info field. The contents of the fields for ordersource or settlement method is not in the correct format. |

### `ME_QM_IDX_SETTL_METHOD_NE` (-491003)

| Field | Value |
|-------|-------|
| **Number** | `-491003` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | The settlement method in the indicative quote offer must be equal to the indicative quote responding to. |
| **User Action** | Respecify the quote offer ME_QM_IDX_MAX_TRADE_REPORT_QUANTITY_PERCENTAGE_OF_TRADABLE_ |

### `QUANTITY` (-491005)

| Field | Value |
|-------|-------|
| **Number** | `-491005` |
| **Subsystem** | QM |
| **Severity** | 🟡 WARNING |
| **Message** | Quantity exceeds the max trade report quantity percentage of tradable quantity. |
| **User Action** | Respecify the quote/trade report. 12 SM Error Messages |

---

## SM Error Messages

### `SM_NOT_INITIALIZED` (-1100002)

| Field | Value |
|-------|-------|
| **Number** | `-1100002` |
| **Subsystem** | SM |
| **Severity** | 🔴 ERROR |
| **Message** | The Session Manager is not initialized. |
| **Explanation** | The Session Manager is considered to be initialized when all initial sessions has been sent. |
| **User Action** | Wait until the Session Manager is fully up and running. |

### `SM_INV_ORDER_ACTION` (-1100005)

| Field | Value |
|-------|-------|
| **Number** | `-1100005` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid value for order action. |
| **Explanation** | The value of order action is not correct. |
| **User Action** | Change the order action value to a correct one. |

### `SM_INV_NO_OF_SESSIONS` (-1100007)

| Field | Value |
|-------|-------|
| **Number** | `-1100007` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid number of sessions entered. Must be between 1-15. |
| **Explanation** | The number of sessions entered is not valid. Must be between 1-15. |
| **User Action** | Change the number of override sessions to be between 1-15. |

### `SM_INV_NO_OF_ENTITIES` (-1100009)

| Field | Value |
|-------|-------|
| **Number** | `-1100009` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid number of entities entered. Must be between 1-200. |
| **Explanation** | The number of entities entered is not valid. Must be between 1-200. |
| **User Action** | Change the number of entities. |

### `SM_INV_SESSION_ORDER` (-1100011)

| Field | Value |
|-------|-------|
| **Number** | `-1100011` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | The sessions are not defined in order. Must be ascending order 1, 2, 3 ... 15. |
| **Explanation** | The sessions must be in ascending order 1, 2, 3 ... 15. |
| **User Action** | Change the session order. |

### `SM_INV_END_STATE` (-1100013)

| Field | Value |
|-------|-------|
| **Number** | `-1100013` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | The END-session (session id = 0) must either be the only entry or the last entry in the session sequence. |
| **Explanation** | The END-session (session id = 0) must either be the only entry or the last entry in the session sequence, and can only occur 1 time. |
| **User Action** | Set the END-session to be the only or the last entry in the session sequence, or remove it all together. |

### `SM_INV_DEFAULT_HALT` (-1100015)

| Field | Value |
|-------|-------|
| **Number** | `-1100015` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid session id when requesting default halt. Session id must be zero. |
| **Explanation** | The session id must be zero when requesting default halt. |
| **User Action** | Change session id to zero or remove the default halt flag. |

### `SM_INV_SESSION_ID` (-1100017)

| Field | Value |
|-------|-------|
| **Number** | `-1100017` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid session id. |
| **Explanation** | The specified session id doesn't correspond to an existing session. |
| **User Action** | Change session id. |

### `SM_INV_TIME_ORDER` (-1100019)

| Field | Value |
|-------|-------|
| **Number** | `-1100019` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Session states are not defined in time order. |
| **Explanation** | All session states must be in the future and defined in order of time. |
| **User Action** | Change session state start times. |

### `SM_INV_SESSION_LEVEL` (-1100021)

| Field | Value |
|-------|-------|
| **Number** | `-1100021` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid session level. |
| **Explanation** | The provided session level doesn't match the given entity or is invalid. |
| **User Action** | Change to a valid session level for override sessions and make sure it matches the given entity. |

### `SM_MARKET_OVERRIDE_NOT_ALLOWED` (-1100023)

| Field | Value |
|-------|-------|
| **Number** | `-1100023` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Override is not allowed for this market / market segment, because it doesn't have a direct connection to a trading session. |
| **Explanation** | In order to enter an override for a specific market / market segment, a trading session must be configured on that level. |
| **User Action** | Make sure the market / market segment have a direct connection to a trading session. |

### `SM_DEFAULT_HALT_NOT_ALLOWED` (-1100025)

| Field | Value |
|-------|-------|
| **Number** | `-1100025` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Default halt is not allowed for orderbook and asset. |
| **Explanation** | A default halt session id is only allowed for market or market segment, and must be configured. |
| **User Action** | Disable the default halt flag on the ingoing message and specify a session id. |

### `SM_ENTITY_NOT_FOUND` (-1100027)

| Field | Value |
|-------|-------|
| **Number** | `-1100027` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Entity not found. |
| **Explanation** | At least one of the given entities doesn't exist in the reference data. |
| **User Action** | Configure the entity or chose another one. |

### `SM_ACTOR_NOT_FOUND` (-1100029)

| Field | Value |
|-------|-------|
| **Number** | `-1100029` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Actor not found. |
| **Explanation** | The given actor doesn't exist in the reference data. |
| **User Action** | Configure the actor or chose another one. |

### `SM_NO_OVERRIDE_TO_END` (-1100031)

| Field | Value |
|-------|-------|
| **Number** | `-1100031` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | At least one of the specified markets / market segments doesn't have an override session to end. |
| **Explanation** | An existing override session must exist in order to end it. |
| **User Action** | Only send END-instruction for overrides that exist. |

### `SM_INVALID_OVERRIDE_NO` (-1100033)

| Field | Value |
|-------|-------|
| **Number** | `-1100033` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | The given override number is not the most recent. Use the most recent number or specify zero to bypass the validation. |
| **Explanation** | The most recent disseminated override number must be used in order to enter the override. The validation can be bypassed by entering value zero. |
| **User Action** | Use the most recent disseminated override number when sending the override, or set it to zero to bypass the validation. |

### `SM_DUPLICATE_ENTITIES` (-1100035)

| Field | Value |
|-------|-------|
| **Number** | `-1100035` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to have duplicate entities. |
| **Explanation** | The override contains the same entity more then once. |
| **User Action** | Remove any duplicate entities and try again. |

### `SM_SUSPENDED_SECURITY` (-1100039)

| Field | Value |
|-------|-------|
| **Number** | `-1100039` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | The override cannot be applied because the security is suspended. |

### `SM_END_OF_BUSINESS_ALREADY_REACHED` (-1100041)

| Field | Value |
|-------|-------|
| **Number** | `-1100041` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to send end of business at this time, because it has already been reached. |
| **Explanation** | End of business has already have been reached. |

### `SM_INV_END_OF_BUSINESS_RESPONSE` (-1100043)

| Field | Value |
|-------|-------|
| **Number** | `-1100043` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to perform end of business at this time, because the end of business request was returned with a negative response. |
| **Explanation** | The end of business request was returned with a negative response. |

### `SM_END_OF_BUSINESS_NOT_REQUESTED` (-1100045)

| Field | Value |
|-------|-------|
| **Number** | `-1100045` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to perform end of business at this time, because the end of business request has not yet been sent. |
| **Explanation** | The end of business request has not yet been sent. |

### `SM_INV_START_OFFSET` (-1100047)

| Field | Value |
|-------|-------|
| **Number** | `-1100047` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Start offsets are not allowed to be set on schedules. |
| **Explanation** | Start offsets must not be set on a schedule's sessions. |
| **User Action** | Remove start offsets from the sessions. |

### `SM_INV_RANDOM_INTERVAL` (-1100049)

| Field | Value |
|-------|-------|
| **Number** | `-1100049` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid random interval. |
| **Explanation** | Random interval must be non-negative. |
| **User Action** | Correct the random intervals on the sessions. |

### `SM_RANDOM_INTERVAL_NOT_ALLOWED` (-1100051)

| Field | Value |
|-------|-------|
| **Number** | `-1100051` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Random interval are not allowed on the session. |
| **Explanation** | The first session in a schedule must not have a random interval set. |
| **User Action** | Remove the random interval on the schedule's first session. |

### `SM_INV_AUCTION_TIME` (-1100053)

| Field | Value |
|-------|-------|
| **Number** | `-1100053` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid auction time for the one sided auction. |
| **Explanation** | The auction time of a one sided auction needs to be a valid time that fits the auction sequence. |
| **User Action** | Set the auction time to a valid time that fits the auction sequence. |

### `SM_NOT_ISSUER` (-1100055)

| Field | Value |
|-------|-------|
| **Number** | `-1100055` |
| **Subsystem** | SM |
| **Severity** | 🟡 WARNING |
| **Message** | Transaction not allowed for non-issuer actor. |
| **Explanation** | The transaction is only allowed for actors that are under the issuer of the order book. |
| **User Action** | None 13 TH Error Messages |

---

## TH Error Messages

### `TH_INV_ORDERBOOK` (-700001)

| Field | Value |
|-------|-------|
| **Number** | `-700001` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Order Book. |
| **Explanation** | The specified order book can not be found. |

### `TH_COMBINATION_ORDERBOOK` (-700003)

| Field | Value |
|-------|-------|
| **Number** | `-700003` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | Not allowed to cancel deal for combination order book. |

### `TH_INV_USER` (-700005)

| Field | Value |
|-------|-------|
| **Number** | `-700005` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid User. |
| **Explanation** | The specified user can not be found. |

### `TH_INV_PARTICIPANT` (-700007)

| Field | Value |
|-------|-------|
| **Number** | `-700007` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | Invalid Participant. |
| **Explanation** | The specified participant can not be found. |

### `TH_INV_DEAL_ID` (-700009)

| Field | Value |
|-------|-------|
| **Number** | `-700009` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | Deal not found. |
| **Explanation** | The deal with the given deal id can not be found. |

### `TH_DEAL_CANCELLED` (-700013)

| Field | Value |
|-------|-------|
| **Number** | `-700013` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | The deal has already been cancelled. |

### `TH_DEAL_RELEASED` (-700015)

| Field | Value |
|-------|-------|
| **Number** | `-700015` |
| **Subsystem** | TH |
| **Severity** | 🟡 WARNING |
| **Message** | The deal has already been released. |

---

*© 2025 Indonesia Stock Exchange — Private and Confidential*

*Disclaimer: The information in this document may change in accordance with the progress of the ongoing project development.*