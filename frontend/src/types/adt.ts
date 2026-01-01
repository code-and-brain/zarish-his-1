export interface Ward {
  ID: number;
  name: string;
  department: string;
  type: string;
  description: string;
  rooms?: Room[];
}

export interface Room {
  ID: number;
  ward_id: number;
  room_number: string;
  type: string;
  description: string;
  beds?: Bed[];
}

export interface Bed {
  ID: number;
  room_id: number;
  bed_number: string;
  status: 'Available' | 'Occupied' | 'Maintenance' | 'Cleaning';
  type: string;
  notes: string;
}

export interface Admission {
  ID: number;
  patient_id: number;
  patient?: any; // Replace with Patient type
  ward_id: number;
  bed_id: number;
  bed?: Bed;
  admission_date: string;
  discharge_date?: string;
  admitting_doctor_id: number;
  diagnosis: string;
  status: 'Admitted' | 'Discharged' | 'Transferred';
  notes: string;
}
