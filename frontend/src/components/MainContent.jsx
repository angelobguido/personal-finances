import TransactionView from './TransactionView.jsx'

function MainContent({ transactions, categories, onChangeTransactions, onChangeCategories }) {

  return (
    <main className="p-8">
        <TransactionView transactions={transactions} onChangeTransactions={onChangeTransactions} />
    </main>
  )
}

export default MainContent
