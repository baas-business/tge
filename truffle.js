/*
 * NB: since truffle-hdwallet-provider 0.0.5 you must wrap HDWallet providers in a 
 * function when declaring them. Failure to do so will cause commands to hang. ex:
 * ```
 * mainnet: {
 *     provider: function() { 
 *       return new HDWalletProvider(mnemonic, 'https://mainnet.infura.io/<infura-key>') 
 *     },
 *     network_id: '1',
 *     gas: 4500000,
 *     gasPrice: 10000000000,
 *   },
 */

module.exports = {
    networks: {
        truffle: {
            host: "127.0.0.1",
            port: 9545,
            network_id: "*"

        },
        development: {
            host: "127.0.0.1",
            port: 7545,
            network_id: "*"
        },
        ropsten: {
            host: "35.196.119.92",
            port: 8545,
            network_id: "3",
            gas: 4700036,
            from: "0x80def17bb76025eaba069de8d5ebfd8fbfc64d51"
        },
        ropstenlocal: {
            host: "localhost",
            port: 8545,
            network_id: "3",
            gas: 4700036,
            from: "0x80def17bb76025eaba069de8d5ebfd8fbfc64d51"
        },
        localdev: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "17",
            from: "0x00a329c0648769a73afac7f9381e08fb43dbea72"
        },
    }
};
