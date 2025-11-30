import { create } from 'zustand';
import { 
  getTransactions, 
  createTransaction, 
  updateTransaction, 
  deleteTransaction,
  getCategories,
  createCategory,
  updateCategory,
  deleteCategory,
  getReport
} from '../services/api';

const useFinanceStore = create((set, get) => ({
  loading: true,
  error: false,
  transactions: [],
  categories: [],
  report: null,

  async loadAllData() {
    set({ loading: true, error: false });
    try{
      const [transactions, categories, report] = await Promise.all([
        getTransactions(),
        getCategories(),
        getReport()
      ]);
      set({ transactions, categories, report });
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

  async addCategory(data) {
    set({ loading: true, error: false });
    try{
      await createCategory(data);
      await get().loadAllData();
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }
  },

  async updateCategory(id, data) {
    set({ loading: true, error: false });
    try{
      await updateCategory(id, data);
      await get().loadAllData();
    } 
    catch {
      set({ error: true });
    } 
    finally {
      set({ loading: false });
    }
  },

  async deleteCategory(id) {
    set({ loading: true, error: false });
    try{
      await deleteCategory(id);
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