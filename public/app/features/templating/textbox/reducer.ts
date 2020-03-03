import { createReducer } from '@reduxjs/toolkit';

import { TextBoxVariableModel, VariableHide, VariableOption } from '../variable';
import { getInstanceState } from '../state/types';
import { createTextBoxOptions } from './actions';
import { initialVariablesState } from '../state/variablesReducer';

export const initialTextBoxVariableModelState: TextBoxVariableModel = {
  global: false,
  index: -1,
  type: 'textbox',
  name: '',
  label: '',
  hide: VariableHide.dontHide,
  query: '',
  current: {} as VariableOption,
  options: [],
  skipUrlSync: false,
  initLock: null,
};

export const textBoxVariableReducer = createReducer(initialVariablesState, builder =>
  builder.addCase(createTextBoxOptions, (state, action) => {
    const instanceState = getInstanceState<TextBoxVariableModel>(state, action.payload.name);
    instanceState.options = [{ text: instanceState.query.trim(), value: instanceState.query.trim(), selected: false }];
    instanceState.current = instanceState.options[0];
  })
);
