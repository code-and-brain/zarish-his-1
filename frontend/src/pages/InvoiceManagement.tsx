import { ArrowLeft, Plus, Printer, Save, Trash2 } from 'lucide-react';
import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { billingService } from '../services/billingService';
import { PatientService } from '../services/patientService';
import type { Patient } from '../types';
import type { Invoice, InvoiceItem } from '../types/billing';

const InvoiceManagement: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const isNew = !id || id === 'new';

  const [invoice, setInvoice] = useState<Partial<Invoice>>({
    invoice_date: new Date().toISOString().split('T')[0],
    due_date: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000)
      .toISOString()
      .split('T')[0],
    items: [],
    status: 'pending',
  });

  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState(false);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    fetchPatients();
    if (!isNew && id) {
      fetchInvoice(parseInt(id));
    }
  }, [id, isNew]);

  const fetchPatients = async () => {
    try {
      const response = await PatientService.list(1, 100);
      setPatients(response.data);
    } catch (error) {
      console.error('Error fetching patients:', error);
    }
  };

  const fetchInvoice = async (invoiceId: number) => {
    try {
      setLoading(true);
      const data = await billingService.getInvoice(invoiceId);
      setInvoice(data);
    } catch (error) {
      console.error('Error fetching invoice:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddItem = () => {
    const newItem: Partial<InvoiceItem> = {
      description: '',
      quantity: 1,
      unit_price: 0,
      amount: 0,
      discount: 0,
      tax: 0,
      net_amount: 0,
      service_date: new Date().toISOString().split('T')[0],
    };

    setInvoice((prev) => ({
      ...prev,
      items: [...(prev.items || []), newItem as InvoiceItem],
    }));
  };

  const handleRemoveItem = (index: number) => {
    setInvoice((prev) => ({
      ...prev,
      items: prev.items?.filter((_, i) => i !== index),
    }));
  };

  const handleItemChange = (
    index: number,
    field: keyof InvoiceItem,
    value: any
  ) => {
    setInvoice((prev) => {
      const newItems = [...(prev.items || [])];
      const item = { ...newItems[index], [field]: value };

      // Recalculate amounts
      if (['quantity', 'unit_price', 'discount', 'tax'].includes(field)) {
        const qty = field === 'quantity' ? parseFloat(value) : item.quantity;
        const price =
          field === 'unit_price' ? parseFloat(value) : item.unit_price;
        const discount =
          field === 'discount' ? parseFloat(value) : item.discount;
        const tax = field === 'tax' ? parseFloat(value) : item.tax;

        item.amount = qty * price;
        item.net_amount = item.amount - discount + tax;
      }

      newItems[index] = item;
      return { ...prev, items: newItems };
    });
  };

  const calculateTotal = () => {
    return (
      invoice.items?.reduce((sum, item) => sum + (item.net_amount || 0), 0) || 0
    );
  };

  const handleSave = async () => {
    try {
      setSaving(true);
      if (isNew) {
        await billingService.createInvoice(invoice);
      } else {
        // Update logic would go here
      }
      navigate('/billing');
    } catch (error) {
      console.error('Error saving invoice:', error);
    } finally {
      setSaving(false);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6 max-w-5xl mx-auto bg-white rounded-xl shadow-sm border border-gray-100">
      <div className="flex justify-between items-center mb-6">
        <div className="flex items-center">
          <button
            onClick={() => navigate('/billing')}
            className="mr-4 p-2 hover:bg-gray-100 rounded-full"
          >
            <ArrowLeft className="w-5 h-5 text-gray-500" />
          </button>
          <h1 className="text-2xl font-bold text-gray-900">
            {isNew ? 'New Invoice' : `Invoice #${invoice.invoice_number}`}
          </h1>
        </div>
        <div className="flex space-x-3">
          {!isNew && (
            <button className="flex items-center px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50">
              <Printer className="w-4 h-4 mr-2" />
              Print
            </button>
          )}
          <button
            onClick={handleSave}
            disabled={saving}
            className="flex items-center px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:opacity-50"
          >
            <Save className="w-4 h-4 mr-2" />
            {saving ? 'Saving...' : 'Save Invoice'}
          </button>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Patient
          </label>
          <select
            value={invoice.patient_id || ''}
            onChange={(e) =>
              setInvoice((prev) => ({
                ...prev,
                patient_id: parseInt(e.target.value),
              }))
            }
            className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
            disabled={!isNew}
          >
            <option value="">Select Patient</option>
            {patients.map((p) => (
              <option key={p.id} value={p.id}>
                {p.given_name} {p.family_name}
              </option>
            ))}
          </select>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Invoice Date
            </label>
            <input
              type="date"
              value={invoice.invoice_date}
              onChange={(e) =>
                setInvoice((prev) => ({
                  ...prev,
                  invoice_date: e.target.value,
                }))
              }
              className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Due Date
            </label>
            <input
              type="date"
              value={invoice.due_date}
              onChange={(e) =>
                setInvoice((prev) => ({ ...prev, due_date: e.target.value }))
              }
              className="w-full border-gray-300 rounded-lg focus:ring-primary-500 focus:border-primary-500"
            />
          </div>
        </div>
      </div>

      <div className="mb-8">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-lg font-semibold text-gray-900">Invoice Items</h2>
          <button
            onClick={handleAddItem}
            className="flex items-center text-sm text-primary-600 hover:text-primary-700"
          >
            <Plus className="w-4 h-4 mr-1" />
            Add Item
          </button>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">
                  Description
                </th>
                <th className="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase w-24">
                  Qty
                </th>
                <th className="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase w-32">
                  Price
                </th>
                <th className="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase w-24">
                  Tax
                </th>
                <th className="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase w-32">
                  Total
                </th>
                <th className="px-4 py-2 w-10"></th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {invoice.items?.map((item, index) => (
                <tr key={index}>
                  <td className="px-4 py-2">
                    <input
                      type="text"
                      value={item.description}
                      onChange={(e) =>
                        handleItemChange(index, 'description', e.target.value)
                      }
                      className="w-full border-gray-300 rounded focus:ring-primary-500 focus:border-primary-500"
                      placeholder="Service or item description"
                    />
                  </td>
                  <td className="px-4 py-2">
                    <input
                      type="number"
                      value={item.quantity}
                      onChange={(e) =>
                        handleItemChange(index, 'quantity', e.target.value)
                      }
                      className="w-full text-right border-gray-300 rounded focus:ring-primary-500 focus:border-primary-500"
                      min="1"
                    />
                  </td>
                  <td className="px-4 py-2">
                    <input
                      type="number"
                      value={item.unit_price}
                      onChange={(e) =>
                        handleItemChange(index, 'unit_price', e.target.value)
                      }
                      className="w-full text-right border-gray-300 rounded focus:ring-primary-500 focus:border-primary-500"
                      min="0"
                      step="0.01"
                    />
                  </td>
                  <td className="px-4 py-2">
                    <input
                      type="number"
                      value={item.tax}
                      onChange={(e) =>
                        handleItemChange(index, 'tax', e.target.value)
                      }
                      className="w-full text-right border-gray-300 rounded focus:ring-primary-500 focus:border-primary-500"
                      min="0"
                      step="0.01"
                    />
                  </td>
                  <td className="px-4 py-2 text-right font-medium text-gray-900">
                    ${item.net_amount?.toFixed(2)}
                  </td>
                  <td className="px-4 py-2 text-center">
                    <button
                      onClick={() => handleRemoveItem(index)}
                      className="text-red-500 hover:text-red-700"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
            <tfoot className="bg-gray-50">
              <tr>
                <td
                  colSpan={4}
                  className="px-4 py-3 text-right font-bold text-gray-900"
                >
                  Total Amount:
                </td>
                <td className="px-4 py-3 text-right font-bold text-primary-600 text-lg">
                  ${calculateTotal().toFixed(2)}
                </td>
                <td></td>
              </tr>
            </tfoot>
          </table>
        </div>
      </div>
    </div>
  );
};

export default InvoiceManagement;
