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
                    Age: pb.Int64(174909),
                    Color: pb.String("distinctio"),
                    ID: pb.String("b679d232-2715-4bf0-8bb1-e31b8b90f344"),
                    Name: pb.String("Mr. Sonya Bradtke"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "consequatur",
                    "est",
                    "repellendus",
                    "porro",
                },
                ID: "f4b92187-9fce-4953-b73e-f7fbc7abd74d",
                Name: "Earl Mosciski DVM",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: pb.Int64(862310),
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Johnnie Wunsch"),
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
            Amount: pb.Float64(7535.7),
            Unit: shared.NestedBirdAgeUnitYears,
        },
        Flight: &shared.NestedBirdFlight{
            CanFly: pb.Bool(false),
            Wings: &shared.NestedBirdFlightWings{
                Count: pb.Int64(4048),
                Span: &shared.NestedBirdFlightWingsSpan{
                    Amount: pb.Float64(6394.73),
                    Unit: pb.String("tempora"),
                },
            },
        },
        Food: []string{
            "ea",
            "aspernatur",
        },
        ID: pb.String("6d436813-f16d-49f5-bce6-c556146c3e25"),
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{
                    Latitude: pb.String("a"),
                    Longitutde: pb.String("libero"),
                },
            },
        },
        Name: pb.String("Jennifer Lesch"),
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
            ID: pb.String("e141aac3-66c8-4dd6-b144-2907474778a7"),
            Name: pb.String("Nicolas Graham"),
            WingSpan: pb.Int64(826871),
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
            "praesentium",
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

