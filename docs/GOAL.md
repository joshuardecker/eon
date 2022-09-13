What is the goal of this project?
----
The Luncheon network so far has gone through many design phases, but I believe I have settled on one. Here we go:

The goal of this network has multiple layers, so lets take a look 1 layer at a time. (For more information on whats happening under the hood, take a look at the Layer 0 file in this docs folder).

Layer 1?
----
Layer 1 of this network is a blockchain, but not a singular blockchain, a threaded chain of blockchains. Similar to how Kadena has there layer 1, Luncheon will have multiple blockchains at once that all thread together, which provides a variety of benefits. First, it increases ```scalability```, as more chains = more data per second. Secondly, ```extra security``` can be added, as multiple chains all dependent on each other means more chains that would have to be attacked in order to attack the network. There are other benifits, but those are the biggest 2 that I will focus on here. 

The types of data that can be stored or called on Layer 1 will only include the native token of the network ```Lunch```, and custom created tokens. Lets talk why smart contracts will not be stored / used here. Smart contracts are very powerful tools, but are bound by limitations, and I will discuss them here: