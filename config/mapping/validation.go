package mapping

import (
	"fmt"

	"github.com/xrpscan/platform/config"
)

func IndexTemplateValidation(shards uint8, replicas uint8) string {
	template := `
	{
		"index_patterns": ["%[1]s.validations-*"],
		"template": {
			"settings": {
				"number_of_shards": %[2]d,
				"number_of_replicas": %[3]d
			},
			%[4]s,
			"aliases": {
				"%[1]s.validations": {}
			}
		},
		"priority": 128,
		"version": 1,
		"_meta": {
			"description": "%[1]s.validations template"
		}
	}
	`
	return fmt.Sprintf(template,
		config.EnvEsNamespace(), // %[1]s
		shards,                  // %[2]d
		replicas,                // %[3]d
		validationMapping(),     // %[4]s
	)
}

func validationMapping() string {
	return `
	"mappings": {
		"_source": {
			"enabled": true
		},
		"properties": {
			"amendments":            { "type": "keyword" },
			"cookie":                { "type": "keyword" },
			"data":                  { "type": "text", "index": false },
			"flags":                 { "type": "long" },
			"full":                  { "type": "boolean" },
			"ledger_hash":           { "type": "keyword" },
			"ledger_index":          { "type": "long" },
			"load_fee":              { "type": "long" },
			"master_key":            { "type": "keyword" },
			"reserve_base":          { "type": "long" },
			"reserve_inc":           { "type": "long" },
			"server_version":        { "type": "keyword" },
			"signature":             { "type": "text", "index": false },
			"signing_time":          { "type": "long" },
			"type":                  { "type": "keyword" },
			"validated_hash":        { "type": "keyword" },
			"validation_public_key": { "type": "keyword" }
		}
	}`
}
