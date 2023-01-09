const { expect } = require('chai');
const { ethers } = require('hardhat');

describe('ContaBlockContract', function () {
    it('should create an entry', async function () {
        const ContaBlockContract = await ethers.getContractFactory(
            'ContaBlockContract',
        );
        const contaBlockContract = await ContaBlockContract.deploy();
        await contaBlockContract.deployed();

        const createdEntry = await contaBlockContract.createEntry(
            'testId',
            'testAccountId',
            'testMultipleEntryId',
            'EUR',
            'testConcept',
            'testDate',
            100,
            0,
        );
        await createdEntry.wait();

        expect((await contaBlockContract.entries('testId')).id).to.equal(
            'testId',
        );
    });
});
