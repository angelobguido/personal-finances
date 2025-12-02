import { useState } from 'react';
import TransactionList from './TransactionList.jsx';
import ReportView from './ReportView.jsx';

function TransactionView({ transactions, onChangeTransactions, categories }) {

  const [subview, setSubview] = useState('list');

  return (
    <div className="w-3/4 mx-auto">
      <div className="flex flex-row gap-5">
        <button className={`text-2xl font-bold pb-4 hover:cursor-pointer hover:opacity-60 ${subview === 'list' ? 'underline' : ''}`} onClick={() => setSubview('list')}>Transactions</button>
        <button className={`text-2xl font-bold pb-4 hover:cursor-pointer hover:opacity-60 ${subview === 'report' ? 'underline' : ''}`} onClick={() => setSubview('report')}>Report</button>
      </div>
      {subview === 'list' && (
        <TransactionList transactions={transactions} onChangeTransactions={onChangeTransactions} categories={categories} />
      )}
      {subview === 'report' && (
        <ReportView />
      )}

    </div>
  )
}

export default TransactionView;