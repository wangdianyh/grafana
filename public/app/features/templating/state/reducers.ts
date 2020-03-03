import { combineReducers } from '@reduxjs/toolkit';
import {
  initialState as initialOptionPickerState,
  optionsPickerReducer,
  OptionsPickerState,
} from '../pickers/OptionsPicker/reducer';
import { initialVariableEditorState, variableEditorReducer, VariableEditorState } from '../editor/reducer';
import { inEditorReducer } from './inEditorReducer';
import { variablesReducer } from './variablesReducer';
import { VariableModel } from '../variable';

export interface TemplatingState {
  variables: Record<string, VariableModel>;
  optionsPicker: OptionsPickerState;
  editor: VariableEditorState;
  inEditor: string | null;
}

export const initialTemplatingState: TemplatingState = {
  variables: {},
  optionsPicker: initialOptionPickerState,
  editor: initialVariableEditorState,
  inEditor: null,
};

export default {
  templating: combineReducers({
    inEditor: inEditorReducer,
    optionsPicker: optionsPickerReducer,
    editor: variableEditorReducer,
    variables: variablesReducer,
  }),
};
