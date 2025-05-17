import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

export const fetchUsers = createAsyncThunk("users/fetchUser", async () => {
  try {
    const response = await axios.get(`https://jsonplaceholder.typicode.com/users`);
    return response.data;
  } catch (err) {
    console.log(err);
  }
});

const userSlice = createSlice({
  name: "users",
  initialState: { user: [], loading: false, error: null },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUsers.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchUsers.fulfilled, (state, action) => {
        console.log("Fetched payload:", action.payload);
        state.user = action.payload;
        state.loading = false;
      })
      .addCase(fetchUsers.rejected, (state, action) => {
        state.error = action.error.message;
        state.loading = false;
      });
  },
});
export default userSlice.reducer;


