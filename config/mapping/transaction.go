package mapping

import (
	"fmt"

	"github.com/xrpscan/platform/config"
)

func IndexTemplateTransaction(shards uint8, replicas uint8) string {
	template := `
	{
		"index_patterns": ["%[1]s.transactions-*"],
		"template": {
			"settings": {
				"number_of_shards": %[2]d,
				"number_of_replicas": %[3]d
			},
			%[4]s,
			"aliases": {
				"%[1]s.transactions": {}
			}
		},
		"priority": 128,
		"version": 1,
		"_meta": {
			"description": "%[1]s.transaction template"
		}
	}
	`
	return fmt.Sprintf(template,
		config.EnvEsNamespace(), // %[1]s
		shards,                  // %[2]s
		replicas,                // %[3]s
		transactionMapping(),    // %[4]s
	)
}

func transactionMapping() string {
	return `
	"mappings": {
		"_source": {
			"enabled": true
		},
		"properties": {
			"Account": { "type": "keyword" },
			"AccountTxnID": { "type": "keyword" },
			"Amendment": { "type": "keyword" },
			"Amount": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}

				}
			},
			"Authorize": { "type": "keyword" },
			"Balance": { "type": "keyword" },
			"CancelAfter": { "type": "long" },
			"Channel": { "type": "keyword" },
			"CheckID": { "type": "keyword" },
			"ClearFlag": { "type": "long" },
			"Condition": { "type": "keyword" },
			"DeliverMax": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"DeliverMin": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"Destination": { "type": "keyword" },
			"DestinationTag": { "type": "long" },
			"Domain": { "type": "keyword" },
			"EmailHash": { "type": "keyword" },
			"Expiration": { "type": "long" },
			"Fee": { "type": "long" },
			"FinishAfter": { "type": "long" },
			"Flags": { "type": "long" },
			"Fulfillment": { "type": "keyword" },
			"InvoiceID": { "type": "keyword" },
			"Issuer": { "type": "keyword" },
			"LastLedgerSequence": { "type": "long" },
			"LedgerSequence": { "type": "long" },
			"LimitAmount": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"Memos": {
				"properties": {
					"Memo": {
						"properties": {
							"MemoData": { "type": "text" },
							"MemoFormat": { "type": "text" },
							"MemoType": { "type": "text" }
						}
					}
				}
			},
			"MessageKey": { "type": "keyword" },
			"NFTokenBrokerFee": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"NFTokenBuyOffer": { "type": "keyword" },
			"NFTokenID": { "type": "keyword" },
			"NFTokenMinter":  { "type": "keyword" },
			"NFTokenOffers": { "type": "keyword" },
			"NFTokenSellOffer": { "type": "keyword" },
			"NFTokenTaxon": { "type": "long" },
			"OfferSequence": { "type": "long" },
			"OperationLimit": { "type": "long" },
			"Owner": { "type": "keyword" },
			"Paths": {
				"properties": {
					"account": { "type": "keyword" },
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"type": { "type": "long" }
				}
			},
			"PublicKey": { "type": "keyword" },
			"QualityIn": { "type": "long" },
			"QualityOut": { "type": "long" },
			"RegularKey": { "type": "keyword" },
			"SendMax": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"Sequence": { "type": "long" },
			"SetFlag": { "type": "long" },
			"SettleDelay": { "type": "long" },
			"Signature": { "type": "keyword" },
			"SignerEntries": {
				"properties": {
					"SignerEntry": {
						"properties": {
							"Account": { "type": "keyword" },
							"SignerWeight": { "type": "long" },
							"WalletLocator": { "type": "keyword" }
						}
					}
				}
			},
			"SignerQuorum": { "type": "long" },
			"Signers": {
				"properties": {
					"Signer": {
						"properties": {
							"Account": { "type": "keyword" },
							"SigningPubKey": { "type": "keyword" },
							"TxnSignature": { "type": "keyword" }
						}
					}
				}
			},
			"SigningPubKey": { "type": "keyword" },
			"SourceTag": { "type": "long" },
			"TakerGets": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"TakerPays": {
				"properties": {
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "long" },
					"native": { "type": "boolean"}
				}
			},
			"TickSize": { "type": "long" },
			"TicketCount": { "type": "long" },
			"TicketSequence": { "type": "long" },
			"TransactionType": { "type": "keyword" },
			"TransferFee": { "type": "long" },
			"TransferRate": { "type": "long" },
			"TxnSignature": { "type": "keyword" },
			"UNLModifyDisabling": { "type": "long" },
			"UNLModifyValidator": { "type": "keyword" },
			"URI": { "type": "keyword" },
			"Unauthorize": { "type": "keyword" },
			"date": { "type": "long" },
			"hash": { "type": "keyword" },
			"ledger_index": { "type": "long" },
			"meta": {
				"properties": {
					"DeliveredAmount": {
						"properties": {
							"currency": { "type": "keyword" },
							"issuer": { "type": "keyword" },
							"value":  { "type": "long" },
							"native": { "type": "boolean"}		
						}
					},
					"TransactionIndex": { "type": "integer" },
					"TransactionResult": { "type": "keyword" },
					"delivered_amount": {
						"properties": {
							"currency": { "type": "keyword" },
							"issuer": { "type": "keyword" },
							"value":  { "type": "long" },
							"native": { "type": "boolean"}		
						}
					}
				}
			},
			"validated": { "type": "boolean" }
		}
	}`
}
