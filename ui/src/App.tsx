
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import PatientList from './pages/PatientList';
import PatientRegistrationForm from './pages/PatientRegistrationForm';
import PatientDetail from './pages/PatientDetail';

import AppointmentsPage from './pages/AppointmentsPage';
import LabOrdersPage from './pages/LabOrdersPage';
import InpatientManagement from './pages/InpatientManagement';
import ReportsDashboard from './pages/ReportsDashboard';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        {/* Navigation Bar */}
        <nav className="bg-white shadow-sm border-b border-gray-200">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex justify-between h-16">
              <div className="flex">
                <div className="flex-shrink-0 flex items-center">
                  <span className="text-xl font-bold text-blue-600">Zarish HIS Module</span>
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
            <Route path="/patients/new" element={<PatientRegistrationForm />} />
            <Route path="/patients/:id" element={<PatientDetail />} />
            <Route path="/appointments" element={<AppointmentsPage />} />
            <Route path="/labs" element={<LabOrdersPage />} />
            <Route path="/inpatient" element={<InpatientManagement />} />
            <Route path="/reports" element={<ReportsDashboard />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
