function Transaction({ transactionData }) {
    
    return (
        <div className="p-4 border rounded shadow flex flex-row flex-none justify-between items-center">
            <div className="flex flex-row gap-4 items-center">
                <span className="text-gray-600 w-16">{transactionData.category}</span>
                <div className="flex flex-col">
                    <span className="font-semibold">{transactionData.name}</span>
                    <span className="text-green-600">R$ {transactionData.amount.toFixed(2)}</span>
                </div>
            </div>
            <span className="text-gray-400">{new Date(transactionData.created_at).toLocaleDateString()}</span>
        </div>
    )
}

export default Transaction;