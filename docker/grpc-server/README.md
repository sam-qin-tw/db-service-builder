    docker build -t samqintw/grpc-server .
    docker run -d -p 9090:9090 --name grpc-server samqintw/grpc-server 