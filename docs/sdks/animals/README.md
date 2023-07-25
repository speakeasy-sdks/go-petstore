# Animals

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
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.CreateAnimal(ctx, operations.CreateAnimalRequestBody{
        Age: pb.Int64(870013),
        Color: "at",
        ID: "f7cc78ca-1ba9-428f-8816-742cb7392059",
        Name: "Sheryl Fadel",
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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAnimalRequestBody](../../models/operations/createanimalrequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateAnimalSecurity](../../models/operations/createanimalsecurity.md)       | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateAnimalResponse](../../models/operations/createanimalresponse.md), error**


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
    res, err := s.Animals.CreateLivingThings(ctx, shared.ComplexObject{
        Data: &shared.ComplexObjectData{
            Animal: []shared.Animals{
                shared.Animals{
                    Age: pb.Int64(902599),
                    Color: pb.String("fuga"),
                    ID: pb.String("7596eb10-faaa-4235-ac59-55907aff1a3a"),
                    Name: pb.String("Jaime O'Hara"),
                },
                shared.Animals{
                    Age: pb.Int64(414369),
                    Color: pb.String("quam"),
                    ID: pb.String("739251aa-52c3-4f5a-9019-da1ffe78f097"),
                    Name: pb.String("Thomas Batz"),
                },
                shared.Animals{
                    Age: pb.Int64(979587),
                    Color: pb.String("dicta"),
                    ID: pb.String("5471b5e6-e13b-499d-888e-1e91e450ad2a"),
                    Name: pb.String("Marty Green"),
                },
                shared.Animals{
                    Age: pb.Int64(397821),
                    Color: pb.String("cupiditate"),
                    ID: pb.String("802d502a-94bb-44f6-bc96-9e9a3efa77df"),
                    Name: pb.String("Keith Gulgowski"),
                },
            },
            Birds: &shared.ComplexObjectDataBirds{
                Food: []string{
                    "aliquid",
                    "laborum",
                },
                ID: "e395efb9-ba88-4f3a-a699-7074ba4469b6",
                Name: "Brandon Brakus V",
            },
            CreatedDate: &shared.ComplexObjectDataCreatedDate{},
            UpdatedDate: pb.Int64(590873),
        },
        Meta: &shared.ComplexObjectMeta{},
        Name: pb.String("Kirk Bartoletti"),
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


## DeleteAnimalsByID

Delete the animal

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
    res, err := s.Animals.DeleteAnimalsByID(ctx, operations.DeleteAnimalsByIDRequest{
        ID: "a563e251-6fe4-4c8b-b11e-5b7fd2ed0289",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteAnimalsByIDRequest](../../models/operations/deleteanimalsbyidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteAnimalsByIDResponse](../../models/operations/deleteanimalsbyidresponse.md), error**


## GetAllAnimals

Get Animals Description

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
    res, err := s.Animals.GetAllAnimals(ctx, operations.GetAllAnimalsRequest{
        Age: pb.String("magni"),
        Color: pb.String("sunt"),
        ID: pb.String("cddc6926-01fb-4576-b0d5-f0d30c5fbb25"),
        Name: pb.String("Lance Becker"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAllAnimalsRequest](../../models/operations/getallanimalsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetAllAnimalsResponse](../../models/operations/getallanimalsresponse.md), error**


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
    res, err := s.Animals.GetAllLivingThings(ctx, operations.GetAllLivingThingsRequest{
        Filter: []interface{}{
            "perferendis",
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


## GetAnimalsByID

Get an animal

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.GetAnimalsByID(ctx, operations.GetAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(170986),
            Color: pb.String("minus"),
            ID: pb.String("73d5fe9b-90c2-4890-9b3f-e49a8d9cbf48"),
            Name: pb.String("Florence Ebert"),
        },
        ID: "3f9b77f3-a410-4067-8ebf-69280d1ba77a",
        PerPage: pb.Int64(536579),
    }, operations.GetAnimalsByIDSecurity{
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

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAnimalsByIDRequest](../../models/operations/getanimalsbyidrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetAnimalsByIDSecurity](../../models/operations/getanimalsbyidsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetAnimalsByIDResponse](../../models/operations/getanimalsbyidresponse.md), error**


## UpdateAnimalsByID

Update the animal object

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PB"
	"PB/pkg/models/operations"
	"PB/pkg/models/shared"
)

func main() {
    s := pb.New()

    ctx := context.Background()
    res, err := s.Animals.UpdateAnimalsByID(ctx, operations.UpdateAnimalsByIDRequest{
        Animals: &shared.Animals{
            Age: pb.Int64(607045),
            Color: pb.String("necessitatibus"),
            ID: pb.String("bf737ae4-203c-4e5e-aa95-d8a0d446ce2a"),
            Name: pb.String("Cory Pfeffer"),
        },
        ID: "cf3be453-f870-4b32-ab5a-73429cdb1a84",
    }, operations.UpdateAnimalsByIDSecurity{
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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateAnimalsByIDRequest](../../models/operations/updateanimalsbyidrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateAnimalsByIDSecurity](../../models/operations/updateanimalsbyidsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateAnimalsByIDResponse](../../models/operations/updateanimalsbyidresponse.md), error**

