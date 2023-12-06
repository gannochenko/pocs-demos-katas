import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const onRunClick = async () => {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)

    console.log(result);
    go.run(result.instance);

    console.log('!!!');
    // @ts-expect-error Need to extend the window interface later
    const res = window.foo({name: "Obi", rank: "Master"});
    console.log(res);
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
