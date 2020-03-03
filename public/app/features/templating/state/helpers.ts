import { NEW_VARIABLE_NAME } from './types';
import { QueryVariableModel, VariableHide, VariableModel } from '../variable';
import { initialQueryVariableModelState } from '../query/reducer';
import { VariablesState } from './variablesReducer';

export const getVariableState = (
  noOfVariables: number,
  inEditorIndex = -1,
  includeEmpty = false
): Record<string, VariableModel> => {
  const variables: Record<string, VariableModel> = {};

  for (let index = 0; index < noOfVariables; index++) {
    variables[`Name-${index}`] = {
      type: 'query',
      name: `Name-${index}`,
      hide: VariableHide.dontHide,
      index,
      label: `Label-${index}`,
      skipUrlSync: false,
    };
  }

  if (includeEmpty) {
    variables[NEW_VARIABLE_NAME] = {
      type: 'query',
      name: `Name-${NEW_VARIABLE_NAME}`,
      hide: VariableHide.dontHide,
      index: noOfVariables,
      label: `Label-${NEW_VARIABLE_NAME}`,
      skipUrlSync: false,
    };
  }

  return variables;
};

export const getVariableTestContext = (variableOverrides: Partial<QueryVariableModel> = {}) => {
  const defaultVariable = {
    ...initialQueryVariableModelState,
    index: 0,
    name: 'Name-0',
  };
  const variable = { ...defaultVariable, ...variableOverrides };
  const initialState: VariablesState = {
    'Name-0': variable,
  };

  return { initialState };
};
