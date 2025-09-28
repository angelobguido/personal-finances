<template>
    <div>
        <h1>Finance List</h1>

        <table class="table table-striped table-hover">
            <thead>
                <tr>
                    <th scope="col">ID</th>
                    <th scope="col">Name</th>
                    <th scope="col">Amount</th>
                    <th scope="col">Type</th>
                    <th scope="col">Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="finance in finances" :key="finance.id">
                    <th scope="row">{{ finance.id }}</th>
                    <td>{{ finance.name }}</td>
                    <td>{{ finance.amount }}</td>
                    <td>{{ finance.type }}</td>
                    <td>
                        <button class="btn btn-warning btn-sm me-2" @click="handleUpdate(finance)">
                            Update ✏️
                        </button>
                        <button class="btn btn-danger btn-sm" @click="handleDelete(finance.id)">
                            Delete 🗑️
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const API_URL = 'http://localhost:8090/finances';
const finances = ref([]);

// Fetches the initial list of finances
async function fetchFinances() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        finances.value = data.data;
    } catch (error) {
        console.error('Error fetching finance data:', error);
    }
}

// 3. Create a method to handle the update action
function handleUpdate(financeItem) {
    console.log('Update item:', financeItem);
    // TODO: Implement your update logic here.
    // This could involve opening a modal with a form pre-filled
    // with `financeItem` data or navigating to a separate edit page.
}

// 4. Create a method to handle the delete action
async function handleDelete(financeId) {
    // A confirmation dialog is good practice for destructive actions
    if (!confirm('Are you sure you want to delete this item?')) {
        return;
    }

    try {
        const response = await fetch(`${API_URL}/${financeId}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        // To update the UI, filter out the deleted item from the local array
        finances.value = finances.value.filter(f => f.id !== financeId);
        console.log(`Item with ID ${financeId} deleted successfully.`);

    } catch (error) {
        console.error('Error deleting finance item:', error);
    }
}

onMounted(fetchFinances);
</script>