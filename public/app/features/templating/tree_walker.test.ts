import { VariableType, VariableModel, VariableActions } from './variable';
import { treeWalkerFactory } from './tree_walker';

describe('TreeWalker', () => {
  const walking = treeWalkerFactory();

  describe('when walking tree', () => {
    // lista med variables ifrån redux + angularjs
    const query1 = createReduxVariable('query1', 'query');
    const custom1 = createAngularVariable('custom1', 'custom', (variable: any) => variable.name === query1.name);

    it('should return variables in dependency order', () => {
      const tree = walking([query1, custom1]);
      expect(tree).toEqual([]);
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
