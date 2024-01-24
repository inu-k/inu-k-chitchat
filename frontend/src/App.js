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
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (
    <div className="App">
      <header>
        <Navbar isLoggedIn={isLoggedIn} setIsLoggedIn={setIsLoggedIn} />
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
                {isLoggedIn ? (
                  <Link to='/create_thread'>
                    Create Thread
                  </Link>) : (
                  <div style={{ padding: "10px" }}>
                    <div>Please <Link to='/login'>login</Link> to create a thread.</div>
                  </div>
                )}
              </p>
            </div>
          </div>
        } />
        <Route path="/threads" element={<ThreadsPage />} />
        <Route path="/posts" element={<PostsPage isLoggedIn={isLoggedIn} />} />
        <Route path="/create_thread" element={<CreateThreadPage isLoggedIn={isLoggedIn} />} />
        <Route path="/login" element={<LoginForm setIsLoggedIn={setIsLoggedIn} />} />
        <Route path="/signup" element={<SignupForm />} />
      </Routes>


    </div>
  );
}

export default App;
