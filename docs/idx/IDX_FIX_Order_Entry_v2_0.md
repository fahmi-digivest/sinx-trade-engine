# IDX DOCUMENT SPECIFICATION — FIX ORDER ENTRY

**Document Version:** 2.0  
**Date:** 14 Oct 2025  
**Issuer:** Indonesia Stock Exchange (Bursa Efek Indonesia)  
**Classification:** Private and Confidential

---

## Summary of Changes

| Version | Date | List of Changes | Details |
|---------|------|-----------------|---------|
| 2.0 | 14/10/2025 | Added: SID / Client Account | Information on the use of SID / Client Account for each order |
| 2.0 | 03/10/2025 | Added: Indicative Quotation (AI, out, reject) | QuoteStatusReport (AI, out, indicative, reject) |
| 2.0 | 03/10/2025 | Removed: Indicative Quotation (AI, in & out) | QuoteStatusReport (AI, in), QuoteStatusReport (AI, out) |
| 2.0 | 03/10/2025 | Removed: Request Quotation | Removed the specification regarding RFQ |
| 2.0 | 22/09/2025 | Added Order Source Information | 1.2.11 Order Source |
| 2.0 | 22/09/2025 | Added (Errata): OrderCancelRequest (F, in), OrderCancelReplaceRequest (G, in) | SecurityID (48), SecurityIDSource (22) |
| 2.0 | 22/09/2025 | Added: SecurityDefinitionRequest (c, in), SecurityDefinition (d, out), BusinessMessageReject (j, out) | Support creation of Repos and TMCs |
| 2.0 | 22/09/2025 | Changed: QuoteCancel (Z) | QuoteType (537) is required |
| 2.0 | 22/09/2025 | Changed: TradeCaptureReportAck (AR), out, 2-sided | Marked 167, 762, 573 as not req'd on rejection |
| 2.0 | 22/09/2025 | Added: TradeCaptureReportAck (AR), out, 2-sided | F = Trade to ExecType (150) |
| 2.0 | 22/09/2025 | Added: TradeCaptureReport (AE), in, 2-sided, 1-sided | TransactTime (60) |
| 2.0 | 22/09/2025 | Added: OrderCancelReplaceRequest (G), OrderCancelRequest (F), ExecutionReport (8), Section: How to update or cancel an order in different scenarios | PartyRole (452) = TraderMnemonic (53) |
| 2.0 | 22/09/2025 | Added: NewOrderSingle (D), ExecutionReport (8) | ExDestination (100) |
| 2.0 | 22/09/2025 | Changed: New layout of Spec | All content, FIX DC messages moved to separate spec |
| 2.0 | 22/09/2025 | Added: MassQuote (i), MassQuoteAck (b) | QuoteResponseLevel = 2 to QuoteResponseLevel (301), Affected entries in message properties table |
| 2.0 | 22/09/2025 | Added: NewOrderSingle (D), ExecutionReport (8) | ExecInstr (18) = 0, 6, 9 for Post only orders |
| 2.0 | 22/09/2025 | Changed: QuoteCancel (Z) | QuoteType (537) is required |
| 1.0 | 05/03/2025 | First release | |

---

## Table of Contents

