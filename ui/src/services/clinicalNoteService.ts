import api from './api';
import type { ClinicalNote } from '../types';

export const ClinicalNoteService = {
  create: async (note: Partial<ClinicalNote>): Promise<ClinicalNote> => {
    const response = await api.post<ClinicalNote>('/clinical-notes', note);
    return response.data;
  },

  getById: async (id: number): Promise<ClinicalNote> => {
    const response = await api.get<ClinicalNote>(`/clinical-notes/${id}`);
    return response.data;
  },

  update: async (id: number, note: Partial<ClinicalNote>): Promise<ClinicalNote> => {
    const response = await api.put<ClinicalNote>(`/clinical-notes/${id}`, note);
    return response.data;
  },

  sign: async (id: number, userId: number): Promise<ClinicalNote> => {
    const response = await api.post<ClinicalNote>(`/clinical-notes/${id}/sign`, { user_id: userId });
    return response.data;
  },

  listByEncounter: async (encounterId: number): Promise<ClinicalNote[]> => {
    const response = await api.get<ClinicalNote[]>(`/encounters/${encounterId}/clinical-notes`);
    return response.data;
  },

  listByPatient: async (patientId: number, limit = 20): Promise<ClinicalNote[]> => {
    const response = await api.get<ClinicalNote[]>(`/patients/${patientId}/clinical-notes`, { params: { limit } });
    return response.data;
  },
};
