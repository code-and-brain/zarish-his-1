import clsx from 'clsx';
import { Activity, ArrowRight, BedDouble, FileText } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { ADTService } from '../services/adtService';
import type { Admission, Ward } from '../types/adt';
import DischargeForm, { type DischargeSummaryData } from './DischargeForm';
import TransferModal, { type TransferData } from './TransferModal';

const BedBoard: React.FC = () => {
  const [wards, setWards] = useState<Ward[]>([]);
  const [admissions, setAdmissions] = useState<Admission[]>([]);
  const [loading, setLoading] = useState(true);

  // Modal states
  const [showTransferModal, setShowTransferModal] = useState(false);
  const [showDischargeForm, setShowDischargeForm] = useState(false);
  const [selectedAdmission, setSelectedAdmission] = useState<Admission | null>(
    null
  );

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const [wardsData, admissionsData] = await Promise.all([
        ADTService.listWards(),
        ADTService.listActiveAdmissions(),
      ]);
      setWards(wardsData);
      setAdmissions(admissionsData);
    } catch (error) {
      console.error('Failed to load data:', error);
    } finally {
      setLoading(false);
    }
  };

  const getAdmissionForBed = (bedId: number): Admission | undefined => {
    return admissions.find((adm) => adm.bed_id === bedId);
  };

  const handleTransfer = (admission: Admission) => {
    setSelectedAdmission(admission);
    setShowTransferModal(true);
  };

  const handleDischarge = (admission: Admission) => {
    setSelectedAdmission(admission);
    setShowDischargeForm(true);
  };

  const submitTransfer = async (data: TransferData) => {
    await ADTService.transferPatient(data);
    setShowTransferModal(false);
    setSelectedAdmission(null);
    loadData();
  };

  const submitDischarge = async (data: DischargeSummaryData) => {
    await ADTService.createDischargeSummary(data);
    setShowDischargeForm(false);
    setSelectedAdmission(null);
    loadData();
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'Available':
        return 'bg-green-100 text-green-800 border-green-200';
      case 'Occupied':
        return 'bg-red-100 text-red-800 border-red-200';
      case 'Maintenance':
        return 'bg-yellow-100 text-yellow-800 border-yellow-200';
      case 'Cleaning':
        return 'bg-blue-100 text-blue-800 border-blue-200';
      default:
        return 'bg-gray-100 text-gray-800 border-gray-200';
    }
  };

  if (loading) return <div>Loading Bed Board...</div>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold flex items-center gap-2">
          <Activity className="h-6 w-6 text-blue-600" />
          Inpatient Bed Board
        </h2>
        <div className="flex gap-2 text-sm">
          <span className="flex items-center gap-1">
            <div className="w-3 h-3 bg-green-500 rounded-full"></div> Available
          </span>
          <span className="flex items-center gap-1">
            <div className="w-3 h-3 bg-red-500 rounded-full"></div> Occupied
          </span>
          <span className="flex items-center gap-1">
            <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>{' '}
            Maintenance
          </span>
        </div>
      </div>

      {wards.map((ward) => (
        <div key={ward.ID} className="bg-white rounded-lg shadow p-6">
          <div className="mb-4 border-b pb-2">
            <h3 className="text-lg font-semibold text-gray-900">{ward.name}</h3>
            <p className="text-sm text-gray-500">
              {ward.department} - {ward.type}
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {ward.rooms?.map((room) => (
              <div key={room.ID} className="border rounded-md p-4 bg-gray-50">
                <h4 className="text-sm font-medium text-gray-700 mb-3">
                  Room {room.room_number} ({room.type})
                </h4>
                <div className="grid grid-cols-2 gap-2">
                  {room.beds?.map((bed) => {
                    const admission = getAdmissionForBed(bed.ID);
                    return (
                      <div
                        key={bed.ID}
                        className={clsx(
                          'p-3 rounded border flex flex-col items-center justify-center text-center transition-colors relative group',
                          getStatusColor(bed.status)
                        )}
                      >
                        <BedDouble className="h-5 w-5 mb-1" />
                        <span className="font-bold text-sm">
                          {bed.bed_number}
                        </span>
                        <span className="text-xs mt-1">{bed.status}</span>

                        {bed.status === 'Occupied' && admission && (
                          <div className="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-10 transition-all rounded flex items-center justify-center opacity-0 group-hover:opacity-100">
                            <div className="flex gap-1">
                              <button
                                onClick={() => handleTransfer(admission)}
                                className="p-1 bg-blue-600 text-white rounded hover:bg-blue-700"
                                title="Transfer"
                              >
                                <ArrowRight className="h-4 w-4" />
                              </button>
                              <button
                                onClick={() => handleDischarge(admission)}
                                className="p-1 bg-green-600 text-white rounded hover:bg-green-700"
                                title="Discharge"
                              >
                                <FileText className="h-4 w-4" />
                              </button>
                            </div>
                          </div>
                        )}
                      </div>
                    );
                  })}
                </div>
              </div>
            ))}
          </div>
        </div>
      ))}

      {showTransferModal && selectedAdmission && (
        <TransferModal
          admissionId={selectedAdmission.ID}
          currentWardId={selectedAdmission.ward_id}
          currentBedId={selectedAdmission.bed_id}
          patientName={`${selectedAdmission.patient?.given_name} ${selectedAdmission.patient?.family_name}`}
          onSubmit={submitTransfer}
          onCancel={() => {
            setShowTransferModal(false);
            setSelectedAdmission(null);
          }}
        />
      )}

      {showDischargeForm && selectedAdmission && (
        <DischargeForm
          admissionId={selectedAdmission.ID}
          patientName={`${selectedAdmission.patient?.given_name} ${selectedAdmission.patient?.family_name}`}
          onSubmit={submitDischarge}
          onCancel={() => {
            setShowDischargeForm(false);
            setSelectedAdmission(null);
          }}
        />
      )}
    </div>
  );
};

export default BedBoard;
