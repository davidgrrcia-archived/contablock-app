const HoldContract = artifacts.require("HoldContract.sol");
var Arg = 100000;

module.exports = function (deployer) {
  deployer.deploy(HoldContract, Arg);
};
