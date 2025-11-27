import React, { useState } from 'react';
import { PatientService } from '../services/patientService';
import type { Patient } from '../types';
import { useNavigate } from 'react-router-dom';

const PatientRegistrationForm: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [formData, setFormData] = useState<Partial<Patient>>({
    active: true,
    nationality: 'bangladeshi', // Default
    gender: 'male',
    country: 'Bangladesh',

  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      // Convert birth_date from yyyy-MM-dd to ISO 8601 format
      const dataToSubmit = { ...formData };
      if (dataToSubmit.birth_date && !dataToSubmit.birth_date.includes('T')) {
        dataToSubmit.birth_date = `${dataToSubmit.birth_date}T00:00:00Z`;
      }
      
      await PatientService.create(dataToSubmit);
      navigate('/patients'); // Redirect to patient list
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to register patient');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-6 bg-white rounded-lg shadow-md">
      <h2 className="text-2xl font-bold mb-6 text-gray-800">New Patient Registration</h2>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Nationality Selection */}
        <div className="bg-blue-50 p-4 rounded-md border border-blue-200">
          <label className="block text-sm font-medium text-gray-700 mb-2">Nationality</label>
          <div className="flex space-x-4">
            <label className="inline-flex items-center">
              <input
                type="radio"
                name="nationality"
                value="bangladeshi"
                checked={formData.nationality === 'bangladeshi'}
                onChange={handleChange}
                className="form-radio h-4 w-4 text-blue-600"
              />
              <span className="ml-2">Bangladeshi Citizen</span>
            </label>
            <label className="inline-flex items-center">
              <input
                type="radio"
                name="nationality"
                value="rohingya"
                checked={formData.nationality === 'rohingya'}
                onChange={handleChange}
                className="form-radio h-4 w-4 text-blue-600"
              />
              <span className="ml-2">Rohingya Refugee</span>
            </label>
            <label className="inline-flex items-center">
              <input
                type="radio"
                name="nationality"
                value="other"
                checked={formData.nationality === 'other'}
                onChange={handleChange}
                className="form-radio h-4 w-4 text-blue-600"
              />
              <span className="ml-2">Other / Foreigner</span>
            </label>
          </div>
        </div>

        {/* Personal Information */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label className="block text-sm font-medium text-gray-700">First Name</label>
            <input
              type="text"
              name="given_name"
              required
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Middle Name</label>
            <input
              type="text"
              name="middle_name"
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Last Name</label>
            <input
              type="text"
              name="family_name"
              required
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            />
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label className="block text-sm font-medium text-gray-700">Date of Birth</label>
            <input
              type="date"
              name="birth_date"
              required
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Gender</label>
            <select
              name="gender"
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            >
              <option value="male">Male</option>
              <option value="female">Female</option>
              <option value="other">Other</option>
              <option value="unknown">Unknown</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Phone Number</label>
            <input
              type="tel"
              name="phone"
              required
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
            />
          </div>
        </div>

        {/* Conditional Fields based on Nationality */}
        {formData.nationality === 'bangladeshi' && (
          <div className="bg-green-50 p-4 rounded-md border border-green-200">
            <h3 className="text-md font-medium text-green-800 mb-3">Bangladeshi Citizen Details</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium text-gray-700">National ID (NID)</label>
                <input
                  type="text"
                  name="national_id"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 p-2 border"
                  placeholder="Enter NID Number"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Birth Registration No</label>
                <input
                  type="text"
                  name="birth_reg_no"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 p-2 border"
                  placeholder="Enter Birth Reg No"
                />
              </div>
            </div>
            <p className="text-xs text-gray-500 mt-2">* At least one ID is required.</p>
          </div>
        )}

        {formData.nationality === 'rohingya' && (
          <div className="bg-orange-50 p-4 rounded-md border border-orange-200">
            <h3 className="text-md font-medium text-orange-800 mb-3">Rohingya Refugee Details</h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label className="block text-sm font-medium text-gray-700">UNHCR Number</label>
                <input
                  type="text"
                  name="unhcr_number"
                  required
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-orange-500 focus:ring-orange-500 p-2 border"
                  placeholder="XXX-XX-XXX"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Camp Name</label>
                <input
                  type="text"
                  name="camp_name"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-orange-500 focus:ring-orange-500 p-2 border"
                  placeholder="e.g. Camp 1E"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Block Number</label>
                <input
                  type="text"
                  name="block_number"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-orange-500 focus:ring-orange-500 p-2 border"
                  placeholder="e.g. Block B"
                />
              </div>
            </div>
          </div>
        )}

        {/* Address */}
        <div>
          <h3 className="text-lg font-medium text-gray-900 mb-3">Address</h3>
          <div className="grid grid-cols-1 gap-6">
            <div>
              <label className="block text-sm font-medium text-gray-700">Address Line 1</label>
              <input
                type="text"
                name="address_line1"
                onChange={handleChange}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
              />
            </div>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label className="block text-sm font-medium text-gray-700">City/Village</label>
                <input
                  type="text"
                  name="city"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">District</label>
                <input
                  type="text"
                  name="district"
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Country</label>
                <input
                  type="text"
                  name="country"
                  value={formData.country}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 p-2 border"
                />
              </div>
            </div>
          </div>
        </div>

        <div className="flex justify-end pt-6">
          <button
            type="button"
            onClick={() => navigate('/patients')}
            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 mr-3"
          >
            Cancel
          </button>
          <button
            type="submit"
            disabled={loading}
            className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
          >
            {loading ? 'Registering...' : 'Register Patient'}
          </button>
        </div>
      </form>
    </div>
  );
};

export default PatientRegistrationForm;
