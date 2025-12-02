# New Features Guide

This document provides a quick reference for the newly added features: JWT Authentication, Swagger Documentation, and Rate Limiting.

## JWT Authentication

### Endpoints

#### Register a New User
```bash
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "password": "secure_password123",
  "email": "john@example.com"
}
```

**Response:**
```json
{
  "result": {"user_id": 1},
  "success": true,
  "resultCode": 0
}
```

#### Login
```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "secure_password123"
}
```

**Response:**
```json
{
  "result": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "success": true,
  "resultCode": 0
}
```

#### Refresh Token
```bash
POST /api/v1/auth/refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Using Protected Endpoints

All CRUD endpoints now require authentication. Include the access token in the Authorization header:

```bash
GET /api/v1/countries/1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Configuration

In `config-*.yml`:
```yaml
jwt:
  secret: "your-256-bit-secret-key"
  accessExpireTime: 60  # minutes
  refreshExpireTime: 10080  # minutes (7 days)
```

⚠️ **Important**: Change the JWT secret in production to a strong, random key.

## Swagger Documentation

### Access Swagger UI

When running in debug/development mode, access interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

### Features

- 📖 Complete API endpoint reference
- 🔍 Request/response schema documentation
- 🧪 Try-it-out functionality for testing endpoints
- 🔐 Bearer token authentication support
- 📝 Auto-generated from code annotations

### Using Authentication in Swagger

1. Click the "Authorize" button in Swagger UI
2. Enter: `Bearer <your-access-token>`
3. Click "Authorize"
4. Test protected endpoints directly from the UI

### Regenerating Documentation

After modifying API endpoints:
```bash
cd src
swag init -g cmd/main.go -o ./docs
```

## Rate Limiting

### Configuration

In `config-*.yml`:
```yaml
rateLimiter:
  enabled: true
  requestsPerMin: 100  # Requests allowed per minute
```

### Default Limits

- **Development**: 100 requests/minute
- **Production**: 60 requests/minute

### Behavior

When rate limit is exceeded:
- **Status Code**: 429 (Too Many Requests)
- **Response**: "Rate limit exceeded. Please try again later."

### Disabling Rate Limiting

Set `enabled: false` in the configuration:
```yaml
rateLimiter:
  enabled: false
  requestsPerMin: 100
```

## Protected vs Public Endpoints

### Public Endpoints (No Authentication Required)
- `GET /api/v1/health/` - Health check
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh

### Protected Endpoints (Authentication Required)
- All `/api/v1/countries/*` endpoints
- All `/api/v1/cities/*` endpoints
- All `/api/v1/colors/*` endpoints

## Testing the Features

### 1. Test Authentication Flow

```bash
# Register a user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123","email":"test@example.com"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# Use the access_token from login response
export TOKEN="your-access-token-here"

# Access protected endpoint
curl http://localhost:8080/api/v1/countries/1 \
  -H "Authorization: Bearer $TOKEN"
```

### 2. Test Rate Limiting

```bash
# Send multiple rapid requests
for i in {1..150}; do
  curl http://localhost:8080/api/v1/health/
done
```

You should see 429 responses after exceeding the limit.

### 3. Test Swagger UI

1. Start the server: `docker-compose up` or `go run cmd/main.go`
2. Open browser: `http://localhost:8080/swagger/index.html`
3. Explore the API documentation
4. Test endpoints using the "Try it out" feature

## Security Best Practices

1. **JWT Secret**: Use a strong, random 256-bit key in production
2. **HTTPS**: Always use HTTPS in production to protect tokens in transit
3. **Token Storage**: Store tokens securely on the client (e.g., httpOnly cookies)
4. **Rate Limiting**: Adjust limits based on your application's needs
5. **CORS**: Configure CORS settings appropriately for your domains
6. **Password Requirements**: Consider enforcing stronger password policies

## Troubleshooting

### "Invalid token" Error
- Check that the token hasn't expired
- Verify you're using the correct JWT secret
- Ensure the Authorization header format is: `Bearer <token>`

### "Rate limit exceeded" Error
- Wait for the rate limit window to reset (1 minute)
- Adjust `requestsPerMin` in configuration if needed
- Consider implementing per-user rate limiting

### Swagger UI Not Accessible
- Ensure the server is running in debug/development mode
- Check that `runMode` in config is set to "debug"
- Verify the server is running on the expected port

## Next Steps

Consider extending the authentication system with:
- Role-based access control (RBAC)
- Email verification
- Password reset functionality
- OAuth2/Social login integration
- Two-factor authentication (2FA)
- API key authentication for service-to-service communication
