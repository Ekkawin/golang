import './App.css';
import './style.css';
import './tailwind.css';
import { Switch, Route } from 'react-router-dom';
import { LoginPage } from './features/LoginPage';

function App() {
  return (
    <Switch>
      <Route path="/">
        <LoginPage />
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
