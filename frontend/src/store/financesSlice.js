import { createSlice } from '@reduxjs/toolkit';
import { createFinance, updateFinance, deleteFinance, getFinance, getFinances } from '../services/api';

const initialState = {
    finances: []
}

export const financesSlice = createSlice({
    name: 'finances',
    initialState,
    reducers: {
        add: async (state, action) => {
            let data = action.payload;
            let finance = await createFinance(data);
            state.finances.push(finance);
            return finance;
        },
        remove: async (state, action) => {
            let id = action.payload;
            await deleteFinance(id);
            state.finances = state.finances.filter(item => item.id !== id);
        },
        update: async (state, action) => {
            let id = action.payload.id;
            let data = action.payload.data;
            let finance = await updateFinance(id, data);
            let index = state.finances.findIndex(item => item.id === finance.id);
            if (index !== -1) {
                state.finances[index] = finance;
            }
            return finance;
        },
        getAll: async (state) => {
            let finances = await getFinances();
            state.finances = finances;
            return finances;
        },
        get: async (state, action) => {
            let id = action.payload;
            let finance = state.finances.find(item => item.id === id);
            if (!finance) {
                finance = await getFinance(id);
                state.finances.push(finance);
            }
            return finance;
        }
    },
});

export const { add, remove, update, getAll } = financesSlice.actions;

export default financesSlice.reducer;