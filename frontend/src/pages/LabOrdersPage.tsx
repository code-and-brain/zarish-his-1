import React, { useEffect, useState } from 'react';
import { LabService } from '../services/labService';
import type { LabTest } from '../types';

const LabOrdersPage: React.FC = () => {
  const [tests, setTests] = useState<LabTest[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTests = async () => {
      setLoading(true);
      try {
        const data = await LabService.listTests();
        setTests(data);
      } catch (error) {
        console.error('Failed to fetch lab tests', error);
      } finally {
        setLoading(false);
      }
    };

    fetchTests();
  }, []);

  return (
    <div className="max-w-6xl mx-auto p-6">
      <h2 className="text-2xl font-bold text-gray-800 mb-6">Lab Dashboard</h2>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Available Tests */}
        <div className="bg-white shadow rounded-lg p-6">
          <h3 className="text-lg font-medium text-gray-900 mb-4">Available Lab Tests</h3>
          {loading ? (
            <p className="text-gray-500">Loading tests...</p>
          ) : (
            <ul className="divide-y divide-gray-200">
              {tests.length === 0 ? (
                <li className="py-3 text-gray-500">No tests defined.</li>
              ) : (
                tests.map((test) => (
                  <li key={test.id} className="py-3 flex justify-between">
                    <div>
                      <p className="text-sm font-medium text-gray-900">{test.name}</p>
                      <p className="text-xs text-gray-500">{test.category}</p>
                    </div>
                    {/* <span className="text-sm text-gray-600">{test.base_price} BDT</span> */}
                  </li>
                ))
              )}
            </ul>
          )}
        </div>

        {/* Recent Orders Placeholder */}
        <div className="bg-white shadow rounded-lg p-6">
          <h3 className="text-lg font-medium text-gray-900 mb-4">Recent Lab Orders</h3>
          <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4">
            <div className="flex">
              <div className="ml-3">
                <p className="text-sm text-yellow-700">
                  Global order view is under development. Please view orders within specific Patient Details.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LabOrdersPage;
