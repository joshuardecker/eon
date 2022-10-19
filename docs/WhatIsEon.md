What is Eon?
---
Eon is an engine that handles consensus between parties, and the detail we dive into that definition will be seperated in the following sections. That way those new to the project, general cryptographic consensus, or just computer science can still understand this beautiful project!

What is Eon on a nice basic level?
---
Eon is a program that allows two different parties (two different entities, whether that be you and a server, two companies, and ect.) to securely communicate data transactions (sent and received data of any kind) knowing that the data has not been messed with, currupted, or hijacked. How does it do this? Lets use the following diagram of a block:

```------------------
| Data: ...      |
|                |
| Signature: ... |
------------------```

- This structure is called a block. Why is it called that? It is a chunk of information, a literal block of information, thats it. Also note that the (...) isnt whats literally stored in the Block, it is just a placeholder for data in this diagram. Speaking of data, I will make it clear that the data section can be any type of data. Whether that be a website, payment information, chat between two servers, or anything sent via the internet. Now to ensure that this data is not currupted or tampered with, we can have the sender of the data sign the block with their signature, similar to how you can sign a contract to prove you agreed to the information signed. Same idea applies here, but the cool thing about a signature is that if the data is changed, the signature is no longer valid! That means two parties can talk to each other using this system, and not have to worry about the data failing. 

How does signing and verification work?

- I will not go into how they work on an algorithmic level, but we can talk about their application. Lets use the following diagram to ade in this explanation:

```
Alice:
Private Key -> Keep to yourself -> Makes signatures
Public Key -> Share to the World -> Verify's Signatures

Bob:
Private Key -> Keep to yourself -> Makes signatures
Public Key -> Share to the World -> Verify's Signatures
```

- As the diamgram shows, the two parties here (Alice and Bob) each have a public and private key. Private keys as the name suggests, should not be shared with others, as it is what allows you to sign data with your name on it. You can however give out your public key, as it has no ability to sign data. It can be used to verify signatures however. 

```
Alice -> (Alice) Private Key + Data -> Signature -|
                                                  |
                                                  v
Verify Signature <- (Alice) Public Key + Data <- Bob
```

- As the diagram shows, Bob only knowing the data sent by Alice and her public key is able to verify whether she actually sent it or not. Verifying the data will return false results if the data has been modified, whether by a malicious party, or a curruption of the data.

That concludes the basic level of Eon, so lets sum things up. Eon is a program that allows parties to send blocks of data with each other, and each block is signed by the sender. This signature allows for verification on whether the data has changed since the block was signed, giving all parties involved extra security when making data transactions.