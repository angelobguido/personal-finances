<template>
  <div>
    <AddFinanceForm @finance-added="fetchFinances" />

    <hr class="my-4"> <h1>Finance List</h1>
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
          <td>${{ finance.amount.toFixed(2) }}</td>
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
import AddFinanceForm from './AddFinanceForm.vue';

const API_URL = 'http://localhost:8090/finances';
const finances = ref([]);

// This function is now reusable for initial load and for refreshing!
async function fetchFinances() {
  try {
    const response = await fetch(API_URL);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    // Assuming your Go backend returns the data directly or in a field
    finances.value = Array.isArray(data) ? data : data.data;
  } catch (error) {
    console.error('Error fetching finance data:', error);
  }
}

function handleUpdate(financeItem) {
  console.log('Update item:', financeItem);
}

async function handleDelete(financeId) {
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
    finances.value = finances.value.filter(f => f.id !== financeId);
  } catch (error) {
    console.error('Error deleting finance item:', error);
  }
}

onMounted(fetchFinances);
</script>