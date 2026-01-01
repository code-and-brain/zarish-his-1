import api from './api';
import type { Encounter } from '../types';

export const EncounterService = {
  create: async (encounter: Partial<Encounter>): Promise<Encounter> => {
    const response = await api.post<Encounter>('/encounters', encounter);
    return response.data;
  },

  getById: async (id: number): Promise<Encounter> => {
    const response = await api.get<Encounter>(`/encounters/${id}`);
    return response.data;
  },

  update: async (id: number, encounter: Partial<Encounter>): Promise<Encounter> => {
    const response = await api.put<Encounter>(`/encounters/${id}`, encounter);
    return response.data;
  },

  updateStatus: async (id: number, status: string): Promise<Encounter> => {
    const response = await api.put<Encounter>(`/encounters/${id}/status`, { status });
    return response.data;
  },

  listByPatient: async (patientId: number, page = 1, limit = 20) => {
    const params = { page, limit };
    const response = await api.get<{ data: Encounter[]; total: number }>(`/patients/${patientId}/encounters`, { params });
    return response.data;
  },
};
