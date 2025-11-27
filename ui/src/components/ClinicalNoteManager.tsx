import React, { useEffect, useState } from 'react';
import { ClinicalNoteService } from '../services/clinicalNoteService';
import { EncounterService } from '../services/encounterService';
import type { ClinicalNote, Encounter } from '../types';

interface Props {
  patientId: number;
}

const ClinicalNoteManager: React.FC<Props> = ({ patientId }) => {
  const [notes, setNotes] = useState<ClinicalNote[]>([]);
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Partial<ClinicalNote>>({
    note_type: 'soap',
    status: 'draft'
  });

  const fetchData = async () => {
    setLoading(true);
    try {
      const [notesData, encountersData] = await Promise.all([
        ClinicalNoteService.listByPatient(patientId),
        EncounterService.listByPatient(patientId)
      ]);
      setNotes(notesData);
      setEncounters(encountersData.data);
      
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
      await ClinicalNoteService.create({ ...formData, patient_id: patientId });
      setShowForm(false);
      fetchData();
      setFormData({ note_type: 'soap', status: 'draft' });
    } catch (error) {
      console.error('Failed to create note', error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Clinical Notes</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ New Note'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
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
                    {new Date(enc.period_start).toLocaleDateString()} - {enc.type}
                  </option>
                ))}
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Type</label>
              <select
                name="note_type"
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.note_type}
                onChange={handleChange}
              >
                <option value="soap">SOAP Note</option>
                <option value="progress">Progress Note</option>
                <option value="consultation">Consultation</option>
                <option value="discharge">Discharge Summary</option>
              </select>
            </div>
          </div>

          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Subjective (Symptoms)</label>
              <textarea name="subjective" rows={3} onChange={handleChange} className="mt-1 block w-full rounded border-gray-300 p-2 border" />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Objective (Findings)</label>
              <textarea name="objective" rows={3} onChange={handleChange} className="mt-1 block w-full rounded border-gray-300 p-2 border" />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Assessment (Diagnosis)</label>
              <textarea name="assessment" rows={2} onChange={handleChange} className="mt-1 block w-full rounded border-gray-300 p-2 border" />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Plan (Treatment)</label>
              <textarea name="plan" rows={3} onChange={handleChange} className="mt-1 block w-full rounded border-gray-300 p-2 border" />
            </div>
          </div>

          <div className="mt-4">
            <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
              Save Note
            </button>
          </div>
        </form>
      )}

      {loading ? (
        <p>Loading notes...</p>
      ) : notes.length === 0 ? (
        <p className="text-gray-500">No clinical notes found.</p>
      ) : (
        <div className="space-y-6">
          {notes.map((note) => (
            <div key={note.id} className="border rounded-lg p-4 hover:bg-gray-50">
              <div className="flex justify-between border-b pb-2 mb-2">
                <div className="font-medium text-blue-600 uppercase">{note.note_type} Note</div>
                <div className="text-sm text-gray-500">{new Date(note.note_date).toLocaleString()}</div>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
                <div>
                  <span className="font-semibold text-gray-700">S:</span> {note.subjective}
                </div>
                <div>
                  <span className="font-semibold text-gray-700">O:</span> {note.objective}
                </div>
                <div>
                  <span className="font-semibold text-gray-700">A:</span> {note.assessment}
                </div>
                <div>
                  <span className="font-semibold text-gray-700">P:</span> {note.plan}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default ClinicalNoteManager;
