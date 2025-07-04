package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
)

// GetQueryCmd returns the query commands for IBC clients
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "IBC client query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryClientStates(),
		GetCmdQueryCounterpartyInfo(),
		GetCmdQueryClientState(),
		GetCmdQueryClientStatus(),
		GetCmdQueryConsensusStates(),
		GetCmdQueryConsensusStateHeights(),
		GetCmdQueryConsensusState(),
		GetCmdQueryHeader(),
		GetCmdSelfConsensusState(),
		GetCmdClientParams(),
		GetCmdClientParams(),
		GetCmdQueryClientCreator(),
		GetCmdQueryClientConfig(),
	)

	return queryCmd
}

// NewTxCmd returns the command to create and handle IBC clients
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "IBC client transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		newCreateClientCmd(),
		newAddCounterpartyCmd(),
		newUpdateClientCmd(),
		newUpgradeClientCmd(),
		newSubmitRecoverClientProposalCmd(),
		newScheduleIBCUpgradeProposalCmd(),
		newUpdateClientConfigCmd(),
		newDeleteClientCreatorCmd(),
	)

	return txCmd
}
