## Aplikasi Golang RESTfull API #
Aplikasi ini dibuat untuk melatih pemahaman praktek membuat WEB API dengan bahasa pemograman Golang.
Berikut beberapa operasi yang bisa dilakukan pada aplikasi ini, yaitu:

1. Mendapatkan semua daftar product
2. Mendapatkan detail product
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

   response (**success**)
   ```response
   {
      "status": "success",
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
      ]
   }
   ```
   
   response (**error**)
   ```response
   {
      "status": "error",
      "errors": null,
      "trace_id": "",
      "documentation_url": ""
   }
   ```
___

2. **GET** ``http://localhost:3000/api/products/:id``
   
   request header
    ```http request
    authorization: basic <credential>
    ```

   response (**success**)
   ```response
   {
      "status": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8799000,
            "quantity": 1
      }
   }
   ```

   response (**error**)
   ```response
   {
      "status": "error",
      "errors": null,
      "trace_id": "",
      "documentation_url": ""
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
       "item": "Laptop Acer Swift Go 14",
       "price": 8799000,
       "quantity": 1
   }
   ```

   response (**success**)
   ```response
   {
      "status": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8799000,
            "quantity": 1
      }
   }
   ```

   response (**error**)
   ```response
   {
      "status": "error",
      "error": {
         "item": ["TO_LONG"],
         "price": ["NUMBER_FORMAT", "MINIMUM 1"]
      },
      "trace_id": "",
      "documentation_url": ""
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

   response (**success**)
   ```response
   {
      "status": "success",
      "data": {
            "id": "product001",
            "item": "Laptop Acer Swift Go 14",
            "price": 8000000,
            "quantity": 1
      }
   }
   ```

   response (**error**)
   ```response
   {
      "status": "error",
      "error": null,
      "trace_id": "",
      "documentation_url": ""
   }
   ```

___

5. **DELETE** ``http://localhost:3000/api/products/:id``

   request header
    ```http request
    authorization: basic <credential>
    ```

   response (**success**)
   ```response
   {
      "status": "success",
      "data": null
   }
   ```
   
   response (**error**)
   ```response
   {
      "status": "error",
      "error": nil,
      "trace_id": "",
      "documentation_url": ""
   }
   ```