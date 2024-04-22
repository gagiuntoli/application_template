import React, {useState} from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [data, setData] = useState([]);

  const fetchData = async () => {
    try {
      const response = await axios.get('http://localhost:8080/stories_titles'); 
      setData(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <button onClick={fetchData} style={{ fontSize: '20px' }}>
          {'Fetch Data'}
        </button>
        {data && (
          <ul>
            {data.map((item, index) => (
              <li key={index}>{item}</li>
            ))}
          </ul>
        )}
      </header>
    </div>
  );
}

export default App;
