import type { ImagingStudy, RadiologyReport } from '../types/radiology';
import api from './api';

export const RadiologyService = {
  // Create a new study
  createStudy: async (study: Partial<ImagingStudy>): Promise<ImagingStudy> => {
    const response = await api.post<ImagingStudy>('/radiology/studies', study);
    return response.data;
  },

  // List studies with filtering
  listStudies: async (
    page = 1,
    limit = 20,
    patientId?: number,
    status?: string
  ) => {
    const params: any = { page, limit };
    if (patientId) params.patient_id = patientId;
    if (status) params.status = status;

    const response = await api.get<{ data: ImagingStudy[]; total: number }>(
      '/radiology/studies',
      { params }
    );
    return response.data;
  },

  // Get study by ID
  getStudy: async (id: number): Promise<ImagingStudy> => {
    const response = await api.get<ImagingStudy>(`/radiology/studies/${id}`);
    return response.data;
  },

  // Update study status
  updateStatus: async (
    id: number,
    status: string,
    techId?: number
  ): Promise<void> => {
    await api.put(`/radiology/studies/${id}/status`, {
      status,
      tech_id: techId,
    });
  },

  // Get worklist
  getWorklist: async (): Promise<ImagingStudy[]> => {
    const response = await api.get<ImagingStudy[]>('/radiology/worklist');
    return response.data;
  },

  // Create report
  createReport: async (
    report: Partial<RadiologyReport>
  ): Promise<RadiologyReport> => {
    const response = await api.post<RadiologyReport>(
      '/radiology/reports',
      report
    );
    return response.data;
  },

  // Update report
  updateReport: async (
    id: number,
    report: Partial<RadiologyReport>
  ): Promise<RadiologyReport> => {
    const response = await api.put<RadiologyReport>(
      `/radiology/reports/${id}`,
      report
    );
    return response.data;
  },
};
