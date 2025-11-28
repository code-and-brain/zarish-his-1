import type { Encounter, Patient } from './index';

export interface Invoice {
  id: number;
  invoice_number: string;
  patient_id: number;
  patient?: Patient;
  encounter_id?: number;
  encounter?: Encounter;
  invoice_date: string;
  due_date: string;
  total_amount: number;
  paid_amount: number;
  balance_amount: number;
  status: 'pending' | 'partial' | 'paid' | 'cancelled' | 'overdue';
  items: InvoiceItem[];
  payments?: Payment[];
  created_at: string;
  updated_at: string;
}

export interface InvoiceItem {
  id: number;
  invoice_id: number;
  description: string;
  quantity: number;
  unit_price: number;
  amount: number;
  discount: number;
  tax: number;
  net_amount: number;
  service_date: string;
}

export interface Payment {
  id: number;
  invoice_id: number;
  invoice?: Invoice;
  amount: number;
  payment_method:
    | 'cash'
    | 'card'
    | 'insurance'
    | 'mobile_money'
    | 'bank_transfer';
  transaction_reference?: string;
  paid_at: string;
  received_by: string;
  notes?: string;
  status: 'completed' | 'pending' | 'failed' | 'refunded';
}

export interface InsuranceClaim {
  id: number;
  claim_number: string;
  patient_id: number;
  patient?: Patient;
  encounter_id?: number;
  encounter?: Encounter;
  invoice_id: number;
  invoice?: Invoice;
  insurance_provider: string;
  policy_number: string;
  diagnosis_codes: string; // JSON string or comma-separated
  procedure_codes: string; // JSON string or comma-separated
  total_amount: number;
  approved_amount: number;
  status:
    | 'draft'
    | 'submitted'
    | 'in_review'
    | 'approved'
    | 'rejected'
    | 'paid';
  submitted_at: string;
  reviewed_at?: string;
  paid_at?: string;
  rejection_reason?: string;
  notes?: string;
}

export interface PaymentReport {
  date: string;
  total_collected: number;
  payment_method_breakdown: { [key: string]: number };
  transaction_count: number;
}
