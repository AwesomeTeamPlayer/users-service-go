# user-service-go

## To build:
```bash
go build -o=app
chmod a+x app
```

## Prepare database:
```mysql
create table users(
  id varchar(10) not null, 
  name varchar(10) not null, 
  email varchar(255) not null, 
  is_active bool not null default false,
  PRIMARY KEY (id)
);

```

## To run built file

Required ENV variables:

* APP_PORT
* MYSQL_HOST
* MYSQL_PORT
* MYSQL_USER
* MYSQL_PASSWORD
* MYSQL_DATABASE
* RABBIT_HOST
* RABBIT_PORT
* RABBIT_USER
* RABBIT_PASSWORD
* EXCAHNGE_NAME

```bash
./app
```

## User object:
* Id (string, 10 characters)
* Email (string)
* Name (string)
* IsActive (bool)

## Available methods:

### App.AddUser
Params:
* email (string)
* name (string)
* isActive (bool)

It returns user object.

It pushes users.created event:
```json
{
    "userId": "ab123"
}
```

### App.GetUserByEmail
Params:
* email (string)

It returns user object.

### App.GetUserById
Params:
* id (string, 10 characters)

It returns user object.

### App.ActiveUser
Params:
* id (string, 10 characters)

It returns true if user was activated. 
Otherwise returns false (when status wasn't changed).

If user was activated it pushes users.activated event:
```json
{
    "userId": 1234
}
```

### App.UpdateName
Params:
* id (string, 10 characters)
* name (string)

It returns true.

It pushes users.name.updated event:
```json
{
    "userId": 1234
}
```

### App.InactiveUser
Params:
* Id (int)

It returns true if user was inactivated. 
Otherwise returns false (when status wasn't changed).


If user was inactivated it pushes users.inactivated event:
```json
{
    "userId": 1234
}
```

### App.GetAllUsers
Params:
* Page (int)
* Limit (int)

It returns:
```json
{
  users: [
    // list of user objects sorted by id
  ],
  countAll: 123 // total number of all users
}
```


## Example:

POST http://localhost:1234/rpc
```json
{
    "id": "user_____1",
    "jsonrpc": "2.0",
    "method": "App.GetAllUsers",
    "params": [
        {
        "Page": 0,
        "Limit": 20
        }
    ]
}

```