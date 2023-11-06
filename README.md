# golang-grpc
golang grpc service with client and server

<b>Installations</b> :-
  1) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  2) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

CMD to generate grpc and protobuffer file :-
  -> protoc --go_out=. --go-grpc_out=. .\proto\user.proto

Running/Executions Steps :-

step 1 : Run main file of Server. Head to the server folder then execute following command
		 command :- go run main.go
		 
step 2: Wait until grpc server gets started (displayed message for server started)

step 3: Run main file of Client. After server gets started, open another terminal with heading in client folder. Execute following command
		command :- go run main.go
		
NOTE :- keep server terminal running to execute client request. keep both the terminals running parallely.
