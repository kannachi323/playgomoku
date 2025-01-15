import './App.css';
import { Conn } from './types';

import { useState } from 'react';

function App() {
  const [messages, setMessages] = useState<string[]>([]);
  const [msg, setMsg] = useState<string>('');
  const [socket, setSocket] = useState<WebSocket | null>(null);

  const joinWebSocket = () => {
    if (socket) {
      console.log('Already connected to the WebSocket');
      return;
    }

    const newSocket: WebSocket = new WebSocket('ws://localhost:3000/ws');
    
    newSocket.onopen = () => {
      const data: Conn = {
        type: 'join',
        roomID: '1',
        player: {
          id: 'abc',
          name: 'Matt',
          color: 1,
        }
        
      };

      newSocket.send(JSON.stringify(data));
      console.log('Attempted to connect');
    };

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
      <button onClick={joinWebSocket}>Join WebSocket</button>

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
