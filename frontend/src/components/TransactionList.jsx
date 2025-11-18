import Transaction from './Transaction.jsx';
import TransactionForms from './TransactionForms.jsx';
import { useFinanceStore } from '../store/useFinanceStore.js';
import { useState } from 'react';

function TransactionList() {
  const {transactions, loading, error, addTransaction} = useFinanceStore();

  const [isAdding, setIsAdding] = useState(false);
  
  const handleClick = () => {
    setIsAdding(true);
  };

  const handleCancel = () => {
    setIsAdding(false);
  };

  const handleAdd = async (payload) => {
    const data = {
        name: payload.name,
        amount: payload.amount,
        category: payload.category,
        created_at: payload.createdAt
    };

    await addTransaction(data);
    setIsAdding(false);
  };

  const addButton = (
  <button className="bg-blue-500 text-white rounded shadow h-10 hover:opacity-80 cursor-pointer p-4" onClick={handleClick}>
    <div className="flex flex-row justify-center items-center h-full">
      <p>Add New Transaction</p>
    </div>
  </button>
  );


  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Something went wrong.</div>;
  }

  return (
    <>
      <div className="flex flex-col gap-4 flex-wrap">
        {isAdding ? (<TransactionForms onAdd={handleAdd} onCancel={handleCancel} />): (addButton)}
        {transactions.length === 0 ? (
          <p>No transactions found</p>
        ) : (
          transactions.map((transaction) => (
            <Transaction key={transaction.id} transactionData={transaction} />
          ))
        )}
      </div>
    </>
  )
}

export default TransactionList;