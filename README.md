## Product REST API

### Module

- MySQL Driver
- HTTP Router
- Validation

### API Doc

#### Get All Products

- Method: `GET`
- Endpoint: `/api/products`
- Response:
  ```json
  "code": "200"
  "status": "OK",
  "data": [
      {
          "id": 1,
          "name": "Beras",
          "description": "Ini adalah beras berkualitas tinggi. Dipanen langsung dari sawah tetangga.",
          "category": "bahan pokok",
          "price": 22000
      },
      {...}
    ]
  ```

#### Get Product By Id

- Method: `GET`
- Endpoint: `/api/products/:id`
- Response:
  ```json
  "code": "200"
  "status": "OK",
  "data": {
      "id": 1,
      "name": "Beras",
      "description": "Ini adalah beras berkualitas tinggi. Dipanen langsung dari sawah tetangga.",
      "category": "bahan pokok",
      "price": 22000
  }
  ```

#### Create Product

- Method: `POST`
- Endpoint: `/api/products`
- Request:
  ```json
  {
  	"name": "Beras Merah",
  	"description": "Ini adalah beras merah berkualitas tinggi. Dipanen langsung dari dusun sebelah. Harga yang tercantum adalah harga per 1 kg",
  	"category": "bahan pokok",
  	"price": 51500
  }
  ```
- Response:
  ```json
  "code": "200"
  "status": "OK",
  "data": {
      "id": 4,
      "name": "Beras Merah",
      "description": "Ini adalah beras merah berkualitas tinggi. Dipanen langsung dari dusun sebelah. Harga yang tercantum adalah harga per 1 kg",
      "category": "bahan pokok",
      "price": 51500
  }
  ```

#### Update/Edit Product

- Method: `PUT`
- Endpoint: `/api/products/:id`
- Request:
  ```json
  {
  	"name": "Beras Merah",
  	"name": "Beras Merah",
  	"description": "Ini adalah beras merah berkualitas tinggi yang sudah diedit",
  	"category": "bahan pokok",
  	"price": 54000
  }
  ```
- Response:
  ```json
  "code": "200"
  "status": "OK",
  "data": {
      "id": 4,
      "name": "Beras Merah",
      "description": "Ini adalah beras merah berkualitas tinggi yang sudah diedit",
      "category": "bahan pokok",
      "price": 54000
  }
  ```

#### Delete Product

- Method: `DELETE`
- Endpoint: `/api/products/:id`
- Response:
  ```json
  "code": "200"
  "status": "OK",
  "data": null
  ```

#### Filter Product

- Method: `GET`
- Endpoint: `/api/products?price=asc`
  - query: `price=asc|desc` & `category=perlengkapan|bahan`

#### Pagination Product

- Method: `GET`
- Endpoint: `/api/products?limit=10&offset=5`
  - query: `limit=[typIint]` & `offset=[typeInt]`
