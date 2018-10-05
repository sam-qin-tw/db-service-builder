    docker build -t samqintw/grpc-gateway .
    
    //getway
    docker network create -d bridge --subnet 172.25.0.0/16 --gateway 172.25.0.1 grpc-gateway
    docker run -d -p 8080:8080 --name grpc_gateway --network=grpc_bridge samqintw/grpc-gateway --endpoint grpc_server:9090