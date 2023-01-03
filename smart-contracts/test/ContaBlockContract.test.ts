const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ContaBlockContract", function () {
  it("test initial value", async function () {
    const ContaBlockContract = await ethers.getContractFactory("ContaBlockContract");
    const contaBlockContract = await ContaBlockContract.deploy();
    await contaBlockContract.deployed();
    console.log('contaBlockContract deployed at:'+ contaBlockContract.address)
    expect((await contaBlockContract.retrieve()).toNumber()).to.equal(0);
  });
  it("test updating and retrieving updated value", async function () {
    const ContaBlockContract = await ethers.getContractFactory("ContaBlockContract");
    const contaBlockContract = await ContaBlockContract.deploy();
    await contaBlockContract.deployed();
    const contaBlockContract2 = await ethers.getContractAt("ContaBlockContract", contaBlockContract.address);
    const setValue = await contaBlockContract2.store(56);
    await setValue.wait();
    expect((await contaBlockContract2.retrieve()).toNumber()).to.equal(56);
  });
});