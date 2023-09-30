<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	pb "PB"
	"PB/pkg/models/shared"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Animals.CreateAnimal(ctx, operations.CreateAnimalRequestBody{
        Age: pb.Int64(239780),
        Color: "maroon",
        ID: "<ID>",
        Name: "Buckinghamshire TLS",
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