import { useState } from 'react'
import './App.css'

async function fetchPong(): Promise<string> {
  const response = await fetch('http://localhost:9000/ping')
  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status}`);
  }

  const data: string = await response.text();
  return data;
}

function App() {
  const [pong, setPong] = useState('empty')

  const getPong = async () => {
    const data = await fetchPong().then((data) => {
      return data
    })
    .catch((error) => {
      console.error('Error:', error);
      return 'Empty'
    });
    console.log('Before setPong ' + data)
    setPong(data)
  }

  return (
    <>
      <h1>Test CORS Gorilla Mux</h1>
      <div className="card">
        <button onClick={() => getPong()}>
          Pong is {pong}
        </button>
      </div>
    </>
  )
}

export default App
