import { useState } from "react";

function Transaction({ transactionData, onUpdate, onDelete }) {

    const [isEditing, setIsEditing] = useState(false);

    const [name, setName] = useState(transactionData.name);
    const [amount, setAmount] = useState(transactionData.amount.toString());
    const [category, setCategory] = useState(transactionData.category);
    const [createdAt, setCreatedAt] = useState(new Date(transactionData.created_at).toISOString().split('T')[0]);

    const editingForm = (
        <form className="p-4 border rounded shadow flex flex-row flex-none justify-between items-center">
            <div className="p-2 flex flex-row flex-none justify-between items-center w-4/5">
                <div className="flex flex-row gap-4 items-center">
                    <select className="text-gray-600 w-16" value={category} onChange={(e) => setCategory(e.target.value)}>
                        <option>Fixed Cost</option>
                        <option>Comfort</option>
                        <option>Goals</option>
                        <option>Pleasures</option>
                        <option>Financial Freedom</option>
                        <option>Knowledge</option>
                        <option>Income</option>
                    </select>
                    <div className="flex flex-col">
                        <input className="font-semibold" value={name} onChange={(e) => setName(e.target.value)} />
                        <input type="number" step="0.01" className="text-green-600" value={amount} onChange={(e) => setAmount(e.target.value)} />
                    </div>
                </div>
                <input type="date" className="text-gray-400" value={createdAt} onChange={(e) => setCreatedAt(e.target.value)} />
            </div>
            <div className="flex flex-row justify-end items-center flex-wrap gap-1">
                <button className="bg-green-500 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={() => {
                    onUpdate(transactionData.id, {
                        name: name,
                        amount: parseFloat(amount),
                        category: category,
                        created_at: new Date(createdAt).toISOString()
                    });
                    setIsEditing(false);
                }}>
                    <p>Save</p>
                </button>
                <button className="bg-red-500 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={() => setIsEditing(false)}>
                    <p>Cancel</p>
                </button>
            </div>
        </form>
    )

    const defaultView = (
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
                <button className="bg-yellow-400 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={() => setIsEditing(true)}>
                    <p>Edit</p>
                </button>
                <button className="bg-red-500 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={() => onDelete(transactionData.id)}>
                    <p>Delete</p>
                </button>
            </div>
        </div>
    )

    return (isEditing ? editingForm : defaultView)
}

export default Transaction;