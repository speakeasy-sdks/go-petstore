<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	pb "PB/v3"
	"PB/v3/pkg/models/operations"
	"PB/v3/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := pb.New(
		pb.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
		Color: "white",
		ID:    "<id>",
		Name:  "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Animals != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->