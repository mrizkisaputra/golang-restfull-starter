## Aplikasi Golang RESTfull API #
Aplikasi ini dibuat untuk melatih pemahaman praktek membuat WEB API dengan bahasa pemograman Golang.
Berikut beberapa operasi yang bisa dilakukan pada aplikasi ini, yaitu:

1. Mendapatkan semua daftar product
2. Mendapatkan detail detail product
3. Membuat product
4. Memperbaharui product
5. Menghapus product
_____
_____

1. **GET** ``http://localhost:3000/api/products``

      request header
    ```http request
    authorization: basic <credential>
    ```

   response header
    ```http response
   Content-Type: application/json
    ```
   response body
   ```response
   {
      "status": "OK",
      "code": 200,
      "message": "success",
      "data": [
         {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8799000,
            "quantity": 1
         },
         {
            "id": "product002",
            "item": "Laptop Acer Swift 3 SF314-41",
            "price": 7899000,
            "quantity": 1
         }
      ],
      "error": nil
   }
   ```
   
   response body (**error**)
   ```response
   {
      "timestamp": "03-agustus-2024 15:41:00",
      "status": "401 unauthorized",
      "message": "creadential not valid",
      "subErrors": nil
   }
   ```
___

2. **GET** ``http://localhost:3000/api/products/:id``
   
   request header
    ```http request
    authorization: basic <credential>
    ```

   response header
    ```http response
   Content-Type: application/json
    ```
   response body
   ```response
   {
      "status": "OK",
      "code": 200,
      "message": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8799000,
            "quantity": 1
      },
      "error": nil
   }
   ```

   response body (**error**) http.method (404, 401)
   ```response
   {
      "timestamp": "03-agustus-2024 15:41:00",
      "status": "404 not found",
      "message": "product is not found",
      "subErrors": nil
   }
   ```
___

3. **POST** ``http://localhost:3000/api/product``

   request header
   ```http request
   authorization: basic <credential>
   Content-Type: application/json
   ```
   
   request body
   ```request
   {
       "id": "product001",
       "item": "Laptop Acer Swift Go 14",
       "price": 8799000,
       "quantity": 1
   }
   ```

   response header
   ```http response
   Content-Type: application/json
   ```
   response body
   ```response
   {
      "status": "Created",
      "code": 201,
      "message": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8799000,
            "quantity": 1
      },
      "error": nil
   }
   ```

   response body (**error**) http.method (400, 401)
   ```response
   {
      "timestamp": "03-agustus-2024 15:41:00",
      "status": "400 bad request",
      "message": "validation error",
      "subErrors": [
         {
            "field": "item",
            "message": "item must be require"
         },
         {
            "field": "quantity",
            "message": "quantity minimum greater than 0"
         }
      ]
   }
   ```
   
___

4. **PUT** ``http://localhost:3000/api/products/:id``

   request header
    ```http request
    authorization: basic <credential>
    ```

   request body
   ```request
   {
       "item": "Laptop Acer Swift Go 14",
       "price": 8000000,
       "quantity": 1
   }
   ```
   response header
   ```http response
   Content-Type: application/json
   ```
   response body
   ```response
   {
      "status": "OK",
      "code": 200,
      "message": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8000000,
            "quantity": 1
      },
      "error": nil
   }
   ```

   response body (**error**) http.method (404, 401, 400)
   ```response
   {
      "timestamp": "03-agustus-2024 15:41:00",
      "status": "404 not found",
      "message": "id not found",
      "subErrors": nil
   }
   ```

___

5. **DELETE** ``http://localhost:3000/api/orders/:id``

   request header
    ```http request
    authorization: basic <credential>
    ```

   response body
   ```response
   {
      "status": "OK",
      "code": 200,
      "message": "success",
      "data": nil,
      "error": nil
   }