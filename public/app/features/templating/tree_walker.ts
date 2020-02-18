import { variableAdapters } from './adapters';
import { VariableActions, VariableModel } from './variable';
import { VariableAdapters } from './adapters/types';

export const treeWalkerFactory = (adapters?: VariableAdapters) => {
  return (variables: Array<VariableModel & VariableActions>) => walk(adapters ?? variableAdapters, variables);
};

type VariableTreeNode = {
  variable?: VariableModel & VariableActions;
  children: VariableTreeNode[];
};

const walk = (adapters: VariableAdapters, variables: Array<VariableModel & VariableActions>): VariableTreeNode => {
  const nodes: Record<string, VariableTreeNode> = variables.reduce((all: Record<string, VariableTreeNode>, current) => {
    all[current.name] = {
      variable: current,
      children: [],
    };
    return all;
  }, {});

  for (const v1 of variables) {
    for (const v2 of variables) {
      if (v1 === v2) {
        continue;
      }

      if (adapters.contains(v1.type)) {
        if (adapters.get(v1.type).dependsOn(v1, v2)) {
          const n1 = nodes[v1.name];
          const n2 = nodes[v2.name];
          n2.children.push(n1);
        }
        continue;
      }

      if (v1.dependsOn(v2)) {
        const n1 = nodes[v1.name];
        const n2 = nodes[v2.name];
        n2.children.push(n1);
      }
    }
  }

  const tree: VariableTreeNode = {
    children: [],
  };

  const values = Object.values(nodes);
  const leafs = values.filter(n => n.children.length === 0);

  for (const leaf of leafs) {
    findParents(leaf, values, tree);
  }

  return tree;
};

function findParents(leaf: VariableTreeNode, nodes: VariableTreeNode[], tree: VariableTreeNode) {
  const parents = nodes.filter(n => !!n.children.find(c => c.variable.name === leaf.variable.name));

  if (parents.length === 0) {
    tree.children.push(leaf);
    return;
  }

  for (const parent of parents) {
    findParents(parent, nodes, tree);
  }
}
