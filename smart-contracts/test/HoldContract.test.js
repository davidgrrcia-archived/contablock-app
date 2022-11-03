const HoldContract = artifacts.require("HoldContract");

let instance;
let initialSupply = 100;

beforeEach(async () => {
  instance = await HoldContract.new(initialSupply);
});

contract("HoldContract", (accounts) => {
  it("should have initial supply", async () => {
    let supply = await instance.totalSupply();
    assert(supply != null);
  });

  it("should have balance", async () => {
    let balance = await instance.balanceOf(accounts[0]);
    assert(balance > 0);
  });

  it("should have 0 initial balance on hold", async () => {
    let balanceOnHold = await instance.balanceOnHoldOf(accounts[0]);
    assert(balanceOnHold == 0);
  });

  it("hold amount should be lower than the token balance", async () => {
    try {
      await instance.hold(accounts[1], 200, 1, { from: accounts[0] });
    } catch (e) {
      return;
    }
    assert.fail();
  });

  it("hold id should be unique", async () => {
    try {
      await instance.hold(accounts[1], 10, 1, { from: accounts[0] });
      await instance.hold(accounts[1], 20, 1, { from: accounts[0] });
    } catch (e) {
      return;
    }
    assert.fail();
  });

  it("should hold correctly", async () => {
    let initial_senderBalance = await instance.balanceOf(accounts[0]);
    let initial_senderBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[0]
    );
    let initial_recipientBalance = await instance.balanceOf(accounts[1]);
    let initial_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    await instance.hold(accounts[1], 10, 1, { from: accounts[0] });

    let final_senderBalance = await instance.balanceOf(accounts[0]);
    let final_senderBalanceOnHold = await instance.balanceOnHoldOf(accounts[0]);
    let final_recipientBalance = await instance.balanceOf(accounts[1]);
    let final_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    assert(final_senderBalance - initial_senderBalance == -10);
    assert(final_senderBalanceOnHold - initial_senderBalanceOnHold == 0);
    assert(final_recipientBalance - initial_recipientBalance == 0);
    assert(final_recipientBalanceOnHold - initial_recipientBalanceOnHold == 10);
  });

  it("should execute hold correctly", async () => {
    let initial_senderBalance = await instance.balanceOf(accounts[0]);
    let initial_senderBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[0]
    );
    let initial_recipientBalance = await instance.balanceOf(accounts[1]);
    let initial_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    await instance.hold(accounts[1], 10, 1, { from: accounts[0] });
    await instance.executeHold(1, { from: accounts[0] });

    let final_senderBalance = await instance.balanceOf(accounts[0]);
    let final_senderBalanceOnHold = await instance.balanceOnHoldOf(accounts[0]);
    let final_recipientBalance = await instance.balanceOf(accounts[1]);
    let final_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    assert(final_senderBalance - initial_senderBalance == -10);
    assert(final_senderBalanceOnHold - initial_senderBalanceOnHold == 0);
    assert(final_recipientBalance - initial_recipientBalance == 10);
    assert(final_recipientBalanceOnHold - initial_recipientBalanceOnHold == 0);
  });

  it("should remove hold correctly", async () => {
    let initial_senderBalance = await instance.balanceOf(accounts[0]);
    let initial_senderBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[0]
    );
    let initial_recipientBalance = await instance.balanceOf(accounts[1]);
    let initial_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    await instance.hold(accounts[1], 10, 1, { from: accounts[0] });
    await instance.removeHold(1, { from: accounts[0] });

    let final_senderBalance = await instance.balanceOf(accounts[0]);
    let final_senderBalanceOnHold = await instance.balanceOnHoldOf(accounts[0]);
    let final_recipientBalance = await instance.balanceOf(accounts[1]);
    let final_recipientBalanceOnHold = await instance.balanceOnHoldOf(
      accounts[1]
    );

    assert(final_senderBalance - initial_senderBalance == 0);
    assert(final_senderBalanceOnHold - initial_senderBalanceOnHold == 0);
    assert(final_recipientBalance - initial_recipientBalance == 0);
    assert(final_recipientBalanceOnHold - initial_recipientBalanceOnHold == 0);
  });
});
