function TransactionForms() {

    const handleAdd = async (event) => {
        event.preventDefault();
        const formData = new FormData(event.target.form);
        const data = {
            name: formData.get('name'),
            amount: parseFloat(formData.get('amount')),
            category: formData.get('category'),
            created_at: new Date(formData.get('created_at')).toISOString()
        };

        console.log('Submitting data:', data);

        try {
            const response = await fetch('/finances', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
            });

            if (!response.ok) {
            throw new Error('Failed to add finance', response.statusText);
            }

            const result = await response.json();
            console.log('Success:', result);
            event.target.form.reset();
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div className="p-4 flex flex-col flex-none items-center">
            <form className="p-4 border rounded shadow flex flex-col w-max">
                <div className="flex flex-row">
                    <p>Name:</p>
                    <input type="text" name="name" required /> 
                </div>
                <div className="flex flex-row">
                    <p>Amount:</p>
                    <input type="number" step="0.01" name="amount" required />
                </div>
                <div className="flex flex-row">
                    <p>Category:</p>
                    <select name="category" required>
                        <option selected>Fixed Cost</option>
                        <option>Comfort</option>
                        <option>Goals</option>
                        <option>Pleasures</option>
                        <option>Financial Freedom</option>
                        <option>Knowledge</option>
                        <option>Income</option>
                    </select>
                </div>
                <div>
                    <p>Created At:</p>
                    <input type="date" name="created_at" required />
                </div>
                <button type="button" onClick={handleAdd}>Add</button>
            </form>
        </div>
    );
}

export default TransactionForms;