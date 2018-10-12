    // install tools
    npm install -g grpc-tools
    // if something is wrong 
    npm install -g grpc-tools --unsafe-perm
    
    // generate pb.js
    grpc_tools_node_protoc -I ../ --js_out=import_style=commonjs,binary:./ --grpc_out=./ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ./../echo-service.proto