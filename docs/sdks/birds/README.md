# Birds
(*Birds*)

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
	pb "PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: pb.Int64(179603),
                    Color: pb.String("atque"),
                    ID: pb.String("0d1ba77a-89eb-4f73-bae4-203ce5e6a95d"),
                    Name: pb.String("Donnie Abbott"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "tempora",
                },
                ID: "6ce2af7a-73cf-43be-853f-870b326b5a73",
                Name: "Norma McGlynn",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: &shared.ComplexObjectDataUpdatedDate{},
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Keith Padberg"),
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
	pb "PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateNewBird(ctx, shared.NestedBird{
        Age: &shared.NestedBirdAge{
            Amount: pb.Float64(1320.68),
            Unit: shared.NestedBirdAgeUnitMonths,
        },
        Flight: &shared.NestedBirdFlight{
            CanFly: pb.Bool(false),
            Wings: &shared.NestedBirdFlightWings{
                Count: pb.Int64(716860),
                Span: &shared.NestedBirdFlightWingsSpan{
                    Amount: pb.Float64(7044.74),
                    Unit: pb.String("aliquid"),
                },
            },
        },
        Food: []string{
            "quam",
        },
        ID: pb.String("9d232271-5bf0-4cbb-9e31-b8b90f3443a1"),
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{
                    Latitude: pb.String("quae"),
                    Longitutde: pb.String("aut"),
                },
            },
        },
        Name: pb.String("Percy Altenwerth"),
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
	pb "PB"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New(
        pb.WithSecurity(shared.Security{
            Key1: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Birds.GetAllBirds(ctx, []shared.Birds{
        shared.Birds{
            CanFly: pb.Bool(false),
            ID: pb.String("cf4b9218-79fc-4e95-bf73-ef7fbc7abd74"),
            Name: pb.String("Gilberto Dickinson"),
            WingSpan: pb.Int64(13236),
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
    res, err := s.Birds.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{
        Filter: []interface{}{
            "voluptatibus",
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

