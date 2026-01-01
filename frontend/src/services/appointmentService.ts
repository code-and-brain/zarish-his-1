import api from './api';
import type { Appointment } from '../types';

export const AppointmentService = {
  create: async (appointment: Partial<Appointment>): Promise<Appointment> => {
    const response = await api.post<Appointment>('/appointments', appointment);
    return response.data;
  },

  getById: async (id: number): Promise<Appointment> => {
    const response = await api.get<Appointment>(`/appointments/${id}`);
    return response.data;
  },

  update: async (id: number, appointment: Partial<Appointment>): Promise<Appointment> => {
    const response = await api.put<Appointment>(`/appointments/${id}`, appointment);
    return response.data;
  },

  cancel: async (id: number, reason: string): Promise<Appointment> => {
    const response = await api.post<Appointment>(`/appointments/${id}/cancel`, { reason });
    return response.data;
  },

  listByDate: async (date: string): Promise<Appointment[]> => {
    const response = await api.get<Appointment[]>('/appointments', { params: { date } });
    return response.data;
  },

  listByPatient: async (patientId: number): Promise<Appointment[]> => {
    const response = await api.get<Appointment[]>(`/patients/${patientId}/appointments`);
    return response.data;
  },
};
