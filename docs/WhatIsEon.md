What is Eon?
---
Eon is an engine that handles consensus between parties, and the detail we dive into that definition will be seperated in the following sections. That way those new to the project, general cryptographic consensus, or just computer science can still understand this beautiful project! The following sections will also build upon each other, so for a full view of the project, I would recommend reading them all in order.

What is Eon on a nice basic level?
---
Eon is a program that allows two different parties (two different entities, whether that be you and a server, two companies, and ect.) to securely communicate data transactions (sent and received data of any kind) knowing that the data has not been messed with, corrupted, or hijacked. How does it do this? Lets use the following diagram of a block:

```
------------------
| Data: ...      |
|                |
| Signature: ... |
------------------
```

- This structure is called a block. Why is it called that? It is a chunk of information, a literal block of information, thats it. Also note that the (...) isnt whats literally stored in the Block, it is just a placeholder for data in this diagram. Speaking of data, I will make it clear that the data section can be any type of data. Whether that be a website, payment information, chat between two servers, or anything sent via the internet. Now to ensure that this data is not currupted or tampered with, we can have the sender of the data sign the block with their signature, similar to how you can sign a contract to prove you agreed to the information signed. Same idea applies here, but the cool thing about a signature is that if the data is changed, the signature is no longer valid! That means two parties can talk to each other using this system, and not have to worry about the data failing. 

How does signing and verification work?

- I will not go into how they work on an algorithmic level, but we can talk about their application. Lets use the following diagram to ade in this explanation:

```
Alice:
Private Key -> Keep to yourself -> Makes signatures
Public Key -> Share to the World -> Verify's Signatures

Bob:
Private Key -> Keep to yourself -> Makes signatures
Public Key -> Share to the World -> Verifies Signatures
```

- As the diamgram shows, the two parties here (Alice and Bob) each have a public and private key. Private keys as the name suggests, should not be shared with others, as it is what allows you to sign data with your name on it. You can however give out your public key, as it has no ability to sign data. It can be used to verify signatures however. 

```
       Alice -> (Alice) Private Key + Data -> Signature
                                                  |       
                                                  v       
Verify Signature <- (Alice) Public Key + Data <- Bob
```

- As the diagram shows, Bob only knowing the data sent by Alice and her public key is able to verify whether she actually sent it or not. Verifying the data will return false results if the data has been modified, whether by a malicious party, or a curruption of the data.

That concludes the basic level of Eon, so lets sum things up. Eon is a program that allows parties to send blocks of data with each other, and each block is signed by the sender. This signature allows for verification on whether the data has changed since the block was signed, giving all parties involved extra security when making data transactions.

A Deeper Dive into Eon Consensus:
---

First, what is consensus? Consensus is an agreed state of information. Consensus is achieved when the majority agrees with the information. Lets take an example: If 5 people all go to a pizza place together, consensus is achieved when the majority (here that is 3/5 people) agree upon a pizza to split. Same idea is applied in eon, consensus is achieved when the majority of Eon Nodes (the computers running the eon software) agree upon a state of the information. 

Lets take an example in Eon. An online shopping service and a shipping company are exchanging logistical information of the orders received to the online store. The online store sorts the orders from most to least important to ship. This data is packed into blocks, and in this example lets say that the most important orders are in ```block 1```, then less important orders are sorted into ```block 2```, and finally the least important orders into ```block 3```. Consensus is achieved when both the online shop and the shipping company agree that block 1 is first, followed by block 2, and finally the least important block 3.

This is where the utility of Eon lies: Establishing consensus between parties in any situtation. 

How does Eon establish consensus between different parties?
---
This is the million dollar question and Eon has multiple solutions to this complicated problem. First, we need to define some basic rules that these consensus algorithms must follow:

- Must be Byzantine Fault Tolerant (see https://en.wikipedia.org/wiki/Byzantine_fault for more info on Byzantine Fault Tolerance).

- Must be scalable to handle many computers at once.

- Must be simplistic on a basic level to avoid complicated code hiding bugs and security vulnerabilities.

Lets talk about the three mechanisms Eon will include to solve these three requirements:

Proof of Authority:
- Proof of Authority (PoA) is the most simple of the consensus mechanisms used by Eon. All this proof is, is a trusted source signs data. If any data is signed by the trusted source, then it is viewed as true. Thats it. This mechanism is used in the first section 'What is Eon on a nice basic level?'. The Eon node operator in a config file specifies who is trusted. If the node receives data from the trusted source, it will assume that it is valid.

- The benifit of this system is that it is the most light weight of the consensus systems. Proof of authority will be configurable, where the node still checks whether the data is consistant with past data, but even with this on, PoA is still the lightest weight consensus protocol.

- PoA benifits end however when trying to be used on a public network, where the majority maintain consensus rather than a singular private entity. Thats where the next two consensus options come in handy.