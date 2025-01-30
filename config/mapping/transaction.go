package mapping

import (
	"fmt"

	"github.com/xrpscan/platform/config"
)

const ScalingFactor = 1000000

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
	return fmt.Sprintf(`
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
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean" }
				}
			},
			"Amount2": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"Asset": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"Asset2": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"AssetClass":  { "type": "keyword" },
			"AuthAccounts": {
				"properties": {
					"AuthAccount": {
						"properties": {
							"Account": { "type": "keyword" }
						}
					}
				}
			},
			"Authorize": { "type": "keyword" },
			"Balance": { "type": "keyword" },
			"BidMin": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"BidMax": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"CancelAfter": { "type": "long" },
			"_CancelAfter": { "type": "date" },
			"Channel": { "type": "keyword" },
			"CheckID": { "type": "keyword" },
			"ClearFlag": { "type": "long" },
			"Condition": { "type": "keyword" },
			"ctid": { "type": "keyword" },
			"Data": { "type": "keyword" },
			"DeliverMax": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"DeliverMin": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"Destination": { "type": "keyword" },
			"DestinationTag": { "type": "long" },
			"DIDDocument": { "type": "keyword" },
			"Domain": { "type": "keyword" },
			"EmailHash": { "type": "keyword" },
			"EPrice": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"Expiration": { "type": "long" },
			"_Expiration": { "type": "date" },
			"Fee": { "type": "long" },
			"_Fee": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
			"FinishAfter": { "type": "long" },
			"_FinishAfter": { "type": "date" },
			"Flags": { "type": "long" },
			"Fulfillment": { "type": "keyword" },
			"Holder": { "type": "keyword" },
			"InvoiceID": { "type": "keyword" },
			"Issuer": { "type": "keyword" },
			"LastLedgerSequence": { "type": "long" },
			"LastUpdateTime": { "type": "long" },
			"_LastUpdateTime": { "type": "date" },
			"LedgerSequence": { "type": "long" },
			"LimitAmount": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"LPTokenIn": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"LPTokenOut": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
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
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
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
			"OracleDocumentID": { "type": "long" },
			"Owner": { "type": "keyword" },
			"Paths": {
				"properties": {
					"account": { "type": "keyword" },
					"currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"type": { "type": "long" }
				}
			},
			"PriceDataSeries": {
				"properties": {
					"PriceData": {
						"properties": {
							"BaseAsset": { "type": "keyword" },
							"QuoteAsset": { "type": "keyword" },
							"AssetPrice": { "type": "long" },
							"Scale": { "type": "short" }
						}
					}
				}
			},
			"Provider": { "type": "keyword" },
			"PublicKey": { "type": "keyword" },
			"QualityIn": { "type": "long" },
			"QualityOut": { "type": "long" },
			"RegularKey": { "type": "keyword" },
			"SendMax": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
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
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"TakerPays": {
				"properties": {
					"currency": { "type": "keyword" },
					"_currency": { "type": "keyword" },
					"issuer": { "type": "keyword" },
					"value":  { "type": "keyword" },
					"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
					"native": { "type": "boolean"}
				}
			},
			"TickSize": { "type": "long" },
			"TicketCount": { "type": "long" },
			"TicketSequence": { "type": "long" },
			"TransactionType": { "type": "keyword" },
			"TradingFee": { "type": "integer" },
			"TransferFee": { "type": "long" },
			"TransferRate": { "type": "long" },
			"TxnSignature": { "type": "keyword" },
			"UNLModifyDisabling": { "type": "long" },
			"UNLModifyValidator": { "type": "keyword" },
			"URI": { "type": "keyword" },
			"Unauthorize": { "type": "keyword" },
			"date": { "type": "long" },
			"_date": { "type": "date" },
			"hash": { "type": "keyword" },
			"inLedger": { "type": "long" },
			"ledger_index": { "type": "long" },
			"meta": {
				"properties": {
					"DeliveredAmount": {
						"properties": {
							"currency": { "type": "keyword" },
							"_currency": { "type": "keyword" },
							"issuer": { "type": "keyword" },
							"value":  { "type": "keyword" },
							"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
							"native": { "type": "boolean"}
						}
					},
					"TransactionIndex": { "type": "integer" },
					"TransactionResult": { "type": "keyword" },
					"delivered_amount": {
						"properties": {
							"currency": { "type": "keyword" },
							"_currency": { "type": "keyword" },
							"issuer": { "type": "keyword" },
							"value":  { "type": "keyword" },
							"_value": { "type": "scaled_float", "scaling_factor": %[1]v, "ignore_malformed": true },
							"native": { "type": "boolean"}
						}
					}
				}
			},
			"validated": { "type": "boolean" }
		}
	}`, ScalingFactor)
}
