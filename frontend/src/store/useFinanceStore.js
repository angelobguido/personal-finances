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

  async addTransaction(data) {
    set({ loading: true, error: false });
    try{
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
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }

    await get().loadAllData();
  },

}));

export {useFinanceStore};