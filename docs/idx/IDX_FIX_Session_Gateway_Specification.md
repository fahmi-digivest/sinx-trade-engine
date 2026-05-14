# IDX FIX Protocol Session and Gateway Specification

**Document Version:** 1.0  
**Date:** 05 March 2025  
**Issuer:** Indonesia Stock Exchange (Bursa Efek Indonesia)  
**Classification:** Private and Confidential

> **Disclaimer:** The information in this document may change in accordance with the progress of the ongoing project development.

---

## Summary of Changes

| Version | Date | List of Changes |
|---------|------|-----------------|
| 1.0 | 05/03/2025 | First release |

---

## 1. About the Document

### 1.1 Specification Status and Refinement

This specification depicts the NFF (Nasdaq Financial Framework) FIX service session layer for the MME (Multi Matching Engine) trading product following the **FIX 5.0SP2** and **FIX 4.4** protocol standards.

The full FIX Rules of Engagement suite of specifications includes:
- Session layer
- Reference Data
- Order and Trade Entry
- Post-trade confirmations
- Market Data

A condensed and consistent set of messages and tags is included to meet demands imposed by authorities and regulation regimes under which financial bodies operate at each individual location.

---

### 1.2 Concepts and Legend

#### 1.2.1 Characteristics Given for a Message

Each message is documented with a leading table summarizing the characteristics and use cases of the message.

| Field | Description |
|-------|-------------|
| **Direction** | "In to marketplace" or "out from marketplace". If both variants exist, separate sections are used. |
| **Message code and pedigree** | The FIX message type: MsgType (tag 35), located in the header. The FIX version the message was introduced, and when it was updated. |
| **Purpose by FIX standard** | The general purpose given by the FIX standard. Often versatile and usually narrows itself at a specific installation. Given for reference only. |
| **FIX Session** | The session over which the message should be communicated (e.g. OE – Order Entry session). |
| **Available to** | The intended category of business actors at the intended Marketplace installation. |
| **Usage and Conditions** | Information on specific use cases and any specific conditions that apply. |
| **Limitations** | Specification of limitations that apply or need to be considered when designing operational processes. |
| **Response/Acknowledgment** | The expected response from the peer, which can differ for successful or unsuccessful outcomes. |

#### 1.2.2 Message Chaining or Entity Chaining

FIX offers two principles of chaining a sequence of orders, trades, and other successive events:

- **Message chaining model** – treats the event and its related message as "the same thing"; a successor message refers to its preceding message in the chain.
- **Entity chaining model** – references the entity the event created (e.g. a trade), which is treated as something that exists in its own right, not requiring a message proxy.

This implementation adheres to the **entity chaining** principle where not required to do otherwise. For example, when cancelling a trade, `TradeID (1003)` is more important than `TradeReportID (571)`. This ensures any entity (such as a trade or order) can be addressed unambiguously without referring to a specific protocol.

> **Note:** There are situations where a TradeID is not yet established (e.g. when reporting an off-exchange trade). In such cases, a reference to `TradeReportID` (if submitted via FIX) is appropriate until the Marketplace creates the TradeID.

#### 1.2.3 About the Req'd (Required) Column

| Symbol | Meaning |
|--------|---------|
| `Y` | Required by FIX standard |
| `Q` | Required by NDAQ standard |
| `C` | Conditionally required depending on use case |
| `O` | Optional (default if not declared) |

> The first tag in a repeating group is always required if the group is used. For the Parties group, all 3 tags (`PartyID`, `PartyIDSource`, `PartyRole`) must be supplied.

#### 1.2.4 About the UD (Union Datatype) Column

The FIX standard allows free-use enumeration ranges for some tags. Custom values can be added (agreed between all parties) without breaking the protocol.

- `R100+`, `R1000+`, `R4000+` — extendable value sets
- `BA+` — standard value set extended by bilateral agreement
- `CL` — values defined in a codelist maintained by the FIX consortium

#### 1.2.5 About the Parties Group (453)

The Parties group is common to almost all business level messages and is vital to messages directed towards trading (orders, quotes, trade reports, etc.). It differs from CompID:s in the header in that it defines **business-level stakeholders**, whereas header CompIDs are pure technical actors.

A party definition is composed of 3 particles: **ID**, **IDSource**, **Role**. All 3 must be provided with valid values. Mandatory parties are designated with `Y` in the message specifications; this varies by message.

