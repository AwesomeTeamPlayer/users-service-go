# user-service-go

## To build:
```bash
go build -o=/app/run ./src/
```

## To run
```bash
/app/run
```

## Available methods:

### App.AddUser
Params:
* email (string)
* name (string)
* isActive (bool)

It returns user object.

It pushes users.created event:
```json{
    "userId": 1234
}
```

### App.GetUserByEmail
Params:
* email (string)

It returns user object.

### App.GetUserById
Params:
* id (int)

It returns user object.

### App.ActiveUser
Params:
* id (int)

It returns true if user was activated. 
Otherwise returns false (when status wasn't changed).

If user was activated it pushes users.activated event:
```json{
    "userId": 1234
}
```

### App.UpdateName
Params:
* id (int)
* name (string)

It returns true.

It pushes users.name.updated event:
```json{
    "userId": 1234
}
```

### App.InactiveUser
Params:
* id (int)

It returns true if user was inactivated. 
Otherwise returns false (when status wasn't changed).


If user was inactivated it pushes users.inactivated event:
```json{
    "userId": 1234
}
```

### App.GetAllUsers
Params:
* page (int)
* limit (int)

It returns:
```json
{
  users: [
    // list of user objects sorted by id
  ],
  countAll: 123 // total number of all users
}
```


```json
{
    "id": 12,
    "jsonrpc": "2.0",
    "method": "Arith.Multiply",
    "params": [
        {
        "A": 2,
        "B": 23
        }
    ]
}
```