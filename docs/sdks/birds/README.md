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
                    Age: pb.Int64(24488),
                    Color: pb.String("sky blue"),
                    ID: pb.String("<ID>"),
                    Name: pb.String("loyalty Officer withdrawal"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "ruddy",
                },
                ID: "<ID>",
                Name: "Fantastic",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: &shared.ComplexObjectDataUpdatedDate{},
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Chicken"),
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
            Amount: pb.Float64(5601.46),
            Unit: shared.NestedBirdAgeUnitYears,
        },
        Flight: &shared.NestedBirdFlight{
            CanFly: pb.Bool(false),
            Wings: &shared.NestedBirdFlightWings{
                Count: pb.Int64(959530),
                Span: &shared.NestedBirdFlightWingsSpan{
                    Amount: pb.Float64(7898.44),
                    Unit: pb.String("katal"),
                },
            },
        },
        Food: []string{
            "digital",
        },
        ID: pb.String("<ID>"),
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{
                    Latitude: pb.String("-69.7312"),
                    Longitutde: pb.String("Response"),
                },
            },
        },
        Name: pb.String("wipe Southwest"),
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
            ID: pb.String("<ID>"),
            Name: pb.String("Creative"),
            WingSpan: pb.Int64(956031),
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
            "qua",
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

