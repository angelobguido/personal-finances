import TransactionList from './TransactionList.jsx';

function TransactionView() {
  
  return (
    <div className="w-3/4 mx-auto">
      <h2 className="text-2xl font-bold pt-5 pb-4">Transactions</h2>
      <TransactionList />
    </div>
  )
}

export default TransactionView;