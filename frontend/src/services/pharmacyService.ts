import api from './api';
import type { PharmacyStock, Dispensing, StockMovement } from '../types/pharmacy';
import type { Prescription } from '../types';

export const PharmacyService = {
    // Stock Management
    addStock: async (stock: Partial<PharmacyStock>) => {
        const response = await api.post<PharmacyStock>('/pharmacy/stock', stock);
        return response.data;
    },

    getStock: async (medicationId: number) => {
        const response = await api.get<PharmacyStock[]>(`/pharmacy/stock/${medicationId}`);
        return response.data;
    },

    getLowStock: async () => {
        const response = await api.get<PharmacyStock[]>('/pharmacy/stock/low');
        return response.data;
    },

    // Dispensing
    dispenseMedication: async (dispensing: Partial<Dispensing>) => {
        const response = await api.post<Dispensing>('/pharmacy/dispense', dispensing);
        return response.data;
    },

    getDispensingQueue: async () => {
        const response = await api.get<Prescription[]>('/pharmacy/dispensing-queue');
        return response.data;
    },

    getPatientHistory: async (patientId: number) => {
        const response = await api.get<Dispensing[]>(`/pharmacy/history/${patientId}`);
        return response.data;
    },

    // Stock Movements
    getStockMovements: async (medicationId: number, startDate?: string, endDate?: string) => {
        const params = new URLSearchParams();
        if (startDate) params.append('start_date', startDate);
        if (endDate) params.append('end_date', endDate);

        const response = await api.get<StockMovement[]>(
            `/pharmacy/movements/${medicationId}?${params.toString()}`
        );
        return response.data;
    },
};
