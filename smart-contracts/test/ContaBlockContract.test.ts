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

        expect((await contaBlockContract.entryExist('testId'))).to.be.true
    });

    it('should not create a duplicated entry', async function () {
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

        await expect(contaBlockContract.createEntry(
            'testId',
            'testAccountId',
            'testMultipleEntryId',
            'EUR',
            'testConcept',
            'testDate',
            100,
            0,
        )).to.be.revertedWith(
            "Entry already exists"
        );
    });

    it('should not add a duplicated account id and multiple entry id', async function () {
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

        const createdEntry2 = await contaBlockContract.createEntry(
            'testId2',
            'testAccountId',
            'testMultipleEntryId',
            'EUR',
            'testConcept',
            'testDate',
            100,
            0,
        );
        await createdEntry2.wait();

        expect(((await contaBlockContract.listAccountIds()).length)).to.be.equal(1)
        expect(((await contaBlockContract.listMultipleEntryIds()).length)).to.be.equal(1)
    });

    it('should add a new account id and multiple entry id', async function () {
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

        const createdEntry2 = await contaBlockContract.createEntry(
            'testId2',
            'testAccountId2',
            'testMultipleEntryId2',
            'EUR',
            'testConcept',
            'testDate',
            100,
            0,
        );
        await createdEntry2.wait();

        expect(((await contaBlockContract.listAccountIds()).length)).to.be.equal(2)
        expect(((await contaBlockContract.listMultipleEntryIds()).length)).to.be.equal(2)
    });
});
