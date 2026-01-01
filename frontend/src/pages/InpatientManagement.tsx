import React, { useState } from 'react';
import BedBoard from '../components/BedBoard';
import AdmissionForm from '../components/AdmissionForm';
import { Plus } from 'lucide-react';

const InpatientManagement: React.FC = () => {
  const [showAdmissionForm, setShowAdmissionForm] = useState(false);

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Inpatient Management</h1>
          <p className="mt-2 text-sm text-gray-700">
            Manage wards, beds, and patient admissions.
          </p>
        </div>
        <button
          onClick={() => setShowAdmissionForm(true)}
          className="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          <Plus className="-ml-1 mr-2 h-5 w-5" />
          New Admission
        </button>
      </div>

      {showAdmissionForm && (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-lg shadow-xl max-w-md w-full p-6">
            <h2 className="text-xl font-bold mb-4">Admit Patient</h2>
            <AdmissionForm
              onSuccess={() => {
                setShowAdmissionForm(false);
                window.location.reload(); // Refresh to show updated bed status
              }}
              onCancel={() => setShowAdmissionForm(false)}
            />
          </div>
        </div>
      )}

      <BedBoard />
    </div>
  );
};

export default InpatientManagement;
