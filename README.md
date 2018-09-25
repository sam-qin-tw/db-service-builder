    go generate github.com/samqintw/dbservice-factory
    
    
    # List all apis
    $ curl http://localhost:8080/swagger/echo_service.swagger.json
    
    # Visit the apis
    $ curl -XPOST http://localhost:8080/v1/example/echo/foo
    {"id":"foo"}
    
    $ curl  http://localhost:8080/v1/example/echo/foo/123
    {"id":"foo","num":"123"}