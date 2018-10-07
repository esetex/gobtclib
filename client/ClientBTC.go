package client

import (
	"github.com/chainlibs/gobtclib/utils"
)

/*
Description:
GetBestBlockHashAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBestBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 13:18
 */
func (c *Client) GetBestBlockHashAsync() FutureGetBestBlockHashResult {
	cmd := NewCommand("getbestblockhash")
	return c.sendCmd(cmd)
}

/*
Description:
GetBestBlockHash returns the hash of the best block in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:15
 */
func (c *Client) GetBestBlockHash() (*Hash, error) {
	return c.GetBestBlockHashAsync().Receive()
}

/*
Description:
GetBlockCountAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockCount for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 12:49
 */
func (c *Client) GetBlockCountAsync() FutureGetBlockCountResult {
	cmd := NewCommand("getblockcount")
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockCount returns the number of blocks in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 12:46
 */
func (c *Client) GetBlockCount() (int64, error) {
	return c.GetBlockCountAsync().Receive()
}

/*
Description:
GetDifficultyAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetDifficulty for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 17:35
 */
func (c *Client) GetDifficultyAsync() FutureGetDifficultyResult {
	cmd := NewCommand("getdifficulty")
	return c.sendCmd(cmd)
}

/*
Description:
GetDifficulty returns the proof-of-work difficulty as a multiple of the
minimum difficulty. The result is bits/params.PowLimitBits
 * Author: architect.bian
 * Date: 2018/09/14 17:34
 */
func (c *Client) GetDifficulty() (float64, error) {
	return c.GetDifficultyAsync().Receive()
}

/*
Description:
GetBlockHashAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/15 15:35
 */
func (c *Client) GetBlockHashAsync(blockHeight int64) FutureGetBlockHashResult {
	cmd := NewCommand("getblockhash", blockHeight)
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHash returns the hash of the block in the best block chain at the given height.
 * Author: architect.bian
 * Date: 2018/09/15 15:34
 */
func (c *Client) GetBlockHash(blockHeight int64) (*Hash, error) {
	return c.GetBlockHashAsync(blockHeight).Receive()
}

/*
Description:
GetBlockHeaderVerboseAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHeader for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/15 17:51
 */
func (c *Client) GetBlockHeaderVerboseAsync(blockHash *Hash) FutureGetBlockHeaderVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblockheader", hash, utils.Basis.Bool(true))
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHeaderVerbose returns a data structure with information about the
blockheader from the server given its hash.

See GetBlockHeader to retrieve a blockheader instead.
 * Author: architect.bian
 * Date: 2018/09/15 17:51
 */
func (c *Client) GetBlockHeaderVerbose(blockHash *Hash) (*GetBlockHeaderVerboseResult, error) {
	return c.GetBlockHeaderVerboseAsync(blockHash).Receive()
}

/*
Description:
GetBlockHeaderAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHeader for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
func (c *Client) GetBlockHeaderAsync(blockHash *Hash) FutureGetBlockHeaderResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblockheader", hash, utils.Basis.Bool(false))
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHeader returns the blockheader from the server given its hash.

See GetBlockHeaderVerbose to retrieve a data structure with information about the
block instead.
 * Author: architect.bian
 * Date: 2018/09/17 15:22
 */
func (c *Client) GetBlockHeader(blockHash *Hash) (*BlockHeader, error) {
	return c.GetBlockHeaderAsync(blockHash).Receive()
}

/*
Description:
GetBlockChainInfoAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function
on the returned instance.

See GetBlockChainInfo for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 14:49
 */
func (c *Client) GetBlockChainInfoAsync() FutureGetBlockChainInfoResult {
	cmd := NewCommand("getblockchaininfo")
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockChainInfo returns information related to the processing state of
various chain-specific details such as the current difficulty from the tip
of the main chain.
 * Author: architect.bian
 * Date: 2018/09/17 14:47
 */
func (c *Client) GetBlockChainInfo() (*GetBlockChainInfoResult, error) {
	return c.GetBlockChainInfoAsync().Receive()
}

/*
Description:
GetBlockAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlock for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetBlockAsync(blockHash *Hash) FutureGetBlockResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblock", hash, utils.Basis.Int(0))
	return c.sendCmd(cmd)
}

/*
Description:
GetBlock returns a raw block from the server given its hash.

See GetBlockVerbose to retrieve a data structure with information about the
block instead.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetBlock(blockHash *Hash) (*MsgBlock, error) {
	return c.GetBlockAsync(blockHash).Receive()
}

/*
Description:
GetBlockVerboseAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBlockVerbose for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:24
 */
func (c *Client) GetBlockVerboseAsync(blockHash *Hash) FutureGetBlockVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblock", hash, utils.Basis.Int(1))
	return c.sendCmd(cmd)
}

/*
Description: 
GetBlockVerbose returns a data structure from the server with information
about a block given its hash.

See GetBlockVerboseTx to retrieve transaction data structures as well.
See GetBlock to retrieve a raw block instead.
 * Author: architect.bian
 * Date: 2018/09/17 16:21
 */
func (c *Client) GetBlockVerbose(blockHash *Hash) (*GetBlockVerboseResult, error) {
	return c.GetBlockVerboseAsync(blockHash).Receive()
}

/*
Description:
GetBlockVerboseTxAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBlockVerboseTx or the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:52
 */
func (c *Client) GetBlockVerboseTxAsync(blockHash *Hash) FutureGetBlockVerboseTXResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblock", hash, utils.Basis.Int(2))
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockVerboseTx returns a data structure from the server with information
about a block and its transactions given its hash.

See GetBlockVerbose if only transaction hashes are preferred.
See GetBlock to retrieve a raw block instead.
 * Author: architect.bian
 * Date: 2018/09/17 16:52
 */
func (c *Client) GetBlockVerboseTx(blockHash *Hash) (*GetBlockVerboseTXResult, error) {
	return c.GetBlockVerboseTxAsync(blockHash).Receive()
}

/*
Description:
GetRawMempoolAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetRawMempool for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 19:51
 */
func (c *Client) GetRawMempoolAsync() FutureGetRawMempoolResult {
	cmd := NewCommand("getrawmempool", utils.Basis.Bool(false))
	return c.sendCmd(cmd)
}

/*
Description:
GetRawMempool returns the hashes of all transactions in the memory pool.

See GetRawMempoolVerbose to retrieve data structures with information about
the transactions instead.
 * Author: architect.bian
 * Date: 2018/09/17 19:50
 */
func (c *Client) GetRawMempool() ([]*Hash, error) {
	return c.GetRawMempoolAsync().Receive()
}

/*
Description:
GetRawMempoolVerboseAsync returns an instance of a type that can be used to
get the result of the RPC at some future time by invoking the Receive
function on the returned instance.

See GetRawMempoolVerbose for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:10
 */
func (c *Client) GetRawMempoolVerboseAsync() FutureGetRawMempoolVerboseResult {
	cmd := NewCommand("getrawmempool", utils.Basis.Bool(true))
	return c.sendCmd(cmd)
}

/*
Description:
GetRawMempoolVerbose returns a map of transaction hashes to an associated
data structure with information about the transaction for all transactions in
the memory pool.

See GetRawMempool to retrieve only the transaction hashes instead.
 * Author: architect.bian
 * Date: 2018/09/17 20:09
 */
func (c *Client) GetRawMempoolVerbose() (map[string]GetRawMempoolVerboseResult, error) {
	return c.GetRawMempoolVerboseAsync().Receive()
}

/*
Description:
GetMempoolEntryAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetMempoolEntry for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:23
 */
func (c *Client) GetMempoolEntryAsync(txHash string) FutureGetMempoolEntryResult {
	cmd := NewCommand("getmempoolentry", txHash)
	return c.sendCmd(cmd)
}

/*
Description:
GetMempoolEntry returns a data structure with information about the
transaction in the memory pool given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 20:23
 */
func (c *Client) GetMempoolEntry(txHash string) (*GetMempoolEntryResult, error) { //TODO txhash changed to hash?
	return c.GetMempoolEntryAsync(txHash).Receive()
}

/*
Description:
VerifyChainAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See VerifyChain for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
func (c *Client) VerifyChainAsync() FutureVerifyChainResult {
	cmd := NewCommand("verifychain")
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChain requests the server to verify the block chain database using
the default check level and number of blocks to verify.

See VerifyChainLevel and VerifyChainBlocks to override the defaults.
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
func (c *Client) VerifyChain() (bool, error) {
	return c.VerifyChainAsync().Receive()
}

/*
Description:
VerifyChainLevelAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See VerifyChainLevel for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainLevelAsync(checkLevel int32) FutureVerifyChainResult {
	cmd := NewCommand("verifychain", &checkLevel)
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChainLevel requests the server to verify the block chain database using
the passed check level and default number of blocks to verify.

The check level controls how thorough the verification is with higher numbers
increasing the amount of checks done as consequently how long the
verification takes.

See VerifyChain to use the default check level and VerifyChainBlocks to
override the number of blocks to verify.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainLevel(checkLevel int32) (bool, error) {
	return c.VerifyChainLevelAsync(checkLevel).Receive()
}

/*
Description:
VerifyChainBlocksAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See VerifyChainBlocks for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainBlocksAsync(checkLevel, numBlocks int32) FutureVerifyChainResult {
	cmd := NewCommand("verifychain", &checkLevel, &numBlocks)
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChainBlocks requests the server to verify the block chain database
using the passed check level and number of blocks to verify.

The check level controls how thorough the verification is with higher numbers
increasing the amount of checks done as consequently how long the
verification takes.

The number of blocks refers to the number of blocks from the end of the
current longest chain.

See VerifyChain and VerifyChainLevel to use defaults.
 * Author: architect.bian
 * Date: 2018/09/17 20:59
 */
func (c *Client) VerifyChainBlocks(checkLevel, numBlocks int32) (bool, error) {
	return c.VerifyChainBlocksAsync(checkLevel, numBlocks).Receive()
}

/*
Description:
GetTxOutAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetTxOut for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:07
 */
func (c *Client) GetTxOutAsync(txHash *Hash, index uint32, mempool bool) FutureGetTxOutResult {
	hash := ""
	if txHash != nil {
		hash = txHash.String()
	}
	cmd := NewCommand("gettxout", hash, index, &mempool)
	return c.sendCmd(cmd)
}

/*
Description:
GetTxOut returns the transaction output info if it's unspent and
nil, otherwise.
 * Author: architect.bian
 * Date: 2018/09/17 21:07
 */
func (c *Client) GetTxOut(txHash *Hash, index uint32, mempool bool) (*GetTxOutResult, error) {
	return c.GetTxOutAsync(txHash, index, mempool).Receive()
}

//// FutureRescanBlocksResult is a future promise to deliver the result of a
//// RescanBlocksAsync RPC invocation (or an applicable error).
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//type FutureRescanBlocksResult chan *response
//
//// Receive waits for the response promised by the future and returns the
//// discovered rescanblocks data.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (r FutureRescanBlocksResult) Receive() ([]RescannedBlock, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	var rescanBlocksResult []RescannedBlock
//	err = json.Unmarshal(res, &rescanBlocksResult)
//	if err != nil {
//		return nil, err
//	}
//
//	return rescanBlocksResult, nil
//}
//
//// RescanBlocksCmd defines the rescan JSON-RPC command.
////
//// NOTE: This is a btcd extension ported from github.com/decred/dcrd/dcrjson
//// and requires a websocket connection.
//type RescanBlocksCmd struct {
//	// Block hashes as a string array.
//	BlockHashes []string
//}
//
//// RescanBlocksAsync returns an instance of a type that can be used to get the
//// result of the RPC at some future time by invoking the Receive function on the
//// returned instance.
////
//// See RescanBlocks for the blocking version and more details.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (c *Client) RescanBlocksAsync(blockHashes []Hash) FutureRescanBlocksResult {
//	strBlockHashes := make([]string, len(blockHashes))
//	for i := range blockHashes {
//		strBlockHashes[i] = blockHashes[i].String()
//	}
//
//	cmd := &RescanBlocksCmd{BlockHashes: strBlockHashes}
//	return c.sendCmd(cmd)
//}
//
//// RescanBlocks rescans the blocks identified by blockHashes, in order, using
//// the client's loaded transaction filter.  The blocks do not need to be on the
//// main chain, but they do need to be adjacent to each other.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (c *Client) RescanBlocks(blockHashes []Hash) ([]RescannedBlock, error) {
//	return c.RescanBlocksAsync(blockHashes).Receive()
//}

//// FutureGetCFilterResult is a future promise to deliver the result of a
//// GetCFilterAsync RPC invocation (or an applicable error).
//type FutureGetCFilterResult chan *response
//
//// Receive waits for the response promised by the future and returns the raw
//// filter requested from the server given its block hash.
//func (r FutureGetCFilterResult) Receive() (*MsgCFilter, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	// Unmarshal result as a string.
//	var filterHex string
//	err = json.Unmarshal(res, &filterHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Decode the serialized cf hex to raw bytes.
//	serializedFilter, err := hex.DecodeString(filterHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the filter bytes to the correct field of the wire message.
//	// We aren't going to set the block hash or extended flag, since we
//	// don't actually get that back in the RPC response.
//	var msgCFilter MsgCFilter
//	msgCFilter.Data = serializedFilter
//	return &msgCFilter, nil
//}
//
//// FilterType is used to represent a filter type.
//type FilterType uint8
//
//// GetCFilterCmd defines the getcfilter JSON-RPC command.
//type GetCFilterCmd struct {
//	Hash       string
//	FilterType FilterType
//}
//
//// GetCFilterAsync returns an instance of a type that can be used to get the
//// result of the RPC at some future time by invoking the Receive function on the
//// returned instance.
////
//// See GetCFilter for the blocking version and more details.
//func (c *Client) GetCFilterAsync(blockHash *Hash, filterType FilterType) FutureGetCFilterResult {
//	hash := ""
//	if blockHash != nil {
//		hash = blockHash.String()
//	}
//
//	cmd := &GetCFilterCmd{
//		Hash:       hash,
//		FilterType: filterType,
//	}
//	return c.sendCmd(cmd)
//}
//
////GetCFilter returns a raw filter from the server given its block hash.
//func (c *Client) GetCFilter(blockHash *Hash, filterType FilterType) (*MsgCFilter, error) {
//	return c.GetCFilterAsync(blockHash, filterType).Receive()
//}

////FutureGetCFilterHeaderResult is a future promise to deliver the result of a
////GetCFilterHeaderAsync RPC invocation (or an applicable error).
//type FutureGetCFilterHeaderResult chan *response
//
//// Receive waits for the response promised by the future and returns the raw
//// filter header requested from the server given its block hash.
//func (r FutureGetCFilterHeaderResult) Receive() (*MsgCFHeaders, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	// Unmarshal result as a string.
//	var headerHex string
//	err = json.Unmarshal(res, &headerHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the decoded header into a hash
//	headerHash, err := NewHashFromStr(headerHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the hash to a headers message and return it.
//	msgCFHeaders := MsgCFHeaders{PrevFilterHeader: *headerHash}
//	return &msgCFHeaders, nil
//
//}
//
//// GetCFilterHeaderCmd defines the getcfilterheader JSON-RPC command.
//type GetCFilterHeaderCmd struct {
//	Hash       string
//	FilterType FilterType
//}
//
//// GetCFilterHeaderAsync returns an instance of a type that can be used to get
//// the result of the RPC at some future time by invoking the Receive function
//// on the returned instance.
////
//// See GetCFilterHeader for the blocking version and more details.
//func (c *Client) GetCFilterHeaderAsync(blockHash *Hash,
//	filterType FilterType) FutureGetCFilterHeaderResult {
//	hash := ""
//	if blockHash != nil {
//		hash = blockHash.String()
//	}
//
//	cmd := &GetCFilterHeaderCmd{
//		Hash:       hash,
//		FilterType: filterType,
//	}
//	return c.sendCmd(cmd)
//}
//
//// GetCFilterHeader returns a raw filter header from the server given its block
//// hash.
//func (c *Client) GetCFilterHeader(blockHash *Hash,
//	filterType FilterType) (*MsgCFHeaders, error) {
//	return c.GetCFilterHeaderAsync(blockHash, filterType).Receive()
//}