#### 1.2.6 About Price and Quantity

Price and Quantity precision is determined by the number of decimals fetched from the `SecurityList` message as Instrument price precision and Instrument quantity precision, entered with a decimal point.

**Examples:**
- Tag 44 (Price): `1.012345`
- Tag 38 (OrderQty): `2.03`

---

### 1.3 FIX Infrastructure and Session Handling

#### 1.3.1 FIX Version

The FIX version of this specification is **FIX 5.0SP2**.  
Full documentation: https://www.fixtrading.org/online-specification/

#### 1.3.2 Standard Header

All FIX messages contain a Standard Header carrying session identifiers (CompIDs), sequence numbers, message type and length, and other metadata.

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| 8 | BeginString | Y | Identifies beginning of new message and protocol version. Always first field. Valid values: `FIXT.1.1`, `FIX.4.4` |
| 9 | BodyLength | Y | Message length in bytes, forward to the CheckSum field. Always second field. |
| 35 | MsgType | Y | Defines message type. Always third field. |
| 49 | SenderCompID | Y | The core MME gateways use `"MME"` |
| 56 | TargetCompID | Y | As specified in separate agreement with the customer |
| 34 | MsgSeqNum | Y | Integer message sequence number |
| 50 | SenderSubID | Q | As specified in separate agreement with the customer |
| 142 | SenderLocationID | — | Sender's LocationID. Not used in MME. |
| 57 | TargetSubID | — | Assigned value to identify specific individual or unit. "ADMIN" reserved for administrative messages. Not used in MME. |
| 143 | TargetLocationID | — | Trading partner LocationID. Not used in MME. |
| 43 | PossDupFlag | — | Indicates possible retransmission. Always required for retransmitted messages. |
| 97 | PossResend | — | Indicates message may contain information sent under another sequence number. Required when message may be a duplicate. |
| 52 | SendingTime | Y | Time of message transmission (UTC). |
| 122 | OrigSendingTime | — | Original transmission time (UTC). Required for message resent as a result of a ResendRequest. |

#### 1.3.3 Session and Gateways

For different service categories, different gateways are used. Service segregation is as follows:

| FIX GWY | Description |
|---------|-------------|
| **OE** (order entry) | Order and quotation messaging, and reporting of off-exchange trades. Individual content messages to recipients. The OE gateway can be configured to provide Reference data. |
| **MD/RD** (market data / reference data) | Market Data and Reference data. Often intermingled and provided on the same session. Messages in this domain convey identical information to all recipients authorized for a specific Market segment. Reference data can also be published via the OE gateway. |
| **PT/DC** (post-trade / drop copy) | Post-trade read-only services providing drop copies of executions as TCRs (TradeCaptureReport). Executing parties themselves must also connect to receive TCRs. Order events are communicated by means of ERs (ExecutionReport). |
| **CL** | Trade executions and reported trades subject to clearing are communicated to a connected clearing system as 2-sided TCRs. The MME matching system acts as the client and logs on to the clearing system. |

> The FIX Gateways support multiple FIX logon users on the same port. Each Gateway process supports only one port. Session reset normally occurs every 24 hours.

#### 1.3.4 FIX Session

##### 1.3.4.1 Company IDs

The Sender- and TargetCompID define the FIX session. Only one session can be active between two hosts simultaneously.

- **Marketplace Company ID** must be set on: `TargetCompID` of inbound transactions, `SenderCompID` of outbound transactions.
- **Client Company ID** must be set on: `SenderCompID` of inbound transactions, `TargetCompID` of outbound transactions.

##### 1.3.4.2 SenderSubID

Each inbound business transaction must have `SenderSubID` (tag 50) set to an authenticated user. The SenderSubID on incoming transactions will be echoed back in `TargetSubID` (tag 57) on outbound transactions.

##### 1.3.4.3 User Authentication

Users are authenticated by setting `Username (553)` and `Password (554)` in the Logon message.

##### 1.3.4.4 Logon

At Logon, clients are identified by CompIDs and IP Address. When authenticated, the system responds with a Logon message to the client.

##### 1.3.4.5 Logout

A failure at Logon may result in a Logout message with a SessionStatus code:

| SessionStatus | Reason |
|---------------|--------|
| 4 | Session logout complete |
| 5 | Invalid username or password |
| 6 | Account is locked or suspended |
| 7 | Logons are not allowed at this time |

