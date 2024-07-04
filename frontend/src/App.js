// src/App.js

import React from 'react';
import URLShortener from './urlshortener.js';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>URL Shortener</h1>
      </header>
      <main className="App-main">
        <URLShortener />
      </main>
    </div>
  );
}

export default App;
