import React, { useState } from 'react';
import axios from 'axios';

const URLShortener = () => {
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/api/shorten', {
        originalUrl: originalUrl
      });
      setShortenedUrl(response.data.shortenedUrl);
      setError('');
    } catch (err) {
      setError('Failed to shorten URL');
    }
  };

  return (
    <div className="URLShortener">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={originalUrl}
          onChange={(e) => setOriginalUrl(e.target.value)}
          placeholder="Enter your URL"
          required
        />
        <button type="submit">Shorten URL</button>
      </form>
      {shortenedUrl && (
        <div className="shortened-url">
          <h2>Shortened URL</h2>
          <a href={shortenedUrl} target="_blank" rel="noopener noreferrer">
            {shortenedUrl}
          </a>
        </div>
      )}
      {error && <p className="error">{error}</p>}
    </div>
  );
};

export default URLShortener;
