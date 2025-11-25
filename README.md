# Zarish-HIS (Hospital Information System Module)

Clinical data management module for ZarishSphere Platform, providing EHR capabilities.

## Features

- Patient Management
- Encounter tracking
- FHIR-compliant data models

## API Endpoints

### Patients
- `POST /api/v1/patients` - Create patient
- `GET /api/v1/patients/:id` - Get patient by ID
- `GET /api/v1/patients` - List patients (paginated)

## Running

```bash
go run cmd/server/main.go
```

Server runs on port **8083**.

## Models

- **Patient**: Patient demographics and contact information
- **Encounter**: Clinical encounters between patients and providers

## Technology Stack

- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL
