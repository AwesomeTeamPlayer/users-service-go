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

### App.GetUserByEmail
Params:
* email (string)

It returns user object.

### App.GetUserById
Params:
* id (string)

It returns user object.

### App.ActiveUser
Params:
* id (string)

It returns true if user was activated. 
Otherwise returns false (when status wasn't changed).
z
### App.InactiveUser
Params:
* id (string)

It returns true if user was unactivated. 
Otherwise returns false (when status wasn't changed).

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