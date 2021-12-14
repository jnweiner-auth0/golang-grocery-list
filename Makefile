gen:
		rm -rf rpc
		mkdir rpc
		protoc ./proto/* --go_out=. --go-grpc_out=. --twirp_out=.

serve:
		go run main.go