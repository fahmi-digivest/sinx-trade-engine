# IDX DOCUMENT SPECIFICATION
## OUCH Protocol Specification

**Document Version:** 2.0  
**Date:** 29 Aug 2025  
**Issuer:** Indonesia Stock Exchange (Bursa Efek Indonesia)

> **Disclaimer:** The information in this document may change in accordance with the progress of the ongoing project development.  
> © 2025 Indonesia Stock Exchange — *Private and Confidential*

---

## Summary of Changes

| Version | Date | List of Changes |
|---------|------|-----------------|
| 2.0 | 29/08/2025 | • Removed timeInForce valid values: `2 = GTC` and `6 = Days`<br>• Added description for Exchange Info (Table 2)<br>• Added Order Source valid value (Section 3.1.2)<br>• Added **Client Account** field description (Section 3.1.3)<br>• Removed OrderType valid values: `4 = BestOrder` and `5 = Imbalance` |
| 1.0 | 05/03/2025 | First release |

---

## Table of Contents

1. [About This Manual](#1-about-this-manual)
2. [OUCH Overview](#2-ouch-overview)
   - 2.1 [Architecture](#21-architecture)
   - 2.2 [Data Types](#22-data-types)
   - 2.3 [Fault Redundancy](#23-fault-redundancy)
   - 2.4 [Operating in 24x7 mode](#24-operating-in-24x7-mode)
   - 2.5 [Quantities in Units or Lots](#25-quantities-in-units-or-lots)
3. [Message Formats](#3-message-formats)
   - 3.1 [Data Structures](#31-data-structures)
     - 3.1.1 [Exchange Info](#311-exchange-info)
     - 3.1.2 [Order Source Valid Value](#312-order-source-valid-value)
     - 3.1.3 [Client Account Field](#313-client-account-field)
   - 3.2 [Inbound Messages](#32-inbound-messages)
     - 3.2.1 [Enter Order](#321-enter-order)
     - 3.2.2 [Replace Order](#322-replace-order)
     - 3.2.3 [Cancel Order](#323-cancel-order)
     - 3.2.4 [Cancel by Order ID](#324-cancel-by-order-id)
   - 3.3 [Outbound Messages](#33-outbound-messages)
     - 3.3.1 [Order Accepted](#331-order-accepted)
     - 3.3.2 [Order Rejected](#332-order-rejected)
     - 3.3.3 [Order Replaced](#333-order-replaced)
     - 3.3.4 [Order Canceled](#334-order-canceled)
     - 3.3.5 [Order Executed](#335-order-executed)

---

## 1 About This Manual

The purpose of this document is to describe the Nasdaq OUCH protocol.

### 1.1 References

For more information, refer to the following documents:

- Nasdaq ITCH Protocol Specification for MME Core
- Nasdaq Transport Protocol Overview Core

---

## 2 OUCH Overview

The OUCH protocol accepts orders from system participants and executes matching orders when possible. Non-matching orders may be added to the order book where they are waiting to be matched according to the matching priority model.

OUCH is a simple protocol that allows Nasdaq Financial Framework users to enter orders, replace and cancel existing orders and receive executions. It is intended to allow participants and their software developers to integrate Nasdaq Financial Framework into their proprietary trading systems or to build custom front ends.

OUCH only provides a method for participants to send orders to Nasdaq Financial Framework and receive updates on those orders entered. For information about all orders entered into and executed on the book, refer to the ITCH protocol specification.

OUCH is the low-level native protocol for connecting to the Nasdaq Financial Framework system. It is designed to offer the maximum possible performance at the cost of flexibility and ease of use. For applications that do not require this extreme level of performance, Nasdaq Financial Framework offers other, more standard interfaces that may be more suitable and easier to develop to.

### 2.1 Architecture

The OUCH protocol is composed of logical messages passed between the OUCH host and the client application. Each message type has a fixed message length. The messages are binary encoded, which means that all numeric values are represented as binary values. Character or string values are composed of non-control ISO 8859-1 (Latin-1) encoded bytes.

All (outbound) messages sent from the OUCH system to the client are assumed to be sequenced, and their delivery is guaranteed by the lower level protocol. The SoupBinTCP protocol (specification available separately) is used to guarantee the delivery and sequencing of OUCH messages sent from the host to the client. Please refer to the Nasdaq Transport Protocol Overview Core for details on the SoupBinTCP protocol.

Messages sent from the OUCH client to the host are inherently non-guaranteed, even if they are carried by a lower level protocol that guarantees delivery (like TCP/IP sockets). Therefore, all host-bound messages are designed so that they can be benignly resent for robust recovery from connection and application failures.

Each physical OUCH host port is bound to an OUCH Account assigned by the marketplace. On a given day, every order entered on OUCH is uniquely identified by the combination of the logical OUCH Account and the participant-created Token field.

### 2.2 Data Types

All Integer fields are composed of binary encoded numbers.

**Table 1: List of data types**

| Type | Size | Description |
|------|------|-------------|
| int8 | 1 byte | 8 bit Signed Big Endian |
| int16 | 2 bytes | 16 bit Signed Big Endian |
| int32 | 4 bytes | 32 bit Signed Big Endian |
| int64 | 8 bytes | 64 bit Signed Big Endian |
| Alpha | variable | Composed of non-control ISO 8859-1 (Latin-1) encoded bytes. Left justified and padded on the right with spaces. |
| Price | 8 bytes | Prices are signed numeric fields. Number of decimals is specified in the Order book Directory message. A Price field with the minimum long value indicates that no price was available. |
| Timestamp | 8 bytes | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC). |

### 2.3 Fault Redundancy

A single OUCH Account can be bound to two physical OUCH gateways. These OUCH gateways then act as mirrors of each other for fault redundancy. In this configuration, the client can connect to any one of the gateways. It is not allowed to be logged on to both gateways simultaneously. The system will log out the first client session when a second is established for the same account.

The OUCH Gateway does not support nor block another OUCH Account accessing the standby gateway port of an already engaged primary gateway port.

### 2.4 Operating in 24x7 mode

Once every 24 hours the Ouch stream will be scrapped of old data and a snapshot is created. SoupBinTCP announces this by sending an unsolicited login accept with sequence number 1 to all clients. This is an instruction to clients to get rid of all states, i.e. open orders. In Ouch a Replace Order message will be sent for each open order on the order book at sequence reset.

### 2.5 Quantities in Units or Lots

All quantities fields have quantities expressed in units or lots. The 'Quantity Expressed In' field in ITCH Order Book Directory message is used to determine if quantities for an order book are expressed in units or lots.

---

## 3 Message Formats

### 3.1 Data Structures

The contents of common data structures used in different messages are outlined below.

#### 3.1.1 Exchange Info

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| orderSource | 0 | 4 | Alpha | The order source refers to the type of trading activity from which the order originated. E.g. online trading or sharia trading.<br>Note: First character is mandatory.<br>Allowed characters:<br>Position 1: 0–9, A–Z<br>Position 2–4: 0–9, A–Z and Space |
| settlementMethod | 4 | 1 | Alpha | Settlement type for settlement instruction.<br>Values:<br>`1` = Delivery vs Payment (DVP)<br>`2` = Delivery Free Of Payment (DFOP) |
| Reserved | 5 | 27 | — | Reserved for future use |

#### 3.1.2 Order Source Valid Value

| Digit 1 (application) | Digit 2 (feature) | Digit 3 (platform) | Digit 4 (AO) |
|-----------------------|-------------------|--------------------|--------------|
| R = Remote Trading | E = ETF | A = Desktop | G = Automated Ordering |
| O = Online Trading | S = Sharia | B = Web base | Z = Other |
| D = DMA | P = MSOP/ESOP | F = Mobile | |
| M = MPPE | C = One Day Trade | Z = Other | |
| Z = Other | Z = Other | | |

#### 3.1.3 Client Account Field

The valid value for the Client Account field is the Single Investor Identification (SID). The Client Account field (16 characters) shall be validated according to the following rules:

- The field is **mandatory**.
- The length must be **exactly 15 characters**. The first 15 characters are used, and the 16th character shall be a null character.
- Only the following characters are allowed: `0–9`, `A–Z`, and `a–z`. Both uppercase and lowercase are allowed.
- The **3rd character** must be `D`, `d`, `F`, or `f`. `D`/`d` indicates Domestic, and `F`/`f` indicates Foreign. Both uppercase and lowercase are allowed.

---

### 3.2 Inbound Messages

Inbound messages are sent from the participant's application to the OUCH host. They are not sequenced. All Inbound Messages may be repeated benignly. This gives the client the ability to resend any inbound message if, in the case of a connection loss or an application error, it is uncertain whether or not the Nasdaq Financial Framework system received it.

The idea of benign inbound message retransmission with end-to-end acknowledgement is fundamental to OUCH's fail-over redundancy. If your connection ever fails, there is no way for you to know if pending messages actually made it over the link before the failure. A robust OUCH client can safely resend any pending messages over a mirrored link without worrying about generating duplicates. This applies to Nasdaq Financial Framework's disaster failover capability as well; if the system ever needs to fail over to the backup site, some messages sent at the moment of the failure may be lost. A robust application can simply resend the pending messages, making the failover seamless to the end user.

All inbound messages on an OUCH port are processed sequentially. This guarantees that if two orders are entered consecutively on the same connection, the first order entered will always be accepted first.

#### 3.2.1 Enter Order

Enter Order is used to enter a new order into the system. The response to a successful Enter Order is an Order Accepted message. If the order is rejected, the Order Rejected message will be returned.

**Table 2: Enter Order Message**

**OuchEnterOrder [O] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `O` — Enter Order Message |
| orderToken | 1 | 8 | int64 | Client generated order identifier |
| orderBookId | 9 | 4 | int32 | Order book identifier |
| side | 13 | 1 | Alpha | Supported values:<br>`B` = Buy<br>`S` = Sell<br>`T` = ShortSell |
| quantity | 14 | 8 | int64 | The quantity of the order |
| price | 22 | 8 | Price | Signed long price. Number of decimals and allowed tick steps are given by the Order book Directory message in ITCH. To signal a market price — Minimum Long value shall be sent. See section Data Types for Price. |
| timeInForce | 30 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Day<br>`3` = Fak<br>`4` = Fok<br>`5` = GTS |
| openClose | 31 | 1 | int8 | Specifies the requested position handling, such as closing or opening position.<br>Supported values:<br>`0` = Default<br>`1` = Open<br>`2` = Close/Net<br>`3` = Mandatory close |
| clientAccount | 32 | 16 | Alpha | See 3.1.3. The name of the Account the order is entered from. The first 15 characters are used, and the last 16th character shall be a null character. |
| customerInfo | 48 | 15 | Alpha | Pass-thru field |
| exchangeInfo | 63 | 32 | Alpha | See 3.1.1 & 3.1.2. This field is used exclusively to carry the `orderSource` value. All other fields within Exchange Info are not applicable and must be left empty (default). |
| displayQuantity | 95 | 8 | int64 | Display quantity if reserve order, otherwise set to zero (`0`) |
| orderType | 103 | 1 | int8 | Defines type of order submitted.<br>Supported values:<br>`1` = Limit<br>`2` = Market<br>`3` = MarketToLimit |
| timeInForceData | 104 | 2 | int16 | For Time In Force GTS and GTD: Dependent on the Time In Force setting for the order. For certain Time In Force, this field is needed to specify for how long the order shall be valid.<br>For a Good till Days order (GTD) the data shall be specified as number of days (e.g. 5 days). The maximum value is defined by the RDM configuration, Max time in order book.<br>For a Good till Session order (GTS) the data shall be specified as the session state type id (e.g. 1 = Pre-morning).<br>For order types that do not need this information, this field is left empty. |
| orderCapacity | 106 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Agency<br>`2` = Proprietary<br>`3` = Individual<br>`4` = Principal<br>`5` = RiskLessPrincipal |
| selfMatchPreventionKey | 107 | 4 | int32 | When set, orders from the same participant with equal keys shall not be matched. Values: `0` = No key. In case a participant is not configured to use Crossing key, and a key is set, the system returns 0. |
| attributes | 111 | 2 | int16 | Supported values:<br>`0` = Undefined<br>`1` = MarketBid<br>`2` = PriceStabilization<br>`3` = Margin |

---

#### 3.2.2 Replace Order

Replace Order is used to modify an existing order entered via OUCH.

Modification of some order parameters may not be allowed, depending on how the system is configured. An Order Rejected will be returned in such case.

There are two order tokens in the Replace message:

- **Existing Order Token** — used to reference the order to be replaced. The Order Token should be from the original Enter Order, not from any intermediate replaces. The current implementation allows intermediate tokens to be used, but this may not be supported in the future.
- **Replacement Order Token** — the new Order Token that will be assigned to the order if the replace is successful. The replacement Order Token must not be a token previously used in Enter Order or Replace Order transactions.

The response to a Replace Order is:

- **Order Replaced** if the modification was successful. The Order Replaced will contain the current state of the returned order. See below for a discussion on order quantities.
- **Order Rejected** if the replace failed.

**Table 3: Replace Order Message**

**OuchReplaceOrder [U] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `U` — Replace Order Message |
| existingOrderToken | 1 | 8 | int64 | The Existing Order Token is used to reference the order to be replaced. The Order Token should be the original Order Token, or the latest successful Replacement Order Token, not any intermediate Replacement Order Token. |
| replacementOrderToken | 9 | 8 | int64 | The Replacement Order Token is the new Order Token that will be assigned to the order if the replace is successful. The replacement Order Token must not be a token previously used in Enter Order or Replace Order transactions. |
| quantity | 17 | 8 | int64 | Desired Open Quantity of the order, in relation to the original order quantity. |
| price | 25 | 8 | Price | The updated/new price. MIN LONG value corresponds to "no change." See section Data Types for Price. The number of decimals and allowed tick steps are given by the Order book Directory message in ITCH. |
| openClose | 33 | 1 | int8 | Specifies the requested position handling, such as closing or opening position.<br>Supported values:<br>`0` = Default<br>`1` = Open<br>`2` = Close/Net<br>`3` = Mandatory close |
| clientAccount | 34 | 16 | Alpha | The name of the requested new Account for the order. If empty, no change shall be done to the Account for the order. The first 15 characters are used, and the last 16th character shall be a null character. |
| customerInfo | 50 | 15 | Alpha | Pass-thru field |
| exchangeInfo | 65 | 32 | Alpha | Pass-thru field |
| displayQuantity | 97 | 8 | int64 | Desired displayed quantity (zero for unchanged) |
| timeInForce | 105 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Day<br>`3` = Fak<br>`4` = Fok<br>`5` = GTS |
| timeInForceData | 106 | 2 | int16 | For Time In Force GTS and GTD: Dependent on the Time In Force setting for the order. For certain Time In Force, this field is needed to specify for how long the order shall be valid.<br>For GTD: specify number of days (e.g. 5 days).<br>For GTS: specify session state type id (e.g. 1 = Pre-morning).<br>For order types that do not need this information, this field is left empty. |
| selfMatchPreventionKey | 108 | 4 | int32 | When set, orders from the same participant with equal keys shall not be matched. Values: `0` = No key. In case a participant is not configured to use Crossing key, and a key is set, the system returns 0. |

##### 3.2.2.1 Order Quantities

In Nasdaq OUCH Replace messages, the Quantity field contains the desired total quantity of the order (Order Quantity = open quantity + executed quantity, where open quantity is the quantity of the order in the order book).

> **Example 1:**
> 1. An order with a quantity of 1000 is entered via OUCH. An Order Accepted with Quantity = 1000 will be returned.
> 2. A partial execution for 200 occurs. A quantity of 800 is left in the order book. An Executed Order with Traded Quantity = 200 will be returned.
> 3. The client wants to decrease the open quantity with 50, that is, from the current 800 to 750. He sends in an Order Replace with Quantity = 950, as a quantity update is always done in relation to the original order quantity. A Replaced Order with Quantity = 750 will be returned.

> **Example 2:**
> 1. An order with a quantity of 1000 is entered via OUCH. An Order Accepted with Quantity = 1000 will be returned.
> 2. The client wants to decrease the open quantity with 600, that is, from the current 1000 to 400. He sends in an Order Replace with Quantity = 400.
> 3. Simultaneously, 500 of the order is (partially) traded, the trade occurred before the Order Replace. An Executed Order with Traded Quantity = 500 will be returned.
> 4. The client receives an Order Replaced with Quantity = 0 and Order Status = Not on book, since there is no quantity left in the order book.

> **Example 3:**
> A reserve order with a total quantity of 1000 and a display quantity of 100 is entered via OUCH. An Order Accepted with Quantity = 1000 and Display Quantity = 100 will be returned. A partial execution for 150 occurs. A quantity of 850 is left in the order book, with a display quantity of 50. Two Executed Orders, one with Traded Quantity = 100 and one with Traded Quantity = 50 will be returned. The client wants to decrease the total quantity with 50, i.e. from the current 850 to 800. At the same time the client wants to increase the displayed quantity to 150. The client sends in an Order Replace with Quantity = 950 and a Display Quantity = 150. A Replaced Order with Quantity = 800 and Display Quantity = 150 will be returned.

> **Note:** The system may be configured to reject replaces that would cause a loss of priority.

##### 3.2.2.2 Retaining Order book Priority

Most order attributes are allowed to change in Nasdaq Financial Framework. A replacement of a certain attribute may cause the order to lose priority, depending on what values are modified, how the system is configured, and so on.

In order to increase the likelihood of keeping an order from losing priority as it is changed, all fields that are intended to be left unchanged should be left unset:

- For **Numeric** Order parameters, this means setting them to `0` (zero).
- For **String (Alpha)** fields, the first byte should be set to binary zero `'\0'` to retain previous value.

---

#### 3.2.3 Cancel Order

Partial cancels are not supported with the Cancel Order message. Use Order Replace to modify an existing order.

The response to a successful Cancel Order is an Order Canceled message. Failed Cancel Order messages are rejected with the Order Reject message.

**Table 4: Cancel Order Message**

**OuchCancelOrder [X] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `X` — Cancel Order Message |
| orderToken | 1 | 8 | int64 | The Order Token is used to reference the order to be canceled. |

---

#### 3.2.4 Cancel by Order ID

Using the system-generated Order ID, this message can be used to cancel any order in the book regardless of session during which it was inserted.

The response to a successful Cancel by Order ID is an Order Canceled message. Failed Cancel by Order ID messages are rejected with the Order Reject message.

**Table 5: Cancel By Order ID Message**

**OuchCancelOrderById [Y] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `Y` — Cancel By Order ID Message |
| orderBookId | 1 | 4 | int32 | Identifier for the Order book |
| side | 5 | 1 | Alpha | Supported values:<br>`B` = Buy<br>`S` = Sell<br>`T` = ShortSell |
| orderId | 6 | 8 | int64 | The identifier assigned to the order by the system |

---

### 3.3 Outbound Messages

Outbound messages are generated by the system and sent to the OUCH client.

> **Note:** If the order is traded or otherwise changed immediately at entry, the Order Accepted will show the state of the order after any such operations. Execution messages and/or cancel messages will follow to account for any differences.

---

#### 3.3.1 Order Accepted

This message acknowledges the receipt and acceptance of a valid Enter Order Message. The data fields from the Enter Order Message are echoed back in this message. Note that the accepted values may differ from the entered values for some fields.

Accepted Messages are guaranteed to come before any Executed Messages or Canceled Messages for an order.

**Table 6: Order Accepted Message**

**OuchOrderAccepted [A] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `A` — Order Accepted Message |
| timestamp | 1 | 8 | Timestamp | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |
| orderToken | 9 | 8 | int64 | The Token for the accepted order |
| orderBookId | 17 | 4 | int32 | Order book identifier |
| side | 21 | 1 | Alpha | Supported values:<br>`B` = Buy<br>`S` = Sell<br>`T` = ShortSell |
| orderId | 22 | 8 | int64 | The identifier assigned to the new order. Note that the number is only unique per Order book and side. |
| quantity | 30 | 8 | int64 | Quantity currently open in the book |
| price | 38 | 8 | Price | Signed integer price. Number of decimals is given by the Order book Directory message in ITCH. |
| timeInForce | 46 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Day<br>`3` = Fak<br>`4` = Fok<br>`5` = GTS |
| openClose | 47 | 1 | int8 | Specifies the requested position handling, such as closing or opening position.<br>Supported values:<br>`0` = Default<br>`1` = Open<br>`2` = Close/Net<br>`3` = Mandatory close |
| clientAccount | 48 | 16 | Alpha | The name of the Account the order is entered from. The first 15 characters are used, and the last 16th character shall be a null character. |
| orderState | 64 | 1 | int8 | Supported values:<br>`1` = On book<br>`2` = Not on book |
| customerInfo | 65 | 15 | Alpha | Pass-thru field |
| exchangeInfo | 80 | 32 | Alpha | Pass-thru field |
| displayQuantity | 112 | 8 | int64 | Display quantity if reserve order, otherwise set to zero (`0`) |
| orderType | 120 | 1 | int8 | Defines type of order submitted.<br>Supported values:<br>`1` = Limit<br>`2` = Market<br>`3` = MarketToLimit |
| timeInForceData | 121 | 2 | int16 | For Time In Force GTS and GTD: Dependent on the Time In Force setting for the order. For certain Time In Force, this field is needed to specify for how long the order shall be valid.<br>For GTD: number of days (e.g. 5).<br>For GTS: session state type id (e.g. 1 = Pre-morning).<br>For order types that do not need this information, this field is left empty. |
| orderCapacity | 123 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Agency<br>`2` = Proprietary<br>`3` = Individual<br>`4` = Principal<br>`5` = RiskLessPrincipal |
| selfMatchPreventionKey | 124 | 4 | int32 | When set, orders from the same participant with equal keys shall not be matched. Values: `0` = No key. In case a participant is not configured to use Crossing key, and a key is set, the system returns 0. |
| attributes | 128 | 2 | int16 | Supported values:<br>`0` = Undefined<br>`1` = MarketBid<br>`2` = PriceStabilization<br>`3` = Margin |

---

#### 3.3.2 Order Rejected

This message is used to reject Enter Order messages, Cancel Order messages, and Replace Order messages.

**Table 7: Order Rejected Message**

**OuchOrderRejected [J] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `J` — Order Rejected Message |
| timestamp | 1 | 8 | Timestamp | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |
| orderToken | 9 | 8 | int64 | The Token for the rejected order |
| orderId | 17 | 8 | int64 | The Order ID for the rejected order |
| rejectCode | 25 | 4 | int32 | Backend Error Code. See System Error Messages Reference for more information. |

---

#### 3.3.3 Order Replaced

This message acknowledges the receipt and acceptance of a valid Replace Order Message. The data fields from the Replace Order Message are echoed back in this message. Note that the accepted values may differ from the entered values for some fields. You will receive one and only one of these two for each replacement.

Like Accepted Messages, Replaced Messages use the Order State field to denote that a replace was accepted and then automatically canceled when the Order State is Not on book. No further Executed Messages nor Canceled Messages will be received when the replaced order for the Order State is Not on book.

> **Note:** If an order is modified by the system, the Replacement Order Token will be blanked out (with spaces). The Previous Order Token will contain the latest successful Order Token as entered by the client.

**Table 8: Order Replaced Message**

**OuchOrderReplaced [U] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `U` — Order Replaced Message |
| timestamp | 1 | 8 | Timestamp | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |
| replacementOrderToken | 9 | 8 | int64 | The new Order Token for the order |
| previousOrderToken | 17 | 8 | int64 | The previous Order Token for the order |
| orderBookId | 25 | 4 | int32 | Order book identifier |
| side | 29 | 1 | Alpha | Supported values:<br>`B` = Buy<br>`S` = Sell<br>`T` = ShortSell |
| orderId | 30 | 8 | int64 | The identifier assigned to the new order. Note that the number is only unique per Order book and side. |
| quantity | 38 | 8 | int64 | Quantity currently open in the book |
| price | 46 | 8 | Price | Signed integer price. Number of decimals is given by the Order book Directory message in ITCH. |
| timeInForce | 54 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Day<br>`3` = Fak<br>`4` = Fok<br>`5` = GTS |
| openClose | 55 | 1 | int8 | Specifies the requested position handling, such as closing or opening position.<br>Supported values:<br>`0` = Default<br>`1` = Open<br>`2` = Close/Net<br>`3` = Mandatory close |
| clientAccount | 56 | 16 | Alpha | The name of the Account. The first 15 characters are used, and the last 16th character shall be a null character. |
| orderState | 72 | 1 | int8 | Supported values:<br>`1` = On book<br>`2` = Not on book |
| customerInfo | 73 | 15 | Alpha | Pass-thru field |
| exchangeInfo | 88 | 32 | Alpha | Pass-thru field |
| displayQuantity | 120 | 8 | int64 | The displayed quantity in case the order is a reserve order or zero if the order is not a reserve order |
| orderType | 128 | 1 | int8 | Supported values:<br>`1` = Limit<br>`2` = Market<br>`3` = MarketToLimit |
| timeInForceData | 129 | 2 | int16 | For Time In Force GTS and GTD: Dependent on the Time In Force setting for the order. For certain Time In Force, this field is needed to specify for how long the order shall be valid.<br>For GTD: number of days (e.g. 5).<br>For GTS: session state type id (e.g. 1 = Pre-morning).<br>For order types that do not need this information, this field is left empty. |
| orderCapacity | 131 | 1 | int8 | Supported values:<br>`0` = Undefined<br>`1` = Agency<br>`2` = Proprietary<br>`3` = Individual<br>`4` = Principal<br>`5` = RiskLessPrincipal |
| selfMatchPreventionKey | 132 | 4 | int32 | When set, orders from the same participant with equal keys shall not be matched. Values: `0` = No key. In case a participant is not configured to use Crossing key, and a key is set, the system returns 0. |
| attributes | 136 | 2 | int16 | Supported values:<br>`0` = Undefined<br>`1` = MarketBid<br>`2` = PriceStabilization<br>`3` = Margin |

---

#### 3.3.4 Order Canceled

A Canceled Message informs you that an order has been canceled. This could be acknowledging a Cancel Order Message, or it could be the result of system cancellation of the order.

> **Note:** Order Canceled messages are sent out when orders are suspended due to connection loss. Orders cannot be reactivated again, but it is possible to cancel a suspended order, so therefore it is possible to receive more than one Order Canceled message for the same order.

**Table 9: Order Canceled Message**

**OuchOrderCanceled [C] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `C` — Order Canceled Message |
| timestamp | 1 | 8 | Timestamp | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |
| orderToken | 9 | 8 | int64 | Order Token for the Canceled order |
| orderBookId | 17 | 4 | int32 | Order book identifier |
| side | 21 | 1 | Alpha | Supported values:<br>`B` = Buy<br>`S` = Sell<br>`T` = ShortSell |
| orderId | 22 | 8 | int64 | The identifier assigned to the canceled order. Note that the number is only unique per Order book and side. |
| cancelReason | 30 | 1 | int8 | Supported values:<br>`1` = Canceled by user<br>`3` = Canceled by system after trade (ex FAK order)<br>`6` = Canceled by system after new order entry (ex FOK/FAK order)<br>`8` = Canceled by system after order converted (ex MTL order)<br>`9` = Canceled by system<br>`10` = Canceled by proxy user<br>`12` = Canceled by system after new order triggered<br>`13` = Canceled by system for hidden order<br>`19` = Canceled by system order changed<br>`20` = Canceled by system due to Instrument Session State<br>`43` = Canceled due to Self Match Prevention<br>`44` = Canceled due to Circuit Breaker<br>`45` = Canceled due to credit limits<br>`58` = Canceled due to Corporate Action<br>`59` = Canceled due to Self Match Prev. — Default (cancel aggressive) |

---

#### 3.3.5 Order Executed

This message is returned when a partial or full fill occurs.

**Table 10: Order Executed Message**

**OuchOrderExecuted [E] — Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| ouchMessageType | 0 | 1 | int8 | `E` — Order Executed Message |
| timestamp | 1 | 8 | Timestamp | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |
| orderToken | 9 | 8 | int64 | Token identifier for the matched order |
| orderBookId | 17 | 4 | int32 | Order book identifier |
| tradeQuantity | 21 | 8 | int64 | The matched quantity |
| tradePrice | 29 | 8 | Price | Signed integer trade price. Number of decimals is given by the Order book Directory message in ITCH. |
| matchId | 37 | 8 | int64 | Backend generated identifier. Assigned by the system to each match executed. |
| comboGroupId | 45 | 4 | int32 | Used to group combo and leg executions together |
| dealSource | 49 | 1 | int8 | The deal source of the order executed.<br>Supported values:<br>`1` = Matched by system automatically<br>`7` = Combination order match combination order automatically<br>`20` = Matched in auction<br>`36` = Tailor Made Combination<br>`43` = Combination matched outright legs |

---

*© 2025 Indonesia Stock Exchange — Private and Confidential*  
*Version 2.0 | 29 Aug 2025*
