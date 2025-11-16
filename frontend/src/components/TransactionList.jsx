import Transaction from './Transaction.jsx';
import { useFinanceStore } from '../store/useFinanceStore.js';

function TransactionList() {
  const {transactions, loading, error} = useFinanceStore();
  
  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Something went wrong.</div>;
  }

  return (
    <>
      <div className="flex flex-col gap-4 flex-wrap">
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