# IDX ITCH and ITCH-MDF MME Protocol Specification

**Document Version:** 4.0  
**Date:** 18 September 2025  
**Code:** A  
**Classification:** Private and Confidential  

© 2025 Indonesia Stock Exchange

> **Disclaimer:** The information in this document may change in accordance with the progress of the ongoing project development.

---

## Table of Contents

1. [Summary of Changes](#1-summary-of-changes)
2. [About This Manual](#2-about-this-manual)
3. [ITCH Overview](#3-itch-overview)
   - 3.1 [Architecture](#31-architecture)
   - 3.2 [Data Types](#32-data-types)
   - 3.3 [Quantities in Units or Lots](#33-quantities-in-units-or-lots)
4. [Message Formats](#4-message-formats)
   - 4.1 [Time Messages](#41-time-messages)
   - 4.2 [Reference Data Messages](#42-reference-data-messages)
   - 4.3 [Event and State Change Messages](#43-event-and-state-change-messages)
   - 4.4 [Market by Order Messages](#44-market-by-order-messages)
   - 4.5 [Add Anonymous Order Messages](#45-add-anonymous-order-messages)
   - 4.6 [Trade Messages](#46-trade-messages)
   - 4.7 [Price Messages](#47-price-messages)
   - 4.8 [Index Constituent Information](#48-index-constituent-information)
5. [State Message Flow](#5-state-message-flow)
6. [Glimpse Overview](#6-glimpse-overview)
   - 6.1 [Glimpse for ITCH](#61-glimpse-for-itch)
   - 6.2 [Glimpse for ITCH-MDF](#62-glimpse-for-itch-mdf)
7. [Glimpse Message Formats](#7-glimpse-message-formats)
   - 7.1 [Glimpse Specific Messages](#71-glimpse-specific-messages)

**Appendices**
- A. [How to Build an Order Book View](#appendix-a-how-to-build-an-order-book-view)
- B. [How to Build a Trade Ticker](#appendix-b-how-to-build-a-trade-ticker)
- C. [Trades in Combination Order Books](#appendix-c-trades-in-combination-order-books)

---

## 1. Summary of Changes

| Version | Date | Summary of Changes |
|---------|------|-------------------|
| 1.0 | 05-03-2025 | First release |
| 2.0 | 28-06-2025 | Second Release<br>- Update Orderbookstate<br>- Remove Circuit breaker message for VCM<br>- Adding Glimpse for ITCH and ITCH-MDF |
| 3.0 | 08-07-2025 | Third Release<br>- Remove 24x7 mode information<br>- List of messages in ITCH-MDF<br>- Update of Datatype for Price in Trade Statistics<br>- Update of Add Anonymous Order description<br>- Update value of field priceType in Reference Price Message<br>- Update list of session states |
| 4.0 | 18-09-2025 | Fourth Release<br>- Update of Offset Information for assetName in Order Book Directory Message ITCH-MDF<br>- Update Description of marketSegmentId in Market Segment Directory Message<br>- Adding field numberOfItems in Market by Price Message table<br>- Additional values for status field in Indicative Quote Message<br>- Update of Index Price Message description<br>- Update of Trade Ticker Message description<br>- Adding State Message Flow table<br>- Adding new value for priceType in Reference Price Message |

---

## 2. About This Manual

This document aims to outline the dissemination of ITCH protocol data, providing field information for each message type. This manual contains the following messages:

### Table 1: ITCH Message Types

| Message Type | Name |
|--------------|------|
| T | Seconds Message |
| R | Order book Directory Message |
| L | Tick Size Table Message |
| S | System Event Message |
| O | Order Book State Message |
| A | Add Anonymous Order Message |
| E | Order Executed Message |
| C | Order Executed with Price Message |
| D | Order Delete Message |
| d | Order Book Clear message |
| P | Trade Message |
| Z | Equilibrium Price Message |

### Table 2: ITCH-MDF Message Types

| Message Type | Name |
|--------------|------|
| T | Seconds Message |
| R | Order book Directory Message |
| e | Exchange Directory Message |
| m | Market Directory Message |
| s | Market Segment Directory Message |
| L | Tick Size Table Message |
| K | Participant Directory Message |
| V | Issuer Directory Message |
| S | System Event Message |
| O | Order Book State Message |
| Z | Equilibrium Price Message |
| b | Market by Price Message |
| I | Trade Statistics Message |
| Q | Reference Price Message |
| J | Index Price Message |
| i | Trade Ticker Message |
| k | Price Limits Message |
| q | Indicative Quote Message |
| D | Index Member Message |

---

## 3. ITCH Overview

The ITCH Protocol is a communication protocol to provide real-time market data. This protocol allows users to receive detailed information about market activities, such as order openings, changes, and executions in the order book. IDX ITCH consists of two channels with details shown below:

### IDX ITCH
IDX ITCH protocol itself is a direct data feed product that displays all public orders and trades occurring on the auto-matched market(s). ITCH is an outbound market data feed only; the protocol does not support order entry.

**ITCH features, among others, the following data elements:**

- **Order level data (MBO) with attribution**  
  The system will provide full order depth using the standard ITCH format. ITCH uses a series of order messages to track the life of a customer order, including order executions. The ITCH message formats support market participant attribution, if used by the marketplace.

- **Trade messages**  
  ITCH supports a trade message to reflect a match of a non-displayable order in the system.

- **Reference Data**  
  - Order book Directory messages provide basic security data such as the ISIN code and Financial Product.
  - Tick Size Table Entry messages to convey Tick Sizes for order books.

- **Event controls**  
  - Order book State message to inform receivers of state changes.

> **Note:**  
> ITCH provides an order-book view and auto-matched trades.
> - Do not assume that the mechanisms of the matching logic can be derived from observing the ITCH feed.
> - ITCH cannot be used to manage private orders.
> - ITCH does not provide privately negotiated trades reported to the marketplace.

### ITCH-MDF (Market Data Feed)
Meanwhile, ITCH - Market Data Feed (MDF) is a direct data feed product that distributes complementary market data not contained in the regular ITCH feed. ITCH-MDF features the following data elements:

- Time Messages - Timestamps
- Reference Data messages
- Event and State Change Messages
- Price Messages

### 3.1 Architecture

The ITCH feed is made up of a series of sequenced messages. Each message is variable in length based on the message type. The messages that make up the ITCH protocol are typically delivered using a higher-level protocol that takes care of sequencing and delivery guarantees.

**SoupBinTCP** is a lightweight point-to-point protocol, built on top of TCP/IP sockets that allow delivery of a set of sequenced messages from a server to a client in real-time. SoupBinTCP guarantees that the client receives each message generated by the server in sequence, even across underlying TCP/IP socket connection failures.

The sequence numbers are implicit, meaning that the client maintains a counter that is increased every time a message is received. At reconnect after a connection loss, the client submits the last seen sequence number in its Logon message, and the server resends every message starting from that sequence number.

### 3.2 Data Types

**Table 3: List of Data Types**

| Type | Size | Notes |
|------|------|-------|
| Int8 | 1 byte | 8 bit Signed Big Endian |
| Int16 | 2 bytes | 16 bit Signed Big Endian |
| Int32 | 4 bytes | 32 bit Signed Big Endian |
| Int64 | 8 bytes | 64 bit Signed Big Endian |
| Alpha | Variable | Composed of non-control ISO 8859-1 (Latin-1) encoded bytes. Left justified and padded on the right with spaces |
| Price | 8 bytes | Prices are signed numeric fields. Number of decimals is specified in the Order book Directory message. A Price field with the minimum long value indicates that no price was available |
| Date | 4 bytes | Four-byte integer value derived from the Numeric data type. The decoded value represents a Date in YYYYMMDD-format |
| Timestamp | 8 bytes | UNIX Time (number of nanoseconds since 1970-01-01 00:00:00 UTC) |

### 3.3 Quantities in Units or Lots

All quantities fields have quantities expressed in units or lots. The 'Quantity Expressed In' field in Order Book Directory message is used to determine if quantities for an order book are expressed in units or lots.

---

## 4. Message Formats

The ITCH feed is composed of a series of messages that describe orders added to, removed from, and executed on the system. It also contains messages for basic reference data of the order books as well as state changes and halts.

### 4.1 Time Messages

For bandwidth efficiency reasons, timestamps are separated into two pieces:

**Table 4: Time Messages**

| Timestamp portion | Message Type | Notes |
|-------------------|--------------|-------|
| Seconds | Standalone message | Unix time (number of seconds since 1970-01-01 00:00:00 UTC). A Timestamp – Second message will be disseminated for every second for which there is at least one payload message |
| Nanoseconds | Field within individual messages | Reflects the number of nanoseconds since the most recent Timestamp-Seconds message that the payload message was generated |

#### 4.1.1 Seconds Messages

This message is sent every second for which at least one ITCH message is being generated. The message contains the number of seconds since the start of 1970-01-01 00:00:00 UTC, also called Unix Time.

**Table 5: Timestamp - Seconds Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | T - Seconds Message |
| seconds | 1 | 4 | Int32 | Unix time (number of seconds since 1970-01-01 00:00:00 UTC) |

### 4.2 Reference Data Messages

#### 4.2.1 Order Book Directory Message (ITCH)

At the start of each trading day, Order book directory messages are disseminated through ITCH for all active securities, including halted securities, in the system.

> **Note:** Intra-day transmissions of this message may occur when new order books are added to the system. Updates to existing order books may also be represented by intra-day Order book Directory messages.

**Table 6: Order Book Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | R - Order book Directory Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderBookId | 5 | 4 | Int32 | Denotes the primary identifier of an order book |
| symbol | 9 | 32 | Alpha | Order book short name |
| longName | 41 | 64 | Alpha | Human-readable long name of order book |
| isin | 105 | 12 | Alpha | ISIN code identifying the Asset the order book belongs to |
| financialProduct | 117 | 1 | Int8 | The financial product category the order book belongs to.<br>Supported values:<br>3 = Derivative - Future<br>5 = Equity<br>6 = Entitlement - Warrant<br>7 = Entitlement - Right<br>10 = Index |
| tradingCurrency | 118 | 3 | Alpha | Trading currency. Field not applicable for combination order books |
| decimalsInPrice | 121 | 2 | Int16 | Number of decimals used in price for this order book. A value of 256 means that the instrument is traded in fractions (each fraction is 1/256) |
| decimalsInNominalValue | 123 | 2 | Int16 | Number of decimals in Nominal Value |
| roundLotSize | 125 | 4 | Int32 | Indicates the quantity that represents a round lot for the issue |
| nominalValue | 129 | 8 | Int64 | Nominal value |
| numberOfLegs | 137 | 1 | Int8 | Number of legs. NOTE: Only applicable for combination instruments |
| underlyingOrderBookId | 138 | 4 | Int32 | The Underlying Asset ID (level 0) related to the Order Book |
| strikePrice | 142 | 8 | Price | Only applicable for derivative instruments |
| expirationDate | 150 | 4 | Date | Date of expiration. Applicable for Derivative and Entitlement instruments such as Warrants |
| decimalsInStrikePrice | 154 | 2 | Int16 | Number of decimals used in Strike Price for this order book. Only applicable for derivative instruments |
| optionType | 156 | 1 | Int8 | Supported values:<br>1 = Call<br>2 = Put |
| decimalsInQuantity | 157 | 2 | Int16 | This value defines the number of decimals used in Quantity for this order book |
| testOrderbook | 159 | 1 | Int8 | If true, the order book is dedicated for Production Realtime Verification.<br>Supported values:<br>1 = Yes<br>2 = No |
| quantityExpressedIn | 160 | 1 | Int8 | Defines if quantities for this order book are expressed in units or lots.<br>Supported values:<br>1 = Units<br>2 = Lots |

#### 4.2.2 Order Book Directory Message (ITCH-MDF)

At the start of each trading day, Order book directory messages are disseminated through ITCH-MDF for all active securities, including halted securities, in the system.

> **Note:** Intra-day transmissions of this message may occur when new order books are added to the system. Updates to existing order books may also be represented by intra-day Order book Directory messages.

**Table 7: Order Book Directory Message (ITCH-MDF)**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | R - Order book Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Denotes the primary identifier of an order book |
| partition | 9 | 1 | Int8 | Identifies the MME Partition for the Order book ID |
| symbol | 10 | 32 | Alpha | Order book short name |
| longName | 42 | 64 | Alpha | Human-readable long name of order book |
| isin | 106 | 12 | Alpha | ISIN code identifying the Asset the order book belongs to |
| assetName | 118 | 32 | Alpha | Name of orderbook's asset |
| financialProduct | 150 | 1 | Int8 | The financial product category the order book belongs to.<br>Supported values:<br>3 = Derivative - Future<br>5 = Equity<br>6 = Entitlement - Warrant<br>7 = Entitlement - Right<br>10 = Index |
| tradingCurrency | 151 | 3 | Alpha | Trading currency. Field not applicable for combination order books |
| decimalsInPrice | 154 | 2 | Int16 | This value defines the number of decimals used in price for this order book. NOTE: A value of 256 means that the instrument is traded in fractions (each fraction is 1/256) |
| decimalsInNominalValue | 156 | 2 | Int16 | This value defines the number of decimals in Nominal Value |
| roundLotSize | 158 | 4 | Int32 | Indicates the quantity that represents a round lot for the issue. The order quantity is a multiple of the lot size |
| nominalValue | 162 | 8 | Int64 | Nominal value |
| numberOfLegs | 170 | 1 | Int8 | Number of legs. NOTE: Only applicable for combination instruments |
| underlyingOrderBookId | 171 | 4 | Int32 | The Underlying Asset ID (level 0) related to the Order Book |
| strikePrice | 175 | 8 | Price | NOTE: Only applicable for derivative instruments |
| expirationDate | 183 | 4 | Date | Date of expiration. Applicable for Derivative and Entitlement instruments such as Warrants |
| decimalsInStrikePrice | 187 | 2 | Int16 | This value defines the number of decimals used in Strike Price for this order book. NOTE: Only applicable for derivative instruments |
| optionType | 189 | 1 | Int8 | Applicable for derivative instruments and Warrants. A value of 0 indicates N/A for the order book.<br>Supported values:<br>1 = Call<br>2 = Put |
| marketId | 190 | 4 | Int32 | Market ID |
| exchangeId | 194 | 4 | Int32 | Exchange ID |
| decimalsInQuantity | 198 | 2 | Int16 | This value defines the number of decimals used in Quantity for this order book |
| sectorCode | 200 | 4 | Alpha | The Asset Sector the order book belongs to |
| tradableQuantity | 204 | 8 | Int64 | The quantity available for trading. The tradable quantity can be less than the outstanding quantity |
| outstandingQuantity | 212 | 8 | Int64 | The total outstanding quantity. The quantity held by all shareholders |
| lastTradedDate | 220 | 4 | Date | This is the last date the asset can be traded |
| contractMultiplier | 224 | 8 | Int64 | The number of Underlying Assets one unit of the derivatives contract is based on. This is also frequently referred to as Contract Size |
| multiplier | 232 | 8 | Int64 | Specifies the ratio or multiply factor used to convert from contracts to shares |
| decimalsInMultiplier | 240 | 2 | Int16 | The number of decimals used for Multiplier or Contract Multiplier |
| minOrderQuantity | 242 | 8 | Int64 | This is the lowest allowed order quantity when an order is entered |
| maxOrderQuantity | 250 | 8 | Int64 | The order quantity cannot exceed this limit when entered |
| numberOfSettlementDays | 258 | 4 | Int32 | The number of settlement days applicable for trades in this order book |
| primary | 262 | 1 | Int8 | If the order book is configured as the primary one for the related Asset.<br>Supported values:<br>1 = Yes<br>2 = No |
| testOrderbook | 263 | 1 | Int8 | If the order book is dedicated to PRV - Production Realtime Verification.<br>Supported values:<br>1 = Yes<br>2 = No |
| listingBoard | 264 | 32 | Alpha | The listing board the order book relates to, if any |
| minOrderValue | 296 | 8 | Int64 | The configured minimum order value accepted in the order book for an order. Partially traded orders where the remainder is below the limit remain in the order book and are not revalidated |
| maxOrderValue | 304 | 8 | Int64 | The maximum order value accepted in the order book |
| decimalsInOrderValue | 312 | 2 | Int16 | The number of decimals used for the minimum and maximum order values |
| assetExtendedName | 314 | 100 | Alpha | Extended name of the order book asset |
| marketSegmentId | 414 | 4 | Int32 | Market Segment ID of the order book |
| issuerId | 418 | 4 | Int32 | Issuer Id of the order book's asset |
| ipoPrice | 422 | 8 | Price | The IPO Price for the order book. The field decimalsInPrice defines the number of decimals used in the IPO Price |
| delistingDate | 430 | 4 | Date | The date when the order book asset is no longer valid |
| remarks | 434 | 40 | Alpha | Remarks of the order book asset |
| quantityExpressedIn | 474 | 1 | Int8 | Defines if quantities for this order book are expressed in units or lots.<br>Supported values:<br>1 = Units<br>2 = Lots |

#### 4.2.3 Tick Size Table Message

This message contains information on a tick size for a price range. Together, all Tick Size Table messages with the same order book ID form a complete Tick Size Table. Each order book has a set of Tick Size Table Messages to define its tick size table.

> **Note:** The number of decimals in prices are given by the Order Book Directory message for this order book.

**Table 8: Tick Size Table Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | L - Tick Size Table Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderBookId | 5 | 4 | Int32 | The order book to which this entry belongs |
| tickSize | 9 | 8 | Int64 | Tick Size for the given price range |
| priceFrom | 17 | 8 | Price | Start of price range for this entry |
| priceTo | 25 | 8 | Price | End of price range for this entry. Zero (0) means infinity |

#### 4.2.4 Exchange Directory Message

At the start of each trading day, Exchange directory messages are disseminated for all active Exchanges in the system.

**Table 9: Exchange Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | e - Exchange Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| exchangeId | 5 | 4 | Int32 | Exchange ID |
| exchangeName | 9 | 32 | Alpha | Exchange Name |

#### 4.2.5 Market Directory Message

At the start of each trading day, Market directory messages are disseminated for all active Markets in the system.

**Table 10: Market Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | m - Market Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| marketId | 5 | 4 | Int32 | Market ID |
| marketName | 9 | 32 | Alpha | Market Name |

#### 4.2.6 Market Segment Directory Message

At the start of each trading day, Market Segment directory messages are disseminated for all active Markets in the system.

**Table 11: Market Segment Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | s - Market Segment Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| marketSegmentId | 5 | 4 | Int32 | Market Segment ID. Please refer to 'Table 34: Trading Session' |
| marketSegmentName | 9 | 32 | Alpha | Name of Market Segment |
| marketSegmentLongName | 41 | 64 | Alpha | Long Name of Market Segment |

#### 4.2.7 Participant Directory Message

At the start of each trading day, Participant Directory messages are disseminated for all active and suspended trading participants.

Intra-day transmissions of this message may occur when new trading participants are added to the system. Updates to existing trading participants may also be represented by intra-day Participant Directory messages.

**Table 12: Participant Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | K – Participant Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| participantId | 5 | 7 | Alpha | Participant identifier |
| participantDescription | 12 | 128 | Alpha | Description of the participant |
| status | 140 | 1 | Int8 | Status indicating whether participant is active or suspended.<br>Supported values:<br>1 = Active<br>2 = Suspended |

#### 4.2.8 Issuer Directory Message

At the start of each trading day, Issuer Directory messages are disseminated for all Participants that are configured as an issuer on any asset. Participants not configured as Issuer on any asset are not disseminated.

> **Note:** Intra-day transmissions of this message may occur when a new issuer is added to the system or when updates are made to the names of an existing issuer. The message will only contain the added or updated issuers.

**Table 13: Issuer Directory Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | V – Issuer Directory Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| issuerId | 5 | 4 | Int32 | Issuer ID |
| name | 9 | 32 | Alpha | Name of the issuer |
| longName | 41 | 64 | Alpha | Long name of the issuer |

### 4.3 Event and State Change Messages

#### 4.3.1 System Event Message

The system event message type is used to signal a market or data feed handler event.

**Table 14: System Event Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | S – System Event Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| event | 5 | 1 | Alpha | Event Code. The system supports the following event codes on a daily basis:<br>"O" = Start of Messages. Outside of time stamp messages, the start of day message is the first message sent in any trading day.<br>"C" = End of Messages. This is always the last message sent in any trading day.<br>Supported values:<br>O = Start of Messages<br>C = End of Messages |

#### 4.3.2 Order Book State Message

The Order book state message relays information on state changes.

**Table 15: Order Book State Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | O – Order Book State Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Orderbook identifier |
| stateName | 9 | 20 | Alpha | Name of the order book session state |

The list of possible state throughout the day depends on the Order Book's Market Segment. Please refer to 'Table 34: Trading Session' for session message flow information.

**Table 16: List of Possible stateName**

| Field | Description |
|-------|-------------|
| Break | Break Session |
| Break_Call | Break Session Call Auction |
| Calculate_CP | Calculate Closing Price |
| Call_RandomClose | Random Closing at Call Auction |
| Close_NG | Close Negotiation |
| Close_RF | Close Regular Futures |
| Close_RG | Close Regular |
| Close_TN | Close Cash |
| Close_INDEX | Close Index |
| EndofDay | End of Day |
| Matching_CallAuction | Matching at Call Auction |
| Matching_Close | Matching at Closing Regular |
| Matching_PreOpen | Matching at Pre-Opening |
| NonCancel | Non cancellation |
| PostTrade | Post Trading |
| PreClose | Pre-Closing |
| PreOpen | Pre-Opening |
| Open | Open for Index Calculation |
| RandomClose | Random Closing |
| Session_1_NG | Session 1 NG |
| Session_1_RF | Session 1 Regular Futures |
| Session_1_RG | Session 1 Regular |
| Session_1_RG_Call | Call Auction Session 1 |
| Session_1_TN | Session 1 Cash |
| Session_1_TN_Call | Call Auction Session 1 - TN |
| Session_2_NG | Session 2 NG |
| Session_2_RF | Session 2 Regular Futures |
| Session_2_RG | Session 2 Regular |
| Session_2_RG_Call | Call Auction Session 2 |
| Session_2_TN_Call | Call Auction Session 2 - TN |
| Session_3_RG_Call | Call Auction Session 3 |
| Session_3_TN_Call | Call Auction Session 3 - TN |
| Session_4_RG_Call | Call Auction Session 4 |
| Session_5_RG_Call | Call Auction Session 5 |
| SOBD | Start of Business Date |
| Suspend | Suspend with Cancel Orders |
| Trading_Halt | Trading Halt by Keeping Orders |

### 4.4 Market by Order Messages

> **Note:** Order IDs are only unique per order book and side. When modifying or deleting orders, be careful to only update the order with the correct side and order book, since the same Order ID may be present in multiple order books and/or sides.

### 4.5 Add Anonymous Order Messages

An Add Anonymous Order Message indicates that a new order has been accepted by components in the system and was added to the displayable book. The message includes an Order ID that is unique per order book used by components in the system to track the order. For order tracking purposes, the combination of orderID and side will be used to identify each order.

**Table 17: Add Anonymous Order Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | A - Add Anonymous Order Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderId | 5 | 8 | Int64 | The identifier assigned to the new order. The number is only unique per Order book and side |
| orderBookId | 13 | 4 | Int32 | Order book identifier |
| side | 17 | 1 | Alpha | Side of the order.<br>Supported values:<br>B = buy order<br>S = sell order |
| orderBookPosition | 18 | 4 | Int32 | Rank within order book. For details, see Appendix A, How to Build an Order Book |
| quantity | 22 | 8 | Int64 | The visible quantity of the order |
| price | 30 | 8 | Price | The display price of the new order |
| exchangeOrderType | 38 | 2 | Int16 | Additional order attributes. Applicable types may be defined by the marketplace. This field is a bit map. Multiple values may be set simultaneously.<br>Supported values:<br>0 = Not applicable<br>2 = Short sell<br>4 = Market bid<br>8192 = Bait (implied order) |
| quantityCondition | 40 | 1 | Int8 | Supported values:<br>1 = Quantity Restricted<br>2 = No restriction |

#### 4.5.1 Modify Order Messages

Modify Order messages always include the Order ID, Order book ID and Side of the Add Order to which the update applies.

To determine the currently displayed quantity for an order, ITCH subscribers must deduct the quantity stated in the Modify message from the original quantity stated in the Add Order message with the same Order ID. Components in the system may send multiple Modify Order messages for the same order and the effects are cumulative. When the quantity displayed for an order reaches zero, the order is dead and should be removed from the book.

##### 4.5.1.1 Order Executed Message

This message is sent whenever an order on the book is executed in whole or in part.

If the incoming order causing the match cannot be fully filled, the remainder will be placed in the book after the match has occurred.

It is possible to receive several Order Executed Messages for the same order if that order is executed in several parts. Multiple Order Executed Messages on the same order are cumulative.

**Table 18: Order Executed Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | E - Order Executed Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderId | 5 | 8 | Int64 | The order ID is associated with the executed order |
| orderBookId | 13 | 4 | Int32 | Order book identifier |
| side | 17 | 1 | Alpha | Supported values:<br>B = buy order<br>S = sell order |
| quantity | 18 | 8 | Int64 | The quantity being executed |
| matchId | 26 | 8 | Int64 | Assigned by the system to each match executed |
| comboGroupId | 34 | 4 | Int32 | Used to group combination order book executions and the trades in the constituent order books together. See Appendix C for details |
| owner | 38 | 7 | Alpha | Participant ID, owner. Will be set to blank (space) for anonymous markets |
| counterparty | 45 | 7 | Alpha | Participant ID, counterparty. Will be set to blank (space) for anonymous markets |

##### 4.5.1.2 Order Executed with Price Message

This message is sent in the relatively rare event that an order on the book is executed in whole or in part with a price different than the initial display price.

If the incoming order causing the match cannot be fully filled, the remainder will be placed in the book after the match has occurred.

It is possible to receive several Order Executed messages for the same order if that order is executed in several parts. Multiple Order Executed messages on the same order are cumulative.

The executions may be marked as non-printable. If a participant is looking to use the ITCH data in trade tickers or volume calculations, it is recommended that participants ignore messages marked as non-printable to prevent double counting.

> **Note:**
> - Combination orders on the book that execute will always be represented by this message.
> - Combination orders that execute will have the Printable flag set to "N". The trades that occur in the legs of the combination will be printable. This avoids double counting of the combination order and its leg trades. Leg trades will be published with the Trade message.

**Table 19: Order Executed with Price Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | C - Order Executed with Price Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderId | 5 | 8 | Int64 | The order ID is associated with the executed order |
| orderBookId | 13 | 4 | Int32 | Order book identifier |
| side | 17 | 1 | Alpha | Supported values:<br>B = buy order<br>S = sell order |
| quantity | 18 | 8 | Int64 | The quantity being executed |
| matchId | 26 | 8 | Int64 | Assigned by the system to each match executed |
| comboGroupId | 34 | 4 | Int32 | Used to group combination order book executions and the trades in the constituent order books together. See Appendix C for details |
| owner | 38 | 7 | Alpha | Participant ID, owner. Will be set to blank (space) for anonymous markets |
| counterparty | 45 | 7 | Alpha | Participant ID, counterparty. Will be set to blank (space) for anonymous markets |
| price | 52 | 8 | Price | Trade price |
| cross | 60 | 1 | Alpha | Trade at cross.<br>Supported values:<br>Y = yes<br>N = no |
| printable | 61 | 1 | Alpha | Indicates if the trade should be included in trade tickers and volume calculations.<br>Supported values:<br>N = non-printable<br>Y = printable |

##### 4.5.1.3 Order Delete Message

This message is sent whenever an order on the book is being deleted. There will be no remaining quantity, so the order should be removed from the book.

Normally, no Order Delete message is sent when an order is completely filled. The receiver needs to keep track of the remaining quantity on all orders by recalculating the remaining quantity on each Order Executed message received. Orders must be removed from the book when the remaining quantity reaches 0.

**Table 20: Order Delete Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | D – Order Delete Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderId | 5 | 8 | Int64 | The original order identifier of the order being deleted. The Order ID is only unique per order book and side |
| orderBookId | 13 | 4 | Int32 | Order book identifier |
| side | 17 | 1 | Alpha | The type of order being deleted.<br>Supported values:<br>B = buy order<br>S = sell order |

##### 4.5.1.4 Order Book Clear Message

Instead of sending individual Order Delete messages, a single Order Book Clear message can be sent to communicate that all orders for the specified order book have been deleted. Typically, this is used when entering a blind auction.

**Table 21: Order Book Clear Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | d – Order Book Clear message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderBookId | 5 | 4 | Int32 | Order book identifier |

### 4.6 Trade Messages

#### 4.6.1 Trade Message

The Trade Message is designed to provide execution details for normal match events involving non-displayable order types. This message is also used to publish individual cross trades.

Since no Add Order Message is generated when a non-displayed order is initially received, the Order Executed message cannot be used for all matches. The Trade Message is used to report a match for a non-displayable order in the book.

It is possible to receive multiple Trade Messages for the same order if that order is executed in several parts. Trade Messages for the same order are cumulative.

Trade Messages should be included in trade tickers as well as volume and other market statistics. Since Trade Messages do not affect the displayed book, however, they may be ignored by participants just looking to build and track the order book view.

**Table 22: Trade Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | P - Trade Message identifier |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| matchId | 5 | 8 | Int64 | Assigned by the system to each match executed |
| comboGroupId | 13 | 4 | Int32 | Used to group combination order book executions and the trades in the constituent order books together. See Appendix C for details |
| side | 17 | 1 | Alpha | Type of non-display order on the book being matched. Will be set to blank (space) for anonymous markets.<br>Supported values:<br>B = buy order<br>S = sell order |
| quantity | 18 | 8 | Int64 | Quantity being matched in this execution |
| orderBookId | 26 | 4 | Int32 | Order book identifier |
| price | 30 | 8 | Price | Trade Price |
| owner | 38 | 7 | Alpha | Participant ID, owner. Will be set to blank (space) for anonymous markets |
| counterparty | 45 | 7 | Alpha | Participant ID, counterparty. Will be set to blank (space) for anonymous markets |
| printable | 52 | 1 | Alpha | Indicates if the trade should be included in trade tickers and volume calculations.<br>Supported values:<br>N = non-printable<br>Y = printable |
| cross | 53 | 1 | Alpha | Trade at Cross.<br>Supported values:<br>Y = yes<br>N = no |

#### 4.6.2 Auction Messages

Markets by order dissemination may be disabled during auctions by configuration. In such cases, every existing order will be removed from the book by an Order Delete message immediately prior to the auction.

> **Note:** Owners of these orders must not interpret this as order cancellations. Use the private order flow to determine the state of your orders.

##### 4.6.2.1 Equilibrium Price Message

This message is used when auctions occur. The message provides the changes in equilibrium price. If any Price field has the minimum long value, this means that no price is available.

**Table 23: Equilibrium Price Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | Z - Equilibrium Price Message |
| nanos | 1 | 4 | Int32 | Timestamp in nanoseconds |
| orderBookId | 5 | 4 | Int32 | Order book ID |
| bidQuantity | 9 | 8 | Int64 | Quantity at equilibrium price on the bid side |
| askQuantity | 17 | 8 | Int64 | Quantity at equilibrium price on the ask side |
| price | 25 | 8 | Price | Equilibrium Price |
| bestBidPrice | 33 | 8 | Price | Best Bid Price |
| bestAskPrice | 41 | 8 | Price | Best Ask Price |
| bestBidQuantity | 49 | 8 | Int64 | Best Bid Quantity |
| bestAskQuantity | 57 | 8 | Int64 | Best Ask Quantity |

### 4.7 Price Messages

The messages stated in this section will be distributed through ITCH-MDF.

#### 4.7.1 Market by Price (MBP, Incremental) Message

The Market by Price Message is used to publish order book Prices per Price level.

**Table 24: Market by Price Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | b – Market by Price Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order book identifier |
| maximumLevel | 9 | 1 | Int8 | The maximum price level (for example 10). The value may be changed during the day. Price levels > Maximum Level are not applicable and shall be deleted by a consumer of the Market by Price Message |
| numberOfItems | 10 | 1 | Int8 | The number of price level items in the array |
| PriceLevelItem group | 10 + n * 20 | 20 * n | Array | Array of Price Update items. The item counter n for the Offset is counted from 0 up to Number of Level items -1 when Number of Level items > 0 |

**Table 25: Price Update Item**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| **Start PriceLevelItem** |  |  |  |  |
| levelUpdateAction | 0 | 1 | Alpha | Type of market update action.<br>Supported values:<br>C = Change<br>D = Delete<br>N = New |
| side | 1 | 1 | Alpha | Supported values:<br>B = Bid<br>A = Ask |
| level | 2 | 1 | Int8 | The numeric order of the price level, where "1" is the first price level |
| price | 3 | 8 | Price | Price for the Level. Value is set to minimum long value for a Price Delete Item (when update action is 'D' = Delete From) |
| quantity | 11 | 8 | Int64 | Visible quantity for the Level. Value is 0 for a Price Delete Item (when update action is 'D' = Delete From) |
| numberOfDeletes | 19 | 1 | Int8 | Number of Levels to Delete. Value is > 0 and used for a Price Delete Item (when update action is 'D' = Delete From). Value is 0 and not used for a Price Update Item (when update action is 'N' or 'C') |
| **End PriceLevelItem** |  |  |  |  |

> **Note:** If Price is set to the minimum long value, this means that no price is available.

#### 4.7.2 Trade Statistics Message

The Trade Statistics Message is used to publish Trade Statistics by Orderbook ID. The values in this message are for the current Trading Day.

**Table 26: Trade Statistics Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | I – Trade Statistics Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order book identifier for the order book the statistics is related to |
| openPrice | 9 | 8 | Price | The first trade price |
| highPrice | 17 | 8 | Price | The highest trade price |
| lowPrice | 25 | 8 | Price | The lowest trade price |
| lastPrice | 33 | 8 | Price | The last trade price |
| lastAuctionPrice | 41 | 8 | Price | The last auction trade price |
| lastQuantity | 49 | 8 | Int64 | The last trade quantity |
| turnOverQuantity | 57 | 8 | Int64 | Calculated as sum of Traded Quantity. Includes Reported Trades |
| reportedTurnOverQuantity | 65 | 8 | Int64 | Calculated as sum of Reported Trades Quantity |
| turnOverValue | 73 | 8 | Int64 | Calculated as sum of Traded Price * Traded Quantity |
| vwap | 81 | 8 | Int64 | The Volume Weighted Average Price |
| dailyNumberOfTrades | 89 | 8 | Int64 | The daily number of trades |

> **Note:** If Price is set to the minimum long value, this means that no price is available.

#### 4.7.3 Reference Price Message

The Reference Price Message is used to publish Reference Prices per Price Types.

**Table 27: Reference Price Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | Q - Price Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order book identifier |
| priceType | 9 | 1 | Int8 | Supported values:<br>**1 = Externally Set** - A price determined externally before the trading session begins, usually set by the Exchange due to corporate actions or special conditions<br>**5 = Settlement** - The reference price used for futures contracts<br>**6 = Ever Last** - The last traded price of security during the trading day<br>**11 = Closing Price** - The closing price of a security at the end of the previous trading day. If no transactions occur on previous day, the reference price for the day will refer to the Everlast price<br>**12 = Adjusted Closing Price** - Reference Price |
| price | 10 | 8 | Price | Reference Price |
| updatedTimestamp | 18 | 8 | Timestamp | The time (in UTC) when the Price was updated (if available) |

> **Note:** If the price is set to the minimum long value, it means that no price is available. The reference price that will be used is the most recently sent price.

#### 4.7.4 Index Price Message

This message is used for sending index price information.

**Table 28: Index Price Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | J - Index Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order book identifier |
| lastPrice | 9 | 8 | Price | Current day Index last price |
| calculatedTimestamp | 17 | 8 | Timestamp | Time of Index computation (in UTC) |

#### 4.7.5 Trade Ticker Message

The Trade Ticker Message is used to publish Trade Ticker.

**Table 29: Trade Ticker Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | i – Trade Ticker Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order book identifier |
| dealId | 9 | 8 | Int64 | Identifier for the deal |
| dealSource | 17 | 1 | Int8 | Supported values:<br>1 = Auction<br>2 = Auto match<br>3 = Trade Report<br>4 = Combo to Combo |
| price | 18 | 8 | Price | The deal price |
| quantity | 26 | 8 | Int64 | The deal quantity |
| dealTime | 34 | 8 | Timestamp | Nano seconds since Epoch |
| action | 42 | 1 | Int8 | Supported values:<br>1 = New<br>2 = Updated<br>3 = Canceled |
| aggressor | 43 | 1 | Alpha | The side that is the aggressive part in the trade (taker). If neither buy nor sell is aggressor, the ascii code 0 is sent.<br>Supported values:<br>NULL = None<br>B = Bid<br>A = Ask |
| tradeReportType | 44 | 2 | Int16 | Trade Report type value if the trade ticker originates from a reported trade.<br>Supported values:<br>0 = Not a trade report<br>1 = Negotiated Deal<br>2 = Indicative |
| crossedTrade | 46 | 1 | Int8 | Supported values:<br>0 = Trade is not a crossed trade<br>1 = Trade is a crossed trade (internal trade within the participant) |

#### 4.7.6 Price Limits Message

The message provides an update to the price limits for an order book, either the static or the dynamic limits.

**Table 30: Price Limits Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | k – Price Limits Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | The order book the Price Limits update is for |
| upperLimit | 9 | 8 | Price | The Upper Price Limit. Min long value means limit is removed/no limits |
| lowerLimit | 17 | 8 | Price | The Lower Price Limit. Min long value means limit is removed/no limits |
| type | 25 | 1 | Int8 | Supported values:<br>1 = Price Limits |
| category | 26 | 1 | Int8 | Supported values:<br>1 = Static Limits<br>2 = Dynamic Limits |
| referencePrice | 27 | 8 | Price | The reference price used for the calculated static or dynamic Price Limits |

#### 4.7.7 Indicative Quote Message

This message is sent to give information about when Market Makers enters new or updates are made to Indicative Quotes.

**Table 32: Indicative Quote Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | q – Indicative Quote Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| orderBookId | 5 | 4 | Int32 | Order Book identifier |
| status | 9 | 1 | Int8 | Supported values:<br>1 = New<br>2 = Accepted<br>3 = Cancelled by user<br>4 = Cancelled not enough quantity<br>5 = Cancelled by system<br>6 = Declined<br>7 = Cancelled due to suspended investor<br>8 = Cancelled due to suspended account |
| indicativeQuoteId | 10 | 8 | Int64 | Identifier for the Indicative Quote |
| side | 18 | 1 | Alpha | Supported values:<br>B = Bid<br>A = Ask |
| price | 19 | 8 | Price | The price of the Indicative Quote |
| quantity | 27 | 8 | Int64 | The quantity of the Indicative Quote |
| leavesQuantity | 35 | 8 | Int64 | The Indicative Quote quantity that is left |

### 4.8 Index Constituent Information

#### 4.8.1 Index Member Message

The Index Member Message is used to publish index constituent information.

**Table 33: Index Member Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | D – Index Member Message |
| nanos | 1 | 4 | Int32 | Nanoseconds portion of the timestamp |
| indexOrderBookId | 5 | 4 | Int32 | The index orderbook |
| memberOrderBookId | 9 | 4 | Int32 | The member order book included in the index order book |
| weight | 13 | 8 | Int64 | The weight of the member orderbook in the index |

---

## 5. State Message Flow

The following information specifies the possible session states available via ITCH for each market segment, with reference to the Market Segment Directory Message of the respective order book.

**Table 34: Trading Session**

| Identity Name | Segment Name | State List Entry |
|---------------|--------------|------------------|
| **Regular Pre-Open** | ORDI-OPEN_RG<br>ORDI-OPEN_RG_SHORTMARGIN<br>ORDI-OPEN_RG_MARGIN-IDX30<br>ORDI-OPEN_RG_MARGIN-LQ45<br>ORDI-OPEN_RG_SHORTMARGIN-IDX80<br>ORDI-OPEN_RG_PST<br>ORDI-OPEN_RG_SHORTMARGIN-LQ45<br>ORDI-OPEN_RG_MARGIN | SOBD<br>PreOpen<br>NonCancel<br>Matching_PreOpen<br>Session_1_RG<br>Break<br>Session_2_RG<br>PreClose<br>NonCancel<br>RandomClose<br>Matching_Close<br>Calculate_CP<br>PostTrade<br>Close_RG<br>EndofDay |
| **Tunai Trading Session** | ORDI_TN<br>ORDI_TN_SHORTMARGIN<br>ETF_TN<br>DINFRA_TN<br>ACCEL_TN<br>DIRE_TN<br>WARI_TN<br>WARI_IPO_TN<br>SWARI_TN<br>SWARI_IPO_TN<br>RIGHT_TN<br>ETF_TN_MARGIN<br>ORDI_TN_MARGIN-LQ45<br>ORDI_TN_MARGIN-IDX80<br>ORDI_TN_SHORTMARGIN-IDX80<br>ORDI_TN_SHORTMARGIN-LQ45<br>ORDI_TN_MARGIN<br>ACCEL_TN_MARGIN<br>ACCEL_TN_SHORTMARGIN | SOBD<br>Session_1_TN<br>Close_TN<br>Calculate_CP<br>EndofDay |
| **Tunai Call Auction Trading Session** | WATCH-CALL_TN<br>RIGHT-CALL_TN<br>WATCH-CALL_TN_SHORTMARGIN<br>WARI-CALL_TN<br>WATCH-CALL_TN_MARGIN<br>WATCH-CALL_TN_MARGIN-IDX80<br>WATCH-CALL_TN_MARGIN-LQ45<br>WATCH-CALL_TN_SHORTMARGIN-IDX80<br>WATCH-CALL_TN_SHORTMARGIN-LQ45 | SOBD<br>Session_1_TN_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Session_2_TN_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Session_3_TN_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Close_TN<br>Calculate_CP<br>EndofDay |
| **Derivative Trading Session*** | IDX30-FUT_RF<br>LQ45-FUT_RF<br>SSF_RF<br>GB-FUT_RF<br>BM-FUT_RF<br>SSF_RF_ASII<br>SSF_RF_BBCA<br>SSF_RF_BBRI<br>SSF_RF_MDKA<br>SSF_RF_TLKM<br>MSLCHK-FUT_RF | SOBD<br>Session_1_RF<br>Break<br>Session_2_RF<br>Close_RF<br>Calculate_CP<br>EndofDay |
| **Negotiation Trading Session** | ORDI_NG<br>DIRE_NG<br>DINFRA_NG<br>ETF_NG<br>ACCEL_NG<br>WARI_NG<br>SWARI_NG | SOBD<br>Session_1_NG<br>Break<br>Session_2_NG<br>Close_NG<br>Calculate_CP<br>EndofDay |
| **Regular Call Auction Trading Session** | WATCH-CALL_RG_SHORTMARGIN<br>WATCH-CALL_RG_SHORTMARGIN-IDX80<br>WATCH-CALL_RG_SHORTMARGIN-LQ45<br>WATCH-CALL_RG_MARGIN<br>WATCH-CALL_RG<br>WARI-CALL_RG<br>WATCH-CALL_RG_MARGIN-IDX80<br>WATCH-CALL_RG_MARGIN-LQ45 | SOBD<br>Session_1_RG_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Session_2_RG_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Session_3_RG_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Break_Call<br>Session_4_RG_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Session_5_RG_Call<br>Call_RandomClose<br>Matching_CallAuction<br>Calculate_CP<br>PostTrade<br>Close_RG<br>EndofDay |
| **Regular Session for ACCEL, WARI, RGHI, MUTI*** | SWARI_RG_LP-ZP<br>ORDI_RG<br>ORDI_RG_MARGIN-LQ45<br>ORDI_RG_MARGIN-IDX80<br>ORDI_RG_SHORTMARGIN<br>ORDI_RG_SHORTMARGIN-IDX80<br>ORDI_RG_SHORTMARGIN-LQ45<br>ETF_RG<br>DINFRA_RG<br>ACCEL_RG<br>DIRE_RG<br>WARI_RG<br>SWARI_RG<br>SWARI_RG_LP-BQ<br>ACCEL_RG_MARGIN<br>SWARI_RG_LP-DR<br>ORDI_RG_MARGIN<br>SWARI_RG_LP-HD<br>ACCEL_RG_SHORTMARGIN<br>SWARI_RG_LP-YU<br>WARI-IPO_RG<br>ETF_RG_MARGIN | SOBD<br>Session_1_RG<br>Break<br>Session_2_RG<br>PreClose<br>NonCancel<br>RandomClose<br>Matching_Close<br>Calculate_CP<br>PostTrade<br>Close_RG<br>EndofDay |
| **Right Negotiation Trading Session** | RIGHT_NG | SOBD<br>Session_1_NG<br>Close_NG<br>Calculate_CP<br>EndofDay |

*Additional segment names in this trading session will be added over time

---

## 6. Glimpse Overview

### 6.1 Glimpse for ITCH

A separate connection for obtaining snapshots, called Glimpse for ITCH, can optionally be provided to enable the user to reconnect intraday and be current with the live stream. Glimpse is thus only used in case snapshot is required for the ITCH feed. Connecting to Glimpse intraday obtains a snapshot of a subset of messages configured for this stream. The snapshot of the stream is taken at the point in time when the user connects and logs in to Glimpse. The snapshot is tagged with a sequence number, the point at which one can listen to the live stream. The Glimpse snapshot is available in SoupBinTCP connections only.

> **Note:** System Events and Equilibrium Price Messages are not included in the GLIMPSE snapshot.

Glimpse for ITCH has the following special data element:
- Snapshot Message

### 6.2 Glimpse for ITCH-MDF

A separate connection for obtaining snapshots, called Glimpse for MDF, can optionally be provided to enable the user to reconnect intraday and be current with the live stream. Glimpse is thus only used in case snapshot is required for the MDF feed. Connecting to Glimpse intraday obtains a snapshot of a subset of messages configured for this stream. For the Market By Price (MBP) message the snapshot includes all Price Levels available at the time of the login. The snapshot of the stream is taken at the point in time when the user connects and logs in to Glimpse. The snapshot is tagged with a sequence number, the point at which one can listen to the live stream. The Glimpse snapshot is available in Soup connections only.

> **Note:** Market Announcements, System Events, Open Interest and Equilibrium Price Messages are not included in the Glimpse snapshot.

Glimpse for MDF has the following special data element:
- Snapshot Message

---

## 7. Glimpse Message Formats

Glimpse for ITCH/ITCH-MDF uses the same messages as ITCH/ITCH-MDF. For Glimpse specific messages, see below.

### 7.1 Glimpse Specific Messages

#### 7.1.1 Glimpse Snapshot Message

The Glimpse snapshot message returns the current ITCH/ITCH-MDF sequence number to be used when connecting to the ITCH/ITCH-MDF feed.

To maintain a real-time order display, firms should begin to process real-time ITCH/ITCH-MDF messages beginning with the sequence number stated in this snapshot message + 1.

**Table 18: Glimpse Snapshot Message**

| Field | Offset | Length | Type | Description |
|-------|--------|--------|------|-------------|
| msgType | 0 | 1 | Alpha | G – Glimpse Snapshot Message |
| itchSequenceNumber | 1 | 20 | Alpha | ITCH/ITCH-MDF SoupBin TCP sequence number when the snapshot was taken. To be used when logging in to the ITCH SoupBin TCP feed.<br>NOTE: While GLIMPSE is a binary feed, the SoupBin TCP uses ASCII characters to represent the sequence number |

---

## Appendix A: How to Build an Order Book View

The information needed to build an order book view from the ITCH message flow is contained in the Add Order Messages and the Modify Order Messages. The messages are:

- Add Order
- Order Executed
- Order Executed with Price
- Order Delete

Orders are ranked by Order book Position. 1 denotes the highest ranked order. For an Order Replace, the order must be removed from its previous position and inserted at New Order Book Position. An order inserted at an existing position shifts the order on that position down (and all orders below as well). A deleted or fully filled order causes existing orders below it to shift their position up one step to fill the "void."

- The Order Executed (with Price) message signals a partial or full fill. The order quantity must be reduced by the quantity of the Order Executed message.
- The Order Delete message tells the recipient to remove the order referenced.

## Appendix B: How to Build a Trade Ticker

The Trade Ticker is based on the following messages:
- Order Executed
- Order Executed with Price
- Trade

> **Note:**
> - Trades and Order Executed with Price messages marked as non-printable are excluded to avoid double booking of trades.
> - Reported trades are not included in ITCH.

## Appendix C: Trades in Combination Order Books

When a Combination order is executed, trades also occur in all legs of the combination. To learn about the Combination instrument and its constituents, query the Combination Order book Directory message.

### Communication of Combo vs. Combo Executions
- Order Executed with Price message for the Combination Order Book, with the Printable flag set to N (to avoid double counting)
- Trade messages in the constituent order books

### Communication of Combo vs. Outright Executions
- Order Executed, and/or Order Executed with Price messages for the constituent order books
- Trade message in the Combination Order book, with the Printable flag set to N (to avoid double counting)

The Combination Order Book execution and the constituent Order book executions have different Match ID, but the same Combo Group ID. Use the Combo Group ID field to group the Order Executed and the Trade messages for a combination execution together. 

> **Note:** The Combo Group ID field should not be assumed to be unique over time.

---

*End of Document*