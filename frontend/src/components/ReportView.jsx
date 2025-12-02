import { useState, useEffect } from 'react';
import { getReport } from '../services/api';

function ReportView() {

    const now = new Date();
    const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1).toISOString().split('T')[0];
    const endOfMonth = new Date(now.getFullYear(), now.getMonth() + 1, 0).toISOString().split('T')[0];

    const [reportData, setReportData] = useState(null);
    const [startDate, setStartDate] = useState(startOfMonth);
    const [endDate, setEndDate] = useState(endOfMonth);

    useEffect(() => {
        const fetchReport = async () => {
            const data = await getReport(startDate, endDate);
            setReportData(data);
        };
        fetchReport();
    }, [startDate, endDate]);

    const handleStartDateChange = (e) => setStartDate(e.target.value);
    const handleEndDateChange = (e) => setEndDate(e.target.value);

    let report = {};
    if (reportData){
        const totalIncome = reportData.categories.reduce((sum, category) => sum + (category.is_income ? category.total : 0), 0);
        const totalExpense = reportData.categories.reduce((sum, category) => sum + (!category.is_income ? category.total : 0), 0);
        const netTotal = totalIncome - totalExpense;
        report = {
            totalIncome,
            totalExpense,
            netTotal
        };
    }

    if (!reportData) {
        return <p>Loading report...</p>;
    }

    return (
    <div className='flex flex-col justify-center items-center gap-6 p-4'>
        <div className='flex flex-row justify-center items-center gap-4 bg-gray-100 p-4 rounded shadow'>
            <div className='flex flex-row items-center gap-2'>
                <label className='font-semibold'>Start Date:</label>
                <input 
                    type='date' 
                    value={startDate} 
                    onChange={handleStartDateChange}
                    className='border rounded px-3 py-2'
                />
            </div>
            <div className='flex flex-row items-center gap-2'>
                <label className='font-semibold'>End Date:</label>
                <input 
                    type='date' 
                    value={endDate} 
                    onChange={handleEndDateChange}
                    className='border rounded px-3 py-2'
                />
            </div>
        </div>
        
        <div className='flex flex-row justify-center items-center gap-6 text-lg'>
            <p className='font-semibold'>Total Income: <span className='text-green-600'>${report.totalIncome.toFixed(2)}</span></p>
            <p className='font-semibold'>Total Expense: <span className='text-red-600'>${report.totalExpense.toFixed(2)}</span></p>
            <p className='font-semibold'>Net Total: <span className={report.netTotal >= 0 ? 'text-green-600' : 'text-red-600'}>${report.netTotal.toFixed(2)}</span></p>
        </div>
        <div className='flex flex-col w-full gap-4'>
            <h3 className='text-xl font-bold text-center'>Category Breakdown</h3>
            
            <div className='flex flex-col gap-6'>
                <div className='border rounded p-4 shadow-md bg-gray-100'>
                    <h4 className='text-lg font-semibold mb-3'>Income</h4>
                    <div className='flex flex-row flex-wrap justify-center items-start gap-4'>
                        {reportData.categories.filter(cat => cat.is_income).map((category) => {
                            const percentage = report.totalIncome > 0 ? (category.total / report.totalIncome * 100) : 0;
                            
                            return (
                                <div key={category.id} className='border rounded p-4 shadow-md min-w-[150px]'>
                                    <p className='font-semibold mb-2'>{category.category}</p>
                                    <p className='text-green-600 font-bold'>
                                        ${category.total.toFixed(2)}
                                    </p>
                                    <p className='text-gray-600 text-sm mt-1'>
                                        {percentage.toFixed(1)}% of income
                                    </p>
                                </div>
                            );
                        })}
                    </div>
                </div>

                <div className='border rounded p-4 shadow-md bg-gray-100'>
                    <h4 className='text-lg font-semibold mb-3'>Expenses</h4>
                    <div className='flex flex-row flex-wrap justify-center items-start gap-4'>
                        {reportData.categories.filter(cat => !cat.is_income).map((category) => {
                            const percentage = report.totalExpense > 0 ? (category.total / report.totalExpense * 100) : 0;
                            
                            return (
                                <div key={category.id} className='border rounded p-4 shadow-md min-w-[150px]'>
                                    <p className='font-semibold mb-2'>{category.category}</p>
                                    <p className='text-red-600 font-bold'>
                                        ${category.total.toFixed(2)}
                                    </p>
                                    <p className='text-gray-600 text-sm mt-1'>
                                        {percentage.toFixed(1)}% of expenses
                                    </p>
                                </div>
                            );
                        })}
                    </div>
                </div>
            </div>
        </div>
    </div>

    )
}

export default ReportView;