package actions

import (
	"log"
	"net/http"

	"github.com/forbole/juno/v2/cmd/parse"
	"github.com/spf13/cobra"

	"github.com/forbole/bdjuno/v2/cmd/actions/handlers"
)

// NewActionsCmd returns the Cobra command allowing to activate hasura actions
func NewActionsCmd(parseCfg *parse.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "hasura-actions",
		Short:   "Activate hasura actions",
		PreRunE: parse.ReadConfig(parseCfg),
		RunE: func(cmd *cobra.Command, args []string) error {

			// HTTP server for the handler
			mux := http.NewServeMux()
			mux.HandleFunc("/account_balances", handlers.AccountBalances)
			mux.HandleFunc("/total_supply", handlers.TotalSupply)

			mux.HandleFunc("/delegator_rewards", handlers.DelegatorRewards)
			mux.HandleFunc("/validator_commission", handlers.ValidatorCommission)

			err := http.ListenAndServe(":3000", mux)
			log.Fatal(err)

			return nil
		},
	}
}
