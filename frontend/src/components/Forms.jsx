function Forms() {

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
        <>
            <form>
                <div>
                    <input type="text" placeholder="Name" name="name" required /> 
                </div>
                <div>
                    <input type="number" step="0.01" placeholder="Amount" name="amount" required />
                </div>
                <div>
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
                    <label>Created At</label>
                    <input type="date" name="created_at" required />
                </div>
                <button type="button" onClick={handleAdd}>Add</button>
            </form>
        </>
    );
}

export default Forms;