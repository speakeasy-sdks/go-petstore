<!-- Start SDK Example Usage -->


```go
package main

import (
	pb "PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<ID>",
		Name:  "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Animals != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->