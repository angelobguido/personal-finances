import { create } from 'zustand';
import { getTransactions } from '../services/api';

const useFinanceStore = create((set, get) => ({
  loading: true,
  error: false,
  transactions: [],

  async loadAllData() {
    set({ loading: true, error: false });
    try{
        const transactions = await getTransactions();
        set({ transactions });
    } catch {
        set({ error: true });
    }
    set({ loading: false });
  },

}));

export {useFinanceStore};