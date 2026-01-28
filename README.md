# PressAndPlay

## Technical Documentation

*A Microservices-Based Sports Court Booking Platform*

**Version 1.0**  
**Author:** Adarsh Srinivasan  
**Repository:** github.com/adarshsrinivasan/PressAndPlay

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [System Architecture Overview](#2-system-architecture-overview)
3. [Microservices Deep Dive](#3-microservices-deep-dive)
   - 3.1 [User Service](#31-user-service)
   - 3.2 [Court Service](#32-court-service)
   - 3.3 [Events Service](#33-events-service)
4. [Inter-Service Communication](#4-inter-service-communication)
5. [Infrastructure Components](#5-infrastructure-components)
6. [API Reference](#6-api-reference)
7. [Data Models](#7-data-models)
8. [Deployment Guide](#8-deployment-guide)

---

## 1. Executive Summary

PressAndPlay is a web application designed to enable users to book time slots at nearby sporting centers. Built with a microservices architecture using Golang, the platform provides a scalable and maintainable solution for sports facility management.

### Key Features

- User registration and authentication with role-based access (Customer/Manager)
- Court creation and management by facility managers
- Real-time slot booking with availability tracking
- Location-based court discovery with distance calculation
- Rating system for courts
- Event-driven notifications via Kafka
- Cloud image storage integration with Google Cloud Storage

### Technology Stack

| Category | Technology |
|----------|------------|
| Language | Golang 1.19 |
| HTTP Framework | Gorilla Mux |
| RPC Framework | gRPC with Protocol Buffers |
| Databases | PostgreSQL (Users, Events), MongoDB (Courts) |
| Cache/Session | Redis |
| Message Queue | Apache Kafka |
| Cloud Storage | Google Cloud Storage |
| Containerization | Docker, Kubernetes |

---

## 2. System Architecture Overview

PressAndPlay follows a microservices architecture pattern with three core services communicating via both synchronous (gRPC) and asynchronous (Kafka) mechanisms.

### Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              CLIENT APPLICATIONS                             │
│                        (Web Browser / Mobile App)                           │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
                    ┌─────────────────┼─────────────────┐
                    │                 │                 │
                    ▼                 ▼                 ▼
           ┌───────────────┐ ┌───────────────┐ ┌───────────────┐
           │  USER SERVICE │ │ COURT SERVICE │ │EVENTS SERVICE │
           │  HTTP: 50000  │ │  HTTP: 50000  │ │  HTTP: 50000  │
           │  gRPC: 50001  │ │  gRPC: 50001  │ │  gRPC: 50001  │
           └───────┬───────┘ └───────┬───────┘ └───────┬───────┘
                   │                 │                 │
         ┌─────────┴─────────┐       │       ┌─────────┴─────────┐
         ▼                   ▼       ▼       ▼                   ▼
    ┌─────────┐         ┌─────────────────────────┐         ┌─────────┐
    │PostgreSQL│         │         KAFKA          │         │PostgreSQL│
    │ (Users) │         │    Message Broker       │         │ (Events)│
    └─────────┘         └─────────────────────────┘         └─────────┘
                                    │
                              ┌─────┴─────┐
                              ▼           ▼
                        ┌─────────┐ ┌─────────┐
                        │ MongoDB │ │  Redis  │
                        │(Courts) │ │(Session)│
                        └─────────┘ └─────────┘
```

### Service Ports Summary

| Service | HTTP Port (External) | gRPC Port (Internal) | Database |
|---------|---------------------|---------------------|----------|
| User Service | 50010 → 50000 | 50011 → 50001 | PostgreSQL |
| Court Service | 50020 → 50000 | 50021 → 50001 | MongoDB |
| Events Service | 50030 → 50000 | 50031 → 50001 | PostgreSQL |

---

## 3. Microservices Deep Dive

### 3.1 User Service

The User Service handles all user-related operations including registration, authentication, and profile management. It supports two user roles: Customer and Manager.

#### Responsibilities

- User registration with email validation
- Password-based authentication
- Session management via Redis
- User profile retrieval and deletion
- Publishing user-deleted events to Kafka
- Exposing gRPC endpoint for inter-service user lookups

#### Components

| File | Purpose |
|------|---------|
| user.go | Main entry point, initializes all components |
| api-handlers.go | HTTP REST endpoint handlers |
| user-dao.go | Data Access Object with PostgreSQL operations |
| user-ops.go | Business logic layer |
| grpc-handler.go | gRPC server implementation |
| session-handler-ops.go | Redis session management |
| message-queue-ops.go | Kafka producer for user events |
| db-ops.go | Database connection and query utilities |

#### Database Schema

Stored in PostgreSQL under schema `user_schema`, table `user_data`:

| Field | Type | Description |
|-------|------|-------------|
| id | UUID (PK) | Unique user identifier |
| firstName | String | User first name |
| lastName | String | User last name |
| dateOfBirth | String | Date of birth (YYYYMMDD) |
| gender | Enum | FEMALE(0), MALE(1), OTHERS(2) |
| address | String | Serialized address object |
| email | String (Unique) | User email with GIN index |
| phone | String | Phone number |
| password | String | User password (plaintext - needs hashing) |
| role | Enum | CUSTOMER(0), MANAGER(1) |
| verified | Boolean | Account verification status |
| lastLogin | Timestamp | Last login time |
| lastSessionID | String | Current session identifier |
| createdAt | Timestamp | Record creation time |
| updatedAt | Timestamp | Last update time |

---

### 3.2 Court Service

The Court Service manages sporting facilities, their time slots, and booking operations. It integrates with Google Cloud Storage for court images.

#### Responsibilities

- Court CRUD operations (create, read, list, delete)
- Slot booking with availability tracking
- Court rating system
- Location-based court discovery with distance calculation
- GCP image upload URL generation (signed URLs)
- Publishing slot-booked and court-deleted events
- Consuming user-deleted events to cascade delete courts

#### Components

| File | Purpose |
|------|---------|
| court.go | Main entry point, initializes all components |
| api-handlers.go | HTTP REST endpoint handlers |
| court-dao.go | Data Access Object with MongoDB operations |
| court-ops.go | Business logic (booking, rating, distance calc) |
| grpc-handler.go | gRPC server and client implementations |
| cloud-bucket-handler.go | GCP Cloud Storage integration |
| session-handler-ops.go | Session validation via Redis |
| message-queue-ops.go | Kafka producer/consumer |
| db-ops.go | MongoDB connection management |

#### Database Schema

Stored in MongoDB database `pressandplay`, collection `court_data`:

| Field | Type | Description |
|-------|------|-------------|
| _id | UUID | Unique court identifier |
| name | String | Court/facility name |
| address | Object | Full address (line1, line2, city, state, country, pincode) |
| location | String | GPS coordinates ("latitude, longitude") |
| phone | String | Contact phone number |
| rating | Float64 | Average rating (0-5) |
| ratingCount | Int | Number of ratings received |
| managerId | UUID | Owner/manager user ID |
| availableSlots | Array[Slot] | Time slots with booking status |
| imageUploadUrl | String | GCP signed URL for uploading |
| imageDownloadUrl | String | Public URL for viewing |
| sportType | String | Type of sport (e.g., BasketBall) |
| verified | Boolean | Court verification status |

---

### 3.3 Events Service

The Events Service tracks all bookings and provides notification capabilities for managers. It serves as the audit log for slot reservations.

#### Responsibilities

- Recording booking events from Kafka messages
- Providing unread notifications for managers
- Booking history retrieval
- Cascade deletion on user-deleted events
- Cascade deletion on court-deleted events
- Exposing gRPC endpoint for event queries

#### Components

| File | Purpose |
|------|---------|
| events.go | Main entry point, initializes all components |
| api-handlers.go | HTTP REST endpoint handlers |
| events-dao.go | Data Access Object with PostgreSQL operations |
| events-ops.go | Business logic layer |
| grpc-handler.go | gRPC server and client implementations |
| session-handler-ops.go | Session validation via Redis |
| message-queue-ops.go | Kafka consumer for multiple topics |
| db-ops.go | Database utilities with pagination support |

#### Database Schema

Stored in PostgreSQL under schema `events_schema`, table `event_data`:

| Field | Type | Description |
|-------|------|-------------|
| id | UUID (PK) | Unique event identifier |
| userId | UUID | Customer who made the booking |
| managerId | UUID | Manager of the booked court |
| courtId | UUID | Court where slot was booked |
| slotId | String | Specific slot identifier |
| time_start_hhmm | Int | Slot start time (HHMM format) |
| time_end_hhmm | Int | Slot end time (HHMM format) |
| bookingTimestamp | Timestamp | When booking was made |
| notified | Boolean | Whether manager was notified |
| createdAt | Timestamp | Record creation time |
| updatedAt | Timestamp | Last update time |

---

## 4. Inter-Service Communication

### 4.1 Synchronous Communication (gRPC)

Services use gRPC for real-time, request-response communication patterns where immediate data retrieval is required.

#### gRPC Service Definitions

| Service | Method | Input | Output | Called By |
|---------|--------|-------|--------|-----------|
| User | GetUser | UserModel (id) | UserModel (full) | Court, Events |
| Court | GetCourt | CourtModel (id) | CourtModel (full) | Events |
| Events | GetEventsByUserIdAndCourtId | EventModel (userId, courtId) | EventResponse (list) | Court |

#### gRPC Call Flow Examples

1. **Court Rating Flow:** Court Service calls User Service to verify the user is a CUSTOMER role before allowing rating
2. **Booking Flow:** Court Service calls User Service to validate user, then publishes Kafka event
3. **Notification Flow:** Events Service calls both User and Court services to enrich booking data with names

### 4.2 Asynchronous Communication (Kafka)

Kafka enables event-driven communication for operations that don't require immediate response, ensuring loose coupling and resilience.

#### Kafka Topics

| Topic | Producer | Consumers | Payload | Purpose |
|-------|----------|-----------|---------|---------|
| user-deleted | User Service | Court, Events | userId (string) | Cascade delete user's courts and events |
| court-deleted | Court Service | Events | courtId (string) | Cascade delete court's events |
| slot-booked | Court Service | Events | SlotBookedNotification (JSON) | Create event record for booking |

#### SlotBookedNotification Structure

```json
{
  "user_id": "uuid",
  "manager_id": "uuid",
  "court_id": "uuid",
  "slot_id": "string",
  "timestamp": "RFC3339 timestamp"
}
```

#### Event Flow Diagram

```
User Deletion Flow:
┌──────────────┐     user-deleted      ┌───────────────┐
│ User Service │ ───────────────────▶  │ Court Service │ → Deletes manager's courts
└──────────────┘          │            └───────────────┘
                          │            ┌────────────────┐
                          └──────────▶ │ Events Service │ → Deletes user's events
                                       └────────────────┘

Booking Flow:
┌───────────────┐     slot-booked      ┌────────────────┐
│ Court Service │ ───────────────────▶ │ Events Service │ → Creates event record
└───────────────┘                      └────────────────┘
```

---

## 5. Infrastructure Components

### 5.1 PostgreSQL

Used by User Service and Events Service for relational data with ACID compliance.

- **User Service:** `user_schema.user_data` with GIN index on email
- **Events Service:** `events_schema.event_data`
- **Extensions:** pg_trgm (text search), btree_gin (index optimization)

### 5.2 MongoDB

Used by Court Service for flexible document storage of court data with nested arrays.

- **Database:** pressandplay
- **Collection:** court_data

### 5.3 Redis

Centralized session management shared across all services.

- **Key:** Session UUID
- **Value:** JSON with sessionId, userId, lastLoginTime
- **TTL:** Configurable via REDIS_DATA_EXPIRATION_IN_HRS (default: 5 hours)

### 5.4 Apache Kafka

Message broker for asynchronous event-driven communication.

- **Broker:** Single node (kafka1:9092)
- **Zookeeper:** zoo1:2181
- **Replication Factor:** 1 (development setup)

### 5.5 Google Cloud Storage

Court images are stored in GCP with signed URL access.

- **Bucket:** pressandplay
- **Upload:** Signed PUT URLs (15-minute expiry)
- **Download:** Public URLs with auth user parameter

---

## 6. API Reference

### 6.1 User Service APIs

**Base URL:** `/api/v1/user`

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /create | None | Register a new user (Customer or Manager) |
| POST | /login | None | Authenticate user, returns session ID in header |
| GET | /{id} | Session | Get user profile by ID |
| DELETE | /{id} | Session | Delete user account |

#### POST /create - Create User

**Request Body:**

```json
{
  "firstName": "string",
  "lastName": "string",
  "dateOfBirth": "YYYYMMDD",
  "gender": 0|1|2,
  "address": {
    "address_line_1": "string",
    "address_line_2": "string",
    "city": "string",
    "state": "string",
    "country": "string",
    "pincode": "string"
  },
  "phone": "string",
  "role": 0|1,
  "email": "string",
  "password": "string"
}
```

**Response (201):** Created user object without password

#### POST /login - User Login

**Request Body:**

```json
{
  "email": "string",
  "password": "string"
}
```

**Response Header:** `User-Session-Id: <uuid>`

**Response (201):** User object

---

### 6.2 Court Service APIs

**Base URL:** `/api/v1/court`

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /create | Session (Manager) | Create a new court |
| GET | / | Optional | List all courts (with distance if location header) |
| GET | /{id} | Optional | Get court details by ID |
| DELETE | /{id} | Session | Delete a court |
| POST | /{id}/rating | Session (Customer) | Rate a court (0-5) |
| POST | /{court-id}/slot/{slot-id}/book | Session (Customer) | Book a specific slot |

#### POST /create - Create Court

**Required Header:** `User-Session-Id: <manager-session-uuid>`

**Request Body:**

```json
{
  "name": "string",
  "address": { ... },
  "location": "latitude, longitude",
  "phone": "string",
  "availableSlots": [
    {
      "slot_id": "string",
      "time_start_hhmm": 1000,
      "time_end_hhmm": 1100,
      "booked": false
    }
  ],
  "sportType": "string"
}
```

**Response (201):** Court object with imageUploadUrl and imageDownloadUrl

#### GET / - List Courts

**Optional Header:** `location: latitude, longitude`

**Response (200):** Array of CourtListModel with distance (if location provided), sorted by distance then rating

#### POST /{court-id}/slot/{slot-id}/book - Book Slot

**Required Header:** `User-Session-Id: <customer-session-uuid>`

**Response (200):** Updated court object with slot marked as booked

**Side Effect:** Publishes slot-booked event to Kafka

---

### 6.3 Events Service APIs

**Base URL:** `/api/v1/events`

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | / | Session | Get booking history for manager |
| POST | /create | Session | Manually create an event |
| GET | /notifications | Session | Get unread booking notifications |

#### GET /notifications - Unread Notifications

**Required Header:** `User-Session-Id: <manager-session-uuid>`

**Response (200):** Array of EventsListModel with enriched user/court names

**Side Effect:** Marks returned events as notified=true

---

## 7. Data Models

### 7.1 Protocol Buffer Definitions

Located in `libraries/proto/` - shared across all services for gRPC communication.

#### Common Types (common.proto)

```protobuf
message Address {
  string addressLine1 = 1;
  string addressLine2 = 2;
  string city = 3;
  string state = 4;
  string country = 5;
  string pincode = 6;
}
```

#### User Types (user.proto)

```protobuf
enum Gender { GENDER_FEMALE = 0; GENDER_MALE = 1; GENDER_OTHERS = 2; }
enum Role { ROLE_CUSTOMER = 0; ROLE_MANAGER = 1; }

message UserModel {
  string id = 1;
  string firstName = 2;
  string lastName = 3;
  string dateOfBirth = 4;
  Gender gender = 5;
  Address address = 6;
  string Phone = 7;
  Role role = 8;
  string email = 9;
  string password = 10;
  bool verified = 11;
}
```

#### Court Types (court.proto)

```protobuf
message Slot {
  string slotId = 1;
  int32 timeStartHHMM = 2;
  int32 timeEndHHMM = 3;
  bool booked = 4;
}

message CourtModel {
  string id = 1;
  string name = 2;
  Address address = 3;
  string location = 4;
  float distance = 5;
  string phone = 6;
  float rating = 7;
  int32 ratingCount = 8;
  string managerId = 9;
  repeated Slot availableSlots = 10;
  string imageUploadUrl = 11;
  string imageDownloadUrl = 12;
  string sportType = 13;
  bool verified = 14;
}
```

---

## 8. Deployment Guide

### 8.1 Docker Compose Deployment

For local development and testing:

```bash
# Build and push all service images
make user
make court
make events

# Deploy the stack
make deploy-docker

# Tear down
make undeploy-docker
```

### 8.2 Kubernetes Deployment

For production-like environments:

```bash
# Deploy all components (with timing for dependencies)
make deploy-kubernetes

# Tear down
make undeploy-kubernetes
```

#### Kubernetes Components

| Component | Type | Replicas | Notes |
|-----------|------|----------|-------|
| PostgreSQL | Deployment | 1 | PersistentVolume at /mnt/postgres-data |
| MongoDB | Deployment | 1 | PersistentVolume at /mnt/mongo-data |
| Redis | Deployment | 1 | PersistentVolume at /mnt/redis-data |
| Zookeeper | Deployment | 1 | Required for Kafka |
| Kafka | Deployment | 1 | Depends on Zookeeper |
| User Service | StatefulSet | 1 | OrderedReady pod management |
| Court Service | StatefulSet | 1 | OrderedReady pod management |
| Events Service | StatefulSet | 1 | OrderedReady pod management |
| Ingress | Ingress | N/A | NGINX ingress with path rewriting |

### 8.3 Environment Variables

Key configuration variables (with defaults):

| Variable | Service | Default | Description |
|----------|---------|---------|-------------|
| GRPC_SERVER_PORT | All | 50001 | gRPC server listen port |
| HTTP_SERVER_PORT | All | 50000 | HTTP server listen port |
| POSTGRES_HOST | User, Events | localhost | PostgreSQL hostname |
| MONGO_HOST | Court | localhost | MongoDB hostname |
| REDIS_HOST | All | localhost | Redis hostname |
| KAFKA_HOST | All | localhost | Kafka broker hostname |
| GCP_BUCKET_NAME | Court | pressandplay | GCP Storage bucket |
| REDIS_DATA_EXPIRATION_IN_HRS | User | 5 | Session TTL in hours |

### 8.4 Ingress Configuration

The Kubernetes ingress routes requests based on path prefix:

| Path | Target Service | Target Port |
|------|----------------|-------------|
| /user/* | user | 50000 |
| /court/* | court | 50000 |
| /events/* | events | 50000 |

---

*End of Document*
