# golang-grpc
golang grpc service with client and server <br><br>

<h2>Installations :- </h2><br>
  1) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 <br>
  2) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 <br><br>

<h2>CMD to generate grpc and protobuffer file :-</h2><br>
  -> protoc --go_out=. --go-grpc_out=. .\proto\user.proto<br><br>

<h3>Running/Executions Steps :-</h3>

step 1 : Run main file of Server. Head to the server folder then execute following command<br>
		 command :- go run main.go<br><br>
		 
step 2: Wait until grpc server gets started (displayed message for server started)<br><br>

step 3: Run main file of Client. After server gets started, open another terminal with heading in client folder. Execute following command<br>
		command :- go run main.go<br><br>
		
NOTE :- keep server terminal running to execute client request. keep both the terminals running parallely.