##### 1.3.4.6 Encryption

The system does **not** support encryption, including password encryption in the Logon transaction.

##### 1.3.4.7 Character Encoding

MME FIX uses standard **US ASCII** encoding.

##### 1.3.4.8 Change of Password

Change of password over MME FIX is **not supported**.

##### 1.3.4.9 Failure Handling

Possible actions on failure:

| Action | Result |
|--------|--------|
| Disconnect | The user will be disconnected. |
| Logout | The user will receive a Logout. A Disconnect will follow within 3 seconds (or sooner if client sends a logout or other message). |
| Reject and Logout | The user will receive a Reject message, then a Logout message. Disconnect follows within 3 seconds. |
| Reject | The user will receive a Reject message. |

Failure-to-action mapping:

| Failure | Action |
|---------|--------|
| UserName in Logon does not exist, tag 553 not provided, or user is not a FIX OE user | Disconnect |
| SenderCompID or SubID not configured for user in Logon | Disconnect |
| SenderCompID or SubID does not match the one configured for a user in Logon | Disconnect |
| Any other message received before Logon | Disconnect |
| TargetCompID is incorrect in Logon | Disconnect |
| Password is incorrect in Logon | Disconnect |
| BeginString is missing or incorrect in Logon | Logout |
| TargetCompID is missing in Logon | Logout |
| HeartBtInt is missing or incorrect (lower than 10) in Logon | Logout |
| DefaultApplVerID is missing or incorrect (not "9") in Logon | Logout |
| Password is missing in Logon | Logout |
| SendingTime is missing or has incorrect format in Standard Header | Logout |
| Seq Num lower than last received and Logon does not have reset sequence number flag | Logout |
| Seq Num lower than last received and message does not have PossDupFlag set | Logout |
| BeginString is missing or incorrect for messages other than Logon | Reject and Logout |
| SenderCompID or TargetCompID is missing or incorrect for messages other than Logon | Reject and Logout |
| PossDup flag is set and OrigSendingTime is after SendingTime in Standard Header | Reject and Logout |
| SubID is missing or incorrect for messages other than Logon | Reject |
| Other message validation errors (e.g. missing required tags) | Reject |

---

### 1.4 Session Level Messages Supported

| FIX Message Type | Message Name | Over FIX Session | FIX 5.0 | FIX 4.4 |
|-----------------|-------------|-----------------|---------|---------|
| A | Logon | ALL | Y | Y |
| 0 | Heartbeat | ALL | Y | Y |
| 1 | Test Request | ALL | Y | Y |
| 2 | ResendRequest | ALL | Y | Y |
| 3 | Reject | ALL | Y | Y |
| 4 | SequenceReset | ALL | Y | Y |
| 5 | Logout | ALL | Y | Y |

---

### 1.5 Application Level Messages Supported

For message details, see the specification for the associated gateway.

| FIX Message Type | Message Name | Over FIX Session | FIX 5.0 | FIX 4.4 |
|-----------------|-------------|-----------------|---------|---------|
| j | BusinessMessageReject | MD, OE, MD | Y | — |
| 8 | ExecutionReport | OE, DC** | Y | Y |
| BT | MarkerDefinitionRequest | MD*** | Y | — |
| X | MarketDataIncrementalRefresh | MD | Y | — |
| V | MarketDataRequest | MD | Y | — |
| Y | MarketDataRequestReject | MD | Y | — |
| W | MarketDataSnapshotFullRefresh | MD | Y | — |
| BU | MarketDefinition | MD*** | Y | — |
| i | MassQuote | OE | Y | Y |
| b | MassQuoteAck | OE | Y | Y |
| D | NewOrderSingle | OE | Y | Y |
| B | News | MD, DC | Y | Only DC EoD |
| 9 | OrderCancelReject | OE | Y | Y |
| G | OrderCancelReplaceRequest | OE | Y | Y |
| F | OrderCancelRequest | OE | Y | Y |
| AF | OrderMassStatusRequest | DC** | Y | — |
| S | Quote | OE | Y | — |
| Z | QuoteCancel | OE | Y | — |
| R | QuoteRequest | OE | Y | — |
| AG | QuoteRequestReject | OE | Y | — |
| AJ | QuoteResponse | OE | Y | — |
| AI | QuoteStatusReport | OE | Y | — |
| d | SecurityDefinition | MD*** | Y | — |
| c | SecurityDefinitionRequest | MD*** | Y | — |
| y | SecurityList | MD***, OE* | Y | — |
| x | SecurityListRequest | MD***, OE* | Y | — |
| BK | SecurityListUpdateReport | MD***, OE* | Y | — |
| CO | SecurityMassStatus | MD | Y | — |
| CN | SecurityMassStatusRequest | MD | Y | — |
| AE | TradeCaptureReport | DC, OE, CL | Y | Only DC |
| AR | TradeCaptureReportAck | OE, CL | Y | Y |
| AD | TradeCaptureReportRequest | DC | Y | Y |
| AQ | TradeCaptureReportRequestAck | DC | Y | — |
| BJ | TradingSessionList | MD***, OE* | Y | — |
| BI | TradingSessionListRequest | MD***, OE* | Y | — |

