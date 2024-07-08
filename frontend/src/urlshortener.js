import React, { useState } from 'react';
import CopyButton from './copybutton.js'; // Import the CopyButton component
import SlidingPanel from './slidingpanel.js'; // Import the SlidingPanel component

const URLShortener = () => {
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState('');
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('/api/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ originalUrl })
      });

      if (!response.ok) {
        throw new Error('Failed to shorten URL');
      }

      const data = await response.json();
      setShortenedUrl(data.shortenedUrl);
      setError('');
    } catch (err) {
      setError('Failed to shorten URL');
    }
  };

  return (
    <div className="URLShortener">
      <button onClick={() => setIsPanelOpen(true)}>Show All URLs</button>
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
          <CopyButton textToCopy={shortenedUrl} />
        </div>
      )}
      {error && <p className="error">{error}</p>}
      <SlidingPanel isOpen={isPanelOpen} onClose={() => setIsPanelOpen(false)} />
    </div>
  );
};

export default URLShortener;
