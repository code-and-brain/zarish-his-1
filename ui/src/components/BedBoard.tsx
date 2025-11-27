import React, { useEffect, useState } from 'react';
import { ADTService } from '../services/adtService';
import type { Ward } from '../types/adt';
import { BedDouble, Activity } from 'lucide-react';
import clsx from 'clsx';

const BedBoard: React.FC = () => {
  const [wards, setWards] = useState<Ward[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadWards();
  }, []);

  const loadWards = async () => {
    try {
      const data = await ADTService.listWards();
      setWards(data);
    } catch (error) {
      console.error('Failed to load wards:', error);
    } finally {
      setLoading(false);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'Available': return 'bg-green-100 text-green-800 border-green-200';
      case 'Occupied': return 'bg-red-100 text-red-800 border-red-200';
      case 'Maintenance': return 'bg-yellow-100 text-yellow-800 border-yellow-200';
      case 'Cleaning': return 'bg-blue-100 text-blue-800 border-blue-200';
      default: return 'bg-gray-100 text-gray-800 border-gray-200';
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
          <span className="flex items-center gap-1"><div className="w-3 h-3 bg-green-500 rounded-full"></div> Available</span>
          <span className="flex items-center gap-1"><div className="w-3 h-3 bg-red-500 rounded-full"></div> Occupied</span>
          <span className="flex items-center gap-1"><div className="w-3 h-3 bg-yellow-500 rounded-full"></div> Maintenance</span>
        </div>
      </div>

      {wards.map((ward) => (
        <div key={ward.ID} className="bg-white rounded-lg shadow p-6">
          <div className="mb-4 border-b pb-2">
            <h3 className="text-lg font-semibold text-gray-900">{ward.name}</h3>
            <p className="text-sm text-gray-500">{ward.department} - {ward.type}</p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {ward.rooms?.map((room) => (
              <div key={room.ID} className="border rounded-md p-4 bg-gray-50">
                <h4 className="text-sm font-medium text-gray-700 mb-3">Room {room.room_number} ({room.type})</h4>
                <div className="grid grid-cols-2 gap-2">
                  {room.beds?.map((bed) => (
                    <div
                      key={bed.ID}
                      className={clsx(
                        "p-3 rounded border flex flex-col items-center justify-center text-center transition-colors",
                        getStatusColor(bed.status)
                      )}
                    >
                      <BedDouble className="h-5 w-5 mb-1" />
                      <span className="font-bold text-sm">{bed.bed_number}</span>
                      <span className="text-xs mt-1">{bed.status}</span>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};

export default BedBoard;
