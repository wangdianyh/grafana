import { VariableModel, VariableOption, VariableType } from '../variable';
import { VariableEditorProps, VariablePickerProps, VariableState } from '../state/types';
import { UrlQueryValue } from '@grafana/runtime';
import { ComponentType } from 'react';
import { Reducer } from 'redux';
import { TemplatingState } from '../state';

export interface VariableAdapter<Model extends VariableModel> {
  description: string;
  initialState: VariableState;
  dependsOn: (variable: Model, variableToTest: Model) => boolean;
  setValue: (variable: Model, option: VariableOption) => Promise<void>;
  setValueFromUrl: (variable: Model, urlValue: UrlQueryValue) => Promise<void>;
  updateOptions: (variable: Model, searchFilter?: string, notifyAngular?: boolean) => Promise<void>;
  getSaveModel: (variable: Model) => Partial<Model>;
  getValueForUrl: (variable: Model) => string | string[];
  picker: ComponentType<VariablePickerProps>;
  editor: ComponentType<VariableEditorProps>;
  reducer: Reducer<TemplatingState>;
}

export interface VariableAdapters {
  contains: (type: VariableType) => boolean;
  get: (type: VariableType) => VariableAdapter<any>;
  set: (type: VariableType, adapter: VariableAdapter<any>) => void;
}
