import { useState } from "react";

function TransactionForms({onAdd, onCancel, categories}) {

    const [name, setName] = useState('Nova Transação');
    const [amount, setAmount] = useState('100.00');
    const [categoryId, setCategoryId] = useState(categories.length > 0 ? categories[0].id : null);
    const [createdAt, setCreatedAt] = useState(new Date().toISOString().split('T')[0]);

    const handleAdd = () => {
        onAdd({
            name: name,
            amount: parseFloat(amount),
            categoryId: categoryId,
            createdAt: createdAt
        });
    };

    const handleOnNameChange = (e) => setName(e.target.value);
    const handleOnAmountChange = (e) => setAmount(e.target.value);
    const handleOnCreatedAtChange = (e) => setCreatedAt(e.target.value);

    return (
        <form className="bg-blue-500 text-white rounded shadow flex flex-col flex-none items-center p-4">
            <div className="flex flex-row flex-wrap gap-5 p-2">
                <div className="flex flex-row gap-2">
                    <p>Name:</p>
                    <input value={name} onChange={handleOnNameChange} className="bg-blue-400" type="text" name="name" placeholder="Enter Name" required /> 
                </div>
                <div className="flex flex-row gap-2">
                    <p>Amount:</p>
                    <input value={amount} onChange={handleOnAmountChange} className="bg-blue-400" type="number" step="0.01" name="amount" placeholder="Enter Amount" required />
                </div>
                <div className="flex flex-row gap-2">
                    <p>Category:</p>
                    <select className="text-gray-600 w-16" value={categoryId} onChange={(e) => setCategoryId(parseInt(e.target.value))}>
                        {categories.map((cat) => (
                            <option key={cat.id} value={cat.id}>{cat.name}</option>
                        ))}
                    </select>
                </div>
                <div className="flex flex-row gap-2">
                    <p>Created At:</p>
                    <input value={createdAt} onChange={handleOnCreatedAtChange} className="bg-blue-400" type="date" name="created_at" required />
                </div>
            </div>
            <div className="p-2 flex flex-row justify-center gap-2">
                <button className="bg-blue-800 cursor-pointer hover:opacity-80 rounded shadow p-2" type="button" onClick={handleAdd}>Add</button>
                <button className="bg-gray-600 cursor-pointer hover:opacity-80 rounded shadow p-2" type="button" onClick={onCancel}>Cancel</button>
            </div>
        </form>
    );
}

export default TransactionForms;