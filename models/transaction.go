package models

type Transaction struct {
	// Transaction response fields - https://xrpl.org/tx.html#response-format
	Date        uint32 `json:"date,omitempty"`
	Hash        string `json:"hash,omitempty"`
	LedgerIndex uint32 `json:"ledger_index,omitempty"`
	InLedger    uint32 `json:"inLedger,omitempty"`
	Validated   bool   `json:"validated,omitempty"`

	// Common fields - https://xrpl.org/transaction-common-fields.html
	Account            string   `json:"Account,omitempty"`
	TransactionType    string   `json:"TransactionType,omitempty"`
	Fee                string   `json:"Fee,omitempty"`
	Sequence           uint32   `json:"Sequence,omitempty"`
	AccountTxnID       string   `json:"AccountTxnID,omitempty"`
	PreviousTxnID      string   `json:"PreviousTxnID,omitempty"`
	Flags              uint32   `json:"Flags,omitempty"`
	LastLedgerSequence uint32   `json:"LastLedgerSequence,omitempty"`
	Memos              []Memo   `json:"Memos,omitempty"`
	NetworkID          uint32   `json:"NetworkID,omitempty"`
	Signers            []Signer `json:"Signers,omitempty"`
	SourceTag          uint32   `json:"SourceTag,omitempty"`
	SigningPubKey      string   `json:"SigningPubKey,omitempty"`
	TicketSequence     uint32   `json:"TicketSequence,omitempty"`
	TxnSignature       string   `json:"TxnSignature,omitempty"`

	// Metadata fields - https://xrpl.org/transaction-metadata.html
	Meta     Meta `json:"meta,omitempty"`
	MetaData Meta `json:"metaData,omitempty"`

	// AccountDelete fields - https://xrpl.org/accountdelete.html#accountdelete-fields
	// Shared fields:
	// Destination    string
	// DestinationTag uint32

	// AccountSet fields - https://xrpl.org/accountset.html#accountset-fields
	ClearFlag     uint32 `json:"ClearFlag,omitempty"`
	Domain        string `json:"Domain,omitempty"`
	EmailHash     string `json:"EmailHash,omitempty"`
	MessageKey    string `json:"MessageKey,omitempty"`
	NFTokenMinter string `json:"NFTokenMinter,omitempty"`
	SetFlag       uint32 `json:"SetFlag,omitempty"`
	TransferRate  uint32 `json:"TransferRate,omitempty"`
	TickSize      uint8  `json:"TickSize,omitempty"`
	WalletLocator string `json:"WalletLocator,omitempty"`
	WalletSize    uint32 `json:"WalletSize,omitempty"`

	// CheckCancel fields - https://xrpl.org/checkcancel.html#checkcancel-fields
	// Shared fields:
	// CheckID string

	// CheckCash fields - https://xrpl.org/checkcash.html#checkcash-fields
	// Shared fields:
	// Amount     Currency
	// DeliverMin Currency
	CheckID string `json:"CheckID,omitempty"`

	// CheckCreate fields - https://xrpl.org/checkcreate.html#checkcreate-fields
	// Shared fields:
	// Destination    string
	// SendMax        Currency
	// DestinationTag uint32
	// InvoiceID      string
	// Expiration     uint32

	// Clawback fields
	// Shared fields:
	// Amount Currency

	// DepositPreauth fields - https://xrpl.org/depositpreauth.html#depositpreauth-fields
	Authorize   string `json:"Authorize,omitempty"`
	Unauthorize string `json:"Unauthorize,omitempty"`

	// EnableAmendment fields - https://xrpl.org/enableamendment.html#enableamendment-fields
	// Shared fields:
	// LedgerSequence uint32
	Amendment string `json:"Amendment,omitempty"`

	// EscrowCancel fields - https://xrpl.org/escrowcancel.html#escrowcancel-fields
	// Shared fields:
	// Owner         string
	// OfferSequence uint32

	// EscrowCreate fields - https://xrpl.org/escrowcreate.html#escrowcreate-fields
	// Shared fields:
	// Amount Currency
	// Destination    string
	// DestinationTag uint32
	CancelAfter uint32 `json:"CancelAfter,omitempty"`
	FinishAfter uint32 `json:"FinishAfter,omitempty"`
	Condition   string `json:"Condition,omitempty"`

	// EscrowFinish fields - https://xrpl.org/escrowfinish.html#escrowfinish-fields
	// Shared fields:
	// Condition string
	// Owner     string
	// OfferSequence uint32
	Fulfillment string `json:"Fulfillment,omitempty"`

	// NFTokenAcceptOffer fields - https://xrpl.org/nftokenacceptoffer.html#nftokenacceptoffer-fields
	NFTokenSellOffer string   `json:"NFTokenSellOffer,omitempty"`
	NFTokenBuyOffer  string   `json:"NFTokenBuyOffer,omitempty"`
	NFTokenBrokerFee Currency `json:"NFTokenBrokerFee,omitempty"`

	// NFTokenBurn fields - https://xrpl.org/nftokenburn.html#nftokenburn-fields
	// Shared fields:
	// NFTokenID string
	// Owner     string

	// NFTokenCancelOffer fields - https://xrpl.org/nftokencanceloffer.html#nftokencanceloffer-fields
	NFTokenOffers []string `json:"NFTokenOffers,omitempty"`

	// NFTokenCreateOffer fields - https://xrpl.org/nftokencreateoffer.html#nftokencreateoffer-fields
	// Shared fields:
	// Owner       String
	// NFTokenID   String
	// Amount      Currency
	// Expiration  Number
	// Destination String

	// NFTokenMint fields - https://xrpl.org/nftokenmint.html#nftokenmint-fields
	NFTokenTaxon uint32 `json:"NFTokenTaxon,omitempty"`
	Issuer       string `json:"Issuer,omitempty"`
	TransferFee  uint16 `json:"TransferFee,omitempty"`
	URI          string `json:"URI,omitempty"`

	// OfferCancel fields - https://xrpl.org/offercancel.html#offercancel-fields
	// Shared fields:
	// OfferSequence uint32

	// OfferCreate fields - https://xrpl.org/offercreate.html#offercreate-fields
	// Shared fields:
	// Expiration    uint32
	// OfferSequence uint32
	TakerGets Currency `json:"TakerGets,omitempty"`
	TakerPays Currency `json:"TakerPays,omitempty"`

	// Payment fields - https://xrpl.org/payment.html#payment-fields
	Amount         Currency `json:"Amount,omitempty"`
	Destination    string   `json:"Destination,omitempty"`
	DestinationTag int      `json:"DestinationTag,omitempty"`
	InvoiceID      string   `json:"InvoiceID,omitempty"`
	Paths          []Path   `json:"Paths,omitempty"`
	SendMax        Currency `json:"SendMax,omitempty"`
	DeliverMin     Currency `json:"DeliverMin,omitempty"`

	// PaymentChannelClaim fields - https://xrpl.org/paymentchannelclaim.html#paymentchannelclaim-fields
	// Shared fields:
	// Amount    string
	// PublicKey string
	Channel   string `json:"Channel,omitempty"`
	Balance   string `json:"Balance,omitempty"`
	Signature string `json:"Signature,omitempty"`

	// PaymentChannelCreate fields - https://xrpl.org/paymentchannelcreate.html#paymentchannelcreate-fields
	// Shared fields:
	// Amount         string
	// Destination    string
	// CancelAfter    uint32
	// DestinationTag uint32
	SettleDelay uint32 `json:"SettleDelay,omitempty"`
	PublicKey   string `json:"PublicKey,omitempty"`

	// PaymentChannelFund fields - https://xrpl.org/paymentchannelfund.html#paymentchannelfund-fields
	// Shared fields:
	// Channel    string
	// Amount     Currency
	// Expiration uint32

	// SetFee fields - https://xrpl.org/setfee.html#setfee-fields
	// Shared fields:
	// LedgerSequence        uint32
	BaseFee               uint64 `json:"BaseFee,omitempty"`
	ReferenceFeeUnits     uint32 `json:"ReferenceFeeUnits,omitempty"`
	ReserveBase           uint32 `json:"ReserveBase,omitempty"`
	ReserveIncrement      uint32 `json:"ReserveIncrement,omitempty"`
	BaseFeeDrops          string `json:"BaseFeeDrops,omitempty"`
	ReserveBaseDrops      string `json:"ReserveBaseDrops,omitempty"`
	ReserveIncrementDrops string `json:"ReserveIncrementDrops,omitempty"`

	// SetRegularKey fields - https://xrpl.org/setregularkey.html#setregularkey-fields
	RegularKey string `json:"RegularKey,omitempty"`

	// SignerListSet fields - https://xrpl.org/signerlistset.html#signerlistset-fields
	SignerQuorum  uint32        `json:"SignerQuorum,omitempty"`
	SignerEntries []SignerEntry `json:"SignerEntries,omitempty"`

	// TicketCreate fields - https://xrpl.org/ticketcreate.html#ticketcreate-fields
	TicketCount uint32 `json:"TicketCount,omitempty"`

	// TrustSet fields - https://xrpl.org/trustset.html#trustset-fields
	LimitAmount Currency `json:"LimitAmount,omitempty"`
	QualityIn   uint32   `json:"QualityIn,omitempty"`
	QualityOut  uint32   `json:"QualityOut,omitempty"`

	// UNLModify fields - https://xrpl.org/unlmodify.html#unlmodify-fields
	// Shared fields:
	// LedgerSequence uint32
	UNLModifyDisabling uint8  `json:"UNLModifyDisabling,omitempty"`
	UNLModifyValidator string `json:"UNLModifyValidator,omitempty"`

	// Shared fields
	Owner          string `json:"Owner,omitempty"`
	LedgerSequence uint32 `json:"LedgerSequence,omitempty"`
	NFTokenID      string `json:"NFTokenID,omitempty"`
	Expiration     uint32 `json:"Expiration,omitempty"`
	OfferSequence  uint32 `json:"OfferSequence,omitempty"`
}
