
# Boilerplate-GO

Boilerplate-GO adalah sebuah rest api sederhana yang di bangun untuk mempersingkat waktu ketika ingin membuath sebuah rest api dengan golang.

🚀 Tech Stack:

Gin Frameworks, MySQL, Docker

📃 Feature:

Auth With Jwt, Signup, Middleware, Clean Architecture

## Daftar Endpoint

### 1. Signup new user

**Endpoint:** `/memberships/signup`  
**Method:** `POST`  
**Deskripsi:** Membuat user baru.

**Request Body:**
```json
{
    "username"  : "urUsernameHere"
    "email"     : "urEmailHere@domain.com"
    "password"  : "ursecretPasswordHere"
}
```

### 2. Login user

**Endpoint:** `/memberships/login`  
**Method:** `POST`  
**Deskripsi:** Login user menggunakan username dan password.

**Request Body:**
```json
{
    "email"  : "urEmailHere@domain.com"
    "password"  : "ursecretPasswordHere"
}
```

## Jangan lupa bintang nya jika merasa terbantu 😊



