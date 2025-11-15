async function createFinance(data) {
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
        console.error('Error creating finance:', error);
        throw error;
    }
}

async function updateFinance(id, data) {
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
        console.error('Error updating finance:', error);
        throw error;
    }
}

async function getFinances() {
    try {
        const response = await fetch('/finances');
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching finances:', error);
        throw error;
    }
}

async function getFinance(id) {
    try {
        const response = await fetch(`/finances/${id}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error fetching finance:', error);
        throw error;
    }
}

async function deleteFinance(id) {
    try {
        const response = await fetch(`/finances/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        return (await response.json()).data;
    } catch (error) {
        console.error('Error deleting finance:', error);
        throw error;
    }
}

export {
    createFinance,
    updateFinance,
    getFinances,
    getFinance,
    deleteFinance
};