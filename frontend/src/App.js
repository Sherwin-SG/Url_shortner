
import React from 'react';
import URLShortener from './urlshortener.js';
import Footer from './Footer.js';
import './App.css';

function App() {
  return (
    <div> <h1>URL Shortener</h1></div>,
    <div className="App">
      <header className="App-header">
        <h1>URL Shortener</h1>
      </header>
      <main className="App-main">
        <URLShortener />
      </main>
      <Footer />
    </div>
  );
}

export default App;
