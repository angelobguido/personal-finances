function Transaction({ transactionData, onUpdate, onDelete }) {

    return (
        <div className="p-4 border rounded shadow flex flex-row flex-none justify-between items-center">
            <div className="p-2 flex flex-row flex-none justify-between items-center w-4/5">
                <div className="flex flex-row gap-4 items-center">
                    <span className="text-gray-600 w-16">{transactionData.category}</span>
                    <div className="flex flex-col">
                        <span className="font-semibold">{transactionData.name}</span>
                        <span className="text-green-600">R$ {transactionData.amount.toFixed(2)}</span>
                    </div>
                </div>
                <span className="text-gray-400">{new Date(transactionData.created_at).toLocaleDateString()}</span>
            </div>
            <div className="flex flex-row justify-end items-center flex-wrap gap-1">
                <button className="bg-yellow-400 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" >
                    <p>Edit</p>
                </button>
                <button className="bg-red-500 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={() => onDelete(transactionData.id)}>
                    <p>Delete</p>
                </button>
            </div>
        </div>
    )
}

export default Transaction;