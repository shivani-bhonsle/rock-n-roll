import { useDispatch, useSelector } from "react-redux";
import { decrement, increment, reset } from "../redux/store";

const Counter = () => {
  const count = useSelector((state) => state.counter.count);
  const dispatch = useDispatch();

  return (
    <>
      <div>
        <h4>Count : {count}</h4>
        <button name="decrement" onClick={() => dispatch(decrement())}>
          -
        </button>
        <button name="reset" onClick={() => dispatch(reset())}>
          Reset
        </button>
        <button name="increment" onClick={() => dispatch(increment())}>
          +
        </button>
      </div>
    </>
  );
};

export default Counter;
