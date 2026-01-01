import { ArrowLeft, CheckCircle, DollarSign } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { billingService } from '../services/billingService';
import type { Invoice, Payment } from '../types/billing';

const PaymentEntry: React.FC = () => {
  const navigate = useNavigate();
  const { invoiceId } = useParams<{ invoiceId: string }>();
  const { user } = useAuth();

  const [invoice, setInvoice] = useState<Invoice | null>(null);
  const [payment, setPayment] = useState<Partial<Payment>>({
    amount: 0,
    payment_method: 'cash',
    received_by: user?.name || 'Current User', // TODO: Get from auth context
    notes: '',
  });

  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [success, setSuccess] = useState(false);

  useEffect(() => {
    if (invoiceId) {
      fetchInvoice(parseInt(invoiceId));
    }
  }, [invoiceId]);

  const fetchInvoice = async (id: number) => {
    try {
      setLoading(true);
      const data = await billingService.getInvoice(id);
      setInvoice(data);
      setPayment((prev) => ({
        ...prev,
        amount: data.balance_amount,
        invoice_id: id,
      }));
    } catch (error) {
      console.error('Error fetching invoice:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!invoice) return;

    try {
      setSubmitting(true);
      await billingService.recordPayment(payment);
      setSuccess(true);
      setTimeout(() => {
        navigate('/billing');
      }, 2000);
    } catch (error) {
      console.error('Error recording payment:', error);
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) return <div>Loading...</div>;
  if (!invoice) return <div>Invoice not found</div>;

  if (success) {
    return (
      <div className="flex flex-col items-center justify-center h-96">
        <div className="bg-green-100 p-4 rounded-full mb-4">
          <CheckCircle className="w-12 h-12 text-green-600" />
        </div>
        <h2 className="text-2xl font-bold text-gray-900 mb-2">
          Payment Recorded!
        </h2>
        <p className="text-gray-500">Redirecting to dashboard...</p>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto p-6">
      <button
        onClick={() => navigate('/billing')}
        className="flex items-center text-gray-500 hover:text-gray-700 mb-6"
      >
        <ArrowLeft className="w-4 h-4 mr-2" />
        Back to Dashboard
      </button>

      <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-6 border-b border-gray-100 bg-gray-50">
          <h1 className="text-xl font-bold text-gray-900">Record Payment</h1>
          <p className="text-gray-500 mt-1">
            Invoice #{invoice.invoice_number} - {invoice.patient?.given_name}{' '}
            {invoice.patient?.family_name}
          </p>
        </div>

        <form onSubmit={handleSubmit} className="p-6 space-y-6">
          <div className="bg-blue-50 p-4 rounded-lg flex justify-between items-center">
            <span className="text-blue-700 font-medium">
              Outstanding Balance
            </span>
            <span className="text-2xl font-bold text-blue-700">
              ${invoice.balance_amount.toFixed(2)}
            </span>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Payment Amount
            </label>
            <div className="relative">
              <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <span className="text-gray-500 sm:text-sm">$</span>
              </div>
              <input
                type="number"
                value={payment.amount}
                onChange={(e) =>
                  setPayment((prev) => ({
                    ...prev,
                    amount: parseFloat(e.target.value),
                  }))
                }
                className="pl-7 w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
                max={invoice.balance_amount}
                min="0.01"
                step="0.01"
                required
              />
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Payment Method
            </label>
            <select
              value={payment.payment_method}
              onChange={(e) =>
                setPayment((prev) => ({
                  ...prev,
                  payment_method: e.target.value as any,
                }))
              }
              className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
            >
              <option value="cash">Cash</option>
              <option value="card">Credit/Debit Card</option>
              <option value="insurance">Insurance</option>
              <option value="mobile_money">Mobile Money</option>
              <option value="bank_transfer">Bank Transfer</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Transaction Reference (Optional)
            </label>
            <input
              type="text"
              value={payment.transaction_reference || ''}
              onChange={(e) =>
                setPayment((prev) => ({
                  ...prev,
                  transaction_reference: e.target.value,
                }))
              }
              className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
              placeholder="e.g. Receipt number, Transfer ID"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Notes
            </label>
            <textarea
              value={payment.notes || ''}
              onChange={(e) =>
                setPayment((prev) => ({ ...prev, notes: e.target.value }))
              }
              className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
              rows={3}
            />
          </div>

          <button
            type="submit"
            disabled={submitting}
            className="w-full flex justify-center items-center px-4 py-3 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:opacity-50 font-medium"
          >
            <DollarSign className="w-5 h-5 mr-2" />
            {submitting ? 'Processing...' : 'Confirm Payment'}
          </button>
        </form>
      </div>
    </div>
  );
};

export default PaymentEntry;
