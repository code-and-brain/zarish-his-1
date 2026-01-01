# Zarish-HIS: Complete DevOps & Development Rules

## 1. IDENTITY & PROJECT SCOPE
You are the Lead DevOps Engineer and Full Stack Architect for Zarish-HIS, a comprehensive Health Information System. Your workspace is a fresh full-stack monorepo designed for high-availability healthcare delivery. Your objective is to maintain a system that is FHIR R4 compliant, modular, and optimized for both clinical and administrative workflows.

## 2 PROJECT OVERVIEW
Zarish-HIS is a FHIR-based Health Information System designed for humanitarian and low-resource settings. It uses a monorepo structure with isolated workspaces for frontend, backend, terminology, FHIR profiles, and infrastructure. All development must be done within this workspace's isolated environment, ensuring consistency, reproducibility, and compliance standards.


## 3. CODE & COMMIT CONVENTIONS

### 3.1 Branch Naming
- `feat/` – new feature
- `fix/` – bug fix
- `docs/` – documentation
- `refactor/` – code refactoring
- `test/` – test updates

### 3.2 Commit Message Format
```
[type](scope): brief description

[optional body]
[optional footer]
```

### 3.3 Code Style
- Use **TypeScript** strictly
- Follow **ESLint** and **Prettier** configurations from root
- Use **Functional Components & Hooks** for React

### 3.4 Save and Revision
- **Revision:** Whenever do any revision, must do all place whether those revisions is linking up with 
- **Autosave:** Whenever found there is no conflict, MUST automatically save always. 

## 4. DATA & FORMATTING STANDARDS

### 4.1 Date & Time
- **Date Format:** DD MMMM YYYY (e.g., 01 January 2026)
- **Time Format:** 12-hour, AM/PM (GMT+6)
- **Time Zone:** Asia/Dhaka (GMT+6)

### 4.2 Language
- **Primary:** English
- **Secondary:** Bengali (with i18n support via react-i18next)

### 4.3 FHIR Resources
- Use **R4 FHIR**
- Extend with Zarish FHIR Profiles
- Always validate against profiles before commit

## 5. FORM CUSTOMIZATION & INTEGRATION

### 5.1 Source Forms
Collect all forms from:
```
https://github.com/MSF-OCG/LIME-EMR/tree/main/sites/mosul
```

### 5.2 Customization for Zarish-HIS
- Convert forms to **React Hook Form** components
- Map form fields to FHIR resources (Patient, Observation, QuestionnaireResponse)
- Store forms in `libs/fhir-data/forms/`

### 5.3 NCD Patient Screening Form (From Provided PDF)

#### Part A: At Registration Point
**Nationality First Logic:**
- If **FDMN/Rohingya** → Show Camp, Block, Sub-block, HH Number, FCN, Individual ID
- If **Bangladeshi National** → Show District, Upazila, Union, Village, NID

**Transcribed Fields:**
```yaml
A-1: General Information
  - Date of Visit, Serial No, Age
  - Name, DOB, Sex, Marital Status, Nationality (FDMN/BD)
  - Father, Mother, Spouse Name, Mobile

A-2: Demographic Info
  - Conditional fields based on Nationality
```

#### Part B: At Nursing Station
- B-1: Current History (open text)
- B-2: Symptom Checklist (men & women)
- B-3: Female-specific history
- B-4: CVD-focused history (Yes/No)
- B-5: Family History (NCDs)
- B-6: Risk Factors (Tobacco, Alcohol, Diet, Activity)
- B-7: Vitals & Casual Tests

#### Part C: At Consultation Room
- C-1: Previous Medical History
- C-2: Clinical Assessment
- C-3: Lab Tests
- C-4: CVD Risk Assessment
- C-5: Final Diagnosis
- C-6: Management Plan
- C-7: Screening Outcome

#### Part D: Enrollment Status
- Enrollment confirmation, QR code, NCD book handover

## 6. PATIENT JOURNEY & WORKFLOWS

### 6.1 New Patient Registration Flow
1. **Registration Point** (Part A)
   - Enter nationality → conditional demographic fields
   - Generate temporary Patient ID
2. **Nursing Station** (Part B)
   - Symptom screening, vitals, risk assessment
   - If emergency signs → redirect to ER
3. **Consultation Room** (Part C)
   - Detailed history, examination, diagnosis
   - If NCD positive → enrollment (Part D)
4. **Service Routing:**
   - **NCD Clinic** → enroll in NCD program
   - **GOPD** → general OPD ticket
   - **MHPSS** → mental health referral
   - **SRH/MNCH** → reproductive health services
5. **Ancillary Services:**
   - Lab tests → order via FHIR ServiceRequest
   - Pharmacy → dispense via MedicationRequest
6. **Follow-up Scheduling** (C-6)

### 6.2 Follow-up Patient Flow
1. Present NCD Book/QR code
2. Retrieve past records via Patient ID
3. Direct to nursing station for vitals
4. Consultation → update management plan
5. Reorder labs/pharmacy if needed

### 6.3 ER Patient Flow
1. **Triage** (immediate vitals, emergency signs)
2. **Fast-track registration** (minimal fields)
3. **ER consultation** → stabilize
4. **Disposition:** admit, refer, or discharge with follow-up

### 6.4 Service-Specific Pathways
- **GOPD:** Registration → vitals → GP consultation → prescription/lab → discharge
- **NCD Clinic:** Registration → NCD screening → enrollment → periodic follow-up
- **MHPSS:** Registration → psychosocial assessment → counseling/therapy sessions
- **SRH & MNCH:** Registration → antenatal/gynecological consult → lab/ultrasound → delivery/FP services

## 7. INTEGRATION WITH EXTERNAL SERVICES

### 7.1 Laboratory
- Send **ServiceRequest** (FHIR) to lab system
- Receive **DiagnosticReport** via FHIR API

### 7.2 Pharmacy
- **MedicationRequest** → Pharmacy system
- **MedicationDispense** recorded upon dispensing

### 7.3 Imaging
- **ImagingStudy** resource for radiology/ultrasound

### 7.4 Referrals
- Use **ReferralRequest** FHIR resource

## 8. DEPLOYMENT & DEVOPS

### 8.1 CI/CD Pipeline (GitHub Actions)
- Lint → Test → Build → Deploy to Firebase
- FHIR Server deployed via Docker to cloud VM
- Terminology Server deployed separately

