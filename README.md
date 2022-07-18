Welcome to the Luncheon Network Github!
- The following document will do the best I can to explain all of the features of Luncheon while also summerizing how it works. This document should grow and update the closer to release the project gets.

Stats about the currency.
- Block rewards will start out giving 200 LNCH (the currency of the network) per block.
- Each block will take on average 1 minute to mine.
- For the first year, 288,000 LNCH will be mined per day.
- Every 365 days, the block reward will half.
- 1/2 of all LNCH ever will be mined in the first year of the blockchain.
- The maximum amount of LNCH is 208,663,200. This means in terms of amount, 10 LCH is about as scarce as 1 BTC.
- The blockchains rewards can only live 7 years (Day 1 of year 8 the block reward will hit 0)

Proof of Work and Mining:
- The blockchain will be secured via a meathod called proof of work. I will explain everything you need to know about it for this in the following: Transactions are lumped together into blocks of information. The people who validate these transactions and add them to the transaction history will be mostly miners. Mining is a term used for a person who donates computational power to the network by allowing their computer to solve complex math problems. Which ever computer finds the answer to this math problem gets rewarded for it in Lunch (also called LNCH)(the token of this blockchain) and gets there list of mined transactions added to the transaction history. This meathod works well because if the majority are good people, the network can not be out-competed in computational power, providing a secure blockchain.

Wallets:
- Having a wallet for Luncheon will be much simplier and easier to install than last time, even if you are not participating in the network by hosting a node or mining.

What is a node:
- A node is simply the term for a computer who is apart of the Luncheon network. These computers will talk to each other about what is going on in the network, like when blocks are mined or new transactions are sent. The reason these are not called servers is because they arent. Any computer can be a node, and if the software existed, even your phone could be a node.

Lite Nodes vs Full Nodes:
- If you are interested in running your own node, there are two types that will be available. First is the Full Node. This node will have all features and make the most profit when paired with mining, but requires you to download the blockchain, which will very slowly grow over time. The lite node is what most users will have, as no downloading of blockchain is needed, but it will allow you to still do things like fetch your wallet balance, or send transactions. A full node will always provide better security and accuracy, but both get the job done.

Security:
- This network will use many of the same mechanisms as bitcoin for its security, and bitcoin has been very well tested over its life span, so this blockchain should be very secure.

The Peg:
- Luncheon will no longer have a value peg to real life lunches anymore, but it still will be able to be used for that. There will not be a fixed rate at which lunches can be redeemed for now, but opportunities will arise for "bidding wars" to let the highest payer get a free lunch.

Future Transaction Types: (LuTx's)

*Note the words coin and token here are used for the same meaning.

1: Payout Luncheon Transaction (PLuX)
- This is the simplest of the transaction types. This is the transaction in the block that gives the block reward tokens to the person who mined the block first. It does not come from anyone or anywhere, it just appears out of no-where, therefore adding coin to the economy that previously did not exist. There is only one of these per block, as only one person can initially get the block reward.

2: Basic Luncheon Transaction (BLuX)
- The basic Luncheon transaction is the transaction that will be used the most. It sends coin from one wallet to another. There can be many of these transactions per block, and that number is set by the max block size.

3: Advanced Luncheon Transaction (ALuX)
- This transaction is the most advanced of the three. It is similar to transaction type 2, where one person is sending tokens to another, but the difference here is that this transaction has a scripting ability with it. By that I mean with some typing, similar to coding, you can give logic to this transaction. Heres an example: You and a friend want to both pay a person for something you will buy from them, and you want to split the cost 50/50, so you set up an ALuX that will pay 50% of the cost if your friend also pays 50%, that way if they chicken out of paying, you dont lose your money and not get what you are trying to buy from the person selling whatever you were splitting with the friend 50/50. This is not the only thing the scripting will allow you to do, this is simply one example of what it can do.