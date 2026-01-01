import { FileText, User } from 'lucide-react';
import React, { useState } from 'react';
import { useAuth } from '../context/AuthContext';

interface DischargeFormProps {
  admissionId: number;
  patientName: string;
  onSubmit: (data: DischargeSummaryData) => Promise<void>;
  onCancel: () => void;
}

export interface DischargeSummaryData {
  admission_id: number;
  discharge_type: string;
  chief_complaint: string;
  diagnosis: string;
  treatment_summary: string;
  medications_on_discharge: string;
  follow_up_instructions: string;
  signed_by: number;
}

const DischargeForm: React.FC<DischargeFormProps> = ({
  admissionId,
  patientName,
  onSubmit,
  onCancel,
}) => {
  const [loading, setLoading] = useState(false);
  const { user } = useAuth();

  const [formData, setFormData] = useState<DischargeSummaryData>({
    admission_id: admissionId,
    discharge_type: 'Regular',
    discharge_date: new Date().toISOString().split('T')[0],
    clinical_summary: '',
    medications_prescribed: '',
    follow_up_instructions: '',
    signed_by: user?.id || 1, // Replaced hardcoded ID with user.id from auth context
  });

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
    >
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      await onSubmit(formData);
    } catch (error) {
      console.error('Failed to create discharge summary:', error);
      alert('Failed to create discharge summary');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-lg shadow-xl max-w-3xl w-full max-h-[90vh] overflow-y-auto">
        <div className="p-6">
          <div className="flex items-center mb-6">
            <FileText className="h-6 w-6 text-blue-600 mr-2" />
            <h2 className="text-2xl font-bold text-gray-900">
              Discharge Summary
            </h2>
          </div>

          <div className="mb-4 p-4 bg-blue-50 rounded-lg">
            <p className="text-sm text-gray-700">
              <User className="inline h-4 w-4 mr-1" />
              <strong>Patient:</strong> {patientName}
            </p>
          </div>

          <form onSubmit={handleSubmit}>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Discharge Type <span className="text-red-500">*</span>
                </label>
                <select
                  name="discharge_type"
                  value={formData.discharge_type}
                  onChange={handleChange}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  required
                >
                  <option value="Regular">Regular Discharge</option>
                  <option value="AMA">Against Medical Advice (AMA)</option>
                  <option value="Transfer">Transfer to Another Facility</option>
                  <option value="Death">Death</option>
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Chief Complaint
                </label>
                <textarea
                  name="chief_complaint"
                  value={formData.chief_complaint}
                  onChange={handleChange}
                  rows={2}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  placeholder="Primary reason for admission..."
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Diagnosis <span className="text-red-500">*</span>
                </label>
                <textarea
                  name="diagnosis"
                  value={formData.diagnosis}
                  onChange={handleChange}
                  rows={3}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  placeholder="Final diagnosis..."
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Treatment Summary
                </label>
                <textarea
                  name="treatment_summary"
                  value={formData.treatment_summary}
                  onChange={handleChange}
                  rows={4}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  placeholder="Summary of treatments provided during admission..."
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Medications on Discharge
                </label>
                <textarea
                  name="medications_on_discharge"
                  value={formData.medications_on_discharge}
                  onChange={handleChange}
                  rows={3}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  placeholder="List medications to continue at home..."
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Follow-up Instructions
                </label>
                <textarea
                  name="follow_up_instructions"
                  value={formData.follow_up_instructions}
                  onChange={handleChange}
                  rows={3}
                  className="w-full border border-gray-300 rounded-md px-3 py-2"
                  placeholder="Follow-up appointments, activity restrictions, diet, etc..."
                />
              </div>
            </div>

            <div className="mt-6 flex justify-end space-x-3">
              <button
                type="button"
                onClick={onCancel}
                className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
                disabled={loading}
              >
                Cancel
              </button>
              <button
                type="submit"
                className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:bg-gray-400"
                disabled={loading}
              >
                {loading ? 'Creating...' : 'Complete Discharge'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default DischargeForm;
