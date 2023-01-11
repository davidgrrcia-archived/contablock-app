// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contablock

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContaBlockContractEntry is an auto generated low-level Go binding around an user-defined struct.
type ContaBlockContractEntry struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	State           string
}

// ContablockMetaData contains all meta data concerning the Contablock contract.
var ContablockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structContaBlockContract.Entry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structContaBlockContract.Entry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryPublished\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATED_STATE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLISHED_STATE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creditBalance\",\"type\":\"uint256\"}],\"name\":\"createEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"entryExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getEntryData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"}],\"internalType\":\"structContaBlockContract.Entry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listAccountIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listMultipleEntryIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"publishEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806395196b911161006657806395196b9114610120578063b4bc4e5514610150578063e521d3f11461016e578063efff528e1461018a578063f8756c01146101a857610093565b8063134760831461009857806352014343146100b657806358217d9a146100d25780636996f410146100f0575b600080fd5b6100a06101e0565b6040516100ad9190611864565b60405180910390f35b6100d060048036038101906100cb91906119cf565b6102b9565b005b6100da610976565b6040516100e79190611a62565b60405180910390f35b61010a600480360381019061010591906119cf565b6109af565b6040516101179190611b8c565b60405180910390f35b61013a600480360381019061013591906119cf565b610dfa565b6040516101479190611bc9565b60405180910390f35b610158610e30565b6040516101659190611a62565b60405180910390f35b61018860048036038101906101839190611c10565b610e69565b005b6101926111d1565b60405161019f9190611864565b60405180910390f35b6101c260048036038101906101bd91906119cf565b6112aa565b6040516101d799989796959493929190611d7d565b60405180910390f35b60606002805480602002602001604051908101604052809291908181526020016000905b828210156102b057838290600052602060002001805461022390611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461024f90611e6a565b801561029c5780601f106102715761010080835404028352916020019161029c565b820191906000526020600020905b81548152906001019060200180831161027f57829003601f168201915b505050505081526020019060010190610204565b50505050905090565b6001816040516102c99190611ed7565b908152602001604051809103902060009054906101000a900460ff16610324576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031b90611f3a565b60405180910390fd5b600080826040516103359190611ed7565b90815260200160405180910390206040518061012001604052908160008201805461035f90611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461038b90611e6a565b80156103d85780601f106103ad576101008083540402835291602001916103d8565b820191906000526020600020905b8154815290600101906020018083116103bb57829003601f168201915b505050505081526020016001820180546103f190611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461041d90611e6a565b801561046a5780601f1061043f5761010080835404028352916020019161046a565b820191906000526020600020905b81548152906001019060200180831161044d57829003601f168201915b5050505050815260200160028201805461048390611e6a565b80601f01602080910402602001604051908101604052809291908181526020018280546104af90611e6a565b80156104fc5780601f106104d1576101008083540402835291602001916104fc565b820191906000526020600020905b8154815290600101906020018083116104df57829003601f168201915b5050505050815260200160038201805461051590611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461054190611e6a565b801561058e5780601f106105635761010080835404028352916020019161058e565b820191906000526020600020905b81548152906001019060200180831161057157829003601f168201915b505050505081526020016004820180546105a790611e6a565b80601f01602080910402602001604051908101604052809291908181526020018280546105d390611e6a565b80156106205780601f106105f557610100808354040283529160200191610620565b820191906000526020600020905b81548152906001019060200180831161060357829003601f168201915b5050505050815260200160058201805461063990611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461066590611e6a565b80156106b25780601f10610687576101008083540402835291602001916106b2565b820191906000526020600020905b81548152906001019060200180831161069557829003601f168201915b5050505050815260200160068201548152602001600782015481526020016008820180546106df90611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461070b90611e6a565b80156107585780601f1061072d57610100808354040283529160200191610758565b820191906000526020600020905b81548152906001019060200180831161073b57829003601f168201915b50505050508152505090506040518060400160405280600781526020017f43524541544544000000000000000000000000000000000000000000000000008152506040516020016107a99190611ed7565b604051602081830303815290604052805190602001208161010001516040516020016107d59190611ed7565b604051602081830303815290604052805190602001201461082b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082290611fa6565b60405180910390fd5b6040518060400160405280600981526020017f5055424c495348454400000000000000000000000000000000000000000000008152508161010001819052508060008360405161087b9190611ed7565b9081526020016040518091039020600082015181600001908161089e9190612172565b5060208201518160010190816108b49190612172565b5060408201518160020190816108ca9190612172565b5060608201518160030190816108e09190612172565b5060808201518160040190816108f69190612172565b5060a082015181600501908161090c9190612172565b5060c0820151816006015560e082015181600701556101008201518160080190816109379190612172565b509050507fe0b946fe76a7369b614182784350f8116efcead2c0382480f15fa859c5609d8d8160405161096a9190611b8c565b60405180910390a15050565b6040518060400160405280600981526020017f5055424c4953484544000000000000000000000000000000000000000000000081525081565b6109b76116c6565b6000826040516109c79190611ed7565b9081526020016040518091039020604051806101200160405290816000820180546109f190611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1d90611e6a565b8015610a6a5780601f10610a3f57610100808354040283529160200191610a6a565b820191906000526020600020905b815481529060010190602001808311610a4d57829003601f168201915b50505050508152602001600182018054610a8390611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610aaf90611e6a565b8015610afc5780601f10610ad157610100808354040283529160200191610afc565b820191906000526020600020905b815481529060010190602001808311610adf57829003601f168201915b50505050508152602001600282018054610b1590611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4190611e6a565b8015610b8e5780601f10610b6357610100808354040283529160200191610b8e565b820191906000526020600020905b815481529060010190602001808311610b7157829003601f168201915b50505050508152602001600382018054610ba790611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd390611e6a565b8015610c205780601f10610bf557610100808354040283529160200191610c20565b820191906000526020600020905b815481529060010190602001808311610c0357829003601f168201915b50505050508152602001600482018054610c3990611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610c6590611e6a565b8015610cb25780601f10610c8757610100808354040283529160200191610cb2565b820191906000526020600020905b815481529060010190602001808311610c9557829003601f168201915b50505050508152602001600582018054610ccb90611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf790611e6a565b8015610d445780601f10610d1957610100808354040283529160200191610d44565b820191906000526020600020905b815481529060010190602001808311610d2757829003601f168201915b505050505081526020016006820154815260200160078201548152602001600882018054610d7190611e6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9d90611e6a565b8015610dea5780601f10610dbf57610100808354040283529160200191610dea565b820191906000526020600020905b815481529060010190602001808311610dcd57829003601f168201915b5050505050815250509050919050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b6040518060400160405280600781526020017f435245415445440000000000000000000000000000000000000000000000000081525081565b600188604051610e799190611ed7565b908152602001604051809103902060009054906101000a900460ff1615610ed5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ecc90612290565b60405180910390fd5b60006040518061012001604052808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020016040518060400160405280600781526020017f435245415445440000000000000000000000000000000000000000000000000081525081525090506001808a604051610f5f9190611ed7565b908152602001604051809103902060006101000a81548160ff0219169083151502179055508060008a604051610f959190611ed7565b90815260200160405180910390206000820151816000019081610fb89190612172565b506020820151816001019081610fce9190612172565b506040820151816002019081610fe49190612172565b506060820151816003019081610ffa9190612172565b5060808201518160040190816110109190612172565b5060a08201518160050190816110269190612172565b5060c0820151816006015560e082015181600701556101008201518160080190816110519190612172565b509050506003876040516110659190611ed7565b908152602001604051809103902060009054906101000a900460ff166110f2576002879080600181540180825580915050600190039060005260206000200160009091909190915090816110b99190612172565b5060016003886040516110cc9190611ed7565b908152602001604051809103902060006101000a81548160ff0219169083151502179055505b6005886040516111029190611ed7565b908152602001604051809103902060009054906101000a900460ff1661118f576004889080600181540180825580915050600190039060005260206000200160009091909190915090816111569190612172565b5060016005896040516111699190611ed7565b908152602001604051809103902060006101000a81548160ff0219169083151502179055505b7fa7da2603d0fde920f9f8f7f7cf2a31131ad16d2f4c353530f700748512cb895c816040516111be9190611b8c565b60405180910390a1505050505050505050565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156112a157838290600052602060002001805461121490611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461124090611e6a565b801561128d5780601f106112625761010080835404028352916020019161128d565b820191906000526020600020905b81548152906001019060200180831161127057829003601f168201915b5050505050815260200190600101906111f5565b50505050905090565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546112e390611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461130f90611e6a565b801561135c5780601f106113315761010080835404028352916020019161135c565b820191906000526020600020905b81548152906001019060200180831161133f57829003601f168201915b50505050509080600101805461137190611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461139d90611e6a565b80156113ea5780601f106113bf576101008083540402835291602001916113ea565b820191906000526020600020905b8154815290600101906020018083116113cd57829003601f168201915b5050505050908060020180546113ff90611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461142b90611e6a565b80156114785780601f1061144d57610100808354040283529160200191611478565b820191906000526020600020905b81548152906001019060200180831161145b57829003601f168201915b50505050509080600301805461148d90611e6a565b80601f01602080910402602001604051908101604052809291908181526020018280546114b990611e6a565b80156115065780601f106114db57610100808354040283529160200191611506565b820191906000526020600020905b8154815290600101906020018083116114e957829003601f168201915b50505050509080600401805461151b90611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461154790611e6a565b80156115945780601f1061156957610100808354040283529160200191611594565b820191906000526020600020905b81548152906001019060200180831161157757829003601f168201915b5050505050908060050180546115a990611e6a565b80601f01602080910402602001604051908101604052809291908181526020018280546115d590611e6a565b80156116225780601f106115f757610100808354040283529160200191611622565b820191906000526020600020905b81548152906001019060200180831161160557829003601f168201915b50505050509080600601549080600701549080600801805461164390611e6a565b80601f016020809104026020016040519081016040528092919081815260200182805461166f90611e6a565b80156116bc5780601f10611691576101008083540402835291602001916116bc565b820191906000526020600020905b81548152906001019060200180831161169f57829003601f168201915b5050505050905089565b6040518061012001604052806060815260200160608152602001606081526020016060815260200160608152602001606081526020016000815260200160008152602001606081525090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561177857808201518184015260208101905061175d565b60008484015250505050565b6000601f19601f8301169050919050565b60006117a08261173e565b6117aa8185611749565b93506117ba81856020860161175a565b6117c381611784565b840191505092915050565b60006117da8383611795565b905092915050565b6000602082019050919050565b60006117fa82611712565b611804818561171d565b9350836020820285016118168561172e565b8060005b85811015611852578484038952815161183385826117ce565b945061183e836117e2565b925060208a0199505060018101905061181a565b50829750879550505050505092915050565b6000602082019050818103600083015261187e81846117ef565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6118dc82611784565b810181811067ffffffffffffffff821117156118fb576118fa6118a4565b5b80604052505050565b600061190e611886565b905061191a82826118d3565b919050565b600067ffffffffffffffff82111561193a576119396118a4565b5b61194382611784565b9050602081019050919050565b82818337600083830152505050565b600061197261196d8461191f565b611904565b90508281526020810184848401111561198e5761198d61189f565b5b611999848285611950565b509392505050565b600082601f8301126119b6576119b561189a565b5b81356119c684826020860161195f565b91505092915050565b6000602082840312156119e5576119e4611890565b5b600082013567ffffffffffffffff811115611a0357611a02611895565b5b611a0f848285016119a1565b91505092915050565b600082825260208201905092915050565b6000611a348261173e565b611a3e8185611a18565b9350611a4e81856020860161175a565b611a5781611784565b840191505092915050565b60006020820190508181036000830152611a7c8184611a29565b905092915050565b6000819050919050565b611a9781611a84565b82525050565b6000610120830160008301518482036000860152611abb8282611795565b91505060208301518482036020860152611ad58282611795565b91505060408301518482036040860152611aef8282611795565b91505060608301518482036060860152611b098282611795565b91505060808301518482036080860152611b238282611795565b91505060a083015184820360a0860152611b3d8282611795565b91505060c0830151611b5260c0860182611a8e565b5060e0830151611b6560e0860182611a8e565b50610100830151848203610100860152611b7f8282611795565b9150508091505092915050565b60006020820190508181036000830152611ba68184611a9d565b905092915050565b60008115159050919050565b611bc381611bae565b82525050565b6000602082019050611bde6000830184611bba565b92915050565b611bed81611a84565b8114611bf857600080fd5b50565b600081359050611c0a81611be4565b92915050565b600080600080600080600080610100898b031215611c3157611c30611890565b5b600089013567ffffffffffffffff811115611c4f57611c4e611895565b5b611c5b8b828c016119a1565b985050602089013567ffffffffffffffff811115611c7c57611c7b611895565b5b611c888b828c016119a1565b975050604089013567ffffffffffffffff811115611ca957611ca8611895565b5b611cb58b828c016119a1565b965050606089013567ffffffffffffffff811115611cd657611cd5611895565b5b611ce28b828c016119a1565b955050608089013567ffffffffffffffff811115611d0357611d02611895565b5b611d0f8b828c016119a1565b94505060a089013567ffffffffffffffff811115611d3057611d2f611895565b5b611d3c8b828c016119a1565b93505060c0611d4d8b828c01611bfb565b92505060e0611d5e8b828c01611bfb565b9150509295985092959890939650565b611d7781611a84565b82525050565b6000610120820190508181036000830152611d98818c611a29565b90508181036020830152611dac818b611a29565b90508181036040830152611dc0818a611a29565b90508181036060830152611dd48189611a29565b90508181036080830152611de88188611a29565b905081810360a0830152611dfc8187611a29565b9050611e0b60c0830186611d6e565b611e1860e0830185611d6e565b818103610100830152611e2b8184611a29565b90509a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611e8257607f821691505b602082108103611e9557611e94611e3b565b5b50919050565b600081905092915050565b6000611eb18261173e565b611ebb8185611e9b565b9350611ecb81856020860161175a565b80840191505092915050565b6000611ee38284611ea6565b915081905092915050565b7f456e74727920646f6573206e6f74206578697374730000000000000000000000600082015250565b6000611f24601583611a18565b9150611f2f82611eee565b602082019050919050565b60006020820190508181036000830152611f5381611f17565b9050919050565b7f456e7472792073686f756c6420626520696e2063726561746564207374617465600082015250565b6000611f90602083611a18565b9150611f9b82611f5a565b602082019050919050565b60006020820190508181036000830152611fbf81611f83565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026120287fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611feb565b6120328683611feb565b95508019841693508086168417925050509392505050565b6000819050919050565b600061206f61206a61206584611a84565b61204a565b611a84565b9050919050565b6000819050919050565b61208983612054565b61209d61209582612076565b848454611ff8565b825550505050565b600090565b6120b26120a5565b6120bd818484612080565b505050565b5b818110156120e1576120d66000826120aa565b6001810190506120c3565b5050565b601f821115612126576120f781611fc6565b61210084611fdb565b8101602085101561210f578190505b61212361211b85611fdb565b8301826120c2565b50505b505050565b600082821c905092915050565b60006121496000198460080261212b565b1980831691505092915050565b60006121628383612138565b9150826002028217905092915050565b61217b8261173e565b67ffffffffffffffff811115612194576121936118a4565b5b61219e8254611e6a565b6121a98282856120e5565b600060209050601f8311600181146121dc57600084156121ca578287015190505b6121d48582612156565b86555061223c565b601f1984166121ea86611fc6565b60005b82811015612212578489015182556001820191506020850194506020810190506121ed565b8683101561222f578489015161222b601f891682612138565b8355505b6001600288020188555050505b505050505050565b7f456e74727920616c726561647920657869737473000000000000000000000000600082015250565b600061227a601483611a18565b915061228582612244565b602082019050919050565b600060208201905081810360008301526122a98161226d565b905091905056fea2646970667358221220b26ba9312366f03143492c0c7843501074cba282f35f42a22360e19519c0f0be64736f6c63430008110033",
}

