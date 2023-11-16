<!-- Start SDK Example Usage -->
```go
package main

import (
	pb "PB/v2"
	"PB/v2/pkg/models/operations"
	"PB/v2/pkg/models/shared"
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