import { Provider } from "react-redux";
import "./App.css";
import CreateStudent from "./pages/CreateStudent";
import store from "./redux/store";
import Counter from "./pages/Counter";
import Users from "./pages/Users";

function App() {
  return (
    <Provider store={store}>
      <div className="App">
        {/* <CreateStudent /> */}
        <Counter />
        <Users/>
      </div>
    </Provider>
  );
}

export default App;
