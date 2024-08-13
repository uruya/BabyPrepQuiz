import React from 'react';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Welcome from './components/Welcome';
import Quiz from './quiz';
import Login from './components/Login';
import Register from './components/Register';

// import logo from './logo.svg';
// import './App.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path='/' Component={Welcome}></Route>
        <Route path='/login' Component={Login}></Route>
        <Route path='/Register' Component={Register}></Route>
        <Route path='/quiz' Component={Quiz}></Route>
      </Routes>
    </Router>
    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.tsx</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React
    //     </a>
    //   </header>
    // </div>
  );
}

export default App;
