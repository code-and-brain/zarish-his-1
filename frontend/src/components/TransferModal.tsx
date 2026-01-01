import { AlertCircle, ArrowRight } from 'lucide-react';
import React, { useEffect, useState } from 'react';

import { useAuth } from '../context/AuthContext';

interface TransferModalProps {
  admissionId: number;
  currentWardId: number;
  currentBedId: number;
  patientName: string;
  onSubmit: (data: TransferData) => Promise<void>;
  onCancel: () => void;
}

export interface TransferData {
  admission_id: number;
  from_ward_id: number;
  from_bed_id: number;
  to_ward_id: number;
  to_bed_id: number;
  reason: string;
  authorized_by: number;
}

interface Ward {
  id: number;
  name: string;
}

interface BedOption {
  id: number;
  bed_number: string;
  room_id: number;
  status: string;
}

const TransferModal: React.FC<TransferModalProps> = ({
  admissionId,
  currentWardId,
  currentBedId,
  patientName,
  onSubmit,
  onCancel,
}) => {
  const [loading, setLoading] = useState(false);
  const [wards, setWards] = useState<Ward[]>([]);
  const [availableBeds, setAvailableBeds] = useState<BedOption[]>([]);
  const { user } = useAuth();

  const [formData, setFormData] = useState<TransferData>({
    admission_id: admissionId,
    from_ward_id: currentWardId,
    from_bed_id: currentBedId,
    to_ward_id: 0,
    to_bed_id: 0,
    reason: '',
    authorized_by: user?.id || 1, // Replaced hardcoded ID with user.id from auth context
  });

  useEffect(() => {
    // Fetch wards
    fetch('/api/v1/wards')
      .then((res) => res.json())
      .then((data) => setWards(data))
      .catch((err) => console.error('Failed to fetch wards:', err));
  }, []);

  useEffect(() => {
    // Fetch available beds when ward is selected
    if (formData.to_ward_id > 0) {
      fetch(`/api/v1/beds?status=Available`)
        .then((res) => res.json())
        .then((data: BedOption[]) => {
          // Filter beds by selected ward (assuming beds have ward_id)
          setAvailableBeds(data);
        })
        .catch((err) => console.error('Failed to fetch beds:', err));
    }
  }, [formData.to_ward_id]);

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
    >
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name.includes('_id') ? parseInt(value) : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (formData.to_bed_id === 0) {
      alert('Please select a destination bed');
      return;
    }
    setLoading(true);
    try {
      await onSubmit(formData);
    } catch (error) {
      console.error('Failed to transfer patient:', error);
      alert('Failed to transfer patient');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-lg shadow-xl max-w-2xl w-full p-6">
        <div className="flex items-center mb-6">
          <ArrowRight className="h-6 w-6 text-blue-600 mr-2" />
          <h2 className="text-2xl font-bold text-gray-900">Transfer Patient</h2>
        </div>

        <div className="mb-4 p-4 bg-blue-50 rounded-lg">
          <p className="text-sm text-gray-700">
            <strong>Patient:</strong> {patientName}
          </p>
        </div>

        <form onSubmit={handleSubmit}>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Destination Ward <span className="text-red-500">*</span>
              </label>
              <select
                name="to_ward_id"
                value={formData.to_ward_id}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                required
              >
                <option value="0">Select Ward</option>
                {wards.map((ward) => (
                  <option key={ward.id} value={ward.id}>
                    {ward.name}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Destination Bed <span className="text-red-500">*</span>
              </label>
              <select
                name="to_bed_id"
                value={formData.to_bed_id}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                required
                disabled={formData.to_ward_id === 0}
              >
                <option value="0">Select Bed</option>
                {availableBeds.map((bed) => (
                  <option key={bed.id} value={bed.id}>
                    Bed {bed.bed_number} (Room {bed.room_id})
                  </option>
                ))}
              </select>
              {formData.to_ward_id > 0 && availableBeds.length === 0 && (
                <p className="mt-1 text-sm text-orange-600 flex items-center">
                  <AlertCircle className="h-4 w-4 mr-1" />
                  No available beds in this ward
                </p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Reason for Transfer
              </label>
              <textarea
                name="reason"
                value={formData.reason}
                onChange={handleChange}
                rows={3}
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                placeholder="Reason for transferring patient..."
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
              {loading ? 'Transferring...' : 'Transfer Patient'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default TransferModal;
