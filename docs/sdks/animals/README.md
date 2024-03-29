# Animals
(*Animals*)

## Overview

Work with Animals.

### Available Operations

* [CreateAnimal](#createanimal) - create an animal
* [CreateLivingThings](#createlivingthings) - create a living thing
* [DeleteAnimalsByID](#deleteanimalsbyid) - Delete Animal Object
* [GetAllAnimals](#getallanimals) - Your GET endpoint
* [GetAllLivingThings](#getalllivingthings) - Get All living things
* [GetAnimalsByID](#getanimalsbyid) - Get Animal
* [UpdateAnimalsByID](#updateanimalsbyid) - Update Animal

## CreateAnimal

Post animals description

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
    res, err := s.Animals.CreateAnimal(ctx, &operations.CreateAnimalRequestBody{
        Color: "white",
        ID: "<id>",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateAnimalRequestBody](../../pkg/models/operations/createanimalrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateAnimalResponse](../../pkg/models/operations/createanimalresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
    res, err := s.Animals.CreateLivingThings(ctx, &shared.ComplexObject{})
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

## DeleteAnimalsByID

Delete the animal

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
    res, err := s.Animals.DeleteAnimalsByID(ctx, operations.DeleteAnimalsByIDRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteAnimalsByIDRequest](../../pkg/models/operations/deleteanimalsbyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteAnimalsByIDResponse](../../pkg/models/operations/deleteanimalsbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllAnimals

Get Animals Description

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
    res, err := s.Animals.GetAllAnimals(ctx, operations.GetAllAnimalsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAllAnimalsRequest](../../pkg/models/operations/getallanimalsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetAllAnimalsResponse](../../pkg/models/operations/getallanimalsresponse.md), error**
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
    res, err := s.Animals.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{})
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

## GetAnimalsByID

Get an animal

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
    res, err := s.Animals.GetAnimalsByID(ctx, operations.GetAnimalsByIDRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAnimalsByIDRequest](../../pkg/models/operations/getanimalsbyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetAnimalsByIDResponse](../../pkg/models/operations/getanimalsbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateAnimalsByID

Update the animal object

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
    res, err := s.Animals.UpdateAnimalsByID(ctx, operations.UpdateAnimalsByIDRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Animals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAnimalsByIDRequest](../../pkg/models/operations/updateanimalsbyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateAnimalsByIDResponse](../../pkg/models/operations/updateanimalsbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
