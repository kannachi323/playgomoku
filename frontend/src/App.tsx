import './App.css';

import { useState } from 'react';

function App() {
  const [messages, setMessages] = useState<string[]>([]);
  const [msg, setMsg] = useState<string>('');
  const [socket, setSocket] = useState<WebSocket | null>(null);

  function joinRoom() {
    if (socket) {
      console.log('Already connected to the WebSocket');
      return;
    }

    const pid = "def";
    const clr = "2"

    const newSocket: WebSocket = new WebSocket(`ws://localhost:3000/ws?pid=${pid}&clr=${clr}`);

    newSocket.onmessage = (event) => {
      console.log('server: ', event.data);
    };

    newSocket.onclose = () => {
      console.log('WebSocket connection closed');
    };

    setSocket(newSocket);
  };

  const sendMessage = (content: string) => {
    if (socket) {
      console.log('Sending message from client: ', content);
      socket.send(content);
      setMessages([...messages, content]);
    } else {
      console.error('WebSocket is not connected');
    }
  };

  return (
    <>
      <button onClick={joinRoom}>Join Room</button>

      <input
        onChange={(e) => {
          const content = e.target.value;
          setMsg(content);
        }}
        type="text"
        value={msg}
      />

      <button onClick={() => sendMessage(msg)}>Send Message to server</button>
    </>
  );
}

export default App;
