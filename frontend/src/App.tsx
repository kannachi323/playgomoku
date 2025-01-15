import './App.css'
import { Conn, Player } from './types'

import { useState, useEffect } from 'react';

function App() {
  const [messages, setMessages] = useState<string[]>([])
  const [msg, setMsg] = useState<string>('')

  const [socket, setSocket] = useState<WebSocket | null>(null)

  useEffect(() => {
    const socket : WebSocket = new WebSocket('ws://localhost:3000/ws')
    socket.onopen = () => {
      const p1 : Player = {
        id: ';adlskfj',
        name: 'Bot',
        color: 1
      }
      const data : Conn = {
        roomID: '1',
        player: p1,
        type: 'join',
      }

      if (socket) {
        socket.send(JSON.stringify(data))
        console.log("attempted to connect")
      }
    }

    socket.onmessage = (event) => {
      console.log("server: ", event.data)
    }

    setSocket(socket)
  }, [])
  

  async function sendMessage(content: string) {
    setMessages([...messages, content])
    if (socket) {
      console.log("sending message from client: ", content)
      socket.send(content)
      //this will call socket.onmessage
    }
  }
  
  return (
    <>
      <input
        onChange={(e) => {
          const content = e.target.value
          console.log(content)
          setMsg(content)
        }}
        type="text"
      >

      </input>


      <button
        onClick={() => sendMessage(msg)}
      >
        Send Message to server
      </button>

    

   
    </>
  )

 
}

export default App
