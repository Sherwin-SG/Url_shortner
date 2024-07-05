import React, { useState } from 'react';

const CopyButton = ({ textToCopy }) => {
  const [copySuccess, setCopySuccess] = useState('');

  const copyToClipboard = () => {
    navigator.clipboard.writeText(textToCopy).then(() => {
      setCopySuccess('Copied!');
      setTimeout(() => setCopySuccess(''), 2000); // Clear success message after 2 seconds
    }, (err) => {
      setCopySuccess('Failed to copy!');
    });
  };

  return React.createElement('div', null,
    React.createElement('button', { onClick: copyToClipboard }, 'Copy URL'),
    copySuccess && React.createElement('span', null, copySuccess)
  );
};

export default CopyButton;
