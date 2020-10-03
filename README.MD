# Software Engineer Coding challenge

## How to submit your code
- Init git repository and commit initial code there (it is not needed to push it somewhere);
- Create as many commits and branches as you wish;
- Archive repository when done and send it back to us.

## Install and Run
### Requirements
1. Docker/Docker Compose
2. Git client

### How to install?

The fast way to run the service is by executing "make" target from root folder of the repository:
```
make init
```

- Check the application from a browser going to  http://localhost:8080/api/v1/products

### How to run?

```
make up
```

### Interact with the API

__Get a page of products__

```curl "http://localhost:8080/api/v1/products"```

```curl "http://localhost:8080/api/v1/products?page=2"```

__Get a product__

```curl "curl "http://localhost:8080/api/v1/product?id=8bc12dec-e2f5-11ea-b308-0242acf00a02"```

__Create a product__

```curl -X POST  -d '{"name":"LED Shoes","brand":"Niko","stock":11,"seller":"8bbf3c90-e2f5-11ea-b308-0242acf00a02"}' localhost:8080/api/v1/product```

__Update a product__

```curl -X PUT -d '{"name":"Berlin S.O.L.I.D. T-Shirt","brand":"Shirts Inc.","stock":150}' "http://localhost:8080/api/v1/product?id=156c764b-f563-11e9-94e7-38baf859afa1"```

__Delete a product__

```curl -X DELETE "http://localhost:8080/api/v1/product?id=156c826e-f563-11e9-94e7-38baf859afa1"```

__Get list of sellers__

```curl "http://localhost:8080/api/v1/sellers"```

## Tasks to DO:
### Task 1

Currently, V1 Product object in API responses has the following structure:
 ```
{
  "brand": "string",
  "name": "string",
  "seller_uuid": "string",
  "stock": 0,
  "uuid": "string"
}
```

 Based on business requirements All Product Get endpoints should be updated to contain links on seller resources. It was decided to release the second version of the API (call it V2) with the following changes:
  - V2 Product endpoints should respond with the Product object of structure:
  ```
{
  "brand": "string",
  "name": "string",
  "seller": {
     "uuid": "string",
     "_links": {
        "self":{
            "href":"http://localhost:8080/sellers/{seller_uuid}"
        }
     }
  },
  "stock": 0,
  "uuid": "string"
}
```
- V2 Product objects sent in requests(POST and PUT actions) should remain the same as in V1.
- API V1 should NOT be changed for end customer (applications that implement API V1 should never break).

### Task 2.

**Current behaviour:**

If Product Update (PUT request) changes current Stock of the products, then the seller of this product receives Email with warning.

**Behaviour to implement:**

If Product Update (PUT request) changes current Stock of the products, then the seller of this product receives Email and/or SMS with warning.
What type of notifications should be sent to seller SHOULD be defined as an application level configuration.

**IMPORTANT:** Implementation of communication with 3rd party SMS providers is not needed. Instead, please create log entry with following text template:
```
SMS Warning sent to {seller_UUID} (Phone: {seller_Phone}): {Product_name} Product stock changed
```

### Task 3

**Behaviour to implement:**
- API V2 should contain ``/api/v2/sellers/top10`` GET Endpoint.
- This endpoint should return array of maximum 10 sellers ordered by count of products they have for sale (count of entries in product table) from the largest to the smallest number.
