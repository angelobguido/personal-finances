import Header from './components/Header.jsx'
import MainContent from './components/MainContent.jsx'
import { useEffect } from 'react';
import { useFinanceStore } from './store/useFinanceStore.js';

function App() {

  const {loadAllData} = useFinanceStore();

  useEffect(() => {
    loadAllData();
  }, []);

  return (
    <>
    <Header />
    <MainContent />
    </>
  )
}

export default App
