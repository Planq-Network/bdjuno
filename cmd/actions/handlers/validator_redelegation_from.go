package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"
)

func RedelegationFromValidator(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload actionstypes.StakingPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload: failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	result, err := getValidatorRedelegation(actionPayload.Input)
	if err != nil {
		errorHandler(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

func getValidatorRedelegation(input actionstypes.StakingArgs) (actionstypes.RedelegationResponse, error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return actionstypes.RedelegationResponse{}, err
	}

	// Get latest node height
	height, err := parseCtx.Node.LatestHeight()
	if err != nil {
		return actionstypes.RedelegationResponse{}, fmt.Errorf("error while getting chain latest block height: %s", err)
	}

	pagination := &query.PageRequest{
		Offset:     input.Offset,
		Limit:      input.Limit,
		CountTotal: input.CountTotal,
	}

	// Get redelegations from a source validator address
	redelegationRequest := &stakingtypes.QueryRedelegationsRequest{
		SrcValidatorAddr: input.Address,
		Pagination:       pagination,
	}
	redelegations, err := sources.StakingSource.GetRedelegations(height, redelegationRequest)
	if err != nil {
		return actionstypes.RedelegationResponse{}, fmt.Errorf("error while getting delegator redelegations: %s", err)
	}

	redelegationsList := make([]actionstypes.Redelegation, len(redelegations.RedelegationResponses))
	for index, del := range redelegations.RedelegationResponses {
		redelegationsList[index] = actionstypes.Redelegation{
			DelegatorAddress:    del.Redelegation.DelegatorAddress,
			ValidatorSrcAddress: del.Redelegation.ValidatorSrcAddress,
			ValidatorDstAddress: del.Redelegation.ValidatorDstAddress,
		}

		// Handle Completion Time and Balance
		RedelegationEntriesList := make([]actionstypes.RedelegationEntry, len(del.Entries))
		for indexEntry, entry := range del.Entries {
			RedelegationEntriesList[indexEntry] = actionstypes.RedelegationEntry{
				CompletionTime: entry.RedelegationEntry.CompletionTime,
				Balance:        entry.Balance,
			}
		}
		redelegationsList[index].RedelegationEntries = RedelegationEntriesList
	}

	return actionstypes.RedelegationResponse{
		Redelegations: redelegationsList,
		Pagination:    redelegations.Pagination,
	}, nil
}