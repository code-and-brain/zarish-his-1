import api from './api';
import type { VitalSigns } from '../types';

export const VitalSignsService = {
  create: async (vitals: Partial<VitalSigns>): Promise<VitalSigns> => {
    const response = await api.post<VitalSigns>('/vital-signs', vitals);
    return response.data;
  },

  getById: async (id: number): Promise<VitalSigns> => {
    const response = await api.get<VitalSigns>(`/vital-signs/${id}`);
    return response.data;
  },

  listByEncounter: async (encounterId: number): Promise<VitalSigns[]> => {
    const response = await api.get<VitalSigns[]>(`/encounters/${encounterId}/vital-signs`);
    return response.data;
  },

  listByPatient: async (patientId: number, limit = 10): Promise<VitalSigns[]> => {
    const response = await api.get<VitalSigns[]>(`/patients/${patientId}/vital-signs`, { params: { limit } });
    return response.data;
  },
};
