import { createReducer } from '@reduxjs/toolkit';
import cloneDeep from 'lodash/cloneDeep';

import {
  addVariable,
  changeVariableOrder,
  changeVariableProp,
  changeVariableType,
  duplicateVariable,
  removeInitLock,
  removeVariable,
  resolveInitLock,
  setCurrentVariableValue,
  storeNewVariable,
} from './actions';
import { VariableModel, VariableWithOptions } from '../variable';
import { ALL_VARIABLE_VALUE, getInstanceState, NEW_VARIABLE_NAME } from './types';
import { variableAdapters } from '../adapters';
import { changeVariableNameSucceeded, variableEditorUnMounted } from '../editor/reducer';
import { Deferred } from '../deferred';
import { changeToEditorEditMode } from './inEditorReducer';
import { initialVariablesState } from './variablesReducer';

export const sharedReducer = createReducer(initialVariablesState, builder =>
  builder
    .addCase(addVariable, (state, action) => {
      state[action.payload.name] = cloneDeep(variableAdapters.get(action.payload.type).initialState);
      state[action.payload.name] = {
        ...state[action.payload.name],
        ...action.payload.data.model,
      };
      state[action.payload.name].index = action.payload.data.index;
      state[action.payload.name].global = action.payload.data.global;
      state[action.payload.name].initLock = new Deferred();
    })
    .addCase(resolveInitLock, (state, action) => {
      const instanceState = getInstanceState(state, action.payload.name);
      instanceState.initLock?.resolve();
    })
    .addCase(removeInitLock, (state, action) => {
      const instanceState = getInstanceState(state, action.payload.name);
      instanceState.initLock = null;
    })
    .addCase(removeVariable, (state, action) => {
      delete state[action.payload.name];
      const variableStates = Object.values(state);
      for (let index = 0; index < variableStates.length; index++) {
        variableStates[index].index = index;
      }
    })
    .addCase(variableEditorUnMounted, (state, action) => {
      const variableState = state[action.payload.name];

      if (action.payload.name === NEW_VARIABLE_NAME && !variableState) {
        return;
      }

      if (state[NEW_VARIABLE_NAME]) {
        delete state[NEW_VARIABLE_NAME];
      }
    })
    .addCase(duplicateVariable, (state, action) => {
      const original = cloneDeep<VariableModel>(state[action.payload.name]);
      const index = Object.keys(state).length;
      const newName = `copy_of_${original.name}`;
      state[newName] = cloneDeep(variableAdapters.get(action.payload.type).initialState);
      state[newName] = original;
      state[newName].name = newName;
      state[newName].index = index;
    })
    .addCase(changeVariableOrder, (state, action) => {
      const variables = Object.values(state).map(s => s);
      const fromVariable = variables.find(v => v.index === action.payload.data.fromIndex);
      const toVariable = variables.find(v => v.index === action.payload.data.toIndex);

      if (fromVariable) {
        state[fromVariable.name].index = action.payload.data.toIndex;
      }

      if (toVariable) {
        state[toVariable.name].index = action.payload.data.fromIndex;
      }
    })
    .addCase(storeNewVariable, (state, action) => {
      const name = action.payload.name;
      const emptyVariable: VariableModel = cloneDeep<VariableModel>(state[NEW_VARIABLE_NAME]);
      state[name] = cloneDeep(variableAdapters.get(action.payload.type).initialState);
      state[name] = emptyVariable;
    })
    .addCase(changeToEditorEditMode, (state, action) => {
      if (action.payload.name === NEW_VARIABLE_NAME) {
        state[NEW_VARIABLE_NAME] = cloneDeep(variableAdapters.get('query').initialState);
        state[NEW_VARIABLE_NAME].index = Object.values(state).length - 1;
      }
    })
    .addCase(changeVariableType, (state, action) => {
      const { name: originalName } = action.payload;
      const initialState = cloneDeep(variableAdapters.get(action.payload.data).initialState);
      const { label, name, index } = state[originalName];

      state[originalName] = {
        ...initialState,
        variable: {
          ...initialState.variable,
          label,
          name,
          index,
        },
      };
    })
    .addCase(changeVariableNameSucceeded, (state, action) => {
      const instanceState = getInstanceState(state, action.payload.name);
      instanceState.name = action.payload.data;
    })
    .addCase(setCurrentVariableValue, (state, action) => {
      const instanceState = getInstanceState<VariableWithOptions>(state, action.payload.name);
      const current = { ...action.payload.data };

      if (Array.isArray(current.text) && current.text.length > 0) {
        current.text = current.text.join(' + ');
      } else if (Array.isArray(current.value) && current.value[0] !== ALL_VARIABLE_VALUE) {
        current.text = current.value.join(' + ');
      }

      instanceState.current = current;
      instanceState.options = instanceState.options.map(option => {
        let selected = false;
        if (Array.isArray(current.value)) {
          for (let index = 0; index < current.value.length; index++) {
            const value = current.value[index];
            if (option.value === value) {
              selected = true;
              break;
            }
          }
        } else if (option.value === current.value) {
          selected = true;
        }
        option.selected = selected;
        return option;
      });
    })
    .addCase(changeVariableProp, (state, action) => {
      const instanceState = getInstanceState(state, action.payload.name);
      (instanceState as Record<string, any>)[action.payload.data.propName] = action.payload.data.propValue;
    })
);
