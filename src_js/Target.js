import Konva from 'konva';
import Form from "./Form.js"

// create both Rect and Text
// and event hanlder
class Target {
  constructor(config_obj) {
    this.stage = config_obj.stage
    this.layer = config_obj.layer
    this.node = config_obj.node
    this.x = config_obj.x
    this.y = config_obj.y
    this.width = config_obj.width
    this.height = config_obj.height
    this.text = config_obj.text
    this.clicked = false
    //this.createRect()
    //this.createText()
    this.draw()
    this.getRelatedNodes()
    this.addEventHandler()
  }
  draw() {
    this.Group = new Konva.Group({
      x: this.x,
      y: this.y,
      width: this.width,
      height: this.height,
      id: this.node.sourcefile,
      draggable: true,
    })
    this.Rect = new Konva.Rect({
      width: this.width,
      height: this.height,
      fill: 'green',
      stroke: 'black',
      strokWidth: 4,
      //shadowColor: 'black',
      //shadowOffset: {x: 10, y:10},
      //shadowBlur: 10,
      //shadowOpacity: 0.5,
      cornerRadius: 5,
      name: this.node.sourcefile,
    })
    this.Group.add(this.Rect)
    this.Text = new Konva.Text({
      text: this.text,
      width: this.width,
      height: this.height,
      fontSize: 10,
      verticalAlign: 'middle',
      align: 'center',
    })
    this.Group.add(this.Text)
    this.zIndex = this.Group.getZIndex()
  }
  createRect() {
    this.Rect = new Konva.Rect({
      //x: this.x,
      //y: this.y,
      width: this.width,
      height: this.height,
      fill: 'green',
      stroke: 'black',
      strokWidth: 4,
      //shadowColor: 'black',
      //shadowOffset: {x: 10, y:10},
      //shadowBlur: 10,
      //shadowOpacity: 0.5,
      cornerRadius: 5,
    })
  }
  createText() {
    this.Text = new Konva.Text({
      //x: this.x,
      //y: this.y,
      text: this.text,
      width: this.width,
      height: this.height,
      fontSize: 30,
      verticalAlign: 'middle',
      align: 'center',
    })
  }
  getRelatedNodes() {
    let parents_id = this.node.parents
    let children_id = this.node.children
    let related_id = parents_id.concat(children_id)
    this.relatedNodes = related_id
  }
  addEventHandler() {
    this.Group.on('mouseover', (evt) => {
      console.log(`mouseover on ${this.node.sourcefile}`)
      this.Rect.fill('lightgreen')
      //evt.cancelBubble = true;
      window.g = this.Group
      if (this.relatedNodes.length > 0) {
        this.relatedNodes.forEach((n, id) => {
          //this.stage.find(n)[0].Rect.fill('lightgreen');
          console.log(n)
          this.stage.find(`.${n}`)[0].fill('lightgreen')
        })
      }

      //let parent_node_sourcefile = this.stage.find(this.node.parents)
      //let children_node_sourcefile = this.stage.find(this.node.children)
    })
    this.Group.on('mouseout', (evt) => {
      console.log(`mouseout on ${this.node.sourcefile}`)
      this.Rect.fill('green')
      if (this.relatedNodes.length > 0) {
        this.relatedNodes.forEach((n, id) => {
          //this.stage.find(n)[0].Rect.fill('lightgreen');
          console.log(n)
          this.stage.find(`.${n}`)[0].fill('green')
        })
      }
      //evt.cancelBubble = true;
    })
    this.Group.on('click', (evt) => {
      console.log(`click on ${this.node.sourcefile}`)
      let form = new Form(this.node)
      this.layer.add(form.Group)
      //if (!this.clicked) {
      //  this.Group.moveToTop()
      //  this.Rect.width(500)
      //  this.Rect.height(300)
      //  this.clicked = ! this.clicked
      //} else {
      //  //this.Group.setZIndex(this.zIndex)
      //  this.Group.moveToTop()
      //  this.Rect.width(this.width)
      //  this.Rect.height(this.height)
      //  this.clicked = ! this.clicked
      //}
    })
  }
}

export default Target
