import { VariableType } from '../variable';
import { VariableAdapter, VariableAdapters } from './types';

const allVariableAdapters: Record<VariableType, VariableAdapter<any> | null> = {
  query: null,
  textbox: null,
  constant: null,
  datasource: null,
  custom: null,
  interval: null,
  adhoc: null,
};

export const variableAdapters: VariableAdapters = {
  contains: (type: VariableType): boolean => !!allVariableAdapters[type],
  get: (type: VariableType): VariableAdapter<any> => {
    if (allVariableAdapters[type] !== null) {
      // @ts-ignore
      // Suppressing strict null check in this case we know that this is an instance otherwise we throw
      // Type 'VariableAdapter<any, any> | null' is not assignable to type 'VariableAdapter<any, any>'.
      // Type 'null' is not assignable to type 'VariableAdapter<any, any>'.
      return allVariableAdapters[type];
    }

    throw new Error(`There is no adapter for type:${type}`);
  },
  set: (type: VariableType, adapter: VariableAdapter<any>): void => {
    allVariableAdapters[type] = adapter;
  },
};
