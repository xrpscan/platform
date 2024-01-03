package mapping

import (
	"fmt"

	"github.com/xrpscan/platform/config"
)

func IndexTemplateLedger(shards uint8, replicas uint8) string {
	template := `
	{
		"index_patterns": ["%[1]s.ledger-*"],
		"template": {
			"settings": {
				"number_of_shards": %[2]d,
				"number_of_replicas": %[3]d
			},
			%[4]s,
			"aliases": {
				"%[1]s.ledger": {}
			}
		},
		"priority": 128,
		"version": 1,
		"_meta": {
			"description": "%[1]s.ledger template"
		}
	}
	`
	return fmt.Sprintf(template,
		config.EnvEsNamespace(), // %[1]s
		shards,                  // %[2]s
		replicas,                // %[3]s
		ledgerMapping(),         // %[4]s
	)
}

func ledgerMapping() string {
	return `
    "mappings": {
		"_source": {
			"enabled": true
		},
		"properties": {
			"account_hash":          { "type": "keyword" },
			"close_flags":           { "type": "integer" },
			"close_time":            { "type": "integer" },
			"close_time_human":      { "type": "keyword" },
			"close_time_resolution": { "type": "integer" },
			"closed":                { "type": "boolean" },
			"fee_base":              { "type": "long" },
			"fee_ref":               { "type": "long" },
			"ledger_hash":           { "type": "keyword" },
			"ledger_index":          { "type": "long" },
			"ledger_time":           { "type": "long" },
			"parent_close_time":     { "type": "integer" },
			"parent_hash":           { "type": "keyword" },
			"reserve_base":          { "type": "long" },
			"reserve_inc":           { "type": "long" },
			"total_coins":           { "type": "long" },
			"transaction_hash":      { "type": "keyword" },
			"txn_count":             { "type": "long" },
			"type":                  { "type": "keyword" },
			"validated_ledgers":     { "type": "keyword" }
		}
	}`
}
