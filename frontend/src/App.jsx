import Header from './components/Header.jsx'
import MainContent from './components/MainContent.jsx'
import { useState, useEffect } from 'react';
import { getTransactions, createTransaction, updateTransaction, deleteTransaction, getCategories, createCategory, updateCategory, deleteCategory } from './services/api.js';

function App() {

  const [categories, setCategories] = useState([]);
  const [transactions, setTransactions] = useState([]);
  
  const loadCategories = async () => {
    const categoriesData = await getCategories();
    setCategories(categoriesData);
  }

  const loadTransactions = async () => {
    const transactionsData = await getTransactions();
    setTransactions(transactionsData);
  }

  useEffect(async () => {
    await loadCategories();
    await loadTransactions();
  }, []);

  const handleChangeTransactions = {
    create: async (data) => {
      await createTransaction(data);
      await loadTransactions();
    },
    update: async (id, payload) => {
      await updateTransaction(id, payload);
      await loadTransactions();
    },
    delete: async (id) => {
      await deleteTransaction(id);
      await loadTransactions();
    }
  };

  const handleChangeCategories = {
    create: async (data) => {
      await createCategory(data);
      await loadCategories();
    },
    update: async (id, payload) => {
      await updateCategory(id, payload);
      await loadCategories();
    },
    delete: async (id) => {
      await deleteCategory(id);
      await loadCategories();
    }
  };

  return (
    <>
    <Header />
    <MainContent categories={categories} transactions={transactions} onChangeTransactions={handleChangeTransactions} onChangeCategories={handleChangeCategories}  />
    </>
  )
}

export default App