> \* Available when reference data is configured for the OE gateway. FIX 5.0 only.  
> \*\* Available when order events are configured for the DC gateway. FIX 5.0 only.  
> \*\*\* Specified in Reference Data spec (RD).

---

## 2. Administrative Messages

### 2.1 Logon (A)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace / Out from NFF marketplace |
| **Message code** | A (Added FIX.2.7) |
| **FIX Session** | Order Entry (OE), Market Data (MD), Drop-copy Services (DC) |
| **Available to** | All FIX Actors recognized by the system |
| **Limitations** | Encryption not supported. Change of password not supported. |
| **Response** | Logon (A) |

**Usage:** A client sends a Logon message to initiate a FIX session. The marketplace acknowledges by returning another Logon message. Authentication uses `Username (553)` and `Password (554)`.

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = A |
| 98 | EncryptMethod | Y | Always unencrypted. Valid value: `0 = None / Other` |
| 108 | HeartBtInt | Y | Seconds; same value used by both sides; must be >= 10 |
| 141 | ResetSeqNumFlag | — | `N` = No reset; `Y` = Both sides reset sequence numbers |
| 553 | Username | Q | — |
| 554 | Password | Q | — |
| 1409 | SessionStatus | — | Session status at time of logon. Used when logon is sent as acknowledgement from acceptor. Valid value: `0 = Session active` |
| 1137 | DefaultApplVerID | Y | Default FIX version for the session. Valid values: `9 = FIX50SP2`, `6 = FIX44` |
| — | Standard Trailer | Y | — |

---

### 2.2 Heartbeat (0)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace / Out from NFF marketplace |
| **Message code** | 0 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | Sent at the interval set at Logon; also the response to a TestRequest message. |
| **Response** | — |

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = 0 |
| 112 | TestReqID | — | Required when heartbeat is the result of a Test Request message. |
| — | Standard Trailer | Y | — |

---

### 2.3 TestRequest (1)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace |
| **Message code** | 1 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | Tests communication and verifies sequence numbers — much like a "ping". |
| **Response** | Heartbeat (0) |

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = 1 |
| 112 | TestReqID | Y | Identifier to be returned in resulting Heartbeat. |
| — | Standard Trailer | Y | — |

---

### 2.4 ResendRequest (2)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace / Out from NFF marketplace |
| **Message code** | 2 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | Used if a sequence number gap is detected, if the receiving application lost a message, or as part of initialization. For a single message: `BeginSeqNo = EndSeqNo`. For all subsequent messages: `EndSeqNo = "0"` (infinity). |
| **Response** | Resent Application messages and/or SequenceReset (4) as GapFill |

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = 2 |
| 7 | BeginSeqNo | Y | Sequence number of first message in range to be resent. |
| 16 | EndSeqNo | Y | Sequence number of last message in range to be resent. |
| — | Standard Trailer | Y | — |

---

### 2.5 Reject (3)

| Field | Value |
|-------|-------|
| **Direction** | Out from NFF marketplace |
| **Message code** | 3 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | Sent when the FIX gateway can at least partially parse the message, but it does not adhere to the specification and cannot be delivered to the back-end system. |
| **Response** | — |

