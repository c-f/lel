import dagre from "dagre";

// https://codesandbox.io/s/yv305rln5z
// stolen from the author of GGeditor

// https://github.com/dagrejs/graphlib/wiki/API-Reference
// export/import graph
//
export const HandleLayout = data => {
  const graph = new dagre.graphlib.Graph({})
    .setGraph({
      marginx: "20",
      marginy: "20",
      rankdir: "LR",
      //minlen: "550",
      ranker: "longest-path",
      labeloffset: "50",
      nodesep: "70",
      edgesep: "70",
      ranksep: "70",
      align: "UL"
    })
    .setDefaultEdgeLabel(() => {
      return {};
    });

  data.nodes.forEach(node => {
    const size = node.size.split("*");
    const width = Number(size[0]);
    const height = Number(size[1]);
    graph.setNode(node.id, { width, height });
  });
  data.edges.forEach(edge => {
    graph.setEdge(edge.source, edge.target);
  });
  dagre.layout(graph);
  const nextNodes = data.nodes.map(node => {
    const graphNode = graph.node(node.id);
    return {
      ...node,
      x: graphNode.x,
      y: graphNode.y,
      labelOffsetY: graphNode.labelOffsetY
    };
  });

  return { nodes: nextNodes, edges: data.edges };
};

// FilterNodes
export const FilterNodes = (data, search) => {
  if (search === undefined || search == "") {
    console.log(search);
    console.log("Display all");
    return data;
  }
  let blacklist = [];
  const nextNodes = data.nodes.filter(node => {
    let ok = node.tags.includes(search);
    if (!ok) {
      blacklist.push(node.id);
      return false;
    } else {
      return true;
    }
  });
  const edges = data.edges.filter(edge => {
    let bad =
      blacklist.includes(edge.source) || blacklist.includes(edge.target);
    if (bad) {
      return false;
    } else {
      return true;
    }
  });
  return { nodes: nextNodes, edges: edges };
};
