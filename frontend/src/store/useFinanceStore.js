import { create } from 'zustand';
import { getTransactions, createTransaction, updateTransaction, deleteTransaction } from '../services/api';

const useFinanceStore = create((set, get) => ({
  loading: true,
  error: false,
  transactions: [],

  async loadAllData() {
    set({ loading: true, error: false });
    try{
      const transactions = await getTransactions();
      set({ transactions });
    } 
    catch {
      set({ error: true });
    }
    finally {
      set({ loading: false });
    }
    
  },

  async addTransaction(data) {
    set({ loading: true, error: false });
    try{
      await createTransaction(data);
      await get().loadAllData();
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }
  },

  async updateTransaction(id, data) {
    set({ loading: true, error: false });
    try{
      await updateTransaction(id, data);
      await get().loadAllData();
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }
  },

  async deleteTransaction(id) {
    set({ loading: true, error: false });
    try{
      await deleteTransaction(id);
      await get().loadAllData();
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }
  },

}));

export {useFinanceStore};