# Device Readings Project
#### - Daniel Maciulla

### Building locally
```azure
go mod tidy
go run server/rest_endpoint.go
```
### Calling API's

First send some readings in (or not, you're call). I have two example files in the examples folder. 
```azure
    curl --data @examples/example.json localhost:8080/v1/reading
```

Then call by device ID for latest or the count. I recommend adding | jq . to make it pretty.
```azure
    curl localhost:8080/v1/latest/36d5658a-6908-479e-887e-a949ec199272 
    curl localhost:8080/v1/count/36d5658a-6908-479e-887e-a949ec199272 | jq .
```

### Project Structure:

- /server/ - The main code. 
- /proto/device_rest/ - Holds the proto files. You could remove the device_rest folder and refactor a tiny bit if you wanted to, but this 
way we have a nice spot to add into kafka and other schemas later.
- /examples/ - Examples for curl and generating protobufs (https://json-to-proto.github.io/)
```azure
.
├── examples
│   ├── example-2.json
│   └── example.json
├── go.mod
├── go.sum
├── proto
│   └── device_rest
│       ├── readings.pb.go
│       └── readings.proto
├── README.md
└── server
    ├── rest_endpoint.go
    └── rest_endpoint_test.go

```
### Modifying proto files

If you wanted to add or update anything to the expected request/reply modify proto/device_rest/readings.proto and then execute from the repos root path:
```azure
 protoc ./proto/device_rest/readings.proto --go_out=./proto/device_rest/
```

### If I had more time...

I time-boxed myself to ~90 minutes. I would like to add more testing. 
- Unit tests to validate the processing
- Swagger docs would be great.
- Mux or gin instead of net/http, but im more familiar with net/http.
- Integration testing as part of a build pipeline
- .gitlab-ci (or your preferred build system)
- Openetelemetry with some metrics
- Liveness/readiness probes 


## External Service Incorporation

### Kafka 

I would write the records to Kafka and then use Flink for all the processing. This way we could scale this up 
and not even need to care about the ordering in the endpoint.

### RDB

Another option would be to store the values in a database and get the results from that.

### Other note:

- Typo was spotted in the handout, ISO-8061 should be ISO-8601
- Thank you for taking the time to read this and check out my code. Best of luck in your search for your new team member!
