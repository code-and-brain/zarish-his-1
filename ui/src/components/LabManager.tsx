import React, { useEffect, useState } from 'react';
import { LabService } from '../services/labService';
import { EncounterService } from '../services/encounterService';
import type { LabOrder, LabTest, Encounter } from '../types';

interface Props {
  patientId: number;
}

const LabManager: React.FC<Props> = ({ patientId }) => {
  const [orders, setOrders] = useState<LabOrder[]>([]);
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [tests, setTests] = useState<LabTest[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Partial<LabOrder>>({
    priority: 'routine',
    status: 'ordered'
  });


  const fetchData = async () => {
    setLoading(true);
    try {
      const [ordersData, encountersData, testsData] = await Promise.all([
        LabService.listPatientOrders(patientId),
        EncounterService.listByPatient(patientId),
        LabService.listTests()
      ]);
      setOrders(ordersData);
      setEncounters(encountersData.data);
      setTests(testsData);
      
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
    // In a real app, we would handle multiple tests in one order
    // For MVP, we'll just create the order structure
    try {
      await LabService.createOrder({ ...formData, patient_id: patientId });
      setShowForm(false);
      fetchData();
      setFormData({ priority: 'routine', status: 'ordered' });
    } catch (error) {
      console.error('Failed to create lab order', error);
    }
  };

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-medium">Lab Orders</h3>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm"
        >
          {showForm ? 'Cancel' : '+ Order Labs'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-50 p-4 rounded mb-6 border border-gray-200">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Encounter</label>
              <select
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
              <label className="block text-sm font-medium text-gray-700">Priority</label>
              <select
                className="mt-1 block w-full rounded border-gray-300 p-2 border"
                value={formData.priority}
                onChange={(e) => setFormData({ ...formData, priority: e.target.value })}
              >
                <option value="routine">Routine</option>
                <option value="urgent">Urgent</option>
                <option value="stat">STAT</option>
              </select>
            </div>
          </div>

          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700 mb-2">Select Tests</label>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-2 max-h-40 overflow-y-auto border p-2 rounded bg-white">
              {tests.map(test => (
                <label key={test.id} className="flex items-center space-x-2 text-sm">
                  <input type="checkbox" />
                  <span>{test.name}</span>
                </label>
              ))}
            </div>
          </div>

          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            Place Order
          </button>
        </form>
      )}

      {loading ? (
        <p>Loading lab orders...</p>
      ) : orders.length === 0 ? (
        <p className="text-gray-500">No lab orders found.</p>
      ) : (
        <div className="space-y-4">
          {orders.map((order) => (
            <div key={order.id} className="border rounded p-4 hover:bg-gray-50">
              <div className="flex justify-between">
                <div>
                  <div className="font-medium text-blue-600">Order #{order.id}</div>
                  <div className="text-sm text-gray-500">{new Date(order.order_date).toLocaleDateString()}</div>
                </div>
                <div className="text-right">
                  <span className={`px-2 py-1 rounded text-xs font-medium uppercase ${order.priority === 'stat' ? 'bg-red-100 text-red-800' : 'bg-gray-100 text-gray-800'}`}>
                    {order.priority}
                  </span>
                  <div className="text-xs mt-1 capitalize">{order.status}</div>
                </div>
              </div>
              {/* Results would go here */}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default LabManager;
