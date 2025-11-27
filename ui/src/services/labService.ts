import api from './api';
import type { LabTest, LabOrder, LabResult } from '../types';

export const LabService = {
  // Lab Tests
  createTest: async (test: Partial<LabTest>): Promise<LabTest> => {
    const response = await api.post<LabTest>('/lab-tests', test);
    return response.data;
  },

  listTests: async (): Promise<LabTest[]> => {
    const response = await api.get<LabTest[]>('/lab-tests');
    return response.data;
  },

  // Lab Orders
  createOrder: async (order: Partial<LabOrder>): Promise<LabOrder> => {
    const response = await api.post<LabOrder>('/lab-orders', order);
    return response.data;
  },

  getOrder: async (id: number): Promise<LabOrder> => {
    const response = await api.get<LabOrder>(`/lab-orders/${id}`);
    return response.data;
  },

  listPatientOrders: async (patientId: number): Promise<LabOrder[]> => {
    const response = await api.get<LabOrder[]>(`/patients/${patientId}/lab-orders`);
    return response.data;
  },

  // Lab Results
  addResult: async (result: Partial<LabResult>): Promise<LabResult> => {
    const response = await api.post<LabResult>('/lab-results', result);
    return response.data;
  },
};
