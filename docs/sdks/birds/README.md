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
	"PB/v3/pkg/models/shared"
	pb "PB/v3"
	"context"
	"log"
)

func main() {
    s := pb.New(
        pb.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateLivingThings(ctx, &shared.ComplexObject{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ComplexObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.ComplexObject](../../pkg/models/shared/complexobject.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.CreateLivingThingsResponse](../../pkg/models/operations/createlivingthingsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateNewBird

Create a new Bird

### Example Usage

```go
package main

import(
	"PB/v3/pkg/models/shared"
	pb "PB/v3"
	"context"
	"log"
)

func main() {
    s := pb.New(
        pb.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Birds.CreateNewBird(ctx, &shared.NestedBird{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.NestedBird](../../pkg/models/shared/nestedbird.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateNewBirdResponse](../../pkg/models/operations/createnewbirdresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllBirds

Get All birds

### Example Usage

```go
package main

import(
	"PB/v3/pkg/models/shared"
	pb "PB/v3"
	"context"
	"log"
)

func main() {
    s := pb.New(
        pb.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Birds.GetAllBirds(ctx, []shared.Birds{
        shared.Birds{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.Birds](../../.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.GetAllBirdsResponse](../../pkg/models/operations/getallbirdsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllLivingThings

get All living things data

### Example Usage

```go
package main

import(
	"PB/v3/pkg/models/shared"
	pb "PB/v3"
	"context"
	"PB/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pb.New(
        pb.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Birds.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAllLivingThingsRequest](../../pkg/models/operations/getalllivingthingsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAllLivingThingsResponse](../../pkg/models/operations/getalllivingthingsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
