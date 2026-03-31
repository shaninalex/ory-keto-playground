package main

import (
	"context"
	"fmt"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User 1:
//f8689777-015f-44b7-bb7d-4aa50b4956e4

// User 2:
//07768613-1074-421e-999a-4426d2d9979a

func main() {
	conn, err := grpc.NewClient("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}

	client := rts.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: "Role",
			Object:    "worker",
			Relation:  "member",
			Subject:   rts.NewSubjectID("07768613-1074-421e-999a-4426d2d9979a"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	if res.Allowed {
		fmt.Println("Allowed")
		return
	}
	fmt.Println("Denied")
}
