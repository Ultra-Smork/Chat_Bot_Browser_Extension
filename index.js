fetch('http://localhost:8080/chat', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    "Access-Control-Allow-Origin": "*"
  },
  body: JSON.stringify({
    title: 'My new post',
    body: 'This is the body of my new post.',
  }),
})
.then(response => {
  if (response.ok) {
    console.log("SENDED")
  } else {
    // The request failed.
  }
})
