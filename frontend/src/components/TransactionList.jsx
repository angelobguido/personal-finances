import { useEffect, useState } from 'react';
import { getTransactions } from '../services/api.js';
import Transaction from './Transaction.jsx';

function TransactionList() {
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(async () => {
    setLoading(true);

    try {
      const response = await getTransactions();
      setTransactions(response);
    }
    catch (err) {
      setError(true);
    }
    finally {
      setLoading(false);
    }

  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Something went wrong.</div>;
  }

  return (
    <>
      <div className="flex flex-row gap-4 flex-wrap">
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