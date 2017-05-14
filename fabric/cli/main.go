/*
Copyright SecureKey Technologies Inc. All Rights Reserved.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"strconv"
    "log"

	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
)

func main() {

	testSetup := BaseSetupImpl{
		ConfigFile:      "../config/config_test.yaml",
		ChainID:         "testchannel",
		ChannelConfig:   "../channel/testchannel.tx",
		ConnectEventHub: true,
	}

	if err := testSetup.Initialize(); err != nil {
		log.Fatalf(err.Error())
	}

	chain := testSetup.Chain

	// Test Query Info - retrieve values before transaction
	tx, err := chain.QueryInfo()
	if err != nil {
		log.Fatalf("QueryInfo return error: %v", err)
	}
    fmt.Println(tx)

    block, err := chain.QueryBlockByHash(tx.CurrentBlockHash)
    if err != nil {
        log.Fatalf("QueryBlockByHash return error: %v", err)                                                                               
    }
    fmt.Println(block)

/*
	testQueryBlock(chain)

	testQueryChannels(testSetup, chain)

	testInstalledChaincodes(testSetup, chain)

	testQueryByChaincode(chain)

	// TODO: Synch with test in node SDK when it becomes available
	// testInstantiatedChaincodes(t, chain)
*/
}

func changeBlockState(testSetup BaseSetupImpl) (string, error) {

	value, err := testSetup.QueryAsset()
	if err != nil {
		return "", fmt.Errorf("getQueryValue return error: %v", err)
	}

	// Start transaction that will change block state
	txID, err := testSetup.MoveFunds()
	if err != nil {
		return "", fmt.Errorf("Move funds return error: %v", err)
	}

	valueAfterInvoke, err := testSetup.QueryAsset()
	if err != nil {
		return "", fmt.Errorf("getQueryValue return error: %v", err)
	}

	// Verify that transaction changed block state
	valueInt, _ := strconv.Atoi(value)
	valueInt = valueInt + 1
	valueAfterInvokeInt, _ := strconv.Atoi(valueAfterInvoke)
	if valueInt != valueAfterInvokeInt {
		return "", fmt.Errorf("SendTransaction didn't change the QueryValue %s", value)
	}

	return txID, nil
}

func testQueryTransaction(chain fabricClient.Chain, txID string) {

	// Test Query Transaction -- verify that valid transaction has been processed
	processedTransaction, err := chain.QueryTransaction(txID)
	if err != nil {
		log.Fatalf("QueryTransaction return error: %v", err)
	}

	if processedTransaction.TransactionEnvelope == nil {
		log.Fatalf("QueryTransaction failed to return transaction envelope")
	}
    fmt.Println(processedTransaction.TransactionEnvelope)

    /*
	// Test Query Transaction -- Retrieve non existing transaction
	processedTransaction, err = chain.QueryTransaction("123ABC")
	if err == nil {
		log.Fatalf("QueryTransaction non-existing didn't return an error")
	}
    */

}

func testQueryBlock(chain fabricClient.Chain) {

	// Retrieve current blockchain info
	bci, err := chain.QueryInfo()
	if err != nil {
		log.Fatalf("QueryInfo return error: %v", err)
	}

	// Test Query Block by Hash - retrieve current block by hash
	block, err := chain.QueryBlockByHash(bci.CurrentBlockHash)
	if err != nil {
		log.Fatalf("QueryBlockByHash return error: %v", err)
	}

	if block.Data == nil {
		log.Fatalf("QueryBlockByHash block data is nil")
	}

	// Test Query Block by Hash - retrieve block by non-existent hash
	block, err = chain.QueryBlockByHash([]byte("non-existent"))
	if err == nil {
		log.Fatalf("QueryBlockByHash non-existent didn't return an error")
	}

	// Test Query Block - retrieve block by number
	block, err = chain.QueryBlock(1)
	if err != nil {
		log.Fatalf("QueryBlock return error: %v", err)
	}

	if block.Data == nil {
		log.Fatalf("QueryBlock block data is nil")
	}

	// Test Query Block - retrieve block by non-existent number
	block, err = chain.QueryBlock(2147483647)
	if err == nil {
		log.Fatalf("QueryBlock non-existent didn't return an error")
	}

}

