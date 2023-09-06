# Birds

## Overview

Birds information.

### Available Operations

* [CreateLivingThings](#createlivingthings) - create a living thing
* [CreateNewBird](#createnewbird) - Create new Bird
* [GetAllBirds](#getallbirds) - Get Birds
* [GetAllLivingThings](#getalllivingthings) - Get All living things

## CreateLivingThings

Create a living thing

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Birds.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: pb.Int64(613966),
                    Color: pb.String("dolorum"),
                    ID: pb.String("8d9cbf48-6333-423f-9b77-f3a4100674eb"),
                    Name: pb.String("Hector Mosciski"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "sit",
                },
                ID: "d1ba77a8-9ebf-4737-ae42-03ce5e6a95d8",
                Name: "James Swaniawski",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: pb.Int64(798047),
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Clarence Parisian"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ComplexObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.ComplexObject](../../models/shared/complexobject.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.CreateLivingThingsResponse](../../models/operations/createlivingthingsresponse.md), error**


## CreateNewBird

Create a new Bird

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Birds.CreateNewBird(ctx, shared.NestedBird{
        Age: &shared.NestedBirdAge{
            Amount: pb.Float64(6874.88),
            Unit: shared.NestedBirdAgeUnitYears,
        },
        Flight: &shared.NestedBirdFlight{
            CanFly: pb.Bool(false),
            Wings: &shared.NestedBirdFlightWings{
                Count: pb.Int64(215507),
                Span: &shared.NestedBirdFlightWingsSpan{
                    Amount: pb.Float64(7887.4),
                    Unit: pb.String("tenetur"),
                },
            },
        },
        Food: []string{
            "amet",
        },
        ID: pb.String("be453f87-0b32-46b5-a734-29cdb1a8422b"),
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{
                    Latitude: pb.String("facilis"),
                    Longitutde: pb.String("aliquid"),
                },
            },
        },
        Name: pb.String("Felicia Spencer"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateNewBird200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.NestedBird](../../models/shared/nestedbird.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.CreateNewBirdResponse](../../models/operations/createnewbirdresponse.md), error**


## GetAllBirds

Get All birds

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Birds.GetAllBirds(ctx, []shared.Birds{
        shared.Birds{
            CanFly: pb.Bool(false),
            ID: pb.String("22715bf0-cbb1-4e31-b8b9-0f3443a1108e"),
            Name: pb.String("Jodi Skiles"),
            WingSpan: pb.Int64(281730),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Birds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.Birds](../../models//.md)                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.GetAllBirdsResponse](../../models/operations/getallbirdsresponse.md), error**


## GetAllLivingThings

get All living things data

### Example Usage

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
    res, err := s.Birds.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{
        Filter: []interface{}{
            "facilis",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllLivingThings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAllLivingThingsRequest](../../models/operations/getalllivingthingsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetAllLivingThingsResponse](../../models/operations/getalllivingthingsresponse.md), error**

