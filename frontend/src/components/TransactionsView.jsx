import { useState } from 'react';
import TransactionList from './TransactionList.jsx';
import TransactionForms from './TransactionForms.jsx';

function TransactionView() {

  const [currentTab, setCurrentTab] = useState('all');
  
  const handleClickTab = (tab) => {
    setCurrentTab(tab);
  }

  return (
    <div className="w-3/4 mx-auto">
      <div className="flex flex-col">
        <h2 className="text-2xl font-bold pb-4">Transactions</h2>
        <div className="flex flex-row gap-4 items-center">
          <button onClick={() => handleClickTab('all')} className={(currentTab === 'all' ? 'font-bold' : '')+" hover:underline hover:cursor-pointer"}>All</button>
          <button onClick={() => handleClickTab('addNew')} className={(currentTab === 'addNew' ? 'font-bold' : '')+" hover:underline hover:cursor-pointer"}>Add New</button>
        </div>
      </div>
      {currentTab === 'addNew' && (
        <TransactionForms />
      )}
      {currentTab === 'all' && (
        <TransactionList />
      )}  
    </div>
  )
}

export default TransactionView;