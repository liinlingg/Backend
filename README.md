# Unicom Backend

Backend service for **Unicom Project**, a multi-tenant campus community platform for universities.  
This repository contains the backend code responsible for posts, events, visibility rules, user's role and permission mangement, and basic notification logic.

> Note: This is the backend part of a larger project (web + mobile).  

---

## Overview

Unicom is a platform where students, clubs, and organizations can:

- Post announcements and activities
- Create, manage and attend events
- Target content to specific faculties / departments / roles
- Explore trending hashtags and categorized feeds

The user's allow to do an action that has permission only.

This backend exposes REST APIs for the web and mobile clients, handling all business logic, and validation.

---

## Main Features (Backend)

- User-based visibility using organization path (`org_path`) and position (`position_key`)
- Post & event creation, update, and soft-delete
- Cursor-based pagination for feed endpoints
- Category/Role-based filtering for posts and events
- Trending hashtag counting and aggregation
- Basic notification references (e.g. when events are updated or deleted)
- Media upload support (for post/event images)

---

## Tech Stack

- **Language:** Go
- **Framework:** Fiber (or Gin – update to match your code)
- **Database:** MongoDB
- **Auth:** JWT-based authentication (if applicable)
- **Other:** Git, Postman

---

## 📁 Project Structure

```bash
.
├── README.md
└── main-webbase/
    ├── app/                       # Fiber application setup (middleware, providers)
    ├── bootstrap/                 # Startup helpers (e.g., Mongo indexes)
    ├── cmd/                       # CLI entrypoints (cmd/main.go)
    ├── config/                    # Environment loading & config structs
    ├── database/                  # MongoDB client wiring
    ├── docs/                      # Swagger (docs.go/json/yaml)
    ├── dto/                       # Request/response DTO definitions
    ├── internal/
    │   ├── accessctx/             # Access context helpers
    │   ├── authctx/               # Authentication context helpers
    │   ├── controllers/           # HTTP handlers
    │   ├── cursor/                # Cursor + pagination utilities
    │   ├── middleware/            # Fiber middlewares (JWT, UID injection, etc.)
    │   ├── models/                # MongoDB models
    │   ├── repository/            # Data access layer
    │   ├── routes/                # Route group definitions
    │   ├── services/              # Business logic
    │   └── utils/                 # Shared helpers (cursor cookies, profanity, time)
    ├── uploads/                   # Local media storage for media payloads
    ├── go.mod
    ├── go.sum
    └── main.go                    # Service bootstrap
```
