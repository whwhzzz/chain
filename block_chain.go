package main

import (
	"fmt"

	ont_rpc "github.com/ontio/ontology-go-sdk/rpc"
	"github.com/ontio/ontology/common"
	ont_types "github.com/ontio/ontology/core/types"
)

func checkPayload(payload *ont_types.Payload) {
	return
}

func CheckTheOfferID(rpc *ont_rpc.RpcClient, res map[string]common.Uint256) {
	count, err := rpc.GetBlockCount()
	if err != nil {
		fmt.Println("GetBlockCount error:%s", err)
		return
	}
	var i uint32 = 0
	for ; i < count; i++ {
		block, err := rpc.GetBlockByHeight(i)
		if err != nil {
			fmt.Println("GetBlock error:Block %d %s", i, err)
			continue
		}
		for _, trans := range block.Transactions {
			payload := trans.Payload
			// TODO: Extract the offer_id from payload
			checkPayload(&payload)
			offer_id := "TODO"
			_, ok := res[offer_id]
			if ok {
				res[offer_id] = trans.Hash()
			}
		}
	}
}