// ContablockABI is the input ABI used to generate the binding from.
// Deprecated: Use ContablockMetaData.ABI instead.
var ContablockABI = ContablockMetaData.ABI

// ContablockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContablockMetaData.Bin instead.
var ContablockBin = ContablockMetaData.Bin

// DeployContablock deploys a new Ethereum contract, binding an instance of Contablock to it.
func DeployContablock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contablock, error) {
	parsed, err := ContablockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContablockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contablock{ContablockCaller: ContablockCaller{contract: contract}, ContablockTransactor: ContablockTransactor{contract: contract}, ContablockFilterer: ContablockFilterer{contract: contract}}, nil
}

// Contablock is an auto generated Go binding around an Ethereum contract.
type Contablock struct {
	ContablockCaller     // Read-only binding to the contract
	ContablockTransactor // Write-only binding to the contract
	ContablockFilterer   // Log filterer for contract events
}

// ContablockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContablockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContablockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContablockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContablockSession struct {
	Contract     *Contablock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContablockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContablockCallerSession struct {
	Contract *ContablockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ContablockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContablockTransactorSession struct {
	Contract     *ContablockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContablockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContablockRaw struct {
	Contract *Contablock // Generic contract binding to access the raw methods on
}

// ContablockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContablockCallerRaw struct {
	Contract *ContablockCaller // Generic read-only contract binding to access the raw methods on
}

// ContablockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContablockTransactorRaw struct {
	Contract *ContablockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContablock creates a new instance of Contablock, bound to a specific deployed contract.
func NewContablock(address common.Address, backend bind.ContractBackend) (*Contablock, error) {
	contract, err := bindContablock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contablock{ContablockCaller: ContablockCaller{contract: contract}, ContablockTransactor: ContablockTransactor{contract: contract}, ContablockFilterer: ContablockFilterer{contract: contract}}, nil
}

// NewContablockCaller creates a new read-only instance of Contablock, bound to a specific deployed contract.
func NewContablockCaller(address common.Address, caller bind.ContractCaller) (*ContablockCaller, error) {
	contract, err := bindContablock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContablockCaller{contract: contract}, nil
}

// NewContablockTransactor creates a new write-only instance of Contablock, bound to a specific deployed contract.
func NewContablockTransactor(address common.Address, transactor bind.ContractTransactor) (*ContablockTransactor, error) {
	contract, err := bindContablock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContablockTransactor{contract: contract}, nil
}

// NewContablockFilterer creates a new log filterer instance of Contablock, bound to a specific deployed contract.
func NewContablockFilterer(address common.Address, filterer bind.ContractFilterer) (*ContablockFilterer, error) {
	contract, err := bindContablock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContablockFilterer{contract: contract}, nil
}

// bindContablock binds a generic wrapper to an already deployed contract.
func bindContablock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContablockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contablock *ContablockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contablock.Contract.ContablockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contablock *ContablockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contablock.Contract.ContablockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contablock *ContablockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contablock.Contract.ContablockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contablock *ContablockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contablock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contablock *ContablockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contablock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contablock *ContablockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contablock.Contract.contract.Transact(opts, method, params...)
}

