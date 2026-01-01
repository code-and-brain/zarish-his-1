import {
  ArrowLeft,
  CheckCircle,
  FileText,
  Image as ImageIcon,
  Play,
  Save,
} from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { RadiologyService } from '../services/radiologyService';
import type { ImagingStudy, RadiologyReport } from '../types/radiology';

const StudyViewer: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [study, setStudy] = useState<ImagingStudy | null>(null);
  const [loading, setLoading] = useState(true);
  const [reportText, setReportText] = useState('');
  const [impression, setImpression] = useState('');
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (id) {
      fetchStudy(parseInt(id));
    }
  }, [id]);

  const fetchStudy = async (studyId: number) => {
    try {
      setLoading(true);
      const data = await RadiologyService.getStudy(studyId);
      setStudy(data);
      if (data.report) {
        setReportText(data.report.findings);
        setImpression(data.report.impression);
      }
    } catch (error) {
      console.error('Error fetching study:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleStatusUpdate = async (newStatus: string) => {
    if (!study) return;
    try {
      await RadiologyService.updateStatus(study.ID, newStatus);
      fetchStudy(study.ID);
    } catch (error) {
      console.error('Error updating status:', error);
    }
  };

  const handleSaveReport = async () => {
    if (!study) return;
    try {
      setSaving(true);
      const reportData: Partial<RadiologyReport> = {
        study_id: study.ID,
        findings: reportText,
        impression: impression,
        status: 'final', // Simplified for MVP
        radiologist_id: 1, // Hardcoded for MVP
      };

      if (study.report) {
        await RadiologyService.updateReport(study.report.ID, reportData);
      } else {
        await RadiologyService.createReport(reportData);
      }

      // Also mark study as completed if not already
      if (study.status !== 'completed') {
        await RadiologyService.updateStatus(study.ID, 'completed');
      }

      fetchStudy(study.ID);
    } catch (error) {
      console.error('Error saving report:', error);
    } finally {
      setSaving(false);
    }
  };

  if (loading) return <div className="p-6 text-center">Loading study...</div>;
  if (!study) return <div className="p-6 text-center">Study not found</div>;

  return (
    <div className="p-6 max-w-7xl mx-auto">
      <div className="flex items-center gap-4 mb-6">
        <button
          onClick={() => navigate('/radiology')}
          className="p-2 hover:bg-gray-100 rounded-full"
        >
          <ArrowLeft size={20} />
        </button>
        <div>
          <h1 className="text-2xl font-bold text-gray-900">
            {study.modality} - {study.body_site}
          </h1>
          <p className="text-gray-500">
            {study.patient?.given_name} {study.patient?.family_name} |{' '}
            {study.accession_number}
          </p>
        </div>
        <div className="ml-auto flex gap-2">
          {study.status === 'scheduled' && (
            <button
              onClick={() => handleStatusUpdate('in-progress')}
              className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 flex items-center gap-2"
            >
              <Play size={18} /> Start Exam
            </button>
          )}
          {study.status === 'in-progress' && (
            <button
              onClick={() => handleStatusUpdate('completed')}
              className="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 flex items-center gap-2"
            >
              <CheckCircle size={18} /> Complete Exam
            </button>
          )}
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Image Viewer / Series List */}
        <div className="lg:col-span-2 space-y-6">
          <div className="bg-black rounded-lg aspect-video flex items-center justify-center text-gray-400">
            {/* Placeholder for DICOM Viewer */}
            <div className="text-center">
              <ImageIcon size={48} className="mx-auto mb-2" />
              <p>DICOM Viewer Placeholder</p>
              <p className="text-sm text-gray-600">PACS Integration Required</p>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-4">
            <h3 className="font-semibold mb-3">Series List</h3>
            <div className="space-y-2">
              {[1, 2, 3].map((series) => (
                <div
                  key={series}
                  className="flex items-center justify-between p-3 border rounded hover:bg-gray-50 cursor-pointer"
                >
                  <div className="flex items-center gap-3">
                    <div className="w-12 h-12 bg-gray-200 rounded flex items-center justify-center">
                      <ImageIcon size={20} className="text-gray-500" />
                    </div>
                    <div>
                      <p className="font-medium">Series {series}</p>
                      <p className="text-xs text-gray-500">50 Images | Axial</p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* Report Editor */}
        <div className="bg-white rounded-lg shadow p-6 h-fit">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold flex items-center gap-2">
              <FileText size={20} />
              Radiology Report
            </h2>
            <span
              className={`px-2 py-1 text-xs rounded-full ${
                study.report
                  ? 'bg-green-100 text-green-800'
                  : 'bg-yellow-100 text-yellow-800'
              }`}
            >
              {study.report ? study.report.status : 'Not Started'}
            </span>
          </div>

          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Findings
              </label>
              <textarea
                className="w-full h-64 p-3 border border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="Enter detailed findings..."
                value={reportText}
                onChange={(e) => setReportText(e.target.value)}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Impression/Conclusion
              </label>
              <textarea
                className="w-full h-32 p-3 border border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="Summary impression..."
                value={impression}
                onChange={(e) => setImpression(e.target.value)}
              />
            </div>

            <button
              onClick={handleSaveReport}
              disabled={saving}
              className="w-full bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700 flex items-center justify-center gap-2 disabled:opacity-50"
            >
              <Save size={18} />
              {saving ? 'Saving...' : 'Finalize Report'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StudyViewer;
