;(function () {
  var drawLine = false

  var canvas = document.getElementById('canvas')
  var finalPos = {x: 0, y: 0}
  var startPos = {x: 0, y: 0}
  var ctx = canvas.getContext('2d')
  canvas.width = 193
  canvas.height = 216

  var canvasOffset = {top: 8, left: 8}

  function line (cnvs) {
    cnvs.beginPath()
    cnvs.moveTo(startPos.x + 2, startPos.y)
    cnvs.lineTo(finalPos.x, finalPos.y)
    cnvs.stroke()
  }

  function clearCanvas () {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
  }

  function getMousePos (canvas, evt) {
    var rect = canvas.getBoundingClientRect()
    return {
      x: evt.clientX - rect.left,
      y: evt.clientY - rect.top
    }
  }

  function drawOnFinalPos () {}
  if (drawLine === true) {
    finalPos = {x: (e.pageX - canvasOffset.left), y: e.pageY - canvasOffset.top}

    clearCanvas()
    line(ctx)
  }

  canvas.addEventListener('mousemove', function (e) {
    if (drawLine === true) {
      finalPos = {x: (e.pageX - canvasOffset.left), y: e.pageY - canvasOffset.top}

      clearCanvas()
      line(ctx)
    }
  }, false)

  canvas.addEventListener('mousedown', function (e) {
    drawLine = true
    ctx.strokeStyle = 'blue'
    ctx.lineWidth = 1
    ctx.lineCap = 'square'
    ctx.beginPath()
    startPos = { x: e.pageX - canvasOffset.left - 1, y: e.pageY - canvasOffset.top}
  })

  window.addEventListener('mouseup', function (e) {
    console.log(startPos, finalPos)
    clearCanvas()
    // Replace with var that is second canvas
    line(ctx)
    finalPos = {x: 0, y: 0}
    startPos = {x: 0, y: 0}
    drawLine = false
  })
})()
