function Transaction({ transactionData }) {
    
    return (
        <div className="p-4 border rounded shadow flex flex-col gap-2 flex-none w-48 ">
            <span className="font-semibold">{transactionData.name}</span>
            <span className="text-green-600">R$ {transactionData.amount.toFixed(2)}</span>
            <span className="text-gray-600">{transactionData.category}</span>
            <span className="text-gray-400">{new Date(transactionData.created_at).toLocaleDateString()}</span>
        </div>
    )
}

export default Transaction;