 # ASSIGNMENT_WEEK_5
 
 This repository contains three independent Go services, one per part of the assignment:
 
 - **Part_1**: Simple in‑memory book API using Gin by using Week 3 assignment api endpoints
 - **Part_2**: Inventory service (products + “tasks” as orders) using Gin + GORM with PostgreSQL using database created in week 4
 - **Part_3**: Enhanced inventory service with a dedicated database package and auto‑migration
 
 Each part is a separate Go module with its own `go.mod`, and can be built and run independently.
 
 ## Prerequisites
 
 - Go (version matching or higher than the `go` directive in each `go.mod`)
 - PostgreSQL running locally (for **Part_2** and **Part_3**)
   - Default expected DSN:
     - host: `localhost`
     - user: `postgres`
     - password: `strongpassword`
     - database: `inventory_db`
     - port: `5432`
 
 You can override the database connection string using the `DB_DSN` environment variable in **Part_2** and **Part_3**.
 
 ---
 
 ## Part_1 – Book API
 
 **Path**: `Part_1/`  
 **Module name**: `orchestrator-api`
 
 ### Run
 
 ```bash
 cd Part_1
 go run .
 ```
 
 The server will start on `http://localhost:8080`.
 
 **Main endpoints** (examples):
 
 - `POST /books` – add a book
 - `GET /books` – list all books
 - `GET /books/:id` – get book by ID
 - `PUT /books/:id` – update book
 - `DELETE /books/:id` – delete book
 
 Middlewares such as logging and CORS are defined in the corresponding files in `Part_1/`.
 
 ---
 
 ## Part_2 – Inventory + Tasks (Orders)
 
 **Path**: `Part_2/`  
 **Module name**: `assignment_week_5/Part_2`
 
 This service manages:
 
 - **Products** (`internal/domain/product.go`)
 - **Tasks** as **orders** (`internal/domain/order.go`)
 
 It uses:
 
 - Gin for HTTP routing (`internal/delivery/http`)
 - GORM + PostgreSQL for persistence (`internal/repository`, `cmd/server/main.go`)
 - A simple use‑case layer in `internal/usecase`
 
 ### Configure DB
 
 - By default, it connects to:
 
   ```text
   host=localhost user=postgres password=strongpassword dbname=inventory_db port=5432 sslmode=disable
   ```
 
 - To override, set `DB_DSN`:
 
   ```bash
   set DB_DSN="host=localhost user=postgres password=yourpassword dbname=inventory_db port=5432 sslmode=disable"   # Windows PowerShell
   # or
   export DB_DSN="host=localhost user=postgres password=yourpassword dbname=inventory_db port=5432 sslmode=disable" # Linux/macOS
   ```
 
 ### Run
 
 ```bash
 cd Part_2
 go run ./cmd/server
 ```
 
 The server will start on `http://localhost:8081`.
 
 **Key routes** (defined in `internal/delivery/http/router.go`):
 
 - `GET /products` – list products
 - `POST /products` – create product
 - `PUT /products/:id` – update product
 - `DELETE /products/:id` – delete product
 - `GET /tasks` – list tasks (orders)
 
 ---
 
 ## Part_3 – Inventory API with DB Package & Auto‑Migration
 
 **Path**: `Part_3/`  
 **Module name**: `assignment_week_5/Part_3`
 
 This is a more structured version of the inventory API:
 
 - Dedicated database package: `internal/database/db.go`
 - Domain models: `internal/domain`
 - Repository layer: `internal/repository`
 - Usecase layer: `internal/usecase`
 - HTTP handlers: `internal/delivery/http`
 - Auto‑migration of models on startup (`cmd/server/main.go`)
 
 ### Configure DB
 
 `internal/database/db.go` reads the DSN from `DB_DSN`, falling back to the same default as Part_2.
 
 ```bash
 set DB_DSN="host=localhost user=postgres password=yourpassword dbname=inventory_db port=5432 sslmode=disable"   # Windows PowerShell
 # or
 export DB_DSN="host=localhost user=postgres password=yourpassword dbname=inventory_db port=5432 sslmode=disable" # Linux/macOS
 ```
 
 ### Run
 
 ```bash
 cd Part_3
 go run ./cmd/server
 ```
 
 The server will start on `http://localhost:8080`.
 
 Main routes are under `/api` (see `cmd/server/main.go` and `internal/delivery/http/handler.go`), for example:
 
 - `GET /api/products`
 - `POST /api/products`
 - `PUT /api/products/:id`
 - `DELETE /api/products/:id`
 - `GET /api/tasks`
 - `POST /api/tasks`