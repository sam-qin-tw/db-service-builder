    docker build -t samqintw/grpc-gateway .
    docker run -d -p 8080:8080 --name grpc-gateway samqintw/grpc-gateway 