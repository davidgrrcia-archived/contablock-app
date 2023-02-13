import { ethers } from 'hardhat';

async function main() {
    const provider = new ethers.providers.JsonRpcProvider('http://hardhat-node:8545');
    const Contract = await ethers.getContractFactory('ContaBlockContract', provider);
    const contract = await Contract.deploy();

    await contract.deployed();

    console.log(`Contract deployed to ${contract.address}`);
}

main().catch(error => {
    console.error(error);
    process.exitCode = 1;
});
