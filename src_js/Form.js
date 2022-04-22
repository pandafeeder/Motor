import Konva from 'konva';

class Form {
  constructor(node) {
    this.node = node
    let canvas = document.getElementById("container")
    let width = canvas.getBoundingClientRect().width;
    let height = canvas.getBoundingClientRect().height;
    this.width = 500
    this.height = 400
    this.center = {x: width/2, y: height/2}
    this.draw()
    this.addEventListener()
  }
  draw() {
    this.Group = new Konva.Group({
      x: this.center.x-this.width/2,
      y: this.center.y-this.height/2,
      width: this.width,
      height: this.height,
      draggable: true,
    })
    this.Rect = new Konva.Rect({
      width: this.width,
      height: this.height,
      fill: 'red',
      shadowColor: 'black',
      shadowOffset: {x: 10, y:10},
      shadowBlur: 10,
      shadowOpacity: 0.5,
      cornerRadius: 10,
    })
    this.Group.add(this.Rect)
    this.Text = new Konva.Text({
      width: this.width,
      height: this.height,
      fontSize: 30,
      text: `
        ${this.node.sourcefile}
        inputs: ${this.node.inputs}
        output: ${this.node.outputs}
      `
    })
    this.Group.add(this.Text)
  }
  addEventListener() {
    this.Group.on('click', () => {
      this.Group.destroy()
    })
  }
}

export default Form
