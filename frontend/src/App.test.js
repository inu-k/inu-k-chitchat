import { render, screen } from '@testing-library/react';
import React from 'react';
import App from './App.js';
import { BrowserRouter } from 'react-router-dom';

test('renders See threads link', () => {
  render(<React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>);
  const linkElement = screen.getByText(/See Threads/i);
  expect(linkElement).toBeInTheDocument();
});
