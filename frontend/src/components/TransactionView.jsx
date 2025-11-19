import TransactionList from './TransactionList.jsx';

function TransactionView() {

  return (
    <div className="w-3/4 mx-auto">
      <div className="flex flex-col">
        <h2 className="text-2xl font-bold pb-4">Transactions</h2>
      </div>
      <TransactionList />
    </div>
  )
}

export default TransactionView;