// CREATEDSTATE is a free data retrieval call binding the contract method 0xb4bc4e55.
//
// Solidity: function CREATED_STATE() view returns(string)
func (_Contablock *ContablockCaller) CREATEDSTATE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "CREATED_STATE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CREATEDSTATE is a free data retrieval call binding the contract method 0xb4bc4e55.
//
// Solidity: function CREATED_STATE() view returns(string)
func (_Contablock *ContablockSession) CREATEDSTATE() (string, error) {
	return _Contablock.Contract.CREATEDSTATE(&_Contablock.CallOpts)
}

// CREATEDSTATE is a free data retrieval call binding the contract method 0xb4bc4e55.
//
// Solidity: function CREATED_STATE() view returns(string)
func (_Contablock *ContablockCallerSession) CREATEDSTATE() (string, error) {
	return _Contablock.Contract.CREATEDSTATE(&_Contablock.CallOpts)
}

// PUBLISHEDSTATE is a free data retrieval call binding the contract method 0x58217d9a.
//
// Solidity: function PUBLISHED_STATE() view returns(string)
func (_Contablock *ContablockCaller) PUBLISHEDSTATE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "PUBLISHED_STATE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PUBLISHEDSTATE is a free data retrieval call binding the contract method 0x58217d9a.
//
// Solidity: function PUBLISHED_STATE() view returns(string)
func (_Contablock *ContablockSession) PUBLISHEDSTATE() (string, error) {
	return _Contablock.Contract.PUBLISHEDSTATE(&_Contablock.CallOpts)
}

