<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.CreateAnimal(ctx, operations.CreateAnimalRequestBody{
        Age: pb.Int64(548814),
        Color: "provident",
        ID: "bd9d8d69-a674-4e0f-867c-c8796ed151a0",
        Name: "Estelle Will",
    }, operations.CreateAnimalSecurity{
        Key1: "",
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