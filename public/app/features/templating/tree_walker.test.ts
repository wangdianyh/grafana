import { QueryVariableModel, VariableActions, VariableModel, VariableType } from './variable';
import { treeWalkerFactory } from './tree_walker';
import { VariableAdapter, VariableAdapters } from './adapters/types';

describe('TreeWalker', () => {
  const queryVariableAdapter: VariableAdapter<QueryVariableModel> = {
    dependsOn: (variable, variableToTest) => {
      if (variable.name === 'query2') {
        return true;
      }
      return false;
    },
  } as VariableAdapter<QueryVariableModel>;
  const adapters: VariableAdapters = {
    get: type => queryVariableAdapter,
    set: type => undefined,
    contains: type => type === 'query',
  };
  const walking = treeWalkerFactory(adapters);

  describe('when walking tree', () => {
    // lista med variables ifrån redux + angularjs
    const query1 = createReduxVariable('query1', 'query');
    const custom1 = createAngularVariable('custom1', 'custom', (variable: any) => variable.name === query1.name);
    const query2 = createReduxVariable('query2', 'query');

    it('should return variables in dependency order', () => {
      const tree = walking([custom1, query2, query1]);
      expect(tree).toEqual({
        children: [
          {
            variable: query1,
            children: [
              { variable: custom1, children: [{ variable: query2, children: [] }] },
              { variable: query2, children: [] },
            ],
          },
        ],
      });
    });
  });
});

function createAngularVariable(
  name: string,
  type: VariableType,
  dependsOn?: Function
): VariableModel & VariableActions {
  return ({
    name,
    type,
    dependsOn: dependsOn || ((variable: any) => false),
  } as unknown) as VariableModel & VariableActions;
}

function createReduxVariable(name: string, type: VariableType): VariableModel & VariableActions {
  return ({
    name,
    type,
  } as unknown) as VariableModel & VariableActions;
}
