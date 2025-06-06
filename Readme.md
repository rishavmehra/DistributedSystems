# Distributed Systems

## Concurrency 
- [Race Condition](/Concurrency/Goroutines/main.go)
- [Synchronization](/Concurrency/)
    - [WaitGroups](/Concurrency/WaitGroups/main.go)
    - [Mutex](/Concurrency/Mutex/main.go)
    - [RWMutex](/Concurrency/RWMutex/main.go)
    - [Cond](/Concurrency/Cond/main.go)
    - [Pool](/Concurrency/Pool/main.go)
- [Deadlock, Livelock, Starvation](/Concurrency/)
    - [Deadlock](/Concurrency/Deadlock/main.go)
    - [Livelock](/Concurrency/Livelock/main.go)
    - [Starvation](/Concurrency/Starvation/main.go)
- [Problems](/Concurrency/Problems/)
    - [Producer Consumer](/Concurrency/Problems/ProducerConsumer/main.go)
    - [Confinement](/Concurrency/Problems/Confinement)
        - [ad hoc](/Concurrency/Problems/Confinement/adhoc/main.go)
        - [lexical](/Concurrency/Problems/Confinement/lexical/main.go)  
## Distributed Systems Patterns
| Type                                        | Category                                  | Patterns                                                                                                                                                                                                                                     |
| ------------------------------------------- | ----------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ❌ **Not Core Distributed Systems Patterns** | **Microservices / Cloud-Native Patterns** | **Ambassador**, **Circuit Breaker**, **Sidecar**, **Bulkhead**, **Cache-Aside**, **CQRS**, **Event Sourcing**                                                                                                                                |
| ✅ **Distributed Systems Patterns**          | **Replication**                           | Write-Ahead Log, Segmented Log, Low-Water Mark, Leader and Followers, Heartbeat, Majority Quorum, Generation Clock, High-Water Mark, Paxos, Replicated Log, Singular Update Queue, Request Waiting List, Idempotent Receiver, Follower Reads |
| ✅ **Distributed Systems Patterns**          | **Partitioning**                          | Fixed Partitions, Key-Range Partitions, Directory-Based Partitioning, Consistent Hashing, Sharded Queue, Composite Partitioning                                                                                                              |
| ✅ **Distributed Systems Patterns**          | **Coordination**                          | Single Leader, Leader and Learner, Shared Log, Two-Phase Commit, Consensus, Gossip Protocols, Membership and Failure Detection, Leader Election, Distributed Transactions and Recovery                                                       |
| ✅ **Distributed Systems Patterns**          | **Distributed Time**                      | Lamport Clock, Hybrid Clock, Clock-Bound Wait                                                                                                                                                                                                |
| ✅ **Distributed Systems Patterns**          | **Cluster Management**                    | Consistent Core, Lease, State Watch, Gossip Dissemination, Emergent Leader                                                                                                                                                                   |
| ✅ **Distributed Systems Patterns**          | **Node Communication**                    | Single-Socket Channel, Request Batch, Request Pipeline                                                                                                                                                                                       |


    
## MapReduce
Users specify a map function that processes a
key/value pair to generate a set of intermediate key/value
pairs, and a reduce function that merges all intermediate
values associated with the same intermediate key

[MapReduce Paper](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf)

![alt text](image.png)


Real-World Example: Counting Website Visits
Imagine counting daily visitors across thousands of websites:

MAP: Each worker counts visitors for a subset of websites

Worker A: "Google": 1,000,000, "Facebook": 800,000
Worker B: "Google": 900,000, "Twitter": 500,000
Worker C: "Facebook": 700,000, "Twitter": 400,000


SHUFFLE: The system organizes results by website name

"Google": [1,000,000, 900,000]
"Facebook": [800,000, 700,000]
"Twitter": [500,000, 400,000]


REDUCE: Workers sum up the values for each website

"Google": 1,900,000
"Facebook": 1,500,000