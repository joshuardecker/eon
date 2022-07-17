Welcome to the Luncheon Network Github!
- The following document will do the best I can to explain all of the features of Luncheon while also summerizing how it works. This document should grow abd update the closer to release the project gets.

Proof of Work and Mining:
- The blockchain will be secured via a meathod called proof of work. I will explain everything you need to know about it for this in the following: Transactions are lumped together into blocks of information. The people who validate these transactions and add them to the transaction history will be mostly miners. Mining is a term used for a person who donates computational power to the network by allowing there computer to solve complex math problems. Which ever computer finds the answer to this math problem gets rewarded for it in Luncheon (the token of this blockchain) and gets there list of mined transactions added to the transaction history. This meathod works well because if the majority are good people, the network can not be outcompeted in computational power, providing a secure blockchain.

Wallets:
- Having a wallet for Luncheon will be much simplier and easier to install than last time, even if you are not participating in the network by hosting a node or mining.

Security:
- This network will use many of the same mechanisms as bitcoin for its security, and bitcoin has been very well tested over its life span, so this blockchain should be very secure.

The Peg:
- Luncheon will no longer have a peg to real life lunches anymore, but it still will be able to be used for that. There will not be a fixed rate at which lunches can be redeemed for now, but oppertunitys will arrise for "bidding wars" to let the highest payer get a free lunch.

Lite Nodes vs Full Nodes:
- If you are interested in running your own node, there are two types that will be available. First is the Full Node. This node will have all features and make the most profit when paired with mining, but requires you to download the blockchain, which will very slowly grow over time. The lite node is what most users will have, as no downloading of blockchain is needed, but it will allow you to still do things like fetch your wallet balance, or send transactions. A full node will always provide better security and accuracy, but both get the job done.

Future Transaction Types: (LuTx's)

1: Payout Luncheon Transaction (PLuX)
- This is the transaction that the block will give as a reward to the miner who finds it. No address the coin comes from, no scripts, and a pre-defined lock time until the miner can access it. A very simple transaction.

2: Basic Luncheon Transaction (BLuX)
- This transaction is the basic transaction that will be used for most use cases. A sender uses this to send money to a receiver, and that receivers identity is unknown until they claim it to spend down the road. 

3: Advanced Luncheon Transaction (ALuX)
- This transaction type will have a very basic scripting ability, which put simply means the transaction can have extra contingencies for the transaction to be valid. Example, the transaction will only be valid if another party pays a certin amount, and expires after 24 hours if the other person does not pay, allowing for basic contracts to be built on the platform. 