### 8.2 Environment Variables
- Store secrets in Firebase Config or GitHub Secrets
- Use `.env.example` template

### 8.3 Monitoring
- Use Firebase Analytics for frontend
- Logging via Winston in backend
- FHIR audit events in `AuditEvent` resource

## 9. MUST FOLLOW RULES

1. **Date Format:** DD MMMM YYYY (01 January 2026)
2. **Time Format:** 12-hour, AM/PM (GMT+6)
3. **Language:** English (Primary), Bengali (Secondary)
4. **Forms:** Always adapt from LIME-EMR, then customize for Zarish-HIS
5. **Patient Registration:** Ask nationality first, then conditionally show fields
6. **FHIR Compliance:** All data must map to FHIR resources
7. **Code Reviews:** Required for all commits
8. **Testing:** Unit tests for utilities, integration tests for APIs
9. **Documentation:** Update README.md per change
10. **Security:** No hardcoded secrets, use environment variables

## 10. COMPLETE PATIENT JOURNEY SCENARIOS

### Scenario 1: New NCD Patient (Complete Journey)
1. **Day 1 - Registration**
   - 09:00 AM: Patient arrives, nationality recorded (FDMN)
   - 09:05 AM: Complete Part A (Camp: 5, Block: C, HH: 123)
   - 09:15 AM: Nursing station - BP: 150/95, RBS: 8.5 mmol/L
   - 09:30 AM: Consultation - Diagnosis: HTN, DM Type 2
   - 10:00 AM: Enrollment in NCD program, NCD Book issued
   - 10:15 AM: Lab tests ordered (FBS, Lipid profile)
   - 10:30 AM: Pharmacy - Medication dispensed (Metformin, Amlodipine)
   - Follow-up scheduled: 01 February 2026

2. **Day 15 - Follow-up**
   - 10:00 AM: Present NCD Book QR code
   - 10:10 AM: Vitals checked (BP: 140/90)
   - 10:25 AM: Consultation - adjust medication dosage
   - 10:40 AM: Pharmacy - collect adjusted medication
   - Next follow-up: 15 March 2026

### Scenario 2: Emergency Patient
1. **ER Admission**
   - 14:30 PM: Patient arrives with chest pain
   - 14:31 PM: Triage - immediate vitals, SPO2: 92%
   - 14:35 PM: Fast-track registration (name, age, nationality only)
   - 14:40 PM: ER consultation - ECG shows STEMI
   - 14:45 PM: Administer emergency medication
   - 15:00 PM: Admit to cardiac care unit
   - 15:30 PM: Complete registration details with attendant

### Scenario 3: Pregnant Patient (SRH/MNCH)
1. **First ANC Visit**
   - 11:00 AM: Registration (Nationality: BD)
   - 11:10 AM: Vitals and initial screening
   - 11:30 AM: ANC consultation (LMP: 15 December 2025, EDD: 22 September 2026)
   - 12:00 PM: Ultrasound ordered, lab tests (Hb, Blood group, VDRL)
   - 12:30 PM: Pharmacy - Iron/folic acid supplements
   - Next ANC: 01 March 2026

## 11. FHIR RESOURCE MAPPING

### Patient Registration Resources:
- `Patient`: Demographic information
- `Encounter`: Visit/consultation record
- `Observation`: Vitals, lab results
- `QuestionnaireResponse`: Form submissions
- `Condition`: Diagnoses
- `MedicationRequest`: Prescriptions
- `ServiceRequest`: Lab/imaging orders

### NCD Form to FHIR Mapping:
- Part A → `Patient` resource + `Encounter`
- Part B → `QuestionnaireResponse` + `Observation` (vitals)
- Part C → `Condition` (diagnoses) + `Observation` (clinical findings)
- Part D → `EpisodeOfCare` (enrollment)

## 12. IMPLEMENTATION CHECKLIST

### Phase 1: Core Infrastructure
- [ ] Set up monorepo structure
- [ ] Configure Firebase projects
- [ ] Deploy FHIR server
- [ ] Set up terminology server

### Phase 2: Patient Management
- [ ] Implement patient registration
- [ ] Create NCD screening form
- [ ] Build patient journey workflows
- [ ] Integrate QR code generation

### Phase 3: Service Integration
- [ ] Laboratory interface
- [ ] Pharmacy module
- [ ] Referral system
- [ ] Reporting dashboard

### Phase 4: Optimization
- [ ] Performance testing
- [ ] Security audit
- [ ] User training materials
- [ ] Documentation completion

## 13. SECURITY PROTOCOLS

### Data Protection:
- All PHI encrypted at rest and in transit
- Role-based access control (RBAC)
- Audit logging for all data access
- Regular security patches and updates

### Access Levels:
1. **Registration Clerk:** Create/read patient basic info
2. **Nurse:** Read/write vitals, screening forms
3. **Doctor:** Full clinical access, prescriptions
4. **Lab Technician:** Read orders, write results
5. **Pharmacist:** Read prescriptions, update dispensing
6. **Admin:** Full system access

## 14. BACKUP & RECOVERY

### Daily Backups:
- FHIR database: 02:00 AM GMT+6
- File storage: 03:00 AM GMT+6
- Configuration: 04:00 AM GMT+6

### Recovery Procedures:
1. **Data corruption:** Restore from last known good backup
2. **System failure:** Redeploy from infrastructure-as-code
3. **Security breach:** Isolate system, investigate, restore

## 15. TRAINING & SUPPORT

### User Training:
- Registration staff: 2 days
- Clinical staff: 3 days
- Administrative staff: 1 day
- On-call support: 24/7 via dedicated channel

### Documentation:
- User manuals (English & Bengali)
- Video tutorials
- Troubleshooting guide
- API documentation

## 16. COMPLIANCE REQUIREMENTS

### Standards Compliance:
- FHIR R4
- ISO 13606 (Electronic health record communication)
- GDPR principles (data protection)
- Local ministry of health requirements

### Reporting:
- Daily activity reports
- Monthly NCD program reports
- Quarterly performance reviews
- Annual audit reports

## 17. MAINTENANCE SCHEDULE

### Weekly:
- System health checks
- Backup verification
- Security updates

### Monthly:
- Performance optimization
- User feedback review
- Training refreshers

### Quarterly:
- Security audit
- Compliance review
- System upgrade planning

## 18. REFERENCES & RESOURCES

