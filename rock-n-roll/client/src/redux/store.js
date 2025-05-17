import { configureStore, createSlice } from "@reduxjs/toolkit";
import userReducer from "./asyncThunk"; // Must be default export of userSlice.reducer

const counterSlice = createSlice({
  name: "counter",
  initialState: { count: 0 },
  reducers: {
    increment: (state) => { state.count += 1; },
    decrement: (state) => { state.count -= 1; },
    reset: (state) => { state.count = 0; },
  },
});

export const { increment, decrement, reset } = counterSlice.actions;

const store = configureStore({
  reducer: {
    counter: counterSlice.reducer,
    user: userReducer, // ✅ Must match useSelector((state) => state.user)
  },
});

export default store;
