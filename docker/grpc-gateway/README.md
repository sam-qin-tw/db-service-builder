    docker build -t samqintw/grpc-gateway .
    docker run -d -p 9090:9090 --name grpc-server samqintw/grpc-gateway 