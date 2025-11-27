import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { PatientService } from '../services/patientService';
import type { Patient } from '../types';
import EncounterManager from '../components/EncounterManager';
import VitalSignsManager from '../components/VitalSignsManager';
import ClinicalNoteManager from '../components/ClinicalNoteManager';
import MedicationManager from '../components/MedicationManager';
import LabManager from '../components/LabManager';
import AppointmentManager from '../components/AppointmentManager';

const PatientDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [patient, setPatient] = useState<Patient | null>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('overview');

  useEffect(() => {
    const fetchPatient = async () => {
      if (!id) return;
      try {
        const data = await PatientService.getById(parseInt(id));
        setPatient(data);
      } catch (error) {
        console.error('Failed to fetch patient', error);
      } finally {
        setLoading(false);
      }
    };
    fetchPatient();
  }, [id]);

  if (loading) return <div className="p-6 text-center">Loading patient data...</div>;
  if (!patient) return <div className="p-6 text-center text-red-600">Patient not found</div>;

  const tabs = [
    { id: 'overview', label: 'Overview' },
    { id: 'encounters', label: 'Encounters' },
    { id: 'vitals', label: 'Vital Signs' },
    { id: 'notes', label: 'Clinical Notes' },
    { id: 'medications', label: 'Medications' },
    { id: 'labs', label: 'Lab Results' },
    { id: 'appointments', label: 'Appointments' },
  ];

  return (
    <div className="max-w-7xl mx-auto p-6">
      {/* Patient Header */}
      <div className="bg-white shadow rounded-lg p-6 mb-6 border-l-4 border-blue-600">
        <div className="flex justify-between items-start">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">
              {patient.given_name} {patient.middle_name} {patient.family_name}
            </h1>
            <div className="mt-2 flex flex-wrap gap-4 text-sm text-gray-600">
              <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded font-medium">MRN: {patient.mrn}</span>
              <span>
                {patient.gender.toUpperCase()} •{' '}
                {patient.birth_date &&
                  `${new Date().getFullYear() - new Date(patient.birth_date).getFullYear()} years`}
              </span>
              <span>📞 {patient.phone}</span>
              <span className="capitalize">🌍 {patient.nationality}</span>
            </div>
            {/* Nationality Specifics */}
            <div className="mt-2 text-sm text-gray-500">
              {patient.nationality === 'bangladeshi' && (
                <>
                  {patient.national_id && <span className="mr-4">NID: {patient.national_id}</span>}
                  {patient.birth_reg_no && <span>BRN: {patient.birth_reg_no}</span>}
                </>
              )}
              {patient.nationality === 'rohingya' && (
                <>
                  <span className="mr-4">UNHCR: {patient.unhcr_number}</span>
                  <span>Camp: {patient.camp_name} ({patient.block_number})</span>
                </>
              )}
            </div>
          </div>
          <div className="text-right">
            <span className={`px-3 py-1 rounded-full text-sm font-medium ${patient.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`}>
              {patient.active ? 'Active' : 'Inactive'}
            </span>
          </div>
        </div>
      </div>

      {/* Tabs Navigation */}
      <div className="mb-6 border-b border-gray-200">
        <nav className="-mb-px flex space-x-8 overflow-x-auto">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`
                whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm
                ${activeTab === tab.id
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'}
              `}
            >
              {tab.label}
            </button>
          ))}
        </nav>
      </div>

      {/* Tab Content */}
      <div className="bg-white shadow rounded-lg p-6 min-h-[400px]">
        {activeTab === 'overview' && (
          <div>
            <h3 className="text-lg font-medium mb-4">Patient Overview</h3>
            <p className="text-gray-500">Timeline view coming soon...</p>
            {/* We can add a summary dashboard here later */}
          </div>
        )}
        {activeTab === 'encounters' && <EncounterManager patientId={patient.id} />}
        {activeTab === 'vitals' && <VitalSignsManager patientId={patient.id} />}
        {activeTab === 'notes' && <ClinicalNoteManager patientId={patient.id} />}
        {activeTab === 'medications' && <MedicationManager patientId={patient.id} />}
        {activeTab === 'labs' && <LabManager patientId={patient.id} />}
        {activeTab === 'appointments' && <AppointmentManager patientId={patient.id} />}
      </div>
    </div>
  );
};

export default PatientDetail;
