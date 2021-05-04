import './App.css';
import './style.css';
import './tailwind.css';
import { Switch, Route } from 'react-router-dom';
import { LoginPage } from './features/LoginPage';
import { RegisterPage } from './features/RegisterPage';

function App() {
  return (
    <Switch>
      <Route exact path="/">
        <LoginPage />
      </Route>
      <Route path="/register">
        <RegisterPage />
      </Route>
      <Route path="/home">
        <div>hi, this is home page</div>
      </Route>
    </Switch>
    // <div className="App">
    //   <header className="App-header">
    //     <p>
    //       Edit <code>src/App.js</code> and save to reload.
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
