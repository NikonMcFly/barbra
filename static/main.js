;(function () {
  var drawLine = false

  var canvas = document.getElementById('canvas')
  var endPos = {x: 0, y: 0}
  var startPos = {x: 0, y: 0}
  var ctx = canvas.getContext('2d')
  canvas.width = 193
  canvas.height = 216

  var canvasOffset = {top: 8, left: 8}

  function line (cnvs) {
    cnvs.beginPath()
    // The +2 in this function is to fix a bug where
    // the line starts before where the mose is positioned.
    // Depending on the zoom and the size of the picture
    // this may need to bo adujusted
    cnvs.moveTo(startPos.x + 2, startPos.y)
    cnvs.lineTo(endPos.x, endPos.y)
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

  function drawOnEndPos () {}
  if (drawLine === true) {
    endPos = {x: (e.pageX - canvasOffset.left), y: e.pageY - canvasOffset.top}

    clearCanvas()
    line(ctx)
  }

  canvas.addEventListener('mousemove', function (e) {
    if (drawLine === true) {
      endPos = {x: (e.pageX - canvasOffset.left), y: e.pageY - canvasOffset.top}

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

    // The -1 in this function is to fix a bug where the line starts before where the mose is positioned
    startPos = { x: e.pageX - canvasOffset.left - 1, y: e.pageY - canvasOffset.top}
  })

  window.addEventListener('mouseup', function (e) {
    console.log(startPos, endPos)
    clearCanvas()
    // Replace with var that is second canvas
    line(ctx)
    endPos = {x: 0, y: 0}
    startPos = {x: 0, y: 0}
    drawLine = false
  })

  document.getElementById('resize').addEventListener('click', function () {
    var post = { 
      line: {
        start: startPos, 
        end: endPos,
      },
      length: document.getElementById("length").value
    }
    console.log('Post: ', post)
    ajax.post('/resize', JSON.stringify(post), function (data) {console.log(data)})
  }, false)

  ajax.get('/healthz', null, function (data) { console.log(data)})
})()
