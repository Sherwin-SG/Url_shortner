import React, { useState, useEffect } from 'react';

const SlidingPanel = ({ isOpen, onClose }) => {
  const [urls, setUrls] = useState([]);

  useEffect(() => {
    if (isOpen) {
      fetch('/api/urls')
        .then(response => response.json())
        .then(data => setUrls(data))
        .catch(err => console.error('Error fetching URLs:', err));
    }
  }, [isOpen]);

  const handleDelete = (shortenedUrl) => {
    fetch(`/api/delete/${shortenedUrl}`, {
      method: 'DELETE',
    })
      .then(response => {
        if (response.ok) {
          setUrls(prevUrls => prevUrls.filter(url => url.shortenedUrl !== shortenedUrl));
        } else {
          throw new Error('Failed to delete URL');
        }
      })
      .catch(err => console.error('Error deleting URL:', err));
  };

  return (
    <div className={`sliding-panel ${isOpen ? 'open' : ''}`}>
      <button className="close-button" onClick={onClose}>Close</button>
      <ul>
        {urls.map((url, index) => (
          <li key={index}>
            <a href={url.originalUrl} target="_blank" rel="noopener noreferrer">{url.shortenedUrl}</a>
            <span>{url.originalUrl}</span>
            <button onClick={() => handleDelete(url.shortenedUrl)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SlidingPanel;
