import type { Medication, Prescription } from '../types';
import api from './api';

export const MedicationService = {
  // Medications
  createMedication: async (med: Partial<Medication>): Promise<Medication> => {
    const response = await api.post<Medication>('/medications', med);
    return response.data;
  },

  searchMedications: async (query: string): Promise<Medication[]> => {
    const response = await api.get<Medication[]>('/medications/search', {
      params: { q: query },
    });
    return response.data;
  },

  // Prescriptions
  createPrescription: async (
    prescription: Partial<Prescription>,
    force: boolean = false
  ): Promise<Prescription> => {
    const url = force ? '/prescriptions?force=true' : '/prescriptions';
    const response = await api.post<Prescription>(url, prescription);
    return response.data;
  },

  getPrescription: async (id: number): Promise<Prescription> => {
    const response = await api.get<Prescription>(`/prescriptions/${id}`);
    return response.data;
  },

  discontinuePrescription: async (
    id: number,
    reason: string
  ): Promise<Prescription> => {
    const response = await api.post<Prescription>(
      `/prescriptions/${id}/discontinue`,
      { reason }
    );
    return response.data;
  },

  listPatientPrescriptions: async (
    patientId: number,
    activeOnly = false
  ): Promise<Prescription[]> => {
    const response = await api.get<Prescription[]>(
      `/patients/${patientId}/prescriptions`,
      { params: { active: activeOnly } }
    );
    return response.data;
  },
};
