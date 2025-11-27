import React, { useEffect, useState } from 'react';
import { EncounterService } from '../services/encounterService';
import type { Encounter } from '../types';

interface Props {
  patientId: number;
}

const EncounterManager: React.FC<Props> = ({ patientId }) => {
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Partial<Encounter>>({
    status: 'planned',
    class: 'amb',
    type: 'consultation',
  });

  const fetchEncounters = async () => {
    setLoading(true);
    try {
      const data = await EncounterService.listByPatient(patientId);
      setEncounters(data.data);
    } catch (error) {
      console.error('Failed to fetch encounters', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEncounters();
  }, [patientId]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await EncounterService.create({ ...formData, patient_id: patientId });
      setShowForm(false);
      fetchEncounters(); // Refresh list
      setFormData({ status: 'planned', class: 'amb', type: 'consultation' }); // Reset form
    } catch (error) {
      console.error('Failed to create encounter', error);
    }
  };

  const handleStatusUpdate = async (id: number, status: string) => {
    try {
      await EncounterService.updateStatus(id, status);
      fetchEncounters();
    } catch (error) {
      console.error('Failed to update status', error);
    }
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Encounters</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ New Encounter'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Type</label>
              <select
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.type}
                onChange={(e) => setFormData({ ...formData, type: e.target.value })}
              >
                <option value="consultation">Consultation</option>
                <option value="emergency">Emergency</option>
                <option value="inpatient">Inpatient</option>
                <option value="follow-up">Follow-up</option>
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Class</label>
              <select
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.class}
                onChange={(e) => setFormData({ ...formData, class: e.target.value })}
              >
                <option value="amb">Ambulatory</option>
                <option value="imp">Inpatient</option>
                <option value="emer">Emergency</option>
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Status</label>
              <select
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.status}
                onChange={(e) => setFormData({ ...formData, status: e.target.value })}
              >
                <option value="planned">Planned</option>
                <option value="arrived">Arrived</option>
                <option value="triaged">Triaged</option>
                <option value="in-progress">In Progress</option>
                <option value="finished">Finished</option>
                <option value="cancelled">Cancelled</option>
              </select>
            </div>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Reason / Chief Complaint</label>
            <textarea
              className="mt-1 block w-full rounded border-gray-300 p-2 border"
              rows={2}
              value={formData.reason || ''}
              onChange={(e) => setFormData({ ...formData, reason: e.target.value })}
            />
          </div>
          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            Save Encounter
          </button>
        </form>
      )}

      {loading ? (
        <p>Loading encounters...</p>
      ) : encounters.length === 0 ? (
        <p className="text-gray-500">No encounters found.</p>
      ) : (
        <div className="space-y-4">
          {encounters.map((encounter) => (
            <div key={encounter.id} className="border rounded p-4 hover:bg-gray-50">
              <div className="flex justify-between">
                <div>
                  <div className="font-medium text-blue-600">
                    {new Date(encounter.period_start).toLocaleDateString()} - {encounter.type}
                  </div>
                  <div className="text-sm text-gray-600 mt-1">
                    {encounter.reason || 'No reason specified'}
                  </div>
                </div>
                <div className="text-right">
                  <span className={`px-2 py-1 rounded text-xs font-medium uppercase ${
                    encounter.status === 'finished' ? 'bg-green-100 text-green-800' :
                    encounter.status === 'in-progress' ? 'bg-blue-100 text-blue-800' :
                    encounter.status === 'cancelled' ? 'bg-red-100 text-red-800' :
                    'bg-gray-100 text-gray-800'
                  }`}>
                    {encounter.status}
                  </span>
                  <div className="mt-2 space-x-2">
                    {encounter.status === 'planned' && (
                      <button
                        onClick={() => handleStatusUpdate(encounter.id, 'arrived')}
                        className="text-xs text-blue-600 hover:underline"
                      >
                        Mark Arrived
                      </button>
                    )}
                    {encounter.status === 'arrived' && (
                      <button
                        onClick={() => handleStatusUpdate(encounter.id, 'in-progress')}
                        className="text-xs text-green-600 hover:underline"
                      >
                        Start Visit
                      </button>
                    )}
                    {encounter.status === 'in-progress' && (
                      <button
                        onClick={() => handleStatusUpdate(encounter.id, 'finished')}
                        className="text-xs text-gray-600 hover:underline"
                      >
                        Finish
                      </button>
                    )}
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default EncounterManager;
