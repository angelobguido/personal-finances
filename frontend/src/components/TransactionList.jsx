import Transaction from './Transaction.jsx';
import TransactionForms from './TransactionForms.jsx';
import { useState } from 'react';

function TransactionList({ transactions, onChangeTransactions }) {
  
  const [isAdding, setIsAdding] = useState(false);
  
  const handleUpdate = async (id, payload) => {
    await onChangeTransactions.update(id, payload);
  };

  const handleDelete = async (id) => {
    await onChangeTransactions.delete(id);
  };

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
        category_id: payload.categoryId,
        created_at: payload.createdAt
    };

    await onChangeTransactions.create(data);
    setIsAdding(false);
  };

  const addButton = (
  <button className="bg-blue-500 text-white rounded shadow h-10 hover:opacity-80 cursor-pointer p-4" onClick={handleClick}>
    <div className="flex flex-row justify-center items-center h-full">
      <p>Add New Transaction</p>
    </div>
  </button>
  );

  return (
    <>
      <div className="flex flex-col gap-4 flex-wrap">
        {isAdding ? (<TransactionForms onAdd={handleAdd} onCancel={handleCancel} />): (addButton)}
        {transactions.length === 0 ? (
          <p>No transactions found</p>
        ) : (
          transactions.map((transaction) => (
            <Transaction key={transaction.id} transactionData={transaction} onUpdate={handleUpdate} onDelete={handleDelete} />
          ))
        )}
      </div>
    </>
  )
}

export default TransactionList;