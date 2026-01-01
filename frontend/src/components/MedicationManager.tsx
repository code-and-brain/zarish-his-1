import axios from 'axios';
import { AlertTriangle } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { EncounterService } from '../services/encounterService';
import { MedicationService } from '../services/medicationService';
import type { Encounter, Medication, Prescription } from '../types';

interface Props {
  patientId: number;
}

const MedicationManager: React.FC<Props> = ({ patientId }) => {
  const [prescriptions, setPrescriptions] = useState<Prescription[]>([]);
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [medications, setMedications] = useState<Medication[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');

  // CDS State
  const [warnings, setWarnings] = useState<string[]>([]);
  const [showWarnings, setShowWarnings] = useState(false);

  const [formData, setFormData] = useState<Partial<Prescription>>({
    status: 'active',
    duration_days: 7,
    frequency: 'twice daily',
    dosage: '1 tablet',
  });

  const fetchData = async () => {
    setLoading(true);
    try {
      const [prescriptionsData, encountersData] = await Promise.all([
        MedicationService.listPatientPrescriptions(patientId),
        EncounterService.listByPatient(patientId),
      ]);
      setPrescriptions(prescriptionsData);
      setEncounters(encountersData.data);

      if (encountersData.data.length > 0) {
        setFormData((prev) => ({
          ...prev,
          encounter_id: encountersData.data[0].id,
        }));
      }
    } catch (error) {
      console.error('Failed to fetch data', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [patientId]);

  const handleSearchMedication = async (query: string) => {
    setSearchQuery(query);
    if (query.length > 2) {
      try {
        const results = await MedicationService.searchMedications(query);
        setMedications(results);
      } catch (error) {
        console.error('Search failed', error);
      }
    } else {
      setMedications([]);
    }
  };

  const selectMedication = (med: Medication) => {
    setFormData((prev) => ({ ...prev, medication_id: med.id }));
    setSearchQuery(med.name);
    setMedications([]); // Hide dropdown
  };

  const submitPrescription = async (force: boolean = false) => {
    try {
      await MedicationService.createPrescription(
        { ...formData, patient_id: patientId },
        force
      );
      setShowForm(false);
      setShowWarnings(false);
      setWarnings([]);
      fetchData();
      setFormData({
        status: 'active',
        duration_days: 7,
        frequency: 'twice daily',
        dosage: '1 tablet',
      });
      setSearchQuery('');
    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.status === 409) {
        setWarnings(error.response.data.warnings || []);
        setShowWarnings(true);
      } else {
        console.error('Failed to create prescription', error);
        alert('Failed to create prescription');
      }
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    submitPrescription(false);
  };

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Medications</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ New Prescription'}
        </button>
      </div>

      {showForm && (
        <div className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          {showWarnings ? (
            <div className="bg-orange-50 border-l-4 border-orange-500 p-4 mb-4">
              <div className="flex items-start">
                <AlertTriangle className="h-5 w-5 text-orange-500 mr-2" />
                <div>
                  <h3 className="text-sm font-medium text-orange-800">
                    Clinical Safety Warnings
                  </h3>
                  <ul className="mt-2 list-disc list-inside text-sm text-orange-700">
                    {warnings.map((w, i) => (
                      <li key={i}>{w}</li>
                    ))}
                  </ul>
                  <div className="mt-4 flex space-x-3">
                    <button
                      onClick={() => setShowWarnings(false)}
                      className="text-sm text-gray-600 hover:text-gray-800"
                    >
                      Cancel
                    </button>
                    <button
                      onClick={() => submitPrescription(true)}
                      className="px-3 py-1 bg-orange-600 text-white text-sm rounded hover:bg-orange-700"
                    >
                      Proceed Anyway
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <form onSubmit={handleSubmit}>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Encounter
                  </label>
                  <select
                    name="encounter_id"
                    className="mt-1 block w-full rounded border-gray-300 p-2 border"
                    value={formData.encounter_id || ''}
                    onChange={(e) =>
                      setFormData({
                        ...formData,
                        encounter_id: parseInt(e.target.value),
                      })
                    }
                    required
                  >
                    <option value="">Select Encounter</option>
                    {encounters.map((enc) => (
                      <option key={enc.id} value={enc.id}>
                        {new Date(enc.period_start).toLocaleDateString()} -{' '}
                        {enc.type}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="relative">
                  <label className="block text-sm font-medium text-gray-700">
                    Medication Search
                  </label>
                  <input
                    type="text"
                    className="mt-1 block w-full rounded border-gray-300 p-2 border"
                    value={searchQuery}
                    onChange={(e) => handleSearchMedication(e.target.value)}
                    placeholder="Type to search..."
                    required
                  />
                  {medications.length > 0 && (
                    <ul className="absolute z-10 w-full bg-white border border-gray-300 rounded mt-1 max-h-40 overflow-y-auto shadow-lg">
                      {medications.map((med) => (
                        <li
                          key={med.id}
                          className="p-2 hover:bg-blue-50 cursor-pointer text-sm"
                          onClick={() => selectMedication(med)}
                        >
                          {med.name} ({med.strength})
                        </li>
                      ))}
                    </ul>
                  )}
                </div>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Dosage
                  </label>
                  <input
                    type="text"
                    name="dosage"
                    value={formData.dosage}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded border-gray-300 p-2 border"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Frequency
                  </label>
                  <input
                    type="text"
                    name="frequency"
                    value={formData.frequency}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded border-gray-300 p-2 border"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Duration (Days)
                  </label>
                  <input
                    type="number"
                    name="duration_days"
                    value={formData.duration_days}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded border-gray-300 p-2 border"
                  />
                </div>
              </div>

              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">
                  Instructions
                </label>
                <textarea
                  name="instructions"
                  rows={2}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded border-gray-300 p-2 border"
                />
              </div>

              <button
                type="submit"
                className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm"
              >
                Prescribe
              </button>
            </form>
          )}
        </div>
      )}

      {loading ? (
        <p>Loading prescriptions...</p>
      ) : prescriptions.length === 0 ? (
        <p className="text-gray-500">No prescriptions found.</p>
      ) : (
        <div className="space-y-4">
          {prescriptions.map((p) => (
            <div
              key={p.id}
              className="border rounded p-4 hover:bg-gray-50 flex justify-between items-center"
            >
              <div>
                <div className="font-medium text-blue-600">
                  {p.medication?.name} {p.medication?.strength}
                </div>
                <div className="text-sm text-gray-700">
                  {p.dosage} - {p.frequency} for {p.duration_days} days
                </div>
                {p.instructions && (
                  <div className="text-xs text-gray-500 mt-1 italic">
                    "{p.instructions}"
                  </div>
                )}
              </div>
              <div className="text-right">
                <span
                  className={`px-2 py-1 rounded text-xs font-medium uppercase ${
                    p.status === 'active'
                      ? 'bg-green-100 text-green-800'
                      : 'bg-gray-100 text-gray-800'
                  }`}
                >
                  {p.status}
                </span>
                <div className="text-xs text-gray-500 mt-1">
                  {new Date(p.start_date).toLocaleDateString()}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default MedicationManager;
