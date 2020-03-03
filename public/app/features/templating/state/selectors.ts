import { cloneDeep } from 'lodash';

import { StoreState } from '../../../types';
import { VariableModel } from '../variable';
import { getState } from '../../../store/store';
import { NEW_VARIABLE_NAME } from './types';

export const getVariable = <T extends VariableModel = VariableModel>(
  name: string,
  state: StoreState = getState()
): T => {
  if (!state.templating.variables[name]) {
    throw new Error(`Couldn't find variable with name:${name}`);
  }

  return state.templating.variables[name] as T;
};

export const getVariableWithName = <T extends VariableModel = VariableModel>(name: string): T => {
  return getState().templating.variables[name] as T;
};

export const getVariables = (state: StoreState = getState()): VariableModel[] => {
  return Object.values(state.templating.variables).filter(variable => variable.name !== NEW_VARIABLE_NAME);
};

export const getVariableClones = (state: StoreState = getState(), includeNewVariable = false): VariableModel[] => {
  const variables = Object.values(state.templating.variables)
    .filter(variable => (includeNewVariable ? true : variable.name !== NEW_VARIABLE_NAME))
    .map(variable => cloneDeep(variable));
  return variables.sort((s1, s2) => s1.index! - s2.index!);
};
