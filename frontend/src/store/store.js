import { configureStore } from '@reduxjs/toolkit';
import financesReducer from './financesSlice';

export const store = configureStore({
  reducer: {
    finances: financesReducer,
  },
});