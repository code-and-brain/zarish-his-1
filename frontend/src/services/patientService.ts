import api from './api';
import type { Patient } from '../types';

export const PatientService = {
  // Create a new patient
  create: async (patient: Partial<Patient>): Promise<Patient> => {
    const response = await api.post<Patient>('/patients', patient);
    return response.data;
  },

  // Get patient by ID
  getById: async (id: number): Promise<Patient> => {
    const response = await api.get<Patient>(`/patients/${id}`);
    return response.data;
  },

  // Update patient
  update: async (id: number, patient: Partial<Patient>): Promise<Patient> => {
    const response = await api.put<Patient>(`/patients/${id}`, patient);
    return response.data;
  },

  // Delete patient
  delete: async (id: number): Promise<void> => {
    await api.delete(`/patients/${id}`);
  },

  // List patients with pagination and search
  list: async (page = 1, limit = 20, search = '', nationality = '') => {
    const params = { page, limit, search, nationality };
    const response = await api.get<{ data: Patient[]; total: number }>('/patients', { params });
    return response.data;
  },

  // Search patients
  search: async (query: string): Promise<Patient[]> => {
    const response = await api.get<Patient[]>('/patients/search', { params: { q: query } });
    return response.data;
  },

  // Get patient history
  getHistory: async (id: number) => {
    const response = await api.get(`/patients/${id}/history`);
    return response.data;
  },
};
