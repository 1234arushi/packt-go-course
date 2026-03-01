Implementing simple port scanning technique using goroutines and chanels

1. Spawn a pool of worker goroutines that concurrently attempt to establish TCP connections to a range of ports on the specified address.

2. If the connection is successfull,the port is considered open,otherwise it is closed

3. The results of the scan are collected,sorted and printed to the console.


TCP -> transmission control protocol : is a communication protocol not just a connection itself.While it is a protocol,TCP creates a connection oriented session between sender and receiver before sending data