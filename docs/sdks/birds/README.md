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
        pb.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateLivingThings(ctx, &shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{},
            },
            Birds: shared.CreateComplexObjectDataBirdsArrayOfany(
                    []interface{}{
                        "input",
                    },
            ),
            CreatedDate: shared.CreateComplexObjectDataCreatedDateInteger(
            248447,
            ),
            UpdatedDate: shared.CreateComplexObjectDataUpdatedDateNumber(
            6866.6,
            ),
        },
        Meta: shared.CreateComplexObjectMetaPagination(
                shared.Pagination{},
        ),
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
        pb.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateNewBird(ctx, &shared.NestedBird{
        Age: &shared.NestedBirdAge{
            Unit: shared.NestedBirdAgeUnitYears,
        },
        Flight: &shared.NestedBirdFlight{
            Wings: &shared.NestedBirdFlightWings{
                Span: &shared.NestedBirdFlightWingsSpan{},
            },
        },
        Food: []string{
            "silver",
        },
        Location: []shared.NestedBirdLocation{
            shared.NestedBirdLocation{
                Geography: &shared.NestedBirdLocationGeography{},
            },
        },
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
        pb.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Birds.GetAllBirds(ctx, &[]shared.Birds{
        shared.Birds{},
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
        pb.WithSecurity(""),
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

    if res.GetAllLivingThings200ApplicationJSONOneOf != nil {
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

