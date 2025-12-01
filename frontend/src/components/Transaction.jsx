import { useState } from "react";

function Transaction({ transactionData, categories, onUpdate, onDelete }) {

    const [isEditing, setIsEditing] = useState(false);

    const [name, setName] = useState(transactionData.name);
    const [amount, setAmount] = useState(transactionData.amount.toString());
    const [categoryId, setCategoryId] = useState(transactionData.category_id);
    const [createdAt, setCreatedAt] = useState(transactionData.created_at);

    const getCategoryNameById = (id) => {
        const category = categories.find((cat) => cat.id === id);
        return category ? category.name : 'Unknown';
    }

    const editingForm = (
        <form className="p-4 border rounded shadow flex flex-row flex-none justify-between items-center bg-gray-100">
            <div className="p-2 flex flex-row flex-none justify-between items-center w-4/5">
                <div className="flex flex-row gap-4 items-center">
                    <select className="text-gray-600 w-16" value={categoryId} onChange={(e) => setCategoryId(parseInt(e.target.value))}>
                        {categories.map((cat) => (
                            <option key={cat.id} value={cat.id}>{cat.name}</option>
                        ))}
                    </select>
                    <div className="flex flex-col">
                        <input className="font-semibold" value={name} onChange={(e) => setName(e.target.value)} />
                        <input type="number" step="0.01" className="text-green-600" value={amount} onChange={(e) => setAmount(e.target.value)} />
                    </div>
                </div>
                <input type="date" className="text-gray-400" value={new Date(createdAt).toISOString().split('T')[0]} onChange={(e) => setCreatedAt(e.target.value)} />
            </div>
            <div className="flex flex-row justify-end items-center flex-wrap gap-1">
                <button className="bg-green-500 text-white rounded shadow h-8 hover:opacity-80 cursor-pointer p-2 flex items-center" onClick={(e) => {
                    e.preventDefault();
                    onUpdate(transactionData.id, {
                        name: name,
                        amount: parseFloat(amount),
                        category_id: categoryId,
                        created_at: createdAt
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
        <div className="p-4 border rounded shadow flex flex-row flex-none justify-between items-center bg-gray-100">
            <div className="p-2 flex flex-row flex-none justify-between items-center w-4/5">
                <div className="flex flex-row gap-4 items-center">
                    <span className="text-gray-600 w-16">{getCategoryNameById(transactionData.category_id)}</span>
                    <div className="flex flex-col">
                        <span className="font-semibold">{transactionData.name}</span>
                        <span className="text-green-600">R$ {transactionData.amount.toFixed(2)}</span>
                    </div>
                </div>
                <span className="text-gray-400">{transactionData.created_at}</span>
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