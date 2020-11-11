# KRIYA PEOPLE BACKEND TEST

## Installation

### On Local machine 
 - Clone this repository
 - Open the source code, open terminal and run <code> go mod tidy </code>
 - Run <code>go run main.go</code>

## Available API's
BaseUrl : localhost:12345

## Authorization for admin access
Basic Auth Authorization 

- username = <code>rudiarta</code>
- pasword = <code>$2a$14$sMIjm/rKKDADIcBcIjKViuuNrJVYCt1VeBNE6tSHxd/FSe4o53Yh2</code>

Note: use this Authorization to access API for admin

### Post User API
 - Endpoint: [BaseUrl]/users/create
 - Method  : POST
 - Authorization : Basic Auth (see on top)
 - Request Body :
  ```javascript
  {
      "username": "<string>",
      "password": "<string>",
      "email": "<string>"
  }
  ```
 - Password must be at least 6 character or more
 - Example:
 - - Request Body:
     ```javascript
        {
            "username": "testmember123",
            "password": "$2a$14$D91hJnIwC0nIIPb.AYKpjOd.u0sMnt2LglwPIZBOnx21LqmG.5ftW",
            "email": "test123@gmail.com"
        }
     ```
 - - Response Body:    
     ```javascript
        {
            "data": "User Inserted Successfully",
            "message": "OK",
            "status_code": 200
        }
     ```


### Update User API

 - Endpoint  : [BaseUrl]/users/update
 - Method    : PUT
 - Authorization : Basic Auth (see on top)
 - Request Body:
 ```javascript
{
    "uuid": "<string>",
    "email": "<string>",
    "username": "<string>",
    "password": "<string>"
}
```
 - Example:
 - - Request Body:
     ```javascript
         {
             "uuid": "e26ed126-6d41-46cf-afa6-a8db49a69edb",
             "email": "test1234@gmail.com",
             "username": "testmember123",
             "password": "$2a$14$D91hJnIwC0nIIPb.AYKpjOd.u0sMnt2LglwPIZBOnx21LqmG.5ftW"
         }
     ```
 - - Response Body: 
      ```javascript
        {
            "data": "User Updated Successfully",
            "message": "OK",
            "status_code": 200
        }
        ```
 
 ### Delete User API
 - Endpoint : [BaseUrl]/users/delete
 - Method   : DELETE
 - Authorization : Basic Auth (see on top)
 - Request Body:
  ```javascript
 {
     "uuid": "<string>"
 }
  ```
 - Example:
 - - Request Body:
     ```javascript
        {
            "uuid": "e26ed126-6d41-46cf-afa6-a8db49a69edb"
        }
     ```
 - - Response Body
     ```javascript
        {
            "data": "User Deleted Successfully",
            "message": "OK",
            "status_code": 200
        }
     ```

### Get List User API
 - Endpoint : [BaseUrl]/user/list?offset=
 - Method   : GET
 - Request URL Param:
  ```javascript
 offset (int)
  ```
 - Example:
 
 - - Response Body
     ```javascript
        {
            "data": [
                {
                    "username": "ada",
                    "email": "memberNine@gmail.com",
                    "status": {
                        "is_active": true
                    }
                },
                {
                    "username": "testMember",
                    "email": "seventhMember7@gmail.com",
                    "status": {
                        "is_active": null
                    }
                },
                {
                    "username": "testmembaer",
                    "email": "memberSix@gmail.com",
                    "status": {
                        "is_active": true
                    }
                },
                {
                    "username": "testmembaer",
                    "email": "memberFive@gmail.com",
                    "status": {
                        "is_active": true
                    }
                },
                {
                    "username": "testmembaer",
                    "email": "memberFour@gmail.com",
                    "status": {
                        "is_active": true
                    }
                }
            ],
            "message": "OK",
            "status_code": 200
        }
     ```

### Get User Detail API
 - Endpoint : [BaseUrl]/user/detail
 - Method   : GET
 - Request Body:
  ```javascript
 {
     "uuid": "<string>"
 }
  ```
 - Example:
 - - Request Body:
     ```javascript
        {
            "uuid": "dcb3d6d3-93d3-4c08-8c8c-2e95aac707ce"
        }
     ```
 - - Response Body
     ```javascript
        {
            "data": {
                "user_id": "dcb3d6d3-93d3-4c08-8c8c-2e95aac707ce",
                "username": "rudiarta",
                "email": "rudiarta@gmail.com",
                "role_name": "Admin"
            },
            "message": "OK",
            "status_code": 200
        }
     ```

     