func testQueryChannels(testSetup BaseSetupImpl, chain fabricClient.Chain) {

	// Our target will be primary peer on this channel
	target := chain.GetPrimaryPeer()
	fmt.Printf("****QueryChannels for %s\n", target.GetURL())
	channelQueryResponse, err := testSetup.Client.QueryChannels(target)
	if err != nil {
		log.Fatalf("QueryChannels return error: %v", err)
	}

	for _, channel := range channelQueryResponse.Channels {
		fmt.Printf("**Channel: %s\n", channel)
	}

}

func testInstalledChaincodes(testSetup BaseSetupImpl,chain fabricClient.Chain) {

	// Our target will be primary peer on this channel
	target := chain.GetPrimaryPeer()

	fmt.Printf("****QueryInstalledChaincodes for %s\n", target.GetURL())
	// Test Query Installed chaincodes for target (primary)
	chaincodeQueryResponse, err := testSetup.Client.QueryInstalledChaincodes(target)
	if err != nil {
		log.Fatalf("QueryInstalledChaincodes return error: %v", err)
	}

	for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		fmt.Printf("**InstalledCC: %s\n", chaincode)
	}

}

func testInstantiatedChaincodes(chain fabricClient.Chain) {

	// Our target will indirectly be primary peer on this channel
	target := chain.GetPrimaryPeer()

	fmt.Printf("QueryInstantiatedChaincodes for primary %s\n", target.GetURL())

	// Test Query Instantiated chaincodes
	chaincodeQueryResponse, err := chain.QueryInstantiatedChaincodes()
	if err != nil {
		log.Fatalf("QueryInstantiatedChaincodes return error: %v", err)
	}

	for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		fmt.Printf("**InstantiatedCC: %s\n", chaincode)
	}

}

func testQueryByChaincode(chain fabricClient.Chain) {

	// Test valid targets
	targets := chain.GetPeers()

	queryResponses, err := chain.QueryByChaincode("lccc", []string{"getinstalledchaincodes"}, targets)
	if err != nil {
		log.Fatalf("QueryByChaincode failed %s", err)
	}

	// Number of responses should be the same as number of targets
	if len(queryResponses) != len(targets) {
		log.Fatalf("QueryByChaincode number of results mismatch. Expected: %d Got: %d", len(targets), len(queryResponses))
	}

	// Create invalid target
	firstInvalidTarget, err := fabricClient.NewPeer("test:1111", "", "")
	if err != nil {
		log.Fatalf("Create NewPeer error(%v)", err)
	}

	// Create second invalid target
	secondInvalidTarget, err := fabricClient.NewPeer("test:2222", "", "")
	if err != nil {
		log.Fatalf("Create NewPeer error(%v)", err)
	}

	// Add invalid targets to targets
	invalidTargets := append(targets, firstInvalidTarget)
	invalidTargets = append(invalidTargets, secondInvalidTarget)

	// Add invalid targets to chain otherwise validation will fail
	chain.AddPeer(firstInvalidTarget)
	chain.AddPeer(secondInvalidTarget)

	// Test valid + invalid targets
	queryResponses, err = chain.QueryByChaincode("lccc", []string{"getinstalledchaincodes"}, invalidTargets)
	if err == nil {
		log.Fatalf("QueryByChaincode failed to return error for non-existing target")
	}

	// Verify that valid targets returned response
	if len(queryResponses) != len(targets) {
		log.Fatalf("QueryByChaincode number of results mismatch. Expected: %d Got: %d", len(targets), len(queryResponses))
	}

	chain.RemovePeer(firstInvalidTarget)
	chain.RemovePeer(secondInvalidTarget)
}
