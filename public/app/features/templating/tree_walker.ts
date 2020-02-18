import { variableAdapters } from './adapters';
import { VariableActions, VariableModel } from './variable';
import { VariableAdapters } from './adapters/types';

export const treeWalkerFactory = (adapters?: VariableAdapters) => {
  return (variables: Array<VariableModel & VariableActions>) => walk(adapters ?? variableAdapters, variables);
};

type Node = {
  name: string;
  children: Node[];
};

const walk = (adapters: VariableAdapters, variables: Array<VariableModel & VariableActions>): Node => {
  const nodes: Record<string, Node> = variables.reduce((all: Record<string, Node>, current) => {
    all[current.name] = {
      name: current.name,
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
          continue;
        }
      }

      if (v1.dependsOn(v2)) {
        const n1 = nodes[v1.name];
        const n2 = nodes[v2.name];
        n2.children.push(n1);
      }
    }
  }

  const tree: Node = {
    name: 'root',
    children: [],
  };

  const values = Object.values(nodes);
  const leafs = values.filter(n => n.children.length === 0);

  for (const leaf of leafs) {
    findParents(leaf, values, tree);
  }

  return tree;
};

function findParents(leaf: Node, nodes: Node[], tree: Node) {
  const parents = nodes.filter(n => !!n.children.find(c => c.name === leaf.name));

  if (parents.length === 0) {
    tree.children.push(leaf);
    return;
  }

  for (const parent of parents) {
    findParents(parent, nodes, tree);
  }
}
