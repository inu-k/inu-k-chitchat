import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import Navbar from './components/NavBar.jsx';
import ThreadsPage from './pages/ThreadsPage.jsx';
import { Routes, Route, Link } from 'react-router-dom';
import PostsPage from './pages/PostsPage.jsx';
import { fetchData } from './functions/utils.jsx';

function App() {
  return (
    <div className="App">
      <header>
        <Navbar />
      </header>

      <Routes>
        <Route path="/" element={
          <div>
            <p>
              Welcome to inu-k-chitchat!
            </p>
            <p>
              <Link to='/threads'>
                See Threads
              </Link>
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
