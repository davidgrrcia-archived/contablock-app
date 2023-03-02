import { HardhatUserConfig } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';

const config: HardhatUserConfig = {
    solidity: '0.8.17',
    ethernal: {
        email: process.env.ETHERNAL_EMAIL,
        password: process.env.ETHERNAL_PASSWORD,
    },
};

export default config;
