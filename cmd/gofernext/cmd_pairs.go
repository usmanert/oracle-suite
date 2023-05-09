package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sort"

	"github.com/spf13/cobra"

	"github.com/chronicleprotocol/oracle-suite/pkg/config"
	"github.com/chronicleprotocol/oracle-suite/pkg/pricenext/provider"
	"github.com/chronicleprotocol/oracle-suite/pkg/util/maputil"
)

func NewPairsCmd(opts *options) *cobra.Command {
	return &cobra.Command{
		Use:     "pairs [PAIR...]",
		Aliases: []string{"pair"},
		Args:    cobra.MinimumNArgs(0),
		Short:   "List all supported asset pairs",
		Long:    `List all supported asset pairs.`,
		RunE: func(c *cobra.Command, args []string) (err error) {
			if err := config.LoadFiles(&opts.Config, opts.ConfigFilePath); err != nil {
				return err
			}
			ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Interrupt)
			defer ctxCancel()
			services, err := opts.Config.Services(opts.Logger())
			if err != nil {
				return err
			}
			if err = services.Start(ctx); err != nil {
				return err
			}
			models, err := services.PriceProvider.Models(ctx, services.PriceProvider.ModelNames(ctx)...)
			if err != nil {
				return err
			}
			marshaled, err := marshalModels(models, opts.Format.format)
			if err != nil {
				return err
			}
			fmt.Println(string(marshaled))
			return nil
		},
	}
}

func marshalModels(models map[string]provider.Model, format string) ([]byte, error) {
	switch format {
	case formatPlain:
		return marshalModelsPlain(models)
	case formatTrace:
		return marshalModelsTrace(models)
	case formatJSON:
		return marshalModelsJSON(models)
	default:
		return nil, fmt.Errorf("unsupported format")
	}
}

func marshalModelsPlain(models map[string]provider.Model) ([]byte, error) {
	var buf bytes.Buffer
	for _, name := range maputil.SortKeys(models, sort.Strings) {
		bts, err := models[name].MarshalText()
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("Model for %s:\n", name))
		buf.Write(bts)
	}
	return buf.Bytes(), nil
}

func marshalModelsTrace(models map[string]provider.Model) ([]byte, error) {
	var buf bytes.Buffer
	for _, name := range maputil.SortKeys(models, sort.Strings) {
		bts, err := models[name].MarshalTrace()
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("Model for %s:\n", name))
		buf.Write(bts)
	}
	return buf.Bytes(), nil
}

func marshalModelsJSON(models map[string]provider.Model) ([]byte, error) {
	return json.Marshal(models)
}
