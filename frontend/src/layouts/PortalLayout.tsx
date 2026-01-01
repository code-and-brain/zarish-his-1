import {
  Calendar,
  FileText,
  LayoutDashboard,
  LogOut,
  User,
} from 'lucide-react';
import React from 'react';
import { Link, Outlet, useLocation } from 'react-router-dom';

const PortalLayout: React.FC = () => {
  const location = useLocation();

  const isActive = (path: string) => {
    return location.pathname.startsWith(path)
      ? 'bg-indigo-50 text-indigo-600'
      : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900';
  };

  return (
    <div className="min-h-screen bg-gray-50 flex">
      {/* Sidebar */}
      <div className="w-64 bg-white border-r border-gray-200 flex flex-col">
        <div className="p-6 border-b border-gray-200">
          <div className="flex items-center gap-2">
            <div className="w-8 h-8 bg-indigo-600 rounded-lg flex items-center justify-center text-white font-bold">
              Z
            </div>
            <span className="text-xl font-bold text-gray-900">
              Patient Portal
            </span>
          </div>
        </div>

        <nav className="flex-1 p-4 space-y-1">
          <Link
            to="/portal/dashboard"
            className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors ${isActive(
              '/portal/dashboard'
            )}`}
          >
            <LayoutDashboard size={20} />
            Dashboard
          </Link>
          <Link
            to="/portal/appointments"
            className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors ${isActive(
              '/portal/appointments'
            )}`}
          >
            <Calendar size={20} />
            Appointments
          </Link>
          <Link
            to="/portal/records"
            className={`flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-colors ${isActive(
              '/portal/records'
            )}`}
          >
            <FileText size={20} />
            Medical Records
          </Link>
        </nav>

        <div className="p-4 border-t border-gray-200">
          <div className="flex items-center gap-3 px-4 py-3 text-gray-600">
            <User size={20} />
            <div className="flex-1 min-w-0">
              <p className="text-sm font-medium text-gray-900 truncate">
                John Doe
              </p>
              <p className="text-xs text-gray-500 truncate">MRN-000123</p>
            </div>
          </div>
          <button className="w-full mt-2 flex items-center gap-3 px-4 py-2 text-sm font-medium text-red-600 hover:bg-red-50 rounded-lg transition-colors">
            <LogOut size={20} />
            Sign Out
          </button>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 overflow-auto">
        <Outlet />
      </div>
    </div>
  );
};

export default PortalLayout;
