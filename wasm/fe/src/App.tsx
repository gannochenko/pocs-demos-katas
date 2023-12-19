import React, {useState} from 'react';
import logo from './logo.svg';
import './App.css';
import {getWASM, Item} from "./wasm";

function App() {
  const [items, setItems] = useState<Item[]>([]);

  const onRunClick = async () => {
    const wasmMethods = await getWASM();

    const result = wasmMethods.getItems({
      amount: 10,
    });

    console.log(result);
    setItems(result.items);
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <button onClick={onRunClick}>Get items!</button>
        <div>
          {
            items.map(item => {
              return (<div key={item.id}>{item.title} delivered on {item.date}</div>);
            })
          }
        </div>
      </header>
    </div>
  );
}

export default App;