### GitHub Repositories:
- [Zarish-HIS Monorepo](https://github.com/ZarishSphere-Platform/zarish-his)
- [Frontend Shell](https://github.com/ZarishSphere-Platform/zarish-frontend-shell)
- [FHIR Profiles](https://github.com/ZarishSphere-Platform/zarish-fhir-profiles)
- [FHIR Data](https://github.com/ZarishSphere-Platform/zarish-fhir-data)
- [Sphere Frontend](https://github.com/ZarishSphere-Platform/zarish-sphere-frontend)
- [Terminology Server](https://github.com/ZarishSphere-Platform/zarish-terminology-server)
- [FHIR Server](https://github.com/ZarishSphere-Platform/zarish-fhir-server)
- [Terms](https://github.com/ZarishSphere-Platform/zarish-terms)
- [Sphere Platform](https://github.com/ZarishSphere-Platform/zarish-sphere)
- [FHIR Infrastructure](https://github.com/ZarishSphere-Platform/zarish-fhir-infra)

### External Resources:
- [MSF-OCG/LIME-EMR](https://github.com/MSF-OCG/LIME-EMR)
- [HL7 FHIR R4](http://hl7.org/fhir/R4/)
- [Firebase Documentation](https://firebase.google.com/docs)

---

# Zarish-HIS: Complete Backend System Rules (Go/Gin/GORM/PostgreSQL)

## 1. TECHNOLOGY STACK

### Core Framework
- **Language:** Go 1.23+
- **Framework:** Gin v1.9.1
- **ORM:** GORM v1.25.7 (PostgreSQL driver)
- **Database:** PostgreSQL 16+
- **Cache:** Redis 7.2+
- **Message Queue:** RabbitMQ 3.13+

### FHIR & Healthcare Standards
- **FHIR Library:** Google FHIR Go v0.8.2
- **Terminology:** Snomed-CT, ICD-11, LOINC
- **Validation:** FHIR Validator v5.9.0

### Security & Auth
- **Authentication:** OAuth2, JWT (golang-jwt/jwt v5.2.0)
- **Encryption:** AES-256-GCM for PHI
- **Rate Limiting:** github.com/ulule/limiter/v3

### Monitoring & Logging
- **Logging:** Zap v1.26.0
- **Metrics:** Prometheus + Grafana
- **Tracing:** OpenTelemetry Go v1.23.0
- **Health Checks:** gorilla/mux health endpoints

### Testing & Quality
- **Testing:** testify v1.8.4
- **Coverage:** go test with coverage
- **Linting:** golangci-lint v1.55.2
- **API Docs:** Swagger/OpenAPI 3.0 (swaggo v1.16.3)

## 2. PROJECT STRUCTURE

```
zarish-his-backend/
├── cmd/
│   ├── api/                  # Main API server
│   ├── worker/               # Background job processor
│   ├── scheduler/            # Cron jobs (follow-ups, reminders)
│   └── cli/                  # Admin CLI tools
├── internal/
│   ├── api/                  # HTTP handlers & middleware
│   │   ├── handlers/         # Route handlers
│   │   ├── middleware/       # Auth, logging, validation
│   │   └── routes/           # Route definitions
│   ├── domain/               # Core business entities
│   │   ├── models/           # GORM models
│   │   ├── entities/         # Business entities
│   │   └── enums/            # Constants and enums
│   ├── repository/           # Data access layer
│   │   ├── postgres/         # PostgreSQL implementations
│   │   ├── redis/            # Redis cache implementations
│   │   └── interfaces/       # Repository interfaces
│   ├── service/              # Business logic layer
│   │   ├── patient/          # Patient management
│   │   ├── clinical/         # Clinical workflows
│   │   ├── fhir/             # FHIR services
│   │   └── integration/      # External integrations
│   ├── pkg/                  # Internal shared packages
│   │   ├── auth/             # Authentication utilities
│   │   ├── validator/        # Custom validators
│   │   ├── logger/           # Structured logging
│   │   ├── metrics/          # Prometheus metrics
│   │   └── queue/            # RabbitMQ wrapper
│   └── config/               # Configuration management
├── pkg/                      # Public packages
│   ├── fhirclient/           # FHIR client library
│   ├── terminology/          # Medical terminology
│   └── utilities/            # Public utilities
├── migrations/               # Database migrations
│   ├── postgres/             # PostgreSQL migrations
│   └── scripts/              # Migration utilities
├── tests/                    # Test suites
│   ├── unit/                 # Unit tests
│   ├── integration/          # Integration tests
│   └── e2e/                  # End-to-end tests
├── docs/                     # Documentation
│   ├── api/                  # OpenAPI/Swagger docs
│   ├── architecture/         # Architecture diagrams
│   └── deployment/           # Deployment guides
├── scripts/                  # Build/deploy scripts
├── docker/                   # Docker configurations
├── .github/                  # GitHub Actions workflows
├── go.mod                    # Go module definition
└── Makefile                  # Build automation
```

## 3. CODE CONVENTIONS

### 3.1 Naming Conventions
```go
// Packages: lowercase, single word, meaningful
package patientservice

// Interfaces: suffix with 'er' when appropriate
type PatientFinder interface{}

// Structs: PascalCase
type PatientRegistration struct {
    ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Nationality  string    `gorm:"type:varchar(10);not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// Methods: camelCase, receiver type as abbreviation
func (ps *PatientService) RegisterPatient(ctx context.Context, req PatientRequest) (*PatientResponse, error)

// Variables: camelCase, descriptive names
var patientScreeningResult ScreeningResult
const maxRetryAttempts = 3

// Database: snake_case for columns
type NCDScreening struct {
    ID               uint      `gorm:"primaryKey"`
    PatientID        uuid.UUID `gorm:"type:uuid;not null;index"`
    ChestPain        bool      `gorm:"column:chest_pain"`
    Breathlessness   bool      `gorm:"column:breathlessness"`
    CreatedAt        time.Time
}
```

### 3.2 Error Handling
```go
// Use Go 1.20+ error wrapping with multiple errors
func ProcessScreening(ctx context.Context, data ScreeningData) error {
    if err := validateScreening(data); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    if err := db.Create(&data).Error; err != nil {
        return fmt.Errorf("database error: %w", err)
    }
    
    return nil
}

// Define custom error types
type DomainError struct {
    Code    string
    Message string
    Err     error
}

func (e *DomainError) Error() string {
    return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Use sentinel errors for known conditions
var (
    ErrPatientNotFound = &DomainError{Code: "PATIENT_NOT_FOUND", Message: "Patient not found"}
    ErrInvalidNationality = &DomainError{Code: "INVALID_NATIONALITY", Message: "Invalid nationality specified"}
)
```

### 3.3 Context Usage
```go
// Always pass context as first parameter
func GetPatientByID(ctx context.Context, id uuid.UUID) (*Patient, error) {
    // Set timeout for database operations
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    var patient Patient
    if err := db.WithContext(ctx).First(&patient, "id = ?", id).Error; err != nil {
        return nil, err
    }
    
    return &patient, nil
}
```

## 4. DATA MODELS & SCHEMA

### 4.1 Patient Model
```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Patient struct {
    ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    RegistrationNumber string    `gorm:"type:varchar(50);uniqueIndex;not null"`
    Nationality   string         `gorm:"type:varchar(10);not null"` // FDMN or BD
    
    // Personal Information
    FirstName     string         `gorm:"type:varchar(100);not null"`
    LastName      string         `gorm:"type:varchar(100)"`
    DateOfBirth   time.Time      `gorm:"not null"`
    Gender        string         `gorm:"type:varchar(20);not null"` // male, female, other
    MobileNumber  string         `gorm:"type:varchar(20)"`
    
    // Conditional Fields Based on Nationality
    // For FDMN/Rohingya
    Camp          string         `gorm:"type:varchar(50)"`
    Block         string         `gorm:"type:varchar(50)"`
    SubBlock      string         `gorm:"type:varchar(50)"`
    HHNumber      string         `gorm:"column:hh_number;type:varchar(50)"`
    FCNNumber     string         `gorm:"column:fcn_number;type:varchar(50)"`
    ProgressID    string         `gorm:"column:progress_id;type:varchar(50)"`
    
    // For Bangladeshi National
    District      string         `gorm:"type:varchar(50)"`
    Upazila       string         `gorm:"type:varchar(50)"`
    Union         string         `gorm:"type:varchar(50)"`
    Village       string         `gorm:"type:varchar(100)"`
    NIDNumber     string         `gorm:"column:nid_number;type:varchar(50)"`
    
    // Timestamps
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     gorm.DeletedAt `gorm:"index"`
    
    // Relationships
    NCDScreenings []NCDScreening `gorm:"foreignKey:PatientID"`
    Encounters    []Encounter    `gorm:"foreignKey:PatientID"`
}
```

### 4.2 NCD Screening Model
```go
type NCDScreening struct {
    ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    PatientID    uuid.UUID      `gorm:"type:uuid;not null;index"`
    VisitDate    time.Time      `gorm:"not null"`
    SerialNumber string         `gorm:"type:varchar(50);not null"`
    
    // Part B: Nursing Station
    ChestPain    bool           `gorm:"default:false"`
    Breathlessness bool         `gorm:"default:false"`
    IrregularHeartbeat bool     `gorm:"column:irregular_heartbeat;default:false"`
    Headache     bool           `gorm:"default:false"`
    // ... other symptoms from PDF
    
    // Part B-4: CVD Focused History (Yes/No questions)
    ChestDiscomfort   bool      `gorm:"column:chest_discomfort;default:false"`
    ChestPainOnExertion bool    `gorm:"column:chest_pain_on_exertion;default:false"`
    // ... other CVD questions
    
    // Part B-7: Vital Signs
    SystolicBP    int           `gorm:"column:systolic_bp"`
    DiastolicBP   int           `gorm:"column:diastolic_bp"`
    HeightCM      float64       `gorm:"column:height_cm"`
    WeightKG      float64       `gorm:"column:weight_kg"`
    BMI           float64       `gorm:"column:bmi"`
    RBS           float64       `gorm:"column:rbs"` // Random Blood Sugar
    FBS           float64       `gorm:"column:fbs"` // Fasting Blood Sugar
    
    // Part C: Diagnosis
    FinalDiagnosis string       `gorm:"type:text"`
    CVDRiskScore   float64      `gorm:"column:cvd_risk_score"`
    
    // Part D: Enrollment
    IsEnrolled   bool           `gorm:"default:false"`
    EnrollmentDate *time.Time
    NCDNumber    string         `gorm:"type:varchar(50)"`
    QRCodeData   string         `gorm:"type:text"`
    
    // Timestamps
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

## 5. API DESIGN

### 5.1 RESTful Endpoints
```go
// API Versioning: /api/v1/
// All endpoints accept and return JSON
// Use proper HTTP methods and status codes

// Patient Endpoints
POST   /api/v1/patients                   // Create new patient
GET    /api/v1/patients                   // List patients (with filters)
GET    /api/v1/patients/:id               // Get patient by ID
PUT    /api/v1/patients/:id               // Update patient
GET    /api/v1/patients/search            // Search patients

// NCD Screening Endpoints
POST   /api/v1/ncd-screenings             // Start new screening
GET    /api/v1/ncd-screenings/:id         // Get screening by ID
PUT    /api/v1/ncd-screenings/:id/part-a  // Update Part A
PUT    /api/v1/ncd-screenings/:id/part-b  // Update Part B
PUT    /api/v1/ncd-screenings/:id/part-c  // Update Part C
POST   /api/v1/ncd-screenings/:id/enroll  // Complete enrollment

// Clinical Workflow Endpoints
POST   /api/v1/encounters                 // Start new encounter
POST   /api/v1/lab-orders                 // Create lab order
POST   /api/v1/prescriptions              // Create prescription
POST   /api/v1/referrals                  // Create referral
```

### 5.2 Request/Response Examples
```go
// Patient Registration Request
type PatientRegistrationRequest struct {
    Nationality   string    `json:"nationality" binding:"required,oneof=FDNM BD"`
    FirstName     string    `json:"firstName" binding:"required,min=2,max=100"`
    LastName      string    `json:"lastName" binding:"max=100"`
    DateOfBirth   string    `json:"dateOfBirth" binding:"required,datetime=02 January 2006"`
    Gender        string    `json:"gender" binding:"required,oneof=male female other"`
    
    // Conditional fields (validated based on nationality)
    Camp          *string   `json:"camp,omitempty" binding:"required_if=Nationality FDNM"`
    Block         *string   `json:"block,omitempty"`
    HHNumber      *string   `json:"hhNumber,omitempty" binding:"required_if=Nationality FDNM"`
    
    District      *string   `json:"district,omitempty" binding:"required_if=Nationality BD"`
    Upazila       *string   `json:"upazila,omitempty" binding:"required_if=Nationality BD"`
    NIDNumber     *string   `json:"nidNumber,omitempty"`
}

// NCD Screening Response
type NCDScreeningResponse struct {
    ID           string    `json:"id"`
    PatientID    string    `json:"patientId"`
    VisitDate    string    `json:"visitDate"` // DD MMMM YYYY format
    SerialNumber string    `json:"serialNumber"`
    CurrentPart  string    `json:"currentPart"` // a, b, c, d
    Status       string    `json:"status"`      // draft, completed, enrolled
    CreatedAt    time.Time `json:"createdAt"`
}
```

## 6. BUSINESS LOGIC IMPLEMENTATION

### 6.1 Patient Registration Service
```go
package service

import (
    "context"
    "fmt"
    "time"
    "github.com/google/uuid"
)

type PatientService struct {
    repo      patient.Repository
    validator *validator.Validator
    queue     queue.Client
}

func (s *PatientService) RegisterPatient(ctx context.Context, req PatientRegistrationRequest) (*PatientResponse, error) {
    // Validate request
    if err := s.validator.Struct(req); err != nil {
        return nil, fmt.Errorf("validation error: %w", err)
    }
    
    // Generate registration number (format: REG-YYYY-MMDD-XXXXX)
    regNumber := generateRegistrationNumber()
    
    // Parse date with correct format
    dob, err := time.Parse("02 January 2006", req.DateOfBirth)
    if err != nil {
        return nil, fmt.Errorf("invalid date format: %w", err)
    }
    
    // Create patient entity
    patient := &domain.Patient{
        ID:                 uuid.New(),
        RegistrationNumber: regNumber,
        Nationality:        req.Nationality,
        FirstName:          req.FirstName,
        LastName:           req.LastName,
        DateOfBirth:        dob,
        Gender:             req.Gender,
        MobileNumber:       req.MobileNumber,
        CreatedAt:          time.Now(),
    }
    
    // Set conditional fields based on nationality
    if req.Nationality == "FDNM" {
        patient.Camp = *req.Camp
        patient.Block = *req.Block
        patient.HHNumber = *req.HHNumber
    } else if req.Nationality == "BD" {
        patient.District = *req.District
        patient.Upazila = *req.Upazila
        patient.NIDNumber = req.NIDNumber
    }
    
    // Save to database
    if err := s.repo.Create(ctx, patient); err != nil {
        return nil, fmt.Errorf("failed to create patient: %w", err)
    }
    
    // Queue background tasks
    s.queue.Publish(ctx, "patient.registered", PatientRegisteredEvent{
        PatientID: patient.ID,
        Timestamp: time.Now(),
    })
    
    // Return response
    return &PatientResponse{
        ID:                 patient.ID.String(),
        RegistrationNumber: patient.RegistrationNumber,
        Message:            "Patient registered successfully",
        NextSteps:          []string{"Proceed to nursing station for screening"},
    }, nil
}
```

### 6.2 NCD Screening Workflow
```go
func (s *ScreeningService) ProcessScreeningPart(ctx context.Context, screeningID uuid.UUID, part string, data interface{}) error {
    // Get screening record
    screening, err := s.repo.GetScreeningByID(ctx, screeningID)
    if err != nil {
        return fmt.Errorf("screening not found: %w", err)
    }
    
    // Validate workflow sequence
    if !isValidPartSequence(screening.CurrentPart, part) {
        return fmt.Errorf("invalid workflow sequence: current part is %s, cannot jump to %s", 
            screening.CurrentPart, part)
    }
    
    // Process based on part
    switch part {
    case "a":
        return s.processPartA(ctx, screening, data.(PartAData))
    case "b":
        return s.processPartB(ctx, screening, data.(PartBData))
    case "c":
        return s.processPartC(ctx, screening, data.(PartCData))
    case "d":
        return s.processEnrollment(ctx, screening, data.(EnrollmentData))
    default:
        return fmt.Errorf("unknown part: %s", part)
    }
}

func (s *ScreeningService) processPartB(ctx context.Context, screening *NCDScreening, data PartBData) error {
    // Calculate CVD risk based on B-4 questions
    if hasCVDWarningSigns(data.CVDQuestions) {
        // Immediate referral for suspected heart attack/stroke
        s.queueEmergencyReferral(ctx, screening.PatientID, "Suspected CVD emergency")
        screening.RequiresEmergencyReferral = true
    }
    
    // Calculate BMI
    if data.HeightCM > 0 && data.WeightKG > 0 {
        screening.BMI = calculateBMI(data.WeightKG, data.HeightCM)
    }
    
    // Update screening record
    screening.ChestPain = data.ChestPain
    screening.Breathlessness = data.Breathlessness
    screening.SystolicBP = data.SystolicBP
    screening.DiastolicBP = data.DiastolicBP
    screening.CurrentPart = "c" // Move to next part
    screening.UpdatedAt = time.Now()
    
    return s.repo.Update(ctx, screening)
}
```

## 7. FHIR INTEGRATION

### 7.1 FHIR Resource Mapping
```go
package fhir

import (
    "github.com/google/fhir/go/jsonformat"
    fhirpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/patient_go_proto"
)

type FHIRService struct {
    converter *jsonformat.Converter
    client    fhirclient.Client
}

func (s *FHIRService) ConvertToFHIRPatient(patient *domain.Patient) (*fhirpb.Patient, error) {
    fhirPatient := &fhirpb.Patient{
        Id: &datatypes_go_proto.Id{Value: patient.ID.String()},
        Identifier: []*datatypes_go_proto.Identifier{
            {
                System: &datatypes_go_proto.Uri{Value: "urn:zarish:patient"},
                Value:  &datatypes_go_proto.String{Value: patient.RegistrationNumber},
            },
        },
        Name: []*datatypes_go_proto.HumanName{
            {
                Given:  []*datatypes_go_proto.String{{Value: patient.FirstName}},
                Family: &datatypes_go_proto.String{Value: patient.LastName},
            },
        },
        Gender: getFHIRGender(patient.Gender),
        BirthDate: &datatypes_go_proto.Date{
            Precision: datatypes_go_proto.Date_DAY,
            ValueUs:   patient.DateOfBirth.UnixNano() / 1000,
        },
    }
    
    // Add address based on nationality
    address := s.createFHIRAddress(patient)
    if address != nil {
        fhirPatient.Address = []*datatypes_go_proto.Address{address}
    }
    
    return fhirPatient, nil
}

func (s *FHIRService) createFHIRAddress(patient *domain.Patient) *datatypes_go_proto.Address {
    if patient.Nationality == "FDNM" {
        return &datatypes_go_proto.Address{
            Line: []*datatypes_go_proto.String{
                {Value: fmt.Sprintf("Camp %s, Block %s", patient.Camp, patient.Block)},
                {Value: fmt.Sprintf("HH: %s", patient.HHNumber)},
            },
            City: &datatypes_go_proto.String{Value: "Cox's Bazar"},
            Country: &datatypes_go_proto.String{Value: "Bangladesh"},
        }
    } else if patient.Nationality == "BD" {
        return &datatypes_go_proto.Address{
            Line: []*datatypes_go_proto.String{
                {Value: patient.Village},
                {Value: patient.Union},
            },
            District: &datatypes_go_proto.String{Value: patient.District},
            State: &datatypes_go_proto.String{Value: patient.Upazila},
            Country: &datatypes_go_proto.String{Value: "Bangladesh"},
        }
    }
    return nil
}
```

## 8. DATABASE MIGRATIONS

### 8.1 Migration Structure
```sql
-- migrations/postgres/001_create_patients_table.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_number VARCHAR(50) UNIQUE NOT NULL,
    nationality VARCHAR(10) NOT NULL CHECK (nationality IN ('FDNM', 'BD')),
    
    -- Personal information
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    date_of_birth DATE NOT NULL,
    gender VARCHAR(20) NOT NULL CHECK (gender IN ('male', 'female', 'other')),
    mobile_number VARCHAR(20),
    
    -- FDMN fields (nullable, used when nationality = 'FDNM')
    camp VARCHAR(50),
    block VARCHAR(50),
    sub_block VARCHAR(50),
    hh_number VARCHAR(50),
    fcn_number VARCHAR(50),
    progress_id VARCHAR(50),
    
    -- Bangladeshi fields (nullable, used when nationality = 'BD')
    district VARCHAR(50),
    upazila VARCHAR(50),
    union_name VARCHAR(50),
    village VARCHAR(100),
    nid_number VARCHAR(50),
    
    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes for performance
CREATE INDEX idx_patients_nationality ON patients(nationality);
CREATE INDEX idx_patients_registration_number ON patients(registration_number);
CREATE INDEX idx_patients_deleted_at ON patients(deleted_at) WHERE deleted_at IS NULL;

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_patients_updated_at 
    BEFORE UPDATE ON patients 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
```

## 9. DEPLOYMENT & DEVOPS

### 9.1 Docker Configuration
```dockerfile
# Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080
ENV TZ=Asia/Dhaka
CMD ["./main"]
```

### 9.2 Docker Compose
```yaml
# docker-compose.yml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: zarish_his
      POSTGRES_USER: zarish_user
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./migrations/postgres:/docker-entrypoint-initdb.d
    ports:
      - "5432:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U zarish_user"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7.2-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes

  rabbitmq:
    image: rabbitmq:3.13-management-alpine
    environment:
      RABBITMQ_DEFAULT_USER: admin
      RABBITMQ_DEFAULT_PASS: ${RABBITMQ_PASSWORD}
    ports:
      - "5672:5672"
      - "15672:15672"

  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://zarish_user:${DB_PASSWORD}@postgres:5432/zarish_his
      - REDIS_URL=redis://redis:6379
      - RABBITMQ_URL=amqp://admin:${RABBITMQ_PASSWORD}@rabbitmq:5672/
      - TZ=Asia/Dhaka
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_started
      rabbitmq:
        condition: service_started
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  worker:
    build: .
    command: ["./worker"]
    environment:
      - DATABASE_URL=postgres://zarish_user:${DB_PASSWORD}@postgres:5432/zarish_his
      - REDIS_URL=redis://redis:6379
      - RABBITMQ_URL=amqp://admin:${RABBITMQ_PASSWORD}@rabbitmq:5672/
    depends_on:
      - postgres
      - redis
      - rabbitmq

volumes:
  postgres_data:
  redis_data:
```

## 10. TESTING STRATEGY

### 10.1 Unit Tests
```go
// internal/service/patient_service_test.go
package service_test

import (
    "context"
    "testing"
    "time"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestPatientService_RegisterPatient_FDMN(t *testing.T) {
    // Setup
    mockRepo := new(MockPatientRepository)
    mockQueue := new(MockQueueClient)
    service := NewPatientService(mockRepo, mockQueue)
    
    ctx := context.Background()
    req := PatientRegistrationRequest{
        Nationality: "FDNM",
        FirstName:   "Mohammed",
        LastName:    "Hossain",
        DateOfBirth: "15 June 1985",
        Gender:      "male",
        Camp:        stringPtr("5"),
        Block:       stringPtr("C"),
        HHNumber:    stringPtr("123"),
    }
    
    // Expectations
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Patient")).
        Return(nil)
    mockQueue.On("Publish", mock.Anything, "patient.registered", mock.Anything).
        Return(nil)
    
    // Execute
    resp, err := service.RegisterPatient(ctx, req)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, resp)
    assert.NotEmpty(t, resp.ID)
    assert.Contains(t, resp.RegistrationNumber, "REG-")
    assert.Equal(t, "Patient registered successfully", resp.Message)
    
    mockRepo.AssertExpectations(t)
    mockQueue.AssertExpectations(t)
}
```

## 11. MONITORING & OBSERVABILITY

### 11.1 Metrics Collection
```go
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // Patient metrics
    patientsRegistered = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "zarish_patients_registered_total",
        Help: "Total number of patients registered",
    }, []string{"nationality", "gender"})
    
    // NCD Screening metrics
    ncdScreeningsCompleted = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "zarish_ncd_screenings_completed_total",
        Help: "Total number of NCD screenings completed",
    }, []string{"outcome"})
    
    // API metrics
    apiRequests = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "zarish_api_requests_total",
        Help: "Total number of API requests",
    }, []string{"method", "endpoint", "status_code"})
    
    apiRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        Name: "zarish_api_request_duration_seconds",
        Help: "API request duration in seconds",
        Buckets: prometheus.DefBuckets,
    }, []string{"method", "endpoint"})
)
```

### 11.2 Health Check Endpoint
```go
// internal/api/handlers/health.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type HealthHandler struct {
    db     *gorm.DB
    redis  *redis.Client
}

func (h *HealthHandler) Check(c *gin.Context) {
    health := gin.H{
        "status": "healthy",
        "timestamp": time.Now().Format("02 January 2006 03:04:05 PM"),
        "timezone": "Asia/Dhaka",
        "services": make(map[string]string),
    }
    
    // Check database
    sqlDB, err := h.db.DB()
    if err != nil || sqlDB.Ping() != nil {
        health["services"]["database"] = "unhealthy"
        health["status"] = "degraded"
    } else {
        health["services"]["database"] = "healthy"
    }
    
    // Check Redis
    if _, err := h.redis.Ping(c).Result(); err != nil {
        health["services"]["redis"] = "unhealthy"
        health["status"] = "degraded"
    } else {
        health["services"]["redis"] = "healthy"
    }
    
    statusCode := http.StatusOK
    if health["status"] == "degraded" {
        statusCode = http.StatusServiceUnavailable
    }
    
    c.JSON(statusCode, health)
}
```

## 12. SECURITY IMPLEMENTATION

### 12.1 JWT Authentication
```go
package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
    secretKey     string
    tokenDuration time.Duration
}

func (m *JWTManager) GenerateToken(userID string, roles []string) (string, error) {
    claims := Claims{
        UserID: userID,
        Roles:  roles,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "zarish-his",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) VerifyToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(m.secretKey), nil
    })
    
    if err != nil {
        return nil, err
    }
    
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    
    return claims, nil
}
```

## 13. API DOCUMENTATION

### 13.1 OpenAPI/Swagger Setup
```go
// docs/swagger.go
package docs

import _ "zarish-his/docs"

// @title Zarish-HIS API
// @version 1.0
// @description Health Information System for humanitarian settings
// @termsOfService http://zarish.org/terms/

// @contact.name API Support
// @contact.url http://zarish.org/support
// @contact.email support@zarish.org

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host api.zarish.org
// @BasePath /api/v1
// @schemes https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token

// Patient Registration
// @Summary Register a new patient
// @Description Register a new patient with nationality-based fields
// @Tags patients
// @Accept json
// @Produce json
// @Param request body PatientRegistrationRequest true "Patient registration data"
// @Success 201 {object} PatientResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /patients [post]
// @Security BearerAuth
func registerPatient() {}
```

## 14. ENVIRONMENT CONFIGURATION

```go
// internal/config/config.go
package config

import (
    "github.com/spf13/viper"
    "time"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
    RabbitMQ RabbitMQConfig
    JWT      JWTConfig
    FHIR     FHIRConfig
}

type ServerConfig struct {
    Port            string        `mapstructure:"PORT"`
    Timezone        string        `mapstructure:"TZ"`
    ReadTimeout     time.Duration `mapstructure:"READ_TIMEOUT"`
    WriteTimeout    time.Duration `mapstructure:"WRITE_TIMEOUT"`
    ShutdownTimeout time.Duration `mapstructure:"SHUTDOWN_TIMEOUT"`
}

type DatabaseConfig struct {
    URL                string `mapstructure:"DATABASE_URL"`
    MaxOpenConnections int    `mapstructure:"DB_MAX_OPEN_CONNS"`
    MaxIdleConnections int    `mapstructure:"DB_MAX_IDLE_CONNS"`
    ConnMaxLifetime    time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

type JWTConfig struct {
    SecretKey     string        `mapstructure:"JWT_SECRET"`
    TokenDuration time.Duration `mapstructure:"JWT_DURATION"`
}

func LoadConfig(path string) (*Config, error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    
    viper.AutomaticEnv()
    
    // Set defaults
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("server.timezone", "Asia/Dhaka")
    viper.SetDefault("database.max_open_connections", 25)
    viper.SetDefault("database.max_idle_connections", 5)
    viper.SetDefault("jwt.token_duration", "24h")
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
```

## 15. MUST FOLLOW RULES

### 15.1 Date & Time Rules
1. **Date Storage:** Always store dates in UTC in database
2. **Date Display:** Convert to Asia/Dhaka (GMT+6) for display
3. **Format:** API accepts/returns "DD MMMM YYYY" (15 June 1985)
4. **Time Format:** 12-hour with AM/PM (02:30 PM)
5. **Birth Dates:** Store exact date, calculate age dynamically

### 15.2 Nationality Rules
1. **First Question:** Always ask nationality first in registration
2. **Conditional Fields:** Show FDMN fields for "FDNM", BD fields for "BD"
3. **Validation:** Validate required fields based on nationality
4. **Data Integrity:** Enforce at database level with CHECK constraints

### 15.3 FHIR Compliance
1. **Resource Mapping:** All entities must map to FHIR resources
2. **Validation:** Validate FHIR resources before storage
3. **Versioning:** Use FHIR R4 standard
4. **Extensions:** Use custom extensions for Zarish-specific data

### 15.4 Security Rules
1. **PHI Encryption:** Encrypt all personally identifiable information
2. **Audit Logging:** Log all data access and modifications
3. **Role-Based Access:** Implement RBAC for all endpoints
4. **Rate Limiting:** Protect APIs from abuse
5. **Input Validation:** Validate all user inputs, prevent SQL injection

### 15.5 Code Quality Rules
1. **Test Coverage:** Minimum 80% test coverage for critical paths
2. **Code Review:** All changes require code review
3. **Static Analysis:** Run golangci-lint on every commit
4. **Documentation:** Update OpenAPI docs with API changes
5. **Error Handling:** Proper error wrapping and logging

## 16. PATIENT JOURNEY IMPLEMENTATION

### 16.1 New Patient Registration Flow
```go
// internal/service/workflow/patient_journey.go
package workflow

type PatientJourneyService struct {
    patientService    PatientService
    screeningService  ScreeningService
    encounterService  EncounterService
    labService        LabService
    pharmacyService   PharmacyService
}

func (s *PatientJourneyService) ProcessNewPatient(ctx context.Context, req PatientRegistrationRequest) (*JourneyResponse, error) {
    // Step 1: Register patient
    patient, err := s.patientService.RegisterPatient(ctx, req)
    if err != nil {
        return nil, fmt.Errorf("registration failed: %w", err)
    }
    
    // Step 2: Start NCD screening
    screening, err := s.screeningService.StartScreening(ctx, patient.ID)
    if err != nil {
        return nil, fmt.Errorf("screening initialization failed: %w", err)
    }
    
    // Step 3: Based on screening results, route to appropriate service
    journey := &JourneyResponse{
        PatientID:      patient.ID,
        ScreeningID:    screening.ID,
        CurrentStep:    "registration_complete",
        NextSteps:      []JourneyStep{},
        EstimatedTime:  30, // minutes
    }
    
    // Check for emergency signs
    if screening.RequiresEmergencyReferral {
        journey.NextSteps = append(journey.NextSteps, JourneyStep{
            Service:    "emergency",
            Location:   "ER Department",
            Priority:   "immediate",
            Instructions: "Proceed directly to Emergency Room",
        })
    } else {
        // Normal routing based on screening outcome
        if screening.NeedsNCDEnrollment {
            journey.NextSteps = append(journey.NextSteps, JourneyStep{
                Service:    "ncd_clinic",
                Location:   "NCD Clinic, Room 201",
                Priority:   "normal",
                Instructions: "Wait for consultation with NCD specialist",
            })
        } else {
            journey.NextSteps = append(journey.NextSteps, JourneyStep{
                Service:    "gopd",
                Location:   "General OPD, Room 101",
                Priority:   "normal",
                Instructions: "Wait for general consultation",
            })
        }
    }
    
    // Generate QR code for patient journey tracking
    qrCode, err := s.generateJourneyQRCode(journey)
    if err != nil {
        return nil, fmt.Errorf("failed to generate QR code: %w", err)
    }
    
    journey.QRCode = qrCode
    return journey, nil
}
```

## 17. DEPLOYMENT CHECKLIST

### 17.1 Pre-Deployment
- [ ] Database migrations tested and ready
- [ ] Environment variables configured
- [ ] SSL certificates installed
- [ ] Backup strategy in place
- [ ] Monitoring alerts configured

### 17.2 Deployment Steps
1. **Database:** Run migrations
2. **Cache:** Clear Redis cache if needed
3. **Application:** Deploy new version with zero-downtime
4. **Health Check:** Verify all services are healthy
5. **Smoke Tests:** Run critical path tests

### 17.3 Post-Deployment
- [ ] Verify API endpoints are responding
- [ ] Check error rates in logs
- [ ] Monitor performance metrics
- [ ] Update API documentation if changed

## 18. DISASTER RECOVERY

### 18.1 Backup Strategy
```yaml
# Backup configuration
backup:
  database:
    schedule: "0 2 * * *"  # Daily at 2 AM
    retention: 30 days
    location: s3://zarish-backups/database/
  
  files:
    schedule: "0 3 * * *"  # Daily at 3 AM
    retention: 90 days
    location: s3://zarish-backups/files/
  
  logs:
    schedule: "0 4 * * *"  # Daily at 4 AM
    retention: 180 days
    location: s3://zarish-backups/logs/
```

### 18.2 Recovery Procedures
1. **Database Corruption:** Restore from latest backup
2. **Data Loss:** Restore from backup, replay transaction logs
3. **System Failure:** Redeploy from infrastructure as code
4. **Security Breach:** Isolate system, investigate, restore clean backup

## 19. PERFORMANCE OPTIMIZATION

### 19.1 Database Optimization
```sql
-- Create performance indexes
CREATE INDEX CONCURRENTLY idx_patients_created_at ON patients(created_at);
CREATE INDEX CONCURRENTLY idx_screenings_patient_date ON ncd_screenings(patient_id, visit_date DESC);

-- Partition large tables
CREATE TABLE ncd_screenings_2024 PARTITION OF ncd_screenings
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
```

### 19.2 Caching Strategy
```go
// Cache frequently accessed data
func (s *PatientService) GetPatientByID(ctx context.Context, id uuid.UUID) (*Patient, error) {
    cacheKey := fmt.Sprintf("patient:%s", id.String())
    
    // Try cache first
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var patient Patient
        if err := json.Unmarshal([]byte(cached), &patient); err == nil {
            return &patient, nil
        }
    }
    
    // Cache miss, get from database
    patient, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Cache for 5 minutes
    data, _ := json.Marshal(patient)
    s.cache.Set(ctx, cacheKey, string(data), 5*time.Minute)
    
    return patient, nil
}
```

## 20. CONTINUOUS INTEGRATION

### 20.1 GitHub Actions Workflow
```yaml
# .github/workflows/ci-cd.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:16-alpine
        env:
          POSTGRES_DB: test_db
          POSTGRES_USER: test_user
          POSTGRES_PASSWORD: test_pass
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.22'
      - name: Run tests
        env:
          DATABASE_URL: postgres://test_user:test_pass@localhost:5432/test_db
        run: |
          go mod download
          go test ./... -v -coverprofile=coverage.out
          go tool cover -func=coverage.out
  
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.55
  
  build:
    needs: [test, lint]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Build Docker image
        run: |
          docker build -t zarish-his:${{ github.sha }} .
          docker build -t zarish-his:latest .
  
  deploy:
    needs: build
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - name: Deploy to production
        run: |
          # Deployment commands here
          echo "Deploying to production..."
```

---

## SUMMARY OF KEY TECHNICAL DECISIONS

1. **Architecture:** Clean Architecture with Domain-Driven Design
2. **Database:** PostgreSQL 16+ with UUID primary keys
3. **Caching:** Redis for session and frequently accessed data
4. **Queue:** RabbitMQ for async job processing
5. **Monitoring:** Prometheus + Grafana + OpenTelemetry
6. **Deployment:** Docker + Docker Compose for development, Kubernetes for production
7. **Security:** JWT tokens, role-based access control, PHI encryption
8. **FHIR Compliance:** Google FHIR Go library with custom extensions
9. **Testing:** Comprehensive unit, integration, and E2E tests
10. **CI/CD:** GitHub Actions with automated testing and deployment

---

*Developer: Arwa Zarish Technology*
*The Code and The Brain*
*Email : zarishsphere@gmail.com*

*This document is maintained by the Zarish Sphere Team*  
*Last Updated: 01 January 2026*  
*Version: 1.0.0*  
*Technology Stack: Go 1.23, Gin, GORM, PostgreSQL 16*