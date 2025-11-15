function Transaction({ transactionData }) {
    
    return (
        <div>
            <span>{transactionData.name}</span>
            <span>R$ {transactionData.amount.toFixed(2)}</span>
            <span>{transactionData.category}</span>
            <span>{new Date(transactionData.created_at).toLocaleDateString()}</span>
        </div>
    )
}

export default Transaction;