import api from './api';
import type { Ward, Room, Bed, Admission } from '../types/adt';

export const ADTService = {
  // Wards
  createWard: async (ward: Partial<Ward>): Promise<Ward> => {
    const response = await api.post<Ward>('/wards', ward);
    return response.data;
  },

  listWards: async (): Promise<Ward[]> => {
    const response = await api.get<Ward[]>('/wards');
    return response.data;
  },

  // Rooms
  createRoom: async (room: Partial<Room>): Promise<Room> => {
    const response = await api.post<Room>('/rooms', room);
    return response.data;
  },

  // Beds
  createBed: async (bed: Partial<Bed>): Promise<Bed> => {
    const response = await api.post<Bed>('/beds', bed);
    return response.data;
  },

  listBeds: async (status?: string): Promise<Bed[]> => {
    const params = status ? { status } : {};
    const response = await api.get<Bed[]>('/beds', { params });
    return response.data;
  },

  // Admissions
  admitPatient: async (admission: Partial<Admission>): Promise<Admission> => {
    const response = await api.post<Admission>('/admissions', admission);
    return response.data;
  },

  dischargePatient: async (id: number): Promise<void> => {
    await api.post(`/admissions/${id}/discharge`);
  },

  listActiveAdmissions: async (): Promise<Admission[]> => {
    const response = await api.get<Admission[]>('/admissions/active');
    return response.data;
  },
};
