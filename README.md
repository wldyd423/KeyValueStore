# KeyValueStore
Distributed Systems Final Project


# Implementation

## Raft
The grpc servers will maintain consistency through RAFT algorithm. 

## Problems
1. How would client connect to grpc server?
    First we assume client already knows a server. However we don't know if this server is a leader. Normally raft would redirect clients to leader. This could be a neat solution since all we need is a middleware to redirect client to leader. However this is not interesting and there is no point of having other servers (other than just in case of a failure)

    Having one leader means all requests will be handled by one server. This is not scalable.

    Since, the assignment didn't specify we exactly use raft as it normally is implemented. We will use a datacenter like scenario. Each client will access different servers and the key value store of each server will be maintained by the leader. That said we have two major problems.

    1. How can we make sure that all servers have the same key value store?
        
        in certain cases the old key and value could be sent. This is a problem but not too severe.
        

    2. How can we write a new key value?

        This certainly requires a leader to handle this request. This indicate that a server once received a put request, it will send this packet to the leader. 

