    docker build -t samqintw/grpc-server .
    
    //getway
    docker network create -d bridge --subnet 172.25.0.0/16 --gateway 172.25.0.1 grpc-gateway
    docker run -d -p 9090:9090 --name grpc_server --network=grpc_bridge samqintw/grpc-server 