import { Activity, AlertTriangle } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';
import { PharmacyService } from '../services/pharmacyService';
import type { Prescription } from '../types';
import type { PharmacyStock } from '../types/pharmacy';

const PharmacyDashboard: React.FC = () => {
  const [lowStock, setLowStock] = useState<PharmacyStock[]>([]);
  const [dispensingQueue, setDispensingQueue] = useState<Prescription[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedPrescription, setSelectedPrescription] =
    useState<Prescription | null>(null);
  const [dispensingForm, setDispensingForm] = useState({
    quantity_dispensed: 0,
    instructions: '',
    notes: '',
  });

  const { user } = useAuth();

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const [lowStockData, queueData] = await Promise.all([
        PharmacyService.getLowStock(),
        PharmacyService.getDispensingQueue(),
      ]);
      setLowStock(lowStockData);
      setDispensingQueue(queueData);
    } catch (error) {
      console.error('Failed to load pharmacy data:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleDispense = async () => {
    if (!selectedPrescription) return;

    try {
      await PharmacyService.dispenseMedication({
        prescription_id: selectedPrescription.id,
        patient_id: selectedPrescription.patient_id,
        medication_id: selectedPrescription.medication_id,
        quantity_dispensed: dispensingForm.quantity_dispensed,
        instructions: dispensingForm.instructions,
        notes: dispensingForm.notes,
        dispensed_by: user?.id || 1, // TODO: Get from auth context
      });

      setSelectedPrescription(null);
      setDispensingForm({ quantity_dispensed: 0, instructions: '', notes: '' });
      loadData();
    } catch (error) {
      console.error('Failed to dispense medication:', error);
      alert('Failed to dispense medication. Please check stock availability.');
    }
  };

  if (loading) return <div>Loading Pharmacy Dashboard...</div>;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <h1 className="text-3xl font-bold text-gray-900 mb-8">
        Pharmacy Dashboard
      </h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
        {/* Low Stock Alerts */}
        <div className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center mb-4">
            <AlertTriangle className="h-6 w-6 text-orange-500 mr-2" />
            <h2 className="text-xl font-semibold">Low Stock Alerts</h2>
          </div>
          <div className="space-y-3">
            {lowStock.length === 0 ? (
              <p className="text-gray-500">No low stock items</p>
            ) : (
              lowStock.map((stock) => (
                <div
                  key={stock.id}
                  className="border-l-4 border-orange-500 pl-3 py-2"
                >
                  <p className="font-medium">{stock.medication?.name}</p>
                  <p className="text-sm text-gray-600">
                    Quantity: {stock.quantity} | Reorder: {stock.reorder_level}
                  </p>
                  <p className="text-xs text-gray-500">
                    Batch: {stock.batch_number}
                  </p>
                </div>
              ))
            )}
          </div>
        </div>

        {/* Dispensing Queue */}
        <div className="lg:col-span-2 bg-white rounded-lg shadow p-6">
          <div className="flex items-center mb-4">
            <Activity className="h-6 w-6 text-blue-500 mr-2" />
            <h2 className="text-xl font-semibold">Dispensing Queue</h2>
          </div>
          <div className="space-y-3">
            {dispensingQueue.length === 0 ? (
              <p className="text-gray-500">No pending prescriptions</p>
            ) : (
              dispensingQueue.map((prescription) => (
                <div
                  key={prescription.id}
                  className="border rounded-lg p-4 hover:bg-gray-50 cursor-pointer"
                  onClick={() => setSelectedPrescription(prescription)}
                >
                  <div className="flex justify-between items-start">
                    <div>
                      <p className="font-medium">
                        {prescription.medication?.name}
                      </p>
                      <p className="text-sm text-gray-600">
                        Patient: {prescription.patient?.given_name}{' '}
                        {prescription.patient?.family_name}
                      </p>
                      <p className="text-sm text-gray-500">
                        Dosage: {prescription.dosage} | Frequency:{' '}
                        {prescription.frequency}
                      </p>
                    </div>
                    <button
                      className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                      onClick={(e) => {
                        e.stopPropagation();
                        setSelectedPrescription(prescription);
                      }}
                    >
                      Dispense
                    </button>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>

      {/* Dispensing Modal */}
      {selectedPrescription && (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-lg shadow-xl max-w-md w-full p-6">
            <h3 className="text-lg font-semibold mb-4">Dispense Medication</h3>

            <div className="mb-4">
              <p className="text-sm text-gray-600">Medication</p>
              <p className="font-medium">
                {selectedPrescription.medication?.name}
              </p>
            </div>

            <div className="mb-4">
              <p className="text-sm text-gray-600">Patient</p>
              <p className="font-medium">
                {selectedPrescription.patient?.given_name}{' '}
                {selectedPrescription.patient?.family_name}
              </p>
            </div>

            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Quantity to Dispense
              </label>
              <input
                type="number"
                value={dispensingForm.quantity_dispensed}
                onChange={(e) =>
                  setDispensingForm({
                    ...dispensingForm,
                    quantity_dispensed: parseInt(e.target.value),
                  })
                }
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                min="1"
              />
            </div>

            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Instructions
              </label>
              <textarea
                value={dispensingForm.instructions}
                onChange={(e) =>
                  setDispensingForm({
                    ...dispensingForm,
                    instructions: e.target.value,
                  })
                }
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                rows={3}
              />
            </div>

            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Notes
              </label>
              <textarea
                value={dispensingForm.notes}
                onChange={(e) =>
                  setDispensingForm({
                    ...dispensingForm,
                    notes: e.target.value,
                  })
                }
                className="w-full border border-gray-300 rounded-md px-3 py-2"
                rows={2}
              />
            </div>

            <div className="flex justify-end space-x-3">
              <button
                onClick={() => setSelectedPrescription(null)}
                className="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50"
              >
                Cancel
              </button>
              <button
                onClick={handleDispense}
                className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
              >
                Dispense
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default PharmacyDashboard;
