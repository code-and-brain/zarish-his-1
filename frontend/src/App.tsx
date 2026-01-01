import {
  Link,
  Navigate,
  Route,
  BrowserRouter as Router,
  Routes,
} from 'react-router-dom';
import PatientDetail from './pages/PatientDetail';
import PatientList from './pages/PatientList';
import PatientRegistrationForm from './pages/PatientRegistrationForm';

import AppointmentsPage from './pages/AppointmentsPage';
import InpatientManagement from './pages/InpatientManagement';
import LabOrdersPage from './pages/LabOrdersPage';
import PharmacyDashboard from './pages/PharmacyDashboard';
import ReportsDashboard from './pages/ReportsDashboard';

import BillingDashboard from './pages/BillingDashboard';
import ClaimManagement from './pages/ClaimManagement';
import InvoiceManagement from './pages/InvoiceManagement';
import PaymentEntry from './pages/PaymentEntry';
import RadiologyDashboard from './pages/RadiologyDashboard';
import StudyViewer from './pages/StudyViewer';

import { AuthProvider } from './context/AuthContext';
import PortalLayout from './layouts/PortalLayout';
import PortalAppointments from './pages/portal/PortalAppointments';
import PortalDashboard from './pages/portal/PortalDashboard';
import PortalRecords from './pages/portal/PortalRecords';

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          {/* Patient Portal Routes - Distinct Layout */}
          <Route path="/portal" element={<PortalLayout />}>
            <Route
              index
              element={<Navigate to="/portal/dashboard" replace />}
            />
            <Route path="dashboard" element={<PortalDashboard />} />
            <Route path="appointments" element={<PortalAppointments />} />
            <Route path="records" element={<PortalRecords />} />
          </Route>

          {/* Main Clinical App Routes */}
          <Route
            path="*"
            element={
              <div className="min-h-screen bg-gray-100">
                {/* Navigation Bar */}
                <nav className="bg-white shadow-sm border-b border-gray-200">
                  <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div className="flex justify-between h-16">
                      <div className="flex">
                        <div className="flex-shrink-0 flex items-center">
                          <span className="text-xl font-bold text-blue-600">
                            Zarish HIS Module
                          </span>
                        </div>
                        <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
                          <Link
                            to="/patients"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Patients
                          </Link>
                          <Link
                            to="/appointments"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Appointments
                          </Link>
                          <Link
                            to="/labs"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Lab Orders
                          </Link>
                          <Link
                            to="/inpatient"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Inpatient
                          </Link>
                          <Link
                            to="/reports"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Reports
                          </Link>
                          <Link
                            to="/pharmacy"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Pharmacy
                          </Link>
                          <Link
                            to="/billing"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Billing
                          </Link>
                          <Link
                            to="/radiology"
                            className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                          >
                            Radiology
                          </Link>
                        </div>
                      </div>
                      <div className="flex items-center">
                        <button className="bg-blue-600 text-white px-3 py-1 rounded-md text-sm">
                          Login
                        </button>
                      </div>
                    </div>
                  </div>
                </nav>

                {/* Main Content */}
                <main className="py-6">
                  <Routes>
                    <Route path="/" element={<PatientList />} />
                    <Route path="/patients" element={<PatientList />} />
                    <Route
                      path="/patients/new"
                      element={<PatientRegistrationForm />}
                    />
                    <Route path="/patients/:id" element={<PatientDetail />} />
                    <Route
                      path="/appointments"
                      element={<AppointmentsPage />}
                    />
                    <Route path="/labs" element={<LabOrdersPage />} />
                    <Route
                      path="/inpatient"
                      element={<InpatientManagement />}
                    />
                    <Route path="/pharmacy" element={<PharmacyDashboard />} />

                    {/* Billing Routes */}
                    <Route path="/billing" element={<BillingDashboard />} />
                    <Route
                      path="/billing/invoices"
                      element={<BillingDashboard />}
                    />
                    <Route
                      path="/billing/invoices/new"
                      element={<InvoiceManagement />}
                    />
                    <Route
                      path="/billing/invoices/:id"
                      element={<InvoiceManagement />}
                    />
                    <Route
                      path="/billing/invoices/:invoiceId/pay"
                      element={<PaymentEntry />}
                    />
                    <Route
                      path="/billing/claims"
                      element={<ClaimManagement />}
                    />
                    <Route
                      path="/billing/claims/new"
                      element={<ClaimManagement />}
                    />

                    {/* Radiology Routes */}
                    <Route path="/radiology" element={<RadiologyDashboard />} />
                    <Route
                      path="/radiology/studies/:id"
                      element={<StudyViewer />}
                    />

                    <Route path="/reports" element={<ReportsDashboard />} />
                  </Routes>
                </main>
              </div>
            }
          />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
