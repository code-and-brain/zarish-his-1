import type { Encounter, Patient } from './index';

export interface ImagingStudy {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  study_uid: string;
  accession_number: string;
  patient_id: number;
  patient?: Patient;
  encounter_id?: number;
  encounter?: Encounter;
  modality: string;
  body_site: string;
  description: string;
  status: 'scheduled' | 'in-progress' | 'completed' | 'cancelled';
  started_at: string;
  completed_at?: string;
  number_of_series: number;
  number_of_instances: number;
  referrer_id?: number;
  performing_tech_id?: number;
  series?: ImagingSeries[];
  report?: RadiologyReport;
}

export interface ImagingSeries {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  series_uid: string;
  study_id: number;
  number: number;
  modality: string;
  description: string;
  body_part: string;
  instances?: ImagingInstance[];
}

export interface ImagingInstance {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  sop_instance_uid: string;
  series_id: number;
  number: number;
  title: string;
  file_path: string;
  content_type: string;
}

export interface RadiologyReport {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  study_id: number;
  radiologist_id: number;
  status: 'draft' | 'preliminary' | 'final' | 'amended';
  findings: string;
  impression: string;
  conclusion: string;
  reported_at: string;
  finalized_at?: string;
}
