import React from 'react';
import ReactDOM from 'react-dom/client';
import './globals.css';
import {Routes, Route} from "react-router-dom";
import reportWebVitals from './reportWebVitals';
import {HomePage} from "./pages/HomePage";
import {ApplicationLayout, Providers} from "./components";
import {NotFoundPage} from "./pages/NotFoundPage";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
      <Providers>
          <ApplicationLayout>
              <Routes>
                  <Route index element={<HomePage />} />
                  <Route
                      path="*"
                      element={<NotFoundPage />}
                  />
              </Routes>
          </ApplicationLayout>
      </Providers>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
