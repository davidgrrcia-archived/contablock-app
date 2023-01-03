import { ethers } from "hardhat";

async function main() {
  const ContaBlock = await ethers.getContractFactory("ContaBlockContract");
  console.log(ContaBlock.interface)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
