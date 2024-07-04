import React, { useState } from 'react';

const CopyButton = ({ textToCopy }) => {
  const [copySuccess, setCopySuccess] = useState('');

  const copyToClipboard = () => {
    navigator.clipboard.writeText(textToCopy).then(() => {
      setCopySuccess('Copied!');
      setTimeout(() => setCopySuccess(''), 2000);
    }, (err) => {
      setCopySuccess('Failed to copy!');
    });
  };

  return (
    <div>
      <button onClick={copyToClipboard}>Copy URL</button>
      {copySuccess && <span>{copySuccess}</span>}
    </div>
  );
};

export default CopyButton;
