### Zarish-HIS Project To-Do List

#### Phase 1: Core Infrastructure (Completed)

- [x] Set up monorepo structure
- [x] Configure Firebase projects
- [x] Deploy FHIR server
- [x] Set up terminology server

#### Phase 2: Patient Management

- [ ] **Patient Registration**
    - [ ] Create patient registration form (Part A) as a React Hook Form component
    - [ ] Implement nationality-first logic in the registration form
    - [ ] Create backend endpoint for patient registration (`POST /api/v1/patients`)
    - [ ] Implement patient registration service with validation and data mapping to the `Patient` FHIR resource
    - [ ] Generate temporary Patient ID upon registration
- [ ] **NCD Patient Screening**
    - [ ] Create NCD screening form (Parts B, C, and D) as React Hook Form components
    - [ ] Create backend endpoints for NCD screening (`POST /api/v1/ncd-screenings`, etc.)
    - [ ] Implement NCD screening service with data mapping to `QuestionnaireResponse` and `Observation` FHIR resources
- [ ] **Patient Journey Workflows**
    - [ ] Implement the new patient registration flow
    - [ ] Implement the follow-up patient flow
    - [ ] Implement the ER patient flow
- [ ] **QR Code Integration**
    - [ ] Implement QR code generation for patient identification
    - [ ] Integrate QR code scanning for patient lookup

#### Phase 3: Service Integration

- [ ] **Laboratory**
    - [ ] Create lab order form
    - [ ] Create backend endpoint for creating lab orders (`POST /api/v1/lab-orders`)
    - [ ] Implement lab service with data mapping to the `ServiceRequest` FHIR resource
    - [ ] Implement interface for receiving `DiagnosticReport` from the lab system
- [ ] **Pharmacy**
    - [ ] Create prescription form
    - [ ] Create backend endpoint for creating prescriptions (`POST /api/v1/prescriptions`)
    - [ ] Implement pharmacy service with data mapping to the `MedicationRequest` FHIR resource
    - [ ] Implement interface for recording `MedicationDispense`
- [ ] **Referrals**
    - [ ] Create referral form
    - [ ] Create backend endpoint for creating referrals (`POST /api/v1/referrals`)
    - [ ] Implement referral service with data mapping to the `ReferralRequest` FHIR resource
- [ ] **Reporting**
    - [ ] Create reporting dashboard
    - [ ] Implement backend service for generating reports

#### Phase 4: DevOps & Deployment

- [ ] **CI/CD Pipeline**
    - [ ] Create GitHub Actions workflow for CI/CD
    - [ ] Implement automated linting, testing, and building
    - [ ] Configure automated deployment to Firebase for the frontend
    - [ ] Configure automated deployment of the FHIR server
- [ ] **Testing**
    - [ ] Write unit tests for all backend services
    - [ ] Write integration tests for all API endpoints
    - [ ] Write end-to-end tests for the patient journey workflows
- [ ] **Monitoring**
    - [ ] Configure Firebase Analytics for the frontend
    - [ ] Configure logging with Winston for the backend
    - [ ] Configure FHIR audit events

#### Phase 5: Documentation & Training

- [ ] **Documentation**
    - [ ] Create user manuals in English and Bengali
    - [ ] Create video tutorials for the user workflows
    - [ ] Create a troubleshooting guide
    - [ ] Generate API documentation
- [ ] **Training**
    - [ ] Develop training materials for registration staff, clinical staff, and administrative staff