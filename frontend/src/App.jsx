import Header from './components/Header.jsx'
import MainContent from './components/MainContent.jsx'
import { useState, useEffect } from 'react';
import { getTransactions, createTransaction, updateTransaction, deleteTransaction, getCategories, createCategory, updateCategory, deleteCategory } from './services/api.js';

function App() {

  const [categories, setCategories] = useState([]);
  const [transactions, setTransactions] = useState([]);
  
  useEffect(async () => {
    const categoriesData = await getCategories();
    setCategories(categoriesData);
    const transactionsData = await getTransactions();
    setTransactions(transactionsData);
  }, []);

  const handleChangeTransactions = {
    create: async (data) => {
      const newTransaction = await createTransaction(data);
      setTransactions([...transactions, newTransaction]);
    },
    update: async (id, payload) => {
      const updatedTransaction = await updateTransaction(id, payload);
      setTransactions(transactions.map((transaction) => (transaction.id === id ? updatedTransaction : transaction)));
    },
    delete: async (id) => {
      await deleteTransaction(id);
      setTransactions(transactions.filter((transaction) => transaction.id !== id));
    }
  };

  const handleChangeCategories = {
    create: async (data) => {
      const newCategory = await createCategory(data);
      setCategories([...categories, newCategory]);
    },
    update: async (id, payload) => {
      const updatedCategory = await updateCategory(id, payload);
      setCategories(categories.map((category) => (category.id === id ? updatedCategory : category)));
    },
    delete: async (id) => {
      await deleteCategory(id);
      setCategories(categories.filter((category) => category.id !== id));
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
