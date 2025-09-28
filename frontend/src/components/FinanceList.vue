<template>
    <div>
        <h1>Finance List</h1>

        <table class="table">
            <thead>
                <tr>
                    <th scope="col">ID</th>
                    <th scope="col">Name</th>
                    <th scope="col">Amount</th>
                    <th scope="col">Type</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="finance in finances" :key="finance.id">
                    <th scope="row">{{ finance.id }}</th>
                    <td>{{ finance.name }}</td>
                    <td>{{ finance.amount }}</td>
                    <td>{{ finance.type }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const API_URL = 'http://localhost:8090/finances';
const finances = ref([]);

async function fetchFinances() {
    try {
        const response = await fetch(`${API_URL}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        finances.value = data.data;
    } catch (error) {
        console.error('Error fetching finance data:', error);
    }
}

onMounted(fetchFinances);

</script>