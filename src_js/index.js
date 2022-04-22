import Target from './Target.js'
import Tree from "./Tree.js"

function createConfigObj() {
}

function main() {
  let width = window.innerWidth
  let height = window.innerHeight

  let stage = new Konva.Stage({
    container: 'container',
    width: width,
    height: height,
  })

  let nodes = Tree.nodes
  let tree_depth = Tree.depth
  let nodes_in_level = {}
  nodes.forEach((node, id) => {
    if (nodes_in_level[node.level]) {
      nodes_in_level[node.level] += 1
    } else {
      nodes_in_level[node.level] = 1
    }
  })
  let tree_width = Math.max(...Object.values(nodes_in_level))

  let node_height = height / tree_depth
  node_height = node_height >= 100 ? 100 : node_height
  let node_width  = width  / tree_width
  node_width  = node_width  >= 200 ? 200 : node_width
  let start_y = 10

  console.log(width)
  console.log(node_height)
  console.log(node_width)

  let shapes = []
  window.shapes = shapes
  let layer = new Konva.Layer()

  let re = /\S\/([^\/]+).sh$/
  for (let i = 0; i < tree_depth; i++) {
    let nodes_in_cur_level = nodes.filter( node =>  node.level == i )
    let cur_width = nodes_in_cur_level.length
    let start_x   = (width / 2) - (cur_width / 2)*node_width
    console.log(start_x)
    nodes_in_cur_level.forEach((node, id) => {
      let configObj = {
        stage: stage,
        layer: layer,
        node: node,
        x: start_x + id * node_width,
        y: start_y + i * node_height,
        width: node_width,
        height: node_height,
        text: re.exec(node.sourcefile)[1]
      }
      let target = new Target(configObj)
      layer.add(target.Group)
    })
  }
  stage.add(layer)
}

main();
