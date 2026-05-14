# IDX FIX Drop Copy Specification

**Document Version:** 1.0  
**Date:** 14 October 2025  
**Issuer:** Indonesia Stock Exchange (Bursa Efek Indonesia)  
**Classification:** Private and Confidential

> **Disclaimer:** The information in this document may change in accordance with the progress of the ongoing project development.

---

## Summary of Changes

| Version | Date | List of Changes | Details |
|---------|------|-----------------|---------|
| 1.0 | 14/10/2025 | First release | — |

---

## Table of Contents

1. [About the Document](#1-about-the-document)
   - 1.1 [About this specification's status and refinement](#11-about-this-specifications-status-and-refinement)
   - 1.2 [Concepts and Legend](#12-concepts-and-legend)
   - 1.3 [Messages and scenarios in this document](#13-messages-and-scenarios-in-this-document)
   - 1.4 [Gateway configuration and optional messages](#14-gateway-configuration-and-optional-messages)
2. [Order Monitoring](#2-order-monitoring)
   - 2.1 [OrderMassStatusRequest (AF, in)](#21-ordermasstatusrequest-af-in)
   - 2.2 [ExecutionReport (8, out)](#22-executionreport-8-out)
3. [Trade Monitoring](#3-trade-monitoring)
   - 3.1 [TradeCaptureReportRequest (AD, in)](#31-tradecapturereportrequest-ad-in)
   - 3.2 [TradeCaptureReport (AE, out)](#32-tradecapturereport-ae-out)
   - 3.3 [TradeCaptureReportRequestAck (AQ, out)](#33-tradecapturereportrequestack-aq-out)
4. [Other Information](#4-other-information)
   - 4.1 [News (B, out)](#41-news-b-out)
   - 4.2 [BusinessMessageReject (j, out)](#42-businessmessagereject-j-out)

---

## 1 About the Document

### 1.1 About this specification's status and refinement

This specification depicts the FIX Drop Copy service for the JATS. The full suite of specifications includes Session layer, Reference Data, Order and Trade Entry, Order and Trade Monitoring (Drop Copy) and Market Data. A condensed and consistent set of tags is included with the pronounced aim to meet demands as imposed by authorities and regulation regimes under which financial bodies operate at each individual location.

As far as possible standard FIX tags and codesets are utilized and where not, custom tags and values have been added.

---

### 1.2 Concepts and Legend

#### 1.2.1 Characteristics given for a message

Each message is documented with a leading summary of the characteristics and use cases of the message.

| Characteristic | Content |
|----------------|---------|
| Direction | Can be "In to JATS" or "out from JATS". If there can be both variants, separate message sections are used. |
| Message code (FIX standard) introduced | The FIX message type: MsgType, tag 35 which is located in the header. The FIX version the message was introduced. Older messages are often less open to custom extensions. |
| FIX Session | The default session over which the message (and its containing flow) is communicated, for example "OE" (Order Entry session). |
| Available to | Defines the intended categories of business users. |
| Usage and Conditions | Contains information on specific use cases of relevance for the intended marketplace and any specific conditions that apply. |
| Limitations | Specifies limitations that may not be obvious, need attention or should be considered when designing operational processes. |
| Response/Acknowledgment | The immediate responses to expect for the successful and unsuccessful cases. |

#### 1.2.2 Message chaining or entity chaining

FIX offers two principles of chaining a sequence of orders, trades, and other successive events: the **message chaining model** and the **entity chaining model**.

- The **message chaining model** treats the event and its related message as "the same thing" and a successor message refers to its preceding in the chain by a message reference which implicitly becomes a reference to the event which entailed the message about it.
- The **entity reference model** references the entity the event created, such as a trade which is treated as something that exists in its own right, not in need of message proxy. In this model the trade is referenced directly.

Where not required to do otherwise this implementation adheres to the **entity chaining principle**. This manifests itself when cancelling a trade for example where, to this end, TradeID (1003) is more important than TradeReportID (571). The motivation is that in a system that makes use of several protocols, potentially partially overlapping, any entity, such as a trade or order, must be able to address unambiguously without referring to a specific protocol.

There are however situations where a TradeID is not yet established, for example when reporting an off-exchange trade: the marketplace (report recipient) has yet to create the TradeID upon a successful validation of the reported content. Until this is done, a suitable way to refer to the trade is by a reference to the TradeReportID (if submitted via FIX).

#### 1.2.3 About the Req'd (Required) column

The following nomenclature is used for the Req'd column:

- **Y:** The field is required (must have a value) by FIX standard or is mandated in JATS. In the latter case this is stated as remark in the column: "Req'd in JATS".
- **N:** The field is not required, i.e. optional, which is the default if not declared. It can be conditional in JATS to distinguish different use cases, which is stated as a remark in the column: "Cond'l in JATS".

The scope of the Req'd property is within the closest enclosure, which is the group, component, or message, in the given scenario.

Where a field is required because of FIX syntax, for example being first field in repeating group or to make an ID designation complete (Party, Security, any type/value construct, etc), this isn't made explicit in the Req'd column.

#### 1.2.4 About the UD (Union Datatype) column

The FIX standard opens for free to use enumeration ranges for some codesets. These are indicated by the union datatype (UD) feature where custom values can be freely added (and agreed between all parties of course) without breaking the protocol.

UD values can be **100+**, **1000+**, and **4000+** where numbers above (and including) these limits are reserved for implementation dependent use.

#### 1.2.5 About the Parties Group (453)

The Parties group is common to almost all business level messages, and is vital to all messages directed towards trading, like orders, quotes, trade reports and so on. The difference as compared to the CompID:s in the header is that the parties in the parties group define or reflect the business level stakeholders whereas the header CompID:s are pure technical actors.

A party definition is composed of 3 parts: **PartyID**, **PartyIDSource**, **PartyRole**. A party definition isn't valid syntactically if not all 3 parts are defined with valid values.

**Table 1: Summary of Parties group usage**

| Entity Description | Tag 448 PartyID | Tag 447 PartyIDSource | Tag 452 PartyRole | Req'd | Val'd |
|---|---|---|---|---|---|
| Firm owning the transaction business wise | Firm name in JATS | D | 1 = ExecutingFirm | Y | Y |
| User owning the transaction business wise | User name in JATS | D | 12 = ExecutingTrader | Y | Y |
| Account or Client to which the transaction is accountable. Only the first 15 characters are used and the last 16th character shall be a null character. The 3rd character must be either "D" (domestic) or "F" (foreign). | Account/Client designation | D | 24 = CustomerAccount | N | N |
| Counterparty to the transaction. Required or optional for certain types of transactions. | Firm name in JATS | D | 17 = ContraFirm | C | Y |
| If the transaction is submitted via a firm different from the owning firm, and that firm is not only a re-routing hub. Not all transaction types allow for this type of on-behalf-of submission. | Firm name in JATS | D | 7 = EnteringFirm | C | Y |
| The user at the on-behalf-of submitter firm sending the transaction. | User name in JATS | D | 36 = EnteringTrader | C | Y |
| A user who can act on an order but is not part of the order's entry. | User name in JATS | D | 53 = TraderMnemonic | C | Y |
| The firm eligible or required to quote or respond to quote requests for a certain instrument. | Firm name in JATS | D | 66 = MarketMaker | N | Y |
| The firm responsible for clearing a trade for an ExecutingParty. | Firm name in JATS | D | 4 = ClearingFirm | N | Y |

> All the PartyIDSource values are "D", however when invoking specific regulations such as MiFID, or other domains of designation, other symbols come into use as well, such as "P" or "G".

#### 1.2.6 About Price and Quantity

For Price and Quantity precision, the number of decimals for each determine the valid amounts. They are fetched from the SecurityList message as Instrument price precision and Instrument quantity precision and are entered with a decimal point. Examples:

- `44 (Price): 1.012345`
- `38 (OrderQty): 2.03`

#### 1.2.7 Quantities in Units or Lots

All quantities fields have quantities expressed in units or lots. Tag `InstrAttribType (871)` / `InstrAttibValue (872)` in SecurityList and SecurityListUpdateReport message is used to determine if quantities for an order book are expressed in units or lots.

#### 1.2.8 About length of String fields

FIX does not constrain string fields to any length. JATS does impose limits to the number of characters that can be entered in a string type tag where there are no pre-declared values to choose from. This is indicated as `[X]` in the type column together with the label `ISO-8859-1`. In practice this means X is the maximum number of characters provided they are coded into 1 byte.

#### 1.2.9 Gateway Configuration Options

The specification volume referred to as "FIX_SG", the JATS FIX specification for session and gateway, contains a section about configuration options for the various gateways. Declaring particular configuration settings at startup will enable additional messages or services to the gateway.

#### 1.2.10 About the Scenario feature

Scenarios allow for messages documentation to look different depending on use case and only contain the fields and values pertaining to it.

Scenario codes look like `P-A-D-S` where the constituents together represent a use case:

1. **1st part:** Product wide or gateway scope (JATS, OE)
2. **2nd part:** Product adaptation (C=Core, Xyz=adaptation Xyz, e.g. an exchange, marketplace, or other installation)
3. **3rd part:** Direction (i/o/io)
4. **4th part:** Specialization of any kind as needed; particular FIX msg (enclosed in `[ ]`), other special conditions

The default scenario is just called **"Base"**.

---

### 1.3 Messages and scenarios in this document

**Table 2: List of Messages**

| Name | Scenario | Description |
|------|----------|-------------|
| TradeCaptureReportRequest | base | Request TradeCaptureReport (TCR) messages for addressed participants. |
| TradeCaptureReportRequestAck | base | JATS accepts or rejects a request for TradeCaptureReport messages. |
| OrderMassStatusRequest | base | Request ExecutionReport messages for addressed participants. |
| News | base | Signal end of business messages stream. |
| BusinessMessageReject | ET-C-o | BMR general, out |
| TradeCaptureReport | DC-IDX-o | Disseminate trade events in JATS via TradeCaptureReport messages. |
| ExecutionReport | DC-IDX-o | Execution Report from an order event in JATS. |

> Note: The abbreviation "ET" is, for space reasons in tables and lists, used for short denotating "JATS".

---

### 1.4 Gateway configuration and optional messages

The specification volume referred to as "FIX_SG", the JATS FIX specification for session and gateway, contains a section about configuration options for the various gateways. Declaring particular configuration settings at startup will enable additional messages or services to the gateway. Please consult this section for information on how to configure the Drop Copy gateway to enable the following message(s):

- `ExecutionReport (8)`

The optional message(s) are described in the Order Entry (`FIX_OE`) document.

---

## 2 Order Monitoring

Providing order drop copies makes it possible for participants or other bodies with a suitable FIX drop copy actor to follow the order flow in real time, either as a service to the own firm, the marketplace, a risk calculator, or other vehicle.

---

### 2.1 OrderMassStatusRequest (AF, in)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | AF (Added FIX.4.3) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants (trading firms, clearing firms), Marketplace, Supervisory bodies, Regulatory bodies |

**Purpose:** The order mass status request message requests the status for orders matching criteria specified within the request.

**Usage and Conditions:**

OrderMassStatusRequest can be requested for/by:
1. Trading firms to receive order copies of their own executions
2. Risk systems to receive order copies to guide risk calculations
3. Bodies granted order read to specified participants

The request can be made as a snapshot or snapshot+subscription to updates.

An actor can continue a discontinued subscription from where it left off when for example re-logging in (given the session is not aborted), by submitting the header field `LastMsgSeqNumProcessed (369)` and setting `369 = 0`.

If the requesting user has `OrderReadLevel = Acting` or `All` set in reference data, no extra parties need to be included in the subscription to get order events where the own participant acts in the capacity of entering or modifying firm.

**Limitations:**
- The order drop copy feature requires OE and DC gateway start configuration
- At least one PartyID for PartyRole = 1 Executing Firm must be supplied, this could be the own participant
- The submitted entries in the Parties group (i.e. Participants) are not validated at the time of request entry: The submitting party is retrieved from the session actor and necessary order read rights and order read level are evaluated as order events arrive

**Response/Acknowledgment:**
- Successful: `ExecutionReport (8)` copies according to the configured level and rights
- Not successful: `BusinessMessageReject (j)`

#### Message Structure — Table 3: OrderMassStatusRequest [AF], in (base)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = AF. Standard header fields apply, only LastMsgSeqNumProcessed (369) shown explicitly. |
| _start StandardHeader_ | | | | |
| 369 | LastMsgSeqNumProcessed | N | SeqNum | Setting this to 0 makes a subscription resume from last message received in the preceding subscription. |
| _end StandardHeader_ | | | | |
| 584 | MassStatusReqID | Y | String ISO-8859-1 [50] | Unique identifier for the OrderMassStatusRequest. |
| 585 | MassStatusReqType | Y | int Reserved100Plus | Supported values: `701` = SnapshotAllOrders, `702` = SnapshotAndUpdatesAllOrders, `703` = UnsubscribeAllOrders. Values 100 and up are ET custom. |
| 1012 | `<Parties>` group | Y (Req'd in JATS) | — | Add each participant (ExecutingFirm) whose all order events you want to receive. Requires order read rights. The requesting participant shall have OrderReadLevel = Acting or All. |
| _start Parties_ | | | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 1..50 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | Message related actor ID, can be: Firm |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Required if NoPartyIDs(453) > 0. Used to identify classification source. Supported values: `D` = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Required if NoPartyIDs(453) > 0. Identifies the type of PartyID(448). PartyRole 1 and 12 are required. Supported values: `1` = ExecutingFirm |
| _end Parties_ | | | | |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

### 2.2 ExecutionReport (8, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | 8 (Added FIX.2.7) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Purpose:** The execution report message is used to:
1. Confirm the receipt of an order
2. Confirm changes to an existing order (i.e. accept cancel and replace requests)
3. Relay order status information
4. Relay fill information on orders
5. Relay fill information on tradeable or restricted tradeable quotes
6. Reject orders
7. Respond to OrderMassStatusRequest (AF)

**Usage and Conditions:**

To acknowledge:
1. A new order
2. An updated (replaced) order
3. A cancelled order
4. A trade execution
5. A restated order

A Market to limit order `OrdType (40) = K` may be returned with a `Price (44)` immediately on submission if it matches upon entry or update.

`DisplayQty` is the quantity currently displayed, after a potential execution conveyed by the message.

**Limitations:**
- The conservation relation `OrderQty = CumQty + LeavesQty + CxlQty` is preserved. A potentially busted execution does not affect the OrderQty (CxlQty not currently used in JATS).
- For executions of MassQuote the relation does not hold since CumQty may increase beyond the BidSize/AskSize.
- For rejected order messages, decimal precision is not validated or adjusted.
- Parties group omitted for a rejected order.
- Trigger details are not included after a trigger order was triggered.
- Peg details are not included in full on executions or cancellations.
- ExecInst (18) and MinQtyMethod (1822) not included on executions or cancellations.
- OrdType (40) is omitted for matched quotes.

#### Message Structure — Table 4: ExecutionReport [8], out (DC-IDX-o)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | — |
| 37 | OrderID | Y | String ISO-8859-1 [18] | — |
| 11 | ClOrdID | N | String ISO-8859-1 [50] | Pass-thru field set by client and echoed back by marketplace. Not set on order entered by interface not supporting order token. In the case of quotes can be mapped to QuoteID(117) of a MassQuote(35=i). Can only be modified from the session that entered the order. Only populated when recipient session is held by the Firm submitting the order. |
| 41 | OrigClOrdID | N | String ISO-8859-1 [50] | Conditionally required for response to a Cancel or Cancel/Replace request (ExecType(150) = 6, 5, or 4) when referring to orders that were electronically submitted over FIX or otherwise assigned a ClOrdID(11). |
| 1012 | `<Parties>` group | N | — | Specifies party information related to the order owner and submitter of the request. Repeating group shall contain unique combinations (triplets) of PartyID, PartyIDSource, and PartyRole. |
| _start Parties_ | | | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | — |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Required if NoPartyIDs(453) > 0. Supported values: `D` = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Required if NoPartyIDs(453) > 0. Supported values: `1` = ExecutingFirm, `12` = ExecutingTrader, `24` = CustomerAccount, `36` = EnteringTrader, `7` = EnteringFirm, `53` = TraderMnemonic |
| → 2077 | `<PtysSubGrp>` group | N (Conditional) | — | — |
| → 802 | NoPartySubIDs | N | NumInGroup | Valid values: 1 for PartyRole = 1 (ExecutingFirm) |
| → → 523 | PartySubID | N | String | When PartySubIDType = OrderSource: Position 1: Char 0-9, A-Z (upper case only); Position 2-4: Char 0-9, A-Z and Space |
| → → 803 | PartySubIDType | N | int Reserved4000Plus | Supported values: `4030` = OrderSource. Values 4000 and up are JATS custom. |
| _end Parties_ | | | | |
| 880 | TrdMatchID | N | String | Match ID assigned by the matching engine. |
| 17 | ExecID | Y | String | An identifier assigned to each unique Execution Report message produced by the marketplace, without duplicates during the entire FIX session. |
| 150 | ExecType | Y | char | Specifies a step in the order's life cycle trajectory. Supported values: `0` = New, `4` = Canceled, `5` = Replaced, `8` = Rejected, `C` = Expired, `D` = Restated, `F` = Trade, `L` = TriggeredOrActivatedBySystem |
| 39 | OrdStatus | Y | char | Describes the current state of a CHAIN of orders. Supported values: `0` = New, `1` = PartiallyFilled, `2` = Filled, `4` = Canceled, `8` = Rejected, `C` = Expired |
| 1328 | RejectText | N | String | Reason description for rejecting the transaction request. |
| 378 | ExecRestatementReason | C | int Reserved100Plus | Required for ExecType = D (Restated). Supported values: `1` = GTRenewal, `4` = BrokerOption, `8` = Market, `17` = CxldSMP, `100` = OrderRejectedWhenTriggered, `101` = CancelledDueToCircuitBreaker, `102` = CancelledDueToCorporateAction, `103` = CxldSMPGroupDefault. Values 100+ are IDX custom. |
| 2220 | `<RegulatoryTradeIDGrp>` group | N | — | Only populated for ExecType = F, H |
| 1907 | NoRegulatoryTradeIDs | N | NumInGroup | Valid values: 1 |
| → 1903 | RegulatoryTradeID | N | String | Trading Venue Transaction Identifier (when 1906 = 5) |
| → 1906 | RegulatoryTradeIDType | N | int | Supported values: `5` = TradingVenueTransactionIdentifier |
| 1003 | `<Instrument>` component | Y | — | — |
| 55 | Symbol | Y | String ISO-8859-1 [32] | Short name of instrument. |
| 54 | Side | Y | char | Supported values: `1` = Buy, `2` = Sell, `5` = SellShort |
| 38 | OrderQty | N | Qty | Quantity ordered. |
| 40 | OrdType | N | char | Supported values: `1` = Market, `2` = Limit, `K` = MarketWithLeftOverAsLimit, `P` = Pegged |
| 44 | Price | N | Price | — |
| 1100 | TriggerType | N | char | Required if any other Triggering tags are specified. Supported values: `2` = SpecifiedTradingSession, `4` = PriceMovement |
| 1101 | TriggerAction | N | char | Supported values: `1` = Activate |
| 1102 | TriggerPrice | N | Price | Only relevant and required for TriggerType = 4 |
| 1103 | TriggerSymbol | N | String ISO-8859-1 [32] | For elaborated trigger orders where another orderbook is followed and governs triggering. |
| 1107 | TriggerPriceType | N | char | Supported values: `1` = BestOffer, `2` = LastTrade, `3` = BestBid |
| 1109 | TriggerPriceDirection | N | char | Supported values: `D` = Down, `U` = Up |
| 1113 | TriggerTradingSessionID | N | int Reserved100Plus | Only relevant and required for TriggerType = 2. |
| 211 | PegOffsetValue | N | int | Valid values: [-127…127] |
| 1094 | PegPriceType | N | int | Supported values: `2` = MidPricePeg, `4` = MarketPeg, `5` = PrimaryPeg |
| 835 | PegMoveType | N | int | Supported values: `0` = Floating, `1` = Fixed |
| 836 | PegOffsetType | N | char | Supported values: `2` = Ticks |
| 1096 | PegSecurityIDSource | N | char | Supported values: `M` = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 1097 | PegSecurityID | N | String ISO-8859-1 [10] | PegSymbol or PegSecurityID (OrderbookID) is only required when pegging order to other instrument than the order's (Symbol takes precedence). |
| 1098 | PegSymbol | N | String ISO-8859-1 [32] | Use [N/A] if PegSymbol is unknown or doesn't apply. |
| 839 | PeggedPrice | N | Price | The current price the order is pegged at. |
| 59 | TimeInForce | N | char | Absence of this field indicates Day order. Supported values: `0` = Day, `1` = GoodTillCancel, `3` = ImmediateOrCancel, `4` = FillOrKill, `6` = GoodTillDate |
| 432 | ExpireDate | N | LocalMktDate | — |
| 18 | ExecInst | N | MultipleCharValue | Supported values: `G` = AllOrNone, `i` = ImbalanceOnly, `0` = StayOnOfferSide, `6` = ParticipateDoNotInitiate, `9` = StayOnBidSide |
| 528 | OrderCapacity | N | char | Supported values: `A` = Agency, `G` = Proprietary, `I` = Individual, `P` = Principal, `R` = RisklessPrincipal |
| 529 | OrderRestrictions | N | MultipleCharValue | Supported values: `C` = IssuePriceStabilization |
| 2593 | NoOrderAttributes | N | NumInGroup | Valid values: 1 |
| → 2594 | OrderAttributeType | N | int Reserved1000Plus | Supported values: `1006` = Margin order |
| → 2595 | OrderAttributeValue | N | String ISO-8859-1 | Supported values: `Y` = Yes |
| 32 | LastQty | N | Qty | Quantity bought/sold on this (last) fill. Required if ExecType(150) = F (Trade) unless FillsGrp or OrderEventGrp is used. |
| 31 | LastPx | N | Price | Price of this (last) fill. Required if ExecType(150) = F (Trade). |
| 100 | ExDestination | N | Exchange | MIC code for requested venue of execution. |
| 336 | TradingSessionID | N | String Reserved100Plus | Supported values: `104` = GoodtillendofSessionsubtype. Values 100+ are JATS custom. |
| 625 | TradingSessionSubID | N | int Reserved100Plus ISO-8859-1 [32] | Session subtype upon discontinuation of which the order shall be cancelled. |
| 151 | LeavesQty | Y | Qty | Quantity open for further execution. |
| 14 | CumQty | Y | Qty | Currently executed quantity for this order chain. |
| 75 | TradeDate | N | LocalMktDate | Trade date according to the running trading session. |
| 60 | TransactTime | N | UTCTimestamp | Time the transaction represented by this ExecutionReport occurred. |
| 110 | MinQty | N | Qty | Min qty required to fill. Included only if not yet met. |
| 1822 | MinQtyMethod | N | int | Supported values: `1` = Once, `2` = Multiple |
| 2362 | SelfMatchPreventionID | N | String ISO-8859-1 [10] | As stated on the order, otherwise not present. |
| 1138 | DisplayQty | N | Qty | The quantity to be displayed. Required for reserve orders. |
| 77 | PositionEffect | N | char | Supported values: `C` = Close, `D` = Default, `O` = Open |
| 58 | Text | N | String ISO-8859-1 [15] | Passthrough text from order if accepted. |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

## 3 Trade Monitoring

Marketplace executions – be they from matched orders, negotiated quotes, or registered off-exchange deals – are disseminated as Trade Capture Reports (TCR). Appropriate parties such as brokers, clearing firms, back office functions, and other bodies to whom access have been granted may subscribe to this information via the FIX Post-trade service. Trade participants formally belong to this category as well but receive their execution TCR:s automatically.

---

### 3.1 TradeCaptureReportRequest (AD, in)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | AD (Added FIX 4.3) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants (trading firms, clearing firms), Marketplace, Supervisory bodies, Regulatory bodies |

**Purpose:** The Trade Capture Report Request can be used to:
- Request one or more trade capture reports based upon selection criteria
- Subscribe for trade capture reports based upon selection criteria

**Usage and Conditions:**

TradeCaptureReport can be requested for/by:
1. Trading firms to receive trade confirmations of their own executions
2. Clearing firms to receive drop copy trades for their associated trading firms
3. Back-office vendors to receive drop copy trades for their associated trading firms
4. Bodies granted trade read to specified participants

The request can be made as a snapshot or snapshot+subscription to updates.

An actor can continue a discontinued subscription from where it left off when for example re-logging in (given the session is not aborted), by submitting the header field `LastMsgSeqNumProcessed (369)` and setting `369 = 0`.

**Limitations:**
- TradeCaptureReportRequest is only provided for trades, order fills, trade cancel
- TradeCaptureReportRequest is only available for the current trade date of the respective trading session
- TradeRequestType = 0 (All eligible trades) is currently the only available business filtering option
- At least one PartyID for PartyRole = 1 Executing Firm must be supplied
- The submitted entries in the Parties group are not validated at the time of request entry

**Response/Acknowledgment:**
- Successful: `TradeCaptureReport (AE)` according to the configured level and rights + Successful `TradeCaptureReportAck (AQ)`
- Not successful: Negative `TradeCaptureReportAck (AQ)`

#### Message Structure — Table 5: TradeCaptureReportRequest [AD], in (base)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = AD |
| 369 | LastMsgSeqNumProcessed | N | SeqNum | Setting this to 0 makes a subscription resume from last message received in the preceding subscription. |
| 568 | TradeRequestID | Y | String ISO-8859-1 [50] | — |
| 569 | TradeRequestType | Y | int | Supported values: `0` = AllTrades |
| 263 | SubscriptionRequestType | N | char | Default if absent: Snapshot (0). Supported values: `0` = Snapshot, `1` = SnapshotAndUpdates, `2` = DisablePreviousSnapshot |
| 1012 | `<Parties>` group | Y (Req'd in JATS) | — | Add each participant (ExecutingFirm) whose all trade events you want to receive. Requires trade read rights. The requesting participant shall have OrderReadLevel = None, Acting, or All. |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 1..50 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | Message related actor ID — Firm |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Required if NoPartyIDs(453) > 0. Supported values: `D` = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Required if NoPartyIDs(453) > 0. PartyRole 1 and 12 are required. Supported values: `1` = ExecutingFirm |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

### 3.2 TradeCaptureReport (AE, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | AE (Added FIX.4.3, Updated FIX.5.0SP2 EP192) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants (trading firms, clearing firms), Marketplace, Supervisory bodies, Regulatory bodies |

**Purpose:** The Trade Capture Report message can be:
1. Used to report trades between counterparties
2. Used to report trades to a trade matching system
3. Sent unsolicited between counterparties
4. Sent as a reply to a Trade Capture Report Request
5. Used to report unmatched and matched trades

**Usage and Conditions:**
- Depending on configuration with respect to counterparty disclosure, the trade capture reports may show counterparty firm (Contra firm) in the parties group.
- A recipient may be entitled to receive messages for buy side, sell side or both.
- OrderID and ClOrdID are assigned to TradeCaptureReports also for reported trades.

**Limitations:**
- TradeCaptureReport is only sent for trades, order fills, trade cancel.
- A TradeCaptureReport may be sent as an aggregation of directly consecutive matches for which individual ExecutionReport:s were sent, typically for reserve orders. In such cases the TradeMatchID (880) takes on the value as of the first ExecutionReport, hence TradeID (1003) corresponds to more than 1 MatchID (880).

#### Message Structure — Table 6: TradeCaptureReport [AE], out (DC-IDX-o)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = AE |
| 571 | TradeReportID | Y (Req'd in JATS) | String ISO-8859-1 [50] | Identifier assigned by the exchange. |
| 1003 | TradeID | Y (Req'd in JATS) | String ISO-8859-1 [39] | Unique identifier for trade. |
| 856 | TradeReportType | Y (Req'd in JATS) | int | Supported values: `0` = Submit (Done trade), `7` = LockedInTradeBreak (Post-match cancellation / Bust) |
| 828 | TrdType | N | int Reserved1000Plus | TPAC (65) flag set on component trades from matches in combination order books. Supported values: `0` = RegularTrade, `65` = PackageTrade |
| 1123 | TradeHandlingInstr | Y (Req'd in JATS) | char | Supported values: `0` = TradeConfirmation |
| 1126 | OrigTradeID | N | String | Used to refer to original trade in case of modifications. |
| 150 | ExecType | N (Req'd in JATS) | char | Supported values: `F` = Trade, `H` = TradeCancel |
| 820 | TradeLinkID | N | String | ID that will be the same for all legs of a reported multileg trade report. |
| 880 | TrdMatchID | Y (Req'd in JATS) | String | ID assigned to a trade by the matching system. |
| 527 | SecondaryExecID | N (Req'd in JATS) | String | The deal ID from the matching system. |
| 1907 | NoRegulatoryTradeIDs | N | NumInGroup | Valid values: 1 |
| → 1903 | RegulatoryTradeID | N | String | Trading Venue Transaction Identifier (when 1906 = 5) |
| → 1906 | RegulatoryTradeIDType | N | int | Supported values: `5` = TradingVenueTransactionIdentifier |
| 570 | PreviouslyReported | N | Boolean | Supported values: `N` = NotReportedToCounterparty, `Y` = PreviouslyReportedToCounterparty |
| 1430 | VenueType | N | char | Supported values: `B` = CentralLimitOrderBook, `E` = Electronic, `N` = QuoteNegotiation, `O` = OffMarket |
| 1300 | MarketSegmentID | N | String | The home listing Market Segment. |
| 1301 | MarketID | N | Exchange | The home listing Market. |
| 55 | Symbol | N | String ISO-8859-1 [32] | Short name of instrument. |
| 48 | SecurityID | N | String ISO-8859-1 [10] | Orderbook ID. |
| 22 | SecurityIDSource | N | String Reserved100Plus ISO-8859-1 [32] | Conditionally req'd if SecurityID is supplied. Supported values: `M` = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 454 | NoSecurityAltID | N | NumInGroup | Valid values: 1 |
| → 455 | SecurityAltID | N | String | — |
| → 456 | SecurityAltIDSource | N | String Reserved100Plus | Supported values: `4` = ISINNumber |
| 167 | SecurityType | N | String ISO-8859-1 [32] | Supported values: `BUYSELL` = BuySellback, `REPO` = Repurchase, `SECLOAN` = SecuritiesLoan, `MLEG` = MultilegInstrument, `NONE` = NoSecurityType |
| 762 | SecuritySubType | N | String | Set to off-book type for a reported trade. Only populated for a reported trade. |
| 1193 | SettlMethod | N | String | Settlement method of trade. Supported values: `1` = DvP, `2` = DFoP |
| 917 | EndDate | N | LocalMktDate | Return/Repurchase date. Req'd for Equity Repo or FI trade. |
| 32 | LastQty | Y (Req'd in JATS) | Qty | — |
| 31 | LastPx | Y (Req'd in JATS) | Price | Trade price. |
| 30 | LastMkt | Y (Req'd in JATS) | Exchange | Valid values: [MIC Code], `SINT` = off-exchange on systematic internaliser, `XOFF` = other off-exchange |
| 75 | TradeDate | Y (Req'd in JATS) | LocalMktDate | — |
| 60 | TransactTime | Y (Req'd in JATS) | UTCTimestamp | — |
| 64 | SettlDate | N | LocalMktDate | — |
| 573 | MatchStatus | Y (Req'd in JATS) | char | Supported values: `0` = Compared |
| 574 | MatchType | N | String | Valid values for trades outside the CLOB order book (1430 != B): `1` = OnePartyTradeReport, `2` = TwoPartyTradeReport, `3` = ConfirmedTradeReport, `10` = AutoMatchLastLook |
| 552 | NoSides | N | NumInGroup | Valid values: 1 |
| → 54 | Side | Y | char | Supported values: `1` = Buy, `2` = Sell, `5` = SellShort |
| → 453 | NoPartyIDs | N | NumInGroup | — |
| → → 448 | PartyID | Y | String ISO-8859-1 [32] | — |
| → → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: `D` = PROPRIETARY_CUSTOM_CODE |
| → → 452 | PartyRole | Y | int | Supported values: `1` = ExecutingFirm, `12` = ExecutingTrader, `24` = CustomerAccount, `36` = EnteringTrader, `7` = EnteringFirm, `17` = ContraFirm, `4` = ClearingFirm |
| → → 802 | NoPartySubIDs | N | NumInGroup | Valid values: 1 for PartyRole = 1 |
| → → → 523 | PartySubID | N | String | Position 1: Char 0-9, A-Z (upper case); Position 2-4: Char 0-9, A-Z and Space |
| → → → 803 | PartySubIDType | N | int Reserved4000Plus | Supported values: `4030` = OrderSource. Values 4000+ are JATS custom. |
| → 625 | TradingSessionSubID | N | int Reserved100Plus | The session state where orders matched for the trade. |
| → 921 | StartCash | N | Amt | Req'd for Repo trade. |
| → 922 | EndCash | N | Amt | Req'd for Repo trade. |
| → 58 | Text | N | String ISO-8859-1 [15] | Passthrough text. |
| → 1057 | AggressorIndicator | N | Boolean | Supported values: `N` = OrderInitiatorIsPassive, `Y` = OrderInitiatorIsAggressor |
| → 1115 | OrderCategory | N | char | Supported values: `1` = Order, `2` = Quote, `3` = PrivatelyNegotiatedTrade, `6` = QuoteRequest |
| → 1851 | StrategyLinkID | N | String | Complex Trade Component ID (match group). Common value to all legs of an electronically matched combination. |
| → 37 | OrderID | Y (Req'd in JATS) | String ISO-8859-1 [18] | OrderID assigned by JATS. |
| → 11 | ClOrdID | N | String ISO-8859-1 [50] | Pass-thru field set by client and echoed back by JATS. Only populated when recipient session is held by the Firm submitting the order. |
| → 528 | OrderCapacity | N | char | Supported values: `A` = Agency, `G` = Proprietary, `I` = Individual, `P` = Principal, `R` = RisklessPrincipal |
| → 2593 | NoOrderAttributes | N | NumInGroup | Valid values: 1 |
| → → 2594 | OrderAttributeType | N | int Reserved1000Plus | Supported values: `1006` = Margin order |
| → → 2595 | OrderAttributeValue | N | String ISO-8859-1 | Supported values: `Y` = Yes |
| → 483 | TransBkdTime | N | UTCTimestamp | — |
| → 1856 | RelatedTradeID | N | String ISO-8859-1 [50] | The ID of a submitted Trade Report or Secondary trade ID from Trade report Ack. |
| → 1857 | RelatedTradeIDSource | N | int | Supported values: `0` = NonFIXSource, `1` = TradeID, `2` = SecondaryTradeID, `3` = TradeReportID |
| 797 | CopyMsgIndicator | N | Boolean | Included and set to 'Y' on Drop Copy messages. |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

### 3.3 TradeCaptureReportRequestAck (AQ, out)

| Property | Value |
|----------|-------|
| **Message type** | AQ (Added FIX.4.4, Updated FIX.5.0SP2 EP192) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants (trading firms, clearing firms), Marketplace, Supervisory bodies, Regulatory bodies |

**Purpose:** The TradeCaptureRequestAck message is used to:
1. Provide an acknowledgement to a Trade Capture Report Request where it is used to specify a subscription or delivery of reports via an out-of-band ResponseTransmissionMethod
2. Provide an acknowledgement when the return of Trade Capture Reports will be delayed or delivered asynchronously
3. Indicate that no trades were found that matched the selection criteria, or that the request was invalid for business reasons

**Usage and Conditions:**

TradeCaptureReportRequestAck is used to:
1. Acknowledge a TradeCaptureRequest (which could be to abort a subscription)
2. Reject a TradeCaptureRequest

**Limitations:**
- TradeCaptureReport is only provided for trades, order fills, trade cancel
- TradeCaptureReport is only available for the current trade date of the respective trading session
- TradeRequestType = 0 (All eligible trades) is the only available business filtering option

#### Message Structure — Table 7: TradeCaptureReportRequestAck [AQ], out (base)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = AQ |
| 568 | TradeRequestID | Y | String ISO-8859-1 [50] | Unique identifier for the trade request. |
| 569 | TradeRequestType | Y | int | Supported values: `0` = AllTrades |
| 263 | SubscriptionRequestType | N | char | Supported values: `0` = Snapshot, `1` = SnapshotAndUpdates, `2` = DisablePreviousSnapshot |
| 748 | TotNumTradeReports | N | int | Number of trade reports returned (if known). |
| 749 | TradeRequestResult | Y | int Reserved100Plus | Result of Trade Request. Supported values: `0` = Successful, `8` = TradeRequestTypeNotSupported, `9` = NotAuthorized, `99` = Other |
| 750 | TradeRequestStatus | Y | int | Status of Trade Request. Supported values: `0` = Accepted, `1` = Completed, `2` = Rejected |
| 58 | Text | N | String ISO-8859-1 [15] | Additional rejection description. |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

## 4 Other Information

### 4.1 News (B, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | B (Added FIX.2.7) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants (trading firms, clearing firms), Marketplace, Supervisory bodies, Regulatory bodies |

**Purpose:** The news message is a general free format message between the broker and institution.

**Usage and Conditions:**

The message is only used to communicate a barrier, i.e. an opening or closing event. The end of upstream business is communicated this way and shall be interpreted as a notification of a completed business day and the last message has been sent downstream. Specifically, a clearinghouse shall not expect any further messages.

#### Message Structure — Table 8: News [B], out (base)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = B |
| 42 | OrigTime | N | UTCTimestamp | Time of message origination as a UTC timestamp. |
| 61 | Urgency | N | char | Supported values: `0` = Normal |
| 1472 | NewsID | Y (Req'd in JATS) | String | Unique identifier of news message. |
| 1473 | NewsCategory | N | int Reserved100Plus | Supported values: `101` = BarrierNewsEndOfBusiness. Values 100+ are JATS custom. |
| 148 | Headline | Y | String | News headline — "End of Business" |
| 33 | NoLinesOfText | N | NumInGroup | Number of text lines. Valid values: 1..10 |
| → 58 | Text | Y | String ISO-8859-1 [80] | Max 80 characters per line. Line 1: "Last business message has been sent downstream for business date:". Line 2: businessDateDetail.date (converted to "YYYY-MM-DD" format). |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

### 4.2 BusinessMessageReject (j, out)

| Property | Value |
|----------|-------|
| **Message type** | j (Added FIX.4.2) |
| **FIX Session** | Drop Copy (DC) |
| **Available to** | Participants |

**Purpose:** The Business Message Reject message can reject an application-level message which fulfills session-level rules and cannot be rejected via any other means. Note if the message fails a session-level rule (e.g. body length is incorrect), a session-level Reject message should be issued.

**Usage and Conditions:**

The Business Message Reject message is used when the NFF Marketplace cannot honor the following MD requests:
- `SecurityMassStatusRequest (CN, in)`

#### Message Structure — Table 9: BusinessMessageReject [j], out (MME-C-o)

| FIX Tag No | FIX Tag Name | Req'd | Type | Description |
|------------|--------------|-------|------|-------------|
| 1024 | `<StandardHeader>` component | Y | — | MsgType = j |
| 45 | RefSeqNum | N | SeqNum | MsgSeqNum of rejected message. |
| 372 | RefMsgType | Y | String | The MsgType of the FIX message being referenced. |
| 380 | BusinessRejectReason | Y | int | Supported values: `0` = OTHER, `1` = UNKNOWN_ID, `5` = CONDITIONALLY_REQUIRED_FIELD_MISSING, `6` = NOT_AUTHORIZED |
| 58 | Text | N | String ISO-8859-1 [15] | Where possible, message to explain reason for rejection. |
| 1025 | `<StandardTrailer>` component | Y | — | — |

---

*© 2025 Indonesia Stock Exchange — Private and Confidential*
