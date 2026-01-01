# Data Inventaris Barang Backend

Backend Go sederhana Data Inventaris barang meggunakan JWT , CRUD & logs.

## Setup & Installation

### 1. Database Setup
1. Buka SQL.
2. Jalankan script di `app/database.sql` 

### 2. Environment Variables
Buat`.env` dengan ketentuan:

```env
DB_USERNAME=root
DB_PASSWORD=your_password
JWT_SECRET=your_super_secret_key_change_this
```

## Run Server
```
go mod tidy
go run main.go
```


## API Endpoints

### User
| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/login` | Public |
| POST | `/api/register` | Public |

### Item
| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/items` | Protected |
| GET | `/api/items` | Protected |
| GET | `/api/items/:itemId` | Protected |
| PUT | `/api/items/:itemId` | Protected |
| DELETE | `/api/items/:itemId` | Protected |

### Logs
| Method | Endpoint | Auth |
|--------|----------|------|
| GET | `/api/logs` | Protected |



