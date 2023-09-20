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
    s := PB.New()

    ctx := context.Background()
    res, err := s.Birds.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: PB.Int64(613966),
                    Color: PB.String("dolorum"),
                    ID: PB.String("8d9cbf48-6333-423f-9b77-f3a4100674eb"),
                    Name: PB.String("Hector Mosciski"),
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
            UpdatedDate: PB.Int64(798047),
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: PB.String("Clarence Parisian"),
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
    s := PB.New()

    ctx := context.Background()
    res, err := s.Birds.CreateNewBird(ctx, shared.NestedBird{
        Age: &shared.NestedBirdAge{
            Amount: PB.Float64(6874.88),
            Unit: shared.NestedBirdAgeUnitYears,
        },
        Flight: &shared.NestedBirdFlight{
            CanFly: PB.Bool(false),
            Wings: &shared.NestedBirdFlightWings{
                Count: PB.Int64(215507),
                Span: &shared.NestedBirdFlightWingsSpan{
                    Amount: PB.Float64(7887.4),
                    Unit: PB.String("tenetur"),
                },
            },
        },
        Food: []string{
            "amet",
        },
        ID: PB.String("be453f87-0b32-46b5-a734-29cdb1a8422b"),
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{
                    Latitude: PB.String("facilis"),
                    Longitutde: PB.String("aliquid"),
                },
            },
        },
        Name: PB.String("Felicia Spencer"),
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
    s := PB.New()

    ctx := context.Background()
    res, err := s.Birds.GetAllBirds(ctx, []shared.Birds{
        shared.Birds{
            CanFly: PB.Bool(false),
            ID: PB.String("22715bf0-cbb1-4e31-b8b9-0f3443a1108e"),
            Name: PB.String("Jodi Skiles"),
            WingSpan: PB.Int64(281730),
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
    s := PB.New()

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

