import type { Appointment, Patient } from '../types';
import api from './api';

// Define types for portal responses
export interface PortalDashboardData {
  patient: Patient;
  appointments: Appointment[];
  alerts: string[];
}

export interface PortalRecordsData {
  clinical_notes: any[];
  prescriptions: any[];
  lab_orders: any[];
}

export const PortalService = {
  // Get dashboard data
  getDashboard: async (patientId: number): Promise<PortalDashboardData> => {
    const response = await api.get<PortalDashboardData>('/portal/dashboard', {
      params: { patient_id: patientId },
    });
    return response.data;
  },

  // Get appointments
  getAppointments: async (
    patientId: number
  ): Promise<{ data: Appointment[]; total: number }> => {
    const response = await api.get<{ data: Appointment[]; total: number }>(
      '/portal/appointments',
      {
        params: { patient_id: patientId },
      }
    );
    return response.data;
  },

  // Get medical records
  getRecords: async (patientId: number): Promise<PortalRecordsData> => {
    const response = await api.get<PortalRecordsData>('/portal/records', {
      params: { patient_id: patientId },
    });
    return response.data;
  },
};
