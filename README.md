Transaction Types: (LuTx's)

0: Payout Luncheon Transaction (PLuX)
- This is the transaction that the block will give as a reward to the miner who finds it. No adress the coin comes from, no scripts, and a pre-defined lock time until the miner can access it. A very simple transaction.

1: Basic Luncheon Transaction (BLuX)
- This transaction is the basic transaction that will be used for most use cases. Nothing is special about it, and the only thing specifiable will be a time lock on the transaction, preventing it from being spent for a time, based on unix time or block height (block number). A real life example of where this would be useful just in general would be routine payments, like a var. This feature even at just the bank allows you to pay early for your car, but have it arrive in the dealers wallet on time, so as to not confuse it with extra payment.

2: Advanced Luncheon Transaction (ALuX)
- This transaction type will have a very basic scripting ability, which put simply means the transaction can have extra contingencies for the transaction to be valid. Example, the transaction will only be valid if another party pays a certin amount, and expires after 24 hours if the other person does not pay, allowing for basic contracts to be built on the platform. 