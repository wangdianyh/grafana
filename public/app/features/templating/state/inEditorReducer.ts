import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { VariablePayload } from './actions';

const inEditorReducerSlice = createSlice({
  name: 'templating/inEditor',
  initialState: null,
  reducers: {
    changeToEditorEditMode: (state, action: PayloadAction<VariablePayload>) => {
      return action.payload.name;
    },
    changeToEditorListMode: (state, action: PayloadAction<undefined>) => {
      return null;
    },
  },
});

export const inEditorReducer = inEditorReducerSlice.reducer;

export const { changeToEditorListMode, changeToEditorEditMode } = inEditorReducerSlice.actions;
