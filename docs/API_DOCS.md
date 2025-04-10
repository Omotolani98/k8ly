# 📦 K8ly API Collections

---

## 📁 Authly API

**Base URL:** `https://authly.k8ly.io`

### POST `/auth/register`
Registers a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "message": "User registered successfully",
  "api_key": "k8ly_sk_xxx"
}
```

---

### POST `/auth/login`
Logs in an existing user.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "message": "Login successful",
  "api_key": "k8ly_sk_xxx"
}
```

---

## 📁 Reqly API

**Base URL:** `https://reqly.k8ly.io`

### POST `/webhook`
Accepts webhook/event requests.

**Headers:**
```
Authorization: Bearer <API_KEY>
```

**Request Body:** `any JSON payload`

**Response:**
```json
{
  "message": "Webhook received"
}
```

---

### GET `/logs`
Fetch request logs for this API key.

**Headers:**
```
Authorization: Bearer <API_KEY>
```

**Response:**
```json
[
  {
    "id": 1,
    "service": "reqly",
    "data": { ... },
    "created_at": "2025-04-09T10:00:00Z"
  }
]
```

---

## 📁 Logly API

**Base URL:** `https://logly.k8ly.io`

### POST `/logs`
Submit a log manually.

**Headers:**
```
Authorization: Bearer <API_KEY>
Content-Type: application/json
```

**Request Body:**
```json
{
  "level": "info",
  "message": "Payment succeeded",
  "service": "checkout"
}
```

**Response:**
```json
{
  "message": "Log recorded"
}
```

---

### GET `/logs`
Fetch logs submitted by this user.

**Headers:**
```
Authorization: Bearer <API_KEY>
```

**Query Params (optional):**
- `service=checkout`
- `level=error`

**Response:**
```json
[
  {
    "id": 1,
    "message": "Payment succeeded",
    "level": "info",
    "service": "checkout",
    "created_at": "2025-04-09T10:00:00Z"
  }
]

