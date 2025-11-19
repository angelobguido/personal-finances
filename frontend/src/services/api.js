async function createTransaction(data) {
    try {
        const response = await fetch('/finances', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error creating transaction:', error);
        throw error;
    }
}

async function updateTransaction(id, data) {
    try {
        const response = await fetch(`/finances/${id}`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error updating transaction:', error);
        throw error;
    }
}

async function getTransactions() {
    try {
        const response = await fetch('/finances');
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching transactions:', error);
        throw error;
    }
}

async function getTransaction(id) {
    try {
        const response = await fetch(`/finances/${id}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching transaction:', error);
        throw error;
    }
}

async function deleteTransaction(id) {
    try {
        const response = await fetch(`/finances/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
    } catch (error) {
        console.error('Error deleting transaction:', error);
        throw error;
    }
}

export {
    createTransaction,
    updateTransaction,
    getTransactions,
    getTransaction,
    deleteTransaction
};