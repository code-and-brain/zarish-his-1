import { Activity, AlertCircle, Calendar, Clock, FileText } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import {
  PortalService,
  type PortalDashboardData,
} from '../../services/portalService';

const PortalDashboard: React.FC = () => {
  const [data, setData] = useState<PortalDashboardData | null>(null);
  const [loading, setLoading] = useState(true);

  // Hardcoded patient ID for MVP
  const patientId = 1;

  useEffect(() => {
    fetchDashboard();
  }, []);

  const fetchDashboard = async () => {
    try {
      setLoading(true);
      const dashboardData = await PortalService.getDashboard(patientId);
      setData(dashboardData);
    } catch (error) {
      console.error('Error fetching dashboard:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading)
    return <div className="p-8 text-center">Loading dashboard...</div>;
  if (!data)
    return <div className="p-8 text-center">Failed to load dashboard</div>;

  return (
    <div className="p-8 max-w-6xl mx-auto">
      <header className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900">
          Welcome back, {data.patient.given_name}
        </h1>
        <p className="text-gray-500 mt-2">
          Here's what's happening with your health today.
        </p>
      </header>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex items-center gap-4">
            <div className="p-3 bg-indigo-50 text-indigo-600 rounded-lg">
              <Calendar size={24} />
            </div>
            <div>
              <p className="text-sm text-gray-500">Next Appointment</p>
              <p className="text-lg font-semibold text-gray-900">
                {data.appointments.length > 0
                  ? new Date(
                      data.appointments[0].scheduled_start
                    ).toLocaleDateString()
                  : 'None scheduled'}
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex items-center gap-4">
            <div className="p-3 bg-green-50 text-green-600 rounded-lg">
              <Activity size={24} />
            </div>
            <div>
              <p className="text-sm text-gray-500">Recent Vitals</p>
              <p className="text-lg font-semibold text-gray-900">Stable</p>
            </div>
          </div>
        </div>

        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex items-center gap-4">
            <div className="p-3 bg-blue-50 text-blue-600 rounded-lg">
              <FileText size={24} />
            </div>
            <div>
              <p className="text-sm text-gray-500">New Reports</p>
              <p className="text-lg font-semibold text-gray-900">2 Available</p>
            </div>
          </div>
        </div>
      </div>

      {/* Upcoming Appointments */}
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden mb-8">
        <div className="p-6 border-b border-gray-200 flex justify-between items-center">
          <h2 className="text-lg font-semibold text-gray-900">
            Upcoming Appointments
          </h2>
          <button className="text-indigo-600 text-sm font-medium hover:text-indigo-700">
            View All
          </button>
        </div>
        <div className="divide-y divide-gray-200">
          {data.appointments.length === 0 ? (
            <div className="p-6 text-center text-gray-500">
              No upcoming appointments
            </div>
          ) : (
            data.appointments.map((apt) => (
              <div
                key={apt.id}
                className="p-6 flex items-center justify-between hover:bg-gray-50"
              >
                <div className="flex items-center gap-4">
                  <div className="text-center w-16">
                    <p className="text-sm font-bold text-gray-900">
                      {new Date(apt.scheduled_start).toLocaleDateString(
                        undefined,
                        {
                          month: 'short',
                        }
                      )}
                    </p>
                    <p className="text-2xl font-bold text-gray-900">
                      {new Date(apt.scheduled_start).getDate()}
                    </p>
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">
                      {apt.appointment_type} Visit
                    </p>
                    <div className="flex items-center gap-2 text-sm text-gray-500 mt-1">
                      <Clock size={14} />
                      {new Date(apt.scheduled_start).toLocaleTimeString([], {
                        hour: '2-digit',
                        minute: '2-digit',
                      })}
                    </div>
                  </div>
                </div>
                <span
                  className={`px-3 py-1 rounded-full text-xs font-medium ${
                    apt.status === 'confirmed'
                      ? 'bg-green-100 text-green-800'
                      : 'bg-yellow-100 text-yellow-800'
                  }`}
                >
                  {apt.status}
                </span>
              </div>
            ))
          )}
        </div>
      </div>

      {/* Health Alerts */}
      {data.alerts.length > 0 && (
        <div className="bg-yellow-50 border border-yellow-200 rounded-xl p-6">
          <div className="flex items-start gap-3">
            <AlertCircle className="text-yellow-600 mt-0.5" size={20} />
            <div>
              <h3 className="text-sm font-medium text-yellow-800">
                Health Reminders
              </h3>
              <ul className="mt-2 space-y-1">
                {data.alerts.map((alert, idx) => (
                  <li key={idx} className="text-sm text-yellow-700">
                    • {alert}
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default PortalDashboard;