| Tag | FIX Tag Name | Req'd | UD | Description |
|-----|-------------|-------|-----|-------------|
| — | Standard Header | Y | — | MsgType = 3 |
| 45 | RefSeqNum | Y | — | MsgSeqNum of rejected message |
| 371 | RefTagID | — | — | Tag number of the FIX field being referenced (if applicable) |
| 372 | RefMsgType | — | — | MsgType of the FIX message being referenced |
| 373 | SessionRejectReason | Q | R100+ | Code identifying reason for session-level Reject. Valid values: `1` = Required Tag Missing, `2` = Tag Not Defined For This Message Type, `3` = Undefined Tag, `4` = Tag Specified Without A Value, `5` = Value Is Incorrect Out Of Range For This Tag, `6` = Incorrect Data Format For Value, `9` = Comp ID Problem, `10` = Sending Time Accuracy Problem, `11` = Invalid Msg Type, `15` = Repeating group fields out of order, `16` = Incorrect NumInGroup count for repeating group, `99` = Other; R100+ for customer-specific use |
| 58 | Text | — | — | Where possible, message to explain reason for rejection. |
| — | Standard Trailer | Y | — | — |

---

### 2.6 SequenceReset (4)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace / Out from NFF marketplace |
| **Message code** | 4 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | Common usage: `GapFillFlag = Y` in response to ResendRequest, indicating a range of messages will not be resent but fulfills the MsgSeqNum sequence (commonly used to avoid resending Heartbeats). Rare usage: reset sequence numbers to a higher number to escape a deadlock — only triggered by manual intervention. |
| **Limitations** | `GapFillFlag = N` should be used with greatest care. |
| **Response** | ExecutionReport (8) |

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = 4 |
| 123 | GapFillFlag | — | Indicates Sequence Reset is replacing messages that will not be resent. `N` = Sequence Reset, Ignore Msg Seq Num; `Y` = Gap Fill Message, Msg Seq Num Field Valid |
| 36 | NewSeqNo | Y | — |
| — | Standard Trailer | Y | — |

---

### 2.7 Logout (5)

| Field | Value |
|-------|-------|
| **Direction** | In to marketplace / Out from NFF marketplace |
| **Message code** | 5 (Added FIX.2.7) |
| **FIX Session** | OE, MD, DC |
| **Available to** | All FIX Actors |
| **Usage** | A client sends a Logout to initiate session termination; the marketplace acknowledges with another Logout (and vice versa). Disconnection without exchanging Logout messages should be interpreted as an abnormal condition. |
| **Response** | Logout (5) |

| Tag | FIX Tag Name | Req'd | UD | Description |
|-----|-------------|-------|-----|-------------|
| — | Standard Header | Y | — | MsgType = 5 |
| 1409 | SessionStatus | — | R100+ | Session status at time of logout. Valid values: `0` = Session active, `4` = Session logout complete, `6` = Account locked. NDAQ+ extensions: `101` = Heartbeat interval too low, `103` = User credentials could not be authenticated; R100+ for customer-specific use |
| 58 | Text | — | — | Placeholder for any additional session status from the marketplace. |
| — | Standard Trailer | Y | — | — |

---

### 2.8 BusinessMessageReject (j)

| Field | Value |
|-------|-------|
| **Direction** | Out from NFF Marketplace |
| **Message code** | j (Added FIX.4.2) |
| **FIX Session** | Reference Data (RD), Order Entry (OE) |
| **Available to** | All FIX Actors |
| **Usage** | Used when the NFF Marketplace cannot honor requests that lack rejection responses, including: `SecurityMassStatusRequest (CN)`, `TradingSessionListRequest (BJ)`, `OrderMassStatusRequest (AF)`. |
| **Response** | — |

| Tag | FIX Tag Name | Req'd | Description |
|-----|-------------|-------|-------------|
| — | Standard Header | Y | MsgType = j |
| 45 | RefSeqNum | — | MsgSeqNum of rejected message |
| 372 | RefMsgType | Y | MsgType of the FIX message being referenced. Valid values: `CN` = SecurityMassStatusRequest, `BJ` = TradingSessionListRequest, `AF` = OrderMassStatusRequest |
| 380 | BusinessRejectReason | Y | Code identifying reason. Valid values: `0` = Other, `1` = Unknown ID, `6` = Not Authorized |
| 58 | Text | — | Where possible, message to explain reason for rejection. |
| — | Standard Trailer | Y | — |

---

*© 2025 Indonesia Stock Exchange — Private and Confidential*
