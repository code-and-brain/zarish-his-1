export interface PharmacyStock {
    id: number;
    medication_id: number;
    medication?: Medication;
    quantity: number;
    batch_number: string;
    expiry_date: string;
    location: string;
    cost_price: number;
    selling_price: number;
    reorder_level: number;
    notes?: string;
}

export interface Dispensing {
    id: number;
    prescription_id: number;
    prescription?: Prescription;
    patient_id: number;
    patient?: Patient;
    medication_id: number;
    medication?: Medication;
    quantity_dispensed: number;
    batch_number: string;
    dispensed_by: number;
    dispensed_at: string;
    instructions?: string;
    notes?: string;
    status: string;
}

export interface StockMovement {
    id: number;
    type: string;
    medication_id: number;
    medication?: Medication;
    quantity: number;
    batch_number: string;
    reference: string;
    reason?: string;
    performed_by: number;
    performed_at: string;
}

import type { Medication } from './index';
import type { Prescription } from './index';
import type { Patient } from './index';