1. [About the Document](#1-about-the-document)
2. [Order Management](#2-order-management)
3. [Continuous Quotation](#3-continuous-quotation)
4. [Indicative Quotation](#4-indicative-quotation)
5. [Trade Matching](#5-trade-matching)
6. [Deal Reporting](#6-deal-reporting)
7. [Rejections](#7-rejections)
8. [Appendices](#8-appendices)

---

## 1. About the Document

### 1.1 About this Specification's Status and Refinement

This specification depicts the FIX Order, Quote, and Trade entry service for the JATS (Jakarta Automated Trading System). The full suite of specifications includes Session layer, Reference Data, Order and Trade Entry, Order and Trade Monitoring (Drop Copy) and Market Data. A condensed and consistent set of tags is included with the pronounced aim to meet demands as imposed by authorities and regulation regimes under which financial bodies operate at each individual location.

As far as possible standard FIX tags and codesets are utilized and where not, custom tags and values have been added.

### 1.2 Concepts and Legend

#### 1.2.1 Characteristics Given for a Message

Each message is documented with a leading summary of characteristics and use cases.

| Characteristic | Content |
|----------------|---------|
| Direction | Can be "In to JATS" or "out from JATS". If there can be both variants, separate message sections are used. |
| Message code (FIX standard) introduced | The FIX message type: MsgType, tag 35 which is located in the header. The FIX version the message was introduced. Older messages are often less open to custom extensions. |
| FIX Session | The default session over which the message (and its containing flow) is communicated, for example "OE" (Order Entry session). |
| Available to | Defines the intended categories of business users. |
| Usage and Conditions | Contains information on specific use cases of relevance for the intended marketplace and any specific conditions that apply. |
| Limitations | Specifies limitations that may not be obvious, need attention or should be considered when designing operational processes. |
| Response/Acknowledgment | The immediate responses to expect for the successful and unsuccessful cases. |

#### 1.2.2 Message Chaining or Entity Chaining

FIX offers two principles of chaining a sequence of orders, trades, and other successive events: the **message chaining model** and the **entity chaining model**.

- The message chaining model treats the event and its related message as "the same thing" and a successor message refers to its preceding in the chain by a message reference which implicitly becomes a reference to the event which entailed the message about it.
- The entity reference model references the entity the event created, such as a trade which is treated as something that exists in its own right, not in need of message proxy. In this model the trade is referenced directly.

Where not required to do otherwise this implementation adheres to the **entity chaining principle**. This manifests itself when cancelling a trade for example where, to this end, TradeID (1003) is more important than TradeReportID (571). The motivation is that in a system that makes use of several protocols, potentially partially overlapping, any entity, such as a trade or order, must be able to address unambiguously without referring to a specific protocol.

There are however situations where a TradeID is not yet established, for example when reporting an off-exchange trade: the marketplace (report recipient) has yet to create the TradeID upon a successful validation of the reported content. Until this is done, a suitable way to refer to the trade is by a reference to the TradeReportID (if submitted via FIX).

#### 1.2.3 About the Req'd (Required) Column

The following nomenclature is used for the Req'd column:

- **Y**: The field is required (must have a value) by FIX standard or is mandated in JATS. In the latter case this is stated as a remark in the column: "Req'd in JATS".
- **N**: The field is not required, i.e. optional, which is the default if not declared. It can be conditional in JATS to distinguish different use cases, which is stated as a remark in the column: "Cond'l in JATS".

The scope of the Req'd property is within the closest enclosure, which is the group, component, or message, in the given scenario.

Where a field is required because of FIX syntax, for example being first field in repeating group or to make an ID designation complete (Party, Security, any type/value construct, etc), this isn't made explicit in the Req'd column.

#### 1.2.4 About the UD (Union Datatype) Column

The FIX standard opens for free to use enumeration ranges for some codesets. These are indicated by the union datatype (UD) feature where custom values can be freely added (and agreed between all parties) without breaking the protocol.

UD values can be `100+`, `1000+`, and `4000+` where numbers above (and including) these limits are reserved for implementation dependent use.

#### 1.2.5 About the Parties Group (453)

The Parties group is common to almost all business level messages, and is vital to all messages directed towards trading, like orders, quotes, trade reports and so on. The difference as compared to the CompID:s in the header is that the parties in the parties group define or reflect the business level stakeholders whereas the header CompID:s are pure technical actors.

A party definition is composed of 3 parts: **PartyID**, **PartyIDSource**, **PartyRole**. A party definition isn't valid syntactically if not all 3 parts are defined with valid values. Some parties, holding specific roles, are mandatory to provide while others are optional. If mandatory roles are missing the message is rejected.

**Table 1: Summary of Parties Group Usage**

| Entity Description | Tag 448 PartyID | Tag 447 PartyIDSource | Tag 452 PartyRole | Req'd | Val'd |
|---|---|---|---|---|---|
| Defines the number of party definitions that follows in the parties group. Mandatory in FIX (but not FIXML). Must be >= 1 if present. | No of party definitions | | | Y | Y |
| Firm owning the transaction business wise | Firm name in JATS | D | 1 = ExecutingFirm | Y | Y |
| User owning the transaction business wise | User name in JATS | D | 12 = ExecutingTrader | Y | Y |
| Account or Client to which the transaction is accountable. Only the first 15 characters are used and the last 16th character shall be a null character. The 3rd character must be either "D" (for domestic) or "F" (for foreign) | Account/Client designation | D | 24 = CustomerAccount | N | N |
| Counterparty to the transaction. Required or optional for certain types of transactions | Firm name in JATS | D | 17 = ContraFirm | C | Y |
| If the transaction is submitted via a firm different from the owning firm, and that firm is not only a re-routing hub. Not all transaction types allow for this type of on-behalf-of submission | Firm name in JATS | D | 7 = EnteringFirm | C | Y |
| The user at the on-behalf-of submitter firm sending the transaction | User name in JATS | D | 36 = EnteringTrader | C | Y |
| A user who can act on an order but is not part of the order's entry | User name in JATS | D | 53 = TraderMnemonic | C | Y |
| The firm eligible or required to quote or respond to quote requests for a certain instrument | Firm name in JATS | D | 66 = MarketMaker | N | Y |
| The firm responsible for clearing a trade for an ExecutingParty | Firm name in JATS | D | 4 = ClearingFirm | N | Y |

All the PartyIDSource values are "D", however when invoking specific regulations such as MiFID, or other domains of designation, other symbols come into use as well, such as "P" or "G".

#### 1.2.6 About Price and Quantity

For Price and Quantity precision, the number of decimals for each determine the valid amounts. They are fetched from the SecurityList message as Instrument price precision and Instrument quantity precision and are entered with a decimal point. Examples:
- 44 (Price): `1.012345`
- 38 (OrderQty): `2.03`

#### 1.2.7 Quantities in Units or Lots

All quantities fields have quantities expressed in units or lots. Tag `InstrAttribType (871)` / `InstrAttibValue (872)` in SecurityList and SecurityListUpdateReport message is used to determine if quantities for an order book are expressed in units or lots.

#### 1.2.8 About Length of String Fields

FIX does not constrain string fields to any length. JATS does impose limits to the number of characters that can be entered in a string type tag where there are no pre-declared values to choose from. This is indicated as `[X]` in the type column together with the label ISO-8859-1. In practice this means X is the maximum number of characters provided they are coded into 1 byte.

#### 1.2.9 Gateway Configuration Options

The specification volume referred to as "FIX_SG", the JATS FIX specification for session and gateway contains a section about configuration options for the various gateways. Declaring particular configuration settings at startup will enable additional messages or services to the gateway.

#### 1.2.10 About the Scenario Feature

Scenarios allow for messages documentation to look different depending on use case and only contain the fields and values pertaining to it. Scenario codes look like `P-A-D-S` where the constituents together represent a use case:

1. **1st part**: Product wide or gateway scope (ET (JATS), OE)
2. **2nd part**: Product adaptation (C=Core, Xyz=adaptation Xyz, e.g. an exchange, marketplace, or other installation)
3. **3rd part**: Direction (i/o/io)
4. **4th part**: Specialization of any kind as needed; particular FIX msg (enclosed in `[ ]`), other special conditions

The default scenario is just called "Base".

#### 1.2.11 Order Source

**Table 2: List of Order Source**

| Digit 1 (application) | Digit 2 (feature) | Digit 3 (platform) | Digit 4 (AO) |
|---|---|---|---|
| R = Remote Trading | E = ETF | A = Desktop | G = Automated Ordering |
| O = Online Trading | S = Sharia | B = Web base | Z = Other |
| D = DMA | P = MSOP/ESOP | F = Mobile | |
| M = MPPE | C = One Day Trade | Z = Other | |
| Z = Other | Z = Other | | |

### 1.3 Messages and Scenarios in this Document

**Table 3: List of Messages**

| Name | Scenario | Description |
|------|----------|-------------|
| OrderCancelReject | base | Order cancel reject message is received for a cancel request or cancel/replace request message which cannot be honored |
| QuoteCancel | OE-C-i-QR | Quote submitter cancels quote for quote request |
| MassQuoteAck | base | Ack returned from JATS for an entered MassQuote message |
| QuoteRequestReject | base | Conveys a quote request rejection from JATS |
| QuoteStatusReport | OE-C-o-QRRej | JATS rejects an invalid message entered as part of quote request workflow |
| QuoteResponse | OE-IDX-i | Enter response to a received quote for a quote request |
| QuoteResponse | OE-C-o | JATS disseminates a submitted quote response to market participants for a quote request |
| TradeCaptureReportAck | OE-C-o-2side | JATS accepts or rejects a 2-sided trade report |
| TradeCaptureReportAck | OE-C-o-1side | JATS accepts or rejects a 1-sided trade report |
| QuoteResponse | OE-C-o-IQ | JATS disseminates response to an indicative quote to quote submitter |
| QuoteStatusReport | OE-C-i-IQ | Indicative quote submitter accepts or declines an offer |
| QuoteStatusReport | OE-C-oIQNote | JATS conveys information about an indicative quote or offer being accepted, declined, or cancelled to the intended recipient |
| OrderCancelRequest | base | Request a cancellation of an order, all variants |
| Quote | OE-C-o | JATS disseminates quote for quote request to requestor |
| QuoteCancel | OE-C-i-IQ | Quote submitter cancels indicative quote |
| QuoteRequest | OE-C-o | JATS disseminates quote request to market makers/participants |
| QuoteStatusReport | OE-C-oQRCanc | JATS conveys to requestor a cancelled quote from respondent (quote submitter) for a quote request |
| QuoteStatusReport | OE-C-o-IQRej | JATS rejects an invalid message entered as part of indicative quote workflow |
| SecurityDefinition | base | |
| SecurityDefinitionRequest | RD-C-i-REP | |
| SecurityDefinitionRequest | RD-C-i-TMC | |
| TradeCaptureReportAck | OE-C-o-OnHold | JATS notifies about the on-hold status |
| NewOrderSingle | OE-IDX-i | Enter order, all order types and variations |
| OrderCancelReplaceRequest | OE-IDX-i | Update an order, all types and variations |
| MassQuote | OE-IDX-i | Enter MassQuote for 1..10 instruments |
| QuoteRequest | OE-IDX-i | Enter quote request, single or double sided |
| Quote | OE-IDX-i-IQ | Enter indicative quote |
| QuoteResponse | OE-IDX-i-IQ | Enter response to an indicative quote |
| TradeCaptureReport | OE-IDX-i-1side | Enter 1-sided trade report for matching with counterparty |
| TradeCaptureReport | OE-IDX-o-1side | JATS propagates a 1-sided trade report to an assigned counterparty, to match against or to notify about a cancellation |
| TradeCaptureReport | OE-IDX-i-2side | Enter 2-sided trade report with one or more trade legs |
| Quote | OE-IDX-i | Enter quote to respond to quote request |
| ExecutionReport | OE-IDX-o | Execution Report from an order event in JATS |
| BusinessMessageReject | ET-C-o | |

> **Note:** The abbreviation "JATS" is, for space reasons in tables and lists, used for short denoting "Jakarta Automated Trading System".

### 1.4 Gateway Configuration and Optional Messages

The specification volume referred to as "FIX_SG", the JATS FIX specification for session and gateway, contains a section about configuration options for the various gateways. Declaring particular configuration settings at startup will enable additional messages or services to the gateway.

Optional messages that can be enabled on the Order Entry gateway:
- `MarketDefinitionRequest (BT)`
- `MarketDefinition (BU)`
- `TradingSessionListRequest (BI)`
- `TradingSessionList (BJ)`
- `SecurityListRequest (x)`
- `SecurityList (y)`
- `SecurityListUpdateReport (BK)`

The optional message(s) are described in the Reference Data (FIX_RD) document.

---

## 2. Order Management

### 2.1 NewOrderSingle (D, in)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | D (Added FIX.2.7) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Purpose:** The new order message type is used by institutions wishing to electronically submit securities and forex orders to a broker for execution.

**Usage and Conditions:**  
Used by a trading user at a firm to submit a new order with unique ClOrdID in a specific instrument. Attributes determining order behavior include: OrdType (40), DisplayQty (1138), TimeInForce (59), ExecInst (18), TradingSessionSubID (625), TriggerType (1100), TriggerPrice (1102), PegMoveType (835), MinQtyMethod (1822), Text (58).

- Post-only (passive) orders will not match at entry; use ExecInst (18) = one of combinations "6", "6 0", "6 9".

**Limitations:**
- The applicability of on-behalf-of designations (EnteringFirm, EnteringTrader) is dependent on user entitlements configured in the system.
- If executing and entering trader are different albeit with same firm, EnteringFirm still must be specified.
- Setting DisplayQty (1138) = 0 is interpreted as no hidden volume, i.e. the full order quantity is displayed to the market. Omitting the field gives the same result.
- SelfMatchPreventionID (2362) must be enabled for the participant to take effect. The tag only accepts characters 0…9.
- For trigger orders, only one trigger type is allowed per order.
- Trigger orders that meet trigger conditions already at entry may receive a rejection instead of cancellation if not appropriate (fails validation) to go into the book.
- FoK orders prohibited to match due to self match prevention leave the passive side stay in the book even if passive is set to be cancelled.
- Peg offset sign (+ or -) adds (+) in the generous direction, e.g. +2 for a sell order will lower the peg price from the price the peg order is based on (and the opposite for buy orders).

**Response/Acknowledgment:** ExecutionReport (8)

**Table 4: NewOrderSingle [D], in (OE-IDX-i) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 11 | ClOrdID | Y | String ISO-8859-1 [50] | |
| | `<Parties>` group | Y (Req'd in JATS) | | Specifies party information related to the order owner and submitter of the request. Repeating group shall contain unique combinations (triplets) of PartyID, PartyIDSource, and PartyRole |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Required if NoPartyIDs(453) > 0. Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Required if NoPartyIDs(453) > 0. Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm |
| → 802 | NoPartySubIDs | N | NumInGroup | Valid values: 1 for PartyRole = 1 (ExecutingFirm) |
| → → 523 | PartySubID | N | String | When PartySubIDType = OrderSource: Position 1: Character 0-9, A-Z (Upper case only); Position 2-4: Character 0-9, A-Z and Space |
| → → 803 | PartySubIDType | N | int Reserved4000Plus | Supported values: 4030 = OrderSource (refer to 1.2.11) |
| 75 | TradeDate | N | LocalMktDate | Can be entered if preferred by the firm, not used in JATS |
| 18 | ExecInst | N | MultipleCharValue | Can contain multiple instructions, space delimited. Supported values: G = AllOrNone, i = ImbalanceOnly, 0 = StayOnOfferSide, 6 = ParticipateDoNotInitiate, 9 = StayOnBidSide |
| 110 | MinQty | N | Qty | Min qty required in fills according to MinQtyMethod (1822) given |
| 1822 | MinQtyMethod | N | int | Supported values: 1 = Once (first fill, default), 2 = Multiple (each fill is a single execution) |
| 2362 | SelfMatchPreventionID | N | String ISO-8859-1 [10] | |
| 1138 | DisplayQty | N | Qty | The quantity to be displayed. Required for reserve orders. |
| 100 | ExDestination | N | Exchange | MIC code for requested venue of execution. |
| 386 | NoTradingSessions | N | NumInGroup | Valid values: 1 |
| → 336 | TradingSessionID | N | String Reserved100Plus | Required if NoTradingSessions is > 0. Supported values: 104 = GoodtillendofSessionsubtype |
| → 625 | TradingSessionSubID | N | int Reserved100Plus | Session subtype upon discontinuation of which the order shall be cancelled |
| 55 | Symbol | Y (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument. If other value than "[N/A]" is supplied, then Symbol has precedence over Tag 48 values. |
| 48 | SecurityID | N (Required when Symbol=[N/A]) | String ISO-8859-1 [10] | Orderbook ID, ignored if Symbol != [N/A] |
| 22 | SecurityIDSource | N | String Reserved100Plus | Conditionally req'd if SecurityID is supplied. Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| 60 | TransactTime | Y | UTCTimestamp | |
| 38 | OrderQty | Y | Qty | Quantity ordered |
| 40 | OrdType | Y | char | Supported values: 1 = Market, 2 = Limit, K = MarketWithLeftOverAsLimit |
| 44 | Price | N | Price | |
| 1100 | TriggerType | N | char | Required if any other Triggering tags are specified. Supported values: 2 = SpecifiedTradingSession, 4 = PriceMovement |
| 1101 | TriggerAction | N | char | Supported values: 1 = Activate |
| 1102 | TriggerPrice | N | Price | The price at which the order shall trigger. Only relevant and required for TriggerType = 4 |
| 1103 | TriggerSymbol | N | String ISO-8859-1 [32] | For elaborated trigger orders where another orderbook is followed |
| 1107 | TriggerPriceType | N | char | Supported values: 1 = BestOffer, 2 = LastTrade, 3 = BestBid |
| 1109 | TriggerPriceDirection | N | char | Supported values: D = Down (trigger if price goes DOWN to or through the specified Trigger Price), U = Up (trigger if price goes UP to or through the specified Trigger Price) |
| 1113 | TriggerTradingSessionID | N | int Reserved100Plus | Defines the session subtype at which the order will be activated. Only relevant and required for TriggerType = 2. |
| 59 | TimeInForce | N | char | Absence of this field indicates Day order. Supported values: 0 = Day (default), 3 = ImmediateOrCancel (IoC/FaK), 4 = FillOrKill (FoK) |
| 432 | ExpireDate | N | LocalMktDate | |
| 528 | OrderCapacity | N | char | Supported values: A = Agency, G = Proprietary, I = Individual, P = Principal, R = RisklessPrincipal |
| 529 | OrderRestrictions | N | MultipleCharValue | Supported values: C = IssuePriceStabilization |
| 2593 | NoOrderAttributes | N | NumInGroup | Valid values: 1 |
| → 2594 | OrderAttributeType | N | int Reserved1000Plus | Supported values: 1006 = Margin order |
| → 2595 | OrderAttributeValue | N | String ISO-8859-1 | Supported values: Y = Yes |
| 58 | Text | N | String ISO-8859-1 [15] | |
| 77 | PositionEffect | N | char | Supported values: C = Close, D = Default, O = Open |
| 211 | PegOffsetValue | N | int | Valid values: [-127…127] |
| 1094 | PegPriceType | N | int | Supported values: 2 = MidPricePeg, 4 = MarketPeg, 5 = PrimaryPeg |
| 835 | PegMoveType | N | int | Supported values: 0 = Floating, 1 = Fixed |
| 836 | PegOffsetType | N | char | Supported values: 2 = Ticks |
| 1096 | PegSecurityIDSource | N | char ISO-8859-1 [32] | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 1097 | PegSecurityID | N | String ISO-8859-1 [10] | Required when PegSymbol=[N/A] |
| 1098 | PegSymbol | N | String Reserved100Plus | Use [N/A] if PegSymbol is unknown or doesn't apply |
| | `<StandardTrailer>` component | Y | | |

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

**Usage and Conditions:** Acknowledges: A new order, An updated (replaced) order, A cancelled order, A trade execution, A restated order.

**Limitations:**
- The conservation relation: `OrderQty = CumQty + LeavesQty + CxlQty` is preserved.
- For executions of MassQuote the relation does not hold since CumQty may increase beyond the BidSize/AskSize.
- For rejected order messages, decimal precision is not validated or adjusted.
- Parties group omitted for a rejected order.
- Trigger details are not included after a trigger order was triggered.
- Peg details are not included in full on executions or cancellations.
- ExecInst (18) and MinQtyMethod (1822) not included on executions or cancellations.
- OrdType (40) is omitted for matched quotes.

**Table 5: ExecutionReport [8], out (OE-IDX-o) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 37 | OrderID | Y | String ISO-8859-1 [18] | |
| 11 | ClOrdID | N | String ISO-8859-1 [50] | Pass-thru field set by client and echoed back by marketplace. |
| 41 | OrigClOrdID | N | String ISO-8859-1 [50] | Conditionally required for response to a Cancel or Cancel/Replace request. |
| | `<Parties>` group | N | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm, 53 = TraderMnemonic |
| 880 | TrdMatchID | N | String | Match ID assigned by the matching engine |
| 17 | ExecID | Y | String | An identifier assigned to each unique Execution Report message produced by the marketplace |
| 150 | ExecType | Y | char | Specifies a step in the order's life cycle trajectory. Supported values: 0 = New, 4 = Canceled, 5 = Replaced, 8 = Rejected, C = Expired, D = Restated, F = Trade, L = TriggeredOrActivatedBySystem |
| 39 | OrdStatus | Y | char | Describes the current state of a CHAIN of orders. Supported values: 0 = New, 1 = PartiallyFilled, 2 = Filled, 4 = Canceled, 8 = Rejected, C = Expired |
| 1328 | RejectText | N | String | Reason description for rejecting the transaction request |
| 378 | ExecRestatementReason | CT | int Reserved100Plus | Required for ExecType = D (Restated). Supported values: 1 = GTRenewal, 4 = BrokerOption, 8 = Market, 17 = CxldSMP, 100 = OrderRejectedWhenTriggered, 101 = CancelledDuetoCircuitBreaker, 102 = CancelledDuetoCorporateAction, 103 = CxldSMPGroupDefault |
| 1907 | NoRegulatoryTradeIDs | N | NumInGroup | Valid values: 1. Only populated for ExecType = F, H |
| → 1903 | RegulatoryTradeID | N | String | Trading Venue Transaction Identifier |
| → 1906 | RegulatoryTradeIDType | N | int | Supported values: 5 = TradingVenueTransactionIdentifier |
| 55 | Symbol | Y | String ISO-8859-1 [32] | Short name of instrument |
| 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| 38 | OrderQty | N | Qty | Quantity ordered |
| 40 | OrdType | N | char | Supported values: 1 = Market, 2 = Limit, K = MarketWithLeftOverAsLimit, P = Pegged |
| 44 | Price | N | Price | |
| 1100 | TriggerType | N | char | Not included for fills or cancellations |
| 1101 | TriggerAction | N | char | |
| 1102 | TriggerPrice | N | Price | |
| 1103 | TriggerSymbol | N | String ISO-8859-1 [32] | |
| 1107 | TriggerPriceType | N | char | |
| 1109 | TriggerPriceDirection | N | char | |
| 1113 | TriggerTradingSessionID | N | int Reserved100Plus | |
| 211 | PegOffsetValue | N | int | |
| 1094 | PegPriceType | N | int | |
| 835 | PegMoveType | N | int | |
| 836 | PegOffsetType | N | char | |
| 1096 | PegSecurityIDSource | N | char | |
| 1097 | PegSecurityID | N | String | |
| 1098 | PegSymbol | N | String | |
| 839 | PeggedPrice | N | Price | The current price the order is pegged at |
| 59 | TimeInForce | N | char | Supported values: 0 = Day, 3 = ImmediateOrCancel, 4 = FillOrKill |
| 432 | ExpireDate | N | LocalMktDate | |
| 18 | ExecInst | N | MultipleCharValue | |
| 528 | OrderCapacity | N | char | |
| 529 | OrderRestrictions | N | MultipleCharValue | |
| 2593 | NoOrderAttributes | N | NumInGroup | |
| → 2594 | OrderAttributeType | N | int Reserved1000Plus | |
| → 2595 | OrderAttributeValue | N | String | |
| 32 | LastQty | N | Qty | Quantity bought/sold on this (last) fill. Required if ExecType(150) = F (Trade). |
| 31 | LastPx | N | Price | Price of this (last) fill. Required if ExecType(150) = F (Trade). |
| 100 | ExDestination | N | Exchange | MIC code for requested venue of execution |
| 336 | TradingSessionID | N | String Reserved100Plus | |
| 625 | TradingSessionSubID | N | int Reserved100Plus | |
| 151 | LeavesQty | Y | Qty | Quantity open for further execution |
| 14 | CumQty | Y | Qty | Currently executed quantity for this order chain |
| 75 | TradeDate | N | LocalMktDate | Trade date according to the running trading session |
| 60 | TransactTime | N | UTCTimestamp | Time the transaction represented by this ExecutionReport occurred |
| 110 | MinQty | N | Qty | Min qty required to fill |
| 1822 | MinQtyMethod | N | int | |
| 2362 | SelfMatchPreventionID | N | String ISO-8859-1 [10] | As stated on the order, otherwise not present |
| 1138 | DisplayQty | N | Qty | The quantity to be displayed |
| 77 | PositionEffect | N | char | Supported values: C = Close, D = Default, O = Open |
| 58 | Text | N | String ISO-8859-1 [15] | Passthrough text from order if accepted |
| | `<StandardTrailer>` component | Y | | |

---

### 2.3 OrderCancelReplaceRequest (G, in)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | G (Added FIX 2.7) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Purpose:** The order cancel/replace request is used to change the parameters of an existing order. Do **not** use this message to cancel the remaining quantity of an outstanding order — use the Order Cancel Request message for this purpose.

**Usage and Conditions (Usage 1: Alter a live order).** The following attributes may be changed:
- DisplayQty (1138), OrdQty (38), Price (44), PartyRole (452): only CustomerAccount (24) can be updated, TIF (59), ExpiryDate (432), PositionEffect (77), Text (58), SelfMatchPreventionID (2362), ExecInst (18): must only be set to G if order is already AoN or FoK, MinQty (110), PegOffsetValue (211), OrderRestrictions (529): only adding, not removing

**Limitations:**
- Fields that are not allowed to change but are required must be re-populated with original values.
- A client cannot change OrdType.
- If 59 = FoK or 18 = G (AoN) then rejection ExecRpt is returned if MinQty is submitted.
- Omission of Price (44) is interpreted as no change to this tag.
- A change of quantity is submitted as the new full order quantity, not a delta quantity.
- Updating OrdQty to a value <= CumQty while CumQty > 0 leads to a filled order.
- Setting DisplayQty (1138) = 0 is interpreted as unaltered.
- SelfMatchPreventionID: To clear/unset use "0"; To set/change use only characters 0…9; To keep the value – omit the tag.

**Response/Acknowledgment:** Successful: ExecutionReport (8) / Not successful: OrderCancelReject (9)

**Table 6: OrderCancelReplaceRequest [G], in (OE-IDX-i) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 37 | OrderID | N | String ISO-8859-1 [18] | Unique (most recent) identifier of the order |
| | `<Parties>` group | Y (Req'd in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm, 53 = TraderMnemonic |
| 41 | OrigClOrdID | N | String ISO-8859-1 [50] | ClOrdID of the previous order. Use "NONE" if updating by OrderID (37). |
| 11 | ClOrdID | Y | String ISO-8859-1 [50] | |
| 18 | ExecInst | N | MultipleCharValue | Supported values: G = AllOrNone, i = ImbalanceOnly |
| 110 | MinQty | N | Qty | Min qty required to fill |
| 2362 | SelfMatchPreventionID | N | String ISO-8859-1 [10] | Must only contain characters 0..9 if used |
| 1138 | DisplayQty | N | Qty | The new quantity to be displayed |
| 55 | Symbol | Y (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N | String ISO-8859-1 [10] | Orderbook ID |
| 22 | SecurityIDSource | N | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 54 | Side | Y | char | Must match original order's side. Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| 60 | TransactTime | Y | UTCTimestamp | |
| 38 | OrderQty | Y | Qty | Quantity ordered (new total quantity) |
| 40 | OrdType | Y | char | Supported values: 1 = Market, 2 = Limit, K = MarketWithLeftOverAsLimit |
| 44 | Price | N | Price | New price if changed, omission is interpreted as no change |
| 211 | PegOffsetValue | N | int | Valid values: omitted = leave offset unchanged; 0 = sets offset to 0; [-127…127] |
| 835 | PegMoveType | N | int | Supported values: 0 = Floating, 1 = Fixed |
| 59 | TimeInForce | N | char | Supported values: 0 = Day, 3 = ImmediateOrCancel, 4 = FillOrKill |
| 432 | ExpireDate | N | LocalMktDate | Conditionally required if TimeInForce = GTD |
| 58 | Text | N | String ISO-8859-1 [15] | New Text if changed |
| 77 | PositionEffect | N | char | New Position effect if changed. Supported values: C = Close, D = Default, O = Open |
| | `<StandardTrailer>` component | Y | | |

---

### 2.4 OrderCancelRequest (F, in)

| Property | Value |
|----------|-------|
| **Direction** | In to Eclipse Trading |
| **Message type** | F (Added FIX 2.7) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Purpose:** The order cancel request message requests the cancellation of all of the remaining quantity of an existing order. Note that the Order Cancel/Replace Request should be used to partially cancel (reduce) an order.

**Table 7: OrderCancelRequest [F], in (base) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 41 | OrigClOrdID | Y (Req'd in JATS) | String ISO-8859-1 [50] | ClOrdID of the previous order. Use "NONE" if cancelling by OrderID (37). |
| 37 | OrderID | N | String ISO-8859-1 [18] | Unique (most recent) identifier of the order |
| 11 | ClOrdID | Y | String ISO-8859-1 [50] | Unique ID of cancel request (this message) as assigned by the institution |
| | `<Parties>` group | Y (Req'd in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2, 4 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 36 = EnteringTrader, 7 = EnteringFirm, 53 = TraderMnemonic |
| 55 | Symbol | Y (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N | String ISO-8859-1 [10] | Orderbook ID |
| 22 | SecurityIDSource | N | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| 60 | TransactTime | Y | UTCTimestamp | |
| | `<StandardTrailer>` component | Y | | |

---

### 2.5 OrderCancelReject (9, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | 9 (Added FIX.2.7) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Purpose:** The order cancel reject message is issued by the broker or exchange upon receipt of a cancel request or cancel/replace request message which cannot be honored.

**Table 8: OrderCancelReject [9], out (base) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 37 | OrderID | Y | String ISO-8859-1 [18] | If CxlRejReason="Unknown order", specify "NONE". |
| 11 | ClOrdID | Y | String ISO-8859-1 [50] | Unique order id assigned by institution to the cancel request or the replacement order |
| 41 | OrigClOrdID | N | String ISO-8859-1 [50] | ClOrdID(11) which could not be canceled/replaced |
| 39 | OrdStatus | Y | char | OrdStatus value after this cancel reject is applied. Supported values: 0 = New, 1 = PartiallyFilled, 2 = Filled, 4 = Canceled, 8 = Rejected, C = Expired |
| 60 | TransactTime | N | UTCTimestamp | |
| 434 | CxlRejResponseTo | Y | char | Supported values: 1 = OrderCancelRequest, 2 = OrderCancelReplaceRequest |
| 102 | CxlRejReason | N | int Reserved100Plus | Supported values: 0 = TooLateToCancel, 1 = UnknownOrder, 2 = BrokerCredit, 3 = OrderAlreadyInPendingStatus, 6 = DuplicateClOrdID, 99 = Other |
| 1328 | RejectText | N | String | Reason description for rejecting the transaction request |
| | `<StandardTrailer>` component | Y | | |

---

## 3. Continuous Quotation

JATS supports continuous quoting where market makers (quote contributors) provide quotes to the public orderbook. A general constraint is that a quote contributor can only have 1 quote per instrument (orderbook) and side; the combination of which implicitly defines a quote id.

Other important concepts that apply are:
- Freshened quotes are submitted and replace previous without canceling
- To update price only on one side of the mass quote, the price field of the opposite side is omitted
- Cancelled quotes are submitted by stating zero size
- MassQuote (i) is the single message used to provide public quotes
- MassQuoteAck (b) conveys rejects only, meaning no positive ack's is produced for solicited actions

QuoteID (117) is used for purposes and references the quote, both as message identifier and in Market Data, Executions (fills), etc.

### 3.1 MassQuote (i, in)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | i (Added FIX.4.2) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Usage and Conditions:**

MassQuote (in) is used for the following purposes:
- For quote contributor/market maker to send new or refresh/replace existing quotes.
- Quote can be one price level per instrument and side, or multi level quote with more than one price level per instrument and side.
- Quote contributor/market maker may cancel existing quotes (by sending qty/size=0)
- A MassQuote is implicitly a limit order valid for the day.

**Limitations:**
- MassQuote message can be used to enter/replace/cancel quotes on: one price level in one or several instruments, or multiple price levels in one instrument.
- Instrument, side, and quote contributor uniquely defines a quote entry.
- QuoteID is limited to 50 chars.
- QuoteID must be unique per mass quote message on the session.
- Number of Quote entries in a mass quote is limited to 10.
- OrderCapacity can only be set to same value for all quote entries.

**Response/Acknowledgment:** Successful: MassQuoteAck (b) or no acknowledgment depending on QuoteResponseLevel (301) / Not successful: MassQuoteAck (b)

**Table 9: MassQuote [i], in (OE-IDX-i) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 117 | QuoteID | Y | String ISO-8859-1 [50] | Set by respondent/quote submitter. Must be unique on the session. |
| 537 | QuoteType | Y (Req'd in JATS) | int | Type of Quote. Supported values: 1 = Tradeable |
| 2403 | QuoteModelType | N | int | Supported values: 1 = QuoteEntry |
| 301 | QuoteResponseLevel | N | int | Supported values: 1 = AcknowledgeOnlyNegativeOrErroneousQuotes, 2 = AcknowledgeEachQuoteMessage |
| | `<Parties>` group | Y (Req'd in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm |
| 296 | NoQuoteSets | N | NumInGroup | Valid values: 1 |
| → 302 | QuoteSetID | Y | String ISO-8859-1 [50] | Sequential number for the Quote Set |
| → 304 | TotNoQuoteEntries | Y | int | Total number of quotes for the quote set across all messages |
| → 893 | LastFragment | N | Boolean | Supported values: N = NotLastMessage, Y = LastMessage |
| → 295 | NoQuoteEntries | N | NumInGroup | Valid values: 1..10 |
| → → 299 | QuoteEntryID | Y | String ISO-8859-1 [50] | Uniquely identifies the quote across the complete set |
| → → 55 | Symbol | Y (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| → → 48 | SecurityID | N | String ISO-8859-1 [10] | Orderbook ID |
| → → 22 | SecurityIDSource | N | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| → → 132 | BidPx | N | Price | Not needed for cancellation; otherwise either BidPx, OfferPx or both must be specified |
| → → 133 | OfferPx | N | Price | Not needed for cancellation; otherwise either BidPx, OfferPx or both must be specified |
| → → 134 | BidSize | N | Qty | Enter 0 (zero) for cancellation, otherwise required if BidPx is given |
| → → 135 | OfferSize | N | Qty | Enter 0 (zero) for cancellation, otherwise required if BidPx is given |
| → → 528 | OrderCapacity | N | char | Supported values: A = Agency, G = Proprietary, I = Individual, P = Principal, R = RisklessPrincipal |
| | `<StandardTrailer>` component | Y | | |

---

### 3.2 MassQuoteAck (b, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | b (Added FIX.4.2, Updated FIX.5.0SP2 EP143) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Usage and Conditions:**  
A MassQuoteAck will be returned to the submitter of either of:
1. A MassQuote where at least 1 QuoteEntry wasn't accepted (which includes an implicit cancel with size=0).
2. A rejected QuoteCancel by Underlying.
3. An accepted MassQuote having QuoteResponseLevel (301) = 2.

**Limitations:** QuoteEntryRejectReason (368) for a quote entry returned in a non-successful MassQuoteAck is =99 or any of the custom values 201..210. The specific error is always described in the Text (58) field.

**Table 10: MassQuoteAck [b], out (base) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 117 | QuoteID | N | String ISO-8859-1 [50] | Required when acknowledgment is in response to a Mass Quote, mass Quote Cancel or mass Quote Status Request message |
| 297 | QuoteStatus | Y | int | Status of the MassQuote. Supported values: 0 = Accepted, 5 = Rejected |
| 300 | QuoteRejectReason | N | int Reserved100Plus | Supported values: 105 = GatewayRejectOnInvalidQuoteEntry, 4 = TooLateToEnter, 9 = NotAuthorizedToQuoteSecurity, 99 = Other |
| 301 | QuoteResponseLevel | N | int | |
| 537 | QuoteType | Y (Req'd in JATS) | int | Supported values: 1 = Tradeable |
| | `<Parties>` group | Y | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..5 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm |
| 58 | Text | N | String ISO-8859-1 [80] | Additional rejection description |
| 296 | NoQuoteSets | N | NumInGroup | Valid values: 1 |
| → 302 | QuoteSetID | N | String ISO-8859-1 [50] | |
| → 304 | TotNoQuoteEntries | N | int | |
| → 1170 | TotNoRejQuotes | N | int | Total number of quotes rejected for the quote set |
| → 893 | LastFragment | N | Boolean | |
| → 295 | NoQuoteEntries | N | NumInGroup | Valid values: 1..10 |
| → → 299 | QuoteEntryID | N | String ISO-8859-1 [50] | |
| → → 1167 | QuoteEntryStatus | N | int | Supported values: 5 = Rejected |
| → → 368 | QuoteEntryRejectReason | N | int Reserved100Plus | Supported values: 4 = TooLateToEnter, 8 = InvalidPrice, 9 = NotAuthorizedToQuoteSecurity, 201-210 = QuoteEntry1-10 |
| | `<StandardTrailer>` component | Y | | |

---

## 4. Indicative Quotation

### 4.1 Quote (S, in, indicative)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | S (Added FIX.4.0) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Limitations:**
- Only 1-sided quotes are supported
- QuoteID is limited to 50 chars
- On-behalf of Quotes are not supported
- Quote update is not supported. Any change must be done through cancelling and entering a new quote

**Response/Acknowledgment:** Successful: no immediate acknowledgment / Not successful: QuoteStatusReport (AI)

**Table 11: Quote [S], in (OE-IDX-i-IQ) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = S |
| 117 | QuoteID | Y | String ISO-8859-1 [50] | |
| 537 | QuoteType | Y (Req'd in JATS) | int | Supported values: 0 = Indicative |
| 1171 | PrivateQuote | Y (Req's in JATS) | Boolean | Supported values: N = PublicQuote |
| 2837 | SingleQuoteIndicator | Y (Req's in JATS) | char | Supported values: N = False/No |
| | `<Parties>` group | Y (Req's in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..3 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char ISO-8859-1 1 | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | PartyRole 1 and 12 are required. Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount |
| 55 | Symbol | Y (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 1193 | SettlMethod | N | String | Supported values: 1 = DvP, 2 = DFoP |
| 54 | Side | Y (Req's in JATS) | char | Supported values: 1 = Buy, 2 = Sell |
| 38 | OrderQty | Y | Qty | Quantity ordered |
| 132 | BidPx | N | Price | |
| 133 | OfferPx | N | Price | |
| 528 | OrderCapacity | N | char | Supported values: A = Agency, G = Proprietary, I = Individual, P = Principal, R = RisklessPrincipal |
| 58 | Text | N | String ISO-8859-1 [15] | Free passthrough text |
| | `<StandardTrailer>` component | Y | | |

---

### 4.2 Quote (S, out, indicative)

A submitted indicative quote is public market data and conveyed through JATS FIX market data feed. Details can be found in JATS FIX specification for Market Data.

---

### 4.3 QuoteResponse (AJ, in, indicative)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | AJ (Added FIX.4.4) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Limitations:**
- QuoteRespID is limited to 50 chars
- On-behalf of quote responses are not supported
- Quote offer update is not supported

**Table 12: QuoteResponse [AJ], in (OE-IDX-i-IQ) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = AJ |
| 693 | QuoteRespID | Y | String ISO-8859-1 [50] | ID as assigned by the respondent |
| 117 | QuoteID | N | String ISO-8859-1 [50] | |
| 694 | QuoteRespType | Y | int | Supported values: 1 = Hit (Hit/Lift), 7 = EndTrade (End/Cancel Quote offer) |
| 528 | OrderCapacity | N | char | |
| | `<Parties>` group | Y (Req's in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..3 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount |
| 55 | Symbol | Y (Req's in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N | String Reserved100Plus | |
| 1193 | SettlMethod | Y | String | Supported values: 1 = DvP, 2 = DFoP |
| 38 | OrderQty | Y | Qty | Required when QuoteRespType = 1 |
| 54 | Side | Y (Req's in JATS) | char | Supported values: 1 = Buy, 2 = Sell |
| 40 | OrdType | Y (Req's in JATS) | char | Supported values: D = PreviouslyQuoted |
| 58 | Text | N | String ISO-8859-1 [15] | Free passthrough text |
| 44 | Price | N | Price | |
| | `<StandardTrailer>` component | Y | | |

---

### 4.4 QuoteResponse (AJ, out, indicative)

**Table 13: QuoteResponse [AJ], out (OE-C-o-IQ) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = AJ |
| 693 | QuoteRespID | Y | String ISO-8859-1 [50] | Set to "NONE" by system |
| 117 | QuoteID | Y (Req's in JATS) | String ISO-8859-1 [50] | The quote provider's quote ID |
| 694 | QuoteRespType | Y | int | Supported values: 1 = Hit, 7 = EndTrade |
| | `<Parties>` group | Y (Req's in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 1 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | Message related actor ID (Firm) |
| → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm |
| 55 | Symbol | Y (Req's in JATS) | String ISO-8859-1 [32] | |
| 48 | SecurityID | N | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N | String Reserved100Plus | |
| 54 | Side | Y (Req's in JATS) | char | Same side as the quote. Supported values: 1 = Buy, 2 = Sell |
| 38 | OrderQty | Y (Req's in JATS) | Qty | The quantity offered to trade |
| 62 | ValidUntilTime | Y (Req's in JATS) | UTCTimestamp | Time when offer expires in UTC |
| 60 | TransactTime | N | UTCTimestamp | Creation time of offer in UTC |
| 40 | OrdType | Y (Req's in JATS) | char | Supported values: D = PreviouslyQuoted |
| 44 | Price | Y (Req's in JATS) | Price | |
| | `<StandardTrailer>` component | Y | | |

---

### 4.5 QuoteCancel (Z, in, indicative)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | Z (Added FIX.4.2) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Limitations:**
- To cancel a single private quote the QuoteID must be provided.
- On-behalf of quote cancellations are not supported.

**Table 14: QuoteCancel [Z], in (OE-C-i-IQ) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = Z |
| 117 | QuoteID | Y (Req's in JATS) | String ISO-8859-1 [50] | QuoteID for the quote to be cancelled |
| 298 | QuoteCancelType | Y | int Reserved100Plus | Supported values: 5 = CancelSpecifiedSingleQuote |
| 537 | QuoteType | Y (Req's in JATS) | int | Supported values: 0 = Indicative |
| | `<Parties>` group | Y (Req's in JATS) | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader |
| 295 | NoQuoteEntries | N | NumInGroup | Valid values: 1 |
| → 55 | Symbol | Y | String ISO-8859-1 [32] | Short name of instrument |
| | `<StandardTrailer>` component | Y | | |

---

### 4.6 QuoteStatusReport (AI, out, indicative, reject)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | AI (Added FIX.4.3) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Usage and Conditions:**
1. To QPro: Used by intermediary to reject an invalid Quote from quote provider
2. To QPro: Used by intermediary to reject an invalid QuoteCancel from quote provider
3. To QPro: Used by intermediary to reject an invalid QuoteStatusReport from quote provider
4. To QOff: Used by intermediary to reject an invalid QuoteResponse from quote offerer

**Table — QuoteStatusReport [AI], out (OE-C-o-IQRej) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = AI |
| 117 | QuoteID | N | String ISO-8859-1 [50] | QuoteID that was sent in |
| 693 | QuoteRespID | N | String ISO-8859-1 [50] | QuoteResponseID that was sent in |
| 537 | QuoteType | N | int | Supported values: 0 = Indicative |
| | `<Parties>` group | N | | |
| 453 | NoPartyIDs | N | NumInGroup | Valid values: 2 |
| → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader |
| 55 | Symbol | N | String ISO-8859-1 [32] | |
| 48 | SecurityID | N | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N | String Reserved100Plus | |
| 54 | Side | N | char | Supported values: 1 = Buy, 2 = Sell |
| 38 | OrderQty | N | Qty | |
| 44 | Price | N | Price | |
| 297 | QuoteStatus | Y (Req'd in ET) | int | Supported values: 5 = Rejected |
| 300 | QuoteRejectReason | Y (Req'd in ET) | int Reserved100Plus | Supported values: 1 = UnknownSymbol, 4 = TooLateToEnter, 5 = UnknownQuote, 6 = DuplicateQuote, 8 = InvalidPrice, 9 = NotAuthorizedToQuoteSecurity, 99 = Other, 199 = InvalidQuantity, 202 = InvalidQuoteRespType, 203 = DuplicateQuoteRespID |
| 1328 | RejectText | N | String | Additional rejection description |
| | `<StandardTrailer>` component | Y | | |

---

## 5. Trade Matching

JATS provides a trade report matching facility. A participant user can send in an agreed or proposed trade to a counterparty in the form of a TradeCaptureReport (TCR), dispatched by the system. The appointed counterparty can either ignore the trade or accept it. An acceptance is communicated by sending a reciprocal trade.

The trade reports are subject to a matching attempt by the system and if the match is successful, trade confirmations are sent out on the post-trade flow. Any trade report sent in that doesn't match will instead result in a new trade proposal to the counterparty.

Any unattended trades for matching are cancelled (deleted) at the end of the trading session.

### 5.1 TradeCaptureReport (AE, in, 1-sided)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | AE (Added FIX.4.3) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Usage and Conditions:**
1. Propose a trade to a counterparty
2. Respond to a trade proposal
3. Cancel an unmatched trade proposal

> `[M]` is used to indicate which fields are matched to the counterparty's Trade report.

**Limitations:** The matched trade is not returned on the OE session. It is available on the Drop Copy (DC) feed.

**Table 15: TradeCaptureReport [AE], in (OE-IDX-i-1side) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 571 | TradeReportID | Y (Req'd in JATS) | String ISO-8859-1 [50] | |
| 1003 | TradeID | N | String ISO-8859-1 [39] | |
| 1040 | SecondaryTradeID | N (Cond'l in JATS) | String ISO-8859-1 [39] | ID that JATS prints on an acknowledgment for a successful reception. Needed if the participant wants to do a cancellation. |
| 856 | TradeReportType | Y (Req'd in JATS) | int | Supported values: 0 = Submit, 6 = TradeReportCancel |
| 1123 | TradeHandlingInstr | Y (Req'd in JATS) | char | Supported values: 1 = TwoPartyReportForMatching, 2 = OnePartyReportForMatching |
| 55 | Symbol | Y [M] (Req's in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N [M] | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N [M] | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| 167 | SecurityType | N | String ISO-8859-1 [32] | Not validated or used in JATS. Supported values: NONE = NoSecurityType, OTHER = Other |
| 762 | SecuritySubType | Y [M] (Req's in JATS) | String | Set to off-book type |
| 1193 | SettlMethod | Y | String | Supported values: 1 = DvP, 2 = DFoP |
| 32 | LastQty | Y [M] (Req's in JATS) | Qty | Trade quantity |
| 31 | LastPx | Y [M] (Req'd in JATS) | Price | Trade price |
| 60 | TransactTime | N | UTCTimestamp | |
| 64 | SettlDate | N | LocalMktDate | Settlement date as agreed between parties |
| 552 | NoSides | N | NumInGroup | Valid values: 1 |
| → 54 | Side | Y [M] | char | Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| → 453 | NoPartyIDs | N | NumInGroup | Valid values: 3..6 |
| → → 448 | PartyID | Y [M] | String ISO-8859-1 [32] | |
| → → 447 | PartyIDSource | Y [M] | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount, 36 = EnteringTrader, 7 = EnteringFirm, 17 = ContraFirm |
| → 77 | PositionEffect | N | char | Supported values: C = Close, D = Default, O = Open |
| → 58 | Text | N | String ISO-8859-1 [15] | |
| → 528 | OrderCapacity | N | char | |
| → 483 | TransBkdTime | N [M*] (Cond'l in JATS) | UTCTimestamp | Time of agreement |
| | `<StandardTrailer>` component | Y | | |

---

### 5.2 TradeCaptureReportAck (AR, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | AR (Added FIX.4.4) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Table 16: TradeCaptureReportAck [AR], out (OE-C-o-1side) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 571 | TradeReportID | N (Cond'l in JATS) | String ISO-8859-1 [50] | |
| 1003 | TradeID | N (Cond'l in ET) | String ISO-8859-1 [39] | Set upon successful validation, serves as a preliminary ID used to reference the trade while not yet matched |
| 1040 | SecondaryTradeID | N (Cond'l in ET) | String ISO-8859-1 [39] | If accepted, the (preliminary) identifier created by JATS upon successful reception |
| 856 | TradeReportType | Y (Req's in JATS) | int | Supported values: 0 = Submit, 6 = TradeReportCancel |
| 1123 | TradeHandlingInstr | N (Cond'l in ET) | char | Supported values: 1 = TwoPartyReportForMatching, 2 = OnePartyReportForMatching |
| 150 | ExecType | N (Cond'l in ET) | char | Supported values: F = Trade |
| 939 | TrdRptStatus | N (Req's in JATS) | int | Supported values: 0 = Accepted, 1 = Rejected, 2 = Cancelled |
| 751 | TradeReportRejectReason | N | int Reserved100Plus | Supported values: 1 = InvalidPartyInformation, 149 = DealCancelledByIntermediary, 2 = UnknownInstrument, 3 = UnauthorizedToReportTrades, 4 = InvalidTradeType, 99 = Other |
| 1328 | RejectText | N | String | |
| 55 | Symbol | Y (Req's in JATS) [M] | String ISO-8859-1 [32] | Short name of instrument |
| 48 | SecurityID | N [M] | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N [M] | String Reserved100Plus | |
| 167 | SecurityType | N | String ISO-8859-1 [32] | |
| 762 | SecuritySubType | N (Not req'd on rejection) [M] | String | |
| 573 | MatchStatus | N (Not req'd on rejection) | char | Supported values: 0 = Compared, 1 = Uncompared |
| 552 | NoSides | N | NumInGroup | Valid values: 1 |
| → 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell |
| | `<StandardTrailer>` component | Y | | |

---

### 5.3 TradeCaptureReport (AE, out, 1-sided, to counterparty)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | AE (Added FIX.4.3) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Usage and Conditions:**
1. To appointed counterparties to get trade notifications (proposals) to respond to
2. To appointed counterparties to inform about cancelled trade proposals from the initiator
3. To appointed counterparties to inform about cancelled trade proposals due to the trade reports being matched

**Table 17: TradeCaptureReport [AE], out (OE-IDX-o-1side) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 571 | TradeReportID | Y | String ISO-8859-1 [50] | |
| 1003 | TradeID | Y (Req's in JATS) | String ISO-8859-1 [39] | Preliminary ID used to reference the trade while not yet matched |
| 1040 | SecondaryTradeID | Y (Req's in JATS) | String ISO-8859-1 [39] | ID usually but not necessarily equal to TradeID (1003). This will survive to the final TCR on the same side. |
| 856 | TradeReportType | Y (Req's in JATS) | int | Supported values: 11 = AllegedNew, 14 = AllegedTradeReportCancel |
| 1123 | TradeHandlingInstr | N | char | Supported values: 2 = OnePartyReportForMatching |
| 55 | Symbol | Y [M] (Req's in JATS) | String ISO-8859-1 [32] | |
| 48 | SecurityID | N [M] | String ISO-8859-1 [10] | |
| 22 | SecurityIDSource | N [M] | String Reserved100Plus | |
| 167 | SecurityType | N | String ISO-8859-1 [32] | |
| 762 | SecuritySubType | N (Not req'd on rejection) [M] | String | |
| 1193 | SettlMethod | Y | String | Supported values: 1 = DvP, 2 = DFoP |
| 32 | LastQty | Y [M] (Req's in JATS) | Qty | Trade quantity |
| 31 | LastPx | Y [M] (Req's in JATS) | Price | |
| 64 | SettlDate | N | LocalMktDate | |
| 573 | MatchStatus | N | char | Supported values: 0 = Compared, 1 = Uncompared |
| 552 | NoSides | N | NumInGroup | Valid values: 1 |
| → 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell |
| → 453 | NoPartyIDs | N | NumInGroup | Valid values: 3 |
| → → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 17 = ContraFirm |
| → 483 | TransBkdTime | N (Cond'l in ET) | UTCTimestamp | |
| | `<StandardTrailer>` component | Y | | |

---

## 6. Deal Reporting

Deal registration is the action of sending trades done off-exchange to the JATS in order to get formal or official deal and trade ID:s from the marketplace and to fulfill regulatory requirements of trade registration and publication.

The system supports double-sided (one of the trade parties reports both sides) binary (1 buyer and 1 seller) trade reports, potentially containing a number of trades (multi-leg) in different instruments.

### 6.1 TradeCaptureReport (AE, in, 2-sided)

| Property | Value |
|----------|-------|
| **Direction** | In to JATS |
| **Message type** | AE (Added FIX.4.3) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Limitations:**
- TradeCaptureReport (in) is 2-sided multileg only.
- The number of leg trades is in the range 1..6.
- OrderCapacity (528) can only be set per trade report side, not per trade leg.

**Response/Acknowledgment:** TradeCaptureReportAck (AR), out (in either case: accepted or rejected)

**Table 18: TradeCaptureReport [AE], in (OE-IDX-i-2side) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 571 | TradeReportID | Y (Req'd in JATS) | String ISO-8859-1 [50] | Client set message identifier |
| 856 | TradeReportType | Y (Req'd in JATS) | int | Supported values: 0 = Submit |
| 1123 | TradeHandlingInstr | Y | char | Supported values: 1 = TwoPartyReport |
| 55 | Symbol | N (Req'd in JATS) | String ISO-8859-1 [32] | Set to "[N/A]" |
| 167 | SecurityType | N | String ISO-8859-1 [32] | Set to "MLEG" |
| 762 | SecuritySubType | Y (Req'd in JATS) | String | Set to off-book type |
| 1193 | SettlMethod | Y | String | Supported values: 1 = DvP, 2 = DFoP |
| 555 | NoLegs | N | NumInGroup | Valid values: 1..6 |
| → 600 | LegSymbol | N (Req'd in JATS) | String ISO-8859-1 [32] | Short name of instrument |
| → 602 | LegSecurityID | N | String ISO-8859-1 [10] | |
| → 603 | LegSecurityIDSource | N | String Reserved100Plus | Supported values: M = MARKETPLACE_ASSIGNED_IDENTIFIER |
| → 624 | LegSide | Y (Req'd in JATS) | char | Supported values: B = AsDefined, C = Opposite |
| → 637 | LegLastPx | Y (Req'd in JATS) | Price | |
| → 1418 | LegLastQty | Y (Req'd in JATS) | Qty | |
| 60 | TransactTime | N | UTCTimestamp | |
| 64 | SettlDate | N | LocalMktDate | Settlement date (expected) |
| 552 | NoSides | N | NumInGroup | Valid values: 2 |
| → 54 | Side | Y | char | Supported values: 1 = Buy, 2 = Sell, 5 = SellShort |
| → 453 | NoPartyIDs | N | NumInGroup | Valid values: 2..3 |
| → → 448 | PartyID | Y | String ISO-8859-1 [32] | |
| → → 447 | PartyIDSource | Y | char | Supported values: D = PROPRIETARY_CUSTOM_CODE |
| → → 452 | PartyRole | Y | int | Supported values: 1 = ExecutingFirm, 12 = ExecutingTrader, 24 = CustomerAccount |
| → 58 | Text | N | String ISO-8859-1 [15] | |
| → 528 | OrderCapacity | N | char | |
| → 483 | TransBkdTime | N | UTCTimestamp | Time of agreement. Can only be set on the reporting session (party) side. |
| | `<StandardTrailer>` component | Y | | |

---

### 6.2 TradeCaptureReportAck (AR, out, 2-sided)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | AR (Added FIX.4.4) |
| **FIX Session** | Order Entry (OE) |
| **Available to** | Participants |

**Table 19: TradeCaptureReportAck [AR], out (OE-C-o-2side) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | |
| 571 | TradeReportID | N | String ISO-8859-1 [50] | |
| 1003 | TradeID | N | String ISO-8859-1 [39] | Set upon successful validation. When matched, a new ID is assigned. |
| 1040 | SecondaryTradeID | N | String ISO-8859-1 [39] | If accepted, the (preliminary) identifier created by JATS upon successful reception |
| 856 | TradeReportType | Y (Req'd in JATS) | int | Supported values: 0 = Submit |
| 1123 | TradeHandlingInstr | N (Cond'l in JATS) | char | Supported values: 1 = TwoPartyReport |
| 150 | ExecType | N (Cond'l in JATS) | char | Supported values: F = Trade |
| 939 | TrdRptStatus | N (Req'd in JATS) | int | Supported values: 0 = Accepted, 1 = Rejected, 2 = Cancelled |
| 751 | TradeReportRejectReason | N | int Reserved100Plus | Supported values: 1 = InvalidPartyInformation, 149 = DealCancelledByIntermediary, 2 = UnknownInstrument, 3 = UnauthorizedToReportTrades, 4 = InvalidTradeType, 99 = Other |
| 1328 | RejectText | N | String | |
| 55 | Symbol | N | String ISO-8859-1 [32] | Set to "[N/A]" |
| 167 | SecurityType | N (Not req'd on rejection) | String | Set to "MLEG" |
| 762 | SecuritySubType | N (Not req'd on rejection) | String | |
| 573 | MatchStatus | N (Not req'd on rejection) | char | Supported values: 0 = Compared, 1 = Uncompared |
| | `<StandardTrailer>` component | Y | | |

---

## 7. Rejections

### 7.1. BusinessMessageReject (j, out)

| Property | Value |
|----------|-------|
| **Direction** | Out from JATS |
| **Message type** | j (Added FIX.4.2) |
| **FIX Session** | Market Data (MD) |
| **Available to** | Participants |

**Purpose:** The Business Message Reject message can reject an application-level message which fulfills session-level rules and cannot be rejected via any other means.

**Usage and Conditions:** The Business Message Reject message is used when JATS cannot honor the following MD requests: SecurityMassStatusRequest (CN, in).

**Table 20: BusinessMessageReject [j], out (ET-C-o) — Message Structure**

| FIX tag no | FIX tag name | Req'd | Type | Description |
|------------|-------------|-------|------|-------------|
| | `<StandardHeader>` component | Y | | MsgType = j |
| 45 | RefSeqNum | N | SeqNum | MsgSeqNum of rejected message |
| 372 | RefMsgType | Y | String | The MsgType of the FIX message being referenced |
| 380 | BusinessRejectReason | Y | int | Supported values: 0 = OTHER, 1 = UNKNOWN_ID, 5 = CONDITIONALLY_REQUIRED_FIELD_MISSING, 6 = NOT_AUTHORIZED |
| 58 | Text | N | String ISO-8859-1 [15] | Where possible, message to explain reason for rejection |
| | `<StandardTrailer>` component | Y | | |

---

## 8. Appendices

### 8.1. How to Update or Cancel an Order in Different Scenarios

Assume there's an order that's been assigned the following values at entry:
- OrderID (37) = 'Order1'
- ClOrdID (11) = 'apple'

The following 4 possible modification scenarios can occur depending on how the order was submitted and who is acting on it:

| Scenario | Parties to enter | Order reference to use | ER response |
|----------|-----------------|----------------------|-------------|
| ExecutingTrader A01 with ExecutingFirm A modifies own order | ExecutingFirm (1) = A, ExecutingTrader (12) = A01 | ClOrdID (11) = 'apple2', OrigClOrdID (41) = 'apple' | ClordID (11) = 'apple2', OrigClOrdID = 'apple', OrderID = 'Order1' |
| ExecutingTrader A01 with ExecutingFirm A modifies own order that was entered on behalf by other EnteringTrader | ExecutingFirm (1) = A, ExecutingTrader (12) = A01 | OrderID = 'Order1', ClOrdID (11) = 'berry' (!), OrigClOrdID (41) = 'NONE' | ClordID (11) = 'berry', OrigClOrdID = 'NONE', OrderID = 'Order1' |
| EnteringTrader T01 with EnteringFirm T (!!) modifies on-behalf entered order for ExecutingFirm A | ExecutingFirm (1) = A, ExecutingTrader (12) = A01 (!!!), EnteringFirm (7) = T, EnteringTrader (36) = T01 | ClOrdID (11) = 'apple2', OrigClOrdID (41) = 'apple' | ClordID (11) = 'apple2', OrigClOrdID = 'apple', OrderID = 'Order1' |
| Trader Z01 not part of order submission modifies order owned by ExecutingTrader A01 (!!!) with ExecutingFirm A | ExecutingFirm (1) = A, ExecutingTrader (12) = A01, TraderMnemonic (53) = Z01 | OrderID = 'Order1' (!!!!), ClOrdID (11) = 'cherry' (!), OrigClOrdID (41) = 'NONE' | ClordID (11) = 'cherry', OrigClOrdID = 'NONE', OrderID = 'Order1' |

**Legend:**
- **!** The modifying actor's ClOrdID does not interfere or impact the order submitter's order access using the ClOrdID returned upon submission or subsequent modification by the submitting session.
- **!!** If Entering and Executing traders are with same firm, same firm occurs twice in different roles.
- **!!!** ExecutingTrader does not have to be a FIX actor.
- **!!!!** Order does not have to be a FIX submitted order.

---

### 8.2. Order State Change Matrices

Please refer to the below location for a comprehensive specification of order state changes:  
https://www.fixtrading.org/online-specification/order-state-changes/

---

### 8.3. Trigger Component Usages

**Stop Order Example**

**Table 21: Submit Standard Stop Loss Order Using Trigger Properties**

| Tag | Field Name | Value | Value Description | Comment |
|-----|-----------|-------|-------------------|---------|
| 1100 | TriggerType | 4 | Price Movement | |
| 1101 | TriggerAction | 1 | Activate | |
| 1102 | TriggerPrice | 10.00 | | |
| 1107 | TriggerPriceType | 2 | Last Trade | |
| 1109 | TriggerPriceDirection | D | Down | Will trigger when price goes below 10.00 or immediately upon entry if price is already below 10.00 |
| 40 | OrdType | 2 | Limit | |
| 44 | Price | 9.50 | | |

---

### 8.4. One-Sided Trade Report for Matching Flow

**Flow description (Party A enters a TCR for matching with Party B):**

1. **Participant A** submits a `TradeCaptureReport` — enter new trade (buy) with:
   - TradeReportID = '<RPT1>'
   - TradeHandlingInstr = 2 (One-party report for matching)
   - TradeReportType = 0 (Submit)
   - NoSides = 1; Side = 1 (buy)
   - NoPartyDetails = 3..6 (ExecutingFirm, ExecutingTrader, CustomerAccount, ContraFirm, EnteringFirm, EnteringTrader)

2. **Marketplace** sends `TradeCaptureReport` — notification to counterparty (Party B):
   - TradeReportType = 11 (AllegedNew)
   - MatchStatus = 1 (Uncompared, unmatched, or unaffirmed)
   - Side = 1 (buy)

3. **Marketplace** sends `TradeCaptureReportAck` — success response to Party A:
   - TrdRptStatus = 0 (Accepted)
   - MatchStatus = 1 (Uncompared, unmatched, or unaffirmed)

4. **Participant B** submits matching `TradeCaptureReport` (sell) to match the trade.

5. **Marketplace**: The 2 trades are matched into a deal. Any notifications sent are cancelled. Matched trades are published on PT/DC/CL flows.

6. **Marketplace** sends `TradeCaptureReport` — cancel notification to counterparty (Party B):
   - TradeReportType = 14 (AllegedTradeReportCancel)
   - MatchStatus = 1 (Uncompared, unmatched, or unaffirmed)

7. **Marketplace** sends `TradeCaptureReport` — confirmations to parties on PT/DC flow.

---

**Flow description (Party A cancels its 1-sided trade report):**

1. **Participant A** submits `TradeCaptureReport` — cancel unmatched trade report:
   - TradeHandlingInstr = 2 (One-party report for matching)
   - TradeReportType = 6 (Trade Report Cancel)

2. **Marketplace** sends notification `TradeCaptureReport` to counterparty (Party B):
   - TradeReportType = 14 (Alleged Trade Report Cancel)
   - MatchStatus = 1 (Uncompared, unmatched, or unaffirmed)

3. **Marketplace** sends `TradeCaptureReportAck` — success response to Party A:
   - TrdRptStatus = 0 (Accepted)
   - MatchStatus = 1

> The trade report with the preliminary designated TradeID is no longer available.

---

*© Indonesia Stock Exchange | Private and Confidential | Version 2.0 | 14 Oct 2025*
