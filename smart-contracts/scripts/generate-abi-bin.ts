import json from '../artifacts/contracts/ContaBlockContract.sol/ContaBlockContract.json';
import { writeFile } from 'fs/promises';

const abi = JSON.stringify(json.abi);
const bin = json.bytecode.replace('0x', '');

writeFile('./build/ContaBlockContract.abi', abi);
writeFile('./build/ContaBlockContract.bin', bin);
