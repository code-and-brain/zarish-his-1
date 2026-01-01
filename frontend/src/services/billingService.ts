import type { InsuranceClaim, Invoice, Payment } from '../types/billing';
import api from './api';

export const billingService = {
  // Invoice Operations
  createInvoice: async (invoice: Partial<Invoice>): Promise<Invoice> => {
    const response = await api.post<Invoice>('/billing/invoices', invoice);
    return response.data;
  },

  getInvoice: async (id: number): Promise<Invoice> => {
    const response = await api.get<Invoice>(`/billing/invoices/${id}`);
    return response.data;
  },

  getPatientInvoices: async (patientId: number): Promise<Invoice[]> => {
    const response = await api.get<Invoice[]>(
      `/billing/invoices/patient/${patientId}`
    );
    return response.data;
  },

  getOutstandingInvoices: async (): Promise<Invoice[]> => {
    const response = await api.get<Invoice[]>('/billing/invoices/outstanding');
    return response.data;
  },

  // Payment Operations
  recordPayment: async (payment: Partial<Payment>): Promise<Payment> => {
    const response = await api.post<Payment>('/billing/payments', payment);
    return response.data;
  },

  getInvoicePayments: async (invoiceId: number): Promise<Payment[]> => {
    const response = await api.get<Payment[]>(
      `/billing/payments/invoice/${invoiceId}`
    );
    return response.data;
  },

  getPaymentReport: async (
    startDate: string,
    endDate: string
  ): Promise<Payment[]> => {
    const response = await api.get<Payment[]>('/billing/payments/report', {
      params: { start_date: startDate, end_date: endDate },
    });
    return response.data;
  },

  // Insurance Claim Operations
  submitClaim: async (
    claim: Partial<InsuranceClaim>
  ): Promise<InsuranceClaim> => {
    const response = await api.post<InsuranceClaim>('/insurance/claims', claim);
    return response.data;
  },

  getClaim: async (id: number): Promise<InsuranceClaim> => {
    const response = await api.get<InsuranceClaim>(`/insurance/claims/${id}`);
    return response.data;
  },

  getPendingClaims: async (): Promise<InsuranceClaim[]> => {
    const response = await api.get<InsuranceClaim[]>(
      '/insurance/claims/pending'
    );
    return response.data;
  },

  approveClaim: async (id: number, approvedAmount: number): Promise<void> => {
    await api.post(`/insurance/claims/${id}/approve`, {
      approved_amount: approvedAmount,
    });
  },

  rejectClaim: async (id: number, reason: string): Promise<void> => {
    await api.post(`/insurance/claims/${id}/reject`, { reason });
  },
};
