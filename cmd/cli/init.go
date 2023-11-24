/**
* This file implements `platform-cli init` subcommand
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/config/mapping"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/platform/signals"
)

const InitCommandName = "init"
const defaultShards int = 16
const defaultReplicas int = 0

type InitCommand struct {
	fs             *flag.FlagSet
	fConfigFile    string
	fShards        int
	fReplicas      int
	fElasticsearch bool
	fVerbose       bool
}

func NewInitCommand() *InitCommand {
	cmd := &InitCommand{
		fs: flag.NewFlagSet(InitCommandName, flag.ExitOnError),
	}

	cmd.fs.BoolVar(&cmd.fElasticsearch, "elasticsearch", false, "Initialize Elasticsearch")
	cmd.fs.IntVar(&cmd.fShards, "shards", defaultShards, "Number of shards")
	cmd.fs.IntVar(&cmd.fReplicas, "replicas", defaultReplicas, "Number of replicas")
	cmd.fs.StringVar(&cmd.fConfigFile, "config", ".env", "Environment config file")
	cmd.fs.BoolVar(&cmd.fVerbose, "verbose", false, "Make the command more talkative")
	return cmd
}

func (cmd *InitCommand) Init(args []string) error {
	err := cmd.fs.Parse(args)
	if err != nil {
		return err
	}

	return cmd.Validate()
}

func (cmd *InitCommand) Validate() error {
	// Ledgers are backfilled in chronological order. Therefore, --from ledger
	// index must be less than --to ledger index.
	if cmd.fShards > 128 {
		return fmt.Errorf("too many shards (%d)", cmd.fShards)
	}
	if cmd.fReplicas > 128 {
		return fmt.Errorf("too many replicas (%d)", cmd.fReplicas)
	}
	return nil
}

func (cmd *InitCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *InitCommand) Run() error {
	// Register command line signal handlers to gracefully shutdown cli
	signals.HandleAll()

	// Load validated config file
	config.EnvLoad(cmd.fConfigFile)

	// Initialize connections to services
	logger.New()
	connections.NewEsClient()

	// List of supported streams for which we can create index templates
	supportedTemplates := map[string]func(uint8, uint8) string{
		models.StreamLedger.String():      mapping.IndexTemplateLedger,
		models.StreamValidation.String():  mapping.IndexTemplateValidation,
		models.StreamTransaction.String(): mapping.IndexTemplateTransaction,
	}

	for template, mapper := range supportedTemplates {
		templateWithNamespace := fmt.Sprintf("%s.%s", config.EnvEsNamespace(), template)
		if cmd.ExistsIndexTemplate(templateWithNamespace) {
			fmt.Printf("IndexTemplate exists: %s\n", templateWithNamespace)
		} else {
			err := cmd.PutIndexTemplate(templateWithNamespace, mapper(uint8(cmd.fShards), uint8(cmd.fReplicas)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

// Check if the IndexTemplate exists
func (cmd *InitCommand) ExistsIndexTemplate(templateName string) bool {
	req := esapi.IndicesExistsIndexTemplateRequest{
		Name: templateName,
	}
	res, err := req.Do(context.Background(), connections.EsClient)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	return res.StatusCode == http.StatusOK
}

// Create index template
func (cmd *InitCommand) PutIndexTemplate(templateName string, template string) error {
	req := esapi.IndicesPutIndexTemplateRequest{
		Name: templateName,
		Body: strings.NewReader(template),
	}
	res, err := req.Do(context.Background(), connections.EsClient)
	if err != nil {
		return fmt.Errorf("cannot PUT IndexTemplate %s", templateName)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("Response: %s\n", string(body))

	return nil
}
