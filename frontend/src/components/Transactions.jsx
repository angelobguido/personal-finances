import { useEffect, useState } from 'react';
import Forms from './Forms.jsx';

function Transactions() {
  const [finances, setFinances] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchFinances = async () => {
      try {
        setLoading(true);
        const response = await fetch('/finances');

        if (!response.ok) {
          throw new Error('Failed to fetch finances');
        }

        const result = await response.json();
        setFinances(result.data || []);
      } catch (err) {
        setError(err.message);
        console.error('Error fetching finances:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchFinances();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <>
      <h2>Transactions</h2>
      <Forms />
      <ul>
        {finances.length === 0 ? (
          <li>No transactions found</li>
        ) : (
          finances.map((finance) => (
            <li key={finance.id}>
              <div>
                <span>{finance.name}</span>
                <span>R$ {finance.amount.toFixed(2)}</span>
                <span>{finance.category}</span>
                <span>{new Date(finance.created_at).toLocaleDateString()}</span>
              </div>
            </li>
          ))
        )}
      </ul>
    </>
  )
}

export default Transactions