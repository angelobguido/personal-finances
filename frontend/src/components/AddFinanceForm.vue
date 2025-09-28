<template>
  <div>
    <h2>Add New Finance Item</h2>
    <form @submit.prevent="handleSubmit">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input type="text" class="form-control" id="name" v-model="newFinance.name" required>
      </div>
      <div class="mb-3">
        <label for="amount" class="form-label">Amount</label>
        <input type="number" step="0.01" class="form-control" id="amount" v-model.number="newFinance.amount" required>
      </div>
      <div class="mb-3">
        <label for="type" class="form-label">Type</label>
        <select class="form-select" id="type" v-model="newFinance.type">
          <option>Income</option>
          <option>Expense</option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Add Item</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const API_URL = 'http://localhost:8090/finances';

// This defines a custom event that the component can send to its parent
const emit = defineEmits(['finance-added']);

// A reactive object to hold the form's data
const newFinance = ref({
  name: '',
  amount: 0,
  type: 'Income', // Default value
});

async function handleSubmit() {
  if (!newFinance.value.name) {
    alert('Please enter a name.');
    return;
  }

  try {
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newFinance.value),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    // Emit the event to notify the parent component
    emit('finance-added');

    // Reset the form for the next entry
    newFinance.value = { name: '', amount: 0, type: 'Income' };

  } catch (error) {
    console.error('Error adding finance item:', error);
  }
}
</script>