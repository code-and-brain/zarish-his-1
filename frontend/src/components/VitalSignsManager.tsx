import React, { useEffect, useState } from 'react';
import { VitalSignsService } from '../services/vitalSignsService';
import { EncounterService } from '../services/encounterService';
import type { VitalSigns, Encounter } from '../types';

interface Props {
  patientId: number;
}

const VitalSignsManager: React.FC<Props> = ({ patientId }) => {
  const [vitalsList, setVitalsList] = useState<VitalSigns[]>([]);
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Partial<VitalSigns>>({});

  const fetchData = async () => {
    setLoading(true);
    try {
      const [vitalsData, encountersData] = await Promise.all([
        VitalSignsService.listByPatient(patientId),
        EncounterService.listByPatient(patientId)
      ]);
      setVitalsList(vitalsData);
      setEncounters(encountersData.data);
      
      // Default to latest encounter if available
      if (encountersData.data.length > 0) {
        setFormData(prev => ({ ...prev, encounter_id: encountersData.data[0].id }));
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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await VitalSignsService.create({ ...formData, patient_id: patientId });
      setShowForm(false);
      fetchData();
      setFormData({}); // Reset
    } catch (error) {
      console.error('Failed to record vitals', error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value ? parseFloat(value) : undefined }));
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Vital Signs</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ Record Vitals'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Encounter</label>
            <select
              name="encounter_id"
              className="mt-1 block w-full rounded border-gray-300 p-2 border"
              value={formData.encounter_id || ''}
              onChange={(e) => setFormData({ ...formData, encounter_id: parseInt(e.target.value) })}
              required
            >
              <option value="">Select Encounter</option>
              {encounters.map(enc => (
                <option key={enc.id} value={enc.id}>
                  {new Date(enc.period_start).toLocaleDateString()} - {enc.type} ({enc.status})
                </option>
              ))}
            </select>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-4">
            <div>
              <label className="block text-xs font-medium text-gray-500">Systolic BP</label>
              <input type="number" name="systolic_bp" onChange={handleChange} className="w-full p-2 border rounded" placeholder="mmHg" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Diastolic BP</label>
              <input type="number" name="diastolic_bp" onChange={handleChange} className="w-full p-2 border rounded" placeholder="mmHg" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Pulse Rate</label>
              <input type="number" name="pulse_rate" onChange={handleChange} className="w-full p-2 border rounded" placeholder="bpm" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Resp Rate</label>
              <input type="number" name="respiratory_rate" onChange={handleChange} className="w-full p-2 border rounded" placeholder="/min" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Temp (°C)</label>
              <input type="number" step="0.1" name="temperature" onChange={handleChange} className="w-full p-2 border rounded" placeholder="°C" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">SpO2 (%)</label>
              <input type="number" name="spo2" onChange={handleChange} className="w-full p-2 border rounded" placeholder="%" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Weight (kg)</label>
              <input type="number" step="0.1" name="weight" onChange={handleChange} className="w-full p-2 border rounded" placeholder="kg" />
            </div>
            <div>
              <label className="block text-xs font-medium text-gray-500">Height (cm)</label>
              <input type="number" name="height" onChange={handleChange} className="w-full p-2 border rounded" placeholder="cm" />
            </div>
          </div>

          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            Save Vitals
          </button>
        </form>
      )}

      {loading ? (
        <p>Loading vitals...</p>
      ) : vitalsList.length === 0 ? (
        <p className="text-gray-500">No vital signs recorded.</p>
      ) : (
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200 text-sm">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-2 text-left">Date</th>
                <th className="px-4 py-2 text-left">BP</th>
                <th className="px-4 py-2 text-left">Pulse</th>
                <th className="px-4 py-2 text-left">Temp</th>
                <th className="px-4 py-2 text-left">SpO2</th>
                <th className="px-4 py-2 text-left">BMI</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {vitalsList.map((vitals) => (
                <tr key={vitals.id}>
                  <td className="px-4 py-2 whitespace-nowrap">{new Date(vitals.measured_at).toLocaleString()}</td>
                  <td className="px-4 py-2">{vitals.systolic_bp}/{vitals.diastolic_bp}</td>
                  <td className="px-4 py-2">{vitals.pulse_rate}</td>
                  <td className="px-4 py-2">{vitals.temperature}°C</td>
                  <td className="px-4 py-2">{vitals.spo2}%</td>
                  <td className="px-4 py-2 font-medium">{vitals.bmi?.toFixed(1)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default VitalSignsManager;
