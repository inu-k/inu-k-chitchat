import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';

function App() {
  const [jsonData, setJsonData] = useState({});

  useEffect(() => {
    async function fetchData() {
      try {
        const response = await fetch('http://localhost:8999/index');

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setJsonData(data);
      } catch (error) {
        console.error('Error fetching data: ', error);
      }
    }

    fetchData();
  }, []);

  console.log('jsonData: ', jsonData);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>

        <p>
          Your name is {jsonData.name}, Your message is "{jsonData.message}".
        </p>
      </header>

    </div>
  );
}

export default App;
