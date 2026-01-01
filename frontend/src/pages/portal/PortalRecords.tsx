import { Activity, Download, FileText, Pill } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import {
  PortalService,
  type PortalRecordsData,
} from '../../services/portalService';

const PortalRecords: React.FC = () => {
  const [data, setData] = useState<PortalRecordsData | null>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'notes' | 'meds' | 'labs'>(
    'notes'
  );

  // Hardcoded patient ID for MVP
  const patientId = 1;

  useEffect(() => {
    fetchRecords();
  }, []);

  const fetchRecords = async () => {
    try {
      setLoading(true);
      const result = await PortalService.getRecords(patientId);
      setData(result);
    } catch (error) {
      console.error('Error fetching records:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div className="p-8 text-center">Loading records...</div>;
  if (!data)
    return <div className="p-8 text-center">Failed to load records</div>;

  return (
    <div className="p-8 max-w-6xl mx-auto">
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">Medical Records</h1>
        <p className="text-gray-500">
          View your clinical notes, medications, and test results
        </p>
      </div>

      {/* Tabs */}
      <div className="flex border-b border-gray-200 mb-6">
        <button
          onClick={() => setActiveTab('notes')}
          className={`pb-4 px-6 text-sm font-medium transition-colors relative ${
            activeTab === 'notes'
              ? 'text-indigo-600'
              : 'text-gray-500 hover:text-gray-700'
          }`}
        >
          <div className="flex items-center gap-2">
            <FileText size={18} />
            Clinical Notes
          </div>
          {activeTab === 'notes' && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-indigo-600" />
          )}
        </button>
        <button
          onClick={() => setActiveTab('meds')}
          className={`pb-4 px-6 text-sm font-medium transition-colors relative ${
            activeTab === 'meds'
              ? 'text-indigo-600'
              : 'text-gray-500 hover:text-gray-700'
          }`}
        >
          <div className="flex items-center gap-2">
            <Pill size={18} />
            Medications
          </div>
          {activeTab === 'meds' && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-indigo-600" />
          )}
        </button>
        <button
          onClick={() => setActiveTab('labs')}
          className={`pb-4 px-6 text-sm font-medium transition-colors relative ${
            activeTab === 'labs'
              ? 'text-indigo-600'
              : 'text-gray-500 hover:text-gray-700'
          }`}
        >
          <div className="flex items-center gap-2">
            <Activity size={18} />
            Lab Results
          </div>
          {activeTab === 'labs' && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-indigo-600" />
          )}
        </button>
      </div>

      {/* Content */}
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 min-h-[400px]">
        {activeTab === 'notes' && (
          <div className="divide-y divide-gray-200">
            {data.clinical_notes.length === 0 ? (
              <div className="p-12 text-center text-gray-500">
                No clinical notes available
              </div>
            ) : (
              data.clinical_notes.map((note: any) => (
                <div key={note.ID} className="p-6">
                  <div className="flex justify-between items-start mb-2">
                    <h3 className="font-semibold text-gray-900">
                      {note.type} Note
                    </h3>
                    <span className="text-sm text-gray-500">
                      {new Date(note.note_date).toLocaleDateString()}
                    </span>
                  </div>
                  <p className="text-gray-600 text-sm line-clamp-3">
                    {note.content}
                  </p>
                  <button className="mt-3 text-indigo-600 text-sm font-medium hover:text-indigo-700 flex items-center gap-1">
                    Read Full Note
                  </button>
                </div>
              ))
            )}
          </div>
        )}

        {activeTab === 'meds' && (
          <div className="divide-y divide-gray-200">
            {data.prescriptions.length === 0 ? (
              <div className="p-12 text-center text-gray-500">
                No active medications
              </div>
            ) : (
              data.prescriptions.map((rx: any) => (
                <div
                  key={rx.ID}
                  className="p-6 flex items-center justify-between"
                >
                  <div>
                    <h3 className="font-semibold text-gray-900">
                      {rx.medication?.name}
                    </h3>
                    <p className="text-sm text-gray-600">
                      {rx.dosage} - {rx.frequency}
                    </p>
                    <p className="text-xs text-gray-500 mt-1">
                      Prescribed: {new Date(rx.start_date).toLocaleDateString()}
                    </p>
                  </div>
                  <span
                    className={`px-3 py-1 rounded-full text-xs font-medium ${
                      rx.status === 'active'
                        ? 'bg-green-100 text-green-800'
                        : 'bg-gray-100 text-gray-800'
                    }`}
                  >
                    {rx.status}
                  </span>
                </div>
              ))
            )}
          </div>
        )}

        {activeTab === 'labs' && (
          <div className="divide-y divide-gray-200">
            {data.lab_orders.length === 0 ? (
              <div className="p-12 text-center text-gray-500">
                No lab results available
              </div>
            ) : (
              data.lab_orders.map((order: any) => (
                <div key={order.ID} className="p-6">
                  <div className="flex justify-between items-center mb-4">
                    <div>
                      <h3 className="font-semibold text-gray-900">
                        Lab Order #{order.ID}
                      </h3>
                      <p className="text-sm text-gray-500">
                        {new Date(order.order_date).toLocaleDateString()}
                      </p>
                    </div>
                    <button className="text-gray-400 hover:text-gray-600">
                      <Download size={20} />
                    </button>
                  </div>
                  <div className="space-y-2">
                    {order.results?.map((result: any) => (
                      <div
                        key={result.ID}
                        className="flex justify-between text-sm p-2 bg-gray-50 rounded"
                      >
                        <span className="font-medium text-gray-700">
                          {result.test_name}
                        </span>
                        <div className="flex items-center gap-3">
                          <span className="text-gray-900">
                            {result.value} {result.unit}
                          </span>
                          {result.is_abnormal && (
                            <span className="px-2 py-0.5 bg-red-100 text-red-800 text-xs rounded font-bold">
                              Abnormal
                            </span>
                          )}
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              ))
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default PortalRecords;
