![](https://github.com/Part001-R/assets/blob/main/assets/netLogIWE.jpeg)

Pet project - collecting messages over the network and archiving them in a database. gRPCS is used.

Server recieve data in format:
```protobuf
message MessageRequest{
    string typeMessage = 1; // I, W, E, T(test)
    string nameProject = 2;
    string locationEvent = 3; 
    string bodyMessage = 4; 
}
``````

If the save is successful, it returns - Ok.
```protobuf
message MessageResponse{
    string status = 1;
}
``````

FaultForGRPC - a project that generates messages.

v0.0.1 - Basic functionality.
