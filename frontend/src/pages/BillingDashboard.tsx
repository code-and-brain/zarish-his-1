import { AlertCircle, DollarSign, Plus, Shield } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { billingService } from '../services/billingService';
import type { InsuranceClaim, Invoice, Payment } from '../types/billing';

const BillingDashboard: React.FC = () => {
  const navigate = useNavigate();
  const [outstandingInvoices, setOutstandingInvoices] = useState<Invoice[]>([]);
  const [pendingClaims, setPendingClaims] = useState<InsuranceClaim[]>([]);
  const [recentPayments, setRecentPayments] = useState<Payment[]>([]);
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    totalOutstanding: 0,
    totalCollectedToday: 0,
    pendingClaimsCount: 0,
  });

  useEffect(() => {
    fetchDashboardData();
  }, []);

  const fetchDashboardData = async () => {
    try {
      setLoading(true);
      const [invoices, claims, payments] = await Promise.all([
        billingService.getOutstandingInvoices(),
        billingService.getPendingClaims(),
        billingService.getPaymentReport(
          new Date().toISOString().split('T')[0],
          new Date().toISOString().split('T')[0]
        ),
      ]);

      setOutstandingInvoices(invoices);
      setPendingClaims(claims);
      setRecentPayments(payments);

      // Calculate stats
      const totalOutstanding = invoices.reduce(
        (sum, inv) => sum + inv.balance_amount,
        0
      );
      const totalCollected = payments.reduce((sum, pay) => sum + pay.amount, 0);

      setStats({
        totalOutstanding,
        totalCollectedToday: totalCollected,
        pendingClaimsCount: claims.length,
      });
    } catch (error) {
      console.error('Error fetching billing data:', error);
    } finally {
      setLoading(false);
    }
  };

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
    }).format(amount);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  return (
    <div className="p-6 space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">
            Billing & Insurance
          </h1>
          <p className="text-gray-500">
            Manage invoices, payments, and insurance claims
          </p>
        </div>
        <div className="flex space-x-3">
          <button
            onClick={() => navigate('/billing/invoices/new')}
            className="flex items-center px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700"
          >
            <Plus className="w-4 h-4 mr-2" />
            New Invoice
          </button>
          <button
            onClick={() => navigate('/billing/claims/new')}
            className="flex items-center px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50"
          >
            <Shield className="w-4 h-4 mr-2" />
            New Claim
          </button>
        </div>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-sm font-medium text-gray-500">
                Outstanding Amount
              </p>
              <h3 className="text-2xl font-bold text-gray-900 mt-1">
                {formatCurrency(stats.totalOutstanding)}
              </h3>
            </div>
            <div className="p-2 bg-red-50 rounded-lg">
              <AlertCircle className="w-6 h-6 text-red-600" />
            </div>
          </div>
          <div className="mt-4 text-sm text-red-600">
            {outstandingInvoices.length} unpaid invoices
          </div>
        </div>

        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-sm font-medium text-gray-500">
                Collected Today
              </p>
              <h3 className="text-2xl font-bold text-gray-900 mt-1">
                {formatCurrency(stats.totalCollectedToday)}
              </h3>
            </div>
            <div className="p-2 bg-green-50 rounded-lg">
              <DollarSign className="w-6 h-6 text-green-600" />
            </div>
          </div>
          <div className="mt-4 text-sm text-green-600">
            {recentPayments.length} transactions
          </div>
        </div>

        <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-sm font-medium text-gray-500">
                Pending Claims
              </p>
              <h3 className="text-2xl font-bold text-gray-900 mt-1">
                {stats.pendingClaimsCount}
              </h3>
            </div>
            <div className="p-2 bg-blue-50 rounded-lg">
              <Shield className="w-6 h-6 text-blue-600" />
            </div>
          </div>
          <div className="mt-4 text-sm text-blue-600">Awaiting approval</div>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Outstanding Invoices */}
        <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
          <div className="p-6 border-b border-gray-100 flex justify-between items-center">
            <h2 className="text-lg font-semibold text-gray-900">
              Outstanding Invoices
            </h2>
            <button
              onClick={() => navigate('/billing/invoices')}
              className="text-sm text-primary-600 hover:text-primary-700"
            >
              View All
            </button>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Invoice
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Patient
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Due Date
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Amount
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Action
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {outstandingInvoices.slice(0, 5).map((invoice) => (
                  <tr key={invoice.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {invoice.invoice_number}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {invoice.patient?.given_name}{' '}
                      {invoice.patient?.family_name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {new Date(invoice.due_date).toLocaleDateString()}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {formatCurrency(invoice.balance_amount)}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      <button
                        onClick={() =>
                          navigate(`/billing/invoices/${invoice.id}/pay`)
                        }
                        className="text-primary-600 hover:text-primary-900"
                      >
                        Pay
                      </button>
                    </td>
                  </tr>
                ))}
                {outstandingInvoices.length === 0 && (
                  <tr>
                    <td
                      colSpan={5}
                      className="px-6 py-8 text-center text-gray-500"
                    >
                      No outstanding invoices
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>

        {/* Pending Claims */}
        <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
          <div className="p-6 border-b border-gray-100 flex justify-between items-center">
            <h2 className="text-lg font-semibold text-gray-900">
              Pending Claims
            </h2>
            <button
              onClick={() => navigate('/billing/claims')}
              className="text-sm text-primary-600 hover:text-primary-700"
            >
              View All
            </button>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Claim #
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Provider
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Submitted
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Amount
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Status
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {pendingClaims.slice(0, 5).map((claim) => (
                  <tr key={claim.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {claim.claim_number}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {claim.insurance_provider}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {new Date(claim.submitted_at).toLocaleDateString()}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {formatCurrency(claim.total_amount)}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">
                        {claim.status}
                      </span>
                    </td>
                  </tr>
                ))}
                {pendingClaims.length === 0 && (
                  <tr>
                    <td
                      colSpan={5}
                      className="px-6 py-8 text-center text-gray-500"
                    >
                      No pending claims
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BillingDashboard;
