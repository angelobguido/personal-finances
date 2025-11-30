async function createTransaction(data) {
    try {
        const response = await fetch('/transactions', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json());
    } catch (error) {
        console.error('Error creating transaction:', error);
        throw error;
    }
}

async function updateTransaction(id, data) {
    try {
        const response = await fetch(`/transactions/${id}`, {
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
        const response = await fetch('/transactions');
        
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
        const response = await fetch(`/transactions/${id}`);
        
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
        const response = await fetch(`/transactions/${id}`, {
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

async function createCategory(data) {
    try {
        const response = await fetch('/categories', {
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
        console.error('Error creating category:', error);
        throw error;
    }
}

async function updateCategory(id, data) {
    try {
        const response = await fetch(`/categories/${id}`, {
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
        console.error('Error updating category:', error);
        throw error;
    }
}

async function getCategories() {
    try {
        const response = await fetch('/categories');
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching categories:', error);
        throw error;
    }
}

async function getCategory(id) {
    try {
        const response = await fetch(`/categories/${id}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching category:', error);
        throw error;
    }
}

async function deleteCategory(id) {
    try {
        const response = await fetch(`/categories/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
    } catch (error) {
        console.error('Error deleting category:', error);
        throw error;
    }
}

async function getReport() {
    try {
        const response = await fetch('/report');
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching report:', error);
        throw error;
    }
}

export {
    createTransaction,
    updateTransaction,
    getTransactions,
    getTransaction,
    deleteTransaction,
    createCategory,
    updateCategory,
    getCategories,
    getCategory,
    deleteCategory,
    getReport
};