import React, { useState, useEffect } from 'react';
import { ADTService } from '../services/adtService';
import type { Bed } from '../types/adt';
import { PatientService } from '../services/patientService';
import type { Patient } from '../types';

interface AdmissionFormProps {
  onSuccess: () => void;
  onCancel: () => void;
}

const AdmissionForm: React.FC<AdmissionFormProps> = ({ onSuccess, onCancel }) => {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [availableBeds, setAvailableBeds] = useState<Bed[]>([]);
  const [formData, setFormData] = useState({
    patient_id: '',
    bed_id: '',
    admitting_doctor_id: 1, // Hardcoded for now
    diagnosis: '',
    notes: '',
  });

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const [patientsData, bedsData] = await Promise.all([
        PatientService.list(1, 100), // Get first 100 patients
        ADTService.listBeds('Available'),
      ]);
      setPatients(patientsData.data);
      setAvailableBeds(bedsData);
    } catch (error) {
      console.error('Error loading data:', error);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await ADTService.admitPatient({
        patient_id: Number(formData.patient_id),
        bed_id: Number(formData.bed_id),
        admitting_doctor_id: Number(formData.admitting_doctor_id),
        diagnosis: formData.diagnosis,
        notes: formData.notes,
      });
      onSuccess();
    } catch (error) {
      console.error('Admission failed:', error);
      alert('Failed to admit patient');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700">Patient</label>
        <select
          required
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          value={formData.patient_id}
          onChange={(e) => setFormData({ ...formData, patient_id: e.target.value })}
        >
          <option value="">Select Patient</option>
          {patients.map((p) => (
            <option key={p.id} value={p.id}>
              {p.given_name} {p.family_name} (MRN: {p.mrn})
            </option>
          ))}
        </select>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700">Bed</label>
        <select
          required
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          value={formData.bed_id}
          onChange={(e) => setFormData({ ...formData, bed_id: e.target.value })}
        >
          <option value="">Select Bed</option>
          {availableBeds.map((b) => (
            <option key={b.ID} value={b.ID}>
              {b.bed_number} ({b.type})
            </option>
          ))}
        </select>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700">Diagnosis</label>
        <input
          type="text"
          required
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          value={formData.diagnosis}
          onChange={(e) => setFormData({ ...formData, diagnosis: e.target.value })}
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700">Notes</label>
        <textarea
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          rows={3}
          value={formData.notes}
          onChange={(e) => setFormData({ ...formData, notes: e.target.value })}
        />
      </div>

      <div className="flex justify-end gap-3">
        <button
          type="button"
          onClick={onCancel}
          className="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
        >
          Cancel
        </button>
        <button
          type="submit"
          className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Admit Patient
        </button>
      </div>
    </form>
  );
};

export default AdmissionForm;
