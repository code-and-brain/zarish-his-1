import { Calendar, Clock, MapPin, Plus } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { PortalService } from '../../services/portalService';
import type { Appointment } from '../../types';

const PortalAppointments: React.FC = () => {
  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [loading, setLoading] = useState(true);

  // Hardcoded patient ID for MVP
  const patientId = 1;

  useEffect(() => {
    fetchAppointments();
  }, []);

  const fetchAppointments = async () => {
    try {
      setLoading(true);
      const result = await PortalService.getAppointments(patientId);
      setAppointments(result.data);
    } catch (error) {
      console.error('Error fetching appointments:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-8 max-w-6xl mx-auto">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">My Appointments</h1>
          <p className="text-gray-500">
            Manage your upcoming visits and view history
          </p>
        </div>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 flex items-center gap-2">
          <Plus size={20} />
          Book New
        </button>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        {loading ? (
          <div className="p-8 text-center text-gray-500">
            Loading appointments...
          </div>
        ) : appointments.length === 0 ? (
          <div className="p-12 text-center">
            <Calendar size={48} className="mx-auto text-gray-300 mb-4" />
            <h3 className="text-lg font-medium text-gray-900">
              No appointments found
            </h3>
            <p className="text-gray-500 mt-2">
              You don't have any upcoming appointments scheduled.
            </p>
          </div>
        ) : (
          <div className="divide-y divide-gray-200">
            {appointments.map((apt) => (
              <div
                key={apt.id}
                className="p-6 hover:bg-gray-50 transition-colors"
              >
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                  <div className="flex items-start gap-4">
                    <div className="w-14 h-14 bg-indigo-50 rounded-lg flex flex-col items-center justify-center text-indigo-600 flex-shrink-0">
                      <span className="text-xs font-bold uppercase">
                        {new Date(apt.scheduled_start).toLocaleDateString(
                          undefined,
                          { month: 'short' }
                        )}
                      </span>
                      <span className="text-xl font-bold">
                        {new Date(apt.scheduled_start).getDate()}
                      </span>
                    </div>
                    <div>
                      <h3 className="font-semibold text-gray-900 text-lg">
                        {apt.appointment_type} Visit
                      </h3>
                      <div className="flex items-center gap-4 mt-2 text-sm text-gray-500">
                        <div className="flex items-center gap-1">
                          <Clock size={16} />
                          {new Date(apt.scheduled_start).toLocaleTimeString(
                            [],
                            {
                              hour: '2-digit',
                              minute: '2-digit',
                            }
                          )}
                        </div>
                        <div className="flex items-center gap-1">
                          <MapPin size={16} />
                          Main Clinic
                        </div>
                      </div>
                    </div>
                  </div>

                  <div className="flex items-center gap-3">
                    <span
                      className={`px-3 py-1 rounded-full text-xs font-medium ${
                        apt.status === 'confirmed'
                          ? 'bg-green-100 text-green-800'
                          : apt.status === 'cancelled'
                          ? 'bg-red-100 text-red-800'
                          : 'bg-yellow-100 text-yellow-800'
                      }`}
                    >
                      {apt.status.charAt(0).toUpperCase() + apt.status.slice(1)}
                    </span>
                    <button className="text-gray-400 hover:text-gray-600 p-2">
                      Details
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default PortalAppointments;
