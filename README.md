# Bank accounts management system APIs:

The following APIs can be used to implement a banking system on which admins can create customers, create accounts(checking|saving), and customers can deposit or withdraw X amount of money

# Stack
<ul>
<li>Go</li>
<li>PostgreSQL </li>
</ul>

# How to run the project 
*   **Start DB**: to start the DB, run the following docker command:
  ```bash
$ docker-compose -f docker-compose-prod.yml build 
    
```

*   **Create Tables**: upload the  sql_dump file called `bank_api.sql` into the db.

*   **Start project**: to start the project , run the following command: 
  ```bash
$ go run cmd/*
```

# Customers APIs
*   **Desc**: Get All customers API 
*   **EndPoint**: int: {domain-name}/customers
*   **Method**: GET
 
*   **Response**:
  ```json
    {

    "data": [{
            "id": 0,
            "first_name": "John",
            "last_name": "Doe",
            "email": "johndoe@doe.com",
            "username": "Gondor_elf"
        },
        {
            "id": 0,
            "first_name": "Janet",
            "last_name": "Doe",
            "email": "janet@doe.com",
            "username": "valinor_elf"
        }
    ]
}
    
```

*   **Desc**: Get customer API 
*   **EndPoint**: {domain-name}/customers/{customer-id}
*   **Method**: GET
 
*   **Response**:
  ```json
    {

        "data": {
            "id": 0,
            "first_name": "Janet",
            "last_name": "Doe",
            "email": "janet@doe.com",
            "username": "valinor_elf_hello"
        }
    }
    
```


*   **Desc**: Create customer
*   **EndPoint**: {domain-name}/customers
*   **Method**: POST
*   **Payload**:
  ```json
    {
        "first_name": "John",
        "last_name": "Doe",
        "email": "johndoe@doe.com",
        "username": "dev_elf"

    }
  ```
*   **Response**:
  ```json
        {
           
            "message": "customer created"
        }
    
```


*   **Desc**: Update customer
*   **EndPoint**:  {domain-name}/customers/{customer-id}
*   **Method**: PUT
*   **Payload**:
  ```json
        {
            "first_name": "Johv",
            "last_name": "Doe",
            "email" : "johndoe@doe.com",
            "username": "Gondor_elf"

        }
  ```
*   **Response**:
  ```json
        {
           
            "message": "customer updated"
        }
    
```

# Bank Account APIs

*   **Desc**: Get all bank accounts for a given account API.
*   **EndPoint**: {domain-name}/customers/{id}/accounts
*   **Method**: GET
 *   **Response**:
  ```json
    {
        "data": [{
            "id": 1,
            "user_id": 1,
            "account_type": "checking",
            "amount": 7302
        }],

    }
    
```


*   **Desc**: Create  Bank Account API 
*   **EndPoint**: {domain-name}/customers/{id}/accounts
*   **Method**: POST
*   **Payload**:
  ```json
    {
        "account_type": "checking || saving",
        "amount": 100

    }
  ```
*   **Response**:
  ```json
        {
           
            "message": "Bank Account created"
        }
    
```


 

# Transaction API

*   **Desc**: Perform transaction API
*   **EndPoint**: {domain-name}/customers<{id}/accounts{account-id}
*   **Method**: PUT
*   **Payload**:
  ```json
    {
      "transaction_type": "withdraw || deposit",
      "age": "number",
      
    }
```
*   **Response**:
  ```json
{
    "error": false,
    "message": "Transaction completed",
    "data": {
        "id": 1,
        "user_id": 1,
        "account_type": "checking",
        "amount": 8802
    }
}
```