// PUBLISHEDSTATE is a free data retrieval call binding the contract method 0x58217d9a.
//
// Solidity: function PUBLISHED_STATE() view returns(string)
func (_Contablock *ContablockCallerSession) PUBLISHEDSTATE() (string, error) {
	return _Contablock.Contract.PUBLISHEDSTATE(&_Contablock.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance, string state)
func (_Contablock *ContablockCaller) Entries(opts *bind.CallOpts, arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	State           string
}, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Id              string
		AccountId       string
		MultipleEntryId string
		Currency        string
		Concept         string
		Date            string
		DebitBalance    *big.Int
		CreditBalance   *big.Int
		State           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AccountId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MultipleEntryId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Currency = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Concept = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.DebitBalance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CreditBalance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[8], new(string)).(*string)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance, string state)
func (_Contablock *ContablockSession) Entries(arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	State           string
}, error) {
	return _Contablock.Contract.Entries(&_Contablock.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance, string state)
func (_Contablock *ContablockCallerSession) Entries(arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	State           string
}, error) {
	return _Contablock.Contract.Entries(&_Contablock.CallOpts, arg0)
}

// EntryExist is a free data retrieval call binding the contract method 0x95196b91.
//
// Solidity: function entryExist(string ) view returns(bool)
func (_Contablock *ContablockCaller) EntryExist(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "entryExist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EntryExist is a free data retrieval call binding the contract method 0x95196b91.
//
// Solidity: function entryExist(string ) view returns(bool)
func (_Contablock *ContablockSession) EntryExist(arg0 string) (bool, error) {
	return _Contablock.Contract.EntryExist(&_Contablock.CallOpts, arg0)
}

// EntryExist is a free data retrieval call binding the contract method 0x95196b91.
//
// Solidity: function entryExist(string ) view returns(bool)
func (_Contablock *ContablockCallerSession) EntryExist(arg0 string) (bool, error) {
	return _Contablock.Contract.EntryExist(&_Contablock.CallOpts, arg0)
}

// GetEntryData is a free data retrieval call binding the contract method 0x6996f410.
//
// Solidity: function getEntryData(string _id) view returns((string,string,string,string,string,string,uint256,uint256,string))
func (_Contablock *ContablockCaller) GetEntryData(opts *bind.CallOpts, _id string) (ContaBlockContractEntry, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "getEntryData", _id)

	if err != nil {
		return *new(ContaBlockContractEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(ContaBlockContractEntry)).(*ContaBlockContractEntry)

	return out0, err

}

// GetEntryData is a free data retrieval call binding the contract method 0x6996f410.
//
// Solidity: function getEntryData(string _id) view returns((string,string,string,string,string,string,uint256,uint256,string))
func (_Contablock *ContablockSession) GetEntryData(_id string) (ContaBlockContractEntry, error) {
	return _Contablock.Contract.GetEntryData(&_Contablock.CallOpts, _id)
}

// GetEntryData is a free data retrieval call binding the contract method 0x6996f410.
//
// Solidity: function getEntryData(string _id) view returns((string,string,string,string,string,string,uint256,uint256,string))
func (_Contablock *ContablockCallerSession) GetEntryData(_id string) (ContaBlockContractEntry, error) {
	return _Contablock.Contract.GetEntryData(&_Contablock.CallOpts, _id)
}

// ListAccountIds is a free data retrieval call binding the contract method 0xefff528e.
//
// Solidity: function listAccountIds() view returns(string[])
func (_Contablock *ContablockCaller) ListAccountIds(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "listAccountIds")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ListAccountIds is a free data retrieval call binding the contract method 0xefff528e.
//
// Solidity: function listAccountIds() view returns(string[])
func (_Contablock *ContablockSession) ListAccountIds() ([]string, error) {
	return _Contablock.Contract.ListAccountIds(&_Contablock.CallOpts)
}

// ListAccountIds is a free data retrieval call binding the contract method 0xefff528e.
//
// Solidity: function listAccountIds() view returns(string[])
func (_Contablock *ContablockCallerSession) ListAccountIds() ([]string, error) {
	return _Contablock.Contract.ListAccountIds(&_Contablock.CallOpts)
}

// ListMultipleEntryIds is a free data retrieval call binding the contract method 0x13476083.
//
// Solidity: function listMultipleEntryIds() view returns(string[])
func (_Contablock *ContablockCaller) ListMultipleEntryIds(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "listMultipleEntryIds")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ListMultipleEntryIds is a free data retrieval call binding the contract method 0x13476083.
//
// Solidity: function listMultipleEntryIds() view returns(string[])
func (_Contablock *ContablockSession) ListMultipleEntryIds() ([]string, error) {
	return _Contablock.Contract.ListMultipleEntryIds(&_Contablock.CallOpts)
}

// ListMultipleEntryIds is a free data retrieval call binding the contract method 0x13476083.
//
// Solidity: function listMultipleEntryIds() view returns(string[])
func (_Contablock *ContablockCallerSession) ListMultipleEntryIds() ([]string, error) {
	return _Contablock.Contract.ListMultipleEntryIds(&_Contablock.CallOpts)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockTransactor) CreateEntry(opts *bind.TransactOpts, _id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.contract.Transact(opts, "createEntry", _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockSession) CreateEntry(_id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.CreateEntry(&_Contablock.TransactOpts, _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockTransactorSession) CreateEntry(_id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.CreateEntry(&_Contablock.TransactOpts, _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// PublishEntry is a paid mutator transaction binding the contract method 0x52014343.
//
// Solidity: function publishEntry(string _id) returns()
func (_Contablock *ContablockTransactor) PublishEntry(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Contablock.contract.Transact(opts, "publishEntry", _id)
}

// PublishEntry is a paid mutator transaction binding the contract method 0x52014343.
//
// Solidity: function publishEntry(string _id) returns()
func (_Contablock *ContablockSession) PublishEntry(_id string) (*types.Transaction, error) {
	return _Contablock.Contract.PublishEntry(&_Contablock.TransactOpts, _id)
}

// PublishEntry is a paid mutator transaction binding the contract method 0x52014343.
//
// Solidity: function publishEntry(string _id) returns()
func (_Contablock *ContablockTransactorSession) PublishEntry(_id string) (*types.Transaction, error) {
	return _Contablock.Contract.PublishEntry(&_Contablock.TransactOpts, _id)
}

// ContablockEntryCreatedIterator is returned from FilterEntryCreated and is used to iterate over the raw logs and unpacked data for EntryCreated events raised by the Contablock contract.
type ContablockEntryCreatedIterator struct {
	Event *ContablockEntryCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContablockEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContablockEntryCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContablockEntryCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContablockEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContablockEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContablockEntryCreated represents a EntryCreated event raised by the Contablock contract.
type ContablockEntryCreated struct {
	Entry ContaBlockContractEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryCreated is a free log retrieval operation binding the contract event 0xa7da2603d0fde920f9f8f7f7cf2a31131ad16d2f4c353530f700748512cb895c.
//
// Solidity: event EntryCreated((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) FilterEntryCreated(opts *bind.FilterOpts) (*ContablockEntryCreatedIterator, error) {

	logs, sub, err := _Contablock.contract.FilterLogs(opts, "EntryCreated")
	if err != nil {
		return nil, err
	}
	return &ContablockEntryCreatedIterator{contract: _Contablock.contract, event: "EntryCreated", logs: logs, sub: sub}, nil
}

// WatchEntryCreated is a free log subscription operation binding the contract event 0xa7da2603d0fde920f9f8f7f7cf2a31131ad16d2f4c353530f700748512cb895c.
//
// Solidity: event EntryCreated((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) WatchEntryCreated(opts *bind.WatchOpts, sink chan<- *ContablockEntryCreated) (event.Subscription, error) {

	logs, sub, err := _Contablock.contract.WatchLogs(opts, "EntryCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContablockEntryCreated)
				if err := _Contablock.contract.UnpackLog(event, "EntryCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEntryCreated is a log parse operation binding the contract event 0xa7da2603d0fde920f9f8f7f7cf2a31131ad16d2f4c353530f700748512cb895c.
//
// Solidity: event EntryCreated((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) ParseEntryCreated(log types.Log) (*ContablockEntryCreated, error) {
	event := new(ContablockEntryCreated)
	if err := _Contablock.contract.UnpackLog(event, "EntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContablockEntryPublishedIterator is returned from FilterEntryPublished and is used to iterate over the raw logs and unpacked data for EntryPublished events raised by the Contablock contract.
type ContablockEntryPublishedIterator struct {
	Event *ContablockEntryPublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContablockEntryPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContablockEntryPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContablockEntryPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContablockEntryPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContablockEntryPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContablockEntryPublished represents a EntryPublished event raised by the Contablock contract.
type ContablockEntryPublished struct {
	Entry ContaBlockContractEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryPublished is a free log retrieval operation binding the contract event 0xe0b946fe76a7369b614182784350f8116efcead2c0382480f15fa859c5609d8d.
//
// Solidity: event EntryPublished((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) FilterEntryPublished(opts *bind.FilterOpts) (*ContablockEntryPublishedIterator, error) {

	logs, sub, err := _Contablock.contract.FilterLogs(opts, "EntryPublished")
	if err != nil {
		return nil, err
	}
	return &ContablockEntryPublishedIterator{contract: _Contablock.contract, event: "EntryPublished", logs: logs, sub: sub}, nil
}

// WatchEntryPublished is a free log subscription operation binding the contract event 0xe0b946fe76a7369b614182784350f8116efcead2c0382480f15fa859c5609d8d.
//
// Solidity: event EntryPublished((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) WatchEntryPublished(opts *bind.WatchOpts, sink chan<- *ContablockEntryPublished) (event.Subscription, error) {

	logs, sub, err := _Contablock.contract.WatchLogs(opts, "EntryPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContablockEntryPublished)
				if err := _Contablock.contract.UnpackLog(event, "EntryPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEntryPublished is a log parse operation binding the contract event 0xe0b946fe76a7369b614182784350f8116efcead2c0382480f15fa859c5609d8d.
//
// Solidity: event EntryPublished((string,string,string,string,string,string,uint256,uint256,string) entry)
func (_Contablock *ContablockFilterer) ParseEntryPublished(log types.Log) (*ContablockEntryPublished, error) {
	event := new(ContablockEntryPublished)
	if err := _Contablock.contract.UnpackLog(event, "EntryPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
