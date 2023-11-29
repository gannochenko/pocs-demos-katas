import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const onRunClick = () => {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      console.log(result);
      go.run(result.instance);
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={onRunClick}>Fucking execute!</button>
      </header>
    </div>
  );
}

export default App;
