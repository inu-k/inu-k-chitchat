import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import Navbar from './components/NavBar.jsx';
import ThreadsPage from './pages/ThreadsPage.jsx';
import { CreateThreadPage } from './pages/CreateThreadPage';
import { Routes, Route, Link } from 'react-router-dom';
import PostsPage from './pages/PostsPage.jsx';
import { fetchData } from './functions/utils.jsx';
import LoginForm from './components/LoginForm.jsx';
import SignupForm from './components/SignupForm.jsx';

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
            <div>
              <p>
                <Link to='/threads'>
                  See Threads
                </Link>
              </p>
              <p>
                <Link to='/create_thread'>
                  Create Thread
                </Link>
              </p>
            </div>
          </div>
        } />
        <Route path="/threads" element={<ThreadsPage />} />
        <Route path="/posts" element={<PostsPage />} />
        <Route path="/create_thread" element={<CreateThreadPage />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/signup" element={<SignupForm />} />
      </Routes>


    </div>
  );
}

export default App;
