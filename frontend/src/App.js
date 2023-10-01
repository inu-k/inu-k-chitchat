import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import Navbar from './components/NavBar.jsx';
import ThreadsPage from './pages/ThreadsPage.jsx';
import { Routes, Route, Link } from 'react-router-dom';
import PostsPage from './pages/PostsPage.jsx';

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

  return (
    <div className="App">
      <header>
        <Navbar />
      </header>

      <Routes>
        <Route path="/" element={
          <div>
            <img src={logo} className="App-logo" alt="logo" />
            <p>
              <Link to='/threads'>
                See Threads
              </Link>
            </p>

            <p>
              Your name is {jsonData.name}, Your message is "{jsonData.message}".
            </p>
          </div>
        } />
        <Route path="/threads" element={<ThreadsPage />} />
        <Route path="/posts" element={<PostsPage />} />
      </Routes>


    </div>
  );
}

export default App;
