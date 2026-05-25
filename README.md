# Car Dealership Platform - Backend

A complete Go backend for a car dealership platform with JWT authentication, role-based access control, and comprehensive API endpoints.

## 📋 Features

- **User Management**: Registration, login with email/password, role-based access (user/admin)
- **Car Management**: CRUD operations, multiple images per car, descriptions
- **Orders**: Create orders, track user purchases, mark cars as sold
- **Admin Dashboard**: Full control over cars, orders, users, and CMS content
- **CMS Pages**: Editable content for home, about, and contact pages
- **Security**: JWT authentication, bcrypt password hashing, admin-only routes

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **Authentication**: JWT + bcrypt
- **Router**: Gorilla Mux
- **Dependencies**:
  - `github.com/golang-jwt/jwt/v4` - JWT token handling
  - `github.com/gorilla/mux` - HTTP routing
  - `github.com/lib/pq` - PostgreSQL driver
  - `golang.org/x/crypto` - bcrypt password hashing

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12+
- Git

### Setup Steps

1. **Clone/Extract the project**:
   ```bash
   cd car-dealership-backend
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Setup PostgreSQL Database**:
   ```bash
   # Create database
   createdb carstore

   # Run schema
   psql carstore < schema.sql
   ```

4. **Configure connection** (in `storage.go`):
   ```go
   connStr := "user=postgres dbname=carstore password=YOUR_PASSWORD sslmode=disable"
   ```

5. **Run the server**:
   ```bash
   go run main.go api.go storage.go types.go
   ```

   Server will start on `http://localhost:2020`

## 📡 API Endpoints

### Authentication (Public)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | Register new user |
| POST | `/auth/login` | Login with email/password |

### Cars (Public)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/cars` | Get all cars |
| GET | `/cars/{id}` | Get car by ID |

### Orders (Protected)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/orders` | Create order (logged-in users) |
| GET | `/orders` | Get user's orders |

### Admin Routes (Admin Only)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/admin/cars` | Create car |
| PUT | `/admin/cars/{id}` | Update car |
| DELETE | `/admin/cars/{id}` | Delete car |
| POST | `/admin/cars/{id}/images` | Add car image |
| DELETE | `/admin/cars/{id}/images/{imageId}` | Delete car image |
| GET | `/admin/orders` | Get all orders |
| DELETE | `/admin/orders/{id}` | Delete order |
| GET | `/admin/users` | Get all users |
| PUT | `/admin/users/{id}/role` | Update user role |
| DELETE | `/admin/users/{id}` | Delete user |
| GET | `/admin/cms` | Get all CMS pages |
| PUT | `/admin/cms/{slug}` | Update CMS page |

### CMS (Public)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/cms/{slug}` | Get CMS page (home, about, contact) |

## 🔐 Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer YOUR_JWT_TOKEN
```

### Register Example

```bash
curl -X POST http://localhost:2020/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

### Login Example

```bash
curl -X POST http://localhost:2020/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "firstName": "John",
    "lastName": "Doe",
    "email": "john@example.com",
    "role": "user",
    "createdAt": "2024-01-15T10:30:00Z"
  }
}
```

## 📝 Request/Response Examples

### Create Car (Admin)

```bash
curl -X POST http://localhost:2020/admin/cars \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "model": "Civic",
    "brand": "Honda",
    "year": 2023,
    "price": 25000,
    "description": "Excellent condition, low mileage"
  }'
```

### Create Order

```bash
curl -X POST http://localhost:2020/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "carId": 1,
    "firstName": "John",
    "lastName": "Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "notes": "Please deliver to my address"
  }'
```

### Add Car Image (Admin)

```bash
curl -X POST http://localhost:2020/admin/cars/1/images \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "imageUrl": "https://example.com/car-image.jpg"
  }'
```

### Update CMS Page (Admin)

```bash
curl -X PUT http://localhost:2020/admin/cms/about \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "title": "About Our Dealership",
    "content": "We are a trusted car dealership..."
  }'
```

## 📊 Database Schema

### Users Table
- `id` (SERIAL PRIMARY KEY)
- `first_name` (VARCHAR)
- `last_name` (VARCHAR)
- `email` (VARCHAR UNIQUE)
- `password` (VARCHAR - hashed)
- `role` (VARCHAR - 'user' or 'admin')
- `created_at` (TIMESTAMP)

### Cars Table
- `id` (SERIAL PRIMARY KEY)
- `model` (VARCHAR)
- `brand` (VARCHAR)
- `year` (INT)
- `price` (DECIMAL)
- `description` (TEXT)
- `is_sold` (BOOLEAN)
- `created_at` (TIMESTAMP)

### Car Images Table
- `id` (SERIAL PRIMARY KEY)
- `car_id` (INT FOREIGN KEY)
- `image_url` (VARCHAR)
- `created_at` (TIMESTAMP)

### Orders Table
- `id` (SERIAL PRIMARY KEY)
- `user_id` (INT FOREIGN KEY)
- `car_id` (INT FOREIGN KEY)
- `first_name` (VARCHAR)
- `last_name` (VARCHAR)
- `email` (VARCHAR)
- `phone` (VARCHAR)
- `notes` (TEXT)
- `total` (DECIMAL)
- `created_at` (TIMESTAMP)

### CMS Pages Table
- `id` (SERIAL PRIMARY KEY)
- `slug` (VARCHAR UNIQUE - 'home', 'about', 'contact')
- `title` (VARCHAR)
- `content` (TEXT)
- `updated_at` (TIMESTAMP)

## 🔒 Security Features

1. **Password Hashing**: All passwords are hashed using bcrypt before storage
2. **JWT Tokens**: Secure token-based authentication with 7-day expiration
3. **Role-Based Access**: Admin-only endpoints protected by role verification
4. **CORS Support**: Cross-origin requests properly handled
5. **Input Validation**: All inputs validated before processing
6. **No Password Exposure**: Passwords never included in API responses

## 🚀 Deployment

### Environment Variables

Create a `.env` file or set environment variables:

```
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=carstore
DB_HOST=localhost
DB_PORT=5432
SERVER_PORT=2020
JWT_SECRET=your_secret_key
```

### Build for Production

```bash
go build -o car-dealership-api
./car-dealership-api
```

## 📝 Notes

- JWT tokens expire after 7 days
- Cars marked as sold cannot be ordered again
- Deleting an order marks the car as available again
- Admin users can manage all content and users
- CMS pages are pre-populated with default content

## 🐛 Troubleshooting

### Database Connection Error
- Ensure PostgreSQL is running
- Check connection string in `storage.go`
- Verify database exists: `psql -l | grep carstore`

### JWT Token Invalid
- Ensure token is passed with "Bearer " prefix
- Check token hasn't expired (7 days)
- Verify JWT secret matches in code

### Admin Routes Forbidden
- Ensure user has "admin" role
- Check token is valid and contains user_id
- Verify user exists in database

## 📞 Support

For issues or questions, check the API endpoints and ensure:
1. PostgreSQL is running
2. Database schema is initialized
3. All dependencies are installed
4. Server is running on port 2020
