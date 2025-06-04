
import { useState, useEffect } from 'react';

import { GameContext } from './GameContext'
import { ServerMessage, GameState } from '../types';
import { createConnection } from '../utils/connection';
import { createGame } from '../utils/game';


export const GameProvider = ({ children }: { children: React.ReactNode }) => {
  useEffect(() => {
    //TODO: check auth here
  
  }, []);

  const [gameState, setGameState] = useState<GameState>(createGame());
  
  const conn = createConnection(gameState.players.p2, update);

  function update(payload: ServerMessage) {
    if (payload.type === "update") {
      setGameState(payload.data.gameState);
    }
  }

  return (
    <GameContext.Provider value={{ gameState, update, conn }}>
      {children}
    </GameContext.Provider>
  );
};

