![](https://github.com/Part001-R/assets/blob/main/assets/netLogIWE.jpeg)

Pet project - collecting messages over the network and archiving them in a database. GPRS is used.

Server recieve data in format:
message MessageRequest{
    string typeMessage = 1; // I, W, E
    string nameProject = 2;
    string locationEvent = 3; 
    string bodyMessage = 4; 
}

If the save is successful, it returns - Ok.
message MessageResponse{
    string status = 1;
}

v0.0.1 - Basic functionality.