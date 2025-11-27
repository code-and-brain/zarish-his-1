import React, { useEffect, useState } from 'react';
import { AppointmentService } from '../services/appointmentService';
import type { Appointment } from '../types';

interface Props {
  patientId: number;
}

const AppointmentManager: React.FC<Props> = ({ patientId }) => {
  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Partial<Appointment>>({
    appointment_type: 'consultation',
    status: 'scheduled'
  });

  const fetchAppointments = async () => {
    setLoading(true);
    try {
      const data = await AppointmentService.listByPatient(patientId);
      setAppointments(data);
    } catch (error) {
      console.error('Failed to fetch appointments', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAppointments();
  }, [patientId]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      // Combine date and time for start/end
      // Simplified for MVP
      await AppointmentService.create({ ...formData, patient_id: patientId });
      setShowForm(false);
      fetchAppointments();
      setFormData({ appointment_type: 'consultation', status: 'scheduled' });
    } catch (error) {
      console.error('Failed to create appointment', error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Appointments</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ New Appointment'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Type</label>
              <select
                name="appointment_type"
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.appointment_type}
                onChange={handleChange}
              >
                <option value="consultation">Consultation</option>
                <option value="follow-up">Follow-up</option>
                <option value="procedure">Procedure</option>
                <option value="vaccination">Vaccination</option>
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Date & Time</label>
              <input
                type="datetime-local"
                name="scheduled_start"
                required
                onChange={handleChange}
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
              />
            </div>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Reason</label>
            <textarea name="reason" rows={2} onChange={handleChange} className="mt-1 block w-full rounded border-gray-300 p-2 border" />
          </div>
          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            Schedule
          </button>
        </form>
      )}

      {loading ? (
        <p>Loading appointments...</p>
      ) : appointments.length === 0 ? (
        <p className="text-gray-500">No appointments found.</p>
      ) : (
        <div className="space-y-4">
          {appointments.map((apt) => (
            <div key={apt.id} className="border rounded p-4 hover:bg-gray-50 flex justify-between items-center">
              <div>
                <div className="font-medium text-blue-600">{new Date(apt.scheduled_start).toLocaleString()}</div>
                <div className="text-sm text-gray-700 capitalize">{apt.appointment_type}</div>
                {apt.reason && <div className="text-xs text-gray-500 mt-1">"{apt.reason}"</div>}
              </div>
              <div className="text-right">
                <span className={`px-2 py-1 rounded text-xs font-medium uppercase ${
                  apt.status === 'scheduled' ? 'bg-blue-100 text-blue-800' :
                  apt.status === 'completed' ? 'bg-green-100 text-green-800' :
                  apt.status === 'cancelled' ? 'bg-red-100 text-red-800' :
                  'bg-gray-100 text-gray-800'
                }`}>
                  {apt.status}
                </span>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default AppointmentManager;
