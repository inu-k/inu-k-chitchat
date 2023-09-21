import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import axios from 'axios';

function App() {
  const [jsonData, setJsonData] = useState({});

  useEffect(() => {
    axios.get('http://localhost:8999/index')
      .then((response) => {
        setJsonData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      })
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
